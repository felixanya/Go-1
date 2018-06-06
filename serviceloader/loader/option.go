package loader

import (
	"reflect"

	"github.com/Sirupsen/logrus"
)

type option struct {
	rpcCertiFile    string
	rpcKeyFile      string
	rpcAddr         string // RPC服务监听地址
	rpcPort         int    // RPC端口号
	rpcServerName   string // 服务器名称
	params          []string
	rpcCAFile       string // RPC客户端的CA文件
	rpcCAServerName string // 证书中的服务器名称
	redisAddr       string // redis 服务地址
	redisPasswd     string // redis 密码
}

var defaultOption = option{
	redisAddr:   "127.0.0.1:6379",
	redisPasswd: "",
}

// ServiceOption ...
type ServiceOption func(opt *option)

// WithParams 参数选项， 参数将透传给 plugin
func WithParams(params []string) ServiceOption {
	return func(opt *option) {
		opt.params = params
	}
}

// WithRedisOption 设置 redis 选项
func WithRedisOption(addr, passwd string) ServiceOption {
	return func(opt *option) {
		opt.redisAddr = addr
		opt.redisPasswd = passwd
	}
}

// WithRPCParams RPC 选项， certiFile 为证书文件， keyFile 为私钥文件， addr 为 RPC 服务监听地址， port 为 RPC 服务监听端口
// serverName 为 RPC 服务名字
func WithRPCParams(certiFile string, keyFile string, addr string, port int, serverName string) ServiceOption {
	return func(opt *option) {
		opt.rpcCertiFile = certiFile
		opt.rpcKeyFile = keyFile
		opt.rpcAddr = addr
		opt.rpcPort = port
		opt.rpcServerName = serverName
	}
}

// WithClientRPCCA 客户端 RPC CA 证书选项， caFile 为 CA 证书文件， serverName 为服务的证书域名字段
func WithClientRPCCA(caFile, serverName string) ServiceOption {
	return func(opt *option) {
		opt.rpcCAFile = caFile
		opt.rpcCAServerName = serverName
	}
}

// infoOption 输出选项信息
func infoOption(opt option) {
	fields := make(logrus.Fields)
	t := reflect.TypeOf(opt)
	v := reflect.ValueOf(opt)
	for i := 0; i < t.NumField(); i++ {
		ft := t.Field(i)
		fName := ft.Name
		fv := v.Field(i)
		fields[fName] = fv
	}
	logrus.WithFields(fields).Info("服务选项列表")
}

// loadOptions 加载服务选项
func loadOptions(options ...ServiceOption) option {
	opt := defaultOption
	for _, option := range options {
		option(&opt)
	}
	infoOption(opt)
	return opt
}
