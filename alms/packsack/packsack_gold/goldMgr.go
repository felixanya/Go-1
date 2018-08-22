package packsack_gold

import (
	"fmt"
	"steve/alms/packsack/packsack_utils"

	"sync"

	"github.com/Sirupsen/logrus"
)

var goldMgr GoldMgr

//GetGoldMgr 获取金币管理
func GetGoldMgr() *GoldMgr {
	return &goldMgr
}

//GoldMgr 金币管理
type GoldMgr struct {
	userList sync.Map                 // 用户列表
	muLock   map[uint64]*sync.RWMutex // 用户锁，一个用户一个锁
}

func init() {
	goldMgr.muLock = make(map[uint64]*sync.RWMutex)
}

//GetMutex 读写
func (gm *GoldMgr) GetMutex(uid uint64) *sync.RWMutex {
	if mu, ok := gm.muLock[uid]; ok {
		return mu
	}
	n := new(sync.RWMutex)
	gm.muLock[uid] = n
	return n
}

// 获取用户
func (gm *GoldMgr) getUser(uid uint64) (*userGold, error) {
	if uid == 0 {
		return nil, nil
	}
	u, ok := gm.userList.Load(uid)
	if !ok { //不存在，从redsi or db 获取
		return gm.getUserFromCacheOrDB(uid)
	}
	return u.(*userGold), nil
}

// 从Redis或者DB获取用户
func (gm *GoldMgr) getUserFromCacheOrDB(uid uint64) (*userGold, error) {
	m, err := packsack_utils.GetPlayerPacksackGold(uid)
	if err == nil {
		return gm.newUser(uid, m), nil
	}

	m, err = packsack_utils.GetGoldFromDB(uid)
	if err != nil {
		return nil, fmt.Errorf("load from db failed err(%v)", err)
	}
	// 从DB获取到后，马上缓存到Redis
	err = packsack_utils.SaveGoldToRedis(uid, m)
	if err != nil {
		// 记录redis写入失败
		logrus.Errorln("save redis error")
	}
	return gm.newUser(uid, m), nil
}

// 新建用户
func (gm *GoldMgr) newUser(uid uint64, m int64) *userGold {
	n := newuserGold(uid, m)
	gm.userList.Store(uid, n)
	return n
}

//AddGold 加金币
func (gm *GoldMgr) AddGold(uid uint64, value int64) (int64, error) {
	// 1. 先获取玩家当前金币值, GetGold()
	// 2. 在内存中对玩家金币进行加减
	// 3. 将变化后的值写到redis和DB
	before := int64(0)
	after := int64(0)

	entry := logrus.WithFields(logrus.Fields{
		"opr":     "add_gold",
		"uid":     uid,
		"before":  before,
		"changed": value,
		"after":   after,
	})
	// 按用户ID进行加锁,一个用户一个锁
	mu := gm.GetMutex(uid)
	mu.Lock()
	defer mu.Unlock()

	u, err := gm.getUser(uid)
	if u == nil {
		entry.Errorln("get user error")
		_ = err
		return 0, fmt.Errorf("no user")
	}
	// 加金币前，玩家当前金币值
	before, err = u.Get()

	// 加金币后，玩家当前金币值
	after, err = u.Add(value)
	if err != nil {
		entry.Errorln("add opr error goldValue%(v)", value)
		return 0, err
	}

	entry = logrus.WithFields(logrus.Fields{
		"opr":     "add_gold",
		"uid":     uid,
		"before":  before,
		"changed": value,
		"after":   after,
	})
	// 交易记录写到日志
	entry.Infoln("add succeed")

	// 交易记录写到redis
	// 交易记录写到DB
	err = gm.saveUserToCacheAndDB(entry, u, after)
	if err != nil {
		entry.Errorln("save cacheordb error")
	}
	return after, nil
}

//GetGold 获取金币
func (gm *GoldMgr) GetGold(uid uint64) (int64, error) {
	// 1.先在内存中查找玩家是否存在。
	// 2.不存在，从Redis获取玩家金币.
	// 3.不存在，从DB获取玩家金币.

	// 按用户ID进行加锁,一个用户一个锁
	mu := gm.GetMutex(uid)
	mu.RLock()
	defer mu.RUnlock()

	u, err := gm.getUser(uid)
	if err != nil {
		logrus.WithError(err).Debugln("get user err")
	}
	if u == nil {
		return 0, fmt.Errorf("no user")
	}
	// 获取玩家金币
	g, err := u.Get()
	if err != nil {
		return 0, err
	}

	return g, nil
}

// 保存玩家变化到Redis和DB
func (gm *GoldMgr) saveUserToCacheAndDB(entry *logrus.Entry, u *userGold, changeValue int64) error {
	// 暂时先保存到Redis
	err := packsack_utils.SaveGoldToRedis(u.uid, changeValue)
	if err != nil {
		// 记录redis写入失败
		entry.Errorln("save redis error")
	}
	// 后续再保存到DB
	err = packsack_utils.SaveGoldToDB(u.uid, changeValue)
	if err != nil {
		// 记录DB写入失败
		entry.Errorln("save db error")
	}
	return nil
}
