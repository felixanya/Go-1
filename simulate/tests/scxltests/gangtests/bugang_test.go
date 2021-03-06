package gangtests

import (
	"steve/client_pb/room"
	"steve/simulate/global"
	"steve/simulate/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestBuGang 补杠模拟测试
// 游戏流程: 1.庄家出12
//          2.下家碰12,下家打出17
//          3.对家摸到18,对家打出22
//          4.尾家摸到18,尾家打出27
//          5.庄家摸到18,庄家打出16
//          6.下家摸到19,此时下家可以补杠12
//          7.下家选择补杠12,期待:所有人收到下家补杠的广播
func TestBuGang(t *testing.T) {
	param := global.NewCommonStartGameParams()
	param.BankerSeat = 0
	param.Cards[0][4] = 16
	param.Cards[0][5] = 16
	param.Cards[0][6] = 16
	param.Cards[1][4] = 12
	param.Cards[1][5] = 12
	param.Cards[1][6] = 12
	param.WallCards = []uint32{18, 18, 18, 19, 33}
	deskData, err := utils.StartGame(param)
	assert.Nil(t, err)
	utils.WaitZixunNtf(deskData, deskData.BankerSeat)
	//庄家出牌,出12
	utils.SendChupaiReq(deskData, deskData.BankerSeat, uint32(12))
	//检查出牌响应
	utils.CheckChuPaiNotify(t, deskData, uint32(12), deskData.BankerSeat)
	//下家请求碰12
	xjSeat := (deskData.BankerSeat + 1) % len(deskData.Players)
	utils.SendPengReq(deskData, xjSeat)
	//检查碰的通知
	utils.CheckPengNotify(t, deskData, xjSeat, 12)
	//碰成功后收到自询通知
	utils.CheckZixunNotify(t, deskData, xjSeat)
	//下家出牌请求
	utils.SendChupaiReq(deskData, xjSeat, 17)
	//下家出牌响应
	utils.CheckChuPaiNotify(t, deskData, 17, xjSeat)
	//对家摸牌(自寻)响应
	djSeat := (xjSeat + 1) % len(deskData.Players)
	utils.CheckMoPaiNotify(t, deskData, djSeat, 18)
	//对家出牌请求
	utils.SendChupaiReq(deskData, djSeat, 22)
	//对家出牌响应
	utils.CheckChuPaiNotify(t, deskData, 22, djSeat)
	//尾家摸牌(自寻)响应
	wjSeat := (djSeat + 1) % len(deskData.Players)
	utils.CheckMoPaiNotify(t, deskData, wjSeat, 18)
	//尾家出牌请求
	utils.SendChupaiReq(deskData, wjSeat, 27)
	//尾家出牌响应
	utils.CheckChuPaiNotify(t, deskData, 27, wjSeat)
	//庄家摸牌(自寻)响应
	utils.CheckMoPaiNotify(t, deskData, deskData.BankerSeat, 18)
	//庄家出牌请求
	utils.SendChupaiReq(deskData, deskData.BankerSeat, 16)
	//庄家出牌响应
	utils.CheckChuPaiNotify(t, deskData, 16, deskData.BankerSeat)
	//下家摸牌(自寻)响应
	utils.CheckMoPaiNotify(t, deskData, xjSeat, 19)
	//下家补杠
	utils.SendGangReq(deskData, xjSeat, 12, room.GangType_BuGang)
	//下家补杠响应
	player := utils.GetDeskPlayerBySeat(xjSeat, deskData)
	utils.CheckGangNotify(t, deskData, player.Player.GetID(), player.Player.GetID(), 12, room.GangType_BuGang)

}
