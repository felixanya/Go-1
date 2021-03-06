package scxlai

import (
	"steve/entity/majong"
	"steve/gutils"
	"time"

	"steve/room/ai"
)

type dingqueStateAI struct {
	maxDingqueTime time.Duration // 最大定缺时间
}

// GenerateAIEvent 生成 AI 事件
// 无论是超时、托管还是机器人，都选最少的牌作为定缺牌， 并且产生相应的事件
func (h *dingqueStateAI) GenerateAIEvent(params ai.AIParams) (result ai.AIResult, err error) {
	result, err = ai.AIResult{
		Events: []ai.AIEvent{},
	}, nil

	mjContext := params.MajongContext
	player := gutils.GetMajongPlayer(params.PlayerID, mjContext)
	if player.GetHasDingque() {
		return
	}
	if event := h.dingque(player); event != nil {
		result.Events = append(result.Events, *event)
	}
	return
}

// allColor 所有的麻将花色
func (h *dingqueStateAI) allColor() []majong.CardColor {
	return []majong.CardColor{majong.CardColor_ColorWan, majong.CardColor_ColorTiao, majong.CardColor_ColorTong}
}

// getColor 获取定缺花色
func (h *dingqueStateAI) getColor(player *majong.Player) majong.CardColor {
	return player.GetDingqueColor() // 在进入定缺状态时，会设置推荐定缺颜色
}

// dingque 生成定缺请求事件
func (h *dingqueStateAI) dingque(player *majong.Player) *ai.AIEvent {
	color := h.getColor(player)
	eventContext := &majong.DingqueRequestEvent{
		Head: &majong.RequestEventHead{
			PlayerId: player.GetPlayerId(),
		},
		Color: color,
	}

	return &ai.AIEvent{
		ID:      int32(majong.EventID_event_dingque_request),
		Context: eventContext,
	}
}
