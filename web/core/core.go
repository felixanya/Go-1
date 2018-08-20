package core

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"steve/common/constant"
	"steve/server_pb/web"
	"steve/structs"
	"steve/structs/service"

	"github.com/Sirupsen/logrus"
	"github.com/spf13/viper"
)

type webcore struct {
	e *structs.Exposer
}

// NewService 创建服务
func NewService() service.Service {
	return new(webcore)
}

func (c *webcore) Init(e *structs.Exposer, param ...string) error {
	viper.SetDefault("http_listen_addr", "0.0.0.0:36900")

	http.HandleFunc("/", transmitHTTPRequest)
	return nil
}

func (c *webcore) Start() error {
	http.ListenAndServe(viper.GetString("http_listen_addr"), nil)
	return nil
}

/*
transmitHTTPRequest 转发 http 请求。请求和响应数据均为 json 格式。

请求 body 结构:
	{
		"server": {servername},
		"cmd": {cmdname},
		"data": {
			{some-json-data}
		}
	}

响应 body 结构:
	{
		"code": {response-code},
		"msg": {result-msg},
		"data": {
			{some-json-data}
		}
	}
*/
func transmitHTTPRequest(writer http.ResponseWriter, request *http.Request) {
	if request.Body == nil {
		writeResponse(writer, "读取数据失败", constant.HTTPFailure, nil)
		logrus.Infoln("获取 body 失败")
		return
	}
	requestData, err := ioutil.ReadAll(request.Body)
	if err != nil {
		writeResponse(writer, "读取数据失败", constant.HTTPFailure, nil)
		logrus.WithError(err).Infoln("读取数据失败")
		return
	}
	entry := logrus.WithField("request_data", string(requestData))

	var requestJSON = requestJSONData{}
	if err := json.Unmarshal(requestData, &requestJSON); err != nil {
		writeResponse(writer, "请求数据反序列化失败", constant.HTTPInvalidRequest, nil)
		entry.WithError(err).Infoln("请求数据反序列化失败")
		return
	}
	entry = entry.WithFields(logrus.Fields{
		"server": requestJSON.Server,
		"cmd":    requestJSON.Cmd,
		"data":   string(requestJSON.Data),
	})
	responseJSONData, err := serverRequester(&requestJSON)
	entry.WithFields(logrus.Fields{
		"response-code": responseJSONData.Code,
		"response-msg":  responseJSONData.Msg,
		"response-data": string(responseJSONData.Data),
	}).WithError(err).Infoln("处理完成")
	writeResponseJSON(writer, responseJSONData)
}

// writeResponse 响应 http 请求
// 格式参考 transmitHTTPRequest
func writeResponse(writer http.ResponseWriter, msg string, code int, data []byte) {
	c := responseJSONData{Code: code, Msg: msg, Data: data}
	writeResponseJSON(writer, c)
}

// writeResponseJSON 响应 http 请求
func writeResponseJSON(writer http.ResponseWriter, responseData responseJSONData) {
	jsonData, err := json.Marshal(responseData)
	if err != nil {
		logrus.WithError(err).Errorln("序列化响应数据失败")
		return
	}
	if _, err := writer.Write(jsonData); err != nil {
		logrus.WithError(err).Errorln("写响应数据失败")
		return
	}
}

type requestJSONData struct {
	Server string          `json:"server"`
	Cmd    string          `json:"cmd"`
	Data   json.RawMessage `json:"data"`
}

type responseJSONData struct {
	Code int             `json:"code"`
	Msg  string          `json:"msg"`
	Data json.RawMessage `json:"data"`
}

var serverRequester = defaultRequestServer

// defaultRequestServer 向指定服务发起请求
func defaultRequestServer(requestJSON *requestJSONData) (responseJSONData, error) {
	responseJSON := responseJSONData{
		Code: constant.HTTPFailure,
		Msg:  "处理失败",
		Data: nil,
	}
	cli, err := structs.GetGlobalExposer().RPCClient.GetConnectByServerName(requestJSON.Server)
	if err != nil {
		return responseJSON, fmt.Errorf("目标服务不存在: %s", err.Error())
	}
	requestCli := web.NewRequestHandlerClient(cli)
	response, err := requestCli.Handle(context.Background(), &web.RequestData{
		Cmd:  requestJSON.Cmd,
		Data: requestJSON.Data,
	})
	if err != nil {
		return responseJSON, fmt.Errorf("RPC 调用失败：%s", err.Error())
	}
	return responseJSONData{
		Code: int(response.GetCode()),
		Msg:  response.GetMsg(),
		Data: response.GetData(),
	}, nil
}
