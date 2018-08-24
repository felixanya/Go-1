package models

import (
	"steve/client_pb/room"
	"steve/external/hallclient"
	"steve/room/desk"
	"steve/room/player"

	"github.com/Sirupsen/logrus"
	"github.com/golang/protobuf/proto"
)

type BaseModel struct {
	desk *desk.Desk
}

func (model *BaseModel) GetDesk() *desk.Desk {
	return model.desk
}

func (model *BaseModel) SetDesk(desk *desk.Desk) {
	model.desk = desk
}

func (model *BaseModel) GetGameContext() interface{} {
	return model.desk.GetConfig().Context
}

// TranslateToRoomPlayer 将 deskPlayer 转换成 RoomPlayerInfo
func TranslateToRoomPlayer(player *player.Player) room.RoomPlayerInfo {

	coin := player.GetCoin()
	playerID := player.GetPlayerID()

	var name string = "player" // 名字
	var gender uint32 = 0      // 性别
	var avatar string = ""     // 头像
	var showUid int64 = 0      // 显示ID

	// 从hall服获取玩家信息
	playerInfoRsp, err := hallclient.GetPlayerInfo(playerID)
	if err != nil || playerInfoRsp == nil {
		logrus.WithError(err).Errorln("TranslateToRoomPlayer() 从hall服获取玩家游戏信息失败")
	}
	name = playerInfoRsp.GetNickName()
	gender = playerInfoRsp.GetGender()
	avatar = playerInfoRsp.GetAvatar()
	showUid = int64(playerInfoRsp.GetShowUid())

	return room.RoomPlayerInfo{
		PlayerId: proto.Uint64(player.GetPlayerID()),
		Name:     proto.String(name),
		Coin:     proto.Uint64(coin),
		Seat:     proto.Uint32(uint32(player.GetSeat())),
		// Location: TODO 没地方拿
		ShowUid: proto.Int64(showUid),
		Quited:  proto.Bool(player.IsQuit()),
		Gender:  proto.Uint32(gender),
		Avatar:  proto.String(avatar),
	}
}
