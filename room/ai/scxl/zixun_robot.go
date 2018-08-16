package scxlai

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/spf13/viper"
	"steve/common/mjoption"
	"steve/entity/majong"
	"steve/gutils"
	"steve/room/ai"
	"steve/room/majong/global"
	"steve/room/majong/utils"
	"strconv"
)

func (h *zixunStateAI) generateRobot(player *majong.Player, mjContext *majong.MajongContext) (aiEvent ai.AIEvent) {
	zxRecord := player.GetZixunRecord()
	handCards := player.GetHandCards()
	canHu := zxRecord.GetEnableZimo()
	if (gutils.IsHu(player) || gutils.IsTing(player)) && canHu {
		aiEvent = h.hu(player)
		return
	}
	// 优先出定缺牌
	if gutils.CheckHasDingQueCard(mjContext, player) {
		for i := len(handCards) - 1; i >= 0; i-- {
			hc := handCards[i]
			if mjoption.GetXingpaiOption(int(mjContext.GetXingpaiOptionId())).EnableDingque &&
				hc.GetColor() == player.GetDingqueColor() {
				aiEvent = h.chupai(player, hc)
				return
			}
		}
	}

	logEntry := logrus.WithField("playerId", player.PlayerId)
	outCard, gang := getOutCard(handCards, mjContext)
	needHu := outCard == majong.Card{}
	if needHu {
		aiEvent = h.hu(player)
		logEntry.Infoln("中级AI自摸胡牌")
		return
	}
	if len(zxRecord.EnableBugangCards) > 0 {
		buGangCard := global.ToMJCard(int(zxRecord.EnableBugangCards[0]))
		aiEvent = h.gang(player, &buGangCard)
		logEntry.WithField("补杠牌", buGangCard).Infoln("中级AI补杠")
		return
	}
	if gang && len(zxRecord.EnableAngangCards) > 0 {
		aiEvent = h.gang(player, &outCard)
		logEntry.WithField("暗杠牌", outCard).Infoln("中级AI暗杠")
		return
	}

	if len(zxRecord.CanTingCardInfo) > 0 {
		max := -1
		var maxCard majong.Card
		remainCardCount := getRemainCardCount(handCards, mjContext)
		for _, canTingInfo := range zxRecord.CanTingCardInfo {
			total := 0
			for _, tingInfo := range canTingInfo.TingCardInfo {
				count := remainCardCount[gutils.Uint32ToServerCard(tingInfo.TingCard)]
				total += count * int(tingInfo.Times)
			}
			if max < total {
				max = total
				maxCard = gutils.Uint32ToServerCard(canTingInfo.OutCard)
			}
		}
		aiEvent = h.chupai(player, &maxCard)
		logEntry.WithField("outCard", maxCard).Infoln("中级AI听牌")
		return
	}

	aiEvent = h.chupai(player, &outCard)
	logEntry.WithField("outCard", outCard).Infoln("中级AI出牌")
	return
}

func getOutCard(handCards []*majong.Card, mjContext *majong.MajongContext) (majong.Card, bool) {
	// 拆牌，比较顺子优先、刻子优先两种拆牌方式，选出最好的结果
	s := SplitCards(NonPointer(handCards))
	if len(s.Gangs) > 0 {
		return s.Gangs[0].cards[0], true //有杠就杠
	}

	if len(s.Singles) == 1 {
		return s.Singles[0].cards[0], false //只有一张单牌，直接出牌
	}
	remainCardCount := getRemainCardCount(handCards, mjContext)
	if len(s.Singles) > 1 {
		outCard := whichSingle(remainCardCount, s.Singles, len(s.Pairs) >= 1)
		return outCard, false
	}

	var twoCards []Split
	for _, singleCha := range s.SingleChas {
		twoCards = append(twoCards, singleCha)
	}
	for _, doubleCha := range s.DoubleChas {
		twoCards = append(twoCards, doubleCha)
	}
	if len(s.Pairs) > 1 { //只有一个对子保留作将，多于一个对子才拆对子
		for _, pair := range s.Pairs {
			twoCards = append(twoCards, pair)
		}
	}

	chances := make(map[*Split]int)
	xpOption := mjoption.GetXingpaiOption(int(mjContext.GetXingpaiOptionId()))
	for _, twoCard := range twoCards {
		count := countValidCard(remainCardCount, getValidCard(twoCard))
		if twoCard.t == PAIR {
			chances[&twoCard] = count * (4 / 1) // 可碰，四家摸到都有可能成刻
		}
		if twoCard.t == DOUBLE_CHA || twoCard.t == SINGLE_CHA {
			if xpOption.EnableChi {
				chances[&twoCard] = count * (4 / 2) // 可吃，两家摸到都有可能成顺
			} else {
				chances[&twoCard] = count * (4 / 4) // 不可吃，只有自家摸到才可能成顺
			}
		}
	}

	var needChai Split
	minChance := 99
	for split, chance := range chances {
		if chance < minChance {
			minChance = chance
			needChai = *split
		} else if chance == minChance && split.t > needChai.t {
			needChai = *split
		}
	}
	singles := SplitSingle(needChai.cards)
	outCard := whichSingle(remainCardCount, singles, needChai.t == PAIR && len(s.Pairs) >= 2 || needChai.t != PAIR && len(s.Pairs) >= 1)
	return outCard, false
}

func getRemainCardCount(handCards []*majong.Card, mjContext *majong.MajongContext) map[majong.Card]int {
	visible := viper.GetBool("ai.visible")
	if visible {
		var wallCards []majong.Card
		for _, wallCard := range mjContext.WallCards {
			wallCards = append(wallCards, *wallCard)
		}

		remainCards := CountCard(wallCards)
		return remainCards
	} else {
		var visibleCards []*majong.Card
		visibleCards = append(visibleCards, handCards...)
		for _, player := range mjContext.Players {
			visibleCards = append(visibleCards, player.OutCards...)
			visibleCards = append(visibleCards, utils.TransChiCard(player.ChiCards)...)
			visibleCards = append(visibleCards, utils.TransPengCard(player.PengCards)...)
			visibleCards = append(visibleCards, utils.TransGangCard(player.GangCards)...)
			visibleCards = append(visibleCards, utils.TransHuCard(player.HuCards)...)
		}

		countMap := CountCard(NonPointer(visibleCards))

		remainCards := make(map[majong.Card]int)
		for k, v := range countMap {
			remainCards[k] = 4 - v
		}
		return remainCards
	}
}

func whichSingle(remainCards map[majong.Card]int, singles []Split, cha bool) majong.Card {
	min := 99
	var outCard majong.Card
	if cha { //有将，比较成茬机会数
		for _, single := range singles {
			validCards := getValidCard(single)
			chance := countValidCard(remainCards, validCards)
			if chance < min {
				min = chance
				outCard = single.cards[0]
			}
		}

	} else { //无将，比较成将机会数
		for _, single := range singles {
			chance := remainCards[single.cards[0]]
			if chance < min {
				min = chance
				outCard = single.cards[0]
			}
		}
	}
	return outCard
}

func countValidCard(remainCards map[majong.Card]int, validCards []majong.Card) int {
	total := 0
	for _, validCard := range validCards {
		total += remainCards[validCard]
	}
	return total
}

func getValidCard(split Split) (result []majong.Card) {
	if split.t == SINGLE { //单牌成茬有效牌
		single := split.cards[0]
		if single.Color == majong.CardColor_ColorHua || single.Color == majong.CardColor_ColorZi {
			return
		}
		for _, addend := range []int32{-2, -1, 1, 2} {
			if single.Point+addend >= 1 && single.Point+addend <= 9 {
				result = append(result, majong.Card{Color: single.Color, Point: single.Point + addend})
			}
		}
	}
	if split.t == PAIR { //对子成刻有效牌
		result = append(result, split.cards[0])
	}
	if split.t == DOUBLE_CHA {
		small := split.cards[0]
		result = append(result, majong.Card{Color: small.Color, Point: small.Point - 1})
		result = append(result, majong.Card{Color: small.Color, Point: small.Point + 2})
	}
	if split.t == SINGLE_CHA {
		small := split.cards[0]
		if ContainsEdge(split.cards) { // 12 89
			if small.Point == 1 {
				result = append(result, majong.Card{Color: small.Color, Point: 3})
			} else {
				result = append(result, majong.Card{Color: small.Color, Point: 7})
			}
		} else { // 13 24 35 ... 79
			result = append(result, majong.Card{Color: small.Color, Point: small.Point + 1})
		}
	}
	return
}

func SplitCards(cards []majong.Card) Splits {
	var splits Splits
	splits1 := splitCardsWithoutGang(cards, true)
	splits2 := splitCardsWithoutGang(cards, false)

	if splits1.BetterThan(splits2) {
		splits = splits1
	} else {
		splits = splits2
	}

	gangs := SplitGang(cards)
	remain := RemoveSplits(cards, gangs)
	splits3 := splitCardsWithoutGang(remain, true)
	splits3.Gangs = gangs
	splits4 := splitCardsWithoutGang(remain, false)
	splits4.Gangs = gangs

	if splits3.BetterThan(splits) {
		splits = splits3
	}

	if splits4.BetterThan(splits) {
		splits = splits4
	}

	logrus.WithFields(logrus.Fields{"手牌": cards, "拆牌": splits}).Debugln("中级AI拆牌结果")
	return splits
}

func splitCardsWithoutGang(cards []majong.Card, shunZiFirst bool) Splits {
	var s Splits
	remain := cards
	if shunZiFirst {
		s.ShunZis = SplitShunZi(remain)
		remain = RemoveSplits(remain, s.ShunZis)
		s.KeZis = SplitKeZi(remain)
		remain = RemoveSplits(remain, s.KeZis)
	} else {
		s.KeZis = SplitKeZi(remain)
		remain = RemoveSplits(remain, s.KeZis)
		s.ShunZis = SplitShunZi(remain)
		remain = RemoveSplits(remain, s.ShunZis)
	}
	s.Pairs = SplitPair(remain)
	remain = RemoveSplits(remain, s.Pairs)
	s.DoubleChas = SplitDoubleCha(remain)
	remain = RemoveSplits(remain, s.DoubleChas)
	s.SingleChas = SplitSingleCha(remain)
	remain = RemoveSplits(remain, s.SingleChas)
	s.Singles = SplitSingle(remain)
	return s
}

type SplitType int

const (
	GANG       SplitType = iota //杠，四张相同的牌，已成牌
	KEZI                        //刻子，三张相同的牌，已成牌
	SHUNZI                      //顺子，如567，已成牌
	PAIR                        //对子，如55，一步成刻
	DOUBLE_CHA                  //双茬，如56，一步成顺
	SINGLE_CHA                  //单茬，如57，89，一步成顺
	SINGLE                      //单牌，如5，两步成牌
)

type Split struct {
	t     SplitType
	cards []majong.Card
}

func (s Split) String() string {
	var str string
	for _, card := range s.cards {
		str = str + card.String() + ","
	}
	if len(str) > 0 {
		str = str[0 : len(str)-1]
	}
	switch s.t {
	case GANG:
		return "杠(" + str + ")"
	case KEZI:
		return "刻子(" + str + ")"
	case SHUNZI:
		return "顺子(" + str + ")"
	case PAIR:
		return "对子(" + str + ")"
	case DOUBLE_CHA:
		return "双茬(" + str + ")"
	case SINGLE_CHA:
		return "单茬(" + str + ")"
	case SINGLE:
		return "单牌(" + str + ")"
	default:
		return strconv.Itoa(int(s.t))
	}
}

type Splits struct {
	Gangs      []Split
	ShunZis    []Split
	KeZis      []Split
	Pairs      []Split
	DoubleChas []Split
	SingleChas []Split
	Singles    []Split
}

func (s Splits) String() string {
	return fmt.Sprint(append(append(append(append(append(append(s.Gangs, s.ShunZis...), s.KeZis...), s.Pairs...), s.DoubleChas...), s.SingleChas...), s.Singles...))
}

func (s Splits) BetterThan(o Splits) bool {
	if len(s.Gangs)+len(s.ShunZis)+len(s.KeZis) > len(o.Gangs)+len(o.ShunZis)+len(o.KeZis) {
		return true
	} else if len(s.Gangs)+len(s.ShunZis)+len(s.KeZis) < len(o.Gangs)+len(o.ShunZis)+len(o.KeZis) {
		return false
	}

	if len(s.Gangs) > len(o.Gangs) {
		return true
	} else if len(s.Gangs) < len(o.Gangs) {
		return false
	}

	if len(s.Pairs) > len(o.Pairs) {
		return true
	} else if len(s.Pairs) < len(o.Pairs) {
		return false
	}

	if len(s.DoubleChas) > len(o.DoubleChas) {
		return true
	} else if len(s.DoubleChas) < len(o.DoubleChas) {
		return false
	}

	if len(s.SingleChas) > len(o.SingleChas) {
		return true
	} else if len(s.SingleChas) < len(o.SingleChas) {
		return false
	}

	return len(s.Singles) < len(o.Singles)
}

// 拆出所有杠
func SplitGang(handCards []majong.Card) (result []Split) {
	gangs := FindAllShunZi(handCards, 4, 1)
	for _, gang := range gangs {
		result = append(result, Split{GANG, gang})
	}
	return
}

// 拆出所有刻子
func SplitKeZi(handCards []majong.Card) (result []Split) {
	keZis := FindAllShunZi(handCards, 3, 1)
	for _, keZi := range keZis {
		result = append(result, Split{KEZI, keZi})
	}
	return
}

// 拆出所有顺子
func SplitShunZi(handCards []majong.Card) (result []Split) {
	shunZis := FindAllShunZi(handCards, 1, 3)
	for _, shunZi := range shunZis {
		result = append(result, Split{SHUNZI, shunZi})
	}
	return
}

// 拆出所有对子
func SplitPair(handCards []majong.Card) (result []Split) {
	pairs := FindAllShunZi(handCards, 2, 1)
	for _, pair := range pairs {
		result = append(result, Split{PAIR, pair})
	}
	return
}

// 拆出所有双茬
func SplitDoubleCha(cards []majong.Card) []Split {
	doubleCha, _ := getNearShunZi(cards)
	return doubleCha
}

func getNearShunZi(handCards []majong.Card) (doubleCha []Split, singleCha []Split) {
	result := FindAllShunZi(handCards, 1, 2)

	for _, split := range result {
		if ContainsEdge(split) {
			singleCha = append(singleCha, Split{SINGLE_CHA, split})
		} else {
			doubleCha = append(doubleCha, Split{DOUBLE_CHA, split})
		}
	}
	return
}

// 拆出所有单茬
func SplitSingleCha(cards []majong.Card) []Split {
	_, singleCha := getNearShunZi(cards)
	remain := RemoveSplits(cards, singleCha)
	singleCha = append(singleCha, getSpaceShunZi(remain)...)
	return singleCha
}

func getSpaceShunZi(handCards []majong.Card) (result []Split) {
	spaceShunZis := FindAllCommonShunZi(handCards, 1, 2, 2)
	for _, spaceShunZi := range spaceShunZis {
		result = append(result, Split{SINGLE_CHA, spaceShunZi})
	}
	return
}

// 拆成单牌
func SplitSingle(cards []majong.Card) []Split {
	var result []Split
	for _, card := range cards {
		result = append(result, Split{t: SINGLE, cards: []majong.Card{card}})
	}
	return result
}

/**
FindAllShunZi 双向夹击，找出手牌中所有顺子长度为shunZiLen，重复次数为duplicateCount的牌
*/
func FindAllShunZi(handCards []majong.Card, duplicateCount int, shunZiLen int) (result [][]majong.Card) {
	return FindAllCommonShunZi(handCards, duplicateCount, shunZiLen, 1)
}
func FindAllCommonShunZi(handCards []majong.Card, duplicateCount int, shunZiLen int, shunZiGap int) (result [][]majong.Card) {
	countMap := CountCard(handCards)
	var matchCards []majong.Card
	for card, count := range countMap {
		if count >= duplicateCount {
			matchCards = append(matchCards, card)
		}
	}
	gap := shunZiLen - 1

	colorCards := divideByColor(matchCards)
	for color, cards := range colorCards {
		MJCardSort(cards)
		if color == majong.CardColor_ColorHua || color == majong.CardColor_ColorZi && shunZiLen != 1 {
			continue //花牌都按单牌处理，字牌没有顺子
		}
		i := 0
		j := len(cards) - 1
		for {
			if i+gap <= len(cards)-1 && cards[i+gap].Point-cards[i].Point == int32(gap*shunZiGap) && existAll(countMap, cards, i, i+gap, duplicateCount) { //从1向9取
				shunZi := cards[i : i+gap+1]
				inflated := InflateAll(shunZi, duplicateCount)
				result = append(result, inflated)
				decreaseAll(countMap, shunZi, duplicateCount)
				continue //重复取
			} else {
				i++
			}

			if j-gap >= 0 && cards[j].Point-cards[j-gap].Point == int32(gap*shunZiGap) && existAll(countMap, cards, j-gap, j, duplicateCount) { //从9向1取
				shunZi := cards[j-gap : j+1]
				inflated := InflateAll(shunZi, duplicateCount)
				result = append(result, inflated)
				decreaseAll(countMap, shunZi, duplicateCount)
				continue //重复取
			} else {
				j--
			}
			if i > j {
				break
			}
		}
	}
	return
}

// 按万、条、筒、字拆分手牌
func divideByColor(cards []majong.Card) map[majong.CardColor][]majong.Card {
	colors := make(map[majong.CardColor][]majong.Card)
	for _, card := range cards {
		colors[card.Color] = append(colors[card.Color], card)
	}
	return colors
}

func decreaseAll(countMap map[majong.Card]int, shunZi []majong.Card, duplicateCount int) {
	for _, card := range shunZi {
		countMap[card] -= duplicateCount
	}
}

func existAll(countMap map[majong.Card]int, matchCards []majong.Card, start int, end int, duplicateCount int) bool {
	for i := start; i <= end; i++ {
		card := matchCards[i]
		if countMap[card] < duplicateCount {
			return false
		}
	}
	return true
}

func InflateAll(cards []majong.Card, duplicateCount int) (result []majong.Card) {
	for _, card := range cards {
		for i := 0; i < duplicateCount; i++ {
			result = append(result, card)
		}
	}
	return
}

func CountCard(cards []majong.Card) map[majong.Card]int {
	countMap := make(map[majong.Card]int)
	for _, card := range cards {
		countMap[card]++
	}
	return countMap
}
