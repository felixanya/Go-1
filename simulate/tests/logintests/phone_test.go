package logintests

import (
	"fmt"
	"steve/client_pb/common"
	"steve/client_pb/hall"
	"steve/client_pb/msgid"
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

// sendDymcCode 发送手机验证码
func sendDymcCode(player interfaces.ClientPlayer, phone uint64, sendCase hall.AuthCodeSendScene) error {
	rsp := hall.AuthCodeRsp{}
	err := player.GetClient().Request(utils.CreateMsgHead(msgid.MsgID_AUTH_CODE_REQ), &hall.AuthCodeReq{
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

	assert.Nil(t, bindPhone(player, "15013701957", "abcd"))
}

func Test_ChangePhone(t *testing.T) {
	if !useAccountSystem { // 需要正式账号系统，不开启此测试
		return
	}
	player, err := utils.LoginPlayerYouke("abacdeea1acx")
	assert.Nil(t, err)
	assert.NotNil(t, player)

	assert.Nil(t, bindPhone(player, "10000000010", "abcd"))

	// 等待 冷却时间
	time.Sleep(time.Second * 10)
	assert.Nil(t, sendDymcCode(player, 10000000010, hall.AuthCodeSendScene_RESET_CELLPHONE))
	assert.Nil(t, sendDymcCode(player, 10000000011, hall.AuthCodeSendScene_RESET_CELLPHONE))

	rsp := hall.ChangePhoneRsp{}
	err = player.GetClient().Request(utils.CreateMsgHead(msgid.MsgID_CHANGE_PHONE_REQ), &hall.ChangePhoneReq{
		OldPhone:     proto.String("10000000010"),
		OldPhoneCode: proto.String("123456"),
		NewPhone:     proto.String("10000000011"),
		NewPhoneCode: proto.String("123456"),
	}, time.Second*5, uint32(msgid.MsgID_BIND_PHONE_RSP), &rsp)

	assert.Nil(t, err)
}
