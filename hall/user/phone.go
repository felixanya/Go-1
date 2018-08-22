package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"steve/client_pb/common"
	"steve/client_pb/hall"
	"steve/client_pb/msgid"
	"steve/entity/db"
	"steve/external/configclient"
	"steve/external/goldclient"
	"steve/hall/data"
	"steve/structs/exchanger"
	"steve/structs/proto/gate_rpc"
	"sync"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/golang/protobuf/proto"
	"github.com/spf13/viper"
)

// bindPhoneRewardConfig 绑定手机奖励配置
type bindPhoneRewardConfig struct {
	MoneyType int    `json:"type"`
	Num       uint64 `json:"num"`
}

// defaultBindPhoneRewardConfig 默认绑定手机奖励，不可以直接访问，要使用 getBindPhoneRewardConfig 函数
// 使用指针，可以原子性的修改
var defaultBindPhoneRewardConfig = &bindPhoneRewardConfig{
	MoneyType: int(common.MoneyType_MT_GOLDINGOT),
	Num:       5,
}

// loadBindPhoneRewardConfig 加载绑定手机奖励配置
var loadBindPhoneRewardConfig = func(maxRetry int) {
	cfg, err := configclient.GetConfigUntilSucc("bindphone", "reward", maxRetry, time.Second*3)
	if err != nil {
		logrus.WithError(err).Errorln("配置获取失败")
		return
	}
	parseBindPhoneRewardConfig(cfg)
}

// 解析绑定手机奖励配置
func parseBindPhoneRewardConfig(cfg string) {
	var rewardConfig bindPhoneRewardConfig
	if err := json.Unmarshal([]byte(cfg), &rewardConfig); err != nil {
		logrus.WithError(err).Errorln("反序列化配置失败")
		return
	}
	defaultBindPhoneRewardConfig = &rewardConfig
	logrus.Infof("绑定手机配置解析成功:%#v", rewardConfig)
}

var loadBindPhoneRewardConfigOnce = sync.Once{}

var dbPlayerGetter = data.GetPlayerInfo
var dbPlayerSetter = data.SetPlayerFields

// getBindPhoneRewardConfig 获取手机绑定奖励配置
func getBindPhoneRewardConfig() *bindPhoneRewardConfig {
	return defaultBindPhoneRewardConfig
}

func init() {
	// 发送验证码 url
	viper.SetDefault("send_code_url", "http://192.168.7.26:18101/account/sendCode")
	// 校验验证码 url
	viper.SetDefault("check_code_url", "http://192.168.7.26:18101/account/checkCode")
	// 绑定手机 url
	viper.SetDefault("bind_phone_url", "http://192.168.7.26:18101/account/bindPhone")
	// 修改手机 url
	viper.SetDefault("change_phone_url", "http://192.168.7.26:18101/account/resetPhone")

	// 绑定手机配置修改订阅
	configclient.SubConfigChange("bindphone", "reward", func(key, subkey, val string) error {
		parseBindPhoneRewardConfig(val)
		return nil
	})
}

// normalHTTPResponse 常规 http 回复数据结构
type normalHTTPResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// requestJSONHTTP 将 requestData 序列化成 json， 请求 url，并且读取回复，然后反序列化到 responseData
func requestJSONHTTP(url string, requestData interface{}, responseData interface{}) error {
	httpRequestData, err := json.Marshal(requestData)
	if err != nil {
		return fmt.Errorf("序列化请求数据失败: err=%s", err.Error())
	}
	httpResponse, err := http.Post(url, "application/json", bytes.NewReader(httpRequestData))
	if err != nil || httpResponse.StatusCode != 200 {
		return fmt.Errorf("发起 HTTP 请求失败， err=%s, code=%d", err.Error(), httpResponse.StatusCode)
	}
	httpResponseRaw, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return fmt.Errorf("读取 HTTP 回复失败: err=%s", err.Error())
	}
	if err := json.Unmarshal(httpResponseRaw, responseData); err != nil {
		fmt.Println(string(httpResponseRaw))
		return fmt.Errorf("反序列化失败: err=%s", err.Error())
	}
	logrus.WithFields(logrus.Fields{
		"url":      url,
		"request":  string(httpRequestData),
		"response": string(httpResponseRaw),
	}).Debugln("HTTP 请求成功")
	return nil
}

// HandleSendAuthCodeReq 发送验证码消息处理器
func HandleSendAuthCodeReq(playerID uint64, header *steve_proto_gaterpc.Header, req hall.AuthCodeReq) (rspMsg []exchanger.ResponseMsg) {
	response := hall.AuthCodeRsp{
		ErrorCode: proto.Uint64(uint64(common.ErrCode_EC_FAIL)),
		ErrorMsg:  proto.String("发送失败"),
	}
	rspMsg = []exchanger.ResponseMsg{
		exchanger.ResponseMsg{
			MsgID: uint32(msgid.MsgID_AUTH_CODE_RSP),
			Body:  &response,
		},
	}

	entry := logrus.WithFields(logrus.Fields{
		"player_id": playerID,
		"phone":     req.GetCellphoneNum(),
		"send_case": req.GetSendCase(),
	})

	url := viper.GetString("send_code_url")
	httpResponseData := normalHTTPResponse{}
	err := requestJSONHTTP(url, map[string]interface{}{
		"productid":     viper.GetInt("product_id"),
		"cellphone_num": fmt.Sprintf("%d", req.GetCellphoneNum()),
		"send_case":     int(req.GetSendCase()),
	}, &httpResponseData)
	if err != nil {
		entry.WithError(err).Errorln("请求失败")
		return
	}
	if httpResponseData.Code == 0 {
		response.ErrorCode = proto.Uint64(0)
		response.ErrorMsg = proto.String("发送成功")
	} else {
		response.ErrorMsg = proto.String(httpResponseData.Msg)
	}
	entry.WithField("http_response", httpResponseData).Debugln("发送验证码完成")
	return
}

// HandleCheckAuthCodeReq 验证码校验处理器
func HandleCheckAuthCodeReq(playerID uint64, header *steve_proto_gaterpc.Header, req hall.CheckAuthCodeReq) (rspMsg []exchanger.ResponseMsg) {
	result := common.Result{
		ErrCode: common.ErrCode_EC_FAIL.Enum(),
		ErrDesc: proto.String("验证码校验失败"),
	}
	response := hall.CheckAuthCodeRsp{
		Result: &result,
	}
	rspMsg = []exchanger.ResponseMsg{
		exchanger.ResponseMsg{
			MsgID: uint32(msgid.MsgID_CHECK_AUTH_CODE_RSP),
			Body:  &response,
		},
	}
	entry := logrus.WithFields(logrus.Fields{
		"player_id": playerID,
		"request":   req.String(),
	})

	url := viper.GetString("check_code_url")
	httpResponseData := normalHTTPResponse{}
	err := requestJSONHTTP(url, map[string]interface{}{
		"productid":     viper.GetInt("product_id"),
		"cellphone_num": req.GetPhone(),
		"send_case":     int(req.GetSendCase()),
		"dymc_code":     req.GetCode(),
	}, &httpResponseData)
	if err != nil {
		entry.WithError(err).Errorln("HTTP 请求失败")
		return
	}
	if httpResponseData.Code == 0 {
		result.ErrCode = common.ErrCode_EC_SUCCESS.Enum()
	} else {
		result.ErrDesc = proto.String(httpResponseData.Msg)
	}
	entry.WithField("response", result.String()).Debugln("验证码校验完成")
	return
}

// HandleGetBindphoneRewardInfoReq 获取绑定手机可获得奖励请求
func HandleGetBindphoneRewardInfoReq(playerID uint64, header *steve_proto_gaterpc.Header, req hall.GetBindPhoneRewardInfoReq) (rspMsg []exchanger.ResponseMsg) {
	cfg := getBindPhoneRewardConfig()
	response := hall.GetBindPhoneRewardInfoRsp{
		Reward: &common.Money{
			MoneyType: common.MoneyType(cfg.MoneyType).Enum(),
			MoneyNum:  proto.Uint64(cfg.Num),
		},
	}
	rspMsg = []exchanger.ResponseMsg{
		exchanger.ResponseMsg{
			MsgID: uint32(msgid.MsgID_GET_BIND_PHONE_REWARD_RSP),
			Body:  &response,
		},
	}
	return
}

var goldAdder = goldclient.AddGold

// HandleBindPhoneReq 处理绑定手机请求
func HandleBindPhoneReq(playerID uint64, header *steve_proto_gaterpc.Header, req hall.BindPhoneReq) (rspMsg []exchanger.ResponseMsg) {
	loadBindPhoneRewardConfigOnce.Do(func() {
		loadBindPhoneRewardConfig(5) // 重试 5 次， 加载绑定手机奖励配置
	})

	response := hall.BindPhoneRsp{
		Result: &common.Result{
			ErrCode: common.ErrCode_EC_FAIL.Enum(),
			ErrDesc: proto.String("绑定失败"),
		},
		Reward:   &common.Money{},
		NewMoney: &common.Money{},
	}
	rspMsg = []exchanger.ResponseMsg{
		exchanger.ResponseMsg{
			MsgID: uint32(msgid.MsgID_BIND_PHONE_RSP),
			Body:  &response,
		},
	}
	entry := logrus.WithFields(logrus.Fields{"player_id": playerID, "request": req.String()})

	dbPlayer, err := dbPlayerGetter(playerID, "phone", "accountID")
	if err != nil {
		entry.WithError(err).Errorln("获取玩家信息失败")
		return
	}
	if dbPlayer.Phone != "" {
		entry.WithField("phone", dbPlayer.Phone).Debugln("已经绑定手机")
		response.Result.ErrDesc = proto.String("已绑定")
		return
	}
	httpResponse := normalHTTPResponse{}
	err = requestJSONHTTP(viper.GetString("bind_phone_url"), map[string]interface{}{
		"productid":        viper.GetInt("product_id"),
		"cellphone_number": req.GetPhone(),
		"dymc_code":        req.GetDymcCode(),
		"send_case":        int(hall.AuthCodeSendScene_BIND_PHONE),
		"guid":             dbPlayer.Accountid,
		"password":         req.GetPasswd(),
	}, &httpResponse)
	if err != nil {
		entry.WithError(err).Errorln("HTTP 请求失败")
		return
	}
	if httpResponse.Code != 0 {
		entry.WithFields(logrus.Fields{"code": httpResponse.Code, "msg": httpResponse.Msg}).Infoln("手机绑定失败")
		response.Result.ErrDesc = proto.String(httpResponse.Msg)
		return
	}
	dbPlayer.Phone = req.GetPhone()
	if err := dbPlayerSetter(playerID, []string{"phone"}, dbPlayer); err != nil {
		entry.WithError(err).Errorln("更新手机号到数据库失败")
		// 不返回，照常发奖励
	}
	rewardCfg := getBindPhoneRewardConfig()
	newMoney, err := goldAdder(playerID, int16(rewardCfg.MoneyType), int64(rewardCfg.Num), 0, 0, 0, 0)
	if err != nil {
		entry.WithError(err).Errorln("奖励发放失败")
		return
	}
	response.Result.ErrCode = common.ErrCode_EC_SUCCESS.Enum()
	response.Result.ErrDesc = proto.String("")
	response.Reward.MoneyNum = proto.Uint64(rewardCfg.Num)
	response.Reward.MoneyType = common.MoneyType(rewardCfg.MoneyType).Enum()
	response.NewMoney.MoneyNum = proto.Uint64(uint64(newMoney))
	response.NewMoney.MoneyType = common.MoneyType(rewardCfg.MoneyType).Enum()
	entry.Debugln("手机绑定成功")
	return
}

// HandleChangePhoneReq 处理修改手机号请求
func HandleChangePhoneReq(playerID uint64, header *steve_proto_gaterpc.Header, req hall.ChangePhoneReq) (rspMsg []exchanger.ResponseMsg) {
	response := hall.ChangePhoneRsp{
		Result: &common.Result{
			ErrCode: common.ErrCode_EC_FAIL.Enum(),
			ErrDesc: proto.String("绑定失败"),
		},
	}
	rspMsg = []exchanger.ResponseMsg{
		exchanger.ResponseMsg{
			MsgID: uint32(msgid.MsgID_CHANGE_PHONE_RSP),
			Body:  &response,
		},
	}
	entry := logrus.WithFields(logrus.Fields{"player_id": playerID, "request": req.String()})
	dbPlayer, err := dbPlayerGetter(playerID, "phone", "accountID")
	if err != nil {
		entry.WithError(err).Errorln("获取玩家信息失败")
		return
	}
	entry = entry.WithField("old_phone", dbPlayer.Phone)
	if dbPlayer.Phone == "" {
		entry.Debugln("未绑定手机")
		response.Result.ErrDesc = proto.String("还未绑定手机")
		return
	}
	if dbPlayer.Phone == req.GetNewPhone() {
		entry.Debugln("新旧手机号相同")
		response.Result.ErrDesc = proto.String("新手机号与旧手机号不能相同")
		return
	}

	httpResponse := normalHTTPResponse{}
	err = requestJSONHTTP(viper.GetString("change_phone_url"), map[string]interface{}{
		"productid":      viper.GetInt("product_id"),
		"old_phone_code": req.GetOldPhoneCode(),
		"new_phone_code": req.GetNewPhoneCode(),
		"send_case":      int(hall.AuthCodeSendScene_RESET_CELLPHONE),
		"guid":           dbPlayer.Accountid,
		"new_phone_num":  req.GetNewPhone(),
	}, &httpResponse)
	if err != nil {
		entry.WithError(err).Errorln("HTTP 请求失败")
		return
	}

	if httpResponse.Code != 0 {
		entry.WithFields(logrus.Fields{"code": httpResponse.Code, "msg": httpResponse.Msg}).Infoln("修改手机失败")
		response.Result.ErrDesc = proto.String(httpResponse.Msg)
		return
	}
	response.Result.ErrCode = common.ErrCode_EC_SUCCESS.Enum()
	response.Result.ErrDesc = proto.String("")

	// 更新到数据库
	if err := dbPlayerSetter(playerID, []string{"phone"}, &db.TPlayer{Phone: req.GetNewPhone()}); err != nil {
		entry.WithError(err).Errorln("更新手机号到数据库失败")
	}
	entry.Debugln("手机更换成功")
	return
}
