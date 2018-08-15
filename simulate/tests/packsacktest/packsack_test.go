package packsacktest

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

// Test_Get_Packsack_Info 获取背包信息

func Test_Get_Packsack_Info(t *testing.T) {
	player, _ := utils.LoginNewPlayer()
	assert.NotNil(t, player)

	player.AddExpectors(msgid.MsgID_PACKSACK_INFO_RSP)
	client := player.GetClient()
	req := &alms.PlayerPacksackInfoRep{}
	client.SendPackage(utils.CreateMsgHead(msgid.MsgID_PACKSACK_INFO_REQ), req)

	expector := player.GetExpector(msgid.MsgID_PACKSACK_INFO_RSP)
	rsq := &alms.PlayerPacksackInfoRsp{}
	assert.Nil(t, expector.Recv(global.DefaultWaitMessageTime, rsq))
	fmt.Println(rsq.GetPacksackGold())
	for _, r := range rsq.GetPropInfo() {
		fmt.Println("---------------")
		fmt.Println(*r.PropId)
		fmt.Println(*r.PropName)
		fmt.Println(*r.PropType)
		fmt.Println(*r.PropDescribe)
		fmt.Println(*r.PropCount)
	}
}

// Test_Add_Packsack_Gold 背包金币添加
func Test_Add_Packsack_Gold(t *testing.T) {
	player, _ := utils.LoginNewPlayer()
	assert.NotNil(t, player)

	player.AddExpectors(msgid.MsgID_PACKSACK_GOLD_RSP)
	client := player.GetClient()
	req := &alms.PacksackGoldReq{}
	req.ChangeGold = proto.Int64(500)
	client.SendPackage(utils.CreateMsgHead(msgid.MsgID_PACKSACK_GOLD_REQ), req)

	expector := player.GetExpector(msgid.MsgID_PACKSACK_GOLD_RSP)
	rsq := &alms.PacksackGoldRsp{}
	assert.Nil(t, expector.Recv(global.DefaultWaitMessageTime, rsq))
	fmt.Println(rsq)
}

// Test_TakeOut_Packsack_Gold 背包金币取出
func Test_TakeOut_Packsack_Gold(t *testing.T) {
	player, _ := utils.LoginNewPlayer()
	assert.NotNil(t, player)

	player.AddExpectors(msgid.MsgID_PACKSACK_GOLD_RSP)
	client := player.GetClient()
	req := &alms.PacksackGoldReq{}
	req.ChangeGold = proto.Int64(-500)
	client.SendPackage(utils.CreateMsgHead(msgid.MsgID_PACKSACK_GOLD_REQ), req)

	expector := player.GetExpector(msgid.MsgID_PACKSACK_GOLD_RSP)
	rsq := &alms.PacksackGoldRsp{}
	assert.Nil(t, expector.Recv(global.DefaultWaitMessageTime, rsq))
	fmt.Println(rsq)
}
