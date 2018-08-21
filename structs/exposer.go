package structs

import (
	"steve/structs/arg"
	"steve/structs/configuration"
	"steve/structs/consul"
	"steve/structs/exchanger"
	"steve/structs/net"
	"steve/structs/pubsub"
	"steve/structs/redisfactory"
	"steve/structs/rpc"
	"steve/structs/sgrpc"
)

// Exposer provide common interfaces for services
type Exposer struct {
	RPCServer       sgrpc.RPCServer
	RPCClient       rpc.Client
	Configuration   configuration.Configuration
	WatchDogFactory net.WatchDogFactory
	Exchanger       exchanger.Exchanger
	MysqlEngineMgr  MysqlEngineMgr
	RedisFactory    redisfactory.RedisFactory
	Publisher       pubsub.Publisher
	Subscriber      pubsub.Subscriber
	Option          arg.Option
	ConsulReq       consul.Requester // consul请求接口
	WebHandleMgr    WebHandleMgr
}

var gExposer *Exposer

// GetGlobalExposer 获取全局 exposer 对象
// 全局对象在 servieloader 调用 Init 函数之前设置
func GetGlobalExposer() *Exposer {
	return gExposer
}

// SetGlobalExposer 设置全局 exposer
func SetGlobalExposer(e *Exposer) {
	gExposer = e
}

// ------------------------------------------------

// WebHandleMgr 运维后台消息处理器管理
type WebHandleMgr interface {
	Register(command string, handler WebHandler) error
}

// WebHandler 运维后台消息处理器
type WebHandler func(requestData []byte) (code int, msg string, responseData []byte)
