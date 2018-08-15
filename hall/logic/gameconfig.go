package logic

import (
	"context"
	entityConf "steve/entity/config"
	"steve/external/configclient"

	"github.com/Sirupsen/logrus"
)

// GameConf 游戏配置
var GameConf []entityConf.GameConfig

// LevelConf 场次配置
var LevelConf []entityConf.GameLevelConfig

// RoleConfig 角色配置
var RoleConfig []entityConf.RoleInitConfig

// InitGameConfig 初始化游戏配置
func InitGameConfig(ctx context.Context) {
	var err error
	for {
		select {
		case <-ctx.Done():
			logrus.Debugf("hall服启动加载的游戏玩法，GameConf:(%v)\n，LevelConf：（%v）", GameConf, LevelConf)
			return
		default:
			GameConf, err = configclient.GetGameConfigMap()
			if err != nil {
				continue
			}
			LevelConf, err = configclient.GetGameLevelConfigMap()
			if err != nil {
				continue
			}
			logrus.Debugf("hall服启动加载的游戏玩法，GameConf:(%v)\n，LevelConf：（%v）", GameConf, LevelConf)
			return
		}
	}
}

// InitRoleConfig 初始化角色配置
func InitRoleConfig() {
	var err error
	RoleConfig, err = configclient.GetRoleInitConfigMap()
	if err != nil {
		logrus.Debugf("hall服启动加载角色失败，RoleConfig:(%v)，error:(%s)", RoleConfig, err.Error())
	}
	logrus.Debugf("hall服启动加载角色配置，RoleConfig:(%v)", RoleConfig)

	return
}
