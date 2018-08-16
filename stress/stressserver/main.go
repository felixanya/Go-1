package main

import (
	"fmt"
	"steve/stress/stressserver/core"
	"steve/stress/common"
)

func main() {
	fmt.Printf("start\n")
	common.Init()
	core.Init()
}
