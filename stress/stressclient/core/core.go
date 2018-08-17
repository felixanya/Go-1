package core

import (
	"github.com/Sirupsen/logrus"
	"time"
	"steve/stress/common"
	"steve/stress/proto"
	"encoding/json"
	"strconv"
	"runtime/debug"
)

func init() {
	common.Init()
	initSys()
	go initPrometheus()
	go recv()
}

//Sprite 压测逻辑
type Sprite interface {
	Init() error
	Start(params []string) error
	Stop() error
}

//控制启动次数的定时器
var ticker *time.Ticker
var exitTicker chan struct{}
//使用的压测逻辑
var sp Sprite
//从web界面传入的附加参数
var args []string

var getSprite func(name string) (Sprite, error)
func SetGetSpriteFunc(f func(name string) (Sprite, error)) {
	getSprite = f
}

func doServerCommand(in *client.ServerCommand) {
	var jsonParam map[string][]string
	var params []string
	if in.Params != "" {
		err := json.Unmarshal([]byte(in.Params), &jsonParam)
		if err != nil {
			logrus.Error("params ", in.Params, err)
		}else {
			params = jsonParam["params"]
		}
	}
	switch in.Cmd {
	case 0: //connected
		logrus.Info("connected")
	case 1: //start
		times, _ := strconv.Atoi(params[0])
		num, _ := strconv.Atoi(params[1])
		interval, _ := strconv.Atoi(params[2])
		fun := params[3]
		args = params[4:]
		go startStress(times, num, interval, fun)
	case 2: //stop
		stopStress()
	}

}

func startStress(times int, num int, interval int, fun string) {
	du := time.Duration(interval) * time.Second
	exitTicker = make(chan struct{})
	ticker = time.NewTicker(du)
	currentTimes := 0
	total := num
	name := fun
	s, err := getSprite(name)
	if err != nil {
		logrus.Info(name, err.Error())
		return
	}
	sp = s
	sp.Init()
	goStart(s, total)
	if times < 2 {
		logrus.Info("ticker over")
		return
	}
	tickerFor:
	for {
		select {
		case <-ticker.C:
			if currentTimes < times - 1 {
				goStart(s, total)
			}else{
				ticker.Stop()
				break tickerFor
			}
			currentTimes++
		case <-exitTicker:
			break tickerFor
		}
	}
	logrus.Info("ticker over")
}

func goStart(s Sprite, total int) {
	logrus.Info("goStart")
	for i := 0; i < total; i++ {
		go start(s)
	}
}
func start(s Sprite) {
	start := time.Now()
	defer recoverPanic(start)
	metrics.ConnectsCounter.WithLabelValues("started").Add(1)
	s.Start(args)
}

func recoverPanic(start time.Time) {
	elapsed := time.Since(start)
	Observe("connect", elapsed.Seconds())
	if x := recover(); x != nil {
		metrics.ConnectsCounter.WithLabelValues("exited").Add(1)
		AddError("panic")
		stack := debug.Stack()
		logrus.Errorln(string(stack))
	}
}

func stopStress() {
	if ticker != nil {
		ticker.Stop()
	}
	if sp != nil {
		sp.Stop()
	}
	close(exitTicker)
	metrics.ConnectsCounter.Reset()
	metrics.histogram.Reset()
}
