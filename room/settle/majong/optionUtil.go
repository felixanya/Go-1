package majong

import (
	"steve/common/mjoption"
	majongpb "steve/server_pb/majong"
)

// GetSettleOption 获取游戏的结算配置
func GetSettleOption(gameID int) *mjoption.SettleOption {
	return mjoption.GetSettleOption(mjoption.GetGameOptions(gameID).SettleOptionID)
}

// GetCardTypeOption 获取游戏的番型配置
func GetCardTypeOption(gameID int) *mjoption.CardTypeOption {
	return mjoption.GetCardTypeOption(mjoption.GetGameOptions(gameID).CardTypeOptionID)
}

// IsGangSettle 是否是杠结算方式
func IsGangSettle(settleType majongpb.SettleType) bool {
	return map[majongpb.SettleType]bool{
		majongpb.SettleType_settle_angang:   true,
		majongpb.SettleType_settle_bugang:   true,
		majongpb.SettleType_settle_minggang: true,
	}[settleType]
}

// IsHuSettle 是否是胡结算方式
func IsHuSettle(settleType majongpb.SettleType) bool {
	return map[majongpb.SettleType]bool{
		majongpb.SettleType_settle_dianpao: true,
		majongpb.SettleType_settle_zimo:    true,
	}[settleType]
}

// IsRoundSettle 是否是单局结算方式
func IsRoundSettle(settleType majongpb.SettleType) bool {
	return map[majongpb.SettleType]bool{
		majongpb.SettleType_settle_yell:      true,
		majongpb.SettleType_settle_flowerpig: true,
		majongpb.SettleType_settle_calldiver: true,
		majongpb.SettleType_settle_taxrebeat: true,
	}[settleType]
}

// CanInstantSettle 能否立即结算
func CanInstantSettle(settleType majongpb.SettleType, settleOption *mjoption.SettleOption) bool {
	if IsGangSettle(settleType) {
		return settleOption.GangInstantSettle
	} else if IsHuSettle(settleType) {
		return settleOption.HuInstantSettle
	}
	return true
}

// CanRoundSettle 玩家是否可以单局结算
func CanRoundSettle(playerID uint64, huQuitPlayers map[uint64]bool, settleOption *mjoption.SettleOption) bool {
	if huQuitPlayers[playerID] {
		return settleOption.HuQuitPlayerSettle.HuQuitPlayerRoundSettle
	}
	return true
}

// NeedBillDetails 是否需要单局结算详情
func NeedBillDetails(gameID int) bool {
	settleOptionID := mjoption.GetGameOptions(gameID).SettleOptionID
	return mjoption.GetSettleOption(settleOptionID).NeedBillDetails
}
