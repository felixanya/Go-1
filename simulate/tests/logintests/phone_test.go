package logintests

import (
	"fmt"
	"steve/client_pb/common"
	"steve/client_pb/hall"
	"steve/client_pb/msgid"
	"steve/simulate/config"
	"steve/simulate/connect"
	"steve/simulate/interfaces"
	"steve/simulate/utils"
	"strconv"
	"testing"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
)

// -----------------------------------------------------------
// 此测试用于和后台联调，正式模拟测试时不开启
// -----------------------------------------------------------

var useAccountSystem = false

func sendDymcCodeByClient(client interfaces.Client, phone uint64, sendCase hall.AuthCodeSendScene) error {
	rsp := hall.AuthCodeRsp{}
	err := client.Request(utils.CreateMsgHead(msgid.MsgID_AUTH_CODE_REQ), &hall.AuthCodeReq{
		CellphoneNum: proto.Uint64(phone),
		SendCase:     sendCase.Enum(),
	}, time.Second*5, uint32(msgid.MsgID_AUTH_CODE_RSP), &rsp)
	if err != nil {
		return fmt.Errorf("请求失败：%s", err.Error())
	}
	if rsp.GetErrorCode() != 0 {
		return fmt.Errorf("发送失败:[%d] %s", rsp.GetErrorCode(), rsp.GetErrorMsg())
	}
	return nil
}

// sendDymcCode 发送手机验证码
func sendDymcCode(player interfaces.ClientPlayer, phone uint64, sendCase hall.AuthCodeSendScene) error {
	return sendDymcCodeByClient(player.GetClient(), phone, sendCase)
}

// checkDymcCode 检查验证码
func checkDymcCode(client interfaces.Client, phone string, code string, sendCase hall.AuthCodeSendScene) error {
	rsp := hall.CheckAuthCodeRsp{}
	err := client.Request(utils.CreateMsgHead(msgid.MsgID_CHECK_AUTH_CODE_REQ), &hall.CheckAuthCodeReq{
		Phone:    proto.String(phone),
		SendCase: sendCase.Enum(),
		Code:     proto.String(code),
	}, time.Second*5, uint32(msgid.MsgID_CHECK_AUTH_CODE_RSP), &rsp)
	if err != nil {
		return fmt.Errorf("请求失败：%s", err.Error())
	}
	result := rsp.GetResult()
	if result.GetErrCode() != common.ErrCode_EC_SUCCESS {
		return fmt.Errorf("发送失败:[%d] %s", result.GetErrCode(), result.GetErrDesc())
	}
	return nil
}

// Test_CheckDymcCode 校验验证码
func Test_CheckDymcCode(t *testing.T) {
	//useAccountSystem = true
	if !useAccountSystem {
		return
	}
	client := connect.NewTestClient(config.GetGatewayServerAddr(), config.GetClientVersion())
	assert.NotNil(t, client)
	assert.Nil(t, sendDymcCodeByClient(client, 18565877566, hall.AuthCodeSendScene_BIND_PHONE))
	assert.Nil(t, checkDymcCode(client, "18565877566", "123456", hall.AuthCodeSendScene_BIND_PHONE))
}

// Test_SendDymcCode 发送验证码
func Test_SendDymcCode(t *testing.T) {
	if !useAccountSystem { // 需要正式账号系统，不开启此测试
		return
	}
	player, err := utils.LoginPlayerYouke("abacdee")
	assert.Nil(t, err)
	assert.NotNil(t, player)

	assert.Nil(t, sendDymcCode(player, 18565877566, hall.AuthCodeSendScene_BIND_PHONE))
}

func bindPhone(player interfaces.ClientPlayer, phone, passwd string) error {
	rsp := hall.BindPhoneRsp{}

	iphone, _ := strconv.ParseUint(phone, 10, 64)
	if err := sendDymcCode(player, iphone, hall.AuthCodeSendScene_BIND_PHONE); err != nil {
		return fmt.Errorf("发送验证码失败：%s", err.Error())
	}
	err := player.GetClient().Request(utils.CreateMsgHead(msgid.MsgID_BIND_PHONE_REQ), &hall.BindPhoneReq{
		Phone:    proto.String(phone),
		DymcCode: proto.String("123456"),
		Passwd:   proto.String(passwd),
	}, time.Second*5, uint32(msgid.MsgID_BIND_PHONE_RSP), &rsp)
	if err != nil {
		return fmt.Errorf("请求失败：%s", err.Error())
	}
	if rsp.GetResult().GetErrCode() != common.ErrCode_EC_SUCCESS {
		return fmt.Errorf("绑定失败:%s", rsp.GetResult().GetErrDesc())
	}
	return nil
}

func Test_BindPhone(t *testing.T) {
	if !useAccountSystem { // 需要正式账号系统，不开启此测试
		return
	}
	player, err := utils.LoginPlayerYouke("abacdeeaa")
	assert.Nil(t, err)
	assert.NotNil(t, player)

	assert.Nil(t, bindPhone(player, "10000000020", "abcd"))
}

func Test_ChangePhone(t *testing.T) {
	if !useAccountSystem { // 需要正式账号系统，不开启此测试
		return
	}
	// 由于验证码发送间隔时间有限制，使用 Test_BindPhone 中的账号
	player, err := utils.LoginPlayerYouke("abacdeeaa")
	assert.Nil(t, err)
	assert.NotNil(t, player)

	assert.Nil(t, sendDymcCode(player, 10000000020, hall.AuthCodeSendScene_RESET_CELLPHONE))
	assert.Nil(t, sendDymcCode(player, 10000000013, hall.AuthCodeSendScene_RESET_CELLPHONE))

	rsp := hall.ChangePhoneRsp{}
	err = player.GetClient().Request(utils.CreateMsgHead(msgid.MsgID_CHANGE_PHONE_REQ), &hall.ChangePhoneReq{
		OldPhone:     proto.String("10000000020"),
		OldPhoneCode: proto.String("123456"),
		NewPhone:     proto.String("10000000013"),
		NewPhoneCode: proto.String("123456"),
	}, time.Second*5, uint32(msgid.MsgID_CHANGE_PHONE_RSP), &rsp)

	assert.Nil(t, err)
}
