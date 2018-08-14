// Code generated by protoc-gen-go. DO NOT EDIT.
// source: majong.proto

package majong

import (
	fmt "fmt"
	math "math"
	"time"

	proto "github.com/golang/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// 杠的类型
type GangType int32

const (
	GangType_gang_angang   GangType = 0
	GangType_gang_minggang GangType = 1
	GangType_gang_bugang   GangType = 2
)

// 胡类型
type HuType int32

const (
	HuType_hu_ganghoupao        HuType = 0
	HuType_hu_qiangganghu       HuType = 1
	HuType_hu_dianpao           HuType = 2
	HuType_hu_gangkai           HuType = 3
	HuType_hu_haidilao          HuType = 4
	HuType_hu_gangshanghaidilao HuType = 5
	HuType_hu_zimo              HuType = 6
	HuType_hu_tianhu            HuType = 7
	HuType_hu_dihu              HuType = 8
	HuType_hu_miaoshouhuichun   HuType = 9
	HuType_hu_renhu             HuType = 10
	HuType_hu_quanqiuren        HuType = 11
	HuType_hu_buqiuren          HuType = 12
	HuType_hu_juezhang          HuType = 13
)

var HuType_name = map[int32]string{
	0:  "hu_ganghoupao",
	1:  "hu_qiangganghu",
	2:  "hu_dianpao",
	3:  "hu_gangkai",
	4:  "hu_haidilao",
	5:  "hu_gangshanghaidilao",
	6:  "hu_zimo",
	7:  "hu_tianhu",
	8:  "hu_dihu",
	9:  "hu_miaoshouhuichun",
	10: "hu_renhu",
	11: "hu_quanqiuren",
	12: "hu_buqiuren",
	13: "hu_juezhang",
}
var HuType_value = map[string]int32{
	"hu_ganghoupao":        0,
	"hu_qiangganghu":       1,
	"hu_dianpao":           2,
	"hu_gangkai":           3,
	"hu_haidilao":          4,
	"hu_gangshanghaidilao": 5,
	"hu_zimo":              6,
	"hu_tianhu":            7,
	"hu_dihu":              8,
	"hu_miaoshouhuichun":   9,
	"hu_renhu":             10,
	"hu_quanqiuren":        11,
	"hu_buqiuren":          12,
	"hu_juezhang":          13,
}

// 玩家行牌状态
type XingPaiState int32

const (
	XingPaiState_normal  XingPaiState = 0
	XingPaiState_hu      XingPaiState = 1
	XingPaiState_give_up XingPaiState = 2
)

// SettleType 结算类型
type SettleType int32

const (
	SettleType_settle_angang    SettleType = 0
	SettleType_settle_minggang  SettleType = 1
	SettleType_settle_bugang    SettleType = 2
	SettleType_settle_dianpao   SettleType = 3
	SettleType_settle_zimo      SettleType = 4
	SettleType_settle_flowerpig SettleType = 5
	SettleType_settle_yell      SettleType = 6
	SettleType_settle_taxrebeat SettleType = 7
	SettleType_settle_calldiver SettleType = 8
)

// 玩家操作,值按优先级排列，值越大，优先级越高
type Action int32

const (
	Action_action_qi   Action = 0
	Action_action_chi  Action = 10
	Action_action_peng Action = 20
	Action_action_gang Action = 30
	// action_bugang       = 31;  // 补杠
	// action_minggang     = 32;  // 明杠
	// action_angang       = 33;  // 暗杠
	Action_action_hu Action = 40
	// action_zimo         = 41;  // 自摸
	// action_dianpao      = 42;  // 点炮
	// action_qiangganghu  = 43;  // 抢杠胡
)

type ActionSlice []Action

func (as ActionSlice) Len() int           { return len(as) }
func (as ActionSlice) Swap(i, j int)      { as[i], as[j] = as[j], as[i] }
func (as ActionSlice) Less(i, j int) bool { return int(as[i]) < int(as[j]) }

// MopaiType 摸牌类型
type MopaiType int32

const (
	MopaiType_MT_NORMAL MopaiType = 0
	MopaiType_MT_GANG   MopaiType = 1
)

// ZixunType 自询类型
type ZixunType int32

const (
	ZixunType_ZXT_NORMAL ZixunType = 0
	ZixunType_ZXT_PENG   ZixunType = 1
	ZixunType_ZXT_CHI    ZixunType = 2
)

// 麻将组类型
type CardsGroupType int32

const (
	CardsGroupType_CGT_HAND     CardsGroupType = 0
	CardsGroupType_CGT_CHI      CardsGroupType = 1
	CardsGroupType_CGT_PENG     CardsGroupType = 2
	CardsGroupType_CGT_MINGGANG CardsGroupType = 3
	CardsGroupType_CGT_ANGANG   CardsGroupType = 4
	CardsGroupType_CGT_BUGANG   CardsGroupType = 5
	CardsGroupType_CGT_HU       CardsGroupType = 6
	CardsGroupType_CGT_HUA      CardsGroupType = 7
	CardsGroupType_CGT_OUT      CardsGroupType = 8
)

// 杠牌数据
type GangCard struct {
	Card                 *Card    `protobuf:"bytes,1,opt,name=card,proto3" json:"card,omitempty"`
	Type                 GangType `protobuf:"varint,2,opt,name=type,proto3,enum=majong.GangType" json:"type,omitempty"`
	SrcPlayer            uint64   `protobuf:"varint,3,opt,name=src_player,json=srcPlayer,proto3" json:"src_player,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GangCard) GetCard() *Card {
	if m != nil {
		return m.Card
	}
	return nil
}

func (m *GangCard) GetType() GangType {
	if m != nil {
		return m.Type
	}
	return GangType_gang_angang
}

func (m *GangCard) GetSrcPlayer() uint64 {
	if m != nil {
		return m.SrcPlayer
	}
	return 0
}

// 碰牌数据
type PengCard struct {
	Card                 *Card    `protobuf:"bytes,1,opt,name=card,proto3" json:"card,omitempty"`
	SrcPlayer            uint64   `protobuf:"varint,2,opt,name=src_player,json=srcPlayer,proto3" json:"src_player,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PengCard) GetCard() *Card {
	if m != nil {
		return m.Card
	}
	return nil
}

func (m *PengCard) GetSrcPlayer() uint64 {
	if m != nil {
		return m.SrcPlayer
	}
	return 0
}

// 玩家胡牌对应的fanType
type HuFanType struct {
	HuCard               *Card    `protobuf:"bytes,1,opt,name=hu_card,json=huCard,proto3" json:"hu_card,omitempty"`
	GenCount             uint64   `protobuf:"varint,2,opt,name=gen_count,json=genCount,proto3" json:"gen_count,omitempty"`
	HuaCount             uint64   `protobuf:"varint,3,opt,name=hua_count,json=huaCount,proto3" json:"hua_count,omitempty"`
	FanTypes             []int64  `protobuf:"varint,4,rep,packed,name=fan_types,json=fanTypes,proto3" json:"fan_types,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HuFanType) GetHuCard() *Card {
	if m != nil {
		return m.HuCard
	}
	return nil
}

func (m *HuFanType) GetGenCount() uint64 {
	if m != nil {
		return m.GenCount
	}
	return 0
}

func (m *HuFanType) GetHuaCount() uint64 {
	if m != nil {
		return m.HuaCount
	}
	return 0
}

func (m *HuFanType) GetFanTypes() []int64 {
	if m != nil {
		return m.FanTypes
	}
	return nil
}

// 胡牌数据
type HuCard struct {
	Card                 *Card    `protobuf:"bytes,1,opt,name=card,proto3" json:"card,omitempty"`
	SrcPlayer            uint64   `protobuf:"varint,2,opt,name=src_player,json=srcPlayer,proto3" json:"src_player,omitempty"`
	Type                 HuType   `protobuf:"varint,3,opt,name=type,proto3,enum=majong.HuType" json:"type,omitempty"`
	IsReal               bool     `protobuf:"varint,4,opt,name=is_real,json=isReal,proto3" json:"is_real,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HuCard) GetCard() *Card {
	if m != nil {
		return m.Card
	}
	return nil
}

func (m *HuCard) GetSrcPlayer() uint64 {
	if m != nil {
		return m.SrcPlayer
	}
	return 0
}

func (m *HuCard) GetType() HuType {
	if m != nil {
		return m.Type
	}
	return HuType_hu_ganghoupao
}

func (m *HuCard) GetIsReal() bool {
	if m != nil {
		return m.IsReal
	}
	return false
}

// 吃牌数据
type ChiCard struct {
	Card                 *Card    `protobuf:"bytes,1,opt,name=card,proto3" json:"card,omitempty"`
	OprCard              *Card    `protobuf:"bytes,2,opt,name=opr_card,json=oprCard,proto3" json:"opr_card,omitempty"`
	SrcPlayer            uint64   `protobuf:"varint,3,opt,name=src_player,json=srcPlayer,proto3" json:"src_player,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChiCard) GetCard() *Card {
	if m != nil {
		return m.Card
	}
	return nil
}

func (m *ChiCard) GetOprCard() *Card {
	if m != nil {
		return m.OprCard
	}
	return nil
}

func (m *ChiCard) GetSrcPlayer() uint64 {
	if m != nil {
		return m.SrcPlayer
	}
	return 0
}

// TingCardInfo 听牌信息
type TingCardInfo struct {
	TingCard             uint32   `protobuf:"varint,1,opt,name=ting_card,json=tingCard,proto3" json:"ting_card,omitempty"`
	Times                uint32   `protobuf:"varint,2,opt,name=times,proto3" json:"times,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TingCardInfo) GetTingCard() uint32 {
	if m != nil {
		return m.TingCard
	}
	return 0
}

func (m *TingCardInfo) GetTimes() uint32 {
	if m != nil {
		return m.Times
	}
	return 0
}

// CanTingCardInfo 出本张牌可以听
type CanTingCardInfo struct {
	OutCard              uint32          `protobuf:"varint,1,opt,name=out_card,json=outCard,proto3" json:"out_card,omitempty"`
	TingCardInfo         []*TingCardInfo `protobuf:"bytes,2,rep,name=ting_card_info,json=tingCardInfo,proto3" json:"ting_card_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *CanTingCardInfo) GetOutCard() uint32 {
	if m != nil {
		return m.OutCard
	}
	return 0
}

func (m *CanTingCardInfo) GetTingCardInfo() []*TingCardInfo {
	if m != nil {
		return m.TingCardInfo
	}
	return nil
}

// ZiXunRecord 记录
type ZiXunRecord struct {
	EnableAngangCards []uint32           `protobuf:"varint,1,rep,packed,name=enable_angang_cards,json=enableAngangCards" json:"enable_angang_cards,omitempty"`
	EnableBugangCards []uint32           `protobuf:"varint,2,rep,packed,name=enable_bugang_cards,json=enableBugangCards" json:"enable_bugang_cards,omitempty"`
	EnableZimo        bool               `protobuf:"varint,3,opt,name=enable_zimo,json=enableZimo" json:"enable_zimo,omitempty"`
	EnableChupaiCards []uint32           `protobuf:"varint,4,rep,packed,name=enable_chupai_cards,json=enableChupaiCards" json:"enable_chupai_cards,omitempty"`
	CanTingCardInfo   []*CanTingCardInfo `protobuf:"bytes,5,rep,name=can_ting_card_info,json=canTingCardInfo" json:"can_ting_card_info,omitempty"`
	EnableQi          bool               `protobuf:"varint,6,opt,name=enable_qi,json=enableQi" json:"enable_qi,omitempty"`
	HuType            HuType             `protobuf:"varint,7,opt,name=hu_type,json=huType,enum=majong.HuType" json:"hu_type,omitempty"`
	EnableTing        bool               `protobuf:"varint,8,opt,name=enable_ting,json=enableTing" json:"enable_ting,omitempty"`
	TingType          TingType           `protobuf:"varint,9,opt,name=ting_type,json=tingType,enum=majong.TingType" json:"ting_type,omitempty"`
	HuFanType         *HuFanType         `protobuf:"bytes,10,opt,name=hu_fan_type,json=huFanType" json:"hu_fan_type,omitempty"`
}

func (m *ZiXunRecord) GetEnableAngangCards() []uint32 {
	if m != nil {
		return m.EnableAngangCards
	}
	return nil
}

func (m *ZiXunRecord) GetEnableBugangCards() []uint32 {
	if m != nil {
		return m.EnableBugangCards
	}
	return nil
}

func (m *ZiXunRecord) GetEnableZimo() bool {
	if m != nil {
		return m.EnableZimo
	}
	return false
}

func (m *ZiXunRecord) GetEnableChupaiCards() []uint32 {
	if m != nil {
		return m.EnableChupaiCards
	}
	return nil
}

func (m *ZiXunRecord) GetCanTingCardInfo() []*CanTingCardInfo {
	if m != nil {
		return m.CanTingCardInfo
	}
	return nil
}

func (m *ZiXunRecord) GetEnableQi() bool {
	if m != nil {
		return m.EnableQi
	}
	return false
}

func (m *ZiXunRecord) GetHuType() HuType {
	if m != nil {
		return m.HuType
	}
	return HuType_hu_ganghoupao
}

func (m *ZiXunRecord) GetEnableTing() bool {
	if m != nil {
		return m.EnableTing
	}
	return false
}

func (m *ZiXunRecord) GetTingType() TingType {
	if m != nil {
		return m.TingType
	}
	return TingType_TT_NORMAL_TING
}

func (m *ZiXunRecord) GetHuFanType() *HuFanType {
	if m != nil {
		return m.HuFanType
	}
	return nil
}

// Player 玩家数据
type Player struct {
	PlayerId          uint64            `protobuf:"varint,1,opt,name=player_id,json=playerId" json:"player_id,omitempty"`
	HandCards         []*Card           `protobuf:"bytes,2,rep,name=hand_cards,json=handCards" json:"hand_cards,omitempty"`
	OutCards          []*Card           `protobuf:"bytes,3,rep,name=out_cards,json=outCards" json:"out_cards,omitempty"`
	ChiCards          []*ChiCard        `protobuf:"bytes,4,rep,name=chi_cards,json=chiCards" json:"chi_cards,omitempty"`
	PengCards         []*PengCard       `protobuf:"bytes,5,rep,name=peng_cards,json=pengCards" json:"peng_cards,omitempty"`
	HuCards           []*HuCard         `protobuf:"bytes,6,rep,name=hu_cards,json=huCards" json:"hu_cards,omitempty"`
	GangCards         []*GangCard       `protobuf:"bytes,7,rep,name=gang_cards,json=gangCards" json:"gang_cards,omitempty"`
	PossibleActions   []Action          `protobuf:"varint,8,rep,packed,name=possible_actions,json=possibleActions,enum=majong.Action" json:"possible_actions,omitempty"`
	HasSelected       bool              `protobuf:"varint,9,opt,name=has_selected,json=hasSelected" json:"has_selected,omitempty"`
	SelectedAction    Action            `protobuf:"varint,10,opt,name=selected_action,json=selectedAction,enum=majong.Action" json:"selected_action,omitempty"`
	HasDingque        bool              `protobuf:"varint,11,opt,name=has_dingque,json=hasDingque" json:"has_dingque,omitempty"`
	DingqueColor      CardColor         `protobuf:"varint,12,opt,name=dingque_color,json=dingqueColor,enum=majong.CardColor" json:"dingque_color,omitempty"`
	HuansanzhangSure  bool              `protobuf:"varint,13,opt,name=huansanzhang_sure,json=huansanzhangSure" json:"huansanzhang_sure,omitempty"`
	HuansanzhangCards []*Card           `protobuf:"bytes,14,rep,name=huansanzhang_cards,json=huansanzhangCards" json:"huansanzhang_cards,omitempty"`
	MopaiCount        int32             `protobuf:"varint,15,opt,name=mopai_count,json=mopaiCount" json:"mopai_count,omitempty"`
	ZixunRecord       *ZiXunRecord      `protobuf:"bytes,16,opt,name=zixun_record,json=zixunRecord" json:"zixun_record,omitempty"`
	XpState           XingPaiState      `protobuf:"varint,17,opt,name=xp_state,json=xpState,enum=majong.XingPaiState" json:"xp_state,omitempty"`
	IsQuit            bool              `protobuf:"varint,18,opt,name=is_quit,json=isQuit" json:"is_quit,omitempty"`
	TingCardInfo      []*TingCardInfo   `protobuf:"bytes,19,rep,name=ting_card_info,json=tingCardInfo" json:"ting_card_info,omitempty"`
	HuaCards          []*Card           `protobuf:"bytes,20,rep,name=hua_cards,json=huaCards" json:"hua_cards,omitempty"`
	TingStateInfo     *TingStateInfo    `protobuf:"bytes,21,opt,name=ting_state_info,json=tingStateInfo" json:"ting_state_info,omitempty"`
	ZixunCount        int32             `protobuf:"varint,22,opt,name=zixun_count,json=zixunCount" json:"zixun_count,omitempty"`
	EnbleChiCards     []uint32          `protobuf:"varint,23,rep,packed,name=enble_chi_cards,json=enbleChiCards" json:"enble_chi_cards,omitempty"`
	DesignChiCards    []*Card           `protobuf:"bytes,24,rep,name=design_chi_cards,json=designChiCards" json:"design_chi_cards,omitempty"`
	MaxCardValue      uint32            `protobuf:"varint,25,opt,name=max_card_value,json=maxCardValue" json:"max_card_value,omitempty"`
	CardsGroup        []*CardsGroup     `protobuf:"bytes,26,rep,name=cards_group,json=cardsGroup" json:"cards_group,omitempty"`
	SelectedTing      bool              `protobuf:"varint,27,opt,name=selected_ting,json=selectedTing" json:"selected_ting,omitempty"`
	ChupaiCount       int32             `protobuf:"varint,28,opt,name=chupai_count,json=chupaiCount" json:"chupai_count,omitempty"`
	Properties        map[string][]byte `protobuf:"bytes,256,rep,name=properties" json:"properties,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *Player) GetPlayerId() uint64 {
	if m != nil {
		return m.PlayerId
	}
	return 0
}

func (m *Player) GetHandCards() []*Card {
	if m != nil {
		return m.HandCards
	}
	return nil
}

func (m *Player) GetOutCards() []*Card {
	if m != nil {
		return m.OutCards
	}
	return nil
}

func (m *Player) GetChiCards() []*ChiCard {
	if m != nil {
		return m.ChiCards
	}
	return nil
}

func (m *Player) GetPengCards() []*PengCard {
	if m != nil {
		return m.PengCards
	}
	return nil
}

func (m *Player) GetHuCards() []*HuCard {
	if m != nil {
		return m.HuCards
	}
	return nil
}

func (m *Player) GetGangCards() []*GangCard {
	if m != nil {
		return m.GangCards
	}
	return nil
}

func (m *Player) GetPossibleActions() []Action {
	if m != nil {
		return m.PossibleActions
	}
	return nil
}

func (m *Player) GetHasSelected() bool {
	if m != nil {
		return m.HasSelected
	}
	return false
}

func (m *Player) GetSelectedAction() Action {
	if m != nil {
		return m.SelectedAction
	}
	return Action_action_peng
}

func (m *Player) GetHasDingque() bool {
	if m != nil {
		return m.HasDingque
	}
	return false
}

func (m *Player) GetDingqueColor() CardColor {
	if m != nil {
		return m.DingqueColor
	}
	return CardColor_ColorWan
}

func (m *Player) GetHuansanzhangSure() bool {
	if m != nil {
		return m.HuansanzhangSure
	}
	return false
}

func (m *Player) GetHuansanzhangCards() []*Card {
	if m != nil {
		return m.HuansanzhangCards
	}
	return nil
}

func (m *Player) GetMopaiCount() int32 {
	if m != nil {
		return m.MopaiCount
	}
	return 0
}

func (m *Player) GetZixunRecord() *ZiXunRecord {
	if m != nil {
		return m.ZixunRecord
	}
	return nil
}

func (m *Player) GetXpState() XingPaiState {
	if m != nil {
		return m.XpState
	}
	return XingPaiState_normal
}

func (m *Player) GetIsQuit() bool {
	if m != nil {
		return m.IsQuit
	}
	return false
}

func (m *Player) GetTingCardInfo() []*TingCardInfo {
	if m != nil {
		return m.TingCardInfo
	}
	return nil
}

func (m *Player) GetHuaCards() []*Card {
	if m != nil {
		return m.HuaCards
	}
	return nil
}

func (m *Player) GetTingStateInfo() *TingStateInfo {
	if m != nil {
		return m.TingStateInfo
	}
	return nil
}

func (m *Player) GetZixunCount() int32 {
	if m != nil {
		return m.ZixunCount
	}
	return 0
}

func (m *Player) GetEnbleChiCards() []uint32 {
	if m != nil {
		return m.EnbleChiCards
	}
	return nil
}

func (m *Player) GetDesignChiCards() []*Card {
	if m != nil {
		return m.DesignChiCards
	}
	return nil
}

func (m *Player) GetMaxCardValue() uint32 {
	if m != nil {
		return m.MaxCardValue
	}
	return 0
}

func (m *Player) GetCardsGroup() []*CardsGroup {
	if m != nil {
		return m.CardsGroup
	}
	return nil
}

func (m *Player) GetSelectedTing() bool {
	if m != nil {
		return m.SelectedTing
	}
	return false
}

func (m *Player) GetChupaiCount() int32 {
	if m != nil {
		return m.ChupaiCount
	}
	return 0
}

func (m *Player) GetProperties() map[string][]byte {
	if m != nil {
		return m.Properties
	}
	return nil
}

type TingStateInfo struct {
	IsTing               bool     `protobuf:"varint,1,opt,name=is_ting,json=isTing,proto3" json:"is_ting,omitempty"`
	IsTianting           bool     `protobuf:"varint,2,opt,name=is_tianting,json=isTianting,proto3" json:"is_tianting,omitempty"`
	BaotingyifaCount     int32    `protobuf:"varint,3,opt,name=baotingyifa_count,json=baotingyifaCount,proto3" json:"baotingyifa_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TingStateInfo) GetIsTing() bool {
	if m != nil {
		return m.IsTing
	}
	return false
}

func (m *TingStateInfo) GetIsTianting() bool {
	if m != nil {
		return m.IsTianting
	}
	return false
}

func (m *TingStateInfo) GetBaotingyifaCount() int32 {
	if m != nil {
		return m.BaotingyifaCount
	}
	return 0
}

// AutoEvent 自动事件
type AutoEvent struct {
	EventId      EventID     `protobuf:"varint,1,opt,name=event_id,json=eventId,enum=majong.EventID" json:"event_id,omitempty"`
	EventContext interface{} `protobuf:"bytes,2,opt,name=event_context,json=eventContext,proto3" json:"event_context,omitempty"`
	WaitTime     uint32      `protobuf:"varint,3,opt,name=wait_time,json=waitTime" json:"wait_time,omitempty"`
}

func (m *AutoEvent) GetEventId() EventID {
	if m != nil {
		return m.EventId
	}
	return EventID_event_invalid
}

func (m *AutoEvent) GetWaitTime() uint32 {
	if m != nil {
		return m.WaitTime
	}
	return 0
}

func (m *AutoEvent) GetEventContext() interface{} {
	if m != nil {
		return m.EventContext
	}
	return nil
}

// SettleInfo 结算信息
type SettleInfo struct {
	Id           uint64           `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Scores       map[uint64]int64 `protobuf:"bytes,2,rep,name=scores" json:"scores,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	SettleType   SettleType       `protobuf:"varint,3,opt,name=settle_type,json=settleType,enum=majong.SettleType" json:"settle_type,omitempty"`
	HuType       HuType           `protobuf:"varint,4,opt,name=hu_type,json=huType,enum=majong.HuType" json:"hu_type,omitempty"`
	CardType     []int64          `protobuf:"varint,5,rep,packed,name=card_type,json=cardType" json:"card_type,omitempty"`
	CardValue    uint32           `protobuf:"varint,6,opt,name=card_value,json=cardValue" json:"card_value,omitempty"`
	GenCount     uint32           `protobuf:"varint,7,opt,name=gen_count,json=genCount" json:"gen_count,omitempty"`
	HuaCount     uint32           `protobuf:"varint,8,opt,name=hua_count,json=huaCount" json:"hua_count,omitempty"`
	CallTransfer bool             `protobuf:"varint,9,opt,name=call_transfer,json=callTransfer" json:"call_transfer,omitempty"`
	GroupId      []uint64         `protobuf:"varint,10,rep,packed,name=group_id,json=groupId" json:"group_id,omitempty"`
	HuPlayers    []uint64         `protobuf:"varint,11,rep,packed,name=hu_players,json=huPlayers" json:"hu_players,omitempty"`
}

func (m *SettleInfo) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SettleInfo) GetScores() map[uint64]int64 {
	if m != nil {
		return m.Scores
	}
	return nil
}

func (m *SettleInfo) GetSettleType() SettleType {
	if m != nil {
		return m.SettleType
	}
	return SettleType_settle_angang
}

func (m *SettleInfo) GetHuType() HuType {
	if m != nil {
		return m.HuType
	}
	return HuType_hu_ganghoupao
}

func (m *SettleInfo) GetCardType() []int64 {
	if m != nil {
		return m.CardType
	}
	return nil
}

func (m *SettleInfo) GetCardValue() uint32 {
	if m != nil {
		return m.CardValue
	}
	return 0
}

func (m *SettleInfo) GetGenCount() uint32 {
	if m != nil {
		return m.GenCount
	}
	return 0
}

func (m *SettleInfo) GetHuaCount() uint32 {
	if m != nil {
		return m.HuaCount
	}
	return 0
}

func (m *SettleInfo) GetCallTransfer() bool {
	if m != nil {
		return m.CallTransfer
	}
	return false
}

func (m *SettleInfo) GetGroupId() []uint64 {
	if m != nil {
		return m.GroupId
	}
	return nil
}

func (m *SettleInfo) GetHuPlayers() []uint64 {
	if m != nil {
		return m.HuPlayers
	}
	return nil
}

// MajongContext 麻将现场
type MajongContext struct {
	GameId          int32         `protobuf:"varint,1,opt,name=game_id,json=gameId,proto3" json:"game_id,omitempty"`
	CurState        StateID       `protobuf:"varint,2,opt,name=cur_state,json=curState,proto3,enum=majong.StateID" json:"cur_state,omitempty"`
	Players         []*Player     `protobuf:"bytes,3,rep,name=players,proto3" json:"players,omitempty"`
	ActivePlayer    uint64        `protobuf:"varint,4,opt,name=active_player,json=activePlayer,proto3" json:"active_player,omitempty"`
	WallCards       []*Card       `protobuf:"bytes,5,rep,name=wall_cards,json=wallCards,proto3" json:"wall_cards,omitempty"`
	SettleInfos     []*SettleInfo `protobuf:"bytes,6,rep,name=settle_infos,json=settleInfos,proto3" json:"settle_infos,omitempty"`
	CurrentSettleId uint64        `protobuf:"varint,7,opt,name=current_settleId,json=currentSettleId,proto3" json:"current_settleId,omitempty"`
	RevertSettles   []uint64      `protobuf:"varint,8,rep,packed,name=revert_settles,json=revertSettles,proto3" json:"revert_settles,omitempty"`
	LastOutCard     *Card         `protobuf:"bytes,9,opt,name=last_out_card,json=lastOutCard,proto3" json:"last_out_card,omitempty"`
	ZhuangjiaIndex  uint32        `protobuf:"varint,10,opt,name=zhuangjia_index,json=zhuangjiaIndex,proto3" json:"zhuangjia_index,omitempty"`
	// bool fix_zhuangjia_index = 11;  // 是否固定庄家位置
	LastHuPlayers       []uint64            `protobuf:"varint,12,rep,packed,name=last_hu_players,json=lastHuPlayers" json:"last_hu_players,omitempty"`
	LastPengPlayer      uint64              `protobuf:"varint,13,opt,name=last_peng_player,json=lastPengPlayer" json:"last_peng_player,omitempty"`
	LastGangPlayer      uint64              `protobuf:"varint,14,opt,name=last_gang_player,json=lastGangPlayer" json:"last_gang_player,omitempty"`
	LastChupaiPlayer    uint64              `protobuf:"varint,15,opt,name=last_chupai_player,json=lastChupaiPlayer" json:"last_chupai_player,omitempty"`
	MopaiPlayer         uint64              `protobuf:"varint,16,opt,name=mopai_player,json=mopaiPlayer" json:"mopai_player,omitempty"`
	LastMopaiPlayer     uint64              `protobuf:"varint,17,opt,name=last_mopai_player,json=lastMopaiPlayer" json:"last_mopai_player,omitempty"`
	LastMopaiCard       *Card               `protobuf:"bytes,18,opt,name=last_mopai_card,json=lastMopaiCard" json:"last_mopai_card,omitempty"`
	GangCard            *Card               `protobuf:"bytes,19,opt,name=gang_card,json=gangCard" json:"gang_card,omitempty"`
	MopaiType           MopaiType           `protobuf:"varint,20,opt,name=mopai_type,json=mopaiType,enum=majong.MopaiType" json:"mopai_type,omitempty"`
	ZixunType           ZixunType           `protobuf:"varint,21,opt,name=zixun_type,json=zixunType,enum=majong.ZixunType" json:"zixun_type,omitempty"`
	Dices               []uint32            `protobuf:"varint,22,rep,packed,name=dices" json:"dices,omitempty"`
	ExcutedHuansanzhang bool                `protobuf:"varint,23,opt,name=excuted_huansanzhang,json=excutedHuansanzhang" json:"excuted_huansanzhang,omitempty"`
	CardTotalNum        uint32              `protobuf:"varint,24,opt,name=card_total_num,json=cardTotalNum" json:"card_total_num,omitempty"`
	LastChiPlayer       uint64              `protobuf:"varint,25,opt,name=last_chi_player,json=lastChiPlayer" json:"last_chi_player,omitempty"`
	NextBankerSeat      uint32              `protobuf:"varint,26,opt,name=next_banker_seat,json=nextBankerSeat" json:"next_banker_seat,omitempty"`
	FixNextBankerSeat   bool                `protobuf:"varint,27,opt,name=fix_next_banker_seat,json=fixNextBankerSeat" json:"fix_next_banker_seat,omitempty"`
	XingpaiOptionId     uint32              `protobuf:"varint,251,opt,name=xingpai_option_id,json=xingpaiOptionId" json:"xingpai_option_id,omitempty"`
	CardtypeOptionId    uint32              `protobuf:"varint,252,opt,name=cardtype_option_id,json=cardtypeOptionId" json:"cardtype_option_id,omitempty"`
	SettleOptionId      uint32              `protobuf:"varint,253,opt,name=settle_option_id,json=settleOptionId" json:"settle_option_id,omitempty"`
	Option              *MajongCommonOption `protobuf:"bytes,254,opt,name=option" json:"option,omitempty"`
	MajongOption        []byte              `protobuf:"bytes,255,opt,name=majong_option,json=majongOption,proto3" json:"majong_option,omitempty"`
	TempData            *TempDatas          `protobuf:"bytes,256,opt,name=TempData" json:"TempData,omitempty"`
	GameStartTime       time.Time
}

func (m *MajongContext) GetGameId() int32 {
	if m != nil {
		return m.GameId
	}
	return 0
}

func (m *MajongContext) GetCurState() StateID {
	if m != nil {
		return m.CurState
	}
	return StateID_state_init
}

func (m *MajongContext) GetPlayers() []*Player {
	if m != nil {
		return m.Players
	}
	return nil
}

func (m *MajongContext) GetActivePlayer() uint64 {
	if m != nil {
		return m.ActivePlayer
	}
	return 0
}

func (m *MajongContext) GetWallCards() []*Card {
	if m != nil {
		return m.WallCards
	}
	return nil
}

func (m *MajongContext) GetSettleInfos() []*SettleInfo {
	if m != nil {
		return m.SettleInfos
	}
	return nil
}

func (m *MajongContext) GetCurrentSettleId() uint64 {
	if m != nil {
		return m.CurrentSettleId
	}
	return 0
}

func (m *MajongContext) GetRevertSettles() []uint64 {
	if m != nil {
		return m.RevertSettles
	}
	return nil
}

func (m *MajongContext) GetLastOutCard() *Card {
	if m != nil {
		return m.LastOutCard
	}
	return nil
}

func (m *MajongContext) GetZhuangjiaIndex() uint32 {
	if m != nil {
		return m.ZhuangjiaIndex
	}
	return 0
}

func (m *MajongContext) GetLastHuPlayers() []uint64 {
	if m != nil {
		return m.LastHuPlayers
	}
	return nil
}

func (m *MajongContext) GetLastPengPlayer() uint64 {
	if m != nil {
		return m.LastPengPlayer
	}
	return 0
}

func (m *MajongContext) GetLastGangPlayer() uint64 {
	if m != nil {
		return m.LastGangPlayer
	}
	return 0
}

func (m *MajongContext) GetLastChupaiPlayer() uint64 {
	if m != nil {
		return m.LastChupaiPlayer
	}
	return 0
}

func (m *MajongContext) GetMopaiPlayer() uint64 {
	if m != nil {
		return m.MopaiPlayer
	}
	return 0
}

func (m *MajongContext) GetLastMopaiPlayer() uint64 {
	if m != nil {
		return m.LastMopaiPlayer
	}
	return 0
}

func (m *MajongContext) GetLastMopaiCard() *Card {
	if m != nil {
		return m.LastMopaiCard
	}
	return nil
}

func (m *MajongContext) GetGangCard() *Card {
	if m != nil {
		return m.GangCard
	}
	return nil
}

func (m *MajongContext) GetMopaiType() MopaiType {
	if m != nil {
		return m.MopaiType
	}
	return MopaiType_MT_NORMAL
}

func (m *MajongContext) GetZixunType() ZixunType {
	if m != nil {
		return m.ZixunType
	}
	return ZixunType_ZXT_NORMAL
}

func (m *MajongContext) GetDices() []uint32 {
	if m != nil {
		return m.Dices
	}
	return nil
}

func (m *MajongContext) GetExcutedHuansanzhang() bool {
	if m != nil {
		return m.ExcutedHuansanzhang
	}
	return false
}

func (m *MajongContext) GetCardTotalNum() uint32 {
	if m != nil {
		return m.CardTotalNum
	}
	return 0
}

func (m *MajongContext) GetLastChiPlayer() uint64 {
	if m != nil {
		return m.LastChiPlayer
	}
	return 0
}

func (m *MajongContext) GetNextBankerSeat() uint32 {
	if m != nil {
		return m.NextBankerSeat
	}
	return 0
}

func (m *MajongContext) GetFixNextBankerSeat() bool {
	if m != nil {
		return m.FixNextBankerSeat
	}
	return false
}

func (m *MajongContext) GetXingpaiOptionId() uint32 {
	if m != nil {
		return m.XingpaiOptionId
	}
	return 0
}

func (m *MajongContext) GetCardtypeOptionId() uint32 {
	if m != nil {
		return m.CardtypeOptionId
	}
	return 0
}

func (m *MajongContext) GetSettleOptionId() uint32 {
	if m != nil {
		return m.SettleOptionId
	}
	return 0
}

func (m *MajongContext) GetOption() *MajongCommonOption {
	if m != nil {
		return m.Option
	}
	return nil
}

func (m *MajongContext) GetMajongOption() []byte {
	if m != nil {
		return m.MajongOption
	}
	return nil
}

func (m *MajongContext) GetTempData() *TempDatas {
	if m != nil {
		return m.TempData
	}
	return nil
}

func (m *MajongContext) GetGameStartTime() time.Time {
	if m != nil {
		return m.GameStartTime
	}
	return time.Time{}
}

// MajongCommonOption 麻将通用玩法选项
type MajongCommonOption struct {
	MaxCartoonTime             uint32          `protobuf:"varint,1,opt,name=max_cartoon_time,json=maxCartoonTime,proto3" json:"max_cartoon_time,omitempty"`
	MaxFapaiCartoonTime        uint32          `protobuf:"varint,2,opt,name=max_fapai_cartoon_time,json=maxFapaiCartoonTime,proto3" json:"max_fapai_cartoon_time,omitempty"`
	MaxHuansanzhangCartoonTime uint32          `protobuf:"varint,3,opt,name=max_huansanzhang_cartoon_time,json=maxHuansanzhangCartoonTime,proto3" json:"max_huansanzhang_cartoon_time,omitempty"`
	HasHuansanzhang            bool            `protobuf:"varint,4,opt,name=has_huansanzhang,json=hasHuansanzhang,proto3" json:"has_huansanzhang,omitempty"`
	Cards                      string          `protobuf:"bytes,5,opt,name=cards,proto3" json:"cards,omitempty"`
	WallcardsLength            uint32          `protobuf:"varint,6,opt,name=wallcards_length,json=wallcardsLength,proto3" json:"wallcards_length,omitempty"`
	HszFx                      *Huansanzhangfx `protobuf:"bytes,7,opt,name=hsz_fx,json=hszFx,proto3" json:"hsz_fx,omitempty"`
	XXX_NoUnkeyedLiteral       struct{}        `json:"-"`
	XXX_unrecognized           []byte          `json:"-"`
	XXX_sizecache              int32           `json:"-"`
}

func (m *MajongCommonOption) GetMaxCartoonTime() uint32 {
	if m != nil {
		return m.MaxCartoonTime
	}
	return 0
}

func (m *MajongCommonOption) GetMaxFapaiCartoonTime() uint32 {
	if m != nil {
		return m.MaxFapaiCartoonTime
	}
	return 0
}

func (m *MajongCommonOption) GetMaxHuansanzhangCartoonTime() uint32 {
	if m != nil {
		return m.MaxHuansanzhangCartoonTime
	}
	return 0
}

func (m *MajongCommonOption) GetHasHuansanzhang() bool {
	if m != nil {
		return m.HasHuansanzhang
	}
	return false
}

func (m *MajongCommonOption) GetCards() string {
	if m != nil {
		return m.Cards
	}
	return ""
}

func (m *MajongCommonOption) GetWallcardsLength() uint32 {
	if m != nil {
		return m.WallcardsLength
	}
	return 0
}

func (m *MajongCommonOption) GetHszFx() *Huansanzhangfx {
	if m != nil {
		return m.HszFx
	}
	return nil
}

type Huansanzhangfx struct {
	NeedDeployFx         bool     `protobuf:"varint,1,opt,name=need_deploy_fx,json=needDeployFx,proto3" json:"need_deploy_fx,omitempty"`
	HuansanzhangFx       int32    `protobuf:"varint,2,opt,name=huansanzhang_fx,json=huansanzhangFx,proto3" json:"huansanzhang_fx,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Huansanzhangfx) GetNeedDeployFx() bool {
	if m != nil {
		return m.NeedDeployFx
	}
	return false
}

func (m *Huansanzhangfx) GetHuansanzhangFx() int32 {
	if m != nil {
		return m.HuansanzhangFx
	}
	return 0
}

// SichuanxueliuOption 四川血流麻将玩法
type SichuanxueliuOption struct {
	OpenHuansanzhang     bool     `protobuf:"varint,1,opt,name=open_huansanzhang,json=openHuansanzhang,proto3" json:"open_huansanzhang,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SichuanxueliuOption) GetOpenHuansanzhang() bool {
	if m != nil {
		return m.OpenHuansanzhang
	}
	return false
}

// SichuanxuezhanOption 四川血流麻将玩法
type SichuanxuezhanOption struct {
	OpenHuansanzhang     bool     `protobuf:"varint,1,opt,name=open_huansanzhang,json=openHuansanzhang,proto3" json:"open_huansanzhang,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SichuanxuezhanOption) GetOpenHuansanzhang() bool {
	if m != nil {
		return m.OpenHuansanzhang
	}
	return false
}

// InitMajongContextParams 麻将现场初始化参数
type InitMajongContextParams struct {
	GameId               int32               `protobuf:"varint,1,opt,name=game_id,json=gameId,proto3" json:"game_id,omitempty"`
	Players              []uint64            `protobuf:"varint,2,rep,packed,name=players,proto3" json:"players,omitempty"`
	Option               *MajongCommonOption `protobuf:"bytes,3,opt,name=option,proto3" json:"option,omitempty"`
	MajongOption         []byte              `protobuf:"bytes,4,opt,name=majong_option,json=majongOption,proto3" json:"majong_option,omitempty"`
	ZhuangIndex          uint32              `protobuf:"varint,5,opt,name=zhuang_index,json=zhuangIndex,proto3" json:"zhuang_index,omitempty"`
	FixZhuangIndex       bool                `protobuf:"varint,6,opt,name=fix_zhuang_index,json=fixZhuangIndex,proto3" json:"fix_zhuang_index,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *InitMajongContextParams) GetGameId() int32 {
	if m != nil {
		return m.GameId
	}
	return 0
}

func (m *InitMajongContextParams) GetPlayers() []uint64 {
	if m != nil {
		return m.Players
	}
	return nil
}

func (m *InitMajongContextParams) GetOption() *MajongCommonOption {
	if m != nil {
		return m.Option
	}
	return nil
}

func (m *InitMajongContextParams) GetMajongOption() []byte {
	if m != nil {
		return m.MajongOption
	}
	return nil
}

func (m *InitMajongContextParams) GetZhuangIndex() uint32 {
	if m != nil {
		return m.ZhuangIndex
	}
	return 0
}

func (m *InitMajongContextParams) GetFixZhuangIndex() bool {
	if m != nil {
		return m.FixZhuangIndex
	}
	return false
}

// ReplyClientMessage 回复给客户端的消息
type ReplyClientMessage struct {
	Players              []uint64 `protobuf:"varint,1,rep,packed,name=players,proto3" json:"players,omitempty"`
	MsgId                int32    `protobuf:"varint,2,opt,name=msg_id,json=msgId,proto3" json:"msg_id,omitempty"`
	Msg                  []byte   `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReplyClientMessage) GetPlayers() []uint64 {
	if m != nil {
		return m.Players
	}
	return nil
}

func (m *ReplyClientMessage) GetMsgId() int32 {
	if m != nil {
		return m.MsgId
	}
	return 0
}

func (m *ReplyClientMessage) GetMsg() []byte {
	if m != nil {
		return m.Msg
	}
	return nil
}

// TimeCheckInfo 时间检测信息
type TimeCheckInfo struct {
	Duration             uint64   `protobuf:"varint,1,opt,name=duration,proto3" json:"duration,omitempty"`
	Type                 int32    `protobuf:"varint,2,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TimeCheckInfo) GetDuration() uint64 {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *TimeCheckInfo) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

// 麻将组
type CardsGroup struct {
	Cards                []uint32       `protobuf:"varint,1,rep,packed,name=cards,proto3" json:"cards,omitempty"`
	Type                 CardsGroupType `protobuf:"varint,2,opt,name=type,proto3,enum=majong.CardsGroupType" json:"type,omitempty"`
	Pid                  uint64         `protobuf:"varint,3,opt,name=pid,proto3" json:"pid,omitempty"`
	IsReal               bool           `protobuf:"varint,4,opt,name=is_real,json=isReal,proto3" json:"is_real,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CardsGroup) GetCards() []uint32 {
	if m != nil {
		return m.Cards
	}
	return nil
}

func (m *CardsGroup) GetType() CardsGroupType {
	if m != nil {
		return m.Type
	}
	return CardsGroupType_CGT_HAND
}

func (m *CardsGroup) GetPid() uint64 {
	if m != nil {
		return m.Pid
	}
	return 0
}

func (m *CardsGroup) GetIsReal() bool {
	if m != nil {
		return m.IsReal
	}
	return false
}

// TempDatas 临时数据存储
type TempDatas struct {
	CartoonReqPlayerIDs  []uint64 `protobuf:"varint,1,rep,packed,name=CartoonReqPlayerIDs,proto3" json:"CartoonReqPlayerIDs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TempDatas) GetCartoonReqPlayerIDs() []uint64 {
	if m != nil {
		return m.CartoonReqPlayerIDs
	}
	return nil
}
