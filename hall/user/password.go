package user

import (
	"steve/client_pb/common"
	"steve/client_pb/hall"
	"steve/client_pb/msgid"
	"steve/structs/exchanger"
	"steve/structs/proto/gate_rpc"

	"github.com/Sirupsen/logrus"
	"github.com/golang/protobuf/proto"
	"github.com/spf13/viper"
)

// ------------------------------------
// 用户密码修改，重置，校验
// ------------------------------------

func init() {
	// 修改密码 url
	viper.SetDefault("change_password_url", "http://192.168.7.26:18101/account/resetPwd")
	// 重置密码 url
	viper.SetDefault("reset_password_url", "http://192.168.7.26:18101/account/resetPwd")
	// 校验密码 url
	viper.SetDefault("check_password_url", "http://192.168.7.26:18101/account/checkPwd")
}

// HandleChangePasswordReq 处理修改密码请求
func HandleChangePasswordReq(playerID uint64, header *steve_proto_gaterpc.Header, req hall.ChangePasswordReq) (rspMsg []exchanger.ResponseMsg) {
	entry := logrus.WithFields(logrus.Fields{"player_id": playerID, "request": req.String()})

	response := hall.ChangePasswordRsp{
		Result: &common.Result{
			ErrCode: common.ErrCode_EC_FAIL.Enum(),
			ErrDesc: proto.String("修改失败"),
		},
	}
	rspMsg = []exchanger.ResponseMsg{
		exchanger.ResponseMsg{
			MsgID: uint32(msgid.MsgID_CHANGE_PASSWORD_RSP),
			Body:  &response,
		},
	}

	dbPlayer, err := dbPlayerGetter(playerID, "accountID")
	if err != nil {
		entry.WithError(err).Errorln("获取玩家信息失败")
		return
	}

	url := viper.GetString("change_password_url")
	httpResponseData := normalHTTPResponse{}
	err = requestJSONHTTP(url, map[string]interface{}{
		"productid":    viper.GetInt("product_id"),
		"guid":         dbPlayer.Accountid,
		"type":         2,
		"old_password": req.GetOldPasswd(),
		"new_password": req.GetNewPasswd(),
	}, &httpResponseData)
	if err != nil {
		entry.WithError(err).Errorln("请求失败")
		return
	}
	if httpResponseData.Code == 0 {
		response.Result.ErrCode = common.ErrCode_EC_SUCCESS.Enum()
	} else {
		response.Result.ErrDesc = proto.String(httpResponseData.Msg)
	}
	entry.WithField("response", response.String()).Debugln("修改密码处理完成")
	return
}

// HandleResetPasswordReq 处理重置密码请求
func HandleResetPasswordReq(playerID uint64, header *steve_proto_gaterpc.Header, req hall.ResetPasswordReq) (rspMsg []exchanger.ResponseMsg) {
	entry := logrus.WithFields(logrus.Fields{"player_id": playerID, "request": req.String()})

	response := hall.ResetPasswordRsp{
		Result: &common.Result{
			ErrCode: common.ErrCode_EC_FAIL.Enum(),
			ErrDesc: proto.String("重置失败"),
		},
	}
	rspMsg = []exchanger.ResponseMsg{
		exchanger.ResponseMsg{
			MsgID: uint32(msgid.MsgID_RESET_PASSWORD_RSP),
			Body:  &response,
		},
	}

	// dbPlayer, err := dbPlayerGetter(playerID, "accountID")
	// if err != nil {
	// 	entry.WithError(err).Errorln("获取玩家信息失败")
	// 	return
	// }

	url := viper.GetString("reset_password_url")
	httpResponseData := normalHTTPResponse{}
	err := requestJSONHTTP(url, map[string]interface{}{
		"productid": viper.GetInt("product_id"),
		// "guid":          dbPlayer.Accountid,
		"type":          1,
		"cellphone_num": req.GetPhone(),
		"dymc_code":     req.GetDymcCode(),
		"send_case":     int(hall.AuthCodeSendScene_RESET_PASSWORD),
		"new_password":  req.GetNewPasswd(),
	}, &httpResponseData)
	if err != nil {
		entry.WithError(err).Errorln("请求失败")
		return
	}
	if httpResponseData.Code == 0 {
		response.Result.ErrCode = common.ErrCode_EC_SUCCESS.Enum()
	} else {
		response.Result.ErrDesc = proto.String(httpResponseData.Msg)
	}
	entry.WithField("response", response.String()).Debugln("重置密码处理完成")
	return
}

// HandleCheckPasswordReq 校验密码请求
func HandleCheckPasswordReq(playerID uint64, header *steve_proto_gaterpc.Header, req hall.CheckPasswordReq) (rspMsg []exchanger.ResponseMsg) {
	entry := logrus.WithFields(logrus.Fields{"player_id": playerID, "request": req.String()})

	response := hall.CheckPasswordRsp{
		Result: &common.Result{
			ErrCode: common.ErrCode_EC_FAIL.Enum(),
			ErrDesc: proto.String("密码校验失败"),
		},
	}
	rspMsg = []exchanger.ResponseMsg{
		exchanger.ResponseMsg{
			MsgID: uint32(msgid.MsgID_CHECK_PASSWORD_RSP),
			Body:  &response,
		},
	}

	dbPlayer, err := dbPlayerGetter(playerID, "accountID")
	if err != nil {
		entry.WithError(err).Errorln("获取玩家信息失败")
		return
	}

	url := viper.GetString("check_password_url")
	httpResponseData := normalHTTPResponse{}
	err = requestJSONHTTP(url, map[string]interface{}{
		"productid": viper.GetInt("product_id"),
		"guid":      dbPlayer.Accountid,
		"password":  req.GetPassword(),
	}, &httpResponseData)
	if err != nil {
		entry.WithError(err).Errorln("请求失败")
		return
	}
	if httpResponseData.Code == 0 {
		response.Result.ErrCode = common.ErrCode_EC_SUCCESS.Enum()
		response.Result.ErrDesc = proto.String("")
	} else {
		response.Result.ErrDesc = proto.String(httpResponseData.Msg)
	}
	entry.WithField("response", response.String()).Debugln("校验密码处理完成")
	return
}
