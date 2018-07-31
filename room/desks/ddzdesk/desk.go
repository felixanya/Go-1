package ddzdesk

import (
	"steve/client_pb/msgid"
	"steve/entity/poker/ddz"
	"steve/room/desks/ddzdesk/flow/ddz/ddzmachine"
	"steve/room/desks/ddzdesk/flow/ddz/procedure"
	"steve/room/desks/deskbase"
	"steve/room/interfaces"
	"steve/room/interfaces/facade"
	"steve/room/interfaces/global"
	"steve/structs/proto/gate_rpc"
	"time"

	"context"
	"runtime/debug"

	"github.com/Sirupsen/logrus"
	"github.com/golang/protobuf/proto"
)

// deskEvent 牌桌事件
type deskEvent struct {
	eventID      int
	eventContext interface{}
	eventType    interfaces.EventType
	playerID     uint64
}

// desk 斗地主牌桌
type desk struct {
	*deskbase.DeskBase
	eventChannel   chan deskEvent
	closingChannel chan struct{}
	ddzContext     *ddz.DDZContext
	cancel         context.CancelFunc // 取消事件处理
}

// initDDZContext 初始化斗地主现场
func (d *desk) initDDZContext() {

	// 牌桌所有玩家的playerID
	// index:座位号 value:playerID
	playersID := facade.GetDeskPlayerIDs(d)

	d.ddzContext = procedure.CreateInitDDZContext(playersID)
}

// Start 启动牌桌逻辑
// finish : 当牌桌逻辑完成时调用
func (d *desk) Start(finish func()) error {
	d.eventChannel = make(chan deskEvent, 4)
	d.closingChannel = make(chan struct{})

	// 初始化操作
	d.initDDZContext()

	// 逻辑线程
	go func() {
		defer func() {
			if x := recover(); x != nil {
				logrus.Errorln(x)
				debug.PrintStack()
			}
		}()

		// 开始运行
		d.run()

		// 执行结束时的函数
		finish()
	}()

	// 定时器线程
	var ctx context.Context
	ctx, d.cancel = context.WithCancel(context.Background())
	go func() {
		defer func() {
			if x := recover(); x != nil {
				logrus.Errorln(x)
				debug.PrintStack()
			}
		}()

		// 定时器
		d.timerTask(ctx)
	}()

	// 游戏开始事件
	d.pushEvent(&deskEvent{
		eventID: int(ddz.EventID_event_start_game),
	})

	return nil
}

// timerTask 定时任务，产生自动事件
func (d *desk) timerTask(ctx context.Context) {
	defer func() {
		if x := recover(); x != nil {
			debug.PrintStack()
		}
	}()

	// 200毫秒的定时器
	t := time.NewTicker(time.Millisecond * 200)

	defer t.Stop()

	for {
		select {
		case <-t.C:
			{
				d.genTimerEvent()
			}
		case <-ctx.Done():
			{
				return
			}
		}
	}
}

// genTimerEvent 生成计时事件，定时器触发时调用
func (d *desk) genTimerEvent() {
	g := global.GetDeskAutoEventGenerator()
	// 先将 context 指针读出来拷贝， 后面的 context 修改都会分配一块新的内存
	dContext := d.ddzContext

	// 牌桌的所有托管玩家
	tuoGuanPlayers := facade.GetTuoguanPlayers(d)

	logEntry := logrus.WithFields(logrus.Fields{
		"func_name":       "desk.genTimerEvent",
		"tuoguan_players": tuoGuanPlayers,
	})

	// 开始时间
	startTime := time.Time{}
	startTime.UnmarshalBinary(dContext.StartTime)

	// 产生AI事件
	result := g.GenerateV2(&interfaces.AutoEventGenerateParams{
		Desk:       d,
		DDZContext: dContext,
		PlayerIds:  dContext.CountDownPlayers,
		StartTime:  startTime,
		Duration:   dContext.Duration,
		RobotLv:    map[uint64]int{},
	})

	// 把AI事件转换为斗地主的牌桌事件
	for _, event := range result.Events {
		logEntry.WithFields(logrus.Fields{
			"event_id":     event.ID,
			"event_player": event.PlayerID,
			"event_type":   event.EventType,
		}).Debugln("注入计时事件")
		d.eventChannel <- deskEvent{
			eventID:      int(event.ID),
			eventContext: event.Context,
			eventType:    event.EventType,
			playerID:     event.PlayerID,
		}
	}
}

// Stop 停止牌桌
func (d *desk) Stop() error {
	d.cancel()
	d.closingChannel <- struct{}{}
	return nil
}

// PushRequest 压入玩家请求
func (d *desk) PushRequest(playerID uint64, head *steve_proto_gaterpc.Header, bodyData []byte) {
	entry := logrus.WithFields(logrus.Fields{
		"func_name": "desk.PushRequest",
		"player_id": playerID,
		"msg_id":    head.GetMsgId(),
	})

	translator := global.GetReqEventTranslator()
	eventID, eventData, err := translator.Translate(playerID, head, bodyData)
	if err != nil {
		entry.WithError(err).Errorln("事件转换失败")
		return
	}
	if eventID == 0 {
		entry.Warningln("没有对应事件")
		return
	}

	d.pushEvent(&deskEvent{
		eventID:      eventID,
		eventContext: eventData,
	})
}

func (d *desk) pushEvent(e *deskEvent) {
	d.eventChannel <- *e
}

// PushEvent 压入事件
func (d *desk) PushEvent(event interfaces.Event) {
	return
}

// run 执行牌桌逻辑
func (d *desk) run() {
	defer d.consumeAllEnterQuit() // 消费完所有的退出进入数据

forstart:
	for {
		select {
		case event := <-d.eventChannel:
			{
				d.processEvent(&event)
				d.recordTuoguanOverTimeCount(event)
			}
		case enterQuitInfo := <-d.PlayerEnterQuitChannel():
			{
				d.handleEnterQuit(enterQuitInfo)
			}
		case <-d.closingChannel:
			{
				break forstart
			}
		}
	}
}

func (d *desk) consumeAllEnterQuit() {
	for {
		select {
		case enterQuitInfo := <-d.PlayerEnterQuitChannel():
			{
				d.handleEnterQuit(enterQuitInfo)
			}
		default:
			return
		}
	}
}

// handleEnterQuit 处理退出进入信息
func (d *desk) handleEnterQuit(eqi interfaces.PlayerEnterQuitInfo) {
	logEntry := logrus.WithFields(logrus.Fields{
		"func_name": "handleEnterQuit",
		"player_id": eqi.PlayerID,
		"quit":      eqi.Quit,
	})
	deskPlayer := facade.GetDeskPlayerByID(d, eqi.PlayerID)
	defer close(eqi.FinishChannel)

	if deskPlayer == nil {
		logEntry.Errorln("玩家不在牌桌上")
		return
	}
	if eqi.Quit {
		deskPlayer.SetTuoguan(true, true)
		logEntry.Debugln("玩家退出")
	} else {
		deskPlayer.SetTuoguan(false, true)

		//生成恢复对局事件
		eventMessage := &ddz.ResumeRequestEvent{
			Head: &ddz.RequestEventHead{PlayerId: eqi.PlayerID},
		}
		eventID := int(ddz.EventID_event_resume_request)
		d.processEvent(&deskEvent{eventID: eventID, eventContext: eventMessage})
		logEntry.Debugln("玩家进入")
	}
}

// recordTuoguanOverTimeCount 记录托管超时计数
func (d *desk) recordTuoguanOverTimeCount(event deskEvent) {
	if event.eventType != interfaces.OverTimeEvent {
		return
	}
	playerID := event.playerID
	if playerID == 0 {
		return
	}
	deskPlayer := facade.GetDeskPlayerByID(d, playerID)
	if deskPlayer != nil {
		deskPlayer.OnPlayerOverTime()
	}
}

func (d *desk) processEvent(e *deskEvent) {
	entry := logrus.WithFields(logrus.Fields{
		"func_name": "desk.processEvent",
		"event_id":  e.eventID,
	})

	params := procedure.Params{
		PlayerMgr:    d.DeskPlayerMgr,
		Context:      *d.ddzContext,
		Sender:       d.getMessageSender(), //TODO: 尽量不要把一个参数拆成多个参数
		EventID:      e.eventID,
		EventContext: e.eventContext,
	}

	result := procedure.HandleEvent(params)
	if !result.Succeed {
		entry.Errorln("处理事件失败")
		return
	}

	d.ddzContext = &result.Context
	// 游戏结束
	if d.ddzContext.GetCurState() == ddz.StateID_state_over {
		d.cancelTuoguanGameOver()
		d.ContinueDesk(false, 0, d.getWinners())
		go func() { d.Stop() }()
		return
	}
	if result.HasAutoEvent {
		if result.AutoEventDuration == time.Duration(0) {
			d.pushEvent(&deskEvent{
				eventID:      result.AutoEventID,
				eventContext: result.AutoEventContext,
			})
		} else {
			go func() {
				timer := time.NewTimer(result.AutoEventDuration)
				<-timer.C
				d.pushEvent(&deskEvent{
					eventID:      result.AutoEventID,
					eventContext: result.AutoEventContext,
				})
			}()
		}
	}
}

func (d *desk) cancelTuoguanGameOver() {
	players := d.GetDeskPlayers()
	for _, player := range players {
		if player.IsTuoguan() {
			player.SetTuoguan(false, true)
		}
	}
}

func (d *desk) getWinners() []uint64 {
	players := d.ddzContext.GetPlayers()
	winners := make([]uint64, 0, len(players))

	for _, player := range players {
		if player.GetWin() {
			winners = append(winners, player.GetPlayerId())
		}
	}
	return winners
}

func (d *desk) getMessageSender() ddzmachine.MessageSender {
	return func(players []uint64, msgID msgid.MsgID, body proto.Message) error {
		return facade.BroadCastDeskMessage(d, players, msgID, body, true)
	}
}

// TODO 待优化
func (d *desk) ChangePlayer(playerID uint64) error {
	return nil
}
