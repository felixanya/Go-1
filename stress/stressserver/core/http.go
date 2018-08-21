package core

import (
	"net/http"
	"fmt"
	"path/filepath"
	"github.com/spf13/viper"
	"steve/serviceloader/pprof"
	"html/template"
	"encoding/json"
	"github.com/Sirupsen/logrus"
	"strconv"
	"steve/stress/proto"
	"steve/stress/common"
	"github.com/spf13/cast"
)

func startHttp() {
	httpMux := http.NewServeMux()
	httpMux.HandleFunc("/log/", http.StripPrefix("/log/", http.FileServer(http.Dir(common.LogPath))).ServeHTTP)
	httpMux.HandleFunc("/report/", http.StripPrefix("/report/", http.FileServer(http.Dir(viper.GetString("report_path")))).ServeHTTP)
	httpMux.HandleFunc("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))).ServeHTTP)
	httpMux.HandleFunc("/api/clients", getClients)
	httpMux.HandleFunc("/api/startClient", startClient)
	httpMux.HandleFunc("/api/stopClient", stopClient)
	httpMux.HandleFunc("/api/finished", finished)
	httpMux.HandleFunc("/view", httpView)
	pprof.Init("stressserver", "svg", viper.GetInt("http_port"), httpMux)
	//if err := http.ListenAndServe(":8885", nil); err != nil {
	//	logrus.Fatal("failed to start stress server", err)
	//}
}

func startClient(w http.ResponseWriter, r *http.Request) {
	id, result := startStopClient(true, r)
	writeResponse(w, id, result)
}
func stopClient(w http.ResponseWriter, r *http.Request) {
	id, result := startStopClient(false, r)
	writeResponse(w, id, result)
}
func writeResponse(w http.ResponseWriter, id int, result int) {
	m := make(map[string]int)
	m["id"] = id
	m["result"] = result
	json, _ := json.Marshal(m)
	fmt.Fprintf(w, string(json))
}
func startStopClient(isStart bool, r *http.Request) (id int, result int) {
	r.ParseForm()
	ID := r.Form["id"]
	if len(ID) == 0 {
		return 0,1
	}
	id, _ = strconv.Atoi(ID[0])
	configs := r.Form["config"]
	var config string
	if len(configs) > 0 {
		config = configs[0]
	}

	lock.Lock()
	cc, ok := Clients.Load(int64(id))
	if ok == false {
		logrus.Error(ID, "not exist")
		return id,2
	}
	ccc := cc.(Client)
	ccc.Started = isStart
	ccc.Config = config
	Clients.Store(int64(id), ccc)
	lock.Unlock()

	var startCmd int32
	if isStart {
		startCmd = 1
	} else {
		startCmd = 2
	}
	serverCmd := &client.ServerCommand{Cmd: startCmd}
	serverCmd.Params = config
	(*ccc.Stream).Send(serverCmd)
	logrus.Infof("%d isStart: %t", ID, isStart)
	return id,0
}

func getClients(w http.ResponseWriter, r *http.Request) {
	jsonmap := make(map[string]interface{})
	list := []Client{}
	Clients.Range(func(key interface{}, value interface{}) bool {
		client := value.(Client)
		list = append(list, client)
		return true
	})
	jsonmap["list"] = list
	grafana := make(map[string]interface{})
	boards := viper.GetStringMap("grafana")
	bb := boards["boards"].([]interface{})
	grafanaBoards := make([]map[string]string, len(bb))
	for i := 0; i < len(bb); i++ {
		grafanaBoards[i] = make(map[string]string)
		bbb := bb[i].(map[interface{}]interface{})
		grafanaBoards[i]["name"] = cast.ToString(bbb["name"])
		grafanaBoards[i]["url"] = cast.ToString(bbb["url"])
	}
	grafana["boards"] = grafanaBoards
	grafana["url"] = viper.GetString("grafana.url")
	grafana["clients_url"] = viper.GetString("grafana.clients_url")
	grafana["finished_url"] = viper.GetString("grafana.finished_url")
	jsonmap["grafana"] = grafana
	j, err := json.Marshal(jsonmap)
	if err != nil {
		logrus.Error(err.Error())
	}
	s := string(j)
	fmt.Fprintf(w, s)
}

func finished(w http.ResponseWriter, r *http.Request) {
	m := make(map[string]int)
	m["id"] = 354
	m["result"] = 75
	json, _ := json.Marshal(m)
	fmt.Fprintf(w, string(json))
}

func httpView(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("GetClients", r.URL.Path)

	lp := filepath.Join("views", "layout.html")
	fp := filepath.Join("views", filepath.Clean("clients.html"))

	tmpl, _ := template.ParseFiles(lp, fp)
	tmpl.ExecuteTemplate(w, "layout", nil)
}
