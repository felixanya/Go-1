// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common.proto

/*
Package common is a generated protocol buffer package.

It is generated from these files:
	common.proto
	errors.proto

It has these top-level messages:
	GeographicalLocation
	GameConfig
	GameLevelConfig
	Property
	Money
	Result
*/
package common

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// GameId 游戏 ID
type GameId int32

const (
	GameId_GAMEID_XUELIU   GameId = 1
	GameId_GAMEID_XUEZHAN  GameId = 2
	GameId_GAMEID_DOUDIZHU GameId = 3
	GameId_GAMEID_ERRENMJ  GameId = 4
)

var GameId_name = map[int32]string{
	1: "GAMEID_XUELIU",
	2: "GAMEID_XUEZHAN",
	3: "GAMEID_DOUDIZHU",
	4: "GAMEID_ERRENMJ",
}
var GameId_value = map[string]int32{
	"GAMEID_XUELIU":   1,
	"GAMEID_XUEZHAN":  2,
	"GAMEID_DOUDIZHU": 3,
	"GAMEID_ERRENMJ":  4,
}

func (x GameId) Enum() *GameId {
	p := new(GameId)
	*p = x
	return p
}
func (x GameId) String() string {
	return proto.EnumName(GameId_name, int32(x))
}
func (x *GameId) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(GameId_value, data, "GameId")
	if err != nil {
		return err
	}
	*x = GameId(value)
	return nil
}
func (GameId) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// LocSourceType 位置信息来源类型
type LocSourceType int32

const (
	LocSourceType_LOC_SOURCE_BAIDU  LocSourceType = 0
	LocSourceType_LOC_SOURCE_JIZHAN LocSourceType = 1
)

var LocSourceType_name = map[int32]string{
	0: "LOC_SOURCE_BAIDU",
	1: "LOC_SOURCE_JIZHAN",
}
var LocSourceType_value = map[string]int32{
	"LOC_SOURCE_BAIDU":  0,
	"LOC_SOURCE_JIZHAN": 1,
}

func (x LocSourceType) Enum() *LocSourceType {
	p := new(LocSourceType)
	*p = x
	return p
}
func (x LocSourceType) String() string {
	return proto.EnumName(LocSourceType_name, int32(x))
}
func (x *LocSourceType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(LocSourceType_value, data, "LocSourceType")
	if err != nil {
		return err
	}
	*x = LocSourceType(value)
	return nil
}
func (LocSourceType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// PlayerState 玩家状态
type PlayerState int32

const (
	PlayerState_PS_IDLE     PlayerState = 1
	PlayerState_PS_GAMEING  PlayerState = 2
	PlayerState_PS_MATCHING PlayerState = 3
)

var PlayerState_name = map[int32]string{
	1: "PS_IDLE",
	2: "PS_GAMEING",
	3: "PS_MATCHING",
}
var PlayerState_value = map[string]int32{
	"PS_IDLE":     1,
	"PS_GAMEING":  2,
	"PS_MATCHING": 3,
}

func (x PlayerState) Enum() *PlayerState {
	p := new(PlayerState)
	*p = x
	return p
}
func (x PlayerState) String() string {
	return proto.EnumName(PlayerState_name, int32(x))
}
func (x *PlayerState) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(PlayerState_value, data, "PlayerState")
	if err != nil {
		return err
	}
	*x = PlayerState(value)
	return nil
}
func (PlayerState) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// PlayerGender 玩家性别
type PlayerGender int32

const (
	PlayerGender_PG_NIL    PlayerGender = 0
	PlayerGender_PG_MALE   PlayerGender = 1
	PlayerGender_PG_FEMALE PlayerGender = 2
)

var PlayerGender_name = map[int32]string{
	0: "PG_NIL",
	1: "PG_MALE",
	2: "PG_FEMALE",
}
var PlayerGender_value = map[string]int32{
	"PG_NIL":    0,
	"PG_MALE":   1,
	"PG_FEMALE": 2,
}

func (x PlayerGender) Enum() *PlayerGender {
	p := new(PlayerGender)
	*p = x
	return p
}
func (x PlayerGender) String() string {
	return proto.EnumName(PlayerGender_name, int32(x))
}
func (x *PlayerGender) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(PlayerGender_value, data, "PlayerGender")
	if err != nil {
		return err
	}
	*x = PlayerGender(value)
	return nil
}
func (PlayerGender) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

// Platform 平台
type Platform int32

const (
	Platform_Android Platform = 1
	Platform_Iphone  Platform = 2
)

var Platform_name = map[int32]string{
	1: "Android",
	2: "Iphone",
}
var Platform_value = map[string]int32{
	"Android": 1,
	"Iphone":  2,
}

func (x Platform) Enum() *Platform {
	p := new(Platform)
	*p = x
	return p
}
func (x Platform) String() string {
	return proto.EnumName(Platform_name, int32(x))
}
func (x *Platform) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Platform_value, data, "Platform")
	if err != nil {
		return err
	}
	*x = Platform(value)
	return nil
}
func (Platform) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

// LevelTag 场次标签
type LevelTag int32

const (
	LevelTag_LT_HOT   LevelTag = 1
	LevelTag_LT_NEW   LevelTag = 2
	LevelTag_LT_OTHER LevelTag = 3
)

var LevelTag_name = map[int32]string{
	1: "LT_HOT",
	2: "LT_NEW",
	3: "LT_OTHER",
}
var LevelTag_value = map[string]int32{
	"LT_HOT":   1,
	"LT_NEW":   2,
	"LT_OTHER": 3,
}

func (x LevelTag) Enum() *LevelTag {
	p := new(LevelTag)
	*p = x
	return p
}
func (x LevelTag) String() string {
	return proto.EnumName(LevelTag_name, int32(x))
}
func (x *LevelTag) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(LevelTag_value, data, "LevelTag")
	if err != nil {
		return err
	}
	*x = LevelTag(value)
	return nil
}
func (LevelTag) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

// PropType 互动道具类型
type PropType int32

const (
	PropType_INVALID_PROP PropType = 0
	PropType_ROSE         PropType = 1
	PropType_BEER         PropType = 2
	PropType_BOMB         PropType = 3
	PropType_GRAB_CHICKEN PropType = 4
	PropType_EGG_GUN      PropType = 5
	PropType_VOUCHER      PropType = 6
	PropType_GAME         PropType = 7
)

var PropType_name = map[int32]string{
	0: "INVALID_PROP",
	1: "ROSE",
	2: "BEER",
	3: "BOMB",
	4: "GRAB_CHICKEN",
	5: "EGG_GUN",
	6: "VOUCHER",
	7: "GAME",
}
var PropType_value = map[string]int32{
	"INVALID_PROP": 0,
	"ROSE":         1,
	"BEER":         2,
	"BOMB":         3,
	"GRAB_CHICKEN": 4,
	"EGG_GUN":      5,
	"VOUCHER":      6,
	"GAME":         7,
}

func (x PropType) Enum() *PropType {
	p := new(PropType)
	*p = x
	return p
}
func (x PropType) String() string {
	return proto.EnumName(PropType_name, int32(x))
}
func (x *PropType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(PropType_value, data, "PropType")
	if err != nil {
		return err
	}
	*x = PropType(value)
	return nil
}
func (PropType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

// MoneyType 货币类型
type MoneyType int32

const (
	MoneyType_MT_INVAID    MoneyType = 0
	MoneyType_MT_COIN      MoneyType = 1
	MoneyType_MT_DIAMOND   MoneyType = 2
	MoneyType_MT_GOLDINGOT MoneyType = 3
)

var MoneyType_name = map[int32]string{
	0: "MT_INVAID",
	1: "MT_COIN",
	2: "MT_DIAMOND",
	3: "MT_GOLDINGOT",
}
var MoneyType_value = map[string]int32{
	"MT_INVAID":    0,
	"MT_COIN":      1,
	"MT_DIAMOND":   2,
	"MT_GOLDINGOT": 3,
}

func (x MoneyType) Enum() *MoneyType {
	p := new(MoneyType)
	*p = x
	return p
}
func (x MoneyType) String() string {
	return proto.EnumName(MoneyType_name, int32(x))
}
func (x *MoneyType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MoneyType_value, data, "MoneyType")
	if err != nil {
		return err
	}
	*x = MoneyType(value)
	return nil
}
func (MoneyType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

// GeographicalLocation 玩家地理位置
type GeographicalLocation struct {
	Type             *LocSourceType `protobuf:"varint,1,opt,name=type,enum=common.LocSourceType" json:"type,omitempty"`
	Longitude        *float64       `protobuf:"fixed64,2,opt,name=longitude" json:"longitude,omitempty"`
	Latitude         *float64       `protobuf:"fixed64,3,opt,name=latitude" json:"latitude,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *GeographicalLocation) Reset()                    { *m = GeographicalLocation{} }
func (m *GeographicalLocation) String() string            { return proto.CompactTextString(m) }
func (*GeographicalLocation) ProtoMessage()               {}
func (*GeographicalLocation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GeographicalLocation) GetType() LocSourceType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return LocSourceType_LOC_SOURCE_BAIDU
}

func (m *GeographicalLocation) GetLongitude() float64 {
	if m != nil && m.Longitude != nil {
		return *m.Longitude
	}
	return 0
}

func (m *GeographicalLocation) GetLatitude() float64 {
	if m != nil && m.Latitude != nil {
		return *m.Latitude
	}
	return 0
}

// GameLevelConfig 游戏玩法
type GameConfig struct {
	GameId           *uint32 `protobuf:"varint,1,opt,name=game_id,json=gameId" json:"game_id,omitempty"`
	GameName         *string `protobuf:"bytes,2,opt,name=game_name,json=gameName" json:"game_name,omitempty"`
	GameType         *uint32 `protobuf:"varint,3,opt,name=game_type,json=gameType" json:"game_type,omitempty"`
	MinPeople        *uint32 `protobuf:"varint,4,opt,name=min_people,json=minPeople" json:"min_people,omitempty"`
	MaxPeople        *uint32 `protobuf:"varint,5,opt,name=max_people,json=maxPeople" json:"max_people,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GameConfig) Reset()                    { *m = GameConfig{} }
func (m *GameConfig) String() string            { return proto.CompactTextString(m) }
func (*GameConfig) ProtoMessage()               {}
func (*GameConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GameConfig) GetGameId() uint32 {
	if m != nil && m.GameId != nil {
		return *m.GameId
	}
	return 0
}

func (m *GameConfig) GetGameName() string {
	if m != nil && m.GameName != nil {
		return *m.GameName
	}
	return ""
}

func (m *GameConfig) GetGameType() uint32 {
	if m != nil && m.GameType != nil {
		return *m.GameType
	}
	return 0
}

func (m *GameConfig) GetMinPeople() uint32 {
	if m != nil && m.MinPeople != nil {
		return *m.MinPeople
	}
	return 0
}

func (m *GameConfig) GetMaxPeople() uint32 {
	if m != nil && m.MaxPeople != nil {
		return *m.MaxPeople
	}
	return 0
}

// GameLevelConfig 游戏等级
type GameLevelConfig struct {
	GameId           *uint32   `protobuf:"varint,1,opt,name=game_id,json=gameId" json:"game_id,omitempty"`
	LevelId          *uint32   `protobuf:"varint,2,opt,name=level_id,json=levelId" json:"level_id,omitempty"`
	LevelName        *string   `protobuf:"bytes,3,opt,name=level_name,json=levelName" json:"level_name,omitempty"`
	BaseScores       *uint32   `protobuf:"varint,4,opt,name=base_scores,json=baseScores" json:"base_scores,omitempty"`
	LowScores        *uint32   `protobuf:"varint,5,opt,name=low_scores,json=lowScores" json:"low_scores,omitempty"`
	HighScors        *uint32   `protobuf:"varint,6,opt,name=high_scors,json=highScors" json:"high_scors,omitempty"`
	ShowPeople       *uint32   `protobuf:"varint,7,opt,name=show_people,json=showPeople" json:"show_people,omitempty"`
	RealPeople       *uint32   `protobuf:"varint,8,opt,name=real_people,json=realPeople" json:"real_people,omitempty"`
	LevelTag         *LevelTag `protobuf:"varint,9,opt,name=level_tag,json=levelTag,enum=common.LevelTag" json:"level_tag,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *GameLevelConfig) Reset()                    { *m = GameLevelConfig{} }
func (m *GameLevelConfig) String() string            { return proto.CompactTextString(m) }
func (*GameLevelConfig) ProtoMessage()               {}
func (*GameLevelConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GameLevelConfig) GetGameId() uint32 {
	if m != nil && m.GameId != nil {
		return *m.GameId
	}
	return 0
}

func (m *GameLevelConfig) GetLevelId() uint32 {
	if m != nil && m.LevelId != nil {
		return *m.LevelId
	}
	return 0
}

func (m *GameLevelConfig) GetLevelName() string {
	if m != nil && m.LevelName != nil {
		return *m.LevelName
	}
	return ""
}

func (m *GameLevelConfig) GetBaseScores() uint32 {
	if m != nil && m.BaseScores != nil {
		return *m.BaseScores
	}
	return 0
}

func (m *GameLevelConfig) GetLowScores() uint32 {
	if m != nil && m.LowScores != nil {
		return *m.LowScores
	}
	return 0
}

func (m *GameLevelConfig) GetHighScors() uint32 {
	if m != nil && m.HighScors != nil {
		return *m.HighScors
	}
	return 0
}

func (m *GameLevelConfig) GetShowPeople() uint32 {
	if m != nil && m.ShowPeople != nil {
		return *m.ShowPeople
	}
	return 0
}

func (m *GameLevelConfig) GetRealPeople() uint32 {
	if m != nil && m.RealPeople != nil {
		return *m.RealPeople
	}
	return 0
}

func (m *GameLevelConfig) GetLevelTag() LevelTag {
	if m != nil && m.LevelTag != nil {
		return *m.LevelTag
	}
	return LevelTag_LT_HOT
}

// Property 互动道具
type Property struct {
	PropId           *int32    `protobuf:"varint,1,opt,name=prop_id,json=propId" json:"prop_id,omitempty"`
	PropName         *string   `protobuf:"bytes,2,opt,name=prop_name,json=propName" json:"prop_name,omitempty"`
	PropType         *PropType `protobuf:"varint,3,opt,name=prop_type,json=propType,enum=common.PropType" json:"prop_type,omitempty"`
	PropCount        *uint32   `protobuf:"varint,4,opt,name=prop_count,json=propCount" json:"prop_count,omitempty"`
	PropCost         *int64    `protobuf:"varint,5,opt,name=prop_cost,json=propCost" json:"prop_cost,omitempty"`
	PropLimit        *int64    `protobuf:"varint,6,opt,name=prop_limit,json=propLimit" json:"prop_limit,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *Property) Reset()                    { *m = Property{} }
func (m *Property) String() string            { return proto.CompactTextString(m) }
func (*Property) ProtoMessage()               {}
func (*Property) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Property) GetPropId() int32 {
	if m != nil && m.PropId != nil {
		return *m.PropId
	}
	return 0
}

func (m *Property) GetPropName() string {
	if m != nil && m.PropName != nil {
		return *m.PropName
	}
	return ""
}

func (m *Property) GetPropType() PropType {
	if m != nil && m.PropType != nil {
		return *m.PropType
	}
	return PropType_INVALID_PROP
}

func (m *Property) GetPropCount() uint32 {
	if m != nil && m.PropCount != nil {
		return *m.PropCount
	}
	return 0
}

func (m *Property) GetPropCost() int64 {
	if m != nil && m.PropCost != nil {
		return *m.PropCost
	}
	return 0
}

func (m *Property) GetPropLimit() int64 {
	if m != nil && m.PropLimit != nil {
		return *m.PropLimit
	}
	return 0
}

// Money 货币
type Money struct {
	MoneyType        *MoneyType `protobuf:"varint,1,opt,name=money_type,json=moneyType,enum=common.MoneyType" json:"money_type,omitempty"`
	MoneyNum         *uint64    `protobuf:"varint,2,opt,name=money_num,json=moneyNum" json:"money_num,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *Money) Reset()                    { *m = Money{} }
func (m *Money) String() string            { return proto.CompactTextString(m) }
func (*Money) ProtoMessage()               {}
func (*Money) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Money) GetMoneyType() MoneyType {
	if m != nil && m.MoneyType != nil {
		return *m.MoneyType
	}
	return MoneyType_MT_INVAID
}

func (m *Money) GetMoneyNum() uint64 {
	if m != nil && m.MoneyNum != nil {
		return *m.MoneyNum
	}
	return 0
}

func init() {
	proto.RegisterType((*GeographicalLocation)(nil), "common.GeographicalLocation")
	proto.RegisterType((*GameConfig)(nil), "common.GameConfig")
	proto.RegisterType((*GameLevelConfig)(nil), "common.GameLevelConfig")
	proto.RegisterType((*Property)(nil), "common.Property")
	proto.RegisterType((*Money)(nil), "common.Money")
	proto.RegisterEnum("common.GameId", GameId_name, GameId_value)
	proto.RegisterEnum("common.LocSourceType", LocSourceType_name, LocSourceType_value)
	proto.RegisterEnum("common.PlayerState", PlayerState_name, PlayerState_value)
	proto.RegisterEnum("common.PlayerGender", PlayerGender_name, PlayerGender_value)
	proto.RegisterEnum("common.Platform", Platform_name, Platform_value)
	proto.RegisterEnum("common.LevelTag", LevelTag_name, LevelTag_value)
	proto.RegisterEnum("common.PropType", PropType_name, PropType_value)
	proto.RegisterEnum("common.MoneyType", MoneyType_name, MoneyType_value)
}

func init() { proto.RegisterFile("common.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 856 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0x4b, 0x8e, 0xdb, 0x46,
	0x10, 0x35, 0x49, 0x8d, 0x86, 0xaa, 0xf9, 0xf5, 0x74, 0xec, 0x44, 0xf9, 0xc1, 0x86, 0xb2, 0x71,
	0x04, 0xf8, 0x83, 0x2c, 0xb2, 0x49, 0x36, 0x12, 0xc9, 0x50, 0x74, 0xf8, 0x03, 0x45, 0x4d, 0x8c,
	0xd9, 0x34, 0x68, 0xa9, 0x2d, 0x11, 0x20, 0xd9, 0x04, 0x45, 0x79, 0x2c, 0xe4, 0x2c, 0xb9, 0x45,
	0x4e, 0x91, 0x53, 0x05, 0x5d, 0x2d, 0x6a, 0x94, 0x55, 0x76, 0xac, 0xf7, 0x5e, 0x55, 0xbd, 0xaa,
	0x2e, 0xc2, 0xe5, 0x52, 0x94, 0xa5, 0xa8, 0x5e, 0xd7, 0x8d, 0x68, 0x05, 0xed, 0xab, 0x68, 0xf4,
	0x27, 0x3c, 0x75, 0xb9, 0x58, 0x37, 0x59, 0xbd, 0xc9, 0x97, 0x59, 0xe1, 0x8b, 0x65, 0xd6, 0xe6,
	0xa2, 0xa2, 0x3f, 0x42, 0xaf, 0xdd, 0xd7, 0x7c, 0xa8, 0xbd, 0xd0, 0x5e, 0x5e, 0xff, 0xf4, 0xec,
	0xf5, 0x21, 0xd9, 0x17, 0xcb, 0xb9, 0xd8, 0x35, 0x4b, 0x9e, 0xee, 0x6b, 0x9e, 0xa0, 0x84, 0x7e,
	0x07, 0x83, 0x42, 0x54, 0xeb, 0xbc, 0xdd, 0xad, 0xf8, 0x50, 0x7f, 0xa1, 0xbd, 0xd4, 0x92, 0x47,
	0x80, 0x7e, 0x03, 0x66, 0x91, 0xb5, 0x8a, 0x34, 0x90, 0x3c, 0xc6, 0xa3, 0xbf, 0x34, 0x00, 0x37,
	0x2b, 0xb9, 0x25, 0xaa, 0x8f, 0xf9, 0x9a, 0x7e, 0x05, 0xe7, 0xeb, 0xac, 0xe4, 0x2c, 0x5f, 0x61,
	0xdb, 0xab, 0xa4, 0x2f, 0x43, 0x6f, 0x45, 0xbf, 0x85, 0x01, 0x12, 0x55, 0x56, 0xaa, 0x0e, 0x83,
	0xc4, 0x94, 0x40, 0x98, 0x95, 0xfc, 0x48, 0xa2, 0x5d, 0x03, 0xf3, 0x90, 0x94, 0x0e, 0xe9, 0xf7,
	0x00, 0x65, 0x5e, 0xb1, 0x9a, 0x8b, 0xba, 0xe0, 0xc3, 0x1e, 0xb2, 0x83, 0x32, 0xaf, 0x62, 0x04,
	0x90, 0xce, 0x3e, 0x77, 0xf4, 0xd9, 0x81, 0xce, 0x3e, 0x2b, 0x7a, 0xf4, 0xb7, 0x0e, 0x37, 0xd2,
	0x9f, 0xcf, 0x3f, 0xf1, 0xe2, 0xff, 0x4c, 0x7e, 0x0d, 0x66, 0x21, 0x75, 0x92, 0xd1, 0x91, 0x39,
	0xc7, 0xd8, 0x5b, 0xc9, 0x36, 0x8a, 0xc2, 0x01, 0x0c, 0x1c, 0x60, 0x80, 0x08, 0x4e, 0xf0, 0x1c,
	0x2e, 0x3e, 0x64, 0x5b, 0xce, 0xb6, 0x4b, 0xd1, 0xf0, 0xed, 0xc1, 0x25, 0x48, 0x68, 0x8e, 0x08,
	0xe6, 0x8b, 0x87, 0x8e, 0x3f, 0xd8, 0x2c, 0xc4, 0xc3, 0x23, 0xbd, 0xc9, 0xd7, 0x1b, 0xe4, 0xb7,
	0xc3, 0xbe, 0xa2, 0x25, 0x22, 0xf9, 0xad, 0x2c, 0xbf, 0xdd, 0x88, 0x87, 0x6e, 0xca, 0x73, 0x55,
	0x5e, 0x42, 0x87, 0x2d, 0x3c, 0x87, 0x8b, 0x86, 0x67, 0x45, 0x27, 0x30, 0x95, 0x40, 0x42, 0x07,
	0xc1, 0x2b, 0x50, 0x6e, 0x59, 0x9b, 0xad, 0x87, 0x03, 0xbc, 0x08, 0x72, 0xbc, 0x08, 0x49, 0xa4,
	0xd9, 0x3a, 0x51, 0xd3, 0xa7, 0xd9, 0x7a, 0xf4, 0x8f, 0x06, 0x66, 0xdc, 0x88, 0x9a, 0x37, 0xed,
	0x5e, 0xee, 0xab, 0x6e, 0x44, 0xdd, 0xed, 0xeb, 0x2c, 0xe9, 0xcb, 0x50, 0x3d, 0x2a, 0x12, 0xa7,
	0x8f, 0x2a, 0x01, 0x5c, 0xc9, 0xab, 0x03, 0x79, 0x7c, 0xd4, 0x93, 0x8e, 0xb2, 0x34, 0x9e, 0x1f,
	0xca, 0xbb, 0x67, 0x46, 0xf9, 0x52, 0xec, 0xaa, 0xb6, 0x7b, 0x66, 0x89, 0x58, 0x12, 0x38, 0xb6,
	0x5a, 0x8a, 0x6d, 0x8b, 0xeb, 0x33, 0x54, 0xae, 0x25, 0xb6, 0xed, 0x31, 0xb7, 0xc8, 0xcb, 0xbc,
	0xc5, 0xed, 0x19, 0x2a, 0xd7, 0x97, 0xc0, 0xe8, 0x0e, 0xce, 0x02, 0x51, 0xf1, 0x3d, 0x7d, 0x0b,
	0x50, 0xca, 0x0f, 0x76, 0xf2, 0x5f, 0xdc, 0x76, 0x9e, 0x50, 0x82, 0xa6, 0x06, 0x65, 0xf7, 0x29,
	0xdb, 0xaa, 0x8c, 0x6a, 0x57, 0xe2, 0x84, 0xbd, 0xc4, 0x44, 0x20, 0xdc, 0x95, 0xe3, 0xf7, 0xd0,
	0x77, 0xd5, 0xe1, 0xdc, 0xc2, 0x95, 0x3b, 0x09, 0x1c, 0xcf, 0x66, 0xef, 0x17, 0x8e, 0xef, 0x2d,
	0x88, 0x46, 0x29, 0x5c, 0x3f, 0x42, 0xf7, 0xb3, 0x49, 0x48, 0x74, 0xfa, 0x05, 0xdc, 0x1c, 0x30,
	0x3b, 0x5a, 0xd8, 0xde, 0xfd, 0x6c, 0x41, 0x8c, 0x13, 0xa1, 0x93, 0x24, 0x4e, 0x18, 0xbc, 0x23,
	0xbd, 0xf1, 0xaf, 0x70, 0xf5, 0x9f, 0xdf, 0x94, 0x3e, 0x05, 0xe2, 0x47, 0x16, 0x9b, 0x47, 0x8b,
	0xc4, 0x72, 0xd8, 0x74, 0xe2, 0xd9, 0x0b, 0xf2, 0x84, 0x3e, 0x83, 0xdb, 0x13, 0xf4, 0x9d, 0x87,
	0x6d, 0xb4, 0xf1, 0x2f, 0x70, 0x11, 0x17, 0xd9, 0x9e, 0x37, 0xf3, 0x36, 0x6b, 0x39, 0xbd, 0x80,
	0xf3, 0x78, 0xce, 0x3c, 0xdb, 0x77, 0x88, 0x46, 0xaf, 0x01, 0xe2, 0x39, 0xc3, 0x86, 0xa1, 0x4b,
	0x74, 0x7a, 0x03, 0x17, 0xf1, 0x9c, 0x05, 0x93, 0xd4, 0x9a, 0x49, 0xc0, 0x18, 0xff, 0x0c, 0x97,
	0x2a, 0xd9, 0xe5, 0xd5, 0x8a, 0x37, 0x14, 0xa0, 0x1f, 0xbb, 0x2c, 0xf4, 0x7c, 0xf2, 0x04, 0x2b,
	0xb9, 0x2c, 0x98, 0x60, 0xa5, 0x2b, 0x18, 0xc4, 0x2e, 0xfb, 0xcd, 0xc1, 0x50, 0x1f, 0xff, 0x00,
	0x66, 0x5c, 0x64, 0xed, 0x47, 0xd1, 0x94, 0x52, 0x37, 0xa9, 0x56, 0x8d, 0xc8, 0x57, 0x44, 0x93,
	0x05, 0xbc, 0x7a, 0x23, 0x2a, 0x4e, 0xf4, 0xf1, 0x5b, 0x30, 0xbb, 0x63, 0x93, 0xb8, 0x9f, 0xb2,
	0x59, 0x94, 0x2a, 0x8d, 0x9f, 0xb2, 0xd0, 0xf9, 0x83, 0xe8, 0xf4, 0x12, 0x4c, 0x3f, 0x65, 0x51,
	0x3a, 0x73, 0x12, 0x62, 0x8c, 0x6b, 0x75, 0x87, 0xb8, 0x04, 0x02, 0x97, 0x5e, 0x78, 0x37, 0xf1,
	0x3d, 0x9b, 0xc5, 0x49, 0x14, 0x93, 0x27, 0xd4, 0x84, 0x5e, 0x12, 0xcd, 0xa5, 0x1b, 0x13, 0x7a,
	0x53, 0xc7, 0x49, 0x88, 0x8e, 0x5f, 0x51, 0x30, 0x25, 0x86, 0xd4, 0xbb, 0xc9, 0x64, 0xca, 0xac,
	0x99, 0x67, 0xfd, 0xee, 0x84, 0xa4, 0x27, 0x8d, 0x39, 0xae, 0xcb, 0xdc, 0x45, 0x48, 0xce, 0x64,
	0x70, 0x17, 0x2d, 0x2c, 0xd9, 0xa7, 0x2f, 0xb3, 0xe4, 0x52, 0xc8, 0xf9, 0xd8, 0x83, 0xc1, 0xf1,
	0x14, 0xe4, 0x90, 0x41, 0xca, 0x64, 0x57, 0xcf, 0x56, 0x0b, 0x08, 0x52, 0x66, 0x45, 0x5e, 0xa8,
	0x56, 0x19, 0xa4, 0xcc, 0xf6, 0x26, 0x41, 0x14, 0xda, 0x44, 0x97, 0xed, 0x82, 0x94, 0xb9, 0x91,
	0x6f, 0x7b, 0xa1, 0x1b, 0xa5, 0xc4, 0x98, 0x0e, 0xef, 0xbf, 0xdc, 0xb6, 0xfc, 0x13, 0x7f, 0xb3,
	0x2c, 0x72, 0x5e, 0xb5, 0xac, 0xfe, 0xf0, 0x46, 0x1d, 0xdb, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x74, 0x2e, 0x9d, 0x9a, 0xca, 0x05, 0x00, 0x00,
}
