package packsack_prop

import (
	"encoding/json"
	"fmt"
	"steve/alms/packsack/packsack_utils"
	"steve/common/data/prop"
	"steve/entity/constant"

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
	if len(props) == 0 {
		logrus.Debugln("GetPlayerAllProps eq 0")
		return propInfos, nil
	}
	// 获取道具配置信息
	propConfig, err := GetPropInfos()
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
				pType := int32(speedy)
				if pConfig.PropID <= 5 {
					pType = int32(interactive)
				}
				p := PropPacksackInfo{
					PropID:    pConfig.PropID,
					PropName:  pConfig.PropName,
					PropType:  pType, //道具类型
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

//GetPropInfos 获取道具信息
func GetPropInfos() ([]constant.PropAttr, error) {
	// redis
	propString, err := packsack_utils.GetPropConfigRedis()
	if err == nil && propString != "" {
		return JSONToPropinfo(propString)
	}
	// 获取道具配置信息
	propConfig, err := prop.GetPropsConfig()
	if err != nil {
		logrus.WithError(err).Debugln("GetPropsConfig")
		return propConfig, err
	}
	if len(propConfig) == 0 {
		logrus.Debugln("GetPropInfos eq 0")
		return propConfig, nil
	}
	// 保存
	propString, err = propinfosToJSON(propConfig)
	if err == nil {
		err = packsack_utils.SetPropConfigRedis(propString)
	}
	if err != nil {
		logrus.WithError(err).Debugln("GetPropInfos propinfosToJSON")
	}
	return propConfig, nil
}

// InitPropInfo 初始化道具信息
func InitPropInfo() error {
	// 获取道具配置信息
	propConfig, err := prop.GetPropsConfig()
	if err != nil {
		logrus.WithError(err).Debugln("获取道具配置信息")
		return err
	}
	// 保存
	propString, err := propinfosToJSON(propConfig)
	if err == nil {
		err = packsack_utils.SetPropConfigRedis(propString)
	}
	if err != nil {
		logrus.WithError(err).Debugln("InitPropInfo GetPropInfos propinfosToJSON")
	}
	return err
}

// propinfosToJSON propinfo to JSON
func propinfosToJSON(propInfo []constant.PropAttr) (string, error) {
	if len(propInfo) == 0 {
		logrus.Debugln("propInfos eq 0")
		return "", nil
	}
	bytes, err := json.Marshal(propInfo)
	return string(bytes), err
}

// JSONToPropinfo JSON TO propinfo
func JSONToPropinfo(propInfoStr string) ([]constant.PropAttr, error) {
	if propInfoStr == "" {
		logrus.Debugln("propInfoStr eq 0")
		return nil, nil
	}
	propInfos := []constant.PropAttr{}
	err := json.Unmarshal([]byte(propInfoStr), &propInfos)
	return propInfos, err
}
