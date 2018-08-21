package ai

import (
	"steve/entity/majong"
	"steve/entity/poker/ddz"
)

// AIType AI 类型
type AIType int

const (
	// OverTimeAI 超时 AI
	OverTimeAI AIType = iota
	// TuoGuangAI 托管 AI
	TuoGuangAI
	// RobotAI 机器人 AI
	RobotAI
)

// PlayerAIInfo 玩家 AI 信息
type PlayerAIInfo struct {
	AIType  AIType // AI 类型
	RobotLv int    // 机器人级别
}

// AIEventGenerateParams 生成 AI 事件需要的参数
type AIEventGenerateParams struct {
	MajongContext *majong.MajongContext
	DDZContext    *ddz.DDZContext
	PlayerID      uint64
	AIType        AIType
	RobotLv       int
}

// AIEvent AI 事件
type AIEvent struct {
	ID      int32
	Context interface{}
}

// AIEventGenerateResult AI 事件生成结果
type AIEventGenerateResult struct {
	Events []AIEvent
}

type CommonAI interface {
	GenerateAIEvent(params AIEventGenerateParams) (AIEventGenerateResult, error)
}
