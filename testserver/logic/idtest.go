package logic

import (
	"time"
	"github.com/Sirupsen/logrus"
	"steve/external/idclient"
)

func startIdServer() {
	testIdServer()
}


func testIdServer() {

	begin := time.Now().UnixNano()

	pid, showid , err := idclient.NewPlayerShowId()
	logrus.Infof("NewPlayerShowId: pid=%d, showid=%d, err=%v++++++++++", pid,showid, err)
	end := time.Now().UnixNano()
	logrus.Infof("GetUserProps1=%d,  useTime=%d(ms), err=%v------------------------1",0,   (end - begin)/1000000, err )

}
