package constant

// GoldFuncType 金币增减调用函数类型
type GoldFuncType int32

const (
	// GFGAMESETTLE 游戏结算
	GFGAMESETTLE GoldFuncType = 0
	// GFGAMEPEIPAI 游戏配牌
	GFGAMEPEIPAI GoldFuncType = 1
	// ALMSFUNC 救济金
	ALMSFUNC GoldFuncType = 10
	// PACKSACKFUNC 背包
	PACKSACKFUNC GoldFuncType = 11
)

// PropsFuncType 道具增减调用函数类型
type PropsFuncType int32

const (
	// PFGAMEUSE 游戏中使用
	PFGAMEUSE PropsFuncType = 0
)

// http 错误码
const (
	// HTTPOK  成功
	HTTPOK = 0

	// HTTPFailure 失败
	HTTPFailure = 1

	// HTTPInvalidRequest 请求数据错误
	HTTPInvalidRequest = 2
)
