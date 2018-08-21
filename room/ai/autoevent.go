package ai

import (
	"steve/client_pb/room"
	"steve/gutils"
	"steve/room/contexts"
	"steve/room/desk"
	"steve/room/fixed"
	playerpkg "steve/room/player"
	"time"

	"github.com/spf13/viper"
)

// AutoEventGenerateParams 生成自动事件的参数
type AutoEventGenerateParams struct {
	Desk      *desk.Desk
	StartTime time.Time
	RobotLv   map[uint64]int
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

	// 未出错时，把产生的每一个AI事件压入结果集
	if err == nil {
		for _, aiEvent := range aiResult.Events {
			event := desk.DeskEvent{EventID: int(aiEvent.ID), EventType: fixed.OverTimeEvent, Context: aiEvent.Context, PlayerID: playerID, Desk: deskObj}
			result.Events = append(result.Events, event)
		}
	}
}

// handleDDZOverTime 斗地主超时处理
// finish : 是否处理完成
// result : 产生的AI事件结果集合
func (aeg *autoEventGenerator) handleDDZOverTime(AI CommonAI, params *AutoEventGenerateParams) (
	finish bool, result AutoEventGenerateResult) {

	finish, result = false, AutoEventGenerateResult{
		Events: []desk.DeskEvent{},
	}

	// 开始时间
	startTime := params.StartTime

	gameContext := params.Desk.GetConfig().Context.(*contexts.DDZDeskContext)
	ddzContext := gameContext.DDZContext

	// 倒计时的时长
	duration := time.Second * time.Duration(ddzContext.Duration)

	// 未到倒计时，不处理
	if duration == 0 || time.Now().Sub(startTime) < duration {
		return
	}

	// 处理每一个处于倒计时的玩家，产生具体的AI事件，并把事件存入result
	players := ddzContext.CountDownPlayers
	for _, player := range players {
		aeg.handleDDZPlayerAI(&result, AI, player, params.Desk, OverTimeAI, 0)
	}

	finish = true
	return
}

// handleDDZTuoGuan 斗地主托管处理
// finish : 是否处理完成
// result : 产生的AI事件结果集合
func (aeg *autoEventGenerator) handleDDZTuoGuan(deskObj *desk.Desk, AI CommonAI, stateTime time.Time) AutoEventGenerateResult {
	result := AutoEventGenerateResult{
		Events: []desk.DeskEvent{},
	}

	// 托管时的操作等待时间
	tuoguanOprTime := 2 * time.Second

	if time.Now().Sub(stateTime) < tuoguanOprTime {
		return result
	}
	playerIDs := deskObj.GetPlayerIds()
	playerMgr := playerpkg.GetPlayerMgr()
	for _, playerID := range playerIDs {
		player := playerMgr.GetPlayer(playerID)
		if player.IsTuoguan() {
			aeg.handleDDZPlayerAI(&result, AI, playerID, deskObj, TuoGuangAI, 0)
		}
	}
	return result
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

	// 斗地主的特殊处理
	if gameID == int(room.GameId_GAMEID_DOUDIZHU) {

		// 先处理超时
		if overTime, result := aeg.handleDDZOverTime(AI, params); overTime {
			return result
		}

		// 未超时，则处理托管
		result = aeg.handleDDZTuoGuan(params.Desk, AI, params.StartTime)
		ddzContext := _gameContext.(*contexts.DDZDeskContext).DDZContext

		// 超过 1s 处理机器人事件
		if time.Now().Sub(params.StartTime) > 1*time.Second {
			players := ddzContext.GetPlayers()
			for _, player := range players {
				playerID := player.GetPlayerId()
				if lv, exist := params.RobotLv[playerID]; exist && lv != 0 {
					aeg.handleDDZPlayerAI(&result, AI, player.GetPlayerId(), params.Desk, RobotAI, lv)
				}
			}
		}
	} else {
		startTime := params.StartTime
		gameContext := params.Desk.GetConfig().Context.(*contexts.MajongDeskContext)
		mjContext := gameContext.MjContext

		result := AutoEventGenerateResult{
			Events: []desk.DeskEvent{},
		}

		players := mjContext.GetPlayers()
		playerMgr := playerpkg.GetPlayerMgr()
		for _, player := range players {
			deskPlayer := playerMgr.GetPlayer(player.GetPlayerId())
			lv, exist := params.RobotLv[player.GetPlayerId()]
			isRobot := exist && lv != 0
			var duration time.Duration
			if gutils.IsHu(player) {
				duration = time.Second * time.Duration(viper.GetInt(fixed.HuStateTimeOut))
			} else if gutils.IsTing(player) {
				duration = time.Second * time.Duration(viper.GetInt(fixed.TingStateTimeOut))
			} else if deskPlayer.IsTuoguan() {
				duration = 1 * time.Second
			} else if isRobot {
				duration = 1 * time.Second
			} else {
				duration = time.Second * time.Duration(viper.GetInt(fixed.XingPaiTimeOut))
			}
			if (duration != 0 || duration == 0 && (deskPlayer.IsTuoguan() || isRobot)) && time.Now().Sub(startTime) >= duration {
				var aiType AIType
				var eventType int
				if gutils.IsTing(player) || gutils.IsHu(player) || deskPlayer.IsTuoguan() {
					aiType = TuoGuangAI
					eventType = fixed.TuoGuanEvent
				} else if isRobot {
					aiType = RobotAI
					eventType = fixed.RobotEvent
				} else {
					aiType = OverTimeAI
					eventType = fixed.OverTimeEvent
				}

				aiResult, err := AI.GenerateAIEvent(AIEventGenerateParams{
					MajongContext: &mjContext,
					PlayerID:      player.GetPlayerId(),
					AIType:        aiType,
					RobotLv:       lv,
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
