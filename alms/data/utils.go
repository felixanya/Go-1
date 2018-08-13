package data

import (
	"encoding/json"
	"strconv"

	"github.com/Sirupsen/logrus"
)

//InterToint64 接口转int64
func InterToint64(param interface{}) int64 {
	if param == nil {
		return 0
	}
	str := param.(string)
	result, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		logrus.WithFields(logrus.Fields{"func_name": "InterToint64",
			"param": param}).Infoln("InterToint64失败")
		return 0
	}
	return result
}

//检验redis 返回的 值
func checkMapStringInterface(m map[string]interface{}, checkString []string) bool {
	if len(m) != len(checkString) {
		return false
	}
	for _, str := range checkString {
		switch m[str].(type) {
		case string:
			if str == GameLeveConfigs && len(JSONToGameLeveConfig(m[str].(string))) <= 0 {
				return false
			}
		case int64:
			if InterToint64(m[str]) <= 0 {
				return false
			}
		default:
			logrus.WithFields(logrus.Fields{"func_name": "checkMapStringInterface",
				"m[str]": m[str]}).Infoln("检验redis 返回的")
			return false
		}
	}
	return true
}

// GameLeveConfigToJSON 游戏场次配置 转 JSON
func GameLeveConfigToJSON(gemeLeveOK []*GameLeveConfig) string {
	if gemeLeveOK == nil {
		return ""
	}
	str, err := json.Marshal(gemeLeveOK)
	if err != nil {
		logrus.WithFields(logrus.Fields{"func_name": "GameLeveConfigToJSON",
			"gemeLeveOK": gemeLeveOK}).Infoln("游戏场次配置 转 JSON失败")
	}
	return string(str)
}

// JSONToGameLeveConfig JSON 转 游戏场次配置
func JSONToGameLeveConfig(gemeLeveOKJSON string) []*GameLeveConfig {
	gemeLeveOK := []*GameLeveConfig{}
	if gemeLeveOKJSON == "" {
		return gemeLeveOK
	}
	globyte := []byte(gemeLeveOKJSON)
	if err := json.Unmarshal(globyte, &gemeLeveOK); err != nil {
		logrus.WithError(err).WithFields(logrus.Fields{"func_name": "JSONToGameLeveConfig",
			"gemeLeveOKJSON": gemeLeveOKJSON}).Infoln("JSON 转 游戏场次配置失败")
	}
	return gemeLeveOK
}
