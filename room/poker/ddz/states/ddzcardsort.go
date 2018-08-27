package states

import (
	"sort"
	. "steve/room/poker"

	"github.com/Sirupsen/logrus"
)

// DDZSort 从小到大排序后返回
func DDZSort(cards []uint32) []uint32 {
	return ddzSort(cards, false)
}

// DDZSortDescend 从大到小排序后返回
func DDZSortDescend(cards []uint32) []uint32 {
	return ddzSort(cards, true)
}

func ddzSort(cards []uint32, reverse bool) []uint32 {
	cs := PokerSlice(ToPokers(cards))
	if reverse {
		sort.Sort(sort.Reverse(cs))
	} else {
		sort.Sort(cs)
	}
	result := make([]uint32, 0, cs.Len())
	for i := range cs {
		result = append(result, cs[i].ToInt())
	}
	logrus.WithFields(logrus.Fields{"in": cards, "out:": result}).Debug("斗地主排序")
	return result
}
