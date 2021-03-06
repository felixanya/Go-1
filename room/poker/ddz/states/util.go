package states

import (
	"fmt"
	"steve/client_pb/msgid"
	"steve/entity/majong"
	"steve/entity/poker/ddz"
	"steve/room/poker/ddz/ddzmachine"
	"steve/room/poker/machine"
	"time"

	"math/rand"
	"steve/client_pb/room"
	"strconv"

	"github.com/Sirupsen/logrus"
	"github.com/golang/protobuf/proto"
)

// PeiPai 配牌工具
func PeiPai(wallCards []uint32, value string) error {
	var cards []uint32
	for i := 0; i < len(value); i = i + 3 {
		card, err := strconv.Atoi(value[i : i+2])
		if err != nil {
			return err
		}
		cards = append(cards, uint32(card))
	}
	for i := 0; i < len(cards); i++ {
		for j := len(wallCards) - 1; j >= 0; j-- {
			if cards[i] == wallCards[j] {
				wallCards[i], wallCards[j] = wallCards[j], wallCards[i]
				break
			}
		}
	}
	logrus.WithFields(logrus.Fields{"wallCards": wallCards, "peipai:": value}).Debug("斗地主配牌成功")
	return nil
}

// getDDZContext 从状态机中获取斗地主现场
func getDDZContext(m machine.Machine) *ddz.DDZContext {
	dm, ok := m.(*ddzmachine.DDZMachine)
	if !ok {
		return nil
	}
	return dm.GetDDZContext()
}

func remove(playerIds []uint64, removeId uint64) []uint64 {
	var result []uint64
	for _, playerId := range playerIds {
		if playerId != removeId {
			result = append(result, playerId)
		}
	}
	return result
}

func getPlayerIds(m machine.Machine) []uint64 {
	dm, ok := m.(*ddzmachine.DDZMachine)
	if !ok {
		return nil
	}

	var players []uint64
	for _, player := range dm.GetDDZContext().GetPlayers() {
		players = append(players, player.GetPlayerId())
	}
	return players
}

func IsValidPlayer(context *ddz.DDZContext, id uint64) bool {
	return GetPlayerByID(context.GetPlayers(), id) != nil
}

func getRandPlayerId(players []*ddz.Player) uint64 {
	i := rand.Intn(len(players))
	playerId := players[i].PlayerId
	return playerId
}

//GetPlayerByID 根据玩家id获取玩家
func GetPlayerByID(players []*ddz.Player, id uint64) *ddz.Player {
	for _, player := range players {
		if player.PlayerId == id {
			return player
		}
	}
	return nil
}

//GetNextPlayerByID 根据玩家id获取下个玩家
func GetNextPlayerByID(players []*ddz.Player, id uint64) *ddz.Player {
	for k, player := range players {
		if player.PlayerId == id {
			index := (k + 1) % len(players)
			return players[index]
		}
	}
	return nil
}

// sendMessage 向玩家发送消息
func sendMessage(m machine.Machine, players []uint64, msgID msgid.MsgID, body proto.Message) error {
	dm, ok := m.(*ddzmachine.DDZMachine)
	if !ok {
		return fmt.Errorf("不是斗地主状态机")
	}
	return dm.SendMessage(players, msgID, body)
}

func sendToPlayer(m machine.Machine, playerID uint64, msgID msgid.MsgID, body proto.Message) error {
	logrus.WithFields(logrus.Fields{"playerID": playerID, "msgId": msgID, "msg": body}).Debug("斗地主玩家响应")
	return sendMessage(m, []uint64{playerID}, msgID, body)
}

func broadcast(m machine.Machine, msgID msgid.MsgID, body proto.Message) error {
	playerIDs := getPlayerIds(m)
	logrus.WithFields(logrus.Fields{"playerIDs": playerIDs, "msgId": msgID, "msg": body}).Debug("斗地主广播")
	return sendMessage(m, playerIDs, msgID, body)
}

func broadcastExcept(m machine.Machine, playerID uint64, msgID msgid.MsgID, body proto.Message) error {
	allPlayers := getPlayerIds(m)
	players := []uint64{}
	for _, pid := range allPlayers {
		if pid != playerID {
			players = append(players, pid)
		}
	}
	return sendMessage(m, players, msgID, body)
}

// setMachineAutoEvent 设置状态机自动事件
func setMachineAutoEvent(m machine.Machine, event machine.Event, duration time.Duration) {
	dm, ok := m.(*ddzmachine.DDZMachine)
	if !ok {
		return
	}
	dm.SetAutoEvent(event, duration)
}

func If(judge bool, trueReturn interface{}, falseReturn interface{}) interface{} {
	if judge {
		return trueReturn
	} else {
		return falseReturn
	}
}

// OnCartoonFinish TODO: 和麻将统一
func OnCartoonFinish(curState int, nextState int, needCartoonType room.CartoonType, eventContext interface{}) (newState int, err error) {
	logEntry := logrus.WithFields(logrus.Fields{
		"func_name":         "OnCartoonFinish",
		"cur_state":         curState,
		"next_state":        nextState,
		"need_cartoon_type": needCartoonType,
	})

	req := eventContext.(*majong.CartoonFinishRequestEvent)
	reqCartoonType := req.GetCartoonType()
	logEntry.WithField("req_cartoon_type", reqCartoonType).Debugln("收到动画完成请求")
	if reqCartoonType != int32(needCartoonType) {
		return curState, nil
	}
	return nextState, nil
}
