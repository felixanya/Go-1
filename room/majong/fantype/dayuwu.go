package fantype

//checkDaYuWu 检查大于五 只能有序数牌并且序数牌>5
func checkDaYuWu(tc *typeCalculator) bool {
	// 所有牌
	cardAll := getPlayerCardAll(tc)
	for _, card := range cardAll {
		// 只能有序数牌并且序数牌>5
		if !IsXuShuCard(card) || card.GetPoint() <= 5 {
			return false
		}
	}
	return true
}
