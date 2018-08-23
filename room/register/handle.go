package registers

import (
	"steve/client_pb/common"
	"steve/client_pb/match"
	"steve/client_pb/msgid"
	"steve/client_pb/room"
	"steve/common/constant"
	propclient "steve/common/data/prop"
	"steve/entity/majong"
	"steve/external/goldclient"
	"steve/external/propsclient"
	"steve/gold/define"
	"steve/gutils"
	"steve/room/contexts"
	"steve/room/majong/utils"
	"steve/room/models"
	modelmanager "steve/room/models"
	player2 "steve/room/player"
	"steve/room/util"
	"steve/server_pb/gold"
	"steve/structs/exchanger"
	"steve/structs/proto/gate_rpc"

	"github.com/Sirupsen/logrus"
	"github.com/golang/protobuf/proto"
)

// HandleRoomChatReq 处理玩家聊天请求
func HandleRoomChatReq(playerID uint64, header *steve_proto_gaterpc.Header, req room.RoomDeskChatReq) (ret []exchanger.ResponseMsg) {
	player := player2.GetPlayerMgr().GetPlayer(playerID)
	if player == nil {
		return
	}
	modelmanager.GetModelManager().GetChatModel(player.GetDesk().GetUid()).RoomChatMsgReq(player, header, req)
	return
}

// HandleRoomDeskQuitReq 处理玩家退出桌面请求
// 失败先不回复
func HandleRoomDeskQuitReq(playerID uint64, header *steve_proto_gaterpc.Header, req room.RoomDeskQuitReq) (rspMsg []exchanger.ResponseMsg) {
	response := room.RoomDeskQuitRsp{
		UserData: proto.Uint32(req.GetUserData()),
		ErrCode:  room.RoomError_FAILED.Enum(),
	}
	player := player2.GetPlayerMgr().GetPlayer(playerID)
	if player == nil {
		util.SendMessageToPlayer(playerID, msgid.MsgID_ROOM_DESK_QUIT_RSP, &response)
		return
	}
	desk := player.GetDesk()
	if desk == nil {
		util.SendMessageToPlayer(playerID, msgid.MsgID_ROOM_DESK_QUIT_RSP, &response)
		return
	}
	modelmanager.GetModelManager().GetPlayerModel(desk.GetUid()).PlayerQuit(player)
	response.ErrCode = room.RoomError_SUCCESS.Enum()
	util.SendMessageToPlayer(playerID, msgid.MsgID_ROOM_DESK_QUIT_RSP, &response)
	return
}

func noGamePlaying() []exchanger.ResponseMsg {
	body := &room.RoomResumeGameRsp{
		ResumeRes: room.RoomError_DESK_NO_GAME_PLAYING.Enum(),
	}
	return []exchanger.ResponseMsg{
		exchanger.ResponseMsg{
			MsgID: uint32(msgid.MsgID_ROOM_RESUME_GAME_RSP),
			Body:  body,
		},
	}
}

// HandleResumeGameReq 恢复对局请求
func HandleResumeGameReq(playerID uint64, header *steve_proto_gaterpc.Header, req room.RoomResumeGameReq) (ret []exchanger.ResponseMsg) {
	entry := logrus.WithField("player_id", playerID)
	player := player2.GetPlayerMgr().GetPlayer(playerID)
	if player == nil {
		entry.Debugln("玩家不存在")
		return noGamePlaying()
	}
	desk := player.GetDesk()
	if desk == nil {
		entry.Debugln("没有对应的牌桌")
		return noGamePlaying()
	}
	entry.Debugln("处理恢复牌局请求开始")
	modelmanager.GetModelManager().GetPlayerModel(desk.GetUid()).PlayerEnter(player)
	entry.Debugln("处理恢复牌局请求完成")
	return
}

// HandleCancelTuoGuanReq 处理取消托管请求
func HandleCancelTuoGuanReq(playerID uint64, header *steve_proto_gaterpc.Header, req room.RoomCancelTuoGuanReq) (ret []exchanger.ResponseMsg) {
	ret = []exchanger.ResponseMsg{}

	logEntry := logrus.WithFields(logrus.Fields{
		"func_name": "HandleCancelTuoGuanReq",
		"player_id": playerID,
	})
	player := player2.GetPlayerMgr().GetPlayer(playerID)
	if player == nil {
		logEntry.Debugln("获取玩家失败")
		return
	}
	desk := player.GetDesk()
	if desk == nil {
		logEntry.Debugln("玩家不在房间中")
		return
	}
	player.SetTuoguan(false, true)
	return
}

// HandleTuoGuanReq 处理托管请求
func HandleTuoGuanReq(playerID uint64, header *steve_proto_gaterpc.Header, req room.RoomTuoGuanReq) (ret []exchanger.ResponseMsg) {
	ret = []exchanger.ResponseMsg{}

	logEntry := logrus.WithFields(logrus.Fields{
		"func_name": "HandleTuoGuanReq",
		"player_id": playerID,
	})
	player := player2.GetPlayerMgr().GetPlayer(playerID)
	if player == nil {
		logEntry.Debugln("获取玩家失败")
		return
	}
	desk := player.GetDesk()
	if desk == nil {
		logEntry.Debugln("玩家不在房间中")
		return
	}
	player.SetTuoguan(req.GetTuoguan(), true)
	return
}

// HandleAutoHuReq 处理自动胡牌请求
func HandleAutoHuReq(playerID uint64, header *steve_proto_gaterpc.Header, req room.RoomAutoHuReq) (ret []exchanger.ResponseMsg) {
	ret = []exchanger.ResponseMsg{}

	logEntry := logrus.WithFields(logrus.Fields{
		"func_name": "HandleAutoHuReq",
		"player_id": playerID,
	})
	player := player2.GetPlayerMgr().GetPlayer(playerID)
	if player == nil {
		logEntry.Debugln("获取玩家失败")
		return
	}
	desk := player.GetDesk()
	if desk == nil {
		logEntry.Debugln("玩家不在房间中")
		return
	}
	player.SetAutoHu(req.GetEnable())

	body := &room.RoomAutoHuRsp{
		ErrCode: room.RoomError_SUCCESS.Enum(),
		Enable:  req.Enable,
	}
	ret = []exchanger.ResponseMsg{{
		MsgID: uint32(msgid.MsgID_ROOM_AUTOHU_RSP),
		Body:  body,
	}}
	return
}

// HandleContinueReq 处理续局请求
func HandleContinueReq(playerID uint64, header *steve_proto_gaterpc.Header, req match.MatchDeskContinueReq) (ret []exchanger.ResponseMsg) {
	entry := logrus.WithField("player_id", playerID)

	response := &match.MatchDeskContinueRsp{
		ErrCode: proto.Int32(0),
		ErrDesc: proto.String("成功"),
	}
	ret = []exchanger.ResponseMsg{{
		MsgID: uint32(msgid.MsgID_MATCH_CONTINUE_RSP),
		Body:  response,
	}}

	player := player2.GetPlayerMgr().GetPlayer(playerID)
	if player == nil {
		response.ErrCode = proto.Int32(int32(common.ErrCode_EC_FAIL))
		response.ErrDesc = proto.String("续局时获取玩家失败")

		entry.Debugln("获取玩家失败")
		return
	}
	desk := player.GetDesk()
	if desk == nil {
		response.ErrCode = proto.Int32(int32(common.ErrCode_EC_FAIL))
		response.ErrDesc = proto.String("续局时玩家不在房间")

		entry.Debugln("续局时玩家不在房间")
		return
	}
	continueModel := models.GetContinueModel(desk.GetUid())
	*response = continueModel.PushContinueRequest(playerID, &req)
	return
}

// HandleUsePropReq 使用道具请求处理
func HandleUsePropReq(playerID uint64, header *steve_proto_gaterpc.Header, req room.RoomUsePropReq) (ret []exchanger.ResponseMsg) {
	logEntry := logrus.WithFields(logrus.Fields{
		"func_name":    "HandleUsePropReq",
		"player_id":    playerID,
		"to_player_id": *req.PlayerId,
		"prop_id":      *req.PropId,
	})
	logEntry.Println("开始处理使用道具请求")
	errDesc := "无法使用该道具"
	rsp := room.RoomUsePropRsp{
		ErrCode: room.RoomError_FAILED.Enum(),
		ErrDesc: &errDesc,
	}
	ret = []exchanger.ResponseMsg{
		exchanger.ResponseMsg{
			MsgID: uint32(msgid.MsgID_ROOM_USE_PROP_RSP),
			Body:  &rsp,
		},
	}

	player := player2.GetPlayerMgr().GetPlayer(playerID)
	if player == nil {
		logEntry.Debugln("获取玩家失败")
		return
	}
	desk := player.GetDesk()
	if desk == nil {
		logEntry.Debugln("玩家不在房间中")
		return
	}

	propID := gutils.PropTypeClient2Server(req.GetPropId())
	prop, err := propclient.GetPlayerOneProp(playerID, propID)

	// 使用道具
	if prop.Count > 0 {
		propsList := map[uint64]int64{
			uint64(propID): -1,
		}
		err = propsclient.AddUserProps(playerID, propsList, int32(constant.PFGAMEUSE), 0, int32(req.GetGameId()), req.GetLevelId())
		if err != nil {
			logEntry.WithError(err).Debugln("增加玩家道具失败")
			return
		}
	} else { // 使用金币购买
		propConfig, err := propclient.GetOnePropsConfig(propID)
		if err != nil {
			logEntry.WithError(err).Debugln("获取道具配置失败")
			return
		}

		coin, err := goldclient.GetGold(playerID, int16(gold.GoldType_GOLD_COIN))
		if coin >= propConfig.Limit {
			goldclient.AddGold(playerID, int16(gold.GoldType_GOLD_COIN), propConfig.Value, 0, 0, int32(desk.GetGameId()), desk.GetLevel())
		} else {
			logEntry.WithError(err).Debugf("玩家金币数不足, limit: %d, coin: %d", propConfig.Limit, coin)
			errDesc = "金币不足，无法使用该道具"
			ret[0].Body.(*room.RoomUsePropRsp).ErrDesc = &errDesc
			return
		}
	}

	// 广播道具
	desPlayerID := req.GetPlayerId()
	ntf := room.RoomUsePropNtf{
		FromPlayerId: &playerID,
		ToPlayerId:   &desPlayerID,
		PropId:       req.PropId,
	}
	logEntry.Println("使用道具玩家 和 接受玩家", playerID, desPlayerID)
	msgBody, err := proto.Marshal(&ntf)
	if err != nil {
		logEntry.WithError(err).Debugln("序列化失败")
		return
	}
	modelmanager.GetModelManager().GetMessageModel(desk.GetUid()).BroadcastMessage([]uint64{}, msgid.MsgID_ROOM_USE_PROP_NTF, msgBody, true)
	return nil
}

// HandlePlayerGameGiveUp 处理玩家游戏内认输
func HandlePlayerGameGiveUp(playerID uint64, header *steve_proto_gaterpc.Header, req room.RoomGiveUpReq) (ret []exchanger.ResponseMsg) {
	logEntry := logrus.WithFields(logrus.Fields{
		"func_name": "HandlePlayerGameGiveUp",
		"player_id": playerID,
	})
	logEntry.Debugf("开始玩家游戏内认输")

	rsp := room.RoomGiveUpRsp{
		ErrCode: room.RoomError_FAILED.Enum(),
	}
	ret = []exchanger.ResponseMsg{
		exchanger.ResponseMsg{
			MsgID: uint32(msgid.MsgID_ROOM_PLAYER_GIVEUP_RSP),
			Body:  &rsp,
		},
	}
	player := player2.GetPlayerMgr().GetPlayer(playerID)
	if player == nil {
		logEntry.Debugln("获取玩家失败")
		return
	}
	desk := player.GetDesk()
	if desk == nil {
		logEntry.Debugln("玩家不在房间中")
		return
	}

	// 斗地主没有认输功能
	if desk.GetGameId() == int(common.GameId_GAMEID_DOUDIZHU) {
		return
	}

	// 麻将context,修改玩家状态为认输
	majongContext := desk.GetConfig().Context.(*contexts.MajongDeskContext)
	majongPlayer := utils.GetMajongPlayer(playerID, &majongContext.MjContext)
	majongPlayer.XpState = majongPlayer.GetXpState() | majong.XingPaiState_give_up

	// 麻将settle, 修改该破产玩家为认输
	majongSettle := desk.GetConfig().Settle.(*models.MajongSettle)
	majongSettle.HandleBrokerPlayer(desk, playerID)

	// 广播该玩家认输
	ntf := room.RoomGiveUpNtf{
		PlayerId: []uint64{playerID},
	}
	logEntry.Println("认输玩家", playerID)
	msgBody, err := proto.Marshal(&ntf)
	if err != nil {
		logEntry.WithError(err).Debugln("序列化失败")
		return
	}
	modelmanager.GetModelManager().GetMessageModel(desk.GetUid()).BroadcastMessage([]uint64{}, msgid.MsgID_ROOM_PLAYER_GIVEUP_NTF, msgBody, true)

	rsp.ErrCode = room.RoomError_SUCCESS.Enum()
	return
}

// HandleRoomBrokerPlayerContinue 处理破产玩家继续游戏
func HandleRoomBrokerPlayerContinue(playerID uint64, header *steve_proto_gaterpc.Header, req room.RoomBrokerPlayerContinueReq) (ret []exchanger.ResponseMsg) {
	logEntry := logrus.WithFields(logrus.Fields{
		"func_name": "HandlePlayerGameRenew",
		"player_id": playerID,
	})
	logEntry.Debugf("开始玩家游戏内续费")

	rsp := room.RoomBrokerPlayerContinueRsp{
		ErrCode: room.RoomError_FAILED.Enum(),
	}
	ret = []exchanger.ResponseMsg{
		exchanger.ResponseMsg{
			MsgID: uint32(msgid.MsgID_ROOM_PLAYER_GIVEUP_RSP),
			Body:  &rsp,
		},
	}
	player := player2.GetPlayerMgr().GetPlayer(playerID)
	if player == nil {
		logEntry.Debugln("获取玩家失败")
		return
	}
	desk := player.GetDesk()
	if desk == nil {
		logEntry.Debugln("玩家不在房间中")
		return
	}

	// 获取游戏金币
	gold, err := goldclient.GetGold(playerID, define.GOLD_COIN)

	if err != nil {
		logEntry.Debugln("获取玩家playerId:(%d)金币错误，error:(%v)", playerID, err.Error())
		return
	}

	if gold == 0 {
		logEntry.Debugln("获取玩家playerId:(%d)金币为0", playerID)
		return
	}

	// 麻将settle, 修改该破产玩家为认输
	majongSettle := desk.GetConfig().Settle.(*models.MajongSettle)
	majongSettle.HandleBrokerPlayer(desk, playerID)

	rsp.ErrCode = room.RoomError_SUCCESS.Enum()
	return
}
