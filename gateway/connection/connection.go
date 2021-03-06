package connection

import (
	"context"
	"fmt"
	"net"
	"steve/external/hallclient"
	"steve/gateway/config"
	user_pb "steve/server_pb/user"
	"strings"
	"time"

	"github.com/Sirupsen/logrus"
)

const (
	// 多长时间没有检测到心跳断开连接
	heartBeatInterval time.Duration = time.Minute
	// 多长时间没有认证断开连接
	attachInterval time.Duration = time.Minute * 10
)

// Connection 连接
type Connection struct {
	clientID uint64
	connMgr  *ConnMgr
	addr     net.Addr

	playerID       uint64
	heartBeatTimer *time.Timer
	attachTimer    *time.Timer
}

func (c *Connection) run(ctx context.Context, finish func()) {
	c.heartBeatTimer = time.NewTimer(heartBeatInterval)
	c.attachTimer = time.NewTimer(attachInterval)

	go func() {
		defer c.heartBeatTimer.Stop()
		defer c.attachTimer.Stop()
		select {
		case <-ctx.Done():
			{
				c.DetachPlayer()
				return
			}
		case <-c.heartBeatTimer.C:
			{
				c.kick("无心跳", finish)
			}
		case <-c.attachTimer.C:
			{
				c.kick("未认证", finish)
			}
		}
	}()
}

// DetachPlayer 解除和 player 的绑定
func (c *Connection) DetachPlayer() {
	logrus.Debugf("DetachPlayer 解除和 player 的绑定,playerId:%v", c.playerID)
	if c.playerID != 0 {
		hallclient.UpdatePlayeServerAddr(c.playerID, user_pb.ServerType_ST_GATE, "")
		c.connMgr.setPlayerConnectionID(c.playerID, 0)
		c.playerID = 0
	}
}

func (c *Connection) kick(reason string, finish func()) {
	entry := logrus.WithFields(logrus.Fields{
		"player_id": c.playerID,
		"client_id": c.clientID,
		"reason":    reason,
	})
	entry.Infoln("踢出玩家")
	c.DetachPlayer()
	finish()
}

// GetPlayerID 获取绑定的玩家 ID
func (c *Connection) GetPlayerID() uint64 {
	return c.playerID
}

// AttachPlayer 绑定玩家 ID
func (c *Connection) AttachPlayer(playerID uint64) bool {
	entry := logrus.WithFields(logrus.Fields{
		"func_name":        "connection.AttachPlayer",
		"player_id":        c.playerID,
		"client_id":        c.clientID,
		"attach_player_id": playerID,
	})
	if c.playerID != 0 {
		entry.Infoln("已绑定")
		return false
	}
	c.playerID = playerID
	c.attachTimer.Stop()

	addr := c.addr.String()
	splitIndex := strings.Index(addr, ":")
	ip := addr
	if splitIndex != -1 {
		ip = addr[:splitIndex]
	}
	entry = entry.WithField("ip", ip)
	if succ, err := hallclient.UpdatePlayerGateInfo(c.playerID, ip, c.getGatewayAddr()); !succ || err != nil {
		entry.WithError(err).Errorln("更新玩家网关地址失败")
	} else {
		entry.Debugln("更新玩家网关地址")
	}
	c.connMgr.setPlayerConnectionID(c.playerID, c.clientID)
	return true
}

// GetClientID 获取连接 ID
func (c *Connection) GetClientID() uint64 {
	return c.clientID
}

// HeartBeat 心跳
func (c *Connection) HeartBeat() {
	c.heartBeatTimer.Reset(heartBeatInterval)
}

func (c *Connection) getGatewayAddr() string {
	return fmt.Sprintf("%s:%d", config.GetRPCAddr(), config.GetRPCPort())
}
