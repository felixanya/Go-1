package logic

import (
	"github.com/Sirupsen/logrus"
	"steve/propserver/data"
	"steve/propserver/define"
	"sync"
	"steve/external/configclient"
	"encoding/json"
	"time"
)

/*
  功能： 道具管理： 加减玩家道具，获取玩家道具,交易序列号去重. 支持redis，db同步存储。交易流水日志对账等.
  作者： SkyWang
  日期： 2018-8-13
*/

var myLogic PropsMgr

func GetMyLogic() *PropsMgr {
	return &myLogic
}

type PropsMgr struct {
	userList map[uint64]*userProps          // 用户列表
	muLock   map[uint64]*sync.Mutex 		// 用户锁，一个用户一个锁

	propsList map[uint64]*propsInfo
}

func (gm *PropsMgr) Init() error {

	gm.userList = make(map[uint64]*userProps)
	gm.muLock = make(map[uint64]*sync.Mutex)

	err := gm.getPropsListFromDB()
	if err != nil {
		logrus.Errorf("get getPropsListFromDB err:", err)
	}

	go gm.runMyTask()

	return nil
}

// 启动道具列表变化检测协程
func (gm *PropsMgr)  runMyTask() error{

	// 1分钟更新一次邮件列表
	for {
		time.Sleep(time.Minute)

		gm.getPropsListFromDB()

		// 清理过期用户和锁
		gm.clearExpiredUser()
	}

	return nil
}

// 每日半夜3-4点，从内存中清理过期玩家信息和玩家的锁
var thisDay = 0
// 清理过期邮件开始点数
var clearBeginHour = 3
// 清理过期邮件结束点数
var clearEndHour = 4
// 7天过期,7天未访问的User，从内存清理出去
var clearTimeOut = int64(3600 * 24 * 7)
func (gm *PropsMgr) clearExpiredUser() {

	now := time.Now()
	// 每日只执行1次
	if thisDay == now.YearDay() {
		return
	}
	if now.Hour() < clearBeginHour || now.Hour() >= clearEndHour {
		return
	}
	thisDay = now.YearDay()

	tick := now.Unix()

	logrus.Infof("begin clearExpiredUser work ...")

	for k, u := range  gm.userList {
		if u.lastVisitTime == 0 {
			continue
		}
		sub :=  tick - u.lastVisitTime
		if sub < clearTimeOut {
			continue
		}

		// 清理此用户
		delete(gm.userList, k)
		// 清理此用户的锁
		delete(gm.muLock, k)

		logrus.Infof("clearExpiredUser one: uid=%d", k)
	}

	logrus.Infof("end clearExpiredUser work ...")
}


func (gm *PropsMgr) getPropsListFromDB() error {
	strJson, err := configclient.GetConfig("prop", "interactive")
	if err != nil {
		logrus.Errorf("GetPropsListFromDB from config err:", err)
		return err
	}

	return gm.parseJsonPropsList(strJson)

}

func (gm *PropsMgr) parseJsonPropsList(strJson string) error {

	jsonObject := make([]*propsInfo,0,2)
	err := json.Unmarshal([]byte(strJson), &jsonObject)
	if err != nil {
		return nil
	}
	gm.propsList = make(map[uint64]*propsInfo, 10)
	for _, one := range  jsonObject {
		gm.propsList[one.PropID] = one
	}

	return nil
}

func (gm *PropsMgr) GetMutex(uid uint64) *sync.Mutex {
	if mu, ok := gm.muLock[uid]; ok {
		return mu
	}
	n := new(sync.Mutex)
	gm.muLock[uid] = n
	return n
}

// 加玩家道具
func (gm *PropsMgr) AddUserProps(uid uint64, propList map[uint64]int64, seq string, funcId int32, channel int64, createTm int64, gameId, level int32) error {
	// 1. 先获取玩家当前金币值, GetGold()
	// 2. 在内存中对玩家金币进行加减
	// 3. 将变化后的值写到redis和DB
	before := int64(0)
	after := int64(0)

	entry := logrus.WithFields(logrus.Fields{
		"opr":        "add_props",
		"gameId":     gameId,
		"level":      level,
		"uid":        uid,
		"funcId":     funcId,
		"propList":     propList,
		"channel":    channel,
		"seq":        seq,
		"createTime": createTm,
	})

	for propId := range  propList {
		if !gm.checkPropId(propId) {
			entry.Errorln("propId error")
			return define.ErrPropId
		}
	}


	// 按用户ID进行加锁,一个用户一个锁
	mu := gm.GetMutex(uid)
	mu.Lock()
	defer mu.Unlock()

	u, err := gm.getUser(uid)
	if u == nil {
		entry.Errorln("get user error")
		_ = err
		return  define.ErrNoUser
	}
	// 设置最后访问时间
	u.lastVisitTime = time.Now().Unix()

	// 判断交易流水号是否有冲突?
	if !u.CheckSeq(seq) {
		entry.Errorf("seq is same: uid=%d, seq=%s", uid, seq)
		return  define.ErrSeqNo
	}

	// 加道具前，玩家当前道具数量
	for propId, num := range  propList {
		if num >= 0 {
			continue
		}
		before, _ = u.Get(propId)

		if before+num < 0 {
			entry.Errorf("prop num < value: uid=%d, before=%d, add=%d", uid, before, num)
			return define.ErrNoProp
		}
	}

	// 加道具后，玩家当前道具数量
	for propId, num := range  propList {
		before, _ = u.Get(propId)
		if num != 0 {
			after, _ = u.Add(propId, num)
		}


		propList[propId] = after

		entry = logrus.WithFields(logrus.Fields{
			"opr":        "add_props",
			"gameId":     gameId,
			"level":      level,
			"uid":        uid,
			"funcId":     funcId,
			"propId":     propId,
			"before":     before,
			"changed":    num,
			"after":      after,
			"channel":    channel,
			"seq":        seq,
			"createTime": createTm,
		})
		// 交易记录写到日志
		entry.Infoln("add succeed")
	}

	// 交易记录写到redis
	// 交易记录写到DB
	err = gm.saveUserToCacheAndDB(entry, uid, propList)
	if err != nil {
		entry.Errorln("saveUserToCacheAndDB error: ", err)
	}

	return nil
}

// 获取玩家道具
func (gm *PropsMgr) GetUserProps(uid uint64, propId uint64) (map[uint64]int64, error) {
	// 1.先在内存中查找玩家是否存在。
	// 2.不存在，从Redis获取玩家道具.
	// 3.不存在，从DB获取玩家道具.

	if propId != 0 {
		if !gm.checkPropId(propId) {
			logrus.Errorf("for={prop id error},uid=%d,goldType=%d", uid, propId)
			return nil, define.ErrNoProp
		}
	}

	// 按用户ID进行加锁,一个用户一个锁
	mu := gm.GetMutex(uid)
	mu.Lock()
	defer mu.Unlock()

	u, _ := gm.getUser(uid)
	if u == nil {
		return nil, define.ErrNoUser
	}
	// 设置最后访问时间
	u.lastVisitTime = time.Now().Unix()
	// 获取玩家指定道具
	g, err := u.GetList(propId)
	if err != nil {
		return nil, err
	}
	return g, nil
}

// 保存玩家变化到Redis和DB
func (gm *PropsMgr) saveUserToCacheAndDB(entry *logrus.Entry, uid uint64, list map[uint64]int64) error {

	// 暂时先保存到Redis
	err := data.SavePropsToRedis(uid, list)
	if err != nil {
		// 记录redis写入失败
		entry.Errorln("SavePropsToRedis error", err)
	}

	// 后续再保存到DB
	for propId, num := range  list {
		err = data.SavePropsToDB(uid, propId, num)
		if err != nil {
			// 记录DB写入失败
			entry.Errorln("SavePropsToDB error:", err)
		}
	}

	return nil
}

// 获取用户
func (gm *PropsMgr) getUser(uid uint64) (*userProps, error) {
	if uid == 0 {
		return nil, nil
	}
	u, ok := gm.userList[uid]
	if !ok {
		return gm.getUserFromCacheOrDB(uid)
	}
	return u, nil
}

// 新建用户
func (gm *PropsMgr) newUser(uid uint64, m map[uint64]int64) *userProps {
	n := newUserProps(uid, m)
	gm.userList[uid] = n
	return n
}

// 从Redis或者DB获取用户
func (gm *PropsMgr) getUserFromCacheOrDB(uid uint64) (*userProps, error) {
	m, err := data.LoadPropsFromRedis(uid)
	if err == nil {
		return gm.newUser(uid, m), nil
	}

	m, err = data.LoadPropsFromDB(uid)
	if err != nil {
		return nil, define.ErrLoadDB
	}
	// 从DB获取到后，马上缓存到Redis
	err = data.SavePropsToRedis(uid, m)
	if err != nil {
		// 记录redis写入失败
		logrus.Errorln("save redis error")
	}
	return gm.newUser(uid, m), nil
}

// 检测道具ID是否有效
func (gm *PropsMgr) checkPropId(propId uint64) bool {

	if _, ok := gm.propsList[propId]; ok {
		return true
	}
	// 先不判断道具是否存在
	return false
}
