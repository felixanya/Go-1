package define

/*
 功能： 基础结构和常量定义
 作者： SkyWang
 日期： 2018-7-24

*/

import (
	"fmt"
	"steve/client_pb/common"
)

// 错误定义
var (
	ErrGoldType = fmt.Errorf("gold type error")
	ErrNoUser   = fmt.Errorf("no user")
	ErrLoadDB   = fmt.Errorf("load from db failed")
	ErrSeqNo    = fmt.Errorf("seq is same")
	ErrNoEnough    = fmt.Errorf("gold no enough")
)

// 货币类型
const (
	GOLD_COIN  = int16(common.MoneyType_MT_COIN)      // 金币
	GOLD_INGOT = int16(common.MoneyType_MT_GOLDINGOT) // 元宝
	GOLD_CARD  = int16(common.MoneyType_MT_CARD)      // 房卡
)
