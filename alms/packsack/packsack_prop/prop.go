package packsack_prop

import (
	"fmt"
	"steve/common/data/prop"

	"github.com/Sirupsen/logrus"
)

//PropPacksackInfo 背包道具信息
type PropPacksackInfo struct {
	PropID    int32  // 道具ID  前端识别的ID与PropTypeID相同
	PropName  string // 道具名称
	PropType  int32  // 道具类型
	Describe  string // 道具描述
	PropCount int64  // 道具数量
}

const (
	interactive = 1 //非便捷道具
	speedy      = 2 //便捷道具
)

//GetPlayerPropInfoAll 获取玩家的所有道具信息
func GetPlayerPropInfoAll(playerID uint64) ([]PropPacksackInfo, error) {
	propInfos := make([]PropPacksackInfo, 0)
	props, err := prop.GetPlayerAllProps(playerID)
	if err != nil {
		return propInfos, fmt.Errorf("get player all prop 失败 err(%v)", err)
	}
	// 获取道具配置信息
	propConfig, err := prop.GetPropsConfig()
	if err != nil {
		return propInfos, err
	}
	for _, prop := range props {
		if prop.Count == 0 {
			continue
		}
		flag := true
		for _, pConfig := range propConfig {
			if pConfig.PropID == prop.PropID {
				p := PropPacksackInfo{
					PropID:    pConfig.PropID,
					PropName:  pConfig.PropName,
					PropType:  pConfig.Type, //道具类型
					PropCount: prop.Count,
					Describe:  pConfig.Describe,
				}
				propInfos = append(propInfos, p)
				flag = false
				break
			}
		}
		if flag {
			logrus.Debugf(fmt.Sprintf("失败的PropID(%d)", prop.PropID))
		}
	}
	return propInfos, nil
}
