package loader

import (
	"context"
	"fmt"
	"steve/common/constant"
	"steve/server_pb/web"
	"steve/structs"
	"steve/structs/sgrpc"

	"github.com/Sirupsen/logrus"
)

// webHandlerMgr 实现 structs.WebHandleMgr
type webHanlderMgr struct {
	handlers          map[string]structs.WebHandler
	registerChannel   chan registerInfo
	getHandlerChannel chan getHandlerInfo
}

type registerInfo struct {
	command    string
	handler    structs.WebHandler
	errChannel chan error
}

type getHandlerInfo struct {
	command        string
	handlerChannel chan structs.WebHandler
}

func (m *webHanlderMgr) Register(command string, handler structs.WebHandler) error {
	errChannel := make(chan error)
	m.registerChannel <- registerInfo{
		command:    command,
		handler:    handler,
		errChannel: errChannel,
	}
	err := <-errChannel
	return err
}

func (m *webHanlderMgr) getHandler(command string) structs.WebHandler {
	handlerChannel := make(chan structs.WebHandler)
	m.getHandlerChannel <- getHandlerInfo{
		command:        command,
		handlerChannel: handlerChannel,
	}
	return <-handlerChannel
}

func (m *webHanlderMgr) run() {
	for {
		select {
		case registerInfo := <-m.registerChannel:
			{
				if _, ok := m.handlers[registerInfo.command]; ok {
					registerInfo.errChannel <- fmt.Errorf("重复注册")
				}
				m.handlers[registerInfo.command] = registerInfo.handler
				close(registerInfo.errChannel)
			}
		case getHandlerInfo := <-m.getHandlerChannel:
			{
				getHandlerInfo.handlerChannel <- m.handlers[getHandlerInfo.command]
			}
		}
	}
}

// 实现 grpc 服务： web.RequestHandlerServer
type requestHandlerServer struct {
	webHandlerMgr *webHanlderMgr
}

func (s *requestHandlerServer) Handle(ctx context.Context, request *web.RequestData) (*web.ResponseData, error) {
	cmd := request.GetCmd()
	handler := s.webHandlerMgr.getHandler(cmd)
	if handler == nil {
		logrus.WithField("cmd", cmd).Infoln("命令不存在")
		return &web.ResponseData{
			Code: constant.HTTPFailure,
			Msg:  "命令不存在",
		}, nil
	}
	code, msg, responseData := handler(request.GetData())
	return &web.ResponseData{
		Code: int32(code),
		Msg:  msg,
		Data: responseData,
	}, nil
}

func createWebHandlerMgr(rpcServer sgrpc.RPCServer) structs.WebHandleMgr {
	m := &webHanlderMgr{
		handlers:          make(map[string]structs.WebHandler, 16),
		registerChannel:   make(chan registerInfo),
		getHandlerChannel: make(chan getHandlerInfo),
	}
	rpcServer.RegisterService(web.RegisterRequestHandlerServer,
		&requestHandlerServer{webHandlerMgr: m})
	go m.run()
	return m
}
