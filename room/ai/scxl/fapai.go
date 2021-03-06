package scxlai

import (
	"steve/client_pb/room"
	"steve/entity/majong"
	"steve/gutils"
	"steve/room/ai"

	"github.com/Sirupsen/logrus"
)

type fapaiStateAI struct {
}

// GenerateAIEvent 生成 发牌AI 事件
// 无论是超时、托管还是机器人，发牌发牌动画完成产生相应的事件
func (f *fapaiStateAI) GenerateAIEvent(params ai.AIParams) (result ai.AIResult, err error) {
	result, err = ai.AIResult{
		Events: []ai.AIEvent{},
	}, nil
	if params.AIType != ai.RobotAI { //不是机器人不能发送动画完成请求
		return
	}
	mjContext := params.MajongContext
	crPlayerIDs := mjContext.GetTempData().GetCartoonReqPlayerIDs()
	if len(crPlayerIDs) == len(mjContext.GetPlayers()) { //所有玩家都发送过动画完成请求
		return
	}
	player := gutils.GetMajongPlayer(params.PlayerID, mjContext)
	for _, playerID := range crPlayerIDs {
		if playerID == player.GetPlayerId() { // 当前玩家已经发送过
			return
		}
	}
	// 发送动画完成请求
	if event := CartoonFinsh(player, int32(room.CartoonType_CTNT_FAPAI)); event != nil {
		result.Events = append(result.Events, *event)
	}
	return
}

//CartoonFinsh 动画完成请求事件
func CartoonFinsh(player *majong.Player, cartoonType int32) *ai.AIEvent {
	event := majong.CartoonFinishRequestEvent{
		CartoonType: cartoonType,
		PlayerId:    player.GetPlayerId(),
	}
	logrus.WithFields(logrus.Fields{"func_name": "ai.CartoonFinsh", "player_id": player.GetPlayerId(), "cartoonType": cartoonType}).Errorln("机器人发送动画完成请求事件")
	return &ai.AIEvent{
		ID:      int32(majong.EventID_event_cartoon_finish_request),
		Context: &event,
	}
}
