package configclient

import (
	"encoding/json"
	entityConf "steve/entity/config"
	"steve/entity/constant"
	"time"

	"github.com/Sirupsen/logrus"
)

// 重试次数
var reTryCount = 20

// 重试时间间隔
var reTryTime = time.Second * 5

// ParseToGameLevelConfigMap 反序列化gameLevelConfig
func ParseToGameLevelConfigMap(jsonStr string) (conf []entityConf.GameLevelConfig) {
	if err := json.Unmarshal([]byte(jsonStr), &conf); err != nil {
		logrus.Errorf("游戏配置数据反序列化失败：%s", err.Error())
	}
	return
}

// GetRoleInitConfigMap 获取角色初始配置
func GetRoleInitConfigMap() (roleConf []entityConf.RoleConfig, err error) {
	roleConfStr, err := GetConfigUntilSucc("role", "init", reTryCount, reTryTime)
	if err != nil {
		logrus.WithError(err).Errorln("获取角色初始属性配置失败")
		return nil, err
	}
	if err := json.Unmarshal([]byte(roleConfStr), &roleConf); err != nil {
		logrus.WithError(err).Errorf("游戏级别配置数据反序列化失败：%s", err.Error())
		return nil, err
	}
	for i := 0; i < len(roleConf); i++ {
		roleConf[i].InitItem()
	}
	return roleConf, nil
}

// GetAlmsConfigMap 获取救济金配置
func GetAlmsConfigMap() (conf []entityConf.AlmsConfig, err error) {
	almsStr, err := GetConfigUntilSucc("game", "alms", reTryCount, reTryTime)
	if err != nil {
		logrus.WithError(err).Errorln("获取救济金配置失败")
		return nil, err
	}
	if err := json.Unmarshal([]byte(almsStr), &conf); err != nil {
		logrus.WithError(err).Errorf("游戏配置数据反序列化失败：%s", err.Error())
		return nil, err
	}
	return
}

// GetGameConfig 获取游戏配置信息
func GetGameConfig(gameID int) (gameConf entityConf.GameConfig, err error) {
	gameConfigMap, err := GetGameConfigMap()
	if err != nil {
		logrus.WithError(err).Errorln("获取游戏级别配置失败！！")
		return
	}
	gameConf, exists := gameConfigMap[gameID]
	if !exists {
		logrus.WithField("gameConfigMap", gameConfigMap).WithField("gameId", gameID).Errorln("找不到游戏配置")
		return
	}
	return
}

// GetAllGameConfig 获取所有游戏配置信息
func GetAllGameConfig() (gameConf []entityConf.GameConfig, err error) {
	gameStr, err := GetConfigUntilSucc(constant.GameConfigKey.Key, constant.GameConfigKey.SubKey, 20, 3*time.Second)
	if err != nil {
		logrus.WithError(err).Errorln("获取游戏配置失败")
		return nil, err
	}

	if err := json.Unmarshal([]byte(gameStr), &gameConf); err != nil {
		logrus.WithError(err).Errorf("游戏配置数据反序列化失败：%s", err.Error())
		return nil, err
	}

	return
}

// GetGameConfigMap 获取游戏配置信息Map<gameId, GameConfig>
func GetGameConfigMap() (gameConfMap map[int]entityConf.GameConfig, err error) {
	gameConfMap = make(map[int]entityConf.GameConfig)

	gameConf, err := GetAllGameConfig()
	if err != nil {
		return
	}

	for _, game := range gameConf {
		gameConfMap[game.GameID] = game
	}
	return
}

// GetGameLevelConfig 获取游戏级别配置信息
func GetGameLevelConfig(gameID int, levelID int) (levelConf entityConf.GameLevelConfig, err error) {
	gameLevelMap, err := GetGameLevelConfigMap()
	if err != nil {
		logrus.WithError(err).Errorln("获取游戏级别配置失败！！")
		return
	}
	levelMap, exists := gameLevelMap[gameID]
	if !exists {
		logrus.WithField("gameLevelMap", gameLevelMap).WithField("gameId", gameID).Errorln("找不到游戏配置")
		return
	}
	levelConf, exists = levelMap[levelID]
	if !exists {
		logrus.WithField("gameLevelMap", gameLevelMap).WithField("gameId", gameID).WithField("levelId", levelID).Errorln("找不到游戏场次")
		return
	}
	return
}

// GetAllGameLevelConfig 获取所有游戏级别配置信息
func GetAllGameLevelConfig() (levelConf []entityConf.GameLevelConfig, err error) {
	levelStr, err := GetConfigUntilSucc(constant.GameLevelConfigKey.Key, constant.GameLevelConfigKey.SubKey, 20, 3*time.Second)
	if err != nil {
		logrus.WithError(err).Errorln("获取游戏级别配置失败")
		return nil, err
	}

	if err := json.Unmarshal([]byte(levelStr), &levelConf); err != nil {
		logrus.WithError(err).Errorf("游戏级别配置数据反序列化失败：%s", err.Error())
		return nil, err
	}

	return
}

// GetGameLevelConfigMap 获取游戏级别配置信息Map<gameId, Map<levelId, GameLevelConfig>>
func GetGameLevelConfigMap() (gameLevelMap map[int]map[int]entityConf.GameLevelConfig, err error) {
	gameLevelMap = make(map[int]map[int]entityConf.GameLevelConfig)
	levelConf, err := GetAllGameLevelConfig()

	for _, level := range levelConf {
		levelMap, exists := gameLevelMap[level.GameID]
		if !exists {
			levelMap = make(map[int]entityConf.GameLevelConfig)
			gameLevelMap[level.GameID] = levelMap
		}
		levelMap[level.LevelID] = level
	}
	return
}
