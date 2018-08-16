package utils

import (
	msgid "steve/client_pb/msgid"
	"steve/client_pb/room"
)

// SendGiveUpReq 发送认输请求
func SendGiveUpReq(seat int, deskData *DeskData) error {
	player := GetDeskPlayerBySeat(seat, deskData)
	client := player.Player.GetClient()
	_, err := client.SendPackage(CreateMsgHead(msgid.MsgID_ROOM_PLAYER_GIVEUP_REQ),
		&room.RoomGiveUpReq{})
	return err
}

// SendBrokerPlayerContinue 破产玩家继续游戏请求
func SendBrokerPlayerContinue(seat int, deskData *DeskData) error {
	player := GetDeskPlayerBySeat(seat, deskData)
	client := player.Player.GetClient()
	_, err := client.SendPackage(CreateMsgHead(msgid.MsgID_ROOM_BROKER_PLAYER_CONTINUE_REQ),
		&room.RoomBrokerPlayerContinueRsp{})
	return err
}
