package config

import (
	"strconv"
	"strings"
)

// RoleConfig 角色配置
type RoleConfig struct {
	ProduceID int    `json:"produceId"`
	Coins     int    `json:"coins"`
	KeyCards  int    `json:"keyCards"`
	Ingots    int    `json:"ingots"`
	Items     string `json:"items"`
	ItemArr   [][]int
}

// InitItem 初始化item
func (config *RoleConfig) InitItem() {
	idAndNum := strings.Split(config.Items, ";")
	config.ItemArr = make([][]int, len(idAndNum))
	for index, val := range idAndNum {
		iN := strings.Split(val, "|")
		id, _ := strconv.Atoi(iN[0])
		num, _ := strconv.Atoi(iN[1])
		config.ItemArr[index] = make([]int, 2)
		config.ItemArr[index][0] = id
		config.ItemArr[index][1] = num
	}
}
