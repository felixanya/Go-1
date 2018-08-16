package giveuptest

import (
	"steve/client_pb/msgid"
	"steve/client_pb/room"
	"steve/simulate/cheater"
	"steve/simulate/global"
	"steve/simulate/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test_GiveUp 认输测试
func Test_GiveUp(t *testing.T) {
	var Int1B uint32 = 31
	var Int9W uint32 = 19
	params := global.NewCommonStartGameParams()

	params.BankerSeat = 0
	zimoSeat := 1
	bankerSeat := params.BankerSeat

	// 庄家的最后一张牌改为 1B
	params.Cards[bankerSeat][13] = 31
	// 1 号玩家最后1张牌改为 9W
	params.Cards[zimoSeat][12] = 19
	// 墙牌改成 9W 。 墙牌有两张，否则就是海底捞了
	params.WallCards = []uint32{19, 31}

	deskData, err := utils.StartGame(params)
	assert.Nil(t, err)
	assert.NotNil(t, deskData)

	bankerPlayer := utils.GetDeskPlayerBySeat(bankerSeat, deskData)
	cheater.SetPlayerCoin(bankerPlayer.Player.GetID(), 1)

	assert.Nil(t, utils.SendChupaiReq(deskData, bankerSeat, Int1B))

	// 1 号玩家收到可自摸通知
	zimoPlayer := utils.GetDeskPlayerBySeat(zimoSeat, deskData)
	expector, _ := zimoPlayer.Expectors[msgid.MsgID_ROOM_ZIXUN_NTF]
	ntf := room.RoomZixunNtf{}
	assert.Nil(t, expector.Recv(global.DefaultWaitMessageTime, &ntf))
	assert.True(t, ntf.GetEnableZimo())

	// 发送胡请求
	assert.Nil(t, utils.SendHuReq(deskData, zimoSeat))

	// 检测所有玩家收到自摸通知
	utils.CheckHuNotify(t, deskData, []int{zimoSeat}, zimoSeat, Int9W, room.HuType_HT_DIHU)

	// 检测所有玩家收到自摸结算通知
	utils.CheckZiMoSettleNotify(t, deskData, []int{zimoSeat}, zimoSeat, Int9W, room.HuType_HT_DIHU)

	// 2号玩家发送认输请求
	err = utils.SendGiveUpReq(bankerSeat, deskData)

}
