package watchdog

import (
	stdnet "net"
	"steve/structs/net"
)

type exchanger interface {
	Recv() ([]byte, error)
	Send([]byte) error
	GetRemoteAddr() stdnet.Addr
}

type workerFunc func(exchanger exchanger) error

type server interface {
	Serve(addr string, worker workerFunc) error
	Close()
}

func newServer(addr string, t net.ServerType) server {
	switch t {
	case net.RPC:
		return new(rpcServer)
	case net.TCP:
		return new(tcpServer)
	default:
		return nil
	}
}
