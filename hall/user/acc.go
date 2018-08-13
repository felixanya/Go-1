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

// getAccountInfoRsp 获取账号信息回复结构
type getAccountInfoRsp struct {
	Code        int         `json:"code"`
	AccountInfo accountInfo `json:"data"`
	Msg         string      `json:"msg"`
}

// accountInfo 账号信息
type accountInfo struct {
	NickName        string `json:"nickname"`         // 昵称
	Image           string `json:"image"`            // 头像
	Sex             int    `json:"sex"`              //	性别，1男 2女
	Country         string `json:"country"`          // 国家
	Province        int    `json:"province"`         // 省 ID
	City            int    `json:"city"`             // 市 ID
	Channel         int    `json:"channel"`          // 渠道 ID
	Phone           string `json:"phone"`            // 电话
	ExtInfo         string `json:"ext_info"`         // 扩展信息(json格式)
	ThirdType       int    `json:"third_type"`       // 第三方账号类型， 0无 1微信
	ThirdInfo       string `json:"third_info"`       // 第三方账号用户扩展信息(json格式)
	RegisterProduct string `json:"register_product"` // 产品id，标识用户在哪个产品注册的
	RegisterTime    string `json:"register_time"`    // 注册时间
}

// getAccountInfo 通过账号系统获取账号信息
func getAccountInfo(accID uint64) (accountInfo, error) {
	accountInfo := accountInfo{}
	// 获取账号信息的 URL
	url := viper.GetString("account_info_url")
	// 参数
	data, err := json.Marshal(map[string]interface{}{
		"productid": viper.GetInt64("product_id"),
		"guid":      accID,
		"is_pull":   true,
	})
	if err != nil {
		return accountInfo, fmt.Errorf("序列化参数失败: %v", err)
	}
	// 发起 post 请求
	resp, err := http.Post(url, "application/json", bytes.NewReader(data))
	if err != nil {
		return accountInfo, fmt.Errorf("发起 post 请求失败: %v", err)
	}
	// 读取回复数据
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return accountInfo, fmt.Errorf("读取回复数据失败: %v", err)
	}
	// 解析回复数据
	var getAccountInfoRsp getAccountInfoRsp
	if err := json.Unmarshal(respData, &getAccountInfoRsp); err != nil {
		return accountInfo, fmt.Errorf("反序列化回复数据失败: %v", err)
	}
	// 错误码
	if getAccountInfoRsp.Code != 0 {
		return accountInfo, fmt.Errorf("账号系统返回错误, code=%d, msg=%s", getAccountInfoRsp.Code, getAccountInfoRsp.Msg)
	}
	return getAccountInfoRsp.AccountInfo, nil
}

// generateNickName 根据账号系统返回的账号数据生成玩家昵称
// 微信账号或者手机账号使用账号昵称作为玩家昵称
// 其他账号生成 游客+playerID 作为玩家昵称
func generateNickName(playerID uint64, accInfo *accountInfo) string {
	if accInfo.Phone != "" || accInfo.ThirdType == 1 {
		return accInfo.NickName
	}
	return fmt.Sprintf("游客%d", playerID)
}

// generateAvartaURL 生成用户头像
// 微信账号或者手机账号使用账号头像作为玩家头像
// 其他账号生成随机头像
func generateAvartaURL(playerID uint64, accInfo *accountInfo) string {
	if accInfo.Phone != "" || accInfo.ThirdType == 1 {
		return accInfo.Image
	}
	return getRandomAvator()
}

// generateGender 生成用户性别
// 微信账号或者手机账号使用账号性别作为玩家别
// 其他账号性别为女
func generateGender(playerID uint64, accInfo *accountInfo) int {
	if accInfo.Phone != "" || accInfo.ThirdType == 1 {
		return accInfo.Sex
	}
	return 1
}

// wxAccountInfo 微信账号信息
type wxAccountInfo struct {
	nickName string
	avatar   string
	gender   int
}

// getAccountWxInfo 获取账号信息的微信数据
// 返回 false 表示账号不是微信账号或者信息获取失败
func getAccountWxInfo(accID uint64) (wxAccountInfo, bool) {
	entry := logrus.WithField("account_id", accID)
	wxInfo := wxAccountInfo{}
	accInfo, err := getAccountInfo(accID)
	entry = entry.WithField("account_info", accInfo)
	if err != nil {
		entry.WithError(err).Warningln("获取账号信息失败")
		return wxInfo, false
	}
	if accInfo.ThirdType != 1 {
		entry.Debugln("不是微信账号")
		return wxInfo, false
	}
	wxInfo.avatar = accInfo.Image
	wxInfo.gender = accInfo.Sex
	wxInfo.nickName = accInfo.NickName
	return wxInfo, true
}
