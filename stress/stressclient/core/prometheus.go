package core

import (
	"steve/stress/stressclient/sprite"
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/viper"
	"github.com/Sirupsen/logrus"
)

func initPrometheus() {
	stage := sprite.GetStage()
	stage.LoginHis = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "start_login",
		Help:    "Number of login operations processed.",
		Buckets: prometheus.LinearBuckets(0, 10, 1),
	})
	prometheus.MustRegister(stage.LoginHis)
	stage.CPUGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "cpu_gauge",
		Help: "CPU LoadPercentage",
	})
	prometheus.MustRegister(stage.CPUGauge)
	stage.MemoryGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "memory_gauge",
		Help: "Memory UsedPercent",
	})
	prometheus.MustRegister(stage.MemoryGauge)

	// Expose the registered metrics via HTTP.
	http.Handle("/metrics", promhttp.Handler())
	addr := viper.GetString("prometheus_addr")
	logrus.Info(http.ListenAndServe(addr, nil))
}