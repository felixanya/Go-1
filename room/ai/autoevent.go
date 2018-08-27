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

// AutoEventParams 生成自动事件的参数
type AutoEventParams struct {
	Desk      *desk.Desk
	StartTime time.Time
}

// AutoEventResult 自动事件生成结果
type AutoEventResult struct {
	Events []desk.DeskEvent
}

// DeskAutoEventGenerator 牌桌自动事件产生器
type DeskAutoEventGenerator interface {
	GenerateV2(params *AutoEventParams) AutoEventResult
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

func (aeg *autoEventGenerator) handlePlayerAI(result *AutoEventResult, AI CommonAI, params AIParams) {
	if params.RobotLv == 0 {
		logrus.WithField("params", params).Debugln("处理自动事件")
	}
	aiResult, err := AI.GenerateAIEvent(params)

	var eventType int
	if params.AIType == OverTimeAI {
		eventType = fixed.OverTimeEvent
	} else if params.AIType == RobotAI {
		eventType = fixed.RobotEvent
	} else if params.AIType == TuoGuangAI {
		eventType = fixed.TuoGuanEvent
	}

	// 未出错时，把产生的每一个AI事件压入结果集
	if err == nil {
		for _, aiEvent := range aiResult.Events {
			event := desk.DeskEvent{EventID: int(aiEvent.ID), EventType: eventType, Context: aiEvent.Context, PlayerID: params.PlayerID}
			result.Events = append(result.Events, event)
		}
	}
}

// GenerateV2 利用 AI 生成自动事件
func (aeg *autoEventGenerator) GenerateV2(params *AutoEventParams) (result AutoEventResult) {
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
	if gameID == int(room.GameId_GAMEID_DOUDIZHU) {
		// 超时事件优先处理
		duration := time.Second * time.Duration(ddzContext.Duration)
		countDownPlayers := ddzContext.CountDownPlayers
		if duration != 0 {
			for _, playerId := range countDownPlayers {
				deskPlayer := playerMgr.GetPlayer(playerId)
				if AddTimeCountDown(deskPlayer, startTime, duration) {
					aeg.handlePlayerAI(&result, AI, AIParams{
						MajongContext: &mjContext,
						DDZContext:    &ddzContext,
						PlayerID:      playerId,
						AIType:        OverTimeAI,
						RobotLv:       deskPlayer.GetRobotLv(),
					})
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

			aiType := OverTimeAI
			duration := time.Second * time.Duration(ddzContext.Duration)
			if isRobot {
				duration = 1 * time.Second
				aiType = RobotAI
			} else if deskPlayer.IsTuoguan() {
				duration = 2 * time.Second
				aiType = TuoGuangAI
			}

			if time.Now().Sub(startTime) > duration && aiType != OverTimeAI {
				aeg.handlePlayerAI(&result, AI, AIParams{
					MajongContext: &mjContext,
					DDZContext:    &ddzContext,
					PlayerID:      player.GetPlayerId(),
					AIType:        aiType,
					RobotLv:       deskPlayer.GetRobotLv(),
				})
			}
		}
	} else {
		duration := time.Second * time.Duration(viper.GetInt(fixed.XingPaiTimeOut))
		actionPlayers := GetActionPlayers(&mjContext)
		for _, playerId := range actionPlayers {
			deskPlayer := playerMgr.GetPlayer(playerId)
			if AddTimeCountDown(deskPlayer, startTime, duration) {
				aeg.handlePlayerAI(&result, AI, AIParams{
					MajongContext: &mjContext,
					DDZContext:    &ddzContext,
					PlayerID:      playerId,
					AIType:        OverTimeAI,
					RobotLv:       deskPlayer.GetRobotLv(),
				})
			}
		}

		if len(result.Events) > 0 {
			return
		}

		players := mjContext.GetPlayers()
		for _, player := range players {
			deskPlayer := playerMgr.GetPlayer(player.GetPlayerId())
			isRobot := deskPlayer.GetRobotLv() != 0

			aiType := OverTimeAI
			duration := time.Second * time.Duration(viper.GetInt(fixed.XingPaiTimeOut))
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

			if time.Now().Sub(startTime) >= duration && aiType != OverTimeAI {
				aeg.handlePlayerAI(&result, AI, AIParams{
					MajongContext: &mjContext,
					DDZContext:    &ddzContext,
					PlayerID:      player.GetPlayerId(),
					AIType:        aiType,
					RobotLv:       deskPlayer.GetRobotLv(),
				})
			}
		}
	}
	return result
}

func AddTimeCountDown(deskPlayer *playerpkg.Player, startTime time.Time, duration time.Duration) bool {
	overTime := time.Now().Sub(startTime) >= duration
	addTimeOver := deskPlayer.AddTime <= 0
	if overTime && !deskPlayer.CountingDown && !addTimeOver {
		deskPlayer.CountingDown = true
		msg := room.RoomCountDownNtf{CountDown: proto.Uint32(0), AddCountDown: proto.Uint32(uint32(math.Round(deskPlayer.AddTime.Seconds())))}
		util.SendMessageToPlayer(deskPlayer.GetPlayerID(), msgid.MsgID_ROOM_COUNT_DOWN_NTF, &msg)
		logrus.WithField("playerId", deskPlayer.PlayerID).WithField("msg", msg).Debugln("发送补时消息")
	}
	if deskPlayer.CountingDown {
		deskPlayer.AddTime = deskPlayer.AddTime - TickTime
		if deskPlayer.AddTime <= 0 {
			deskPlayer.AddTime = 0
			deskPlayer.CountingDown = false
			addTimeOver = true
			logrus.WithField("playerId", deskPlayer.PlayerID).Debugln("玩家补时用尽")
		}
	}

	if addTimeOver && overTime {
		if !deskPlayer.IsTuoguan() {
			deskPlayer.SetTuoguan(true, true)
		}
	}

	return !deskPlayer.CountingDown && overTime || deskPlayer.CountingDown && addTimeOver
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
