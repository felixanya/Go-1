package tests

import (
	"steve/client_pb/msgid"
	"steve/client_pb/room"
	"steve/simulate/cheater"
	"steve/simulate/global"
	"steve/simulate/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test_Peng 测试碰
func Test_Peng(t *testing.T) {
	// 清空所有的匹配
	cheater.ClearAllMatch()

	var Int5w uint32 = 15
	params := global.NewCommonStartGameParams()
	// 0 号玩家的最后一张牌改成 5W， 打出后 1 号玩家可碰
	params.Cards[0][13] = 15
	params.Cards[1][0] = 14
	// 修改换三张的牌
	params.HszCards = [][]uint32{
		{13, 13, 13},
		{17, 17, 17},
		{23, 23, 23},
		{27, 27, 27},
	}

	deskData, err := utils.StartGame(params)
	assert.Nil(t, err)

	// 庄家出 5W
	assert.Nil(t, utils.WaitZixunNtf(deskData, params.BankerSeat))
	assert.Nil(t, utils.SendChupaiReq(deskData, params.BankerSeat, Int5w))

	// 1 号玩家收到出牌问询通知，可以碰
	pengSeat := (params.BankerSeat + 1) % len(deskData.Players)
	pengPlayer := utils.GetDeskPlayerBySeat(pengSeat, deskData)

	expector, _ := pengPlayer.Expectors[msgid.MsgID_ROOM_CHUPAIWENXUN_NTF]
	ntf := room.RoomChupaiWenxunNtf{}
	assert.Nil(t, expector.Recv(global.DefaultWaitMessageTime, &ntf))
	assert.Equal(t, Int5w, ntf.GetCard())
	assert.True(t, ntf.GetEnablePeng())
	assert.True(t, ntf.GetEnableQi())

	// 请求碰
	pengClient := pengPlayer.Player.GetClient()
	pengClient.SendPackage(utils.CreateMsgHead(msgid.MsgID_ROOM_XINGPAI_ACTION_REQ), &room.RoomXingpaiActionReq{
		ActionId: room.XingpaiAction_XA_PENG.Enum(),
	})

	from := utils.GetDeskPlayerBySeat(params.BankerSeat, deskData)
	// 所有玩家收到碰通知
	checkPengNotify(t, deskData, pengPlayer.Player.GetID(), from.Player.GetID(), Int5w)
	// 碰的玩家收到自询通知，且只能出牌
	expector, _ = pengPlayer.Expectors[msgid.MsgID_ROOM_ZIXUN_NTF]
	zixunNtf := room.RoomZixunNtf{}
	assert.Nil(t, expector.Recv(global.DefaultWaitMessageTime, &zixunNtf))
	assert.Equal(t, len(zixunNtf.GetEnableAngangCards()), 0)
	assert.Equal(t, len(zixunNtf.GetEnableBugangCards()), 0)
	assert.False(t, zixunNtf.GetEnableZimo())
	assert.NotEqual(t, 0, len(zixunNtf.GetEnableChupaiCards()))
}

func checkPengNotify(t *testing.T, deskData *utils.DeskData, to uint64, from uint64, card uint32) {
	for _, player := range deskData.Players {
		expector, _ := player.Expectors[msgid.MsgID_ROOM_PENG_NTF]
		pengNtf := room.RoomPengNtf{}
		expector.Recv(global.DefaultWaitMessageTime, &pengNtf)
		assert.Equal(t, to, pengNtf.GetToPlayerId())
		assert.Equal(t, from, pengNtf.GetFromPlayerId())
		assert.Equal(t, card, pengNtf.GetCard())
	}
}
