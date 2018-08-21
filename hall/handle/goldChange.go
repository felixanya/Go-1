package handle

import (
	"encoding/json"
	"fmt"
	pb_common "steve/client_pb/common"
	pb_hall "steve/client_pb/hall"
	"steve/client_pb/msgid"
	"steve/gutils"
	"steve/server_pb/gold"

	"github.com/Sirupsen/logrus"
	"github.com/golang/protobuf/proto"
	nsq "github.com/nsqio/go-nsq"
)

type GoldChanngleHandler struct {
}

func (plh *GoldChanngleHandler) HandleMessage(message *nsq.Message) error {
	entry := logrus.WithFields(logrus.Fields{
		"func_name": "HandleMessage",
		"desc":      "处理玩家金币变化通知",
		"message":   message,
	})
	goldChangeNtf := &gold.GoldChangeNtf{}
	if err := json.Unmarshal(message.Body, &goldChangeNtf); err != nil {
		return fmt.Errorf("消息反序列化失败：%v", err)
	}
	sendMoneyChangeNtf(goldChangeNtf)
	entry.Debugf("处理玩家金币变化通知完成")
	return nil
}

// sendMoneyChangeNtf 通知客户端金币变化
func sendMoneyChangeNtf(ntf *gold.GoldChangeNtf) {
	// 玩家ID
	playerID := ntf.GetUid()
	// 玩家当前金币数
	money := uint64(ntf.GetAfterGold())

	// 下发通知
	moneyChangeNtf := pb_hall.MoneyChangeNtf{
		PlayerId: proto.Uint64(playerID),
		Money: &pb_common.Money{
			MoneyType: pb_common.MoneyType(ntf.GoldType).Enum(),
			MoneyNum:  proto.Uint64(money),
		},
	}

	gutils.SendMessage(playerID, msgid.MsgID_MONEY_CHANGE_NTF, &moneyChangeNtf)
}
