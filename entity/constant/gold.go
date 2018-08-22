package constant

import "steve/client_pb/common"

// 货币类型
const (
	GOLD_COIN  = int(common.MoneyType_MT_COIN)      // 金币
	GOLD_INGOT = int(common.MoneyType_MT_GOLDINGOT) // 元宝
	GOLD_CARD  = int(common.MoneyType_MT_CARD)      // 房卡
)
