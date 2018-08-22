package ddz

import (
	"fmt"
	"steve/entity/poker/ddz"
	"steve/room/ai"
	. "steve/room/flows/ddzflow/ddz/states"

	"steve/entity/poker"

	"github.com/Sirupsen/logrus"
)

type playStateAI struct {
}

// GenerateAIEvent 生成 出牌AI 事件
// 无论是超时、托管还是机器人，胡过了自动胡，没胡过的其他操作都默认弃， 并且产生相应的事件
func (playAI *playStateAI) GenerateAIEvent(params ai.AIEventGenerateParams) (result ai.AIEventGenerateResult, err error) {
	logEntry := logrus.WithFields(logrus.Fields{
		"func_name": "play.go:GenerateAIEvent()"})

	// 产生的事件结果
	result, err = ai.AIEventGenerateResult{
		Events: []ai.AIEvent{},
	}, nil

	ddzContext := params.DDZContext

	// 没到自己打牌
	if ddzContext.GetCurrentPlayerId() != params.PlayerID {
		return result, nil
	}

	// 当前玩家
	var curPlayer *ddz.Player
	for _, player := range ddzContext.GetPlayers() {
		if player.GetPlayerId() == params.PlayerID {
			curPlayer = player
		}
	}

	// 无效玩家
	if curPlayer == nil {
		logEntry.Errorf("无效玩家%d", params.PlayerID)
		return result, fmt.Errorf("无效玩家%d", params.PlayerID)
	}

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
	handCards := ToDDZCards(player.GetHandCards())
	// 上家出的牌，转换为poke
	outCards := ToDDZCards(ddzContext.GetCurOutCards())

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
	logrus.Info("被动出牌：%s", biggerCards)
	return play(player, ToInts(biggerCards))
}

// Play 生成出牌请求事件(主动出牌)
func (playAI *playStateAI) getActivePlayCardEvent(ddzContext *ddz.DDZContext, player *ddz.Player) ai.AIEvent {
	// 转换为poke
	handCards := ToDDZCards(player.GetHandCards())
	// 按照排序权重进行排序
	DDZPokerSort(handCards)
	logrus.Info("主动出牌：%v", handCards[0])
	return play(player, ToInts([]Poker{handCards[0]}))
}

func play(player *ddz.Player, cards []uint32) ai.AIEvent {
	request := &ddz.PlayCardRequestEvent{
		Head: &ddz.RequestEventHead{
			PlayerId: player.GetPlayerId()},
		Cards: cards,
	}
	event := ai.AIEvent{
		ID:      int32(ddz.EventID_event_chupai_request),
		Context: request,
	}
	return event
}
