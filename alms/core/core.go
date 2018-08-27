package core

/*
	功能： 服务控制逻辑中心，实现服务定义，Client消息分派初始化，和服务启动逻辑

*/
import (
	"runtime"
	"steve/alms/almsserver"
	"steve/alms/data"
	"steve/structs"
	"steve/structs/service"

	s_alms "steve/server_pb/alms"

	"github.com/Sirupsen/logrus"
)

type AlmsCore struct {
}

// NewService 创建服务
func NewService() service.Service {
	return new(AlmsCore)
}

func (a *AlmsCore) Init(e *structs.Exposer, param ...string) error {
	runtime.GOMAXPROCS(1) //单线程序
	entry := logrus.WithField("name", "AlmsCore.Init")
	// 注册当前模块RPC服务处理器
	if err := e.RPCServer.RegisterService(s_alms.RegisterPacksackServerServer, &almsserver.PacksackServer{}); err != nil {
		entry.WithError(err).Error("PacksackServer 注册RPC服务处理器失败")
		return err
	}
	// 注册客户端Client消息处理器
	if err := registerHandles(e.Exchanger); err != nil {
		entry.WithError(err).Error("注册客户端Client消息处理器失败")
		return err
	}
	//初始化救济金配置
	err := data.InitAlmsConfig()
	entry.Debugf("AlmsCoreserver init succeed ...")
	return err
}

func (a *AlmsCore) Start() error {
	logrus.Debugf("AlmsCore server start succeed ...")
	return nil
}
