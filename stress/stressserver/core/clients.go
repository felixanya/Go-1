package core

import (
	"encoding/json"
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc/peer"
	"net"
	"os"
	"steve/stress/proto"
	"strconv"
	"sync"
	"time"
)

var Clients *sync.Map
var lock sync.Mutex
var PrometheusJson []*PrometheusClient

//JSONTime ... JSONTime
type JSONTime time.Time

//Client ... Client
type Client struct {
	ID        int64                          `json:"ID"`
	StartTime JSONTime                       `json:"StartTime"`
	IP        string                         `json:"IP"`
	Server    string                         `json:"Server"`
	Started   bool                           `json:"Started"`
	Global    string                         `json:"Global"`
	Config    string                         `json:"Config"`
	Path      string                         `json:"Path"`
	MD5       string                         `json:"MD5"`
	Version   string                         `json:"Version"`
	Stream    *client.Push_PushCommandServer `json:"-"`
}

type PrometheusLabel struct {
	Env string `json:"env"`
	Job string `json:"job"`
}
type PrometheusClient struct {
	Targets []string        `json:"targets"`
	Labels  PrometheusLabel `json:"labels"`
}

type ClientServer struct {
}

func (s *ClientServer) PushCommand(c *client.Client, stream client.Push_PushCommandServer) error {
	pp, _ := peer.FromContext(stream.Context())
	ip := pp.Addr.(*net.TCPAddr).IP.String()
	swap := viper.GetString("prometheus_swap_localhost")
	if swap != "" && (ip == "127.0.0.1" || ip == "localhost") {
		ip = swap
	}
	clientAddr := fmt.Sprintf("%s:%d", ip, c.Port)

	currentTime := time.Now()
	clientID := currentTime.UnixNano() / 1000
	jsonTime := JSONTime(currentTime)
	logrus.Infof("%s connected(%d)", clientAddr, clientID)

	lock.Lock()
	var cc Client
	cc = Client{
		ID:        clientID,
		StartTime: jsonTime,
		IP:        clientAddr,
		Server:    "",
		Config:    "",
		Path:      "",
		MD5:       "",
		Version:   "",
		Stream:    &stream,
	}
	Clients.Store(clientID, cc)
	total := len(PrometheusJson)
	clientJson := PrometheusClient{}
	target := ip + ":" + strconv.Itoa(int(c.Port))
	clientJson.Targets = []string{target}
	clientJson.Labels = PrometheusLabel{Env: "prod", Job: "client" + strconv.Itoa(total)}
	PrometheusJson = append(PrometheusJson, &clientJson)
	lock.Unlock()
	writeJson()

	idmap := make(map[string][]string)
	idstr := strconv.Itoa(int(clientID))
	idmap["params"] = []string{idstr}
	params, _ := json.Marshal(idmap)
	serverCmd := &client.ServerCommand{Cmd: 0, Params: string(params)}
	stream.Send(serverCmd)

	select {
	case <-stream.Context().Done():
		logrus.Infof("%s disconnected", clientAddr)
	}
	lock.Lock()
	for i, v := range PrometheusJson {
		addr := v.Targets[0]
		if addr == target {
			PrometheusJson = append(PrometheusJson[:i], PrometheusJson[i+1:]...)
			break
		}
	}
	Clients.Delete(clientID)
	lock.Unlock()
	writeJson()

	return nil
}

func writeJson() {
	j, _ := json.MarshalIndent(PrometheusJson, "", "  ")
	ss := string(j)
	filePath := viper.GetString("prometheus_config")
	w, err := os.OpenFile(filePath, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0644)
	if err != nil {
		logrus.Error("error Open", filePath, err)
	}
	n, err := fmt.Fprintf(w, ss)
	if err != nil {
		logrus.Error("error Fprintf", n, filePath, err)
	}
}

//MarshalJSON ... MarshalJSON
func (t JSONTime) MarshalJSON() ([]byte, error) {
	//do your serializing here
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}
