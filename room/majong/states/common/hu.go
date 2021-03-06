package common

import (
	"steve/client_pb/msgid"
	"steve/client_pb/room"
	majongpb "steve/entity/majong"
	"steve/gutils"
	"steve/room/majong/fantype"
	"steve/room/majong/global"
	"steve/room/majong/interfaces"
	"steve/room/majong/interfaces/facade"
	"steve/room/majong/utils"

	"github.com/Sirupsen/logrus"
	"github.com/golang/protobuf/proto"
)

// HuState 胡状态
// 进入胡状态时， 执行胡操作。设置胡完成事件
// 收到胡完成事件时，设置摸牌玩家，返回摸牌状态
type HuState struct {
}

var _ interfaces.MajongState = new(HuState)

// ProcessEvent 处理事件
func (s *HuState) ProcessEvent(eventID majongpb.EventID, eventContext interface{}, flow interfaces.MajongFlow) (newState majongpb.StateID, err error) {
	if eventID == majongpb.EventID_event_hu_finish {
		return majongpb.StateID_state_hu_settle, nil
	}
	return majongpb.StateID_state_hu, global.ErrInvalidEvent
}

// OnEntry 进入状态
func (s *HuState) OnEntry(flow interfaces.MajongFlow) {
	s.doHu(flow)
	flow.SetAutoEvent(majongpb.AutoEvent{
		EventId:      majongpb.EventID_event_hu_finish,
		EventContext: nil,
	})
}

// OnExit 退出状态
func (s *HuState) OnExit(flow interfaces.MajongFlow) {

}

// addHuCard 添加胡的牌
func (s *HuState) addHuCard(card *majongpb.Card, player *majongpb.Player, srcPlayerID uint64, isReal bool) {
	AddHuCard(card, player, srcPlayerID, majongpb.HuType_hu_dianpao, isReal)
}

// doHu 执行胡操作
func (s *HuState) doHu(flow interfaces.MajongFlow) {
	mjContext := flow.GetMajongContext()
	card := mjContext.GetLastOutCard()
	players := mjContext.GetLastHuPlayers()
	srcPlayerID := mjContext.GetLastChupaiPlayer()
	srcPlayer := utils.GetMajongPlayer(srcPlayerID, mjContext)

	logEntry := logrus.WithFields(logrus.Fields{
		"func_name": "HuState.doHu",
	})
	logEntry = utils.WithMajongContext(logEntry, mjContext)

	// 从被碰玩家的outCards移除被碰牌
	srcOutCards := srcPlayer.GetOutCards()
	srcPlayer.OutCards = removeLastCard(logEntry, srcOutCards, card)

	realPlayerID := utils.GetRealHuCardPlayer(players, srcPlayerID, mjContext)
	huType := -1

	for _, playerID := range players {
		player := utils.GetMajongPlayer(playerID, mjContext)
		isReal := false
		if playerID == realPlayerID {
			isReal = true
		}
		s.addHuCard(card, player, srcPlayerID, isReal)
		// 玩家胡状态
		player.XpState = player.GetXpState() | majongpb.XingPaiState_hu

		if huType == -1 {
			huCard := player.GetHuCards()[len(player.GetHuCards())-1]
			fanTypes, _, _ := fantype.CalculateFanTypes(mjContext, playerID, player.GetHandCards(), huCard)
			cardOptionID := int(mjContext.GetCardtypeOptionId())
			huType = int(gutils.ServerFanType2ClientHuType(cardOptionID, fanTypes))
		}
	}
	s.notifyHu(flow, realPlayerID, huType)
	gutils.SetNextZhuangIndex(players, srcPlayerID, mjContext)
	return
}

// isAfterGang 是否为杠后炮
// 杠后摸牌、自询出牌则为杠后炮
func (s *HuState) isAfterGang(mjContext *majongpb.MajongContext) bool {
	zxType := mjContext.GetZixunType()
	mpType := mjContext.GetMopaiType()
	return mpType == majongpb.MopaiType_MT_GANG && zxType == majongpb.ZixunType_ZXT_NORMAL
}

// HuState 广播胡
func (s *HuState) notifyHu(flow interfaces.MajongFlow, realPlayerID uint64, rhuType int) {
	mjContext := flow.GetMajongContext()
	huType := room.HuType_HT_DIANPAO.Enum()
	if rhuType != -1 {
		huType = room.HuType(int32(rhuType)).Enum()
	}
	if s.isAfterGang(mjContext) {
		huType = room.HuType_HT_GANGHOUPAO.Enum()
	}
	body := room.RoomHuNtf{
		Players:      mjContext.GetLastHuPlayers(),
		FromPlayerId: proto.Uint64(mjContext.GetLastChupaiPlayer()),
		Card:         proto.Uint32(uint32(utils.ServerCard2Number(mjContext.GetLastOutCard()))),
		HuType:       huType,
		RealPlayerId: proto.Uint64(realPlayerID),
	}
	facade.BroadcaseMessage(flow, msgid.MsgID_ROOM_HU_NTF, &body)
}
