package global

import (
	"steve/client_pb/common"
	"steve/client_pb/room"
	"steve/gutils"
	"steve/simulate/structs"
)

func makeRoomCards(card ...room.Card) []*room.Card {
	return gutils.MakeRoomCards(card...)
}

// NewCommonStartGameParams 创建通用启动参数
func NewCommonStartGameParams() structs.StartGameParams {
	return structs.StartGameParams{
		GameID: common.GameId_GAMEID_XUELIU,
		Cards: [][]uint32{
			{11, 11, 11, 11, 12, 12, 12, 12, 13, 13, 13, 13, 14, 14},
			{15, 15, 15, 15, 16, 16, 16, 16, 17, 17, 17, 17, 18},
			{21, 21, 21, 21, 22, 22, 22, 22, 23, 23, 23, 23, 24},
			{25, 25, 25, 25, 26, 26, 26, 26, 27, 27, 27, 27, 28},
		},
		WallCards:  []uint32{31},
		HszDir:     room.Direction_AntiClockWise,
		BankerSeat: 0,

		IsDq:  true,
		IsHsz: true,
		HszCards: [][]uint32{
			{11, 11, 11},
			{15, 15, 15},
			{21, 21, 21},
			{25, 25, 25},
		},
		PlayerNum:      4,
		DingqueColor:   []room.CardColor{room.CardColor_CC_TIAO, room.CardColor_CC_TIAO, room.CardColor_CC_TIAO, room.CardColor_CC_TIAO},
		PlayerSeatGold: map[int]uint64{0: 100000, 1: 100000, 2: 100000, 3: 100000},
		PeiPaiGame:     "scxl",
		DiFen:          100,
	}
}
