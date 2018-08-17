package config

import (
	"strings"
	"strconv"
)

type RoleInitConfig struct {
	GameId  int    `json:"gameId"`
	Channel int    `json:"channel"`
	GoldNum int    `json:"goldNum"`
	CardNum int    `json:"cardNum"`
	YbNum   int    `json:"ybNum"`
	Item    string `json:"item"`
	ItemArr [][]int
}

func (config *RoleInitConfig) InitItem() {
	idAndNum := strings.Split(config.Item, ";")
	config.ItemArr = make([][]int,len(idAndNum))
	for index, val := range idAndNum {
		iN := strings.Split(val,"|")
		id,_ := strconv.Atoi(iN[0])
		num,_ := strconv.Atoi(iN[1])
		config.ItemArr[index] = make([]int,2)
		config.ItemArr[index][0] = id
		config.ItemArr[index][1] = num
	}
}
