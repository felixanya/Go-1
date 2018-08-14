package data

import (
	"errors"
	"fmt"
	"steve/structs"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/go-redis/redis"
)

const (
	almsPlayerKey = "almsPlayerID:%v" // 救济玩家对应已经领取的数量
)

var redisClifunc = getAlmsRedis //获取redisClien
var errRobotRedisGain = errors.New("robot_redis 获取失败")
var errRobotRedisOpertaion = errors.New("robot_redis 操作失败")

// RedisTimeOut 过期时间 1小时
var RedisTimeOut = time.Hour

// getAlmsRedis 获取redis
func getAlmsRedis() *redis.Client {
	e := structs.GetGlobalExposer()
	redis, err := e.RedisFactory.NewClient()
	if err != nil {
		logrus.WithError(err).Errorln(errRobotRedisGain)
		return nil
	}
	return redis
}

//GetAlmsPlayerGotTimes 获取玩家已经领取的数量
func GetAlmsPlayerGotTimes(playerID uint64) (int, error) {
	entry := logrus.WithFields(logrus.Fields{
		"func_name": "GetAlmsPlayerGotTimes",
		"playerID":  playerID,
	})
	client := redisClifunc()
	key := fmt.Sprintf(almsPlayerKey, playerID)
	data, err := client.Get(key).Int64()
	if err != nil {
		entry.WithError(err).Errorln("redis 命令执行失败")
		return 0, err
	}
	return int(data), nil
}

//UpdateAlmsPlayerGotTimes 修改玩家救济金已领取数量
func UpdateAlmsPlayerGotTimes(playerID uint64, val int, date time.Duration) error {
	entry := logrus.WithFields(logrus.Fields{
		"func_name": "UpdateAlmsPlayerGotTimes",
		"playerID":  playerID,
		"val":       val,
	})
	redisCli := redisClifunc()
	key := fmt.Sprintf(almsPlayerKey, playerID)
	err := redisCli.Watch(func(tx *redis.Tx) error {
		err := tx.Get(key).Err()
		if err != nil && err != redis.Nil {
			return err
		}
		_, err = tx.Pipelined(func(pipe redis.Pipeliner) error {
			pipe.Set(key, val, date)
			return nil
		})
		return err
	}, key)
	if err == redis.TxFailedErr {
		entry.WithError(err).Errorln("重试")
		return UpdateAlmsPlayerGotTimes(playerID, val, date)
	}
	return nil
}
