package packsack

import (
	"fmt"
	"steve/common/data/redis"

	"github.com/Sirupsen/logrus"
)

//SaveGoldToRedis 保存玩家金币到Redis
func SaveGoldToRedis(uid uint64, goldList map[int16]int64) error {
	r := redis.GetRedisClient()
	key := fmtPlayerKey(uid)

	list := make(map[string]interface{}, len(goldList))
	for k, v := range goldList {
		strKey := fmt.Sprintf("%d", k)
		list[strKey] = v
	}
	cmd := r.HMSet(key, list)
	if cmd.Err() != nil {
		//logic.ErrNoUser.WithError(cmd.Err()).Errorln(errRedisOperation)
		logrus.Errorf("save gold to redis err:key=%s,err=%s", key, cmd.Err())
		return fmt.Errorf("set redis err:%v", cmd.Err())
	}
	r.Expire(key, redisTimeOut)
	return nil
}
