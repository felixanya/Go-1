package scxlai

import (
	"steve/common/mjoption"
	"steve/entity/majong"
	"steve/gutils"
	"steve/room/ai"
)

func (h *zixunStateAI) getNormalZiXunAIEvent(player *majong.Player, mjContext *majong.MajongContext) (aiEvent ai.AIEvent) {
	zxRecord := player.GetZixunRecord()
	handCards := player.GetHandCards()
	canHu := zxRecord.GetEnableZimo()
	if (gutils.IsHu(player) || gutils.IsTing(player)) && canHu {
		aiEvent = hu(player)
		return
	}
	// 优先出定缺牌
	if gutils.CheckHasDingQueCard(mjContext, player) {
		for i := len(handCards) - 1; i >= 0; i-- {
			hc := handCards[i]
			if mjoption.GetXingpaiOption(int(mjContext.GetXingpaiOptionId())).EnableDingque &&
				hc.GetColor() == player.GetDingqueColor() {
				aiEvent = chupai(player, hc)
				return
			}
		}
	}

	// 正常出牌
	if player.GetMopaiCount() == 0 || mjContext.GetZixunType() == majong.ZixunType_ZXT_CHI || mjContext.GetZixunType() == majong.ZixunType_ZXT_PENG {
		aiEvent = chupai(player, handCards[len(handCards)-1])
	} else {
		aiEvent = chupai(player, mjContext.GetLastMopaiCard())
	}
	return
}

func chupai(player *majong.Player, card *majong.Card) ai.AIEvent {
	eventContext := majong.ChupaiRequestEvent{
		Head: &majong.RequestEventHead{
			PlayerId: player.GetPlayerId(),
		},
		Cards: card,
	}
	return ai.AIEvent{
		ID:      int32(majong.EventID_event_chupai_request),
		Context: &eventContext,
	}
}

func chi(player *majong.Player, chiCards []*majong.Card) ai.AIEvent {
	eventContext := majong.ChiRequestEvent{
		Head: &majong.RequestEventHead{
			PlayerId: player.GetPlayerId(),
		},
		Cards: chiCards,
	}
	return ai.AIEvent{
		ID:      int32(majong.EventID_event_chi_request),
		Context: &eventContext,
	}
}

func peng(player *majong.Player) ai.AIEvent {
	eventContext := majong.PengRequestEvent{
		Head: &majong.RequestEventHead{
			PlayerId: player.GetPlayerId(),
		},
	}
	return ai.AIEvent{
		ID:      int32(majong.EventID_event_peng_request),
		Context: &eventContext,
	}
}

func gang(player *majong.Player, card *majong.Card) ai.AIEvent {
	eventContext := majong.GangRequestEvent{
		Head: &majong.RequestEventHead{
			PlayerId: player.GetPlayerId(),
		},
		Card: card,
	}
	return ai.AIEvent{
		ID:      int32(majong.EventID_event_gang_request),
		Context: &eventContext,
	}
}

func hu(player *majong.Player) ai.AIEvent {
	eventContext := majong.HuRequestEvent{
		Head: &majong.RequestEventHead{
			PlayerId: player.GetPlayerId(),
		},
	}
	return ai.AIEvent{
		ID:      int32(majong.EventID_event_hu_request),
		Context: &eventContext,
	}
}

func qi(player *majong.Player) ai.AIEvent {
	eventContext := majong.QiRequestEvent{
		Head: &majong.RequestEventHead{
			PlayerId: player.GetPlayerId(),
		},
	}
	return ai.AIEvent{
		ID:      int32(majong.EventID_event_qi_request),
		Context: &eventContext,
	}
}
