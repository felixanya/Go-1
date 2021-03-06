package player

import (
	"steve/client_pb/msgid"
	"steve/client_pb/room"
	"steve/external/goldclient"
	"steve/external/hallclient"
	"steve/room/desk"
	"steve/room/util"
	"steve/server_pb/gold"
	"sync"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/golang/protobuf/proto"
)

type Player struct {
	PlayerID     uint64
	seat         uint32        // 座号
	ecoin        uint64        // 进牌桌金币数
	quit         bool          // 是否已经退出牌桌
	tuoguan      bool          // 是否在托管中
	autoHu       bool          // 是否自动胡牌
	CountingDown bool          // 开始补时倒计时
	AddTime      time.Duration // 补时时间,毫秒为单位
	robotLv      int           // 机器人等级
	brokerCount  int           // 破产次数
	desk         *desk.Desk

	mu sync.RWMutex
}

// GetDesk 获取玩家所在牌桌
func (dp *Player) GetDesk() *desk.Desk {
	return dp.desk
}

// SetDesk 设置玩家所在牌桌
func (dp *Player) SetDesk(deskObj *desk.Desk) {
	dp.mu.Lock()
	dp.desk = deskObj
	logrus.Debugln("设置 desk")
	dp.mu.Unlock()
}

// SetQuit 设置玩家退出状态
func (dp *Player) SetQuit(quit bool) {
	dp.mu.Lock()
	dp.quit = quit
	dp.mu.Unlock()
}

// GetPlayerID 获取玩家 ID
func (dp *Player) GetPlayerID() uint64 {
	dp.mu.RLock()
	defer dp.mu.RUnlock()
	return dp.PlayerID
}

// GetSeat 获取座号
func (dp *Player) GetSeat() int {
	dp.mu.RLock()
	defer dp.mu.RUnlock()
	return int(dp.seat)
}

func (dp *Player) SetSeat(seat uint32) {
	dp.mu.RLock()
	defer dp.mu.RUnlock()
	dp.seat = seat
}

// GetEcoin 获取进入时金币数
func (dp *Player) GetEcoin() uint64 {
	return dp.ecoin
}

func (dp *Player) SetEcoin(coin uint64) {
	dp.ecoin = coin
}

func (p *Player) SetRobotLv(lv int) {
	p.mu.RLock()
	defer p.mu.RUnlock()
	p.robotLv = lv
}

func (p *Player) GetRobotLv() int {
	return p.robotLv
}

// IsQuit 是否已经退出
func (dp *Player) IsQuit() bool {
	return dp.quit
}

// IsTuoguan 玩家是否在托管中
func (dp *Player) IsTuoguan() bool {
	dp.mu.RLock()
	defer dp.mu.RUnlock()
	return dp.tuoguan
}

// SetTuoguan 设置托管
func (dp *Player) SetTuoguan(tuoguan bool, notify bool) {
	dp.mu.Lock()
	defer dp.mu.Unlock()
	toggle := dp.tuoguan != tuoguan
	dp.tuoguan = tuoguan
	if toggle && notify {
		dp.notifyTuoguan(dp.PlayerID, tuoguan)
	}
}

// IsAutoHu 玩家是否自动胡牌
func (dp *Player) IsAutoHu() bool {
	dp.mu.RLock()
	defer dp.mu.RUnlock()
	return dp.autoHu
}

// SetAutoHu 设置自动胡牌
func (dp *Player) SetAutoHu(autoHu bool) {
	dp.mu.Lock()
	defer dp.mu.Unlock()
	dp.autoHu = autoHu
}

func (p *Player) GetCoin() uint64 {
	coin, err := goldclient.GetGold(p.PlayerID, int16(gold.GoldType_GOLD_COIN))
	if err != nil {
		return 0
	}
	return uint64(coin)
}

// IsOnline 判断玩家是否在线
func (p *Player) IsOnline() bool {
	online, _ := hallclient.GetGateAddr(p.PlayerID)
	return online != ""
}

func (dp *Player) notifyTuoguan(playerID uint64, tuoguan bool) {
	util.SendMessageToPlayer(playerID, msgid.MsgID_ROOM_TUOGUAN_NTF, &room.RoomTuoGuanNtf{
		Tuoguan: proto.Bool(tuoguan),
	})
	logrus.WithFields(logrus.Fields{"player_id": playerID, "tuoguan": tuoguan}).Debugln("通知托管")
}

// AddBrokerCount 增加破产次数
func (dp *Player) AddBrokerCount() {
	dp.mu.Lock()
	dp.brokerCount++
	dp.mu.Unlock()
}

// GetBrokerCount 获取破产次数
func (dp *Player) GetBrokerCount() int {
	dp.mu.RLock()
	defer dp.mu.RUnlock()
	return dp.brokerCount
}
