package ai

import (
	"steve/entity/majong"
	"steve/entity/poker/ddz"
)

// AIType AI 类型
type AIType int

func (ai AIType) String() string {
	switch ai {
	case OverTimeAI:
		return "OverTime"
	case TuoGuangAI:
		return "TuoGuan"
	case RobotAI:
		return "Robot"
	default:
		return "Unknow AI"
	}
}

const (
	// OverTimeAI 超时 AI
	OverTimeAI AIType = iota
	// TuoGuangAI 托管 AI
	TuoGuangAI
	// RobotAI 机器人 AI
	RobotAI
)

// AIParams 生成 AI 事件需要的参数
type AIParams struct {
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

// AIResult AI 事件生成结果
type AIResult struct {
	Events []AIEvent
}

type CommonAI interface {
	GenerateAIEvent(params AIParams) (AIResult, error)
}
