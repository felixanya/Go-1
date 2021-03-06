package states

import (
	"steve/back/logic"
	"steve/client_pb/msgid"
	"steve/client_pb/room"
	"steve/common/constant"
	"steve/entity/gamelog"
	"steve/entity/poker/ddz"
	"steve/external/goldclient"
	"steve/room/poker/machine"
	server_gold "steve/server_pb/gold"

	"github.com/Sirupsen/logrus"
	"github.com/golang/protobuf/proto"
)

type settleState struct{}

func (s *settleState) OnEnter(m machine.Machine) {
	logrus.WithField("context", getDDZContext(m)).Debugln("进入Settle状态")
}

func (s *settleState) OnExit(m machine.Machine) {
	logrus.WithField("context", getDDZContext(m)).Debugln("离开Settle状态")
}

func (s *settleState) OnEvent(m machine.Machine, event machine.Event) (int, error) {
	if event.EventID == int(ddz.EventID_event_showhand_finish) {
		s.settle(m)
		return int(ddz.StateID_state_over), nil
	}
	return int(ddz.StateID_state_settle), nil
}

func (s *settleState) settle(m machine.Machine) {
	context := getDDZContext(m)
	context.CurStage = ddz.DDZStage_DDZ_STAGE_OVER

	totalGrab := context.TotalGrab
	totalDouble := context.TotalDouble
	base := context.BaseScore
	multiple := uint64(totalGrab * totalDouble * context.TotalBomb)
	if context.Spring || context.AntiSpring {
		multiple = multiple * 2
	}
	score := base * multiple
	winnerId := context.WinnerId
	lordId := context.LordPlayerId
	lordWin := winnerId == lordId

	//找出每个人最大输赢金币
	maxScores := make(map[uint64]uint64)
	for _, player := range context.GetPlayers() {
		playerId := player.PlayerId
		coin, err := goldclient.GetGold(playerId, int16(server_gold.GoldType_GOLD_COIN))
		if err != nil {
			logrus.Errorf("settle 获取金币失败：%v", err.Error())
		}
		s := If(playerId == lordId, score*2, score).(uint64)
		maxScores[playerId] = If(s > uint64(coin), uint64(coin), s).(uint64)
	}

	settleScores := make(map[uint64]uint64)

	lordMax := maxScores[lordId]
	score = lordMax / 2
	lordScore := uint64(0)
	for _, player := range context.GetPlayers() {
		playerId := player.PlayerId
		if playerId == lordId {
			continue
		}
		maxScore := maxScores[playerId]
		settleScore := If(score > maxScore, maxScore, score).(uint64)
		settleScores[playerId] = settleScore
		lordScore += settleScore
	}
	if lordMax%2 == 1 { //如果地主金豆数是奇数
		if lordWin {
			lordScore-- //少赢一分(赢的分不能超过本人金豆上限)
		} else {
			lordScore++ //多扣一分，系统扣成偶数(华华的要求)
		}
	}
	settleScores[lordId] = lordScore

	var billPlayers []*room.DDZBillPlayerInfo
	for _, player := range context.GetPlayers() {
		playerId := player.PlayerId
		billPlayer := room.DDZBillPlayerInfo{}
		billPlayer.PlayerId = &playerId
		var isWin bool
		var mul int32
		if playerId == lordId {
			isWin = lordWin
			mul = int32(multiple * 2)
		} else {
			isWin = !lordWin
			mul = int32(multiple)
		}
		player.Win = isWin
		billPlayer.Win = &isWin
		billPlayer.Base = proto.Int32(int32(base))
		billPlayer.Multiple = &mul
		settleScore := settleScores[playerId]
		var amount int64
		if isWin {
			amount = int64(settleScore)
			goldclient.AddGold(playerId, int16(server_gold.GoldType_GOLD_COIN), amount, int32(constant.GFGAMESETTLE), 0, 0, 0) //赢钱
		} else {
			amount = int64(-settleScore)
			goldclient.AddGold(playerId, int16(server_gold.GoldType_GOLD_COIN), amount, int32(constant.GFGAMESETTLE), 0, 0, 0) //输钱
		}
		billPlayer.Score = proto.Int64(int64(settleScore))
		gold, err := goldclient.GetGold(playerId, int16(server_gold.GoldType_GOLD_COIN))
		if err != nil {
			logrus.Errorf("settle 获取金币失败：%v", err.Error())
		}
		billPlayer.CurrentScore = proto.Int64(int64(gold))
		billPlayer.Lord = &player.Lord
		billPlayer.OutCards = player.OutCards
		billPlayer.HandCards = player.HandCards
		billPlayers = append(billPlayers, &billPlayer)

		gameDetail := gamelog.TGameDetail{
			Playerid: playerId,
			Gameid:   int(context.GetGameId()),
			Amount:   amount,
			MaxTimes: uint32(mul),
		}
		logic.UpdatePlayerGameInfo(gameDetail)
	}

	antiSpring := !context.Spring && context.AntiSpring

	settleMsg := &room.DDZGameOverNtf{
		WinnerId:     &context.WinnerId,
		ShowHandTime: proto.Uint32(4),
		Spring:       &context.Spring,
		AntiSpring:   &antiSpring,
		Bills:        billPlayers,
	}
	broadcast(m, msgid.MsgID_ROOM_DDZ_GAME_OVER_NTF,
		settleMsg,
	)
	logrus.WithField("settle msg", settleMsg).Debugln("斗地主广播结算消息")
}
