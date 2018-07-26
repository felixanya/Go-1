package matchv2

import (
	"fmt"
	"steve/client_pb/common"
	"steve/client_pb/match"
	"steve/client_pb/msgid"
	"steve/common/data/player"
	server_pb_match "steve/server_pb/match"
	"steve/server_pb/user"
	"steve/structs"
	"steve/structs/exchanger"
	"steve/structs/proto/gate_rpc"

	"github.com/Sirupsen/logrus"
	"github.com/golang/protobuf/proto"
	nsq "github.com/nsqio/go-nsq"
)

// HandleMatchReq 匹配请求的处理(来自网关服)
func HandleMatchReq(playerID uint64, header *steve_proto_gaterpc.Header, req match.MatchReq) (ret []exchanger.ResponseMsg) {

	response := &match.MatchRsp{
		ErrCode: proto.Int32(0),
		ErrDesc: proto.String("成功"),
	}
	ret = []exchanger.ResponseMsg{{
		MsgID: uint32(msgid.MsgID_MATCH_RSP),
		Body:  response,
	}}

	state := player.GetPlayerPlayState(playerID)
	if state != int(common.PlayerState_PS_IDLE) {
		response.ErrCode = proto.Int32(int32(common.ErrCode_EC_MATCH_ALREADY_GAMEING))
		response.ErrDesc = proto.String("已经在游戏中了")
		return
	}
	if player.GetPlayerCoin(playerID) == 0 {
		response.ErrCode = proto.Int32(1)
		response.ErrDesc = proto.String("金豆数为0，不能参加匹配")
		return
	}

	defaultMgr.addPlayer(playerID, int(req.GetGameId()))
	return
}

// HandleContinueReq 处理续局请求
func HandleContinueReq(playerID uint64, header *steve_proto_gaterpc.Header, req match.MatchDeskContinueReq) (ret []exchanger.ResponseMsg) {

	response := &match.MatchDeskContinueRsp{
		ErrCode: proto.Int32(0),
		ErrDesc: proto.String("成功"),
	}
	ret = []exchanger.ResponseMsg{{
		MsgID: uint32(msgid.MsgID_MATCH_CONTINUE_RSP),
		Body:  response,
	}}

	// 不是取消先判断金币数
	if !req.GetCancel() {
		if player.GetPlayerCoin(playerID) == 0 {
			response.ErrCode = proto.Int32(1)
			response.ErrDesc = proto.String("金豆数为0，不能参加匹配")
			return
		}
	}
	defaultMgr.addContinueApply(playerID, req.GetCancel(), int(req.GetGameId()))
	return
}

// AddContinueDesk 添加续局牌桌
func AddContinueDesk(request *server_pb_match.AddContinueDeskReq) *server_pb_match.AddContinueDeskRsp {
	logrus.WithFields(logrus.Fields{
		"func_name": "AddContinueDesk",
		"request":   request.String(),
	}).Debugln("添加续局牌桌")
	response := &server_pb_match.AddContinueDeskRsp{}

	players := make([]deskPlayer, 0, len(request.GetPlayers()))
	continuePlayers := request.GetPlayers()
	for _, continuePlayer := range continuePlayers {
		players = append(players, deskPlayer{
			playerID: continuePlayer.GetPlayerId(),
			robotLv:  int(continuePlayer.GetRobotLevel()),
			seat:     int(continuePlayer.GetSeat()),
			winner:   continuePlayer.GetWin(),
		})
	}
	defaultMgr.addContinueDesk(players, int(request.GetGameId()), request.GetFixBanker(), int(request.GetBankerSeat()))
	return response
}

type playerLoginHandler struct {
}

func (plh *playerLoginHandler) HandleMessage(message *nsq.Message) error {
	loginPb := user.PlayerLogin{}
	if err := proto.Unmarshal(message.Body, &loginPb); err != nil {
		logrus.WithError(err).Errorln("消息反序列化失败")
		return fmt.Errorf("消息反序列化失败：%v", err)
	}
	defaultMgr.addLoginData(loginPb.PlayerId)
	return nil
}

func init() {
	exposer := structs.GetGlobalExposer()
	if err := exposer.Subscriber.Subscribe("player_login", "match", &playerLoginHandler{}); err != nil {
		logrus.WithError(err).Panicln("订阅登录消息失败")
	}
}
