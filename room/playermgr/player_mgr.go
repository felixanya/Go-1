package playermgr

import (
	"steve/room/interfaces"
	"steve/room/interfaces/global"
	"sync"

	"github.com/Sirupsen/logrus"
)

type playerMgr struct {
	playerMap   sync.Map // playerID -> player
	clientMap   sync.Map // clientID -> playerID
	userNameMap sync.Map // userName-> playerID

	mu sync.RWMutex
}

func (pm *playerMgr) AddPlayer(p interfaces.Player) {
	pm.mu.Lock()
	defer pm.mu.Unlock()
	pm.playerMap.Store(p.GetID(), p)
	pm.clientMap.Store(p.GetClientID(), p.GetID())
	pm.userNameMap.Store(p.GetUserName(), p.GetID())
}

func (pm *playerMgr) GetPlayer(playerID uint64) interfaces.Player {
	return pm.getPlayer(playerID)
}

func (pm *playerMgr) getPlayer(playerID uint64) interfaces.Player {
	v, ok := pm.playerMap.Load(playerID)
	if !ok {
		return nil
	}
	return v.(interfaces.Player)
}

func (pm *playerMgr) OnClientDisconnect(clientID uint64) {
	player := pm.GetPlayerByClientID(clientID)
	pm.clientMap.Delete(clientID)
	if player != nil {
		player.SetClientID(0)
	}
}

func (pm *playerMgr) GetPlayerByClientID(clientID uint64) interfaces.Player {
	pm.mu.RLock()
	defer pm.mu.RUnlock()
	v, ok := pm.clientMap.Load(clientID)
	if !ok {
		return nil
	}
	playerID := v.(uint64)
	return pm.getPlayer(playerID)
}

func (pm *playerMgr) GetPlayerByUserName(userName string) interfaces.Player {
	pm.mu.RLock()
	defer pm.mu.RUnlock()
	v, ok := pm.userNameMap.Load(userName)
	if !ok {
		return nil
	}
	playerID := v.(uint64)
	return pm.getPlayer(playerID)
}

func (pm *playerMgr) UpdatePlayerClientID(playerID uint64, clientID uint64) {
	pm.mu.Lock()
	defer pm.mu.Unlock()
	player := pm.getPlayer(playerID)
	if player == nil {
		return
	}
	oldClientID := player.GetClientID()
	pm.clientMap.Delete(oldClientID)
	pm.clientMap.Store(clientID, playerID)
	player.SetClientID(clientID)
}

var setupOnce = sync.Once{}

func init() {
	global.SetPlayerMgr(&playerMgr{})
	logrus.Debugln("设置用户管理器")
}
