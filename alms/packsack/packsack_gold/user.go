package packsack_gold

type userGold struct {
	uid  uint64 // 玩家ID
	gold int64  // 货币金币
}

// 对指定货币加金币
func (ug *userGold) Get() (int64, error) {
	return ug.gold, nil
}

// 对指定货币加金币
func (ug *userGold) Add(value int64) (int64, error) {
	// 可能需要判断加减金币后，金币值变成负值！
	ug.gold += value
	return ug.gold, nil
}

// userGold
func newuserGold(uid uint64, gold int64) *userGold {
	return &userGold{
		uid:  uid,
		gold: gold,
	}
}
