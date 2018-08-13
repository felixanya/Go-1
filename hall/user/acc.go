package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Sirupsen/logrus"

	"github.com/spf13/viper"
)

func init() {
	// 账号信息获取 url 默认值
	viper.SetDefault("account_info_url", "http://192.168.7.26:8086/mock/24/account/getByGuid")

	// 账号系统本产品 ID 默认值
	viper.SetDefault("prodcut_id", 9999)
}

// getAccountInfo 通过账号系统获取账号信息
func getAccountInfo(accID uint64) (map[string]interface{}, error) {
	// 获取账号信息的 URL
	url := viper.GetString("account_info_url")
	// 参数
	data, err := json.Marshal(map[string]interface{}{
		"productid": viper.GetInt64("product_id"),
		"guid":      accID,
		"is_pull":   true,
	})
	if err != nil {
		return nil, fmt.Errorf("序列化参数失败: %v", err)
	}

	// 发起 post 请求
	resp, err := http.Post(url, "application/json", bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("发起 post 请求失败: %v", err)
	}

	// 读取回复数据
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取回复数据失败: %v", err)
	}

	// 解析回复数据
	result := map[string]interface{}{}
	if err := json.Unmarshal(respData, &result); err != nil {
		return nil, fmt.Errorf("反序列化回复数据失败: %v", err)
	}

	// 解析错误码
	code, ok := result["code"].(float64)
	if !ok || int(code) != 0 {
		msg, _ := result["msg"]
		return nil, fmt.Errorf("账号系统返回错误, ok=%t, code=%d, msg=%s", ok, int(code), msg)
	}

	resultData, ok := result["data"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("data 解析失败")
	}
	return resultData, nil
}

// isYouke 判断账号是否为游客
func isYouke(accInfo map[string]interface{}) bool {
	// 有第三方账号，不是游客
	thirdType, ok := accInfo["third_type"].(float64)
	if ok && int(thirdType) != 0 {
		return false
	}

	// 设置了手机号，不是游客
	phone, ok := accInfo["phone"].(string)
	if ok && phone != "" {
		return false
	}

	return true
}

// generateNickName 根据账号系统返回的账号数据生成玩家昵称
func generateNickName(playerID uint64, accInfo map[string]interface{}) string {
	// log
	entry := logrus.WithField("player_id", playerID)

	// 游客昵称，失败时使用
	defaultNickName := fmt.Sprintf("游客%d", playerID)

	// 账号昵称
	accNickName, ok := accInfo["nickname"].(string)
	if !ok {
		entry.Warningln("获取账号昵称失败")
		return defaultNickName
	}

	// 不是游客使用账号昵称
	if !isYouke(accInfo) {
		return accNickName
	}

	// 都没有设置，生成游客昵称
	return defaultNickName
}

// generateAvartaURL 生成用户头像
func generateAvartaURL(playerID uint64, accInfo map[string]interface{}) string {
	// log
	entry := logrus.WithField("player_id", playerID)

	// 默认随机头像
	defaultAvatar := getRandomAvator()

	// 账号头像
	accAvatar, ok := accInfo["image"].(string)
	if !ok {
		entry.Warningln("获取账号头像失败")
		return defaultAvatar
	}

	// 不是游客使用账号头像
	if !isYouke(accInfo) {
		return accAvatar
	}

	// 都没有设置，使用默认随机头像
	return defaultAvatar
}

// getGender 获取性别， 1 女 2 男
func getGender(playerID uint64, accInfo map[string]interface{}) int {
	// log
	entry := logrus.WithField("player_id", playerID)

	// 默认性别： 女
	defaultGender := 1

	// 账号性别
	accGender, ok := accInfo["sex"].(float64)
	if !ok {
		entry.Warningln("获取账号头像失败")
		return defaultGender
	}

	// 不是游客使用账号性别
	if !isYouke(accInfo) {
		return int(accGender)
	}

	// 都没有设置，使用默认性别
	return defaultGender
}
