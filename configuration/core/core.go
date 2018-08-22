package core

import (
	"encoding/json"
	"fmt"
	"net/http"
	commonconstant "steve/common/constant"
	"steve/configuration/data"
	"steve/entity/constant"
	"steve/server_pb/config"
	"steve/structs"
	"steve/structs/service"

	"github.com/Sirupsen/logrus"
	"github.com/golang/protobuf/proto"
	"github.com/spf13/viper"
)

type configService struct{}

func (s *configService) Init(e *structs.Exposer, param ...string) error {
	http.HandleFunc("/update", handleConfigUpdate)
	if err := e.RPCServer.RegisterService(config.RegisterConfigServer, &configServer{}); err != nil {
		return fmt.Errorf("服务注册失败：%v", err)
	}
	if err := e.WebHandleMgr.Register("update", handleConfigUpdatev2); err != nil {
		return fmt.Errorf("注册配置更新处理器失败:%s", err.Error())
	}
	return nil
}

func (s *configService) Start() error {
	if httpAddr := viper.GetString("http_addr"); httpAddr != "" {
		logrus.WithField("addr", httpAddr).Debugln("启动 http 服务")
		http.ListenAndServe(httpAddr, nil)
	}
	return nil
}

// NewService 创建服务
func NewService() service.Service {
	return &configService{}
}

// pubConfigUpdate 发布配置更新
func pubConfigUpdate(key, subkey string) error {
	val, err := data.GetConfig(key, subkey)
	if err != nil {
		return fmt.Errorf("获取配置失败:%v", err)
	}
	message := config.ConfigUpdate{
		Key:    key,
		Subkey: subkey,
		Val:    val,
	}
	md, err := proto.Marshal(&message)
	if err != nil {
		return fmt.Errorf("消息序列化失败:%v", err)
	}

	publisher := structs.GetGlobalExposer().Publisher
	if err := publisher.Publish(constant.UpdateConfig, md); err != nil {
		return fmt.Errorf("消息发布失败：%v", err)
	}
	return nil
}

func handleConfigUpdate(w http.ResponseWriter, request *http.Request) {
	key := request.FormValue("key")
	subkey := request.FormValue("subkey")

	if key == "" || subkey == "" {
		w.Write([]byte("参数错误, key 或者 subkey 为空"))
		return
	}
	if err := pubConfigUpdate(key, subkey); err != nil {
		s := fmt.Sprintf("发布通知失败，错误信息：%v", err)
		w.Write([]byte(s))
		return
	}
	w.Write([]byte("更新成功"))
	return
}

/* handleConfigUpdatev2 处理配置更新
参数：
requestData (json) :
{
	"key": "{some-key}",
	"subkey": "{some-sub-key}",
}

responseData (json) : 空
*/
func handleConfigUpdatev2(requestData []byte) (code int, msg string, responseData []byte) {
	entry := logrus.WithField("request_data", string(requestData))
	request := struct {
		Key    string `json:"key"`
		Subkey string `json:"subkey"`
	}{}

	if err := json.Unmarshal(requestData, &request); err != nil {
		entry.WithError(err).Infoln("反序列化失败")
		return commonconstant.HTTPFailure, "参数错误", nil
	}

	if request.Key == "" {
		entry.Infoln("key 为空")
		return commonconstant.HTTPFailure, "参数错误", nil
	}
	if err := pubConfigUpdate(request.Key, request.Subkey); err != nil {
		entry.WithError(err).Errorln("发布配置更新失败")
		return commonconstant.HTTPFailure, "系统异常", nil
	}

	entry.Debugln("配置更新发布成功")
	return commonconstant.HTTPOK, "成功", nil
}
