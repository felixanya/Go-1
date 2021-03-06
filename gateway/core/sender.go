package core

import (
	"context"
	"steve/client_pb/msgid"
	"steve/gateway/connection"
	"steve/gateway/watchdog"
	"steve/structs/proto/base"
	"steve/structs/proto/gate_rpc"

	"github.com/Sirupsen/logrus"
	"github.com/golang/protobuf/proto"
)

type sender struct{}

var _ steve_proto_gaterpc.MessageSenderServer = new(sender)



func (mss *sender) SendMessage(ctx context.Context, req *steve_proto_gaterpc.SendMessageRequest) (*steve_proto_gaterpc.SendMessageResult, error) {
	msgID := req.GetHeader().GetMsgId()
	playerIDs := req.GetPlayerId()
	logEntry := logrus.WithFields(logrus.Fields{
		"func_name":  "sender.SendMessage",
		"msg_id":     msgid.MsgID(msgID),
		"player_ids": req.GetPlayerId(),
	})
	header := base.Header{
		MsgId:   proto.Uint32(msgID),
		Version: proto.String("1.0"), // TODO
	}
	result := &steve_proto_gaterpc.SendMessageResult{}

	clientIDs := mss.fetchConnectionIDs(playerIDs)
	logEntry = logEntry.WithField("client_ids", clientIDs)

	if len(clientIDs) != 0 {
		dog := watchdog.Get()
		err := dog.BroadPackage(clientIDs, &header, req.GetData())
		if err != nil {
			logEntry.WithError(err).Warningln("广播消息失败")
			result.Ok = false
		} else {
			logEntry.Debugln("广播消息完成")
			result.Ok = true
		}
	}
	return result, nil
}




func (mss *sender) fetchConnectionIDs(playerIDs []uint64) []uint64 {
	result := make([]uint64, 0, len(playerIDs))

	connMgr := connection.GetConnectionMgr()
	for _, playerID := range playerIDs {
		connection := connMgr.GetPlayerConnection(playerID)
		if connection != nil {
			result = append(result, connection.GetClientID())
		}
	}
	return result
}
