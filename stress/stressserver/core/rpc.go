package core

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"net"
	"steve/stress/proto"
	"sync"
)

func createRPCServer(keyFile string, certFile string) *grpc.Server {
	Clients = new(sync.Map)
	PrometheusJson = []*PrometheusClient{}
	writeJson()
	logEntry := logrus.WithFields(logrus.Fields{
		"func_name": "createRPCServer",
		"key_file":  keyFile,
		"cert_file": certFile,
	})
	rpcOption := []grpc.ServerOption{}
	if keyFile != "" {
		cred, err := credentials.NewServerTLSFromFile(certFile, keyFile)
		if err != nil {
			logEntry.Panicln("创建 TLS 证书失败")
		}
		rpcOption = append(rpcOption, grpc.Creds(cred))
	}
	return grpc.NewServer(rpcOption...)
}

// runRPCServer 启动 RPC 服务
func runRPCServer(addr string, port int) {
	client.RegisterPushServer(Grpc, &ClientServer{})
	logEntry := logrus.WithFields(logrus.Fields{
		"func_name": "runRPCServer",
		"addr":      addr,
		"port":      port,
	})
	logEntry.Infoln("启动 RPC 服务")
	if addr == "" || port == 0 {
		logEntry.Info("未配置 RPC 地址或者端口，不启动 RPC 服务")
		return
	}
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", addr, port))
	if err != nil {
		return
	}
	if err := Grpc.Serve(lis); err != nil {
		logEntry.WithError(err).Panicln("启动 RPC 服务失败")
	}
	logEntry.Infoln("RPC 服务完成")
}
