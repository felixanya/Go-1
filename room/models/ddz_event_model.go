package models

import (
	"context"
	"runtime/debug"
	"steve/client_pb/msgid"
	"steve/client_pb/room"
	"steve/entity/poker/ddz"
	"steve/room/ai"
	context2 "steve/room/contexts"
	"steve/room/desk"
	"steve/room/fixed"
	"steve/room/player"
	"steve/room/poker/ddz/ddzmachine"
	"steve/room/poker/ddz/procedure"
	"steve/room/poker/ddz/states"
	"steve/structs/proto/gate_rpc"
	"sync"
	"time"

	fixed2 "steve/datareport/fixed"
	"steve/external/datareportclient"

	"github.com/Sirupsen/logrus"
	"github.com/golang/protobuf/proto"
)

// DDZEventModel 斗地主事件 model
type DDZEventModel struct {
	event chan desk.DeskEvent // 牌桌事件通道
	BaseModel
	closed bool
	mu     sync.Mutex
}

// NewDDZEventModel 创建斗地主事件 model
func NewDDZEventModel(desk *desk.Desk) DeskModel {
	result := &DDZEventModel{}
	result.SetDesk(desk)
	return result
}

// GetName 获取 model 名称
func (model *DDZEventModel) GetName() string {
	return fixed.EventModelName
}

// Active 激活 model
func (model *DDZEventModel) Active() {}

// Start 启动 model
func (model *DDZEventModel) Start() {
	model.event = make(chan desk.DeskEvent, 16)

	model.StartProcessEvents()
}

// StartProcessEvents 开始处理事件
func (model *DDZEventModel) StartProcessEvents() {
	go func() {
		model.processEvents(context.Background())

		gameContext := model.GetGameContext().(*context2.DDZDeskContext)
		ddzContext := gameContext.DDZContext
		players := ddzContext.GetPlayers()
		statistics := make(map[uint64]int64, len(players))

		for _, player := range players {
			if player.GetWin() {
				statistics[player.GetPlayerId()] = 1
			} else {
				statistics[player.GetPlayerId()] = -1
			}
		}
		continueModel := GetContinueModel(model.GetDesk().GetUid())
		continueModel.ContinueDesk(false, 0, statistics)
	}()
	event := desk.DeskEvent{EventID: int(ddz.EventID_event_start_game), EventType: fixed.NormalEvent}
	model.PushEvent(event)
}

// Stop 停止
func (model *DDZEventModel) Stop() {
	model.mu.Lock()
	if model.closed {
		model.mu.Unlock()
		return
	}
	model.closed = true
	close(model.event)
	model.mu.Unlock()
}

// PushEvent 压入事件
func (model *DDZEventModel) PushEvent(event desk.DeskEvent) {
	model.mu.Lock()
	if model.closed {
		model.mu.Unlock()
		return
	}
	model.event <- event
	model.mu.Unlock()
}

// PushRequest 压入玩家请求
func (model *DDZEventModel) PushRequest(playerID uint64, head *steve_proto_gaterpc.Header, bodyData []byte) {
	entry := logrus.WithFields(logrus.Fields{
		"player_id": playerID,
		"msg_id":    head.GetMsgId(),
	})

	trans := GetTranslator()
	eventID, eventData, err := trans.Translate(playerID, head, bodyData)
	if err != nil {
		entry.WithError(err).Errorln("事件转换失败")
		return
	}
	if eventID == 0 {
		entry.Warningln("没有对应事件")
		return
	}
	event := desk.DeskEvent{EventID: eventID, EventType: fixed.NormalEvent, Context: eventData, PlayerID: playerID}
	model.PushEvent(event)
}

func (model *DDZEventModel) processEvents(ctx context.Context) {
	logEntry := logrus.WithFields(logrus.Fields{
		"desk_uid": model.GetDesk().GetUid(),
		"game_id":  model.GetDesk().GetGameId(),
	})
	defer func() {
		if x := recover(); x != nil {
			logEntry.Errorln(x)
			debug.PrintStack()
		}
	}()
	playerModel := GetModelManager().GetPlayerModel(model.GetDesk().GetUid())
	playerEnterChannel := playerModel.getEnterChannel()
	playerLeaveChannel := playerModel.getLeaveChannel()
	tick := time.NewTicker(ai.TickTime)
	defer tick.Stop()

	for {
		select {
		case <-ctx.Done():
			{
				logEntry.Infoln("done")
				return
			}
		case enterInfo := <-playerEnterChannel:
			{
				model.handlePlayerEnter(enterInfo)
			}
		case leaveInfo := <-playerLeaveChannel:
			{
				model.handlePlayerLeave(leaveInfo)
			}
		case event := <-model.event:
			{
				playerMgr := player.GetPlayerMgr()
				eventPlayer := playerMgr.GetPlayer(event.PlayerID)
				if eventPlayer != nil {
					eventPlayer.CountingDown = false
				}

				eventContext := model.getEventContext(event)
				if model.processEvent(event.EventID, eventContext) {
					return
				}
			}
		case <-tick.C:
			{
				events := model.genTimerEvent()
				for _, event := range events {
					context := model.getEventContext(event)
					if model.processEvent(event.EventID, context) {
						return
					}
				}
			}
		}
	}
}

// 处理恢复对局的请求
// eventContext : 事件体
// machine		: 状态机
// ddzContext	: 斗地主牌局信息
// bool 		: 成功/失败
func (model *DDZEventModel) dealResumeRequest(playerID uint64, ddzContext *ddz.DDZContext) error {
	// 请求的玩家ID
	reqPlayerID := playerID

	if !states.IsValidPlayer(ddzContext, reqPlayerID) {
		logrus.WithField("context", ddzContext).WithField("player", reqPlayerID).Warnln("玩家不在本牌桌上!")
		return nil
	}

	// 存在的话则发送游戏信息
	var playersInfo []*room.DDZPlayerInfo

	playerMgr := player.GetPlayerMgr()

	// 把所有的玩家压入playersInfo
	players := ddzContext.GetPlayers()
	for index := 0; index < len(players); index++ {
		player := players[index]

		// Player转为RoomPlayer
		roomPlayerInfo := procedure.TranslateDDZPlayerToRoomPlayer(*player, uint32(index))
		lord := player.GetLord()
		deskPlayer := playerMgr.GetPlayer(player.GetPlayerId())

		tuoguan := deskPlayer.IsTuoguan()

		ddzPlayerInfo := room.DDZPlayerInfo{}

		ddzPlayerInfo.PlayerInfo = &roomPlayerInfo
		ddzPlayerInfo.OutCards = player.GetOutCards()

		// 只发送自己的手牌，其他人的手牌为空
		if player.GetPlayerId() == reqPlayerID {
			ddzPlayerInfo.HandCards = player.GetHandCards()
		} else {
			ddzPlayerInfo.HandCards = []uint32{}
		}
		ddzPlayerInfo.HandCardsCount = proto.Uint32(uint32(len(player.GetHandCards())))

		ddzPlayerInfo.Lord = &lord
		ddzPlayerInfo.Tuoguan = &tuoguan

		// 叫/抢地主
		if ddzContext.CurState == ddz.StateID_state_grab {

			grabLord := room.GrabLordType_GLT_CALLLORD      // 叫地主
			notGrabLord := room.GrabLordType_GLT_NOTCALLORD // 不叫
			grab := room.GrabLordType_GLT_GRAB              // 抢地主
			notGrab := room.GrabLordType_GLT_NOTGRAB        // 不抢
			noneOpe := room.GrabLordType_GLT_NONE           // 未操作

			// 首个叫地主玩家的playerID
			firstPlayerID := ddzContext.GetFirstGrabPlayerId()

			// 自己是否操作过
			grabbed := false
			allOpePlayers := ddzContext.GetGrabbedPlayers()
			for i := 0; i < len(allOpePlayers); i++ {
				if allOpePlayers[i] == player.GetPlayerId() {
					grabbed = true
					break
				}
			}

			// 没人叫过地主
			if firstPlayerID == 0 {
				// 操作过说明不叫，否则说明未操作
				if grabbed {
					ddzPlayerInfo.GrabLord = &notGrabLord
				} else {
					ddzPlayerInfo.GrabLord = &noneOpe
				}
			} else { // 已经有人叫过地主了

				// 自己叫/抢过
				if player.GetGrab() {

					// 首次的是自己，说明是叫地主; 首次的不是自己，说明是抢地主
					if ddzContext.GetFirstGrabPlayerId() == player.GetPlayerId() {
						ddzPlayerInfo.GrabLord = &grabLord
					} else {
						ddzPlayerInfo.GrabLord = &grab
					}
				} else { // 自己没叫/抢过
					// 操作过说明是不叫/不抢，否则说明未操作
					if grabbed {
						// 指定叫地主的是自己，说明是不叫;否则说明是不抢
						if ddzContext.GetCallPlayerId() == player.GetPlayerId() {
							ddzPlayerInfo.GrabLord = &notGrabLord
						} else {
							ddzPlayerInfo.GrabLord = &notGrab
						}
					} else {
						ddzPlayerInfo.GrabLord = &noneOpe
					}
				}
			}
		}

		// 加倍状态
		if ddzContext.CurState == ddz.StateID_state_double {

			double := room.DoubleType_DT_DOUBLE
			notDouble := room.DoubleType_DT_NOTDOUBLE
			noneOpera := room.DoubleType_DT_NONE

			// 加倍
			if player.GetIsDouble() {
				ddzPlayerInfo.Double = &double
			} else {
				// 是否已操作过
				doubled := false
				allOpePlayers := ddzContext.GetDoubledPlayers()
				for i := 0; i < len(allOpePlayers); i++ {
					if allOpePlayers[i] == player.GetPlayerId() {
						doubled = true
						break
					}
				}

				// 存在说明不加倍，不存在说明未操作
				if doubled {
					ddzPlayerInfo.Double = &notDouble
				} else {
					ddzPlayerInfo.Double = &noneOpera
				}
			}
		}

		playersInfo = append(playersInfo, &ddzPlayerInfo)
	}

	// 下面压入公共数据

	// 限制时间
	duration := time.Second * time.Duration(ddzContext.Duration)
	logrus.Debugf("duration = %v", duration)

	// 剩余时间
	leftTime := duration - time.Now().Sub(ddzContext.StartTime)
	logrus.Debugf("leftTime = %v", leftTime)

	if leftTime < 0 {
		leftTime = 0
	}
	leftTimeInt32 := uint32(leftTime.Seconds())

	curStage := room.DDZStage(int32(ddzContext.CurStage))
	curCardType := room.CardType(ddzContext.GetCurCardType())

	totalGrab := ddzContext.GetTotalGrab()
	if totalGrab == 0 {
		totalGrab = 1
	}

	// 斗地主游戏信息
	ddzInfo := room.DDZDeskInfo{
		Players: playersInfo, // 每个人的信息
		Stage: &room.NextStage{
			Stage: &curStage,
			Time:  proto.Uint32(leftTimeInt32),
		},
		CurPlayerId:  proto.Uint64(ddzContext.GetCurrentPlayerId()), // 当前操作的玩家
		Dipai:        ddzContext.GetDipai(),
		TotalGrab:    &totalGrab,
		TotalDouble:  proto.Uint32(ddzContext.GetTotalDouble()),
		TotalBomb:    proto.Uint32(ddzContext.GetTotalBomb()),
		CurCardType:  &curCardType,
		CurCardPivot: proto.Uint32(ddzContext.GetCardTypePivot()),
		CurOutCards:  ddzContext.CurOutCards,
	}

	// 恢复信息
	resumeMsg := &room.RoomResumeGameRsp{
		ResumeRes: room.RoomError_SUCCESS.Enum(),
		GameId:    proto.Uint32(uint32(model.GetDesk().GetGameId())),
		LevelId:   proto.Uint32(uint32(model.GetDesk().GetLevel())),
		DdzInfo:   &ddzInfo,
	}

	messageModel := GetModelManager().GetMessageModel(model.GetDesk().GetUid())
	messageModel.BroadCastDeskMessage([]uint64{reqPlayerID}, msgid.MsgID_ROOM_RESUME_GAME_RSP, resumeMsg, false)
	logrus.WithField("resumeMsg", resumeMsg).WithField("playerId", reqPlayerID).Infoln("斗地主发送恢复对局消息")

	return nil
}

// handlePlayerEnter 处理玩家进入牌桌
func (model *DDZEventModel) handlePlayerEnter(enterInfo playerIDWithChannel) {
	modelMgr := GetModelManager()
	modelMgr.GetPlayerModel(model.GetDesk().GetUid()).handlePlayerEnter(enterInfo.playerID)
	gameContext := model.GetGameContext().(*context2.DDZDeskContext)
	model.dealResumeRequest(enterInfo.playerID, &gameContext.DDZContext)
	close(enterInfo.finishChannel)
}

// handlePlayerLeave 处理玩家离开牌桌
func (model *DDZEventModel) handlePlayerLeave(leaveInfo playerIDWithChannel) {
	modelMgr := GetModelManager()
	playerID := leaveInfo.playerID

	modelMgr.GetPlayerModel(model.GetDesk().GetUid()).handlePlayerLeave(playerID, true)
	logrus.WithField("player_id", playerID).Debugln("玩家退出")
	close(leaveInfo.finishChannel)
}

// getEventPlayerID  获取事件玩家
func (model *DDZEventModel) getEventPlayerID(event desk.DeskEvent) uint64 {
	return event.PlayerID
}

// getEventContext 获取事件现场
func (model *DDZEventModel) getEventContext(event desk.DeskEvent) interface{} {
	return event.Context
}

func (model *DDZEventModel) getMessageSender() ddzmachine.MessageSender {
	messageModel := GetModelManager().GetMessageModel(model.GetDesk().GetUid())
	return func(players []uint64, msgID msgid.MsgID, body proto.Message) error {
		return messageModel.BroadCastDeskMessage(players, msgID, body, true)
	}
}

// processEvent 处理单个事件
// step 1. 调用麻将逻辑的接口来处理事件(返回最新麻将现场, 自动事件， 发送给玩家的消息)， 并且更新 mjContext
// step 2. 将消息发送给玩家
// step 3. 调用 room 的结算逻辑来处理结算
// step 4. 如果有自动事件， 将自动事件写入自动事件通道
// step 5. 如果当前状态是游戏结束状态， 调用 cancel 终止游戏
// 返回： 游戏是否结束
func (model *DDZEventModel) processEvent(eventID int, eventContext interface{}) bool {
	entry := logrus.WithFields(logrus.Fields{
		"event_id": eventID,
	})
	gameContext := model.GetGameContext().(*context2.DDZDeskContext)
	params := procedure.Params{
		Context:      gameContext.DDZContext,
		Sender:       model.getMessageSender(),
		EventID:      eventID,
		EventContext: eventContext,
	}

	result := procedure.HandleEvent(params)
	if !result.Succeed {
		entry.Errorln("处理事件失败")
		return false
	}
	gameContext.DDZContext = result.Context

	// 自动事件不为空，继续处理事件
	if result.HasAutoEvent {
		if result.AutoEventDuration == time.Duration(0) {
			return model.processEvent(result.AutoEventID, result.AutoEventContext)
		}
		go func() {
			timer := time.NewTimer(result.AutoEventDuration)
			<-timer.C
			model.PushEvent(desk.DeskEvent{EventID: result.AutoEventID, EventType: fixed.NormalEvent, Context: result.AutoEventContext})
		}()
	}
	return model.checkGameOver(entry)
}

// checkGameOver 检查游戏结束
func (model *DDZEventModel) checkGameOver(logEntry *logrus.Entry) bool {
	gameContext := model.GetGameContext().(*context2.DDZDeskContext)
	ddzContext := gameContext.DDZContext

	if ddzContext.GetCurState() == ddz.StateID_state_over {
		for _, pID := range model.GetDesk().GetPlayerIds() {
			datareportclient.DataReport(fixed2.LOG_TYPE_GAM, 0, 0, 0, pID, "1")
		}
		return true
	}
	return false
}

// genTimerEvent 生成计时事件
func (model *DDZEventModel) genTimerEvent() []desk.DeskEvent {
	dContext := model.GetDesk().GetConfig().Context.(*context2.DDZDeskContext)
	ddzContext := &dContext.DDZContext

	// 产生AI事件
	result := ai.GetAtEvent().GenerateV2(&ai.AutoEventParams{
		Desk:      model.GetDesk(),
		StartTime: ddzContext.GetStartTime(),
	})
	return result.Events
}
