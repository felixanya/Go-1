package matchv3

import (
	"context"
	"fmt"
	"steve/client_pb/match"
	"steve/client_pb/msgid"
	"steve/external/gateclient"
	"steve/external/hallclient"
	"steve/server_pb/room_mgr"
	"steve/structs"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/golang/protobuf/proto"
)

// randSeat 给desk的所有玩家分配座位号
func randSeat(desk *matchDesk) {
	var i int32 = 0
	for _, player := range desk.players {
		player.seat = i
		i++
	}
}

// sendCreateDesk 向room服请求创建牌桌，创建失败时则重新请求
func sendCreateDesk(desk matchDesk, globalInfo *levelGlobalInfo) {
	logEntry := logrus.WithFields(logrus.Fields{
		"func_name": "sendCreateDesk",
		"desk":      desk,
	})

	logEntry.Debugln("进入函数，准备向room服请求创建桌子")

	exposer := structs.GetGlobalExposer()

	// 获取room的service
	rs, err := exposer.RPCClient.GetConnectByServerName("room")
	if err != nil || rs == nil {
		logEntry.WithError(err).Errorln("获得room服的gRPC失败，桌子被丢弃!!!")
		return
	}

	// 给desk的所有玩家分配座位号
	randSeat(&desk)

	// 通知玩家，匹配成功，创建桌子
	// matchPlayer转换为deskPlayerInfo
	deskPlayers := []*match.DeskPlayerInfo{}
	for _, player := range desk.players {
		// 给客户端扣除费用后的金币,实际扣除由room服开始游戏时扣除,李全林要求
		player.gold = player.gold - globalInfo.fee

		pDeskPlayer := translateToDeskPlayer(player)
		if pDeskPlayer == nil {
			logEntry.Errorln("把matchPlayer转换为deskPlayerInfo失败，跳过")
			continue
		}
		deskPlayers = append(deskPlayers, pDeskPlayer)
	}

	// 通知消息体
	ntf := match.MatchSucCreateDeskNtf{
		GameId:  &desk.gameID,
		LevelId: &desk.levelID,
		Players: deskPlayers,
	}

	logEntry.Debugf("匹配成功，发给客户端的消息:%v", ntf)

	// 广播给桌子内的所有真实玩家
	for playerID, player := range desk.players {
		if player.robotLv == 0 {
			gateclient.SendPackageByPlayerID(playerID, uint32(msgid.MsgID_MATCH_SUC_CREATE_DESK_NTF), &ntf)
		}
	}

	// 该桌子所有的玩家信息
	createPlayers := []*roommgr.DeskPlayer{}
	for playerID, player := range desk.players {

		deskPlayer := &roommgr.DeskPlayer{
			PlayerId:   playerID,
			RobotLevel: player.robotLv,
			Seat:       uint32(player.seat),
		}

		createPlayers = append(createPlayers, deskPlayer)
	}

	roomMgrClient := roommgr.NewRoomMgrClient(rs)

	req := &roommgr.CreateDeskRequest{
		GameId:   desk.gameID,
		LevelId:  desk.levelID,
		DeskId:   desk.deskID,
		Players:  createPlayers,
		MinCoin:  uint64(globalInfo.minGold),
		MaxCoin:  uint64(globalInfo.maxGold),
		BaseCoin: uint64(globalInfo.bottomScore),
	}

	logEntry.Debugf("匹配成功，发给room服的消息:%v", req)

	// 调用room服的创建桌子
	rsp, err := roomMgrClient.CreateDesk(context.Background(), req)

	// 不成功时，报错，应该重新调用或者重新匹配，暂时丢弃，todo
	if err != nil || rsp.GetErrCode() != roommgr.RoomError_SUCCESS {
		logEntry.WithError(err).Errorln("room服创建桌子失败，桌子被丢弃!!!")

		// 处理错误桌子
		if !dealErrorDesk(&desk) {
			logEntry.WithError(err).Errorln("room服创建桌子失败，处理该桌子时再次失败!!!")
		}

		return
	}

	// 成功时的处理

	// 记录匹配成功的真实玩家同桌信息
	for playerID, player := range desk.players {
		if player.robotLv == 0 {
			globalInfo.sucPlayers[playerID] = desk.deskID
		}
	}

	// 记录匹配成功的桌子信息
	newSucDesk := sucDesk{
		gameID:  desk.gameID,
		levelID: desk.levelID,
		sucTime: time.Now().Unix(),
	}
	globalInfo.sucDesks[desk.deskID] = &newSucDesk

	logEntry.Debugln("离开函数，room服创建桌子成功")

	return
}

// 把 matchPlayer 转换为 match.DeskPlayerInfo
func translateToDeskPlayer(player *matchPlayer) *match.DeskPlayerInfo {

	// 从hall服获取玩家基本信息
	playerInfo, err := hallclient.GetPlayerInfo(player.playerID)
	if err != nil || playerInfo == nil {
		logrus.WithError(err).Errorln("从hall服获取玩家信息失败，玩家ID:%v", player.playerID)
		return nil
	}

	playerName := playerInfo.GetNickName()
	playerGender := playerInfo.GetGender()
	playerAvatar := playerInfo.GetAvatar()
	playerShowuid := playerInfo.GetShowUid()

	// 机器人暂时临时写入,因为hall服没有机器人的信息,以后需更改
	if player.robotLv != 0 {
		playerName = fmt.Sprintf("robot_%v", player.playerID)
	}

	deskPlayer := match.DeskPlayerInfo{
		PlayerId: &player.playerID,
		Name:     proto.String(playerName),
		Coin:     &player.gold,
		Seat:     &player.seat,
		Gender:   proto.Uint32(playerGender),
		Avatar:   proto.String(playerAvatar),
		ShowUid:  proto.Int64(int64(playerShowuid)),
	}

	return &deskPlayer
}
