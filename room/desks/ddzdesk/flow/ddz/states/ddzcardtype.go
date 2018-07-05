package states

import (
	"steve/server_pb/ddz"
)

func canBiggerThan(mine ddz.CardType, other ddz.CardType) bool {
	if mine == ddz.CardType_CT_NONE {
		return false
	}
	if mine == ddz.CardType_CT_KINGBOMB {
		return true
	} else if mine == ddz.CardType_CT_BOMB && other != ddz.CardType_CT_KINGBOMB {
		return true
	} else {
		return mine == other
	}
}

func getCardType(cards []Poker) (ddz.CardType, *Poker) {
	if yes, pivot:= isKingBomb(cards); yes {
		return ddz.CardType_CT_KINGBOMB, pivot
	} else if yes, pivot:= isBomb(cards); yes {
		return ddz.CardType_CT_BOMB, pivot
	} else if yes, pivot:= isBombAndPairs(cards); yes {
		return ddz.CardType_CT_4SAND2S, pivot
	} else if yes, pivot:= isBombAndSingles(cards); yes {
		return ddz.CardType_CT_4SAND1S, pivot
	} else if yes, pivot:= isTriples(cards); yes {
		return ddz.CardType_CT_TRIPLES, pivot
	} else if yes, pivot:= isTriplesAndPairs(cards); yes {
		return ddz.CardType_CT_3SAND2S, pivot
	} else if yes, pivot:= isTriplesAndSingles(cards); yes{
		return ddz.CardType_CT_3SAND1S, pivot
	} else if yes, pivot:= isPairs(cards); yes {
		return ddz.CardType_CT_PAIRS, pivot
	} else if yes, pivot:= isShunZi(cards); yes {
		return ddz.CardType_CT_SHUNZI, pivot
	} else if yes, pivot:= isTriple(cards); yes {
		return ddz.CardType_CT_TRIPLE, pivot
	} else if yes, pivot:= isTripleAndPair(cards); yes {
		return ddz.CardType_CT_3AND2, pivot
	} else if yes, pivot:= isTripleAndSingle(cards); yes {
		return ddz.CardType_CT_3AND1, pivot
	} else if yes, pivot:= isPair(cards); yes {
		return ddz.CardType_CT_PAIR, pivot
	} else if yes, pivot:= isSingle(cards); yes {
		return ddz.CardType_CT_SINGLE, pivot
	}

	return ddz.CardType_CT_NONE, nil
}

// 火箭
func isKingBomb(cards []Poker) (bool, *Poker) {
	if len(cards) != 2 {
		return false, nil
	}
	return Contains(cards, blackJoker) && Contains(cards, redJoker), &redJoker
}

// 炸弹
func isBomb(cards []Poker) (bool, *Poker) {
	if len(cards) != 4 {
		return false, nil
	}
	return isAllSamePoint(cards), getMaxCard(cards)
}

// 四带两对
func isBombAndPairs(cards []Poker) (bool, *Poker) {
	if len(cards) != 8 {
		return false, nil
	}

	bomb := getMaxSamePointCards(cards)
	if len(bomb) != 4 {
		return false, nil //没有炸弹
	}

	remain := RemoveAll(cards, bomb)
	firstPair := getMaxSamePointCards(remain)
	if len(firstPair) != 2 {
		return false, nil
	}

	remain = RemoveAll(remain, firstPair)
	if remain[0].Point != remain[1].Point {
		return false, nil
	}

	return true, getMaxCard(bomb)
}

// 四带二单张
func isBombAndSingles(cards []Poker) (bool, *Poker) {
	if len(cards) != 6 {
		return false, nil
	}

	bomb := getMaxSamePointCards(cards)
	if len(bomb) != 4 {
		return false, nil
	}

	return true, getMaxCard(bomb)
}

// 飞机
func isTriples(cards []Poker) (bool, *Poker) {
	planeCount := len(cards)/3
	if planeCount < 2 {
		return false, nil
	}

	remain := cards
	for i:=0; i<planeCount; i++ {
		triple := getMaxSamePointCards(remain)
		if len(triple) != 3 {
			return false, nil
		}
		remain = RemoveAll(remain, triple)
	}
	return true, getMaxCard(cards)
}

// 飞机带对子
func isTriplesAndPairs(cards []Poker) (bool, *Poker) {
	planeCount := len(cards)/5
	if planeCount < 2 {
		return false, nil
	}

	triples := []Poker{}
	remain := cards
	for i:=0; i<planeCount; i++ {
		triple := getMaxSamePointCards(remain)
		if len(triple) != 3 {
			return false, nil
		}
		for j := range triple{
			triples = append(triples, triple[j])
		}
		remain = RemoveAll(remain, triple)
	}

	for i:=0; i<planeCount; i++ {
		pair := getMaxSamePointCards(remain)
		if len(pair) != 2 {
			return false, nil
		}
		remain = RemoveAll(remain, pair)
	}
	return true, getMaxCard(triples)
}

// 飞机带单张
func isTriplesAndSingles(cards []Poker) (bool, *Poker) {
	planeCount := len(cards)/4
	if planeCount < 2 {
		return false, nil
	}

	triples := []Poker{}
	remain := cards
	for i:=0; i<planeCount; i++ {
		triple := getMaxSamePointCards(cards)
		if len(triple) != 3 {
			return false, nil
		}
		for j := range triple{
			triples = append(triples, triple[j])
		}
		remain = RemoveAll(remain, triple)
	}

	return true, getMaxCard(triples)
}

// 连对
func isPairs(cards []Poker) (bool, *Poker) {
	pairs := len(cards)/2
	if pairs < 3 {
		return false, nil
	}

	shunzi := []Poker{}
	remain := cards
	for i:=0; i< pairs; i++ {
		pair := getMaxSamePointCards(cards)
		if len(pair) != 2 {
			return false, nil
		}
		shunzi = append(shunzi, pair[0])
		remain = RemoveAll(remain, pair)
	}

	return isMinShunZi(shunzi, 3)
}

// 顺子
func isShunZi(cards []Poker) (bool, *Poker) {
	return isMinShunZi(cards, 5)
}

// 带最小长度的顺子判断
func isMinShunZi(cards []Poker, minLen int) (bool, *Poker) {
	if len(cards) < minLen {
		return false, nil
	}

	if ContainsPoint(cards, p2) || ContainsPoint(cards, pRedJoker) || ContainsPoint(cards, pBlackJoker) { //有2或着大小王直接返回
		return false, nil
	}

	ddzPointSort(cards)
	for i:=0; i<len(cards)-1; i++ {
		if cards[i+1].PointWeight- cards[i].PointWeight != 1 {
			return false, nil
		}
	}

	return true, getMaxCard(cards)
}

// 三张
func isTriple(cards []Poker) (bool, *Poker) {
	if len(cards) != 3 {
		return false, nil
	}
	return isAllSamePoint(cards), getMaxCard(cards)
}

// 三张带对子
func isTripleAndPair(cards []Poker) (bool, *Poker) {
	if len(cards) != 5 {
		return false, nil
	}

	triple := getMaxSamePointCards(cards)
	if len(triple) != 3 {
		return false, nil
	}

	remain := RemoveAll(cards, triple)
	return isAllSamePoint(remain), getMaxCard(triple)
}

// 三张带单张
func isTripleAndSingle(cards []Poker) (bool, *Poker) {
	if len(cards) != 4 {
		return false, nil
	}

	triple := getMaxSamePointCards(cards)
	if len(triple) != 3 {
		return false, nil
	}

	return true, getMaxCard(triple)
}

// 对子
func isPair(cards []Poker) (bool, *Poker) {
	if len(cards) != 2 {
		return false, nil
	}
	return isAllSamePoint(cards), getMaxCard(cards)
}

// 单张
func isSingle(cards []Poker) (bool, *Poker) {
	if len(cards) != 1 {
		return false, nil
	}
	return true, &cards[0]
}

func getMaxCard(cards []Poker) *Poker{
	ddzPokerSort(cards)
	return &cards[len(cards)-1]
}

// 获取最大相同点数的牌, 如 444555533 返回 5555
func getMaxSamePointCards(cards []Poker) []Poker {
	point, count := getMaxSamePoint(cards)
	maxSamePointCards := make([]Poker, count)
	for _, card := range cards {
		if card.Point == point {
			maxSamePointCards = append(maxSamePointCards, card)
		}
	}
	return maxSamePointCards
}

// 获取最大相同点数, 如 444555533 返回 5,4
func getMaxSamePoint(cards []Poker) (maxCountPoint uint32, maxCount uint32) {
	counts := make(map[uint32]uint32) //Map<Point, count>
	for _, card := range cards {
		point := card.Point
		count, exists := counts[point]
		if !exists {
			counts[point] = 1
		} else {
			counts[point] = count + 1
		}
	}

	for point, count := range counts {
		if count > maxCount {
			maxCount = count
			maxCountPoint = point
		}
	}
	return
}

func isAllSamePoint(cards []Poker) bool {
	for i:=0;i<len(cards)-1;i++ {
		if cards[i].Point != cards[i+1].Point {
			return false
		}
	}
	return true
}