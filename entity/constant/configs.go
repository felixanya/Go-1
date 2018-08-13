package constant

// 配置 key 和 subkey 的定义

// ConfigKey 配置键值
type ConfigKey struct {
	Key, SubKey string
}

// PropKey 道具配置主键
var PropKey = "prop"

// PropSubKey 道具配置子健 互动道具，表情道具，非便捷道具
var PropSubKey = "interactive"

// PropSubKeySpeedy 道具配置子健 便捷道具（比赛卷。。。）
var PropSubKeySpeedy = "speedy"

var (
	// ChargeItemListKey 充值系统商品列表配置
	ChargeItemListKey = ConfigKey{Key: "charge", SubKey: "item_list"}
	// ChargeDayMaxKey 每日充值上限配置
	ChargeDayMaxKey = ConfigKey{Key: "charge", SubKey: "day_max"}
	// PropInteractiveKey 互动道具PropSubKey 道具配置子健
	PropInteractiveKey = ConfigKey{Key: PropKey, SubKey: PropSubKey}
	// PropSpeedyKey 便捷道具PropSubKeySpeedy 道具配置子健
	PropSpeedyKey = ConfigKey{Key: PropKey, SubKey: PropSubKeySpeedy}
)

// PropAttr 道具属性
type PropAttr struct {
	PropID   int32  `json:"propID"`    // 道具ID
	PropName string `json:"name"`      // 道具名称
	Type     int32  `json:"attrType"`  // 属性类型：货币         | 道具
	TypeID   int32  `json:"attrID"`    // 属性ID:金币、钻石、房卡 | 道具ID
	Value    int64  `json:"attrValue"` // 属性值：操作数量
	Limit    int64  `json:"attrLimit"` // 使用限制
}
