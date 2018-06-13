package desks

import (
	"steve/client_pb/room"
	"steve/gutils"
	server_pb "steve/server_pb/majong"
	"time"

	"github.com/golang/protobuf/proto"
)

func getStateCostTime(entryTime int64) (costTime uint32) {
	nowTime := time.Now().Unix()
	if nowTime > entryTime {
		costTime = uint32(nowTime - entryTime)
	}
	return
}

func getOperatePlayerID(mjContext *server_pb.MajongContext) *uint64 {
	state := mjContext.GetCurState()
	var playerID uint64
	switch state {
	case 1: // 自询状态 刚开局是庄家，其他情况是最近摸牌者
		playerID = mjContext.GetLastMopaiPlayer()
	case 2: // 他询状态， 出牌者
		playerID = mjContext.GetLastChupaiPlayer()
	case 3: // 等待抢杠胡 杠牌玩家
		playerID = mjContext.GetLastGangPlayer()
	default:
		return nil
	}
	return &playerID
}

func getGameStage(curState server_pb.StateID) (stage room.GameStage) {
	switch curState {
	case server_pb.StateID_state_huansanzhang:
		stage = room.GameStage_GAMESTAGE_HUANSANZHANG
	case server_pb.StateID_state_dingque:
		stage = room.GameStage_GAMESTAGE_DINGQUE
	default:
		stage = room.GameStage_GAMESTAGE_PLAYCARD
	}
	return
}

func getDoorCard(mjContext *server_pb.MajongContext) *uint32 {
	if mjContext.GetCurState() == server_pb.StateID_state_zixun {
		DoorCard := uint32(mjContext.GetLastMopaiCard().GetPoint())
		return &DoorCard
	}
	return nil
}

func getRecoverPlayerInfo(d *desk) (recoverPlayerInfo []*room.GamePlayerInfo) {
	mjContext := &d.dContext.mjContext
	roomPlayerInfos := d.GetPlayers()
	for _, roomPlayerInfo := range roomPlayerInfos {
		var player *server_pb.Player
		// 这里假设总能找到一个对应玩家
		for _, player = range mjContext.GetPlayers() {
			if player.GetPalyerId() == roomPlayerInfo.GetPlayerId() {
				break
			}
		}
		playerID := player.GetPalyerId()
		svrHandCard := player.GetHandCards()
		handCardCount := uint32(len(svrHandCard))
		gamePlayerInfo := &room.GamePlayerInfo{
			PlayerInfo:    roomPlayerInfo,
			Color:         gutils.ServerColor2ClientColor(player.DingqueColor).Enum(),
			HandCardCount: &handCardCount,
		}

		// 手牌组
		cltHandCard := gutils.ServerCards2Numbers(svrHandCard)
		handCardGroup := &room.CardsGroup{
			Cards: cltHandCard,
			Type:  room.CardsGroupType_CGT_HAND.Enum(),
		}
		gamePlayerInfo.CardsGroup = append(gamePlayerInfo.CardsGroup, handCardGroup)
		// 吃牌组

		// 碰牌组,每一次碰牌填1张还是三张
		var pengCardGroups []*room.CardsGroup
		for _, pengCard := range player.GetPengCards() {
			srcPlayerID := pengCard.GetSrcPlayer()
			pengCardGroup := &room.CardsGroup{
				Cards: []uint32{gutils.ServerCard2Number(pengCard.GetCard())},
				Type:  room.CardsGroupType_CGT_PENG.Enum(),
				Pid:   &srcPlayerID,
			}
			pengCardGroups = append(pengCardGroups, pengCardGroup)
		}
		gamePlayerInfo.CardsGroup = append(gamePlayerInfo.CardsGroup, pengCardGroups...)
		// 杠牌组
		var gangCardGroups []*room.CardsGroup
		for _, gangCard := range player.GetGangCards() {
			groupType := gutils.GangTypeSvr2Client(gangCard.GetType())
			srcPlayerID := gangCard.GetSrcPlayer()
			gangCardGroup := &room.CardsGroup{
				Cards: []uint32{gutils.ServerCard2Number(gangCard.GetCard())},
				Type:  &groupType,
				Pid:   &srcPlayerID,
			}
			gangCardGroups = append(gangCardGroups, gangCardGroup)
		}
		gamePlayerInfo.CardsGroup = append(gamePlayerInfo.CardsGroup, gangCardGroups...)
		// 胡牌组
		var huCardGroups []*room.CardsGroup
		for _, huCard := range player.GetHuCards() {
			srcPlayerID := huCard.GetSrcPlayer()
			huCardGroup := &room.CardsGroup{
				Cards: []uint32{gutils.ServerCard2Number(huCard.GetCard())},
				Type:  room.CardsGroupType_CGT_HU.Enum(),
				Pid:   &srcPlayerID,
			}
			huCardGroups = append(huCardGroups, huCardGroup)
		}
		gamePlayerInfo.CardsGroup = append(gamePlayerInfo.CardsGroup, huCardGroups...)
		// 花牌组

		// 出牌组
		outCardGroup := &room.CardsGroup{
			Cards: gutils.ServerCards2Numbers(player.GetOutCards()),
			Type:  room.CardsGroupType_CGT_OUT.Enum(),
			Pid:   &playerID,
		}
		gamePlayerInfo.CardsGroup = append(gamePlayerInfo.CardsGroup, outCardGroup)
		recoverPlayerInfo = append(recoverPlayerInfo, gamePlayerInfo)
	}
	return
}

func getZixunInfo(playerID uint64, mjContext *server_pb.MajongContext) *room.RoomZixunNtf {
	if mjContext.GetCurState() != server_pb.StateID_state_zixun {
		return nil
	}

	if mjContext.GetLastMopaiPlayer() != playerID {
		return nil
	}
	player := gutils.GetMajongPlayer(playerID, mjContext)
	return zixunTransform(player.GetZixunRecord())
}

func getWenxunInfo(playerID uint64, mjContext *server_pb.MajongContext) *room.RoomChupaiWenxunNtf {
	if mjContext.GetCurState() != server_pb.StateID_state_chupaiwenxun {
		return nil
	}

	player := gutils.GetMajongPlayer(playerID, mjContext)
	enableActions := player.GetPossibleActions()
	if len(enableActions) == 0 || player.GetHasSelected() {
		return nil
	}

	outCard := gutils.ServerCard2Number(mjContext.GetLastOutCard())
	wenXunInfo := &room.RoomChupaiWenxunNtf{
		Card: &outCard,
	}
	for _, action := range enableActions {
		switch action {
		case server_pb.Action_action_peng:
			wenXunInfo.EnablePeng = proto.Bool(true)
		case server_pb.Action_action_gang:
			wenXunInfo.EnableMinggang = proto.Bool(true)
		case server_pb.Action_action_hu:
			wenXunInfo.EnableDianpao = proto.Bool(true)
		case server_pb.Action_action_qi:
			wenXunInfo.EnableQi = proto.Bool(true)
		}
	}
	return wenXunInfo
}

func getQghInfo(playerID uint64, mjContext *server_pb.MajongContext) *room.RoomWaitQianggangHuNtf {
	if mjContext.GetCurState() != server_pb.StateID_state_waitqiangganghu {
		return nil
	}

	player := gutils.GetMajongPlayer(playerID, mjContext)
	enableActions := player.GetPossibleActions()
	if len(enableActions) == 0 || player.GetHasSelected() {
		return nil
	}

	outCard := gutils.ServerCard2Number(mjContext.GetLastOutCard())
	gangPlayerID := mjContext.GetLastGangPlayer()
	qghInfo := &room.RoomWaitQianggangHuNtf{
		Card:         &outCard,
		SelfCan:      proto.Bool(len(player.GetPossibleActions()) != 0),
		FromPlayerId: &gangPlayerID,
	}
	return qghInfo
}

func zixunTransform(record *server_pb.ZixunRecord) *room.RoomZixunNtf {
	zixunNtf := &room.RoomZixunNtf{}
	zixunNtf.EnableAngangCards = record.GetEnableAngangCards()
	zixunNtf.EnableBugangCards = record.GetEnableBugangCards()
	zixunNtf.EnableChupaiCards = record.GetEnableChupaiCards()
	zixunNtf.EnableQi = proto.Bool(record.GetEnableQi())
	zixunNtf.EnableZimo = proto.Bool(record.GetEnableZimo())
	huType := gutils.HuTypeSvr2Client(record.GetHuType())
	if huType != nil {
		zixunNtf.HuType = huType
	}
	zixunNtf.CanTingCardInfo = gutils.CanTingCardInfoSvr2Client(record.GetCanTingCardInfo())

	return zixunNtf
}
