package utils

import (
	"fmt"
	"steve/client_pb/common"
	"steve/client_pb/hall"
	msgid "steve/client_pb/msgid"
	"steve/client_pb/room"
	"steve/simulate/global"
	"steve/simulate/interfaces"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func createMsgHead(msgID msgid.MsgID) interfaces.SendHead {
	return interfaces.SendHead{
		Head: interfaces.Head{
			MsgID: uint32(msgID),
		},
	}
}

// CreateMsgHead 创建消息头
func CreateMsgHead(msgID msgid.MsgID) interfaces.SendHead {
	return createMsgHead(msgID)
}

// MakeRoomCards 构造牌切片
func MakeRoomCards(card ...room.Card) []*room.Card {
	result := []*room.Card{}
	for i := range card {
		result = append(result, &card[i])
	}
	return result
}

// CheckChuPaiNotifyWithSeats 指定玩家检查出牌广播
func CheckChuPaiNotifyWithSeats(t *testing.T, deskData *DeskData, card uint32, seat int, expectedSeats []int) {
	activePlayer := GetDeskPlayerBySeat(seat, deskData)
	for _, s := range expectedSeats {
		player := GetDeskPlayerBySeat(s, deskData)
		messageExpector := player.Expectors[msgid.MsgID_ROOM_CHUPAI_NTF]
		ntf := &room.RoomChupaiNtf{}
		assert.Nil(t, messageExpector.Recv(global.DefaultWaitMessageTime, ntf))
		assert.Equal(t, card, ntf.GetCard())
		assert.Equal(t, activePlayer.Player.GetID(), ntf.GetPlayer())
	}
}

//CheckChuPaiNotify 检查出牌广播
func CheckChuPaiNotify(t *testing.T, deskData *DeskData, card uint32, seat int) {
	CheckChuPaiNotifyWithSeats(t, deskData, card, seat, []int{0, 1, 2, 3})
}

//CheckMoPaiNotify 检查摸牌广播
func CheckMoPaiNotify(t *testing.T, deskData *DeskData, mopaiSeat int, card uint32) *DeskPlayer {
	player := GetDeskPlayerBySeat(mopaiSeat, deskData)
	for _, deskPlayer := range deskData.Players {
		expector, _ := deskPlayer.Expectors[msgid.MsgID_ROOM_MOPAI_NTF]
		ntf := &room.RoomMopaiNtf{}
		assert.Nil(t, expector.Recv(global.DefaultWaitMessageTime, ntf))
		assert.Equal(t, false, ntf.GetBack())
		if player.Seat == deskPlayer.Seat {
			assert.Equal(t, card, ntf.GetCard())
		} else {
			assert.Equal(t, uint32(0), ntf.GetCard())
		}
	}
	return player
}

//CheckPengNotify 检查碰广播
func CheckPengNotify(t *testing.T, deskData *DeskData, seat int, card uint32) {
	xjPlayer := GetDeskPlayerBySeat(seat, deskData)
	messageExpector := xjPlayer.Expectors[msgid.MsgID_ROOM_PENG_NTF]
	ntf := &room.RoomPengNtf{}
	assert.Nil(t, messageExpector.Recv(2*time.Second, ntf))
	assert.Equal(t, card, ntf.GetCard())
	assert.Equal(t, xjPlayer.Player.GetID(), ntf.GetToPlayerId())
	assert.Equal(t, GetDeskPlayerBySeat(deskData.BankerSeat, deskData).Player.GetID(), ntf.GetFromPlayerId())
}

//CheckZixunNotify 检查自询广播
func CheckZixunNotify(t *testing.T, deskData *DeskData, seat int) {
	xjPlayer := GetDeskPlayerBySeat(seat, deskData)
	zxMessageExpector := xjPlayer.Expectors[msgid.MsgID_ROOM_ZIXUN_NTF]
	zxNtf := &room.RoomZixunNtf{}
	assert.Nil(t, zxMessageExpector.Recv(2*time.Second, zxNtf))
}

// GetPlayerState 获取玩家游戏状态
func GetPlayerState(player interfaces.ClientPlayer) (common.PlayerState, common.GameId, error) {
	client := player.GetClient()

	player.AddExpectors(msgid.MsgID_HALL_GET_PLAYER_STATE_RSP)
	response := hall.HallGetPlayerStateRsp{}
	if _, err := client.SendPackage(createMsgHead(msgid.MsgID_HALL_GET_PLAYER_STATE_REQ), &hall.HallGetPlayerStateReq{}); err != nil {
		return common.PlayerState_PS_IDLE, common.GameId(0), fmt.Errorf("请求获取状态失败:%v", err.Error())
	}
	expector := player.GetExpector(msgid.MsgID_HALL_GET_PLAYER_STATE_RSP)
	if err := expector.Recv(global.DefaultWaitMessageTime, &response); err != nil {
		return common.PlayerState_PS_IDLE, common.GameId(0), fmt.Errorf("收取状态回复失败:%v", err.Error())
	}
	return response.GetPlayerState(), response.GetGameId(), nil
}
