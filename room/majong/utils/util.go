package utils

import (
	"fmt"
	"steve/client_pb/room"
	"steve/common/mjoption"
	majongpb "steve/entity/majong"
	"steve/gutils"
	"steve/room/majong/interfaces"

	"github.com/golang/protobuf/proto"
)

var cardAll = []Card{11, 12, 13, 14, 15, 16, 17, 18, 19, 21, 22, 23, 24, 25, 26, 27, 28, 29, 31, 32, 33, 34, 35, 36, 37, 38, 39, 41, 42, 43, 44, 45, 46, 47}

//GetPlayerByID 根据玩家id获取玩家
func GetPlayerByID(players []*majongpb.Player, id uint64) *majongpb.Player {
	for _, player := range players {
		if player.PlayerId == id {
			return player
		}
	}
	return nil
}

//GetNextPlayerByID 根据玩家id获取下个玩家
func GetNextPlayerByID(players []*majongpb.Player, id uint64) *majongpb.Player {
	for k, player := range players {
		if player.PlayerId == id {
			index := (k + 1) % len(players)
			return players[index]
		}
	}
	return nil
}

//CardsToUtilCards 用来辅助查ting胡的工具,将Card转为适合查胡的工具
func CardsToUtilCards(cards []*majongpb.Card) []Card {
	cardsCard := make([]Card, 0)
	for _, v := range cards {
		cardInt := ServerCard2Number(v)
		cardsCard = append(cardsCard, Card(cardInt))
	}
	return cardsCard
}

//HuCardsToUtilCards 用来辅助查ting胡的工具,将Card转为适合查胡的工具
func HuCardsToUtilCards(cards []*majongpb.HuCard) []Card {
	cardsCard := make([]Card, 0)
	for _, v := range cards {

		cardInt := ServerCard2Uint32(v.Card)
		cardsCard = append(cardsCard, Card(cardInt))
	}
	return cardsCard
}

//DeleteIntCardFromLast 从int32类型的牌组中，找到对应的目标牌，并且移除
func DeleteIntCardFromLast(cards []int32, targetCard int32) ([]int32, bool) {
	index := -1
	l := len(cards)
	if l == 0 {
		return cards, false
	}
	for i := l - 1; i >= 0; i-- {
		if targetCard == cards[i] {
			index = i
			break
		}
	}
	if index != -1 {
		cards = append(cards[:index], cards[index+1:]...)
	}
	return cards, index != -1
}

//CardEqual 判断两张牌是否一样
func CardEqual(card1 *majongpb.Card, card2 *majongpb.Card) bool {
	return card1.GetColor() == card2.GetColor() && card1.GetPoint() == card2.GetPoint()
}

//DeleteCardFromLast 从majongpb.Card类型的牌组中，找到对应的目标牌，并且移除
func DeleteCardFromLast(cards []*majongpb.Card, targetCard *majongpb.Card) ([]*majongpb.Card, bool) {
	index := -1
	l := len(cards)
	if l == 0 {
		return cards, false
	}
	for i := l - 1; i >= 0; i-- {
		if CardEqual(targetCard, cards[i]) {
			index = i
			break
		}
	}
	if index != -1 {
		cards = append(cards[:index], cards[index+1:]...)
	}
	return cards, index != -1
}

// RemoveCards 从玩家的手牌中移除指定数量的某张牌
func RemoveCards(cards []*majongpb.Card, card *majongpb.Card, count int) ([]*majongpb.Card, bool) {
	newCards := []*majongpb.Card{}
	removeCount := 0
	for index, c := range cards {
		if CardEqual(c, card) {
			removeCount++
			if removeCount == count {
				newCards = append(newCards, cards[index+1:]...)
				break
			}
		} else {
			newCards = append(newCards, c)
		}
	}
	if removeCount != count {
		return cards, false
	}
	return newCards, true
}

//IntToCard int32类型转majongpb.Card类型
func IntToCard(cardValue int32) (*majongpb.Card, error) {
	colorValue := cardValue / 10
	value := cardValue % 10
	var color majongpb.CardColor
	switch colorValue {
	case 1:
		color = majongpb.CardColor_ColorWan
	case 2:
		color = majongpb.CardColor_ColorTiao
	case 3:
		color = majongpb.CardColor_ColorTong
	case 4:
		color = majongpb.CardColor_ColorZi
	case 5:
		color = majongpb.CardColor_ColorHua
	default:
		return nil, fmt.Errorf("cant trans card %d", cardValue)
	}
	return &majongpb.Card{
		Color: color,
		Point: value,
	}, nil
}

//IntToRoomCard int32类型转room.Card类型
func IntToRoomCard(cardValue int32) (*room.Card, error) {
	colorValue := cardValue / 10
	value := cardValue % 10
	var color room.CardColor
	switch colorValue {
	case 1:
		color = room.CardColor_CC_WAN
	case 2:
		color = room.CardColor_CC_TIAO
	case 3:
		color = room.CardColor_CC_TONG
	default:
		return nil, fmt.Errorf("cant trans card %d", cardValue)
	}
	return &room.Card{
		Color: color.Enum(),
		Point: proto.Int32(value),
	}, nil
}

//CardToRoomCard majongpb.card类型转room.Card类型
func CardToRoomCard(card *majongpb.Card) (*room.Card, error) {
	var color room.CardColor
	if card.Color == majongpb.CardColor_ColorWan {
		color = room.CardColor_CC_WAN
	}
	if card.Color == majongpb.CardColor_ColorTiao {
		color = room.CardColor_CC_TIAO
	}
	if card.Color == majongpb.CardColor_ColorTong {
		color = room.CardColor_CC_TONG
	}
	if card.Color == majongpb.CardColor_ColorHua {
		color = room.CardColor_CC_HUA
	}
	return &room.Card{
		Color: color.Enum(),
		Point: proto.Int32(card.Point),
	}, nil
}

// ServerCard2UtilCard pb的 Card 转 Card
func ServerCard2UtilCard(card *majongpb.Card) Card {
	return Card(ServerCard2Number(card))
}

// ServerCards2UtilsCards pb 的 Card 数组转 Card 数组
func ServerCards2UtilsCards(cards []*majongpb.Card) []Card {
	result := []Card{}
	for _, card := range cards {
		result = append(result, ServerCard2UtilCard(card))
	}
	return result
}

// ServerCard2Number 服务器的 Card 转换成数字
func ServerCard2Number(card *majongpb.Card) int {
	var color int
	if card.Color == majongpb.CardColor_ColorWan {
		color = 1
	} else if card.Color == majongpb.CardColor_ColorTiao {
		color = 2
	} else if card.Color == majongpb.CardColor_ColorTong {
		color = 3
	} else if card.Color == majongpb.CardColor_ColorZi {
		color = 4
	} else if card.Color == majongpb.CardColor_ColorHua {
		color = 5
	}
	value := color*10 + int(card.Point)
	return value
}

// ServerCards2Numbers 服务器的 Card 数组转 int 数组
func ServerCards2Numbers(cards []*majongpb.Card) []int {
	result := []int{}
	for _, c := range cards {
		result = append(result, ServerCard2Number(c))
	}
	return result
}

// ServerCard2Uint32 服务器的 Card 转换成数字
func ServerCard2Uint32(card *majongpb.Card) uint32 {
	return uint32(ServerCard2Number(card))
}

// ServerCards2Uint32 服务器的 Card 数组转 int 数组
func ServerCards2Uint32(cards []*majongpb.Card) []uint32 {
	result := []uint32{}
	for _, c := range cards {
		result = append(result, ServerCard2Uint32(c))
	}
	return result
}

// CardsToRoomCards 将Card转换为room package中的Card
func CardsToRoomCards(cards []*majongpb.Card) []*room.Card {
	var rCards []*room.Card
	for i := 0; i < len(cards); i++ {
		rCards = append(rCards, &room.Card{
			Color: gutils.ServerColor2ClientColor(cards[i].Color).Enum(),
			Point: &cards[i].Point,
		})
	}
	return rCards
}

// ContainCard 验证card是否存在于玩家手牌中，存在返回true,否则返回false
func ContainCard(cards []*majongpb.Card, card *majongpb.Card) bool {
	for i := 0; i < len(cards); i++ {
		if CardEqual(cards[i], card) {
			return true
		}
	}
	return false
}

//IntToUtilCard uint32类型的数组强转成类型
func IntToUtilCard(cards []int32) []Card {
	cardsCard := make([]Card, 0, 0)
	for _, v := range cards {

		utilCard := Card(v)
		cardsCard = append(cardsCard, utilCard)
	}
	return cardsCard
}

//ContainHuCards 判断当前可以胡的牌中是否包含已经胡过的所有牌
func ContainHuCards(targetHuCards []Card, HuCards []Card) bool {
	sameHuCards := 0
	for _, tagetCard := range targetHuCards {
		for _, Card := range HuCards {
			if tagetCard == Card {
				sameHuCards++
			}
		}
	}
	if len(HuCards) == sameHuCards {
		return true
	}
	return false
}

// CheckHuResult 查胡结果
type CheckHuResult struct {
	Can      bool
	Combines Combines // 推倒胡组合
}

// CheckHu 用来辅助胡牌查胡工具 cards玩家的所有牌，huCard点炮的牌（自摸时huCard为0）
// needCombines 是否需要返回所有组合
func CheckHu(cards []*majongpb.Card, huCard uint32, needCombines bool) CheckHuResult {
	result := CheckHuResult{}
	cardsCard := CardsToUtilCards(cards)
	if huCard > 0 {
		cardsCard = append(cardsCard, Card(huCard))
	}
	laizi := make(map[Card]bool)
	flag, combines := FastCheckHuV2(cardsCard, laizi, needCombines) // 检测玩家能否推倒胡
	canQidui := FastCheckQiDuiHu(cardsCard)
	result.Can = result.Can || flag || canQidui
	result.Combines = combines
	return result
}

//CheckHuUtilCardsToHandCards 将推到胡工具的util.Card转为玩家手牌的类型
func CheckHuUtilCardsToHandCards(cards []Card) ([]*majongpb.Card, error) {
	handCards := make([]*majongpb.Card, 0)
	for i := 0; i < len(cards); i++ {
		handCard, err := IntToCard(int32(cards[i]))
		if err != nil {
			return []*majongpb.Card{}, err
		}
		handCards = append(handCards, handCard)
	}
	return handCards, nil
}

//SeekCardSum 相同的牌的数量
func SeekCardSum(cards []*majongpb.Card, targetCard *majongpb.Card) int {
	count := 0
	for i := 0; i < len(cards); i++ {
		if CardEqual(cards[i], targetCard) {
			count++
		}
	}
	return count
}

//GetTingPlayerIDAndMultiple 获取所有听玩家,和返回每个听玩家最大倍数
func GetTingPlayerIDAndMultiple(mjContext *majongpb.MajongContext, players []*majongpb.Player, laizi map[Card]bool) (map[uint64]int64, error) {
	tingPlayers := make(map[uint64]int64, 0)
	for i := 0; i < len(players); i++ {
		// 胡过的不算
		if len(players[i].HuCards) > 0 {
			continue
		}
		// 查能不能听，能听，返回返回最大番型，及ID
		isTing, multiple, err := IsCanTingAndGetMultiple(mjContext, players[i], laizi)
		if err != nil {
			return nil, err
		}
		if isTing {
			tingPlayers[players[i].GetPlayerId()] = multiple
		}
	}
	return tingPlayers, nil
}

//IsCanTingAndGetMultiple 判断玩家是否能听,和返回能听玩家的最大倍数 TODO
//未上听者需赔上听者最大可能番数（杠后炮、杠上开花、抢杠胡、海底捞、海底炮不参与）的牌型钱。注：查大叫时，
//若上听者牌型中有根，则根也要未上听者包给上听者。
func IsCanTingAndGetMultiple(mjContext *majongpb.MajongContext, player *majongpb.Player, laizi map[Card]bool) (bool, int64, error) {
	var max int64
	handCardSum := len(player.HandCards)
	//只差1张牌就能胡，并且玩家手牌不存在花牌
	if handCardSum%3 == 1 && !gutils.CheckHasDingQueCard(mjContext, player) {
		tingCards, err := GetTingCards(player.HandCards, laizi)
		if err != nil {
			return false, 0, err
		}
		handCards := player.GetHandCards()
		for _, card := range tingCards {
			pbCard, _ := IntToCard(int32(card))
			handCards = append(handCards, pbCard)
			// TODO 获取最大番型
			mult := int64(2)
			if max < mult {
				max = mult
			}
			handCards = player.GetHandCards()
		}
	}
	return max > 0, max, nil
}

//GetTingCards 获取玩家能胡的牌,必须是缺一张
func GetTingCards(handCards []*majongpb.Card, laizi map[Card]bool) ([]Card, error) {
	result := []Card{}

	if len(handCards)%3 != 1 {
		return result, fmt.Errorf("获取玩家能胡的牌,必须是缺一张")
	}
	cardsCard := CardsToUtilCards(handCards)
	// 推倒胡
	result = FastCheckTingV2(cardsCard, laizi)
	// 七对
	qiCards := FastCheckQiDuiTing(cardsCard, cardAll)
	return MergeAndNoRepeat(result, qiCards), nil
}

//MergeAndNoRepeat 合并去重复UtilCard
func MergeAndNoRepeat(srcCards1 []Card, srcCards2 []Card) []Card {
	resultMap := map[Card]struct{}{}
	for _, card := range srcCards1 {
		resultMap[card] = struct{}{}
	}
	for _, card := range srcCards2 {
		resultMap[card] = struct{}{}
	}
	result := make([]Card, 0, len(resultMap))
	for card := range resultMap {
		result = append(result, card)
	}
	return result
}

//GetFirstHuPlayerByID 获取第一个胡的玩家,winPlayers源数组， loserPlayerID输家ID
func GetFirstHuPlayerByID(playerAll, winPlayers []*majongpb.Player, loserPlayerID uint64) *majongpb.Player {
	// 获取输家的下家
	nextPlayer := GetNextPlayerByID(playerAll, loserPlayerID)
	for nextPlayer != nil {
		// 判断赢家里面是否有输家的下家
		for i := 0; i < len(winPlayers); i++ {
			if winPlayers[i].PlayerId == nextPlayer.PlayerId {
				return winPlayers[i]
			}
		}
		// 获取输家的下家的下家
		nextPlayer = GetNextPlayerByID(playerAll, nextPlayer.PlayerId)
	}
	return nil
}

// //GetPlayCardCheckTing 出牌查听，获取可以出那些牌，和出了这张牌，可以胡那些牌，返回map[Card][]Card
// func GetPlayCardCheckTing(handCards []*majongpb.Card, laizi map[Card]bool) map[Card][]Card {
// 	tingInfo := make(map[Card][]Card)
// 	// 不能少一张
// 	if len(handCards)%3 != 2 {
// 		return tingInfo
// 	}
// 	// 手牌转查胡的工具牌
// 	cardsCard := CardsToUtilCards(handCards)
// 	// 推倒胡查胡，打那张牌可以胡那些牌
// 	tingCombines := FastCheckTingInfoV2(cardsCard, laizi)
// 	for card, cardCombines := range tingCombines {
// 		tingcards := []Card{}
// 		for card := range cardCombines {
// 			tingcards = append(tingcards, card)
// 		}
// 		tingInfo[card] = tingcards
// 	}

// 	// 1-9所有牌
// 	cardAll := []Card{11, 12, 13, 14, 15, 16, 17, 18, 19, 21, 22, 23, 24, 25, 26, 27, 28, 29, 31, 32, 33, 34, 35, 36, 37, 38, 39}
// 	// 七对查胡，打那张牌可以胡那些牌
// 	qiStrategy := FastCheckQiDuiTingInfo(cardsCard, cardAll)
// 	// 存在相同的playCard,去重复
// 	for playCard, huCard := range tingInfo {
// 		tInfo, exite := qiStrategy[playCard]
// 		if exite {
// 			tingInfo[playCard] = MergeAndNoRepeat(tInfo, huCard)
// 		}
// 	}
// 	// 存在不相同的playCard,合并,把推倒胡中不存在的听，加进去
// 	for playCard, huCards := range qiStrategy {
// 		_, exite := tingInfo[playCard]
// 		if !exite {
// 			tingInfo[playCard] = huCards
// 		}
// 	}
// 	return tingInfo
// }

//GetPlayCardCheckTing 出牌查听，获取可以出那些牌，和出了这张牌，可以胡那些牌
// 返回可胡的牌与对应的组合
func GetPlayCardCheckTing(handCards []*majongpb.Card, laizi map[Card]bool) map[Card][]Card {
	result := make(map[Card][]Card)
	// 不能少一张
	if len(handCards)%3 != 2 {
		return result
	}
	// 手牌转查胡的工具牌
	cardsCard := CardsToUtilCards(handCards)
	// 推倒胡查胡，打那张牌可以胡那些牌
	result = FastCheckTingInfoV2(cardsCard, laizi)
	// 七对查胡，打那张牌可以胡那些牌
	qiStrategy := FastCheckQiDuiTingInfo(cardsCard, cardAll)

	for card, huCards := range qiStrategy {
		if result[card] != nil {
			result[card] = MergeAndNoRepeat(result[card], huCards)
		} else {
			result[card] = huCards
		}
	}
	return result
}

// TransChiCard 吃牌转Card
func TransChiCard(chiCards []*majongpb.ChiCard) []*majongpb.Card {
	cards := make([]*majongpb.Card, 0)
	for _, chiCard := range chiCards {
		cards = append(cards, chiCard.Card)
	}
	return cards
}

// TransPengCard 碰牌转Card
func TransPengCard(pengCards []*majongpb.PengCard) []*majongpb.Card {
	cards := make([]*majongpb.Card, 0)
	for _, pengCard := range pengCards {
		cards = append(cards, pengCard.Card)
	}
	return cards
}

// TransGangCard 杠牌转Card
func TransGangCard(gangCards []*majongpb.GangCard) []*majongpb.Card {
	cards := make([]*majongpb.Card, 0)
	for _, gangCard := range gangCards {
		cards = append(cards, gangCard.Card)
	}
	return cards
}

// TransHuCard 胡牌转Card
func TransHuCard(huCards []*majongpb.HuCard) []*majongpb.Card {
	cards := make([]*majongpb.Card, 0)
	for _, huCard := range huCards {
		cards = append(cards, huCard.Card)
	}
	return cards
}

// GetAllMopaiCount 获取所有人的摸牌数总和
func GetAllMopaiCount(mjContext *majongpb.MajongContext) int {
	count := 0
	for _, player := range mjContext.GetPlayers() {
		count += int(player.GetMopaiCount())
		count += len(player.GetHuaCards())
	}
	return count
}

// HasAvailableWallCards 判断是否有墙牌可摸
func HasAvailableWallCards(flow interfaces.MajongFlow) bool {
	// 由配牌控制是否gameover,配牌长度为0走正常gameover,配牌长度不为0走配牌长度流局
	if GetAvailableWallCardsNum(flow) > 0 {
		return true
	}
	return false
}

// GetAvailableWallCardsNum 获取可用的墙牌数量
func GetAvailableWallCardsNum(flow interfaces.MajongFlow) int {
	context := flow.GetMajongContext()
	length := context.GetOption().GetWallcardsLength()
	if length == 0 {
		return len(context.GetWallCards())
	}
	fapaiCards := 0
	if mjoption.GetXingpaiOption(int(context.GetXingpaiOptionId())).EnableKaijuAddflower {
		fapaiCards = (len(context.GetPlayers()) * 13)
	} else {
		fapaiCards = (len(context.GetPlayers())*13 + 1)
	}
	return int(length) - (GetAllMopaiCount(context) + fapaiCards)
}

// CardsToInt card 转换
func CardsToInt(cards []*majongpb.Card) ([]int32, error) {
	return gutils.ServerCards2Int32(cards), nil
}

// CheckHuByRemoveGangCards 移除杠牌进行查胡
func CheckHuByRemoveGangCards(player *majongpb.Player, gangCard *majongpb.Card, gangCardNum int) bool {
	handCards := player.GetHandCards()
	newcards := make([]*majongpb.Card, 0, len(handCards))
	newcards = append(newcards, handCards...)
	newcards, _ = RemoveCards(newcards, gangCard, gangCardNum)
	laizi := make(map[Card]bool)
	huCards, _ := GetTingCards(newcards, laizi)
	if len(huCards) > 0 && ContainHuCards(huCards, HuCardsToUtilCards(player.HuCards)) {
		return true
	}
	return false
}
