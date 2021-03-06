package mjoption

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/Sirupsen/logrus"
	yaml "gopkg.in/yaml.v2"
)

// FanType 番型
type FanType struct {
	ID      int   `yaml:"id"`      // 番型 ID
	FuncID  int   `yaml:"func_id"` // 计算函数 ID
	Mutex   []int `yaml:"mutex"`   // 互斥番型列表
	Method  int   `yaml:"method"`  // 分数计算方式，0为相加，1为相乘
	Score   int   `yaml:"score"`   // 番数
	Type    int   `yaml:"type"`    // 番/倍
	SubGeng int   `yaml:"subgeng"` // 扣除的根数量
}

// HuType 胡牌类型
type HuType struct {
	ID int `yaml:"huType"` // 胡牌类型 ID
}

// SettleType 结算类型
type SettleType struct {
	ID int `yaml:"settleType"` // 结算类型 ID
}

// CardTypeOption 牌型选项
type CardTypeOption struct {
	ID                int                `yaml:"id"`                     // 选项 ID
	Fantypes          map[int]FanType    `yaml:"enable_fan_types"`       // 支持的番型
	FanType2HuType    map[int]HuType     `yaml:"fan_type_2_hu_type"`     // 番型转胡类型
	FanType2Settle    map[int]SettleType `yaml:"fan_type_2_settle_type"` // 番型转结算类型
	EnableGeng        bool               `yaml:"enable_geng"`            // 是否启用根
	GengScore         int                `yaml:"geng_score"`             // 根的番数
	GengMethod        int                `yaml:"geng_method"`            // 根的计算方式，0为相加，1为相乘，2幂乘
	EnableHua         bool               `yaml:"enable_hua"`             // 是否启用花
	HuaScore          int                `yaml:"hua_score"`              // 花的番数
	HuaMethod         int                `yaml:"hua_method"`             // 花的计算方式，0为相加，1为相乘，2幂乘
	EnableFanTypeDeal bool               `yaml:"enable_fan_type_deal"`   // 番型处理，是否将胡类型从，番型拿出
}

// CardTypeOptionManager 选项管理器
type CardTypeOptionManager struct {
	cardTypeOptionMap map[int]*CardTypeOption
}

// GetCardTypeOption 获取牌型选项
func (som *CardTypeOptionManager) GetCardTypeOption(optID int) *CardTypeOption {
	if opt, ok := som.cardTypeOptionMap[optID]; ok {
		return opt
	}
	return nil
}

func (som *CardTypeOptionManager) loadOption(path string) {
	entry := logrus.WithFields(logrus.Fields{
		"func_name": "CardTypeOptionManager.loadOption",
		"path":      path,
	})
	if !strings.HasSuffix(path, "yaml") {
		return
	}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		entry.WithError(err).Panicln("读取文件失败")
	}
	opt := CardTypeOption{}
	if err := yaml.Unmarshal(data, &opt); err != nil {
		entry.WithError(err).Panicln("反序列化失败")
	}
	if _, exist := som.cardTypeOptionMap[opt.ID]; exist {
		entry.WithField("id", opt.ID).Panicln("结算选项 ID 重复")
	}
	som.cardTypeOptionMap[opt.ID] = &opt
}

// loadOptions 加载选项文件
func (som *CardTypeOptionManager) loadOptions(optionDir string) {
	som.cardTypeOptionMap = make(map[int]*CardTypeOption)
	filepath.Walk(optionDir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			som.loadOption(path)
		}
		return nil
	})
}

// NewCardTypeOptionManager is CardType option manager creator
func NewCardTypeOptionManager(optDir string) *CardTypeOptionManager {
	som := &CardTypeOptionManager{}
	som.loadOptions(optDir)
	return som
}
