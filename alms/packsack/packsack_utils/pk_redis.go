package packsack_utils

import (
	"errors"
	"fmt"
	"steve/structs"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/go-redis/redis"
)

// redis 过期时间
var redisTimeOut time.Duration = time.Minute * 60 * 24 * 30
var errRobotRedisGain = errors.New("robot_redis 获取失败")

var Myredis = getRedis

// getRedis 获取redis
func getRedis() *redis.Client {
	e := structs.GetGlobalExposer()
	redis, err := e.RedisFactory.NewClient()
	if err != nil {
		logrus.WithError(err).Errorln(errRobotRedisGain)
		return nil
	}
	return redis
}

//GetPlayerPacksackGold 从redis加载玩家金币
func GetPlayerPacksackGold(uid uint64) (int64, error) {
	// r := common_redis.GetRedisClient()
	r := Myredis()
	key := fmtPlayerKey(uid)
	cmd := r.Get(key)
	if cmd.Err() != nil {
		return 0, fmt.Errorf("get redis err:%v", cmd.Err())
	}
	m, err := cmd.Int64()
	if err != nil {
		logrus.WithError(err).Errorln("redis 命令执行失败")
		return 0, err
	}
	return m, nil
}

//SaveGoldToRedis 保存玩家金币到Redis
func SaveGoldToRedis(uid uint64, gold int64) error {
	entry := logrus.WithFields(logrus.Fields{
		"func_name": "SaveGoldToRedis",
		"playerID":  uid,
		"gold":      gold,
	})
	redisCli := Myredis()
	key := fmtPlayerKey(uid)
	err := redisCli.Watch(func(tx *redis.Tx) error {
		err := tx.Get(key).Err()
		if err != nil && err != redis.Nil {
			return err
		}
		_, err = tx.Pipelined(func(pipe redis.Pipeliner) error {
			pipe.Set(key, gold, redisTimeOut)
			return nil
		})
		return err
	}, key)
	if err == redis.TxFailedErr {
		entry.WithError(err).Errorln("重试")
		return SaveGoldToRedis(uid, gold)
	}
	return nil
}

// 格式化Redis Key
func fmtPlayerKey(uid uint64) string {
	return fmt.Sprintf("packSack_%v", uid)
}

const packsackPropInfo = "packsackPropInfo"

//SetPropConfigRedis 设置道具配置
func SetPropConfigRedis(field string) error {
	redisCli := Myredis()
	key := packsackPropInfo
	status := redisCli.Set(key, field, redisTimeOut)
	if status.Err() != nil {
		return fmt.Errorf("设置失败(%v)", status.Err())
	}
	redisCli.Expire(key, redisTimeOut)
	return nil
}

//GetPropConfigRedis 获取道具配置
func GetPropConfigRedis() (string, error) {
	client := Myredis()
	key := packsackPropInfo
	data, err := client.Get(key).Result()
	if err != nil {
		logrus.WithError(err).Errorln("redis 命令执行失败")
		return "", err
	}
	return data, nil
}
