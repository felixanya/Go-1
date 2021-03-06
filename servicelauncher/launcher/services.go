package launcher

import (
	gatewaycore "steve/gateway/core"
	goldcore "steve/gold/core"
	hallcore "steve/hall/core"
	logincore "steve/login/core"
	matchcore "steve/match/core"
	msgcore "steve/msgserver/core"
	"steve/serviceloader/loader"
	"steve/structs/service"
	testcore "steve/testserver/core"

	"github.com/Sirupsen/logrus"

	"steve/servicelauncher/cmd"
	"steve/structs"
	"steve/configuration/core"
)

// LoadService load service appointed by name
func LoadService() {
	var svr service.Service
	switch cmd.ServiceName {
	case "hall":
		svr = hallcore.NewService()
	case "login":
		svr = logincore.NewService()
	case "match":
		svr = matchcore.NewService()
	// case "room":
	// 	svr = roomcore.NewService()
	case "testserver":
		svr = testcore.NewService()
	case "msgserver":
		svr = msgcore.NewService()
	case "gateway":
		svr = gatewaycore.NewService()
	case "gold":
		svr = goldcore.NewService()
	case "datareport":
		//svr = datareport.NewService()
	case "config":
		svr = core.NewService()
	}
	if svr != nil {
		exposer := structs.GetGlobalExposer()
		svr.Init(exposer)
		loader.Run(svr, exposer, cmd.Option)
	} else {
		logrus.Errorln("no service found service name : ", svr)
		panic("no service found")
	}
}
