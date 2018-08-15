package db

import (
	"time"
)

type TPlayerPacksack struct {
	Id         int64     `xorm:"pk autoincr comment('ID') BIGINT(20)"`
	Playerid   int64     `xorm:"not null comment('玩家ID') BIGINT(20)"`
	Gold       int       `xorm:"default 0 comment('背包金币数') INT(11)"`
	Createtime time.Time `xorm:"DATETIME"`
	Createby   string    `xorm:"VARCHAR(64)"`
	Updatetime time.Time `xorm:"DATETIME"`
	Updateby   string    `xorm:"VARCHAR(64)"`
}
