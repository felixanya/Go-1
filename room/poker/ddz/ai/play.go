package ddz

import (
	"errors"
	"steve/entity/poker/ddz"
	"steve/room/ai"
	. "steve/room/poker"
	. "steve/room/poker/ddz/states"

	"steve/entity/poker"

	"github.com/Sirupsen/logrus"
)

type playStateAI struct {
}

// GenerateAIEvent 生成 出牌AI 事件
// 无论是超时、托管还是机器人，胡过了自动胡，没胡过的其他操作都默认弃， 并且产生相应的事件
func (playAI *playStateAI) GenerateAIEvent(params ai.AIParams) (result ai.AIResult, err error) {
	ddzContext := params.DDZContext
	// 没到自己打牌
	if ddzContext.GetCurrentPlayerId() != params.PlayerID {
		return
	}

	if !IsValidPlayer(ddzContext, params.PlayerID) {
		logrus.WithField("players", ddzContext.GetPlayers()).WithField("playerId", params.PlayerID).Errorln("斗地主无效玩家")
		return result, errors.New("无效玩家")
	}

	// 当前玩家
	curPlayer := GetPlayerByID(ddzContext.GetPlayers(), params.PlayerID)

	// 没有牌型时说明是主动打牌
	if ddzContext.GetCurCardType() == poker.CardType_CT_NONE {
		event := playAI.getActivePlayCardEvent(ddzContext, curPlayer) //主动出牌，出最小的牌
		result.Events = append(result.Events, event)
	} else {
		if params.AIType == ai.RobotAI {
			event := playAI.getPassivePlayCardEvent(ddzContext, curPlayer) //能压就压
			result.Events = append(result.Events, event)
		} else {
			result.Events = append(result.Events, play(curPlayer, nil)) //超时和托管不出
		}
	}
	return
}

// Play 生成出牌请求事件(被动出牌)
func (playAI *playStateAI) getPassivePlayCardEvent(ddzContext *ddz.DDZContext, player *ddz.Player) ai.AIEvent {
	// 转换为poke
	handCards := ToPokers(player.GetHandCards())
	// 上家出的牌，转换为poke
	outCards := ToPokers(ddzContext.GetCurOutCards())

	bigger, biggerCards := GetMinBiggerCards(handCards, outCards)
	if !bigger {
		curCardType := ddzContext.GetCurCardType()
		// 无压制的牌，且当前牌型是炸弹，则判断自己有无火箭
		if curCardType == poker.CardType_CT_BOMB {
			if hasKingBomb, kingBomb := GetKingBomb(handCards); hasKingBomb {
				biggerCards = kingBomb
			}
		}

		// 无压制的牌，且当前牌型不是炸弹，也不是火箭，则判断自己有无炸弹，无炸弹时再检测火箭
		if curCardType != poker.CardType_CT_BOMB && curCardType != poker.CardType_CT_KINGBOMB {
			// 优先检测炸弹
			if hasBomb, bomb := GetBomb(handCards); hasBomb {
				biggerCards = bomb
			} else {
				if hasKingBomb, kingBomb := GetKingBomb(handCards); hasKingBomb {
					biggerCards = kingBomb
				}
			}
		}
	}
	logrus.WithField("handCards", handCards).WithField("biggerCards", biggerCards).Debugln("被动出牌")
	return play(player, biggerCards)
}

// Play 生成出牌请求事件(主动出牌)
func (playAI *playStateAI) getActivePlayCardEvent(ddzContext *ddz.DDZContext, player *ddz.Player) ai.AIEvent {
	// 转换为poke
	handCards := ToPokers(player.GetHandCards())
	// 按照排序权重进行排序
	PokerSort(handCards)
	outCard := handCards[0]
	logrus.WithField("handCards", handCards).WithField("outCard", outCard).Debugln("主动出牌")
	return play(player, []Poker{outCard})
}

func play(player *ddz.Player, cards []Poker) ai.AIEvent {
	request := &ddz.PlayCardRequestEvent{
		Head: &ddz.RequestEventHead{
			PlayerId: player.GetPlayerId()},
		Cards: ToInts(cards),
	}
	event := ai.AIEvent{
		ID:      int32(ddz.EventID_event_chupai_request),
		Context: request,
	}
	return event
}
