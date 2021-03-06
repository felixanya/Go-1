package gamelog

import "time"

// TGameSummary 游戏记录汇总
type TGameSummary struct {
	Sumaryid      int64           `json:"	Sumaryid      "`
	Deskid        int64           `json:"	Deskid        "`
	Gameid        int             `json:"	Gameid        "`
	Levelid       int             `json:"	Levelid       "`
	Playerids     []uint64        `json:"	Playerids     "`
	Scoreinfo     []int64         `json:" 	Scoreinfo     "`
	Winnerids     []uint64        `json:"	Winnerids     "`
	Roundcurrency []RoundCurrency `json:"	Roundcurrency "`
	Gamestarttime time.Time       `json:"   Gamestarttime  "`
	Gameovertime  time.Time       `json:"   GameoverTime  "`
	Createtime    time.Time       `json:"	Createtime    "`
	Createby      string          `json:"	Createby      "`
	Updatetime    time.Time       `json:"	Updatetime    "`
	Updateby      string          `json:"	Updateby      "`
}

// TGameDetail 游戏明细
type TGameDetail struct {
	Detailid    int64     `json:" Detailid   "`
	Sumaryid    int64     `json:" Sumaryid   "`
	Playerid    uint64    `json:" Playerid   "`
	Deskid      int64     `json:" Deskid     "`
	Gameid      int       `json:" Gameid     "`
	Levelid     int       `json:" Levelid    "`
	Amount      int64     `json:" Amount     "`
	Iswinner    int       `json:" Iswinner   "`
	MaxTimes    uint32    `json:" MaxTimes   "`
	BrokerCount int       `json:" BrokerCount"`
	Createtime  time.Time `json:" Createtime "`
	Createby    string    `json:" Createby   "`
	Updatetime  time.Time `json:" Updatetime "`
	Updateby    string    `json:" Updateby   "`
}

// RoundCurrency 对局金币流水
type RoundCurrency struct {
	Settletype    int32          `json:" Settletype    "`
	Settledetails []SettleDetail `json:" Settledetails "`
}

// SettleDetail 对局金币流水明细
type SettleDetail struct {
	Playerid  uint64 `json:" Playerid  "`
	ChangeVal int64  `json:" ChangeVal "`
}
