package data

import (
	"fmt"
	"steve/entity/cache"
	"steve/entity/db"
	"steve/structs"
	"strconv"
	"time"

	"github.com/Sirupsen/logrus"

	"github.com/go-redis/redis"
)

const redisName = "back"
const playerRedisName = "player"

// redis 过期时间
var redisTimeOut = time.Hour * 24

func getRedisCli(redis string, db int) (*redis.Client, error) {
	exposer := structs.GetGlobalExposer()
	redisCli, err := exposer.RedisFactory.GetRedisClient(redis, db)
	if err != nil {
		return nil, fmt.Errorf("获取 redis client 失败: %v", err)
	}
	return redisCli, nil
}

// RedisCliGetter 单元测试通过这两个值修改 mysql 引擎获取和 redis cli 获取
var RedisCliGetter = getRedisCli

// SetPlayerMaxwinningstream 储存最大连胜
func SetPlayerMaxwinningstream(key string, maxStream int) error {
	redisCli, err := RedisCliGetter(redisName, 0)
	if err != nil {
		return err
	}
	err = redisCli.Set(key, maxStream, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

// GetPlayerMaxwinningstream 获取最大连胜
func GetPlayerMaxwinningstream(key string) (int, error) {
	redisCli, err := RedisCliGetter(redisName, 0)
	if err != nil {
		return 0, err
	}
	streamCmd := redisCli.Get(key)
	MaxStream, err := streamCmd.Int64()
	if err != nil {
		return 0, err
	}
	return int(MaxStream), nil
}

// UpdatePlayerGameToredis 更新玩家胜率
func UpdatePlayerGameToredis(tpg *db.TPlayerGame) error {
	redisCli, err := RedisCliGetter(playerRedisName, 0)
	if err != nil {
		return fmt.Errorf("获取 redis 客户端失败(%s)。", err.Error())
	}
	playerGameKey := cache.FmtPlayerGameInfoKey(uint64(tpg.Playerid), uint32(tpg.Gameid))
	logrus.Debugf("更新玩家胜率key:(%v),Maxwinningstream:(%v)", playerGameKey, tpg.Maxwinningstream)
	kv := map[string]interface{}{
		cache.WinningBurea:     tpg.Winningburea,
		cache.WinningRate:      tpg.Winningrate,
		cache.TotalBurea:       tpg.Totalbureau,
		cache.MaxMultiple:      tpg.Maxmultiple,
		cache.MaxWinningStream: tpg.Maxwinningstream,
	}

	status := redisCli.HMSet(playerGameKey, kv)
	if status.Err() != nil {
		return fmt.Errorf("设置失败(%v)", status.Err())
	}
	redisCli.Expire(playerGameKey, redisTimeOut)
	return err
}

func getPlayerGameFieldsFromRedis(playerID uint64, gameID uint32, fields []string) (*db.TPlayerGame, error) {
	redisCli, err := RedisCliGetter(playerRedisName, 0)
	if err != nil {
		return nil, fmt.Errorf("获取 redis 客户端失败(%s)。", err.Error())
	}
	playerGameKey := cache.FmtPlayerGameInfoKey(playerID, gameID)
	cmd := redisCli.HMGet(playerGameKey, fields...)
	if cmd.Err() != nil {
		return nil, fmt.Errorf("执行 redis 命令失败(%s)", cmd.Err().Error())
	}
	result, err := cmd.Result()
	if err != nil || len(result) != len(fields) {
		return nil, fmt.Errorf("获取 redis 结果失败(%s) fields=(%v)", err.Error(), fields)
	}
	var dbPlayerGame db.TPlayerGame
	for index, field := range fields {
		v, ok := result[index].(string)
		if !ok {
			return nil, fmt.Errorf("错误的数据类型。field=(%s) val=(%v)", field, result[index])
		}
		if err = setDBPlayerGameByField(&dbPlayerGame, field, v); err != nil {
			return nil, err
		}
	}
	redisCli.Expire(playerGameKey, redisTimeOut)
	return &dbPlayerGame, nil
}

// setDBPlayerGameFieldByName 设置 dbPlayerGame 中的指定字段
func setDBPlayerGameByField(dbPlayerGame *db.TPlayerGame, field string, val string) error {
	switch field {
	case "id":
		dbPlayerGame.Id, _ = strconv.ParseInt(val, 10, 64)
	case "userID":
		dbPlayerGame.Playerid, _ = strconv.ParseInt(val, 10, 64)
	case "gameID":
		dbPlayerGame.Gameid, _ = strconv.Atoi(val)
	case "gameName":
		dbPlayerGame.Gamename = val
	case "winningRate":
		dbPlayerGame.Winningrate, _ = strconv.ParseFloat(val, 64)
	case "winningBurea":
		dbPlayerGame.Winningburea, _ = strconv.Atoi(val)
	case "totalBureau":
		dbPlayerGame.Totalbureau, _ = strconv.Atoi(val)
	case "maxWinningStream":
		dbPlayerGame.Maxwinningstream, _ = strconv.Atoi(val)
	case "maxMultiple":
		dbPlayerGame.Maxmultiple, _ = strconv.Atoi(val)
	case "createTime":
	case "createBy":
	case "updateTime":
	case "updateBy":
		return nil
	default:
		return fmt.Errorf("未处理的字段:(%s)", field)
	}
	return nil
}
