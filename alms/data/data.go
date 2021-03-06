package data

import (
	"encoding/json"
	"errors"
	"fmt"
	"steve/external/configclient"

	"github.com/Sirupsen/logrus"
)

//MyAlmsConfig  救济金配置
type MyAlmsConfig struct {
	GetNorm          int64 // 救济线
	GetTimes         int   // 最多领取次数
	GetNumber        int64 // 领取数量
	AlmsCountDonw    int   // 救济倒计时，时间是秒
	DepositCountDonw int   // 快充倒计时，时间是秒
	PlayerGotTimes   int   // 玩家已领取数量
	Version          int   // 救济金配置表版本号,初始1
}

//UpdatePlayerGotTimesByPlayerID 根据玩家ID修改玩家已经领取数量(db,redis)
func UpdatePlayerGotTimesByPlayerID(playerID uint64, changeTimes int) error {
	entry := logrus.WithFields(logrus.Fields{
		"func_name":   "getAlmsConfigByPlayerID",
		"changeTimes": changeTimes,
	})
	// redis
	if err := UpdateAlmsPlayerGotTimes(playerID, changeTimes, RedisTimeOut); err != nil {
		entry.WithError(err).Errorln("修改玩家已经领取数量 redis 失败 playerID(%v)", playerID)
		return err
	}
	//db
	if err := updateMysqlPlayerGotTimesByPlayerID(playerID, changeTimes); err != nil {
		entry.WithError(err).Errorln("修改玩家已经领取数量 DB 失败 playerID(%v)", playerID)
		return err
	}
	entry.Debugln("修改玩家已经领取数量")
	return nil
}

//GetPlayerGotTimesByPlayerID 根据玩家ID获取救济金已领取数量，先从redis,不存在从db，再存入redis
func GetPlayerGotTimesByPlayerID(playerID uint64) (int, error) {
	entry := logrus.WithFields(logrus.Fields{
		"func_name": "GetPlayerGotTimesByPlayerID",
	})
	// redis获取当前玩家救济领取数量
	times, err := GetAlmsPlayerGotTimes(playerID)
	if err == nil {
		return times, nil
	}
	entry.WithError(err).Debugln("重新从db获取已领取次数")
	dbtimes, dberr := getMysqlPlayerGotTimesByPlayerID(playerID) //t_hall_info 可能不存在该玩家id,
	if dberr != nil {
		entry.WithError(dberr).Errorln("获取救济金已领取数量失败")
		return 0, dberr
	}
	// 存入redis
	if err := UpdateAlmsPlayerGotTimes(playerID, dbtimes, RedisTimeOut); err != nil {
		entry.WithError(err).Errorln("存储玩家救济金领取次数失败")
	}
	return dbtimes, nil
}

const (
	// alms
	Almsnorm         = "norm"
	Almsnumber       = "number"
	Almstimes        = "times"
	AlmsCountDonw    = "almsCountDonw"
	DepositCountDonw = "depositCountDonw"
	Version          = "version"
	// gamelevel
)

//InitAlmsConfig 初始话救济金配置
func InitAlmsConfig() error {
	// gamelevelConfig
	gameLevels := make([]*GameLevel, 0)
	// 获取救济金配置
	gameLeveConfigMaps, err := configclient.GetAllGameLevelConfig()
	if err != nil {
		logrus.WithError(err).Debugln("获取gameLevel配置失败")
		return err
	}
	for _, gameLeveConfigMap := range gameLeveConfigMaps {
		gl := &GameLevel{
			GameID:   gameLeveConfigMap.GameID,
			LevelID:  gameLeveConfigMap.LevelID,
			LowSorce: gameLeveConfigMap.LowScores,
			IsOpen:   gameLeveConfigMap.IsAlms,
		}
		gameLevels = append(gameLevels, gl)
	}
	jons, err := GameLevelsToJSON(gameLevels)
	if err == nil {
		// 保存
		err = SetGameLevlConfig(jons, RedisTimeOut)
	}
	if err != nil {
		logrus.WithError(err).Debugln("init save gamelevelconfig")
	}

	// almsConfig
	// 获取救济金配置
	almsConfigMap, err := configclient.GetAlmsConfigMap()
	if err != nil {
		logrus.WithError(err).Debugln("获取救济金配置失败")
		return err
	}
	if len(almsConfigMap) == 0 {
		return errors.New("救济金配置为0")
	}
	acf := almsConfigMap[0]
	resultAlmsConfig := &MyAlmsConfig{
		GetNorm:          int64(acf.GetNorm),
		GetNumber:        int64(acf.GetNumber),
		GetTimes:         acf.GetTimes,
		AlmsCountDonw:    acf.AlmsCountDown,
		DepositCountDonw: acf.DepositCountDown,
		Version:          acf.Version,
	}
	// 保存
	return SaveAlmsConfigRedis(resultAlmsConfig)
}

//SaveAlmsConfigRedis 保存配置到redis
func SaveAlmsConfigRedis(almsConfig *MyAlmsConfig) error {
	if almsConfig == nil {
		logrus.Debugln("save almsConfig 失败  almsConfig eq nil")
		return fmt.Errorf("save almsConfig 失败  almsConfig eq nil")
	}
	fields := make(map[string]interface{})
	fields[Almsnorm] = almsConfig.GetNorm
	fields[Almsnumber] = almsConfig.GetNumber
	fields[Almstimes] = almsConfig.GetTimes
	fields[AlmsCountDonw] = almsConfig.AlmsCountDonw
	fields[DepositCountDonw] = almsConfig.DepositCountDonw
	fields[Version] = almsConfig.Version
	return SetAlmsConfig(fields, RedisTimeOut)
}

//GetAlmsConfig 获取救济金配置
func GetAlmsConfig(playerID uint64) (*MyAlmsConfig, error) {
	// 获取玩家已经领取次数
	gotTimes, err := GetPlayerGotTimesByPlayerID(playerID)
	if err != nil {
		return nil, err
	}
	// 先从redis
	resultAlmsConfig, err := GetAlmsConfigRedis(Almsnorm, Almsnumber, Almstimes, AlmsCountDonw, DepositCountDonw, Version)
	if err == nil {
		resultAlmsConfig.PlayerGotTimes = gotTimes
		return resultAlmsConfig, nil
	}
	logrus.WithError(err).Debugln("GetAlmsConfig redis")
	// 获取救济金配置
	almsConfigMap, err := configclient.GetAlmsConfigMap()
	if err != nil {
		logrus.WithError(err).Debugln("获取救济金配置失败")
		return nil, err
	}
	if len(almsConfigMap) == 0 {
		return nil, errors.New("救济金配置为0")
	}
	acf := almsConfigMap[0]
	resultAlmsConfig = &MyAlmsConfig{
		GetNorm:          int64(acf.GetNorm),
		GetNumber:        int64(acf.GetNumber),
		GetTimes:         acf.GetTimes,
		AlmsCountDonw:    acf.AlmsCountDown,
		DepositCountDonw: acf.DepositCountDown,
		Version:          acf.Version,
		PlayerGotTimes:   gotTimes,
	}
	// 重新保存
	SaveAlmsConfigRedis(resultAlmsConfig)
	return resultAlmsConfig, nil
}

//GameLevel gamelevelconfig
type GameLevel struct {
	GameID   int `json:"gameID"`
	LevelID  int `json:"levelID"`
	LowSorce int `json:"lowSorce"`
	IsOpen   int `json:"isOpen"`
}

//GetGameLevelConfig 获取游戏场次配置
func GetGameLevelConfig() ([]*GameLevel, error) {
	// redis
	if gamelevelJSON, err := GetGameLevlConfigRedis(); err == nil {
		return JSONToGameLevels(gamelevelJSON)
	}
	// configserver
	gameLevels := make([]*GameLevel, 0)
	// 获取救济金配置
	gameLeveConfigMaps, err := configclient.GetAllGameLevelConfig()
	if err != nil {
		logrus.WithError(err).Debugln("获取gameLevel配置失败")
		return nil, err
	}
	for _, gameLeveConfigMap := range gameLeveConfigMaps {
		gl := &GameLevel{
			GameID:   gameLeveConfigMap.GameID,
			LevelID:  gameLeveConfigMap.LevelID,
			LowSorce: gameLeveConfigMap.LowScores,
			IsOpen:   gameLeveConfigMap.IsAlms,
		}
		gameLevels = append(gameLevels, gl)
	}
	jons, err := GameLevelsToJSON(gameLevels)
	if err == nil {
		// 保存
		err = SetGameLevlConfig(jons, RedisTimeOut)
	}
	if err != nil {
		logrus.WithError(err).Debugln("save gamelevelconfig")
	}
	return gameLevels, nil
}

// GameLevelsToJSON gameLevel to JSON
func GameLevelsToJSON(gameLevels []*GameLevel) (string, error) {
	if len(gameLevels) == 0 {
		logrus.Debugln("gameLevels eq 0")
		return "", nil
	}
	bytes, err := json.Marshal(gameLevels)
	return string(bytes), err
}

// JSONToGameLevels JSON TO gameLevel
func JSONToGameLevels(gameLevelJSON string) ([]*GameLevel, error) {
	if gameLevelJSON == "" {
		logrus.Debugln("gameLevelJSON eq 0")
		return nil, nil
	}
	gameLevels := []*GameLevel{}
	err := json.Unmarshal([]byte(gameLevelJSON), &gameLevels)
	return gameLevels, err
}
