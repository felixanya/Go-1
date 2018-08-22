package config

type AlmsConfig struct {
	AlmsCountDown    int `json:"almsCountDown"`
	DepositCountDown int `json:"depositCountDown"`
	GetNorm          int `json:"getNorm"`
	GetTimes         int `json:"getTimes"`
	GetNumber        int `json:"getNumber"`
	Version          int `json:"version"`
}
