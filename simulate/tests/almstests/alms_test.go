package almstests

import (
	"fmt"
	"steve/client_pb/alms"
	"steve/client_pb/msgid"
	"steve/simulate/global"
	"steve/simulate/utils"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
)

//玩家登陆接受到，救济金配合通知
func Test_Alms_Login(t *testing.T) {
	player, _ := utils.LoginNewPlayer(msgid.MsgID_ALMS_LOGIN_GOLD_CONFIG_NTF)
	assert.NotNil(t, player)
	expector := player.GetExpector(msgid.MsgID_ALMS_LOGIN_GOLD_CONFIG_NTF)
	ntf := &alms.AlmsConfigNtf{}
	assert.Nil(t, expector.Recv(global.DefaultWaitMessageTime, ntf))
	fmt.Println(ntf)
}

// Test_Apply_Alms_Login 测试申请救济金登录

func Test_Apply_Alms_Login(t *testing.T) {
	player, _ := utils.LoginNewPlayer()
	assert.NotNil(t, player)

	player.AddExpectors(msgid.MsgID_ALMS_GET_GOLD_RSP)
	client := player.GetClient()
	reqType := alms.AlmsReqType_LOGIN
	req := &alms.AlmsGetGoldReq{
		ReqType: &reqType,
		Version: proto.Int32(55),
	}
	client.SendPackage(utils.CreateMsgHead(msgid.MsgID_ALMS_GET_GOLD_REQ), req)

	expector := player.GetExpector(msgid.MsgID_ALMS_GET_GOLD_RSP)
	rsq := &alms.AlmsGetGoldRsp{}
	assert.Nil(t, expector.Recv(global.DefaultWaitMessageTime, rsq))
	fmt.Println(rsq)
}

// Test_Apply_Alms_XuanChang 测试申请救济金选场
func Test_Apply_Alms_XuanChang(t *testing.T) {
	player, _ := utils.LoginNewPlayer()
	assert.NotNil(t, player)

	player.AddExpectors(msgid.MsgID_ALMS_GET_GOLD_RSP)
	client := player.GetClient()
	reqType := alms.AlmsReqType_SELECTED
	req := &alms.AlmsGetGoldReq{
		ReqType: &reqType,
		Version: proto.Int32(55),
		GameId:  proto.Int32(4),
		LevelId: proto.Int32(4),
	}
	client.SendPackage(utils.CreateMsgHead(msgid.MsgID_ALMS_GET_GOLD_REQ), req)

	expector := player.GetExpector(msgid.MsgID_ALMS_GET_GOLD_RSP)
	rsq := &alms.AlmsGetGoldRsp{}
	assert.Nil(t, expector.Recv(global.DefaultWaitMessageTime, rsq))
	fmt.Println(rsq)
}
