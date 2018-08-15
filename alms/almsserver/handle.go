package almsserver

import (
	"steve/alms/packsack/packsack_gold"
	"steve/alms/packsack/packsack_prop"
	client_alms "steve/client_pb/alms"
	"steve/client_pb/msgid"
	"steve/external/goldclient"
	"steve/server_pb/gold"
	"steve/structs/exchanger"
	"steve/structs/proto/gate_rpc"

	"github.com/Sirupsen/logrus"
	"github.com/golang/protobuf/proto"
)

// HandlePacksackInfo 获取背包信息
func HandlePacksackInfo(playerID uint64, header *steve_proto_gaterpc.Header, req client_alms.PlayerPacksackInfoRep) (rspMsg []exchanger.ResponseMsg) {
	entry := logrus.WithFields(logrus.Fields{
		"func_name": "HandlePacksackInfo",
		"request":   req,
		"playerID":  playerID,
	})
	// 返回消息
	response := &client_alms.PlayerPacksackInfoRsp{}
	rspMsg = []exchanger.ResponseMsg{
		exchanger.ResponseMsg{
			MsgID: uint32(msgid.MsgID_PACKSACK_INFO_RSP),
			Body:  response,
		},
	}
	gold, err := packsack_gold.GetGoldMgr().GetGold(playerID)
	if err != nil {
		logrus.WithError(err).Debugln("获取背包金币 ERR")
		return
	}
	propInfos, err := packsack_prop.GetPlayerPropInfoAll(playerID)
	if err != nil {
		logrus.WithError(err).Debugln("获取所有道具信息 ERR")
		return
	}
	ntfPropInfos := make([]*client_alms.PacksackPropInfo, 0)
	for _, prop := range propInfos {
		pt := client_alms.PropType_INTERACTIVE //互动道具
		ntfPropInfo := &client_alms.PacksackPropInfo{
			PropId:       proto.Int32(prop.PropID),
			PropName:     proto.String(prop.PropName),
			PropDescribe: proto.String(prop.Describe),
			PropCount:    proto.Int64(prop.PropCount),
			PropType:     &pt,
		}
		ntfPropInfos = append(ntfPropInfos, ntfPropInfo)
	}
	response.PacksackGold = proto.Int64(gold)
	response.PropInfo = ntfPropInfos
	entry.Debugln("获取背包信息成功")
	return
}

// HandlePackSackGold 处理背包金币请求
func HandlePackSackGold(playerID uint64, header *steve_proto_gaterpc.Header, req client_alms.PacksackGoldReq) (rspMsg []exchanger.ResponseMsg) {
	entry := logrus.WithFields(logrus.Fields{
		"func_name":  "HandlePackSackGold",
		"changeGold": req.GetChangeGold(),
		"playerID":   playerID,
	})
	response := &client_alms.PacksackGoldRsp{}
	rspMsg = []exchanger.ResponseMsg{
		exchanger.ResponseMsg{
			MsgID: uint32(msgid.MsgID_PACKSACK_GOLD_RSP),
			Body:  response,
		},
	}
	changeGold := req.GetChangeGold()
	pkgold, err := packsack_gold.GetGoldMgr().AddGold(playerID, changeGold)
	if err != nil {
		entry.WithError(err).Debugln("背包金币修改失败")
		return
	}
	// 从金币服获取
	// 更改玩家身上的金币 TODO almsFuncID 渠道ID
	almsFuncID := int32(11)
	gold, err := goldclient.AddGold(playerID, int16(gold.GoldType_GOLD_COIN), -changeGold, almsFuncID, 0, 0, 0)
	if err != nil {
		entry.WithError(err).Debugln("金币服修改金币失败")
		return
	}
	response.Gold = proto.Int64(gold)
	response.PacksackGold = proto.Int64(pkgold)
	entry.WithFields(logrus.Fields{
		"Gold":   gold,
		"pkgold": pkgold,
	}).Debugln("处理背包金币请求成功")
	return rspMsg
}
