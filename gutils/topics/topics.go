package topics

const (
	// ClientDisconnect 客户端断开连接
	// 目前订阅者：
	//	room@[host]
	ClientDisconnect string = "client_disconnect"

	// GameDetailRecord 游戏记录明细
	// 目前订阅者：
	//
	GameDetailRecord string = "game_detail_record"

	// GameSummaryRecord 游戏记录明细
	// 目前订阅者：
	//
	GameSummaryRecord string = "game_summary_record"

	// BroadcastMsg 通过GateWay广播消息给客户端
	// 目前订阅者：
	//
	BroadcastMsg string = "broadcast_msg"

	// GoldChangeNtf 玩家金币数变化
	// 目前订阅者：
	//
	GoldChangeNtf string = "gold_change"
)
