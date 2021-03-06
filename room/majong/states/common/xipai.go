package common

import (
	"math/rand"
	"steve/client_pb/msgid"
	"steve/client_pb/room"
	majongpb "steve/entity/majong"
	"steve/room/majong/global"
	"steve/room/majong/interfaces"
	"steve/room/majong/interfaces/facade"
	"steve/room/majong/utils"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/golang/protobuf/proto"
)

// XipaiState 洗牌状态
type XipaiState struct {
}

var _ interfaces.MajongState = new(XipaiState)

// ProcessEvent 处理事件
func (s *XipaiState) ProcessEvent(eventID majongpb.EventID, eventContext interface{}, flow interfaces.MajongFlow) (newState majongpb.StateID, err error) {
	if eventID == majongpb.EventID_event_xipai_finish {
		return majongpb.StateID(majongpb.StateID_state_fapai), nil
	}
	return majongpb.StateID(majongpb.StateID_state_xipai), nil
}

func (s *XipaiState) genOriginCards(flow interfaces.MajongFlow) []*majongpb.Card {
	mjContext := flow.GetMajongContext()
	return global.GetOriginCards(mjContext)
}

func (s *XipaiState) xipai(flow interfaces.MajongFlow) []*majongpb.Card {
	cards := s.genOriginCards(flow)
	rand.Seed(int64(time.Now().Nanosecond()))
	rand.Shuffle(len(cards), func(i, j int) {
		cards[i], cards[j] = cards[j], cards[i]
	})
	mjContext := flow.GetMajongContext()
	mjContext.CardTotalNum = uint32(len(cards))
	utils.PeiPai(cards, mjContext.GetOption().GetCards())

	logrus.WithFields(logrus.Fields{
		"func_name":    "XipaiState.xipai",
		"option_cards": mjContext.GetOption().GetCards(),
	}).Debugln("洗牌")

	return cards
}

// randDices 随机筛子
func (s *XipaiState) randDices() []uint32 {
	rand.Seed(int64(time.Now().Nanosecond()))
	return []uint32{
		uint32(rand.Int31n(6)) + 1,
		uint32(rand.Int31n(6)) + 1,
	}
}

// selectZhuangjia 选择庄家
// func (s *XipaiState) selectZhuangjia(mjContext *majongpb.MajongContext, dices []uint32, gameID int) int {
// 	totalDice := int(dices[0] + dices[1])

// 	mjContext.ZhuangjiaIndex = uint32(totalDice % len(mjContext.Players))
// 	zhuang := mjContext.GetOption().GetZhuang()
// 	if zhuang.GetNeedDeployZhuang() {
// 		mjContext.ZhuangjiaIndex = uint32(zhuang.GetZhuangIndex() % int32(len(mjContext.Players)))
// 	}
// 	return int(mjContext.ZhuangjiaIndex)
// }

// pushMessages 发送消息给玩家
func (s *XipaiState) pushMessages(cardCount int, dices []uint32, zjIndex uint32, flow interfaces.MajongFlow) {
	facade.BroadcaseMessage(flow, msgid.MsgID_ROOM_XIPAI_NTF, &room.RoomXipaiNtf{
		Dices:      dices,
		TotalCard:  proto.Uint32(uint32(cardCount)),
		BankerSeat: proto.Uint32(zjIndex),
	})
}

// OnEntry 进入状态
func (s *XipaiState) OnEntry(flow interfaces.MajongFlow) {
	mjContext := flow.GetMajongContext()

	mjContext.WallCards = s.xipai(flow)
	dices := s.randDices()
	mjContext.Dices = append(mjContext.Dices, dices...)
	s.pushMessages(len(mjContext.WallCards), dices, mjContext.GetZhuangjiaIndex(), flow)
	flow.SetAutoEvent(majongpb.AutoEvent{
		EventId:      majongpb.EventID_event_xipai_finish,
		EventContext: nil,
		// WaitTime:     mjContext.GetOption().GetMaxCartoonTime(),
	})
}

// OnExit 退出状态
func (s *XipaiState) OnExit(flow interfaces.MajongFlow) {
}
