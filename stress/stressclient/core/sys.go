package core

import (
			"time"
	"github.com/StackExchange/wmi"
		"github.com/spf13/viper"
	"github.com/shirou/gopsutil/mem"
	"steve/stress/stressclient/sprite"
)

type Win32_Processor struct {
	LoadPercentage            *uint16
	Family                    uint16
	Manufacturer              string
	Name                      string
	NumberOfLogicalProcessors uint32
	ProcessorId               *string
	Stepping                  *string
	MaxClockSpeed             uint32
}

var t *time.Ticker
func initSys() {
	t = time.NewTicker(time.Second * time.Duration(viper.GetInt("refresh_rate")))
	go refresh()
}
func refresh() {
	for{
		select {
		case <- t.C:
			v, _ := mem.VirtualMemory()
			var dst []Win32_Processor
			q := wmi.CreateQuery(&dst, "")
			err := wmi.Query(q, &dst)
			if err != nil {
				//return ret, err
				continue
			}
			stage := sprite.GetStage()
			stage.CPUGauge.Set(float64(*dst[0].LoadPercentage))
			stage.MemoryGauge.Set(float64(v.UsedPercent))
			//fmt.Println("cpu: ", *dst[0].LoadPercentage, dst)
			//fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)
		}
	}
}