package logic

import (
	"time"
	"github.com/Sirupsen/logrus"
	"steve/external/idclient"
	"sync"
)

func startIdServer() {
	testIdServer()
}


func testIdServer() {

	begin := time.Now().UnixNano()

	pid, showid , err := idclient.NewPlayerShowId()
	logrus.Infof("NewPlayerShowId: pid=%d, showid=%d, err=%v++++++++++", pid,showid, err)
	end := time.Now().UnixNano()
	logrus.Infof("NewPlayerShowId=%d,  useTime=%d(ms), err=%v------------------------1",0,   (end - begin)/1000000, err )
	begin = time.Now().UnixNano()

	wg := sync.WaitGroup{}

	for i:= 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			pid, showid, err := idclient.NewPlayerShowId()

			if err != nil {
				logrus.Errorf("NewPlayerShowId err: pid=%d, showid=%d, err=%v++++++++++", pid, showid, err)
			}
		}()
	}

	wg.Wait()
	end = time.Now().UnixNano()
	logrus.Infof("NewPlayerShowId 2=%d,  useTime=%d(ms), err=%v------------------------1",0,   (end - begin)/1000000, err )

}
