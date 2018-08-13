package packsack

import (
	"encoding/json"
	propfunc "steve/common/data/prop"
	"steve/entity/constant"
	"steve/external/configclient"

	"github.com/Sirupsen/logrus"
)

//PropPacksackInfo 背包道具信息
type PropPacksackInfo struct {
	PropID    int32  // 道具ID
	PropName  string // 道具名称
	PropType  int    // 道具类型（便捷道具2，非便捷道具1）
	Describe  string // 道具描述
	PropCount int64  // 道具数量
}

const (
	interactive = 1 //非便捷道具
	speedy      = 2 //便捷道具
)

// GetPropPacksConfig 获取道具配置信息
func GetPropPacksConfig(propSubKey string) ([]constant.PropAttr, error) {
	var err error
	// 现在直接从数据库获取，后面改为先从redis获取；订阅更新消息，更新时删掉redis数据 TODO
	logrus.Debugf("GetPropPacksConfig PropKey:(%v),PropSubKey:(%v)", constant.PropKey, propSubKey)

	val, err := configclient.GetConfig(constant.PropKey, propSubKey)

	if err != nil {
		return nil, err
	}

	var propConfig []constant.PropAttr
	err = json.Unmarshal([]byte(val), &propConfig)
	if err != nil {
		return nil, err
	}

	return propConfig, err
}

func getPlayerProp(playerID uint64, propType int, propConfig []constant.PropAttr) (props []PropPacksackInfo, err error) {
	// 获取玩家的道具信息
	props = make([]PropPacksackInfo, len(propConfig))
	for index, attr := range propConfig {
		pcount, _ := propfunc.GetPlayerOneProp(playerID, attr.PropID) //获取该道具的数量
		props[index] = PropPacksackInfo{
			PropID:    attr.PropID,
			PropName:  attr.PropName,
			PropCount: pcount.Count,
			PropType:  propType, //便捷2 还是 非便捷道具1
			Describe:  "道具描述",
		}
	}
	return props, nil
}

//GetPlayerPropInfoAll 获取玩家所有非便捷道具信息
func GetPlayerPropInfoAll(playerID uint64) (props []PropPacksackInfo, err error) {
	// 获取道具配置信息
	propConfig, err := GetPropPacksConfig(constant.PropSubKey)
	if err != nil {
		return nil, err
	}
	return getPlayerProp(playerID, interactive, propConfig)
}

//GetPlayerSpeedyPropInfoAll 获取玩家所有便捷道具信息
func GetPlayerSpeedyPropInfoAll(playerID uint64) (props []PropPacksackInfo, err error) {
	// 获取道具配置信息
	propSpeedyConfig, err := GetPropPacksConfig(constant.PropSubKeySpeedy)
	if err != nil {
		return nil, err
	}
	return getPlayerProp(playerID, speedy, propSpeedyConfig)
}
