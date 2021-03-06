package common

import (
	"steve/room/majong/interfaces"
	"testing"

	majongpb "steve/entity/majong"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

// Test_InitState_ProcessEventStartGame 测试初始状态接收游戏开始事件
func Test_InitState_ProcessEventStartGame(t *testing.T) {
	mc := gomock.NewController(t)
	defer mc.Finish()

	flow := interfaces.NewMockMajongFlow(mc)
	flow.EXPECT().PushMessages(gomock.Any(), gomock.Any()).AnyTimes()
	flow.EXPECT().GetMajongContext().Return(&majongpb.MajongContext{}).AnyTimes()

	initState := new(InitState)
	newStateID, err := initState.ProcessEvent(majongpb.EventID_event_start_game, nil, flow)

	assert.Nil(t, err)
	assert.Equal(t, majongpb.StateID_state_xipai, newStateID)
}
