package scxlai

import (
	"fmt"
	"github.com/magiconair/properties/assert"
	"math/rand"
	"steve/entity/majong"
	"steve/room/majong/global"
	"testing"
)

func Test_GetShunZi(t *testing.T) {
	result := SplitShunZi([]majong.Card{global.Card1W, global.Card3W, global.Card4W, global.Card5W, global.Card6W, global.Card7W, global.Card8W, global.Card9W}) // 13456789万
	assert.Equal(t, result, []Split{
		{SHUNZI, []majong.Card{global.Card7W, global.Card8W, global.Card9W}},
		{SHUNZI, []majong.Card{global.Card3W, global.Card4W, global.Card5W}},
	})
}

func Test_GetShunZi1(t *testing.T) {
	result := SplitShunZi([]majong.Card{global.Card1W, global.Card2W, global.Card3W, global.Card4W, global.Card5W, global.Card6W, global.Card7W, global.Card8W, global.Card9W}) // 123456789万
	assert.Equal(t, result, []Split{
		{SHUNZI, []majong.Card{global.Card1W, global.Card2W, global.Card3W}},
		{SHUNZI, []majong.Card{global.Card7W, global.Card8W, global.Card9W}},
		{SHUNZI, []majong.Card{global.Card4W, global.Card5W, global.Card6W}},
	})
}

func Test_GetKeZi(t *testing.T) {
	result := SplitKeZi([]majong.Card{global.Card3W, global.Card3W, global.Card3W, global.Card4W, global.Card5W, global.Card6W, global.Card7W, global.Card8W, global.Card9W}) // 333456789万
	assert.Equal(t, result, []Split{
		{KEZI, []majong.Card{global.Card3W, global.Card3W, global.Card3W}},
	})
}

func Test_Remove(t *testing.T) {
	cards := []majong.Card{global.Card3W, global.Card3W, global.Card3W}
	cards = Remove(cards, global.Card3W)
	assert.Equal(t, cards, []majong.Card{global.Card3W, global.Card3W})
}

func Test_JoinGang(t *testing.T) {
	s := Splits{KeZis: []Split{{KEZI, []majong.Card{global.Card1T, global.Card1T, global.Card1T}},
		{KEZI, []majong.Card{global.Card2T, global.Card2T, global.Card2T}}},
		Singles: []Split{{SINGLE, []majong.Card{global.Card1T}},
			{SINGLE, []majong.Card{global.Card2T}}},
	}
	s.JoinGang()
	assert.Equal(t, s.Gangs, []Split{{GANG, []majong.Card{global.Card1T, global.Card1T, global.Card1T, global.Card1T}}, {GANG, []majong.Card{global.Card2T, global.Card2T, global.Card2T, global.Card2T}}})
}

func Test_SplitShunZiFirst(t *testing.T) {
	cards := []majong.Card{global.Card5W, global.Card5W, global.Card5W, global.Card6W, global.Card6W, global.Card6W, global.Card6W, global.Card7W, global.Card7W, global.Card7W}
	splits := splitCardsWithoutGang(cards, true)

	assert.Equal(t, splits.ShunZis, []Split{
		{SHUNZI, []majong.Card{global.Card5W, global.Card6W, global.Card7W}},
		{SHUNZI, []majong.Card{global.Card5W, global.Card6W, global.Card7W}},
		{SHUNZI, []majong.Card{global.Card5W, global.Card6W, global.Card7W}},
	})
	assert.Equal(t, splits.Singles, []Split{{SINGLE, []majong.Card{global.Card6W}}})
}

func Test_SplitKeZiFirst(t *testing.T) {
	cards := []majong.Card{global.Card5W, global.Card5W, global.Card5W, global.Card6W, global.Card6W, global.Card6W, global.Card6W, global.Card7W, global.Card7W, global.Card7W}
	splits := splitCardsWithoutGang(cards, false)

	assert.Equal(t, splits.KeZis, []Split{
		{KEZI, []majong.Card{global.Card5W, global.Card5W, global.Card5W}},
		{KEZI, []majong.Card{global.Card7W, global.Card7W, global.Card7W}},
		{KEZI, []majong.Card{global.Card6W, global.Card6W, global.Card6W}},
	})
	assert.Equal(t, splits.Singles, []Split{{SINGLE, []majong.Card{global.Card6W}}})
}

func Test_SplitRemoveShunZi(t *testing.T) {
	cards := []majong.Card{global.Card5W, global.Card5W, global.Card5W, global.Card6W, global.Card6W, global.Card6W, global.Card6W, global.Card7W, global.Card7W, global.Card7W}
	splits := splitCardsWithoutGang(cards, false)
	remain := RemoveSplits(cards, splits.ShunZis)
	gangs := SplitGang(remain)
	assert.Equal(t, gangs, []Split{{GANG, []majong.Card{global.Card6W, global.Card6W, global.Card6W, global.Card6W}}})
}

func Test_SplitSpaceShunZi(t *testing.T) {
	cards := []majong.Card{global.Card5W, global.Card7W}
	splits := splitCardsWithoutGang(cards, false)

	assert.Equal(t, splits.SingleChas, []Split{
		{SINGLE_CHA, []majong.Card{global.Card5W, global.Card7W}},
	})
}

func Test_SplitColorCards(t *testing.T) {
	cards := []majong.Card{global.Card5W, global.Card5T, global.Card5B, global.Card6W, global.Card6T, global.Card6B, global.Card6W, global.Card7W, global.Card7T, global.Card7B}
	splits := splitCardsWithoutGang(cards, false)

	fmt.Println(splits)
	//assert.Equal(t, ShunZis, []Split{
	//	{SHUNZI, []majong.Card{global.Card5W, global.Card6W, global.Card7W}},
	//	{SHUNZI, []majong.Card{global.Card5B, global.Card6B, global.Card7B}},
	//	{SHUNZI, []majong.Card{global.Card5T, global.Card6T, global.Card7T}},
	//})
	assert.Equal(t, splits.Singles, []Split{{SINGLE, []majong.Card{global.Card6W}}})
}

func Test_SplitZiCards(t *testing.T) {
	cards := []majong.Card{global.Card5Z, global.Card5Z, global.Card5Z, global.Card6Z, global.Card6Z, global.Card6Z, global.Card6Z, global.Card7Z, global.Card7Z, global.Card7Z}
	splits := splitCardsWithoutGang(cards, true)

	assert.Equal(t, splits.KeZis, []Split{
		{KEZI, []majong.Card{global.Card5Z, global.Card5Z, global.Card5Z}},
		{KEZI, []majong.Card{global.Card7Z, global.Card7Z, global.Card7Z}},
		{KEZI, []majong.Card{global.Card6Z, global.Card6Z, global.Card6Z}},
	})
	assert.Equal(t, splits.Singles, []Split{{SINGLE, []majong.Card{global.Card6Z}}})
}

func Test_SplitCards(t *testing.T) {
	cards := []majong.Card{global.Card5W, global.Card6W, global.Card7W, global.Card7W, global.Card8W, global.Card9W, global.Card9W}
	splits := splitCardsWithoutGang(cards, true)
	assert.Equal(t, splits.ShunZis, []Split{
		{SHUNZI, []majong.Card{global.Card5W, global.Card6W, global.Card7W}},
		{SHUNZI, []majong.Card{global.Card7W, global.Card8W, global.Card9W}},
	})
}

func Test_SplitCards1(t *testing.T) {
	s := SplitCards([]majong.Card{global.Card4W, global.Card5W, global.Card5W, global.Card6W, global.Card6W, global.Card7W, global.Card7W, global.Card8W, global.Card9W}) // 455667789
	assert.Equal(t, s.ShunZis, []Split{
		{SHUNZI, []majong.Card{global.Card4W, global.Card5W, global.Card6W}},
		{SHUNZI, []majong.Card{global.Card7W, global.Card8W, global.Card9W}},
		{SHUNZI, []majong.Card{global.Card5W, global.Card6W, global.Card7W}},
	})
}

func Test_SplitCards2(t *testing.T) {
	s := SplitCards([]majong.Card{global.Card1W, global.Card1W, global.Card2W, global.Card2W, global.Card3W, global.Card3W, global.Card4W, global.Card4W, global.Card5W, global.Card5W, global.Card6W, global.Card6W}) // 112233445566
	assert.Equal(t, s.ShunZis, []Split{
		{SHUNZI, []majong.Card{global.Card1W, global.Card2W, global.Card3W}},
		{SHUNZI, []majong.Card{global.Card1W, global.Card2W, global.Card3W}},
		{SHUNZI, []majong.Card{global.Card4W, global.Card5W, global.Card6W}},
		{SHUNZI, []majong.Card{global.Card4W, global.Card5W, global.Card6W}},
	})
}

func Test_SplitCards3(t *testing.T) {
	s := SplitCards([]majong.Card{global.Card1W, global.Card2W, global.Card3W, global.Card4W, global.Card5W, global.Card6W, global.Card7W, global.Card8W, global.Card9W, global.Card2W, global.Card3W, global.Card4W, global.Card5W, global.Card6W, global.Card7W}) // 122334455667789
	assert.Equal(t, s.ShunZis, []Split{
		{SHUNZI, []majong.Card{global.Card1W, global.Card2W, global.Card3W}},
		{SHUNZI, []majong.Card{global.Card7W, global.Card8W, global.Card9W}},
		{SHUNZI, []majong.Card{global.Card2W, global.Card3W, global.Card4W}},
		{SHUNZI, []majong.Card{global.Card4W, global.Card5W, global.Card6W}},
		{SHUNZI, []majong.Card{global.Card5W, global.Card6W, global.Card7W}},
	})
}

func Test_SplitCards4(t *testing.T) {
	s := SplitCards([]majong.Card{global.Card1W, global.Card2W, global.Card3W, global.Card4W, global.Card5W, global.Card6W, global.Card7W, global.Card8W, global.Card9W, global.Card2W, global.Card3W, global.Card4W, global.Card6W, global.Card7W, global.Card8W}) // 122334456677889
	assert.Equal(t, s.ShunZis, []Split{
		{SHUNZI, []majong.Card{global.Card1W, global.Card2W, global.Card3W}},
		{SHUNZI, []majong.Card{global.Card7W, global.Card8W, global.Card9W}},
		{SHUNZI, []majong.Card{global.Card2W, global.Card3W, global.Card4W}},
		{SHUNZI, []majong.Card{global.Card6W, global.Card7W, global.Card8W}},
		{SHUNZI, []majong.Card{global.Card4W, global.Card5W, global.Card6W}},
	})
}

func Test_SplitCards5(t *testing.T) {
	s := SplitCards([]majong.Card{global.Card6W, global.Card6W, global.Card1B, global.Card1B, global.Card1B, global.Card1B, global.Card2B, global.Card3B, global.Card3B, global.Card4B, global.Card4B, global.Card5B, global.Card8B, global.Card9B})
	assert.Equal(t, s.Gangs, []Split{{GANG, []majong.Card{global.Card1B, global.Card1B, global.Card1B, global.Card1B}}})
}

func Test_SplitCards6(t *testing.T) {
	s := SplitCards([]majong.Card{global.Card1T, global.Card1T, global.Card1T, global.Card1T, global.Card2T, global.Card2T, global.Card2T, global.Card3T, global.Card3T, global.Card3T, global.Card3T, global.Card4T, global.Card5T, global.Card5T})
	assert.Equal(t, s.Gangs, []Split{{GANG, []majong.Card{global.Card1T, global.Card1T, global.Card1T, global.Card1T}}})
}

func Test_ValidCard(t *testing.T) {
	//单牌
	single := getValidCard(Split{SINGLE, []majong.Card{global.Card1W}})
	assert.Equal(t, single, []majong.Card{global.Card2W, global.Card3W})
	single1 := getValidCard(Split{SINGLE, []majong.Card{global.Card9W}})
	assert.Equal(t, single1, []majong.Card{global.Card7W, global.Card8W})
	single2 := getValidCard(Split{SINGLE, []majong.Card{global.Card5W}})
	assert.Equal(t, single2, []majong.Card{global.Card3W, global.Card4W, global.Card6W, global.Card7W})

	//双茬
	doubleCha := getValidCard(Split{DOUBLE_CHA, []majong.Card{global.Card2W, global.Card3W}})
	assert.Equal(t, doubleCha, []majong.Card{global.Card1W, global.Card4W})

	//单茬
	singleCha := getValidCard(Split{SINGLE_CHA, []majong.Card{global.Card1W, global.Card2W}})
	assert.Equal(t, singleCha, []majong.Card{global.Card3W})
	singleCha1 := getValidCard(Split{SINGLE_CHA, []majong.Card{global.Card8W, global.Card9W}})
	assert.Equal(t, singleCha1, []majong.Card{global.Card7W})
	singleCha2 := getValidCard(Split{SINGLE_CHA, []majong.Card{global.Card3W, global.Card5W}})
	assert.Equal(t, singleCha2, []majong.Card{global.Card4W})
}

func randCard() int {
	suit := rand.Intn(4) + 1
	point := rand.Intn(9) + 1
	return suit*10 + point
}

func Benchmark_SplitCards(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var cards []int
		for j := 0; j < 13; j++ {
			cards = append(cards, randCard())
		}
		SplitCards(global.ToMJCards(cards))
	}
}
