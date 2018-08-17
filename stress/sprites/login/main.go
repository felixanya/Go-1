package login

import (
	"github.com/Sirupsen/logrus"
	"steve/client_pb/gate"
	"steve/client_pb/msgid"
	"steve/simulate/global"
	"steve/simulate/utils"
	"time"
	"steve/simulate/config"
	"steve/stress/stressclient/core"
)

type login struct {
	ch chan struct{}
}

// GetSprite 获取服务接口，被 stressclient 调用
func GetSprite() core.Sprite {
	return &login{}
}

func (s *login) Init() error {
	s.ch = make(chan struct{})
	return nil
}
func (s *login) Start(params []string) error {
	gatewayServerAddr := params[0]
	config.SetGatewayServerAddr(gatewayServerAddr)
	start := time.Now()
	player, _ := utils.LoginNewPlayer()

	if player == nil {
		core.AddError("login")
		return nil
	}
	elapsed := time.Since(start)
	core.Observe("login", elapsed.Seconds())
	player.AddExpectors(msgid.MsgID_GATE_HEART_BEAT_RSP)
	client := player.GetClient()
	client.SendPackage(utils.CreateMsgHead(msgid.MsgID_GATE_HEART_BEAT_REQ), &gate.GateHeartBeatReq{})

	expector := player.GetExpector(msgid.MsgID_GATE_HEART_BEAT_RSP)
	err := expector.Recv(global.DefaultWaitMessageTime, nil)

	logrus.Info(err)
	return nil
	select {
	case <-s.ch:
		return nil
	}
}
func (s *login) Stop() error {
	close(s.ch)
	return nil
}

func main() {}
