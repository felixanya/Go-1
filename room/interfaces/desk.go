package interfaces

import (
	"steve/client_pb/room"
	"steve/structs/proto/gate_rpc"

	"github.com/golang/protobuf/proto"
)

// Desk 牌桌
type Desk interface {
	// GetUID 获取牌桌 UID
	GetUID() uint64

	// GetGameID 获取游戏 ID
	GetGameID() int

	// GetPlayers 获取牌桌玩家数据
	GetPlayers() []*room.RoomPlayerInfo

	// Start 启动牌桌逻辑
	// finish : 当牌桌逻辑完成时调用
	Start(finish func()) error

	// PushRequest 压入玩家请求
	PushRequest(playerID uint64, head *steve_proto_gaterpc.Header, body proto.Message)
}

// DeskMgr 牌桌管理器
type DeskMgr interface {
	// RunDesk 运转牌桌
	RunDesk(desk Desk) error

	// HandlePlayerRequest 处理玩家请求
	HandlePlayerRequest(playerID uint64, head *steve_proto_gaterpc.Header, body proto.Message)
}

// CreateDeskOptions 创建牌桌选项
type CreateDeskOptions struct{}

// CreateDeskResult 创建房间结果
type CreateDeskResult struct {
	Desk Desk
}

// DeskFactory 牌桌工厂
type DeskFactory interface {
	// CreateDesk 创建牌桌
	CreateDesk(players []uint64, gameID int, opt CreateDeskOptions) (CreateDeskResult, error)
}
