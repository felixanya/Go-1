package packsack

import (
	"fmt"
	"steve/alms/packsack/packsack_gold"
	"steve/alms/packsack/packsack_prop"
	"steve/client_pb/alms"
	"steve/client_pb/common"
	"steve/client_pb/msgid"
	"steve/external/gateclient"
	"steve/server_pb/user"

	"github.com/Sirupsen/logrus"
	"github.com/golang/protobuf/proto"
	nsq "github.com/nsqio/go-nsq"
)

func init() {
	// exposer := structs.GetGlobalExposer()
	// if err := exposer.Subscriber.Subscribe(constant.PlayerLogin, "alms", &PlayerPacksackInfo{}); err != nil {
	// 	logrus.WithError(err).Panicln("订阅登录消息失败")
	// }
}

type PlayerPacksackInfo struct {
}

func (ppi *PlayerPacksackInfo) HandleMessage(message *nsq.Message) error {
	logrus.Debugln("玩家登陆订阅背包信息")
	loginPb := user.PlayerLogin{}
	if err := proto.Unmarshal(message.Body, &loginPb); err != nil {
		logrus.WithError(err).Errorln("消息反序列化失败")
		return fmt.Errorf("消息反序列化失败：%v", err)
	}
	if err := PlayerLoginBacksackInfo(loginPb.PlayerId); err != nil {
		logrus.WithError(err).Panicln("发送玩家背包信息失败")
	}
	return nil
}

// type packsackInof struct {
// 	propinfo []packsack_prop.PropPacksackInfo
// 	gold     int64
// }

//PlayerLoginBacksackInfo 玩家登陆获取背包信息
func PlayerLoginBacksackInfo(playerID uint64) error {
	//1.获取所有道具信息
	//2.获取金币信息
	gold, err := packsack_gold.GetGoldMgr().GetGold(playerID)
	if err != nil {
		return fmt.Errorf("获取背包金币 err(%v)", err)
	}
	propInfos, err := packsack_prop.GetPlayerPropInfoAll(playerID)
	if err != nil {
		return fmt.Errorf("获取所有道具信息 err(%v)", err)
	}

	ntfPropInfos := make([]*alms.PacksackPropInfo, 0)
	for _, prop := range propInfos {
		propType := common.PropType(1)
		ntfPropInfo := &alms.PacksackPropInfo{
			PropId:       proto.Int32(prop.PropID),
			PropName:     proto.String(prop.PropName),
			PropDescribe: proto.String(prop.Describe),
			PropCount:    proto.Int64(prop.PropCount),
			PropType:     &propType,
		}
		ntfPropInfos = append(ntfPropInfos, ntfPropInfo)
	}

	pkinfoNtf := &alms.PlayerPacksackInfoNtf{
		PacksackGold: proto.Int64(gold),
		PropInfo:     ntfPropInfos,
	}
	return gateclient.SendPackageByPlayerID(playerID, uint32(msgid.MsgID_PACKSACK_LOGIN_INFO_NTF), pkinfoNtf)
}
