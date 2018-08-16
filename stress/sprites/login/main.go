package login

import (
	"steve/stress/stressclient/sprite"
	"steve/simulate/utils"
	"steve/client_pb/msgid"
	"steve/client_pb/gate"
	"steve/simulate/global"
	"github.com/Sirupsen/logrus"
	"time"
)

type login struct {
	ch chan int
}

// GetSprite 获取服务接口，被 stressclient 调用
func GetSprite() sprite.Sprite {
	return login{}
}

func (s login) Init() error {
	s.ch = make(chan int)
	return nil
}
func (s login) Start() error {
	stage := sprite.GetStage()
	start := time.Now()
	player, _ := utils.LoginNewPlayer()

	if player == nil {
		return nil
	}
	elapsed := time.Since(start)
	stage.LoginHis.Observe(elapsed.Seconds())
	player.AddExpectors(msgid.MsgID_GATE_HEART_BEAT_RSP)
	client := player.GetClient()
	client.SendPackage(utils.CreateMsgHead(msgid.MsgID_GATE_HEART_BEAT_REQ), &gate.GateHeartBeatReq{})

	expector := player.GetExpector(msgid.MsgID_GATE_HEART_BEAT_RSP)
	err := expector.Recv(global.DefaultWaitMessageTime, nil)

	logrus.Info(err)
	select {
		case <-s.ch:
			return  nil
	}
}
func (s login) Stop() error {
	s.ch <- 0
	return nil
}

func main() {}
