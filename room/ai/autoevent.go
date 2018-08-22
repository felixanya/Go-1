package ai

import (
	"github.com/Sirupsen/logrus"
	"github.com/golang/protobuf/proto"
	"math"
	"steve/client_pb/msgid"
	"steve/client_pb/room"
	"steve/entity/majong"
	"steve/entity/poker/ddz"
	"steve/gutils"
	"steve/room/contexts"
	"steve/room/desk"
	"steve/room/fixed"
	playerpkg "steve/room/player"
	"steve/room/util"
	"time"

	"github.com/spf13/viper"
)

var TickTime = time.Millisecond * 200

// AutoEventGenerateParams 生成自动事件的参数
type AutoEventGenerateParams struct {
	Desk      *desk.Desk
	StartTime time.Time
}

// AutoEventGenerateResult 自动事件生成结果
type AutoEventGenerateResult struct {
	Events []desk.DeskEvent
}

// DeskAutoEventGenerator 牌桌自动事件产生器
type DeskAutoEventGenerator interface {
	GenerateV2(params *AutoEventGenerateParams) AutoEventGenerateResult
	RegisterAI(gameID int, stateID int32, AI CommonAI) // 注册 AI
}

type autoEventGenerator struct {
	commonAIs map[int](map[int32]CommonAI) // Key：游戏ID，对应枚举GameId，Value:该游戏各个状态对应的AI产生器
}

var generator = &autoEventGenerator{
	commonAIs: make(map[int](map[int32]CommonAI)),
}

// GetAtEvent 获取生成器
func GetAtEvent() DeskAutoEventGenerator {
	return generator
}

func (aeg *autoEventGenerator) handlePlayerAI(result *AutoEventGenerateResult, AI CommonAI,
	playerID uint64, deskObj *desk.Desk, aiType AIType, robotLv int, ddzContext *ddz.DDZContext, mjContext *majong.MajongContext) {

	// 由该AI产生具体的AI事件
	aiResult, err := AI.GenerateAIEvent(AIEventGenerateParams{
		MajongContext: mjContext,
		DDZContext:    ddzContext,
		PlayerID:      playerID,
		AIType:        aiType,
		RobotLv:       robotLv,
	})

	var eventType int
	if aiType == OverTimeAI {
		eventType = fixed.OverTimeEvent
	} else if aiType == RobotAI {
		eventType = fixed.RobotEvent
	} else if aiType == TuoGuangAI {
		eventType = fixed.TuoGuanEvent
	}

	// 未出错时，把产生的每一个AI事件压入结果集
	if err == nil {
		for _, aiEvent := range aiResult.Events {
			event := desk.DeskEvent{EventID: int(aiEvent.ID), EventType: eventType, Context: aiEvent.Context, PlayerID: playerID, Desk: deskObj}
			result.Events = append(result.Events, event)
		}
	}
}

// GenerateV2 利用 AI 生成自动事件
func (aeg *autoEventGenerator) GenerateV2(params *AutoEventGenerateParams) (result AutoEventGenerateResult) {
	mydesk := params.Desk

	// 该桌子所属的游戏ID
	gameID := mydesk.GetGameId()

	// 该游戏各个状态对应的AI产生器
	gameAIs, ok := aeg.commonAIs[int(gameID)]
	if !ok {
		//logrus.WithField("gameId", gameID).Debug("Can't find game AI")
		return
	}
	_gameContext := mydesk.GetConfig().Context

	// 当前状态ID
	var state int32
	var ddzContext ddz.DDZContext
	var mjContext majong.MajongContext
	if gameID == int(room.GameId_GAMEID_DOUDIZHU) {
		ddzContext = _gameContext.(*contexts.DDZDeskContext).DDZContext
		state = int32(ddzContext.GetCurState())
	} else {
		mjContext = _gameContext.(*contexts.MajongDeskContext).MjContext
		state = int32(mjContext.GetCurState())
	}

	// 当前状态的AI产生器
	AI, ok := gameAIs[int32(state)]
	if !ok {
		//logrus.WithField("gameId", gameID).WithField("state", state).Debug("Can't find state AI")
		return
	}

	startTime := params.StartTime
	playerMgr := playerpkg.GetPlayerMgr()
	var aiType AIType
	var duration time.Duration
	if gameID == int(room.GameId_GAMEID_DOUDIZHU) {
		// 超时事件优先处理
		duration = time.Second * time.Duration(ddzContext.Duration)
		countDownPlayers := ddzContext.CountDownPlayers
		if duration != 0 {
			for _, playerId := range countDownPlayers {
				deskPlayer := playerMgr.GetPlayer(playerId)
				if AddTimeCountDown(deskPlayer, startTime, duration) {
					deskPlayer.CountingDown = false
					aeg.handlePlayerAI(&result, AI, playerId, params.Desk, OverTimeAI, 0, &ddzContext, nil)
				}
			}
		}

		if len(result.Events) > 0 {
			return
		}

		// 处理其他事件
		players := ddzContext.GetPlayers()
		for _, player := range players {
			deskPlayer := playerMgr.GetPlayer(player.GetPlayerId())
			isRobot := deskPlayer.GetRobotLv() != 0

			if isRobot {
				duration = 1 * time.Second
				aiType = RobotAI
			} else if deskPlayer.IsTuoguan() {
				duration = 2 * time.Second
				aiType = TuoGuangAI
			}

			if time.Now().Sub(startTime) > duration && aiType != 0 {
				aeg.handlePlayerAI(&result, AI, player.GetPlayerId(), params.Desk, aiType, deskPlayer.GetRobotLv(), &ddzContext, nil)
			}
		}
	} else {
		players := mjContext.GetPlayers()
		duration = time.Second * time.Duration(viper.GetInt(fixed.XingPaiTimeOut))

		actionPlayers := GetActionPlayers(&mjContext)
		for _, player := range players {
			deskPlayer := playerMgr.GetPlayer(player.GetPlayerId())

			if ContainsUint64(actionPlayers, player.GetPlayerId()) && AddTimeCountDown(deskPlayer, startTime, duration) {
				deskPlayer.CountingDown = false
				aeg.handlePlayerAI(&result, AI, player.GetPlayerId(), params.Desk, OverTimeAI, 0, nil, &mjContext)
			}
			if len(result.Events) > 0 {
				continue
			}

			isRobot := deskPlayer.GetRobotLv() != 0
			if isRobot {
				duration = 1 * time.Second
				aiType = RobotAI
			} else if gutils.IsHu(player) {
				duration = time.Second * time.Duration(viper.GetInt(fixed.HuStateTimeOut))
				aiType = TuoGuangAI
			} else if gutils.IsTing(player) {
				duration = time.Second * time.Duration(viper.GetInt(fixed.TingStateTimeOut))
				aiType = TuoGuangAI
			} else if deskPlayer.IsTuoguan() {
				duration = 2 * time.Second
				aiType = TuoGuangAI
			}
			if time.Now().Sub(startTime) >= duration && aiType != 0 {
				aeg.handlePlayerAI(&result, AI, player.GetPlayerId(), params.Desk, aiType, 0, nil, &mjContext)
			}
		}
	}
	return result
}

func AddTimeCountDown(deskPlayer *playerpkg.Player, startTime time.Time, duration time.Duration) bool {
	overTime := time.Now().Sub(startTime) >= duration
	if overTime && !deskPlayer.CountingDown && deskPlayer.AddTime > 0 {
		deskPlayer.CountingDown = true
		msg := room.RoomCountDownNtf{CountDown: proto.Uint32(0), AddCountDown: proto.Uint32(uint32(math.Round(deskPlayer.AddTime.Seconds())))}
		util.SendMessageToPlayer(deskPlayer.GetPlayerID(), msgid.MsgID_ROOM_COUNT_DOWN_NTF, &msg)
		logrus.WithField("playerId", deskPlayer.PlayerID).WithField("msg", msg).Debugln("发送补时消息")
	}
	if deskPlayer.CountingDown {
		deskPlayer.AddTime = deskPlayer.AddTime - TickTime
		if deskPlayer.AddTime <= 0 {
			deskPlayer.AddTime = 0
			deskPlayer.SetTuoguan(true, true)
		}
		logrus.WithField("playerId", deskPlayer.PlayerID).WithField("addTime", deskPlayer.AddTime).Debugln("玩家补时")
	}
	if deskPlayer.AddTime <= 0 && overTime {
		deskPlayer.SetTuoguan(true, true)
	}
	return !deskPlayer.CountingDown && overTime || deskPlayer.CountingDown && deskPlayer.AddTime <= 0
}

func (aeg *autoEventGenerator) RegisterAI(gameID int, stateID int32, AI CommonAI) {
	//logrus.WithField("gameID", gameID).WithField("stateID", stateID).Debug("Register AI")
	if _, exist := aeg.commonAIs[gameID]; !exist {
		aeg.commonAIs[gameID] = make(map[int32]CommonAI)
	}
	aeg.commonAIs[gameID][stateID] = AI
}

func GetActionPlayers(ctx *majong.MajongContext) []uint64 {
	players := ctx.GetPlayers()
	actionPlayers := make([]uint64, 0, 4)

	switch ctx.GetCurState() {
	case majong.StateID_state_dingque:
		{
			for _, player := range players {
				if !player.GetHasDingque() {
					actionPlayers = append(actionPlayers, player.GetPlayerId())
				}
			}
		}
	case majong.StateID_state_huansanzhang:
		{
			for _, player := range players {
				if !player.GetHuansanzhangSure() {
					actionPlayers = append(actionPlayers, player.GetPlayerId())
				}
			}
		}
	case majong.StateID_state_zixun:
		{
			actionPlayers = append(actionPlayers, gutils.GetZixunPlayer(ctx))
		}
	case majong.StateID_state_chupaiwenxun, majong.StateID_state_waitqiangganghu:
		{

			for _, player := range players {
				if !player.GetHasSelected() && len(player.GetPossibleActions()) > 0 {
					actionPlayers = append(actionPlayers, player.GetPlayerId())
				}
			}
		}
	}
	return actionPlayers
}

func ContainsUint64(players []uint64, inPlayer uint64) bool {
	for _, player := range players {
		if player == inPlayer {
			return true
		}
	}
	return false
}
