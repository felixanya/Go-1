package poker

// ContainsAll handCards是否包含所有outCards
func ContainsAll(handCards []Poker, outCards []Poker) bool {
	for _, outCard := range outCards {
		if !Contains(handCards, outCard) {
			return false
		}
	}
	return true
}

// Contains cards是否包含card
func Contains(cards []Poker, card Poker) bool {
	for _, value := range cards {
		if value.Equals(card) {
			return true
		}
	}
	return false
}

// ContainsPoint cards是否包含点数
func ContainsPoint(cards []Poker, point uint32) bool {
	for _, card := range cards {
		if card.Point == point {
			return true
		}
	}
	return false
}

// RemoveByPoint 从cards中删除所有与removeCards相同点数的牌
func RemoveByPoint(cards []Poker, removeCards []Poker) []Poker {
	var result []Poker
	for _, card := range cards {
		if !ContainsPoint(removeCards, card.PointWeight) {
			result = append(result, card)
		}
	}
	return result
}

// RemoveAll 从cards中删除removeCards
func RemoveAll(cards []Poker, removeCards []Poker) []Poker {
	var result []Poker
	for _, card := range cards {
		if !Contains(removeCards, card) {
			result = append(result, card)
		}
	}
	return result
}
