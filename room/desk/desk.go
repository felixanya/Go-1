package desk

type Desk struct {
	uid       uint64
	gameID    int
	config    *DeskConfig
	playerIds []uint64
	level     int32
	// Context   context.Context
	// Cancel    context.CancelFunc // 取消事件处理
}

func NewDesk(uid uint64, gameId int,levelId int32, playerIds []uint64, config *DeskConfig) Desk {
	desk := Desk{uid: uid,
		gameID:    gameId,
		config:    config,
		playerIds: playerIds,
		level:levelId,
	}

	return desk
}

func (desk *Desk) GetPlayerIds() []uint64 {
	return desk.playerIds
}

func (desk *Desk) GetUid() uint64 {
	return desk.uid
}

func (desk *Desk) GetGameId() int {
	return desk.gameID
}

func (desk *Desk) GetConfig() *DeskConfig {
	return desk.config
}

func (desk *Desk) GetLevel() int32 {
	return desk.level
}
