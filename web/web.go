package main

import (
	"steve/structs/service"
	"steve/web/core"
)

// GetService 获取服务接口，被 serviceloader 调用
func GetService() service.Service {
	return core.NewService()
}

func main() {}
