package user

import (
	"steve/client_pb/common"
	"steve/client_pb/hall"
	"steve/client_pb/msgid"
	"steve/external/goldclient"
	"steve/hall/data"
	"steve/server_pb/gold"
	"steve/structs/exchanger"
	"steve/structs/proto/gate_rpc"

	"github.com/Sirupsen/logrus"
	"github.com/golang/protobuf/proto"
)

// HandleGetPlayerInfoReq2 处理获取玩家个人资料请求
func HandleGetPlayerInfoReq2(playerID uint64, header *steve_proto_gaterpc.Header, req hall.HallGetPlayerInfoReq) (rspMsg []exchanger.ResponseMsg) {
	logrus.Debugln("Handle get player info req", req)

	// 默认返回消息
	response := &hall.HallGetPlayerInfoRsp{
		ErrCode: proto.Uint32(1),
	}
	rspMsg = []exchanger.ResponseMsg{
		exchanger.ResponseMsg{
			MsgID: uint32(msgid.MsgID_HALL_GET_PLAYER_INFO_RSP),
			Body:  response,
		},
	}

	// 获取玩家基本个人资料
	player, err := data.GetPlayerInfo(playerID)
	if err == nil {
		response.ErrCode = proto.Uint32(0)
		response.NickName = proto.String(player.NickName)
		response.Avator = proto.String(player.Avator)
		response.Gender = proto.Uint32(player.Gender)
	}

	// 获取玩家货币信息
	coin, err := goldclient.GetGold(playerID, int16(gold.GoldType_GOLD_COIN))
	if err == nil {
		response.Coin = proto.Uint64(uint64(coin))
	}

	// 获取玩家游戏信息
	state, gameID, _, _, _, _, err := data.GetPlayerState(playerID)
	if err == nil {
		response.PlayerState = common.PlayerState(state).Enum()
		response.GameId = common.GameId(gameID).Enum()
	}

	return
}

// HandleUpdatePlayerInoReq 处理更新玩家个人资料请求
func HandleUpdatePlayerInoReq(playerID uint64, header *steve_proto_gaterpc.Header, req hall.HallUpdatePlayerInfoReq) (rspMsg []exchanger.ResponseMsg) {
	logrus.Debugln("Handle update player info req", req)

	// 默认返回消息
	response := &hall.HallUpdatePlayerInfoRsp{
		ErrCode: proto.Uint32(1),
		Result:  proto.Bool(false),
	}
	rspMsg = []exchanger.ResponseMsg{
		exchanger.ResponseMsg{
			MsgID: uint32(msgid.MsgID_HALL_UPDATE_PLAYER_INFO_REQ),
			Body:  response,
		},
	}

	// 参数校验
	if req.Gender != nil && (req.GetGender() != 1 || req.GetGender() != 2) {
		response.ErrCode = proto.Uint32(1)
	}

	if req.NickName != nil && req.GetNickName() == "" {
		response.ErrCode = proto.Uint32(1)
	}

	if req.Avator != nil && req.GetAvator() == "" {
		response.ErrCode = proto.Uint32(1)
	}

	// 逻辑处理
	exist, result, _ := data.UpdatePlayerInfo(playerID, req.GetNickName(), req.GetAvator(), req.GetGender())

	if !result {
		response.ErrCode = proto.Uint32(1)
	}

	if exist {
		response.ErrCode = proto.Uint32(0)
		response.Result = proto.Bool(true)
	}

	return
}

// HandleGetPlayerStateReq2 获取玩家游戏状态信息
func HandleGetPlayerStateReq2(playerID uint64, header *steve_proto_gaterpc.Header, req hall.HallGetPlayerStateReq) (rspMsg []exchanger.ResponseMsg) {
	logrus.Debugln("Handle get player state req", req)

	// 默认返回消息
	response := &hall.HallGetPlayerStateRsp{
		UserData: proto.Uint64(req.GetUserData()),
	}
	rspMsg = []exchanger.ResponseMsg{
		exchanger.ResponseMsg{
			MsgID: uint32(msgid.MsgID_HALL_GET_PLAYER_STATE_RSP),
			Body:  response,
		},
	}

	// 逻辑处理
	state, gameID, _, _, _, _, err := data.GetPlayerState(playerID)

	// 返回结果
	if err == nil {
		response.PlayerState = common.PlayerState(state).Enum()
		response.GameId = common.GameId(gameID).Enum()
	}

	return
}

// HandleGetGameInfoReq client-> 获取游戏信息列表请求
func HandleGetGameInfoReq(playerID uint64, header *steve_proto_gaterpc.Header, req hall.HallGetGameListInfoReq) (rspMsg []exchanger.ResponseMsg) {
	logrus.Debugln("Handle get game info req", req)

	// 默认返回消息
	response := &hall.HallGetGameListInfoRsp{
		ErrCode: proto.Uint32(1),
	}
	rspMsg = []exchanger.ResponseMsg{
		exchanger.ResponseMsg{
			MsgID: uint32(msgid.MsgID_HALL_GET_PLAYER_INFO_RSP),
			Body:  response,
		},
	}

	// 逻辑处理
	gameInfos, gameLevelInfos, err := data.GetGameInfoList()

	// 返回结果
	if err == nil {
		response.GameConfig = ServerGameConfig2Client(gameInfos)
		response.GameLevelConfig = ServerGameLevelConfig2Client(gameLevelInfos)
	}

	return
}
