package ai

import (
	"github.com/Sirupsen/logrus"
	"github.com/golang/protobuf/proto"
	"math"
	"steve/client_pb/msgid"
	"steve/client_pb/room"
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

// handlePlayerAI 处理玩家 AI
// result 		: 存放AI事件的结果
// AI			: 具体的AI产生器
// playerID 	: AI针对的玩家的playerID
// ddzContext 	: 牌局信息
// aiType		: 托管，超时等，对应枚举 AIType
// robotLv		: 机器人级别
func (aeg *autoEventGenerator) handleDDZPlayerAI(result *AutoEventGenerateResult, AI CommonAI,
	playerID uint64, deskObj *desk.Desk, aiType AIType, robotLv int) {

	gameContext := deskObj.GetConfig().Context.(*contexts.DDZDeskContext)
	ddzContext := &gameContext.DDZContext

	// 由该AI产生具体的AI事件
	aiResult, err := AI.GenerateAIEvent(AIEventGenerateParams{
		DDZContext: ddzContext,
		PlayerID:   playerID,
		AIType:     aiType,
		RobotLv:    robotLv,
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
	if gameID == int(room.GameId_GAMEID_DOUDIZHU) {
		ddzContext := _gameContext.(*contexts.DDZDeskContext).DDZContext
		state = int32(ddzContext.GetCurState())
	} else {
		mjContext := _gameContext.(*contexts.MajongDeskContext).MjContext
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
		gameContext := params.Desk.GetConfig().Context.(*contexts.DDZDeskContext)
		ddzContext := gameContext.DDZContext

		// 超时事件优先处理
		duration := time.Second * time.Duration(ddzContext.Duration)
		countDownPlayers := ddzContext.CountDownPlayers
		if duration != 0 {
			for _, playerId := range countDownPlayers {
				deskPlayer := playerMgr.GetPlayer(playerId)
				if time.Now().Sub(startTime) >= duration && !deskPlayer.CountingDown && deskPlayer.AddTime >= 0 {
					deskPlayer.CountingDown = true
					msg := room.RoomCountDownNtf{CountDown: proto.Uint32(0), AddCountDown: proto.Uint32(uint32(math.Round(deskPlayer.AddTime.Seconds())))}
					util.SendMessageToPlayer(playerId, msgid.MsgID_ROOM_COUNT_DOWN_NTF, &msg)
					logrus.WithField("playerId", deskPlayer.PlayerID).WithField("msg", msg).Debugln("发送补时消息")
				}
				if deskPlayer.CountingDown {
					deskPlayer.AddTime = deskPlayer.AddTime - TickTime
					if deskPlayer.AddTime < 0 {
						deskPlayer.AddTime = 0
						deskPlayer.SetTuoguan(true, true)
					}
					logrus.WithField("playerId", deskPlayer.PlayerID).WithField("addTime", deskPlayer.AddTime).Debugln("玩家补时")
				}
				if !deskPlayer.CountingDown && time.Now().Sub(startTime) >= duration || deskPlayer.CountingDown && deskPlayer.AddTime <= 0 {
					aeg.handleDDZPlayerAI(&result, AI, playerId, params.Desk, OverTimeAI, 0)
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

			var duration time.Duration
			if isRobot {
				duration = 1 * time.Second
			} else if deskPlayer.IsTuoguan() {
				duration = 2 * time.Second
			}

			if time.Now().Sub(startTime) > duration {
				if isRobot {
					aeg.handleDDZPlayerAI(&result, AI, player.GetPlayerId(), params.Desk, RobotAI, deskPlayer.GetRobotLv())
				} else if deskPlayer.IsTuoguan() {
					aeg.handleDDZPlayerAI(&result, AI, player.GetPlayerId(), params.Desk, TuoGuangAI, 0)
				}
			}
		}
	} else {
		gameContext := params.Desk.GetConfig().Context.(*contexts.MajongDeskContext)
		mjContext := gameContext.MjContext

		players := mjContext.GetPlayers()
		for _, player := range players {
			deskPlayer := playerMgr.GetPlayer(player.GetPlayerId())
			isRobot := deskPlayer.GetRobotLv() != 0
			var duration time.Duration
			if isRobot {
				duration = 1 * time.Second
			} else if gutils.IsHu(player) {
				duration = time.Second * time.Duration(viper.GetInt(fixed.HuStateTimeOut))
			} else if gutils.IsTing(player) {
				duration = time.Second * time.Duration(viper.GetInt(fixed.TingStateTimeOut))
			} else if deskPlayer.IsTuoguan() {
				duration = 2 * time.Second
			} else {
				duration = time.Second * time.Duration(viper.GetInt(fixed.XingPaiTimeOut))
				if time.Now().Sub(startTime) >= duration && !deskPlayer.CountingDown && deskPlayer.AddTime >= 0 {
					deskPlayer.CountingDown = true
					msg := room.RoomCountDownNtf{CountDown: proto.Uint32(0), AddCountDown: proto.Uint32(uint32(math.Round(deskPlayer.AddTime.Seconds())))}
					util.SendMessageToPlayer(player.GetPlayerId(), msgid.MsgID_ROOM_COUNT_DOWN_NTF, &msg)
					logrus.WithField("playerId", deskPlayer.PlayerID).WithField("msg", msg).Debugln("发送补时消息")
				}
				if deskPlayer.CountingDown {
					deskPlayer.AddTime = deskPlayer.AddTime - TickTime
					if deskPlayer.AddTime < 0 {
						deskPlayer.AddTime = 0
						deskPlayer.SetTuoguan(true, true)
					}
					logrus.WithField("playerId", deskPlayer.PlayerID).WithField("addTime", deskPlayer.AddTime).Debugln("玩家补时")
				}
			}
			if !deskPlayer.CountingDown && time.Now().Sub(startTime) >= duration || deskPlayer.CountingDown && deskPlayer.AddTime <= 0 {
				deskPlayer.CountingDown = false
				var aiType AIType
				var eventType int
				if isRobot {
					aiType = RobotAI
					eventType = fixed.RobotEvent
				} else if gutils.IsTing(player) || gutils.IsHu(player) || deskPlayer.IsTuoguan() {
					aiType = TuoGuangAI
					eventType = fixed.TuoGuanEvent
				} else {
					aiType = OverTimeAI
					eventType = fixed.OverTimeEvent
				}

				aiResult, err := AI.GenerateAIEvent(AIEventGenerateParams{
					MajongContext: &mjContext,
					PlayerID:      player.GetPlayerId(),
					AIType:        aiType,
					RobotLv:       deskPlayer.GetRobotLv(),
				})
				if err == nil {
					for _, aiEvent := range aiResult.Events {
						event := desk.DeskEvent{EventID: int(aiEvent.ID), EventType: eventType, Context: aiEvent.Context, PlayerID: player.GetPlayerId(), StateNumber: gameContext.StateNumber, Desk: params.Desk}
						result.Events = append(result.Events, event)
					}
				}
			}
		}
	}
	return result
}

func (aeg *autoEventGenerator) RegisterAI(gameID int, stateID int32, AI CommonAI) {
	//logrus.WithField("gameID", gameID).WithField("stateID", stateID).Debug("Register AI")
	if _, exist := aeg.commonAIs[gameID]; !exist {
		aeg.commonAIs[gameID] = make(map[int32]CommonAI)
	}
	aeg.commonAIs[gameID][stateID] = AI
}
