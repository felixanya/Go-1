package core

import (
	"github.com/shirou/gopsutil/mem"
	"github.com/spf13/viper"
	"steve/stress/stressclient/sprite"
	"time"
	"github.com/shirou/gopsutil/cpu"
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
	for {
		select {
		case <-t.C:
			v, _ := mem.VirtualMemory()
			//var dst []Win32_Processor
			//q := wmi.CreateQuery(&dst, "")
			//err := wmi.Query(q, &dst)
			//if err != nil {
			//	//return ret, err
			//	continue
			//}
			stage := sprite.GetStage()
			cpu, _ := cpu.Percent(2, false)
			stage.CPUGauge.Set(cpu[0])
			stage.MemoryGauge.Set(float64(v.UsedPercent))
			//fmt.Println("cpu: ", *dst[0].LoadPercentage, dst)
			//fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)
		}
	}
}
