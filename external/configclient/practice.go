package configclient

import (
	"encoding/json"
	entityConf "steve/entity/config"
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
func GetRoleInitConfigMap() (roleConf []entityConf.RoleInitConfig, err error) {
	roleConfStr, err := GetConfigUntilSucc("role", "init", reTryCount, reTryTime)
	if err != nil {
		logrus.WithError(err).Errorln("获取角色初始属性配置失败")
		return nil, err
	}
	if err := json.Unmarshal([]byte(roleConfStr), &roleConf); err != nil {
		logrus.WithError(err).Errorf("游戏级别配置数据反序列化失败：%s", err.Error())
		return nil, err
	}

	for _, config := range roleConf {
		config.InitItem()
	}
	return
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

// GetGameConfigMap 获取游戏配置信息
func GetGameConfigMap() (gameConf []entityConf.GameConfig, err error) {
	gameStr, err := GetConfigUntilSucc("game", "config", reTryCount, reTryTime)
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
	levelStr, err := GetConfigUntilSucc("game", "levelconfig", reTryCount, reTryTime)
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
