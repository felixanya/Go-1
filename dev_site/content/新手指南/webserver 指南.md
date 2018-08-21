---
title: "webserver 指南"
date: 2018-08-21T13:05:09+08:00
author: 安佳玮
draft: false
---

# 概述

* webserver 用于转发平台的请求到其它应用服
* webserver 接收的请求格式为 json， 回复格式也是 json
* 应用服务可以通过 serviceloader 注册命令处理器，收到请求时系统会回调处理器

# 请求格式

```json
{
    "server": "{server-name}", 
    "cmd": "{command-name}",
    "data": {
        "{app-json-format-data}"
    }
}
```

* server 表示服务名称，需要与应用服的 rpc_server_name 一致
* cmd 表示命令，各应用服独立分配
* data 表示请求数据，json 格式

# 响应格式

```json
{
    "code": 0,
    "msg": "{some-msg}",
    "data": {"app-response-json-format-data"}
}
```

* code 表示错误码，在 common/constant/constant.go/HTTP* 定义
* msg 表示回复信息
* data 表示响应数据，json 格式


# 处理器注册

* 使用全局 Exposer.WebHandlerMgr 注册命令处理器
* 处理器签名为： func(requestData []byte) (code int, msg string, responseData []byte)
    * requestData 表示请求数据， json 格式
    * code， msg, responseData 分别对应响应数据中的三个字段

# 示例

以配置服为例，配置服中需要接收平台发来的配置更新请求，定义的命令处理器如下：

```go
func handleConfigUpdatev2(requestData []byte) (code int, msg string, responseData []byte) {
	entry := logrus.WithField("request_data", string(requestData))
	request := struct {
		Key    string `json:"key"`
		Subkey string `json:"subkey"`
	}{}
	if err := json.Unmarshal(requestData, &request); err != nil {
		return commonconstant.HTTPFailure, "参数错误", nil
    }
    // 
    // do-with-request
    //
	return commonconstant.HTTPOK, "成功", nil
}
```

在服务初始化时，注册命令处理器： 

```go 
func (s *configService) Init(e *structs.Exposer, param ...string) error {
	// some-other-code-init
	if err := e.WebHandleMgr.Register("update", handleConfigUpdatev2); err != nil {
		return fmt.Errorf("注册配置更新处理器失败:%s", err.Error())
	}
	return nil
}
```

然后通过 http://web-server-address 可以向配置服发送命令(推荐使用 postman )， 请求数据可以为:

```json
{
	"server": "configuration",
	"cmd": "update",
	"data": {
		"key": "charge",
		"subkey": "item_list"
	}
}
```

收到的响应数据为：

```json
{"code":0,"msg":"成功","data":null}
```


