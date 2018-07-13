package gutils

import (
	"steve/client_pb/room"
	majongpb "steve/server_pb/majong"
)

//

// IsTing 玩家是否是听的状态
func IsTing(player *majongpb.Player) bool {
	tingState := player.GetTingStateInfo()
	if tingState.GetIsTing() || tingState.GetIsTianting() {
		return true
	}
	return false
}

// GetTingType 获取玩家听的类型
func GetTingType(player *majongpb.Player) (tingType room.TingType) {
	tingState := player.GetTingStateInfo()
	if tingState.GetIsTing() {
		tingType = room.TingType_TT_NORMAL_TING
	}
	if tingState.GetIsTianting() {
		tingType = room.TingType_TT_TIAN_TING
	}
	return tingType
}

// IsHu 玩家是否时胡的状态
func IsHu(player *majongpb.Player) bool {
	if len(player.GetHuCards()) > 0 {
		return true
	}
	return false
}
