package core

/*
	功能： 服务控制逻辑中心，实现服务定义，Client消息分派初始化，和服务启动逻辑

*/
import (
	"steve/structs"
	"steve/structs/service"

	"github.com/Sirupsen/logrus"
)

type AlmsCore struct {
}

// NewService 创建服务
func NewService() service.Service {
	return new(AlmsCore)
}

func (a *AlmsCore) Init(e *structs.Exposer, param ...string) error {
	entry := logrus.WithField("name", "AlmsCore.Init")
	// 注册客户端Client消息处理器
	if err := registerHandles(e.Exchanger); err != nil {
		entry.WithError(err).Error("注册客户端Client消息处理器失败")
		return err
	}
	entry.Debugf("AlmsCoreserver init succeed ...")
	return nil
}

func (a *AlmsCore) Start() error {
	logrus.Debugf("AlmsCore server start succeed ...")
	return nil
}
