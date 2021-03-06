package common

import (
	majongpb "steve/entity/majong"
	"steve/room/majong/interfaces"
	"steve/room/majong/utils"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/golang/protobuf/proto"

	"github.com/Sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestZixunState_angang(t *testing.T) {
	mc := gomock.NewController(t)
	flow := interfaces.NewMockMajongFlow(mc)
	flow.EXPECT().PushMessages(gomock.Any(), gomock.Any()).AnyTimes()
	flow.EXPECT().GetMajongContext().Return(
		&majongpb.MajongContext{
			GameId: 1,
			Players: []*majongpb.Player{
				&majongpb.Player{
					PlayerId:  1,
					HandCards: []*majongpb.Card{&Card1W, &Card1W, &Card1W, &Card1W, &Card2W, &Card2W, &Card2W, &Card2W, &Card3W, &Card3W, &Card3W, &Card3W, &Card4W, &Card4W},
				},
			},
			MopaiPlayer: 1,
			// LastMopaiPlayer: 1,
			WallCards: []*majongpb.Card{&Card1T},
		},
	).AnyTimes()

	// s := ZiXunState{}
	gangRequestEvent := &majongpb.GangRequestEvent{
		Head: &majongpb.RequestEventHead{
			PlayerId: 1,
		},
		Card: &Card1W,
	}
	context := flow.GetMajongContext()
	player := utils.GetPlayerByID(context.GetPlayers(), context.GetMopaiPlayer())
	logrus.WithFields(FmtPlayerInfo(player)).Info("暗杠前")
	// beforeResults := ""
	// beforeResults += fmt.Sprintln("before暗杠：")
	// beforeResults += FmtPlayerInfo(player)
	// logrus.Info(beforeResults)
	// stateID, err := s.ProcessEvent(majongpb.EventID_event_gang_request, requestEvent, flow)
	flow.ProcessEvent(majongpb.EventID_event_gang_request, gangRequestEvent)
	// assert.Nil(t, err)
	// assert.Equal(t, majongpb.StateID_state_angang, stateID, "执行暗杠操作成功后，状态应该为暗杠状态")
	// results := ""
	// results += fmt.Sprintln("after暗杠：")
	// results += FmtPlayerInfo(player)
	// logrus.Info(results)
	logrus.WithFields(FmtPlayerInfo(player)).Info("暗杠后")
}

func TestZixunState_zimo(t *testing.T) {
	mc := gomock.NewController(t)
	flow := interfaces.NewMockMajongFlow(mc)
	flow.EXPECT().PushMessages(gomock.Any(), gomock.Any()).AnyTimes()
	flow.EXPECT().GetMajongContext().Return(
		&majongpb.MajongContext{
			GameId: 1,
			Players: []*majongpb.Player{
				&majongpb.Player{
					PlayerId:        1,
					HandCards:       []*majongpb.Card{&Card1W, &Card1W, &Card1W, &Card1W, &Card2W, &Card2W, &Card2W, &Card2W, &Card3W, &Card3W, &Card3W, &Card3W, &Card4W, &Card4W},
					PossibleActions: []majongpb.Action{majongpb.Action_action_gang, majongpb.Action_action_hu},
					DingqueColor:    majongpb.CardColor_ColorTiao,
				},
			},
			ActivePlayer: 1,
			WallCards:    []*majongpb.Card{&Card1T},
		},
	).AnyTimes()

	s := ZiXunState{}
	huRequestEvent := &majongpb.HuRequestEvent{
		Head: &majongpb.RequestEventHead{
			PlayerId: 1,
		},
	}

	requestEvent, err := proto.Marshal(huRequestEvent)
	context := flow.GetMajongContext()
	player := utils.GetPlayerByID(context.GetPlayers(), context.GetActivePlayer())
	logrus.WithFields(FmtPlayerInfo(player)).Info("自摸前")
	// beforeResults := ""
	// beforeResults += fmt.Sprintln("before自摸：")
	// beforeResults += FmtPlayerInfo(player)
	// logrus.Info(beforeResults)

	stateID, err := s.ProcessEvent(majongpb.EventID_event_hu_request, requestEvent, flow)
	assert.Nil(t, err)
	assert.Equal(t, majongpb.StateID_state_zimo, stateID, "执行自摸操作成功后，状态应该为自摸状态")
	logrus.WithFields(FmtPlayerInfo(player)).Info("自摸后")
	// results := ""
	// results += fmt.Sprintln("after自摸：")
	// results += FmtPlayerInfo(player)
	// logrus.Info(results)
}

func TestZixunState_bugang(t *testing.T) {
	mc := gomock.NewController(t)
	flow := interfaces.NewMockMajongFlow(mc)
	// playersID := []uint64{1}
	// ntf := &room.RoomBugangNtf{
	// 	Player: proto.Uint64(1),
	// 	Card: &room.Card{
	// 		Color: room.CardColor_CC_WAN.Enum(),
	// 		Point: proto.Int32(1),
	// 	},
	// }
	// toClientMessage := interfaces.ToClientMessage{
	// 	MsgID: int(msgid.MsgID_room_bugang_ntf),
	// 	Msg:   ntf,
	// }
	// flow.EXPECT().PushMessages(playersID, toClientMessage).DoAndReturn(
	// 	func(playerIDs []uint64, msgs ...interfaces.ToClientMessage) {},
	// )
	flow.EXPECT().PushMessages(gomock.Any(), gomock.Any()).AnyTimes()
	flow.EXPECT().GetMajongContext().Return(
		&majongpb.MajongContext{
			GameId: 1,
			Players: []*majongpb.Player{
				&majongpb.Player{
					PlayerId:        1,
					HandCards:       []*majongpb.Card{&Card1W, &Card2W, &Card2W, &Card2W, &Card2W, &Card3W, &Card3W, &Card3W, &Card3W, &Card4W, &Card4W},
					PossibleActions: []majongpb.Action{majongpb.Action_action_gang, majongpb.Action_action_hu},
					DingqueColor:    majongpb.CardColor_ColorTiao,
					PengCards: []*majongpb.PengCard{
						&majongpb.PengCard{
							Card:      &Card1W,
							SrcPlayer: 2,
						},
					},
				},
			},
			ActivePlayer: 1,
			WallCards:    []*majongpb.Card{&Card1T},
		},
	).AnyTimes()

	s := ZiXunState{}
	gangRequestEvent := &majongpb.BugangRequestEvent{
		Head: &majongpb.RequestEventHead{
			PlayerId: 1,
		},
		Cards: &Card1W,
	}

	context := flow.GetMajongContext()
	player := utils.GetPlayerByID(context.GetPlayers(), context.GetActivePlayer())
	logrus.WithFields(FmtPlayerInfo(player)).Info("补杠前")
	// beforeResults := ""
	// beforeResults += fmt.Sprintln("before补杠：")
	// beforeResults += FmtPlayerInfo(player)
	// logrus.Info(beforeResults)
	stateID, err := s.ProcessEvent(majongpb.EventID_event_gang_request, gangRequestEvent, flow)
	assert.Nil(t, err)
	assert.Equal(t, majongpb.StateID_state_bugang, stateID, "执行补杠操作成功后，状态应该为补杠状态")
	logrus.WithFields(FmtPlayerInfo(player)).Info("补杠后")
	// results := ""
	// results += fmt.Sprintln("after补杠：")
	// results += FmtPlayerInfo(player)
	// logrus.Info(results)
}

func TestZixunState_chupai(t *testing.T) {
	mc := gomock.NewController(t)
	flow := interfaces.NewMockMajongFlow(mc)
	// playersID := []uint64{1}
	// ntf := &room.RoomChupaiNtf{
	// 	Player: proto.Uint64(1),
	// 	Card: &room.Card{
	// 		Color: room.CardColor_CC_WAN.Enum(),
	// 		Point: proto.Int32(1),
	// 	},
	// }
	// toClientMessage := interfaces.ToClientMessage{
	// 	MsgID: int(msgid.MsgID_room_chupai_ntf),
	// 	Msg:   ntf,
	// }
	// flow.EXPECT().PushMessages(playersID, toClientMessage).DoAndReturn(
	// 	func(playerIDs []uint64, msgs ...interfaces.ToClientMessage) {},
	// )
	flow.EXPECT().PushMessages(gomock.Any(), gomock.Any()).AnyTimes()
	flow.EXPECT().GetMajongContext().Return(
		&majongpb.MajongContext{
			GameId: 1,
			Players: []*majongpb.Player{
				&majongpb.Player{
					PlayerId:        1,
					HandCards:       []*majongpb.Card{&Card1W, &Card2W, &Card2W, &Card2W, &Card2W, &Card3W, &Card3W, &Card3W, &Card3W, &Card4W, &Card4W},
					PossibleActions: []majongpb.Action{majongpb.Action_action_gang, majongpb.Action_action_hu},
					DingqueColor:    majongpb.CardColor_ColorTiao,
					PengCards: []*majongpb.PengCard{
						&majongpb.PengCard{
							Card:      &Card1W,
							SrcPlayer: 2,
						},
					},
				},
			},
			ActivePlayer: 1,
			WallCards:    []*majongpb.Card{&Card1T},
		},
	).AnyTimes()

	s := ZiXunState{}
	bugangRequestEvent := &majongpb.ChupaiRequestEvent{
		Head: &majongpb.RequestEventHead{
			PlayerId: 1,
		},
		Cards: &Card1W,
	}
	context := flow.GetMajongContext()
	player := utils.GetPlayerByID(context.GetPlayers(), context.GetActivePlayer())
	logrus.WithFields(FmtPlayerInfo(player)).Info("出牌前")
	// beforeResults := ""
	// beforeResults += fmt.Sprintln("before出牌：")
	// beforeResults += FmtPlayerInfo(player)
	// logrus.Info(beforeResults)
	// stateID, err := s.bugang(flow, bugangRequestEvent)
	stateID, err := s.ProcessEvent(majongpb.EventID_event_chupai_request, bugangRequestEvent, flow)
	assert.Nil(t, err)
	assert.Equal(t, majongpb.StateID_state_chupai, stateID, "执行出牌操作成功后，状态应该为出牌状态")
	logrus.WithFields(FmtPlayerInfo(player)).Info("出牌后")
	// results := ""
	// results += fmt.Sprintln("after出牌：")
	// results += FmtPlayerInfo(player)
	// logrus.Info(results)
}

// 1w 1w 1w 1w 2w 2w 2w 2w 3w 3w 9t 3w 4w 3w
func Test_CheckAngang(t *testing.T) {
	player := &majongpb.Player{}
	player.HuCards = []*majongpb.HuCard{}
	player.DingqueColor = majongpb.CardColor_ColorTiao
	player.HandCards = []*majongpb.Card{
		&Card1W, &Card1W, &Card1W, &Card1W, &Card2W, &Card2W, &Card2W, &Card2W, &Card3W, &Card3W, &Card9T, &Card3W, &Card4W, &Card3W,
	}
	s := ZiXunState{}
	enables := s.checkPlayerAngang(player)
	assert.Equal(t, 3, len(enables))
	assert.Subset(t, enables, []uint32{11, 12, 13})
}
