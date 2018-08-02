package data

import (
	"encoding/json"
	"fmt"
	"steve/entity/cache"
	"steve/entity/db"
	"steve/gutils"
	"steve/server_pb/user"
	"strconv"
	"strings"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/spf13/viper"
)

// idAllocObject id分配
var idAllocObject *gutils.Node

// redis 过期时间
var redisTimeOut = time.Hour * 24 * 30

// 玩家基本信息列表
var playerInfoList = map[int32]string{
	1: "nickname",
	2: "avatar",
	3: "gender",
	4: "channelID",
	5: "provinceID",
	6: "cityID",
}

const (
	playerRedisName          = "player"
	playerMysqlName          = "steve"
	playerTableName          = "t_player"
	playerCurrencyTableName  = "t_player_currency"
	playerGameTableName      = "t_player_game"
	gameconfigTableName      = "t_game_config"
	gamelevelconfigTableName = "t_game_level_config"
)

type gameConfigDetail struct {
	db.TGameConfig      `xorm:"extends"`
	db.TGameLevelConfig `xorm:"extends"`
}

// GetPlayerIDByAccountID 根据账号 ID 获取其关联的玩家 ID
func GetPlayerIDByAccountID(accountID uint64) (exist bool, playerID uint64, err error) {
	exist, playerID, err = false, 0, nil

	redisKey := cache.FmtAccountPlayerKey(accountID)
	playerID, err = getRedisUint64Val(playerRedisName, redisKey)
	if err == nil {
		return
	}
	engine, err := mysqlEngineGetter(playerMysqlName)
	if err != nil {
		return
	}
	where := fmt.Sprintf("accountID=%d", accountID)
	var dbPlayerID struct {
		ID uint64 `xorm:"playerID"`
	}
	exist, err = engine.Table(playerTableName).Select("playerID").Where(where).Get(&dbPlayerID)
	if err != nil {
		err = fmt.Errorf("select sql err：err=%v", err)
		return
	}
	if exist {
		playerID = dbPlayerID.ID
		if err := setRedisVal(playerRedisName, redisKey, playerID, time.Hour*24); err != nil {
			err = fmt.Errorf("save playerId into redis fail： %v", err)
		}
	}
	return
}

// GetPlayerInfo 根据玩家id获取玩家的基本信息
func GetPlayerInfo(playerID uint64) (player *PlayerInfo, err error) {
	player, err = new(PlayerInfo), nil

	engine, err := mysqlEngineGetter(playerMysqlName)
	if err != nil {
		return
	}
	strCol := ""
	for _, col := range playerInfoList {
		if len(strCol) > 0 {
			strCol += ","
		}
		strCol += col
	}

	sql := fmt.Sprintf("select %s from t_player  where playerID='%d';", strCol, playerID)
	res, err := engine.QueryString(sql)
	if err != nil {
		err = fmt.Errorf("select sql err：sql=%s,err=%v", sql, err)
		return
	}
	if len(res) != 1 {
		err = fmt.Errorf("玩家存在多条信息记录： %v", err)
		return
	}
	player.generatePlayerInfo(res[0])

	return
}

// UpdatePlayerInfo 修改玩家个人信息
func UpdatePlayerInfo(playerID uint64, nickName, avatar string, gender uint32) (exist, result bool, err error) {
	entry := logrus.WithFields(logrus.Fields{
		"opr":      "update_player_info",
		"playerID": playerID,
		"nickName": nickName,
		"avatar":   avatar,
	})
	exist, result, err = true, true, nil

	rfields := make(map[string]interface{}, 0)
	if nickName != "" {
		rfields[cache.NickName] = nickName
	}
	if avatar != "" {
		rfields[cache.Avatar] = avatar
	}
	if gender == 1 || gender == 2 {
		rfields[cache.Gender] = gender
	}

	strCol := "playerID="
	strCol += fmt.Sprintf("'%v'", playerID)
	for key, field := range rfields {
		strCol += ","
		strCol += key
		strCol += "="
		strCol += fmt.Sprintf("'%v'", field)
	}
	engine, err := mysqlEngineGetter(playerMysqlName)
	sql := fmt.Sprintf("update t_player set %s  where playerID=?;", strCol)
	res, sqlerror := engine.Exec(sql, playerID)
	if sqlerror != nil {
		entry.WithError(sqlerror).Errorln("update t_player mysql fail,sql:=%s", sql)
		exist, result, err = true, false, sqlerror
	}
	if aff, aerr := res.RowsAffected(); aff == 0 {
		entry.WithError(err).Errorln("update t_player playerId:%d 不存在", playerID)
		exist, result, err = false, false, aerr
	}

	// list := make(map[string]string, len(rfields))
	// for key, field := range rfields {
	// 	list[key] = field.(string)
	// }
	// if err = SavePlayerInfoToRedis(playerID, list, playerRedisName); err != nil {
	// 	err = fmt.Errorf("save playerInfo  into redis fail： %v", err)
	// }
	return
}

// GetPlayerGameInfo 获取玩家游戏信息
func GetPlayerGameInfo(playerID uint64, gameID uint32) (exist bool, info *db.TPlayerGame, err error) {
	exist, info, err = false, new(db.TPlayerGame), nil

	engine, err := mysqlEngineGetter(playerMysqlName)

	where := fmt.Sprintf("playerID=%d and gameID='%d'", playerID, gameID)
	exist, err = engine.Table(playerGameTableName).Select("gameID").Where(where).Get(info)

	if err != nil {
		err = fmt.Errorf("select t_player_game sql err：%v", err)
		return
	}
	return
}

// GetPlayerState 获取游戏状态,游戏id,ip地址
func GetPlayerState(playerID uint64) (pState *PlayerState, err error) {
	enrty := logrus.WithFields(logrus.Fields{
		"func_name": GetPlayerState,
		"playerID":  playerID,
	})
	pState, err = new(PlayerState), nil

	val, err := loadFromRedis(playerID, playerRedisName)

	if err != nil {
		enrty.WithError(err).Warningln("get player state from redis fail")
		return pState, err
	}
	pState.generatePlayerState(val)
	return
}

// UpdatePlayerState 修改玩家游戏状态
func UpdatePlayerState(playerID uint64, oldState, newState, reqServerType uint32, serverAddr string) (result bool, err error) {
	result, err = true, nil
	redisKey := cache.FmtPlayerIDKey(uint64(playerID))

	val, _ := getRedisField(playerRedisName, redisKey, cache.GameState)
	state, _ := strconv.Atoi(val[0].(string))

	if oldState != uint32(state) {
		return
	}

	serverType := map[user.ServerType]string{
		user.ServerType_ST_GATE:  cache.GateAddr,
		user.ServerType_ST_MATCH: cache.MatchAddr,
		user.ServerType_ST_ROOM:  cache.RoomAddr,
	}[user.ServerType(reqServerType)]

	rfields := map[string]string{
		cache.GameState: fmt.Sprintf("%d", newState),
		serverType:      serverAddr,
	}
	if err = setRedisWatch(playerRedisName, redisKey, rfields, redisTimeOut); err != nil {
		err = fmt.Errorf("save playerInfo  into redis fail： %v", err)
	}
	return
}

// GetGameInfoList 获取游戏配置信息
func GetGameInfoList() (gameInfos []*user.GameConfig, gamelevelInfos []*user.GameLevelConfig, err error) {
	gameInfos, gamelevelInfos, err = make([]*user.GameConfig, 0), make([]*user.GameLevelConfig, 0), nil

	gameConfigKey := "gameconfig"
	gameLevelConfigKey := "gamelevelconfig"

	var dbgameConfigs []db.TGameConfig
	var dbgamelevelConfigs []db.TGameLevelConfig

	gameConfigdata, err := getRedisByteVal(playerRedisName, gameConfigKey)
	if gameConfigdata != nil && len(gameConfigdata) != 0 {
		err = json.Unmarshal(gameConfigdata, &dbgameConfigs)
	}
	gameLeveldata, err := getRedisByteVal(playerRedisName, gameLevelConfigKey)
	if gameConfigdata != nil && len(gameConfigdata) != 0 {
		err = json.Unmarshal(gameLeveldata, &dbgamelevelConfigs)
	}
	if err == nil {
		dbGameConfig2serverGameConfig(dbgameConfigs)
		dbGamelevelConfig2serverGameConfig(dbgamelevelConfigs)
		return
	}

	engine, err := mysqlEngineGetter(playerMysqlName)
	if err != nil {
		return
	}
	err = engine.Table(gameconfigTableName).Find(&dbgameConfigs)

	if err != nil {
		err = fmt.Errorf("select sql error： %v", err)
		return
	}

	err = engine.Table(gamelevelconfigTableName).Find(&dbgamelevelConfigs)

	if err != nil {
		err = fmt.Errorf("select sql error： %v", err)
		return
	}
	dbGameConfig2serverGameConfig(dbgameConfigs)
	dbGamelevelConfig2serverGameConfig(dbgamelevelConfigs)
	// 写入redis
	data, _ := json.Marshal(dbgameConfigs)
	if err = setRedisVal(playerRedisName, gameConfigKey, data, redisTimeOut); err != nil {
		err = fmt.Errorf("save game_config  into redis fail： %v", err)
	}
	data, _ = json.Marshal(dbgamelevelConfigs)
	if err = setRedisVal(playerRedisName, gameLevelConfigKey, data, redisTimeOut); err != nil {
		err = fmt.Errorf("save game_level_config  into redis fail： %v", err)
	}
	return
}

// AllocPlayerID 生成玩家 ID
func AllocPlayerID() uint64 {
	return uint64(idAllocObject.Generate().Int64())
}

// InitPlayerData 初始化玩家数据
func InitPlayerData(player db.TPlayer) error {
	engine, err := mysqlEngineGetter(playerMysqlName)
	if err != nil {
		return err
	}
	affected, err := engine.Table(playerTableName).Insert(&player)
	if err != nil || affected == 0 {
		return fmt.Errorf("insert sql error：%v， affect=%d", err, affected)
	}
	return nil
}

// InitPlayerCoin 初始化玩家货币信息
func InitPlayerCoin(currency db.TPlayerCurrency) error {
	engine, err := mysqlEngineGetter(playerMysqlName)
	if err != nil {
		return err
	}
	affected, err := engine.Table(playerCurrencyTableName).Insert(&currency)
	if err != nil || affected == 0 {
		return fmt.Errorf("insert t_player_cuccency sql：%v， affect=%d", err, affected)
	}
	return nil
}

// InitPlayerState 初始化玩家状态
func InitPlayerState(playerID int64) (err error) {
	redisKey := cache.FmtPlayerIDKey(uint64(playerID))

	rfields := map[string]string{
		cache.GameState: fmt.Sprintf("%d", user.PlayerState_PS_IDIE),
		cache.IPAddr:    fmt.Sprintf("%s", "127.0.0.1"),
	}

	if err = setRedisWatch(playerRedisName, redisKey, rfields, redisTimeOut); err != nil {
		err = fmt.Errorf("save player_state into redis fail： %v", err)
	}
	return
}

// loadFromRedis 从redis查找信息
func loadFromRedis(playerID uint64, redisName string) (map[string]string, error) {

	r, err := redisCliGetter(redisName, 0)
	if err != nil {
		return nil, err
	}

	redisKey := cache.FmtPlayerIDKey(playerID)

	cmd := r.HGetAll(redisKey)
	if cmd.Err() != nil {
		return nil, fmt.Errorf("get redis err:%v", cmd.Err())
	}
	m := cmd.Val()
	if len(m) == 0 {
		return nil, fmt.Errorf("redis no user: playerID=%d", playerID)
	}
	list := make(map[string]string, len(m))
	for k, v := range m {
		sp := strings.Split(k, "_")
		if len(sp) == 2 {
			k = sp[1]
		}
		list[k] = v
	}

	return list, nil
}

// SavePlayerInfoToRedis 玩家信息保存到redis
func SavePlayerInfoToRedis(playerID uint64, pinfo map[string]string, redisName string) error {
	r, err := redisCliGetter(redisName, 0)
	if err != nil {
		return err
	}

	redisKey := cache.FmtPlayerIDKey(playerID)
	list := make(map[string]interface{}, len(pinfo))
	for k, v := range pinfo {
		list[k] = v
	}
	cmd := r.HMSet(redisKey, list)
	if cmd.Err() != nil {
		return fmt.Errorf("set redis err:%v", cmd.Err())
	}
	r.Expire(redisKey, redisTimeOut)
	return nil
}

func init() {
	node := viper.GetInt("node")
	var err error
	idAllocObject, err = gutils.NewNode(int64(node))
	if err != nil {
		logrus.Panicf("创建 id 生成器失败: %v", err)
	}
}
