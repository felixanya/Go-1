package data

import (
	"errors"
	"fmt"
	"steve/structs"
	"strconv"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/go-redis/redis"
)

const (
	almsPlayerKey = "almsPlayerID:%v" // 救济玩家对应已经领取的数量
	almsConfigKey = "almsConfig"      // 救济金配置
)

var redisClifunc = getAlmsRedis //获取redisClien
var errRobotRedisGain = errors.New("robot_redis 获取失败")
var errRobotRedisOpertaion = errors.New("robot_redis 操作失败")

// RedisTimeOut 过期时间 24小时
var RedisTimeOut = time.Hour * 24

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

//SetAlmsConfig 设置救济金配置
func SetAlmsConfig(fields map[string]interface{}, date time.Duration) error {
	redisCli := redisClifunc()
	key := almsConfigKey
	status := redisCli.HMSet(key, fields)
	if status.Err() != nil {
		return fmt.Errorf("设置失败(%v)", status.Err())
	}
	redisCli.Expire(key, date)
	return nil
}

//GetAlmsConfigRedis 获取救济金配置
func GetAlmsConfigRedis(fields ...string) (*MyAlmsConfig, error) {
	redisCli := redisClifunc()
	key := almsConfigKey
	result, err := redisCli.HMGet(key, fields...).Result()
	if err != nil || len(result) != len(fields) {
		return nil, fmt.Errorf("获取 redis 结果失败(%s) fields=(%v)", err.Error(), fields)
	}
	almsConfig := &MyAlmsConfig{}
	for index, field := range fields {
		v, ok := result[index].(string)
		if !ok {
			return nil, fmt.Errorf("错误的数据类型。field=(%s) val=(%v)", field, result[index])
		}
		setMyAlmsConfiByFiedld(almsConfig, v, field)
	}
	return almsConfig, nil
}

func setMyAlmsConfiByFiedld(almsConfig *MyAlmsConfig, val string, field string) error {
	newVal, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return err
	}
	switch field {
	case Almsnorm:
		almsConfig.GetNorm = newVal
	case Almsnumber:
		almsConfig.GetNumber = newVal
	case Almstimes:
		almsConfig.GetTimes = int(newVal)
	case AlmsCountDonw:
		almsConfig.AlmsCountDonw = int(newVal)
	case DepositCountDonw:
		almsConfig.DepositCountDonw = int(newVal)
	case Version:
		almsConfig.Version = int(newVal)
	}
	return nil
}
