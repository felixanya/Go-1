package interfaces

import msgid "steve/client_pb/msgId"

// ClientPlayer 客户端玩家
type ClientPlayer interface {
	GetID() uint64
	GetCoin() uint64
	GetClient() Client
	GetUsrName() string
	GetAccountID() uint64
	AddExpectors(msgID ...msgid.MsgID)
	GetExpector(msgID msgid.MsgID) MessageExpector
}
