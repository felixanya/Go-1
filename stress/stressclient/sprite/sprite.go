package sprite

import "github.com/prometheus/client_golang/prometheus"

var stage *Stage
func InitStage() {
	stage = &Stage{}
}
func GetStage() *Stage {
	return stage
}

type Stage struct {
	StartLoginCounter prometheus.Counter
	CPUGauge prometheus.Gauge
	MemoryGauge prometheus.Gauge
	LoginHis prometheus.Histogram
}

type Sprite interface {
	Init() error
	Start(params []string) error
	Stop() error
}
