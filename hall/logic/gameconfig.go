package logic

import (
	"encoding/json"
	"fmt"
	entityConf "steve/entity/config"
	"steve/entity/constant"
	"steve/external/configclient"
	"time"

	"github.com/Sirupsen/logrus"
)

// GameConf 游戏配置
var GameConf []entityConf.GameConfig

// LevelConf 场次配置
var LevelConf []entityConf.GameLevelConfig

// loadGameConfig load game config from configuration server
func loadGameConfig(retry int) (gameConf []entityConf.GameConfig, err error) {
	gameConfigJSON, err := configclient.GetConfigUntilSucc(constant.GameConfigKey.Key, constant.GameConfigKey.SubKey, retry, 5*time.Second)
	if err != nil {
		return nil, fmt.Errorf("获取game config失败：%s", err.Error())
	}
	return parseGameConfig(gameConfigJSON)
}

// loadGameLevelConfig load game level config from configuration server
func loadGameLevelConfig(retry int) (gameConf []entityConf.GameLevelConfig, err error) {
	gameLevelConfigJSON, err := configclient.GetConfigUntilSucc(constant.GameLevelConfigKey.Key, constant.GameLevelConfigKey.SubKey, retry, 5*time.Second)
	if err != nil {
		return nil, fmt.Errorf("获取game level config失败：%s", err.Error())
	}
	return parseGameLevelConfig(gameLevelConfigJSON)
}

// parseGameConfig 解析游戏玩法配置
func parseGameConfig(config string) (gameConf []entityConf.GameConfig, err error) {
	if err := json.Unmarshal([]byte(config), &gameConf); err != nil {
		return nil, fmt.Errorf("反序列化失败：%s", err.Error())
	}
	logrus.Debugf("解析游戏玩法配置解析成功: %v", gameConf)
	return
}

// parseGameLevelConfig 解析游戏玩法配置
func parseGameLevelConfig(config string) (gameConf []entityConf.GameLevelConfig, err error) {
	if err := json.Unmarshal([]byte(config), &gameConf); err != nil {
		return nil, fmt.Errorf("反序列化失败：%s", err.Error())
	}
	logrus.Debugf("解析游戏场次配置解析成功: %v", gameConf)
	return
}

// InitGameConfig 初始化游戏配置
func InitGameConfig() {
	var err error
	// 游戏配置
	GameConf, err = loadGameConfig(20)
	if err != nil {
		logrus.WithError(err).Errorln("initGameConfig 获取游戏玩法配置失败")
	}
	// 场次配置
	LevelConf, err = loadGameLevelConfig(20)
	if err != nil {
		logrus.WithError(err).Errorln("initGameConfig 获取游戏场次配置失败")
	}
	return
}
