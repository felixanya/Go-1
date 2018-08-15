package logintests

import (
	"fmt"
	"steve/client_pb/common"
	"steve/client_pb/hall"
	"steve/client_pb/msgid"
	"steve/simulate/interfaces"
	"steve/simulate/utils"
	"testing"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
)

// 测试使用游客登录
func Test_YoukeLogin(t *testing.T) {
	if true { // 模拟测试时不开启此测试
		return
	}
	player, err := utils.LoginPlayerYouke("abcdeeadfasfe")
	assert.Nil(t, err)
	assert.NotNil(t, player)
}

// Test_SendDymcCode 发送验证码
func Test_SendDymcCode(t *testing.T) {
	player, err := utils.LoginPlayerYouke("abacdee")
	assert.Nil(t, err)
	assert.NotNil(t, player)

	rsp := hall.AuthCodeRsp{}
	player.GetClient().Request(utils.CreateMsgHead(msgid.MsgID_AUTH_CODE_REQ), &hall.AuthCodeReq{
		CellphoneNum: proto.Uint64(18565877566),
		SendCase:     hall.AuthCodeSendScene_BIND_PHONE.Enum(),
	}, time.Second*5, uint32(msgid.MsgID_AUTH_CODE_RSP), &rsp)

	assert.Zero(t, rsp.GetErrorCode())
}

func bindPhone(player interfaces.ClientPlayer) error {
	rsp := hall.BindPhoneRsp{}
	err := player.GetClient().Request(utils.CreateMsgHead(msgid.MsgID_BIND_PHONE_REQ), &hall.BindPhoneReq{
		Phone:    proto.String("13564581234"),
		DymcCode: proto.String("123456"),
		Passwd:   proto.String("abcd"),
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
	player, err := utils.LoginPlayerYouke("abacdee")
	assert.Nil(t, err)
	assert.NotNil(t, player)

	assert.Nil(t, bindPhone(player))
}

func Test_ChangePhone(t *testing.T) {
	player, err := utils.LoginPlayerYouke("abacdefg")
	assert.Nil(t, err)
	assert.NotNil(t, player)

	assert.Nil(t, bindPhone(player))
	rsp := hall.ChangePhoneRsp{}
	err = player.GetClient().Request(utils.CreateMsgHead(msgid.MsgID_CHANGE_PHONE_REQ), &hall.ChangePhoneReq{
		OldPhone:     proto.String("13564581234"),
		OldPhoneCode: proto.String("123456"),
		NewPhone:     proto.String("13564581235"),
		NewPhoneCode: proto.String("123456"),
	}, time.Second*5, uint32(msgid.MsgID_BIND_PHONE_RSP), &rsp)

	assert.Nil(t, err)
}
