package prop

import (
	"fmt"
	"steve/entity/cache"
	"steve/entity/db"
	"steve/entity/prop"
	"steve/external/propsclient"
	"steve/structs"
	"strconv"
	"time"

	"github.com/Sirupsen/logrus"

	"github.com/go-redis/redis"
	"github.com/go-xorm/xorm"
)

// redis 过期时间
var redisTimeOut = time.Hour * 24

const (
	playerRedisName          = "player"
	playerMysqlName          = "player"
	playerTableName          = "t_player"
	playerCurrencyTableName  = "t_player_currency"
	playerGameTableName      = "t_player_game"
	gameconfigTableName      = "t_game_config"
	gamelevelconfigTableName = "t_game_level_config"
	playerPropsTableName     = "t_player_props"
)

func getRedisCli(redis string, db int) (*redis.Client, error) {
	exposer := structs.GetGlobalExposer()
	redisCli, err := exposer.RedisFactory.GetRedisClient(redis, db)
	if err != nil {
		return nil, fmt.Errorf("获取 redis client 失败: %v", err)
	}
	return redisCli, nil
}

func getMysqlEngine(mysqlName string) (*xorm.Engine, error) {
	exposer := structs.GetGlobalExposer()
	engine, err := exposer.MysqlEngineMgr.GetEngine(mysqlName)
	if err != nil {
		return nil, fmt.Errorf("获取 mysql 引擎失败：%v", err)
	}
	return engine, nil
}

// 单元测试通过这两个值修改 mysql 引擎获取和 redis cli 获取
var mysqlEngineGetter = getMysqlEngine
var redisCliGetter = getRedisCli

// GetPlayerAllProps 获取玩家的所有道具
func GetPlayerAllProps(playerID uint64) (props []prop.Prop, err error) {
	// 获取道具配置信息
	propConfig, err := GetPropsConfig()
	if err != nil {
		return nil, err
	}

	// 获取玩家的所有道具
	items, err := propsclient.GetUserProps(playerID, 0)

	propsCount := make(map[uint64]int64, 0)
	for _, item := range items {
		propsCount[item.PropsId] = item.PropsNum
	}

	props = make([]prop.Prop, len(propConfig))
	for index, attr := range propConfig {
		props[index] = prop.Prop{
			PropID: int32(attr.PropID),
			Count:  propsCount[uint64(attr.PropID)],
		}
	}
	logrus.Debugf("获取玩家playerID:(%d)的所有道具,道具:(%v)", playerID, props)
	return
}

// getPlayerProps 获取玩家的道具,获取单个或多个道具，通过fields参数区分
func getPlayerProps(playerID uint64, propID int32, fields ...string) (prop prop.Prop, err error) {
	//从 redis 获取
	prop, err = getPlayerPropFieldsFromRedis(playerID, propID, fields)
	if err == nil {
		return
	}
	//从 DB 获取
	prop, err = getPlayerPropFieldsFromDB(playerID, propID, fields)
	return
}

func getPlayerPropFieldsFromRedis(playerID uint64, propID int32, fields []string) (prop prop.Prop, err error) {
	redisCli, err := redisCliGetter(playerRedisName, 0)
	if err != nil {
		return prop, fmt.Errorf("获取 redis 客户端失败(%s)。", err.Error())
	}

	propKey := cache.FmtPlayerPropKey(playerID, propID)
	cmd := redisCli.HMGet(propKey, fields...)
	if cmd.Err() != nil {
		return prop, fmt.Errorf("执行 redis 命令失败(%s)", cmd.Err().Error())
	}

	result, err := cmd.Result()
	if err != nil || len(result) != len(fields) {
		return prop, fmt.Errorf("获取 redis 结果失败(%s) fields=(%v)", err.Error(), fields)
	}

	for index, field := range fields {
		v, ok := result[index].(string)
		if !ok {
			return prop, fmt.Errorf("错误的数据类型。field=%s val=%v", field, result[index])
		}
		if err = parsePropByField(&prop, field, v); err != nil {
			return prop, fmt.Errorf("解析数据错误%s。field=%s val=%v", err.Error(), field, result[index])
		}
	}
	redisCli.Expire(propKey, redisTimeOut)
	return
}

func getPlayerPropFieldsFromDB(playerID uint64, propID int32, fields []string) (prop prop.Prop, err error) {
	// 从数据库获取
	engine, err := mysqlEngineGetter(playerMysqlName)
	if err != nil {
		return
	}
	strCol := ""
	for _, col := range fields {
		if len(strCol) > 0 {
			strCol += ","
		}
		strCol += col
	}

	sql := fmt.Sprintf("select %s from t_player_props  where playerID='%d' and propID='%d';", strCol, playerID, propID)
	res, err := engine.QueryString(sql)

	if err != nil {
		err = fmt.Errorf("select t_player_props sql:%s ,err：%v", sql, err)
		return
	}

	if len(res) == 0 { // 数据库里面没有也需要更新到redis，避免重复查询db
		prop.PropID = propID
		prop.Count = 0
		goto update
	}

	if len(res) != 1 {
		err = fmt.Errorf("玩家(%d)存在多条 propID:%d 信息记录： %v", playerID, propID, err)
		return
	}

	prop, err = generateDbPlayerProp(playerID, propID, res[0], fields...)
	if err != nil {
		err = fmt.Errorf("generate dbPlayerGame 失败(%v)", err.Error())
	}

update:
	// 更新redis
	if err = updatePlayerPropFieldsToRedis(playerID, propID, fields, &prop); err != nil {
		err = fmt.Errorf("更新 redis 失败(%v)", err.Error())
	}
	return
}

func updatePlayerPropFieldsToRedis(playerID uint64, propID int32, fields []string, prop *prop.Prop) error {
	redisCli, err := redisCliGetter(playerRedisName, 0)
	if err != nil {
		return fmt.Errorf("获取 redis 客户端失败(%s)。", err.Error())
	}
	playerPropKey := cache.FmtPlayerPropKey(playerID, propID)
	kv := make(map[string]interface{}, len(fields))
	for _, field := range fields {
		v, err := getDBPlayerPropField(field, prop)
		if err != nil {
			return err
		}
		if v == nil {
			continue
		}
		kv[field] = v
	}
	status := redisCli.HMSet(playerPropKey, kv)
	if status.Err() != nil {
		return fmt.Errorf("设置失败(%v)", status.Err())
	}
	redisCli.Expire(playerPropKey, redisTimeOut)
	return nil
}

// updatePropCountToRedis 更新redis中玩家道具的个数
// param  playerID:玩家ID  propID : 道具ID  oldCount : 当前道具个数  newCount ： 要更新成的个数
func updatePropCountToRedis(playerID uint64, propID int32, oldCount, newCount uint32) error {
	redisCli, err := redisCliGetter(playerRedisName, 0)
	if err != nil {
		return fmt.Errorf("获取 redis 客户端失败(%s)。", err.Error())
	}

	playerPropKey := cache.FmtPlayerPropKey(playerID, propID)

	kv := map[string]interface{}{
		"propID": propID,
		"count":  newCount,
	}

	err = redisCli.Watch(func(tx *redis.Tx) error {
		_, err = tx.Pipelined(func(pipe redis.Pipeliner) error {
			pipe.HMSet(playerPropKey, kv)
			return nil
		})
		return err
	}, playerPropKey)

	if err == nil {
		redisCli.Expire(playerPropKey, redisTimeOut)
	}
	return err
}

// expirePlayerProp 设置玩家道具redis key失效
func expirePlayerProp(playerID uint64, propID int32) bool {
	redisCli, err := redisCliGetter(playerRedisName, 0)
	if err != nil {
		return false
	}

	playerPropKey := cache.FmtPlayerPropKey(playerID, propID)

	return redisCli.Del(playerPropKey).Val() == 1
}

// updatePropCountToMysql 更新Mysql中玩家道具的个数
// param  playerID:玩家ID  propID : 道具ID  oldCount : 当前道具个数  newCount ： 要更新成的个数
func updatePropCountToMysql(playerID uint64, propID int32, oldCount, newCount uint32) error {
	// 从数据库获取
	engine, err := mysqlEngineGetter(playerMysqlName)
	if err != nil {
		return err
	}
	fields := []string{"count"}

	sql := fmt.Sprintf("select count from t_player_props  where playerID='%d' and propID='%d';", playerID, propID)
	res, err := engine.QueryString(sql)

	if err != nil {
		err = fmt.Errorf("select t_player_props sql:%s ,err：%v", sql, err)
		return err
	}

	dbCount := 0

	if len(res) > 1 {
		err = fmt.Errorf("玩家(%d)存在多条 propID:%d 信息记录： %v", playerID, propID, err)
		return err
	}

	if len(res) == 1 {
		dbCount, _ = strconv.Atoi(res[0][fields[0]])
	}

	if uint32(dbCount) != oldCount {
		err = fmt.Errorf("msqyl中玩家(%d) 道具propID:(%d) 个数为:(%d),与 redis中道具个数:(%d)不一致 ", playerID, propID, dbCount, oldCount)
		return err
	}

	dbProp := db.TPlayerProps{
		Playerid:   int64(playerID),
		Propid:     int64(propID),
		Count:      int64(newCount),
		Createtime: time.Now(),
		Createby:   "programmer",
		Updatetime: time.Now(),
		Updateby:   "programmer",
	}

	if len(res) == 1 {
		_, err = engine.Table(playerPropsTableName).Where("playerID = ? and propID = ?", playerID, propID).Cols(fields...).Update(dbProp)
		if err != nil {
			return fmt.Errorf("更新失败 (%v)", err.Error())
		}
	} else {
		affected, err := engine.Table(playerPropsTableName).Insert(&dbProp)
		if err != nil || affected == 0 {
			return fmt.Errorf("insert t_player_props sql error：(%v)， affect=(%d)", err, affected)
		}
	}

	return nil
}

func generateDbPlayerProp(playerID uint64, propID int32, info map[string]string, fields ...string) (prop prop.Prop, err error) {
	for _, field := range fields {
		v, ok := info[field]
		if !ok {
			return prop, fmt.Errorf("错误的数据类型。field=%s val=%v", field, info)
		}
		if err = parsePropByField(&prop, field, v); err != nil {
			return prop, err
		}
	}
	return
}

func parsePropByField(prop *prop.Prop, field string, val string) (err error) {
	switch field {
	case "propID":
		temp, _ := strconv.ParseInt(val, 10, 64)
		prop.PropID = int32(temp)
	case "count":
		prop.Count, _ = strconv.ParseInt(val, 10, 64)
	case "createTime":
	case "createBy":
	case "updateTime":
	case "updateBy":
		return nil
	default:
		return fmt.Errorf("未处理的字段:%s", field)
	}
	return nil
}

func getDBPlayerPropField(field string, prop *prop.Prop) (val interface{}, err error) {
	switch field {
	case "propID":
		val = prop.PropID
	case "count":
		val = prop.Count
	case "playerID", "createTime", "createBy", "updateTime", "updateBy":
		val = nil
	default:
		val = nil
		err = fmt.Errorf("未处理字段：%s", field)
	}

	return
}
