package cache

import "fmt"

// key formats
const (
	// AccountPlayerKey 账号关联的玩家
	AccountPlayerKey = "account:player:%v"

	// playerTokenKeyFmt
	playerTokenKeyFmt = "playertoken:%d"

	playerChargeKeyFmt = "playercharge:%d"

	// playerWxInfoUpdateTimeKeyFmt 玩家上次更新微信信息的时间
	playerWxInfoUpdateTimeKeyFmt = "playerwxupdatetime:%d"
)

// Player 字段
const (
	// AccountID 账号 ID
	AccountID = "accountID"
	// NickName ...昵称
	NickName = "nickname"
	// Avatar ...头像
	Avatar = "avatar"
	// Gender  ...性别
	Gender = "gender"
	// ShowUID ...显示ID
	ShowUID = "showUID"
	// ChannelID ...渠道ID
	ChannelID = "channelID"
	// ProvinceID ...省份ID
	ProvinceID = "provinceID"
	// CityID ...城市ID
	CityID = "cityID"
	// Name ...名称
	Name = "name"
	// IDCard ... 身份证
	IDCard = "idCard"

	// Phone 手机号
	Phone = "phone"

	// GameState ...玩家游戏状态
	GameState = "game_state"
	// GameID ...正在进行的游戏id
	GameID = "game_id"
	// LevelID ...正在进行的游戏场次id
	LevelID = "level_id"
	// IPAddr ... 玩家ip地址
	IPAddr = "ip_addr"
	// GateAddr ...网关服地址
	GateAddr = "gate_addr"
	// MatchAddr ...匹配服地址
	MatchAddr = "match_addr"
	// RoomAddr ...房间服地址
	RoomAddr = "room_addr"

	// WinningRate ... 对应gameID：游戏胜率
	WinningRate = "winningRate"
	// WinningBurea ... 对应gameID：赢的局数
	WinningBurea = "winningBurea"
	// TotalBurea ... 对应gameID：总局数
	TotalBurea = "totalBureau"
	// MaxWinningStream ... 对应gameID：最大连胜
	MaxWinningStream = "maxWinningStream"
	// MaxMultiple ... 对应gameID：最大倍数
	MaxMultiple = "maxMultiple"

	// TodayChargeKey 今日充值数量
	TodayChargeKey = "today_charge_count"
	// TodayChargeTime 最近充值时间
	LastChargeTime = "today_charge_time"
)

// FmtAccountPlayerKey 账号所关联玩家 key
func FmtAccountPlayerKey(accountID uint64) string {
	return fmt.Sprintf(AccountPlayerKey, accountID)
}

// FmtGameInfoConfigKey 游戏信息 key
func FmtGameInfoConfigKey() string {
	return "gameInfoconfig"
}

// FmtPlayerIDKey 玩家ID key
func FmtPlayerIDKey(playerID uint64) string {
	return fmt.Sprintf("player:%v", playerID)
}

// FmtPlayerGameInfoKey 玩家游戏信息
func FmtPlayerGameInfoKey(playerID uint64, gameID uint32) string {
	return fmt.Sprintf("player:%v gameId:%v", playerID, gameID)
}

// FmtPlayerTokenKey format player's token key
func FmtPlayerTokenKey(playerID uint64) string {
	return fmt.Sprintf(playerTokenKeyFmt, playerID)
}

// FmtPlayerPropKey 玩家道具 key
func FmtPlayerPropKey(playerID uint64, propID int32) string {
	return fmt.Sprintf("player:%v prop_%v", playerID, propID)
}

// FmtPlayerChargeKey fomat player's charge key
func FmtPlayerChargeKey(playerID uint64) string {
	return fmt.Sprintf(playerChargeKeyFmt, playerID)
}

func FmtGameReportKey(gameId int, level int) string {
	return fmt.Sprintf("gamereport:%v-%v", gameId, level)
}

// FmtPlayerWxInfoUpdateTimeKey 返回玩家上次同步微信信息的时间
func FmtPlayerWxInfoUpdateTimeKey(playerID uint64) string {
	return fmt.Sprintf(playerWxInfoUpdateTimeKeyFmt, playerID)
}
func FmtGameReportKeyGame() string {
	return fmt.Sprintf("gamereport:")
}

func FmtRedisLockKeyReport() string {
	return "datareportlock"
}
