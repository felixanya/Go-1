package packsack

const (
	defaultRestrict     = 30000 // 存入和存出限制
	defaultProcedureFee = 0.05  // 手续费
)

//GoldPacksackInfo 背包金币信息
type GoldPacksackInfo struct {
	PkGlod       int64   // 背包金币
	Restrict     int64   // 存入和存出限制
	ProcedureFee float64 //手续费
}

//GetPacksackGoldInfo 获取背包的金币信息
func GetPacksackGoldInfo(playerID uint64) (*GoldPacksackInfo, error) {
	gpi := &GoldPacksackInfo{}
	return gpi, nil
}
