package core

import (
	"github.com/prometheus/client_golang/prometheus"
)

var metrics *Metrics

func GetMetrics() *Metrics {
	return metrics
}

type Metrics struct {
	ErrorCounter  prometheus.CounterVec
	histogram     prometheus.HistogramVec
	ConnectsCounter prometheus.CounterVec

	CPUGauge    prometheus.Gauge
	MemoryGauge prometheus.Gauge
}

func initPrometheus() {
	initMetrics()
	startHttp()
}

func initMetrics() {
	labels := []string{"type"}
	metrics = &Metrics{}
	metrics.histogram = *prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "client_histogram",
		Help:    "Number of login operations processed.",
		Buckets: prometheus.LinearBuckets(0, 10, 1),
	}, labels)
	prometheus.MustRegister(metrics.histogram)
	metrics.ErrorCounter = *prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "error_counter",
		Help: "ccc",
	}, labels)
	prometheus.MustRegister(metrics.ErrorCounter)
	metrics.ConnectsCounter = *prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "connects_counter",
		Help: "Started and stopped connects",
	}, labels)
	prometheus.MustRegister(metrics.ConnectsCounter)

	metrics.CPUGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "cpu_gauge",
		Help: "CPU LoadPercentage",
	})
	prometheus.MustRegister(metrics.CPUGauge)
	metrics.MemoryGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "memory_gauge",
		Help: "Memory UsedPercent",
	})
	prometheus.MustRegister(metrics.MemoryGauge)
}

func Observe(observeType string, value float64) {
	metrics.histogram.WithLabelValues(observeType).Observe(value)
}

func AddError(errorType string) {
	metrics.ErrorCounter.WithLabelValues(errorType).Add(1)
}
