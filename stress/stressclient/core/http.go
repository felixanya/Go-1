package core

import (
	"net/http"
	"github.com/spf13/viper"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/Sirupsen/logrus"
	"steve/stress/common"
)

func startHttp() {
	httpMux := http.NewServeMux()
	httpMux.HandleFunc("/log/", http.StripPrefix("/log/", http.FileServer(http.Dir(common.LogPath))).ServeHTTP)
	// Expose the registered metrics via HTTP.
	httpMux.Handle("/metrics", promhttp.Handler())
	addr := viper.GetString("prometheus_addr")
	logrus.Info(http.ListenAndServe(addr, httpMux))
}