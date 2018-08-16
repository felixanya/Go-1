package core

import (
	"github.com/Sirupsen/logrus"
	"steve/stress/stressclient/sprite"
	"time"
	"steve/stress/common"
	"steve/stress/proto"
	"encoding/json"
	"strconv"
	"runtime/debug"
)

func init() {
	common.Init()
	sprite.InitStage()
	initSys()
	go initPrometheus()
	go recv()
}

var ticker *time.Ticker
var sp sprite.Sprite
var getSprite func(name string) (sprite.Sprite, error)

func SetGetSpriteFunc(f func(name string) (sprite.Sprite, error)) {
	getSprite = f
}

func doServerCommand(in *client.ServerCommand) {
	var jsonparam map[string][]string
	var params []string
	if in.Params != "" {
		err := json.Unmarshal([]byte(in.Params), &jsonparam)
		if err != nil {
			logrus.Error("params ", in.Params, err)
		}else {
			params = jsonparam["params"]
		}
	}
	switch (in.Cmd) {
	case 0: //connected
		logrus.Info("connected")
	case 1: //start
		times, _ := strconv.Atoi(params[0])
		num, _ := strconv.Atoi(params[1])
		interval, _ := strconv.Atoi(params[2])
		fun := params[3]
		StartStress(times, num, interval, fun)
	case 2: //stop
		StopStress()
	}

}

func StartStress(times int, num int, interval int, fun string) {
	ticker = time.NewTicker(time.Duration(interval) * time.Second)
	currentTimes := 0
	total := num
	name := fun
	s, err := getSprite(name)
	if err != nil {
		logrus.Info(name, err.Error())
		return
	}
	sp = s
	s.Init()
	select {
	case <-ticker.C:
		gostart(s, total)
		currentTimes++
		if currentTimes >= times {
			ticker.Stop()
		}
	}
}

func gostart(s sprite.Sprite, total int) {
	for i := 0; i < total; i++ {
		go start(s)
	}
}
func start(s sprite.Sprite) {
	defer recoverPanic()
	s.Start()
}

func recoverPanic() {
	if x := recover(); x != nil {
		stack := debug.Stack()
		logrus.Errorln(string(stack))
	}
}

func StopStress() {
	if ticker != nil {
		ticker.Stop()
	}
	if sp != nil {
		sp.Stop()
	}
}
