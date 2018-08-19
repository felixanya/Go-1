package logintests

import (
	"steve/client_pb/hall"
	"steve/client_pb/msgid"
	"steve/simulate/utils"
	"testing"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
)

// -----------------------------------------------------------
// 此测试用于和后台联调，正式模拟测试时不开启
// -----------------------------------------------------------

// 修改密码测试
func Test_ChangePasswd(t *testing.T) {
	// useAccountSystem = true
	if !useAccountSystem {
		return
	}
	// 绑定手机后才可以修改密码
	player, err := utils.LoginPlayerYouke("abacdeeazz1acfxy")
	assert.Nil(t, err)
	assert.NotNil(t, player)

	assert.Nil(t, bindPhone(player, "10000000022", "abcd"))

	rsp := hall.ChangePasswordRsp{}
	err = player.GetClient().Request(utils.CreateMsgHead(msgid.MsgID_CHANGE_PASSWORD_REQ), &hall.ChangePasswordReq{
		OldPasswd: proto.String("abcd"),
		NewPasswd: proto.String("dcef"),
	}, time.Second*5, uint32(msgid.MsgID_CHANGE_PASSWORD_RSP), &rsp)
	assert.Nil(t, err)
}

// Test_VerifyPasswd 校验密码测试
func Test_VerifyPasswd(t *testing.T) {
	useAccountSystem = true
	if !useAccountSystem {
		return
	}
	player, err := utils.LoginPlayerYouke("abacdeeazz1acfxyd")
	assert.Nil(t, err)
	assert.NotNil(t, player)
	assert.Nil(t, bindPhone(player, "10000000014", "abcd"))

	rsp := hall.CheckPasswordRsp{}

	err = player.GetClient().Request(utils.CreateMsgHead(msgid.MsgID_CHECK_PASSWORD_REQ), &hall.CheckPasswordReq{
		Password: proto.String("abcd"),
	}, time.Second*5, uint32(msgid.MsgID_CHECK_PASSWORD_RSP), &rsp)

	assert.Nil(t, err)
	assert.Zerof(t, rsp.GetResult().GetErrCode(), rsp.GetResult().GetErrDesc())
}

// Test_ResetPassword 重置密码测试
// 由于发送验证码有冷却时间，使用已经存在的绑定手机的账号
func Test_ResetPassword(t *testing.T) {
	useAccountSystem = true
	if !useAccountSystem {
		return
	}
	player, err := utils.LoginPlayerYouke("abacdeeazz1acfxy")
	assert.Nil(t, err)
	assert.NotNil(t, player)

	assert.Nil(t, sendDymcCode(player, 10000000013, hall.AuthCodeSendScene_RESET_PASSWORD))

	rsp := hall.ResetPasswordRsp{}
	err = player.GetClient().Request(utils.CreateMsgHead(msgid.MsgID_RESET_PASSWORD_REQ), &hall.ResetPasswordReq{
		Phone:     proto.String("10000000013"),
		DymcCode:  proto.String("123456"),
		NewPasswd: proto.String("dddd"),
	}, time.Second*5, uint32(msgid.MsgID_RESET_PASSWORD_RSP), &rsp)

	assert.Nil(t, err)
	assert.Zerof(t, rsp.GetResult().GetErrCode(), rsp.GetResult().GetErrDesc())
}
