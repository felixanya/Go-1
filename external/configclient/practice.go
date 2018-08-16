package configclient

import (
	"encoding/json"
	entityConf "steve/entity/config"
	"steve/entity/constant"
	"time"

	"github.com/Sirupsen/logrus"
)

func ParseToGameLevelConfigMap(jsonStr string) (conf []entityConf.GameLevelConfig) {
	if err := json.Unmarshal([]byte(jsonStr), &conf); err != nil {
		logrus.Errorf("游戏配置数据反序列化失败：%s", err.Error())
	}
	return
}

//获取救济金配置
func GetAlmsConfigMap() (conf []entityConf.AlmsConfig, err error) {
	almsStr, err := GetConfig("game", "alms")
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

// GetGameConfigMap 获取游戏配置信息
func GetGameConfigMap() (gameConf []entityConf.GameConfig, err error) {
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

// GetGameLevelConfigMap 获取游戏级别配置信息
func GetGameLevelConfigMap() (levelConf []entityConf.GameLevelConfig, err error) {
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
