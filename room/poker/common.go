package poker

import (
	"sort"
	"strconv"
)

var (
	RedJoker    = ToPoker(0x0F) //大王
	BlackJoker  = ToPoker(0x0E) //小王
	SDiamond    = uint32(0x10)  //方块
	SClub       = uint32(0x20)  //梅花
	SHeart      = uint32(0x30)  //红桃
	SSpade      = uint32(0x40)  //黑桃
	PA          = uint32(0x01)
	P2          = uint32(0x02)
	P3          = uint32(0x03)
	P4          = uint32(0x04)
	P5          = uint32(0x05)
	P6          = uint32(0x06)
	P7          = uint32(0x07)
	P8          = uint32(0x08)
	P9          = uint32(0x09)
	P10         = uint32(0x0A)
	PJ          = uint32(0x0B)
	PQ          = uint32(0x0C)
	PK          = uint32(0x0D)
	PBlackJoker = uint32(0x0E)
	PRedJoker   = uint32(0x0F)
)

type Poker struct {
	Suit        uint32 //花色 0x00,0x10,0x20,0x30,xx40
	Point       uint32 //点数 0x01-0x0D(A-K), 0x0E(小王), 0x0F(大王)
	Weight      uint32 //带花色权重,用于带花色大小比较，同点数的在一起
	PointWeight uint32 //无花色权重，用于无花色大小比较
}

func (c Poker) String() string {
	if c.Suit == SDiamond {
		return "♦" + c.GetPointString()
	} else if c.Suit == SClub {
		return "♣" + c.GetPointString()
	} else if c.Suit == SHeart {
		return "♥" + c.GetPointString()
	} else if c.Suit == SSpade {
		return "♠" + c.GetPointString()
	} else {
		return c.GetPointString()
	}
}

func (c Poker) GetPointString() string {
	if c.Point == PA {
		return "A"
	} else if c.Point == PJ {
		return "J"
	} else if c.Point == PQ {
		return "Q"
	} else if c.Point == PK {
		return "K"
	} else if c.Point == PBlackJoker {
		return "小王"
	} else if c.Point == PRedJoker {
		return "大王"
	} else {
		return strconv.Itoa(int(c.Point))
	}
}

func (c Poker) ToInt() uint32 {
	return c.Suit + c.Point
}

func (c Poker) Equals(other Poker) bool {
	return c.Suit == other.Suit && c.Point == other.Point
}

// 带花色比较，黑桃A 和 方块A比较返回true
func (c Poker) BiggerThan(other Poker) bool {
	return c.Weight > other.Weight
}

// 无花色比较，黑桃A 和 方块A比较返回false
func (c Poker) PointBiggerThan(other Poker) bool {
	return c.PointWeight > other.PointWeight
}

type PokerSlice []Poker

func (cs PokerSlice) Len() int           { return len(cs) }
func (cs PokerSlice) Swap(i, j int)      { cs[i], cs[j] = cs[j], cs[i] }
func (cs PokerSlice) Less(i, j int) bool { return cs[i].Weight < cs[j].Weight }

func ToPoker(card uint32) Poker {
	result := Poker{}
	result.Suit = card / 16 * 16
	result.Point = card % 16

	// 计算无花色权重
	if result.Point == PA {
		result.PointWeight = PK + PA //A为K加1
	} else if result.Point == P2 {
		result.PointWeight = PK + P2 + 1 //2为A加1,方便断开顺子,连对等
	} else if result.Point == PBlackJoker || result.Point == PRedJoker {
		result.PointWeight = SSpade + PK + result.Point //大小王，加大权重
	} else {
		result.PointWeight = result.Point
	}
	result.Weight = result.PointWeight*5 + result.Suit/16 //带花色权重
	return result
}

func ToPokers(cards []uint32) []Poker {
	result := make([]Poker, 0, len(cards))
	for _, card := range cards {
		result = append(result, ToPoker(card))
	}
	return result
}

func ToInts(cards []Poker) []uint32 {
	result := make([]uint32, 0, len(cards))
	for _, card := range cards {
		result = append(result, card.ToInt())
	}
	return result
}

func PokerSort(cards []Poker) {
	cs := PokerSlice(cards)
	sort.Sort(cs)
}

func PokerSortDesc(cards []Poker) {
	cs := PokerSlice(cards)
	sort.Sort(sort.Reverse(cs))
}
