package scxlai

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/spf13/viper"
	"math/rand"
	"sort"
	"steve/entity/majong"
	"steve/gutils"
	"steve/room/ai"
)

type chupaiWenxunStateAI struct {
}

// GenerateAIEvent 生成 出牌问询AI 事件
// 无论是超时、托管还是机器人，胡过了自动胡，没胡过的其他操作都默认弃， 并且产生相应的事件
func (h *chupaiWenxunStateAI) GenerateAIEvent(params ai.AIParams) (result ai.AIResult, err error) {
	result, err = ai.AIResult{
		Events: []ai.AIEvent{},
	}, nil

	mjContext := params.MajongContext
	player := gutils.GetMajongPlayer(params.PlayerID, mjContext)
	if h.checkAIEvent(player, mjContext, params) != nil {
		return
	}
	// if len(player.GetPossibleActions()) == 0 {
	// 	return
	// }
	switch params.AIType {
	case ai.RobotAI:
		{
			event := h.askMiddleAI(player, *mjContext.LastOutCard)
			result.Events = append(result.Events, event)
		}
	case ai.OverTimeAI, ai.TuoGuangAI:
		{
			if gutils.IsTing(player) {
				return
			}
			if gutils.IsHu(player) && h.containAction(player, majong.Action_action_gang) {
				return
			}
			if viper.GetBool("ai.test") {
				event := h.askMiddleAI(player, *mjContext.LastOutCard)
				result.Events = append(result.Events, event)
			} else {
				event := h.chupaiWenxun(player)
				result.Events = append(result.Events, event)
			}
		}
	}

	return
}

func (h *chupaiWenxunStateAI) askMiddleAI(player *majong.Player, lastOutCard majong.Card) ai.AIEvent {
	logEntry := logrus.WithField("playerId", player.PlayerId)
	var (
		event ai.AIEvent
	)
	actions := player.GetPossibleActions()
	sort.Sort(sort.Reverse(majong.ActionSlice(actions))) //按优先级从高到低排列

	for _, action := range actions {
		switch action {
		case majong.Action_action_hu:
			logEntry.WithField("点炮牌", lastOutCard).Infoln("中级AI点炮胡牌")
			return hu(player)
		case majong.Action_action_gang:
			s := SplitCards(NonPointer(player.HandCards))
			if len(s.KeZis) > 0 && Contains(s.KeZis, lastOutCard) {
				logEntry.WithField("明杠牌", lastOutCard).Infoln("中级AI明杠")
				return gang(player, &lastOutCard)
			}
		case majong.Action_action_peng:
			s := SplitCards(NonPointer(player.HandCards))
			if len(s.Pairs) > 0 && Contains(s.Pairs, lastOutCard) {
				r := rand.Intn(100)
				if len(s.Pairs) >= 2 && r < 90 || len(s.Pairs) == 1 && r < 10 { //多于1对时，碰牌概率90%；等于1对时，碰牌概率10%
					logEntry.WithField("碰牌", lastOutCard).Infoln("中级AI碰牌")
					return peng(player)
				}
			}
		case majong.Action_action_chi:
			s := SplitCards(NonPointer(player.HandCards))
			notOKCards := append(s.GetNotOKCards(), lastOutCard)
			notOKSplits := SplitCards(notOKCards)
			if len(notOKSplits.ShunZis) >= 1 && len(notOKSplits.Pairs) >= 1 {
				shunZi := notOKSplits.ShunZis[0].cards
				logEntry.WithField("吃牌", lastOutCard).Infoln("中级AI吃牌")
				return chi(player, []*majong.Card{&shunZi[0], &shunZi[1], &shunZi[2]})
			}
			// 老的吃策略
			//if len(s.SingleChas)+len(s.DoubleChas) > 0 {
			//	for _, cha := range append(s.SingleChas, s.DoubleChas...) { //优先处理单茬
			//		validCards := getValidCard(cha)
			//		if ContainsCard(validCards, lastOutCard) {
			//			logEntry.WithField("吃牌", lastOutCard).Infoln("中级AI吃牌")
			//			return chi(player, []*majong.Card{&cha.cards[0], &cha.cards[1], &lastOutCard})
			//		}
			//	}
			//}
		default:
			event = qi(player)
		}
	}

	return event
}

func (h *chupaiWenxunStateAI) containAction(player *majong.Player, action majong.Action) bool {
	for _, possibleAction := range player.GetPossibleActions() {
		if possibleAction == action {
			return true
		}
	}
	return false
}

// chupaiWenxun 生成出牌问询请求事件
func (h *chupaiWenxunStateAI) chupaiWenxun(player *majong.Player) ai.AIEvent {
	if (gutils.IsTing(player) || gutils.IsHu(player)) && h.containAction(player, majong.Action_action_hu) {
		return hu(player)
	}
	return qi(player)
}

func (h *chupaiWenxunStateAI) checkAIEvent(player *majong.Player, mjContext *majong.MajongContext, params ai.AIParams) error {
	err := fmt.Errorf("不生成自动事件")
	if player.GetHasSelected() {
		return err
	}
	if len(player.GetPossibleActions()) == 0 {
		return err
	}
	if mjContext.GetCurState() != majong.StateID_state_chupaiwenxun {
		return err
	}

	return nil
}
