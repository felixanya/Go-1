package fantype

// calcGengCount 计算花的数量
func (tc *typeCalculator) calcHuaCount() int {
	return len(tc.getPlayer().GetHuaCards())
}
