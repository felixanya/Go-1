package core

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"fmt"
	"steve/stress/proto"
	"github.com/spf13/viper"
	"strings"
	"strconv"
	"io"
	"github.com/Sirupsen/logrus"
	"context"
)

var rpcconn *grpc.ClientConn

// 长连接，和服务器双向通信
func recv() {
	var err error
	rpcconn, err = createRPCClient()
	if err != nil {
		return
	}
	c := client.NewPushClient(rpcconn)
	p := &client.Client{}
	addr := viper.GetString("prometheus_addr")
	port := addr[strings.LastIndex(addr, ":")+1:]
	port32, _ := strconv.Atoi(port)
	p.Port = int32(port32)
	stream, _ := c.PushCommand(context.Background(), p)
	fmt.Println(stream)
	waitc := make(chan struct{})
	go func() {
		for{
			in, err := stream.Recv()
			if err == io.EOF {
				close(waitc)
			}
			if err != nil {
				logrus.Fatalf("Failed to receive a note : %v", err)
			}
			logrus.Printf("Got server command: %d, %s", in.Cmd, in.Params)
			doServerCommand(in)
		}
	}()
	<-waitc
}

func createRPCClient() (*grpc.ClientConn, error) {
	caFile := viper.GetString("ca_file")
	caServerName := viper.GetString("ca_server")
	addr := viper.GetString("rpc_addr")
	logEntry := logrus.WithFields(logrus.Fields{
		"func_name":      "createRPCClient",
		"ca_file":        caFile,
		"ca_server_name": caServerName,
		"addr":    addr,
	})
	opts := []grpc.DialOption{}
	if caFile != "" {
		c, err := credentials.NewClientTLSFromFile(caFile, caServerName)
		if err != nil {
			logEntry.Panicln("创建 RPC 客户端失败")
			return nil, fmt.Errorf("create client tls failed:%v", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(c))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}
	logEntry.Info("dial rpc server")
	cc, err := grpc.Dial(addr, opts...)
	return cc, err
}