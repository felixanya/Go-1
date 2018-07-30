// Code generated by protoc-gen-go. DO NOT EDIT.
// source: game_ddz.proto

package room

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// DDZStage:当前游戏状态(恢复对局用)
type DDZStage int32

const (
	DDZStage_DDZ_STAGE_NONE    DDZStage = 0
	DDZStage_DDZ_STAGE_DEAL    DDZStage = 1
	DDZStage_DDZ_STAGE_CALL    DDZStage = 2
	DDZStage_DDZ_STAGE_GRAB    DDZStage = 3
	DDZStage_DDZ_STAGE_DOUBLE  DDZStage = 4
	DDZStage_DDZ_STAGE_PLAYING DDZStage = 5
	DDZStage_DDZ_STAGE_OVER    DDZStage = 6
)

var DDZStage_name = map[int32]string{
	0: "DDZ_STAGE_NONE",
	1: "DDZ_STAGE_DEAL",
	2: "DDZ_STAGE_CALL",
	3: "DDZ_STAGE_GRAB",
	4: "DDZ_STAGE_DOUBLE",
	5: "DDZ_STAGE_PLAYING",
	6: "DDZ_STAGE_OVER",
}
var DDZStage_value = map[string]int32{
	"DDZ_STAGE_NONE":    0,
	"DDZ_STAGE_DEAL":    1,
	"DDZ_STAGE_CALL":    2,
	"DDZ_STAGE_GRAB":    3,
	"DDZ_STAGE_DOUBLE":  4,
	"DDZ_STAGE_PLAYING": 5,
	"DDZ_STAGE_OVER":    6,
}

func (x DDZStage) Enum() *DDZStage {
	p := new(DDZStage)
	*p = x
	return p
}
func (x DDZStage) String() string {
	return proto.EnumName(DDZStage_name, int32(x))
}
func (x *DDZStage) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(DDZStage_value, data, "DDZStage")
	if err != nil {
		return err
	}
	*x = DDZStage(value)
	return nil
}
func (DDZStage) EnumDescriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

type CardType int32

const (
	CardType_CT_NONE     CardType = 0
	CardType_CT_SINGLE   CardType = 1
	CardType_CT_PAIR     CardType = 2
	CardType_CT_SHUNZI   CardType = 3
	CardType_CT_PAIRS    CardType = 4
	CardType_CT_TRIPLE   CardType = 5
	CardType_CT_3AND1    CardType = 6
	CardType_CT_3AND2    CardType = 7
	CardType_CT_TRIPLES  CardType = 8
	CardType_CT_3SAND1S  CardType = 9
	CardType_CT_3SAND2S  CardType = 10
	CardType_CT_4SAND1S  CardType = 11
	CardType_CT_4SAND2S  CardType = 12
	CardType_CT_BOMB     CardType = 13
	CardType_CT_KINGBOMB CardType = 14
)

var CardType_name = map[int32]string{
	0:  "CT_NONE",
	1:  "CT_SINGLE",
	2:  "CT_PAIR",
	3:  "CT_SHUNZI",
	4:  "CT_PAIRS",
	5:  "CT_TRIPLE",
	6:  "CT_3AND1",
	7:  "CT_3AND2",
	8:  "CT_TRIPLES",
	9:  "CT_3SAND1S",
	10: "CT_3SAND2S",
	11: "CT_4SAND1S",
	12: "CT_4SAND2S",
	13: "CT_BOMB",
	14: "CT_KINGBOMB",
}
var CardType_value = map[string]int32{
	"CT_NONE":     0,
	"CT_SINGLE":   1,
	"CT_PAIR":     2,
	"CT_SHUNZI":   3,
	"CT_PAIRS":    4,
	"CT_TRIPLE":   5,
	"CT_3AND1":    6,
	"CT_3AND2":    7,
	"CT_TRIPLES":  8,
	"CT_3SAND1S":  9,
	"CT_3SAND2S":  10,
	"CT_4SAND1S":  11,
	"CT_4SAND2S":  12,
	"CT_BOMB":     13,
	"CT_KINGBOMB": 14,
}

func (x CardType) Enum() *CardType {
	p := new(CardType)
	*p = x
	return p
}
func (x CardType) String() string {
	return proto.EnumName(CardType_name, int32(x))
}
func (x *CardType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(CardType_value, data, "CardType")
	if err != nil {
		return err
	}
	*x = CardType(value)
	return nil
}
func (CardType) EnumDescriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

// 叫/抢地主操作的类型
type GrabLordType int32

const (
	GrabLordType_GLT_NONE       GrabLordType = 0
	GrabLordType_GLT_CALLLORD   GrabLordType = 1
	GrabLordType_GLT_NOTCALLORD GrabLordType = 2
	GrabLordType_GLT_GRAB       GrabLordType = 3
	GrabLordType_GLT_NOTGRAB    GrabLordType = 4
)

var GrabLordType_name = map[int32]string{
	0: "GLT_NONE",
	1: "GLT_CALLLORD",
	2: "GLT_NOTCALLORD",
	3: "GLT_GRAB",
	4: "GLT_NOTGRAB",
}
var GrabLordType_value = map[string]int32{
	"GLT_NONE":       0,
	"GLT_CALLLORD":   1,
	"GLT_NOTCALLORD": 2,
	"GLT_GRAB":       3,
	"GLT_NOTGRAB":    4,
}

func (x GrabLordType) Enum() *GrabLordType {
	p := new(GrabLordType)
	*p = x
	return p
}
func (x GrabLordType) String() string {
	return proto.EnumName(GrabLordType_name, int32(x))
}
func (x *GrabLordType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(GrabLordType_value, data, "GrabLordType")
	if err != nil {
		return err
	}
	*x = GrabLordType(value)
	return nil
}
func (GrabLordType) EnumDescriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

// 加倍操作的类型
type DoubleType int32

const (
	DoubleType_DT_NONE      DoubleType = 0
	DoubleType_DT_DOUBLE    DoubleType = 1
	DoubleType_DT_NOTDOUBLE DoubleType = 2
)

var DoubleType_name = map[int32]string{
	0: "DT_NONE",
	1: "DT_DOUBLE",
	2: "DT_NOTDOUBLE",
}
var DoubleType_value = map[string]int32{
	"DT_NONE":      0,
	"DT_DOUBLE":    1,
	"DT_NOTDOUBLE": 2,
}

func (x DoubleType) Enum() *DoubleType {
	p := new(DoubleType)
	*p = x
	return p
}
func (x DoubleType) String() string {
	return proto.EnumName(DoubleType_name, int32(x))
}
func (x *DoubleType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(DoubleType_value, data, "DoubleType")
	if err != nil {
		return err
	}
	*x = DoubleType(value)
	return nil
}
func (DoubleType) EnumDescriptor() ([]byte, []int) { return fileDescriptor4, []int{3} }

// 通用操作结果，客户端直接弹出err_desc提示即可
type Result struct {
	ErrCode          *uint32 `protobuf:"varint,1,opt,name=err_code,json=errCode" json:"err_code,omitempty"`
	ErrDesc          *string `protobuf:"bytes,2,opt,name=err_desc,json=errDesc" json:"err_desc,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Result) Reset()                    { *m = Result{} }
func (m *Result) String() string            { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()               {}
func (*Result) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *Result) GetErrCode() uint32 {
	if m != nil && m.ErrCode != nil {
		return *m.ErrCode
	}
	return 0
}

func (m *Result) GetErrDesc() string {
	if m != nil && m.ErrDesc != nil {
		return *m.ErrDesc
	}
	return ""
}

// “Stage切换”可以独立一条消息，也可以嵌入到其他会发生阶段切换的消息体中
type NextStage struct {
	Stage            *DDZStage `protobuf:"varint,1,opt,name=stage,enum=room.DDZStage" json:"stage,omitempty"`
	Time             *uint32   `protobuf:"varint,2,opt,name=time" json:"time,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *NextStage) Reset()                    { *m = NextStage{} }
func (m *NextStage) String() string            { return proto.CompactTextString(m) }
func (*NextStage) ProtoMessage()               {}
func (*NextStage) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

func (m *NextStage) GetStage() DDZStage {
	if m != nil && m.Stage != nil {
		return *m.Stage
	}
	return DDZStage_DDZ_STAGE_NONE
}

func (m *NextStage) GetTime() uint32 {
	if m != nil && m.Time != nil {
		return *m.Time
	}
	return 0
}

// DDZStartGameNtf 斗地主游戏开始通知
// CMD: 0x15000
type DDZStartGameNtf struct {
	PlayerId         *uint64    `protobuf:"varint,1,opt,name=player_id,json=playerId" json:"player_id,omitempty"`
	NextStage        *NextStage `protobuf:"bytes,2,opt,name=next_stage,json=nextStage" json:"next_stage,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *DDZStartGameNtf) Reset()                    { *m = DDZStartGameNtf{} }
func (m *DDZStartGameNtf) String() string            { return proto.CompactTextString(m) }
func (*DDZStartGameNtf) ProtoMessage()               {}
func (*DDZStartGameNtf) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

func (m *DDZStartGameNtf) GetPlayerId() uint64 {
	if m != nil && m.PlayerId != nil {
		return *m.PlayerId
	}
	return 0
}

func (m *DDZStartGameNtf) GetNextStage() *NextStage {
	if m != nil {
		return m.NextStage
	}
	return nil
}

// DDZDealNtf 发牌通知
// CMD: 0x15001
type DDZDealNtf struct {
	Cards            []uint32   `protobuf:"varint,1,rep,name=cards" json:"cards,omitempty"`
	CallPlayerId     *uint64    `protobuf:"varint,2,opt,name=call_player_id,json=callPlayerId" json:"call_player_id,omitempty"`
	NextStage        *NextStage `protobuf:"bytes,3,opt,name=next_stage,json=nextStage" json:"next_stage,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *DDZDealNtf) Reset()                    { *m = DDZDealNtf{} }
func (m *DDZDealNtf) String() string            { return proto.CompactTextString(m) }
func (*DDZDealNtf) ProtoMessage()               {}
func (*DDZDealNtf) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{3} }

func (m *DDZDealNtf) GetCards() []uint32 {
	if m != nil {
		return m.Cards
	}
	return nil
}

func (m *DDZDealNtf) GetCallPlayerId() uint64 {
	if m != nil && m.CallPlayerId != nil {
		return *m.CallPlayerId
	}
	return 0
}

func (m *DDZDealNtf) GetNextStage() *NextStage {
	if m != nil {
		return m.NextStage
	}
	return nil
}

// DDZGrabLordReq 叫/抢地主请求，叫地主和抢地主用同一个请求
// CMD: 0x15002
type DDZGrabLordReq struct {
	Grab             *bool  `protobuf:"varint,1,opt,name=grab" json:"grab,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *DDZGrabLordReq) Reset()                    { *m = DDZGrabLordReq{} }
func (m *DDZGrabLordReq) String() string            { return proto.CompactTextString(m) }
func (*DDZGrabLordReq) ProtoMessage()               {}
func (*DDZGrabLordReq) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{4} }

func (m *DDZGrabLordReq) GetGrab() bool {
	if m != nil && m.Grab != nil {
		return *m.Grab
	}
	return false
}

// DDZGrabLordNtf 叫/抢地主广播
// CMD: 0x15004
type DDZGrabLordNtf struct {
	PlayerId         *uint64    `protobuf:"varint,1,opt,name=player_id,json=playerId" json:"player_id,omitempty"`
	Grab             *bool      `protobuf:"varint,2,opt,name=grab" json:"grab,omitempty"`
	TotalGrab        *uint32    `protobuf:"varint,3,opt,name=total_grab,json=totalGrab" json:"total_grab,omitempty"`
	NextPlayerId     *uint64    `protobuf:"varint,4,opt,name=next_player_id,json=nextPlayerId" json:"next_player_id,omitempty"`
	NextStage        *NextStage `protobuf:"bytes,5,opt,name=next_stage,json=nextStage" json:"next_stage,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *DDZGrabLordNtf) Reset()                    { *m = DDZGrabLordNtf{} }
func (m *DDZGrabLordNtf) String() string            { return proto.CompactTextString(m) }
func (*DDZGrabLordNtf) ProtoMessage()               {}
func (*DDZGrabLordNtf) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{5} }

func (m *DDZGrabLordNtf) GetPlayerId() uint64 {
	if m != nil && m.PlayerId != nil {
		return *m.PlayerId
	}
	return 0
}

func (m *DDZGrabLordNtf) GetGrab() bool {
	if m != nil && m.Grab != nil {
		return *m.Grab
	}
	return false
}

func (m *DDZGrabLordNtf) GetTotalGrab() uint32 {
	if m != nil && m.TotalGrab != nil {
		return *m.TotalGrab
	}
	return 0
}

func (m *DDZGrabLordNtf) GetNextPlayerId() uint64 {
	if m != nil && m.NextPlayerId != nil {
		return *m.NextPlayerId
	}
	return 0
}

func (m *DDZGrabLordNtf) GetNextStage() *NextStage {
	if m != nil {
		return m.NextStage
	}
	return nil
}

// DDZLordNtf 地主广播
// CMD: 0x15005
type DDZLordNtf struct {
	PlayerId         *uint64    `protobuf:"varint,1,opt,name=player_id,json=playerId" json:"player_id,omitempty"`
	TotalGrab        *uint32    `protobuf:"varint,2,opt,name=total_grab,json=totalGrab" json:"total_grab,omitempty"`
	Dipai            []uint32   `protobuf:"varint,3,rep,name=dipai" json:"dipai,omitempty"`
	NextStage        *NextStage `protobuf:"bytes,4,opt,name=next_stage,json=nextStage" json:"next_stage,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *DDZLordNtf) Reset()                    { *m = DDZLordNtf{} }
func (m *DDZLordNtf) String() string            { return proto.CompactTextString(m) }
func (*DDZLordNtf) ProtoMessage()               {}
func (*DDZLordNtf) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{6} }

func (m *DDZLordNtf) GetPlayerId() uint64 {
	if m != nil && m.PlayerId != nil {
		return *m.PlayerId
	}
	return 0
}

func (m *DDZLordNtf) GetTotalGrab() uint32 {
	if m != nil && m.TotalGrab != nil {
		return *m.TotalGrab
	}
	return 0
}

func (m *DDZLordNtf) GetDipai() []uint32 {
	if m != nil {
		return m.Dipai
	}
	return nil
}

func (m *DDZLordNtf) GetNextStage() *NextStage {
	if m != nil {
		return m.NextStage
	}
	return nil
}

// DDZDoubleReq 斗地主加倍请求
// CMD: 0x15006
type DDZDoubleReq struct {
	IsDouble         *bool  `protobuf:"varint,1,opt,name=is_double,json=isDouble" json:"is_double,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *DDZDoubleReq) Reset()                    { *m = DDZDoubleReq{} }
func (m *DDZDoubleReq) String() string            { return proto.CompactTextString(m) }
func (*DDZDoubleReq) ProtoMessage()               {}
func (*DDZDoubleReq) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{7} }

func (m *DDZDoubleReq) GetIsDouble() bool {
	if m != nil && m.IsDouble != nil {
		return *m.IsDouble
	}
	return false
}

// DDZDoubleNtf 加倍广播
// CMD: 0x15008
type DDZDoubleNtf struct {
	PlayerId         *uint64    `protobuf:"varint,1,opt,name=player_id,json=playerId" json:"player_id,omitempty"`
	IsDouble         *bool      `protobuf:"varint,2,opt,name=is_double,json=isDouble" json:"is_double,omitempty"`
	TotalDouble      *uint32    `protobuf:"varint,3,opt,name=total_double,json=totalDouble" json:"total_double,omitempty"`
	NextStage        *NextStage `protobuf:"bytes,4,opt,name=next_stage,json=nextStage" json:"next_stage,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *DDZDoubleNtf) Reset()                    { *m = DDZDoubleNtf{} }
func (m *DDZDoubleNtf) String() string            { return proto.CompactTextString(m) }
func (*DDZDoubleNtf) ProtoMessage()               {}
func (*DDZDoubleNtf) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{8} }

func (m *DDZDoubleNtf) GetPlayerId() uint64 {
	if m != nil && m.PlayerId != nil {
		return *m.PlayerId
	}
	return 0
}

func (m *DDZDoubleNtf) GetIsDouble() bool {
	if m != nil && m.IsDouble != nil {
		return *m.IsDouble
	}
	return false
}

func (m *DDZDoubleNtf) GetTotalDouble() uint32 {
	if m != nil && m.TotalDouble != nil {
		return *m.TotalDouble
	}
	return 0
}

func (m *DDZDoubleNtf) GetNextStage() *NextStage {
	if m != nil {
		return m.NextStage
	}
	return nil
}

// DDZPlayCardReq 出牌请求
// CMD: 0x15009
type DDZPlayCardReq struct {
	Cards            []uint32  `protobuf:"varint,1,rep,name=cards" json:"cards,omitempty"`
	CardType         *CardType `protobuf:"varint,2,opt,name=card_type,json=cardType,enum=room.CardType" json:"card_type,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *DDZPlayCardReq) Reset()                    { *m = DDZPlayCardReq{} }
func (m *DDZPlayCardReq) String() string            { return proto.CompactTextString(m) }
func (*DDZPlayCardReq) ProtoMessage()               {}
func (*DDZPlayCardReq) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{9} }

func (m *DDZPlayCardReq) GetCards() []uint32 {
	if m != nil {
		return m.Cards
	}
	return nil
}

func (m *DDZPlayCardReq) GetCardType() CardType {
	if m != nil && m.CardType != nil {
		return *m.CardType
	}
	return CardType_CT_NONE
}

// DDZPlayCardRsp 出牌响应
// CMD: 0x1500A
type DDZPlayCardRsp struct {
	Result           *Result `protobuf:"bytes,1,opt,name=result" json:"result,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *DDZPlayCardRsp) Reset()                    { *m = DDZPlayCardRsp{} }
func (m *DDZPlayCardRsp) String() string            { return proto.CompactTextString(m) }
func (*DDZPlayCardRsp) ProtoMessage()               {}
func (*DDZPlayCardRsp) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{10} }

func (m *DDZPlayCardRsp) GetResult() *Result {
	if m != nil {
		return m.Result
	}
	return nil
}

// DDZPlayCardNtf 出牌广播
// CMD: 0x1500B
type DDZPlayCardNtf struct {
	PlayerId         *uint64    `protobuf:"varint,1,opt,name=player_id,json=playerId" json:"player_id,omitempty"`
	Cards            []uint32   `protobuf:"varint,2,rep,name=cards" json:"cards,omitempty"`
	CardType         *CardType  `protobuf:"varint,3,opt,name=card_type,json=cardType,enum=room.CardType" json:"card_type,omitempty"`
	TotalBomb        *uint32    `protobuf:"varint,4,opt,name=total_bomb,json=totalBomb" json:"total_bomb,omitempty"`
	NextPlayerId     *uint64    `protobuf:"varint,5,opt,name=next_player_id,json=nextPlayerId" json:"next_player_id,omitempty"`
	NextStage        *NextStage `protobuf:"bytes,6,opt,name=next_stage,json=nextStage" json:"next_stage,omitempty"`
	CardTypePivot    *uint32    `protobuf:"varint,7,opt,name=card_type_pivot,json=cardTypePivot" json:"card_type_pivot,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *DDZPlayCardNtf) Reset()                    { *m = DDZPlayCardNtf{} }
func (m *DDZPlayCardNtf) String() string            { return proto.CompactTextString(m) }
func (*DDZPlayCardNtf) ProtoMessage()               {}
func (*DDZPlayCardNtf) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{11} }

func (m *DDZPlayCardNtf) GetPlayerId() uint64 {
	if m != nil && m.PlayerId != nil {
		return *m.PlayerId
	}
	return 0
}

func (m *DDZPlayCardNtf) GetCards() []uint32 {
	if m != nil {
		return m.Cards
	}
	return nil
}

func (m *DDZPlayCardNtf) GetCardType() CardType {
	if m != nil && m.CardType != nil {
		return *m.CardType
	}
	return CardType_CT_NONE
}

func (m *DDZPlayCardNtf) GetTotalBomb() uint32 {
	if m != nil && m.TotalBomb != nil {
		return *m.TotalBomb
	}
	return 0
}

func (m *DDZPlayCardNtf) GetNextPlayerId() uint64 {
	if m != nil && m.NextPlayerId != nil {
		return *m.NextPlayerId
	}
	return 0
}

func (m *DDZPlayCardNtf) GetNextStage() *NextStage {
	if m != nil {
		return m.NextStage
	}
	return nil
}

func (m *DDZPlayCardNtf) GetCardTypePivot() uint32 {
	if m != nil && m.CardTypePivot != nil {
		return *m.CardTypePivot
	}
	return 0
}

// BillPlayersInfo 结算玩家账单
type DDZBillPlayerInfo struct {
	PlayerId         *uint64  `protobuf:"varint,1,opt,name=player_id,json=playerId" json:"player_id,omitempty"`
	Win              *bool    `protobuf:"varint,2,opt,name=win" json:"win,omitempty"`
	Base             *int32   `protobuf:"varint,3,opt,name=base" json:"base,omitempty"`
	Multiple         *int32   `protobuf:"varint,4,opt,name=multiple" json:"multiple,omitempty"`
	Score            *int64   `protobuf:"varint,5,opt,name=score" json:"score,omitempty"`
	CurrentScore     *int64   `protobuf:"varint,6,opt,name=current_score,json=currentScore" json:"current_score,omitempty"`
	Lord             *bool    `protobuf:"varint,7,opt,name=lord" json:"lord,omitempty"`
	OutCards         []uint32 `protobuf:"varint,8,rep,name=out_cards,json=outCards" json:"out_cards,omitempty"`
	HandCards        []uint32 `protobuf:"varint,9,rep,name=hand_cards,json=handCards" json:"hand_cards,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *DDZBillPlayerInfo) Reset()                    { *m = DDZBillPlayerInfo{} }
func (m *DDZBillPlayerInfo) String() string            { return proto.CompactTextString(m) }
func (*DDZBillPlayerInfo) ProtoMessage()               {}
func (*DDZBillPlayerInfo) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{12} }

func (m *DDZBillPlayerInfo) GetPlayerId() uint64 {
	if m != nil && m.PlayerId != nil {
		return *m.PlayerId
	}
	return 0
}

func (m *DDZBillPlayerInfo) GetWin() bool {
	if m != nil && m.Win != nil {
		return *m.Win
	}
	return false
}

func (m *DDZBillPlayerInfo) GetBase() int32 {
	if m != nil && m.Base != nil {
		return *m.Base
	}
	return 0
}

func (m *DDZBillPlayerInfo) GetMultiple() int32 {
	if m != nil && m.Multiple != nil {
		return *m.Multiple
	}
	return 0
}

func (m *DDZBillPlayerInfo) GetScore() int64 {
	if m != nil && m.Score != nil {
		return *m.Score
	}
	return 0
}

func (m *DDZBillPlayerInfo) GetCurrentScore() int64 {
	if m != nil && m.CurrentScore != nil {
		return *m.CurrentScore
	}
	return 0
}

func (m *DDZBillPlayerInfo) GetLord() bool {
	if m != nil && m.Lord != nil {
		return *m.Lord
	}
	return false
}

func (m *DDZBillPlayerInfo) GetOutCards() []uint32 {
	if m != nil {
		return m.OutCards
	}
	return nil
}

func (m *DDZBillPlayerInfo) GetHandCards() []uint32 {
	if m != nil {
		return m.HandCards
	}
	return nil
}

// DDZGameOverNtf 斗地主游戏结束通知
// CMD: 0x1500C
type DDZGameOverNtf struct {
	WinnerId         *uint64              `protobuf:"varint,1,opt,name=winner_id,json=winnerId" json:"winner_id,omitempty"`
	Spring           *bool                `protobuf:"varint,2,opt,name=spring" json:"spring,omitempty"`
	AntiSpring       *bool                `protobuf:"varint,3,opt,name=anti_spring,json=antiSpring" json:"anti_spring,omitempty"`
	ShowHandTime     *uint32              `protobuf:"varint,4,opt,name=show_hand_time,json=showHandTime" json:"show_hand_time,omitempty"`
	Bills            []*DDZBillPlayerInfo `protobuf:"bytes,5,rep,name=bills" json:"bills,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *DDZGameOverNtf) Reset()                    { *m = DDZGameOverNtf{} }
func (m *DDZGameOverNtf) String() string            { return proto.CompactTextString(m) }
func (*DDZGameOverNtf) ProtoMessage()               {}
func (*DDZGameOverNtf) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{13} }

func (m *DDZGameOverNtf) GetWinnerId() uint64 {
	if m != nil && m.WinnerId != nil {
		return *m.WinnerId
	}
	return 0
}

func (m *DDZGameOverNtf) GetSpring() bool {
	if m != nil && m.Spring != nil {
		return *m.Spring
	}
	return false
}

func (m *DDZGameOverNtf) GetAntiSpring() bool {
	if m != nil && m.AntiSpring != nil {
		return *m.AntiSpring
	}
	return false
}

func (m *DDZGameOverNtf) GetShowHandTime() uint32 {
	if m != nil && m.ShowHandTime != nil {
		return *m.ShowHandTime
	}
	return 0
}

func (m *DDZGameOverNtf) GetBills() []*DDZBillPlayerInfo {
	if m != nil {
		return m.Bills
	}
	return nil
}

// DDZResumeGameReq 恢复对局请求
// CMD: 0x15010
type DDZResumeGameReq struct {
	Reserve          *uint32 `protobuf:"varint,1,opt,name=reserve" json:"reserve,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *DDZResumeGameReq) Reset()                    { *m = DDZResumeGameReq{} }
func (m *DDZResumeGameReq) String() string            { return proto.CompactTextString(m) }
func (*DDZResumeGameReq) ProtoMessage()               {}
func (*DDZResumeGameReq) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{14} }

func (m *DDZResumeGameReq) GetReserve() uint32 {
	if m != nil && m.Reserve != nil {
		return *m.Reserve
	}
	return 0
}

// DDZPlayerInfo 游戏中的玩家信息
type DDZPlayerInfo struct {
	PlayerInfo       *RoomPlayerInfo `protobuf:"bytes,1,opt,name=player_info,json=playerInfo" json:"player_info,omitempty"`
	OutCards         []uint32        `protobuf:"varint,2,rep,name=out_cards,json=outCards" json:"out_cards,omitempty"`
	HandCards        []uint32        `protobuf:"varint,3,rep,name=hand_cards,json=handCards" json:"hand_cards,omitempty"`
	Lord             *bool           `protobuf:"varint,4,opt,name=lord" json:"lord,omitempty"`
	Tuoguan          *bool           `protobuf:"varint,5,opt,name=tuoguan" json:"tuoguan,omitempty"`
	GrabLord         *GrabLordType   `protobuf:"varint,6,opt,name=grab_lord,json=grabLord,enum=room.GrabLordType" json:"grab_lord,omitempty"`
	Double           *DoubleType     `protobuf:"varint,7,opt,name=double,enum=room.DoubleType" json:"double,omitempty"`
	HandCardsCount   *uint32         `protobuf:"varint,8,opt,name=hand_cards_count,json=handCardsCount" json:"hand_cards_count,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *DDZPlayerInfo) Reset()                    { *m = DDZPlayerInfo{} }
func (m *DDZPlayerInfo) String() string            { return proto.CompactTextString(m) }
func (*DDZPlayerInfo) ProtoMessage()               {}
func (*DDZPlayerInfo) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{15} }

func (m *DDZPlayerInfo) GetPlayerInfo() *RoomPlayerInfo {
	if m != nil {
		return m.PlayerInfo
	}
	return nil
}

func (m *DDZPlayerInfo) GetOutCards() []uint32 {
	if m != nil {
		return m.OutCards
	}
	return nil
}

func (m *DDZPlayerInfo) GetHandCards() []uint32 {
	if m != nil {
		return m.HandCards
	}
	return nil
}

func (m *DDZPlayerInfo) GetLord() bool {
	if m != nil && m.Lord != nil {
		return *m.Lord
	}
	return false
}

func (m *DDZPlayerInfo) GetTuoguan() bool {
	if m != nil && m.Tuoguan != nil {
		return *m.Tuoguan
	}
	return false
}

func (m *DDZPlayerInfo) GetGrabLord() GrabLordType {
	if m != nil && m.GrabLord != nil {
		return *m.GrabLord
	}
	return GrabLordType_GLT_NONE
}

func (m *DDZPlayerInfo) GetDouble() DoubleType {
	if m != nil && m.Double != nil {
		return *m.Double
	}
	return DoubleType_DT_NONE
}

func (m *DDZPlayerInfo) GetHandCardsCount() uint32 {
	if m != nil && m.HandCardsCount != nil {
		return *m.HandCardsCount
	}
	return 0
}

// DDZDeskInfo 游戏基本信息
type DDZDeskInfo struct {
	Players          []*DDZPlayerInfo `protobuf:"bytes,1,rep,name=players" json:"players,omitempty"`
	Stage            *NextStage       `protobuf:"bytes,2,opt,name=stage" json:"stage,omitempty"`
	CurPlayerId      *uint64          `protobuf:"varint,3,opt,name=cur_player_id,json=curPlayerId" json:"cur_player_id,omitempty"`
	Dipai            []uint32         `protobuf:"varint,4,rep,name=dipai" json:"dipai,omitempty"`
	TotalGrab        *uint32          `protobuf:"varint,5,opt,name=total_grab,json=totalGrab" json:"total_grab,omitempty"`
	TotalDouble      *uint32          `protobuf:"varint,6,opt,name=total_double,json=totalDouble" json:"total_double,omitempty"`
	TotalBomb        *uint32          `protobuf:"varint,7,opt,name=total_bomb,json=totalBomb" json:"total_bomb,omitempty"`
	CurCardType      *CardType        `protobuf:"varint,8,opt,name=cur_card_type,json=curCardType,enum=room.CardType" json:"cur_card_type,omitempty"`
	CurCardPivot     *uint32          `protobuf:"varint,9,opt,name=cur_card_pivot,json=curCardPivot" json:"cur_card_pivot,omitempty"`
	CurOutCards      []uint32         `protobuf:"varint,10,rep,name=cur_out_cards,json=curOutCards" json:"cur_out_cards,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *DDZDeskInfo) Reset()                    { *m = DDZDeskInfo{} }
func (m *DDZDeskInfo) String() string            { return proto.CompactTextString(m) }
func (*DDZDeskInfo) ProtoMessage()               {}
func (*DDZDeskInfo) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{16} }

func (m *DDZDeskInfo) GetPlayers() []*DDZPlayerInfo {
	if m != nil {
		return m.Players
	}
	return nil
}

func (m *DDZDeskInfo) GetStage() *NextStage {
	if m != nil {
		return m.Stage
	}
	return nil
}

func (m *DDZDeskInfo) GetCurPlayerId() uint64 {
	if m != nil && m.CurPlayerId != nil {
		return *m.CurPlayerId
	}
	return 0
}

func (m *DDZDeskInfo) GetDipai() []uint32 {
	if m != nil {
		return m.Dipai
	}
	return nil
}

func (m *DDZDeskInfo) GetTotalGrab() uint32 {
	if m != nil && m.TotalGrab != nil {
		return *m.TotalGrab
	}
	return 0
}

func (m *DDZDeskInfo) GetTotalDouble() uint32 {
	if m != nil && m.TotalDouble != nil {
		return *m.TotalDouble
	}
	return 0
}

func (m *DDZDeskInfo) GetTotalBomb() uint32 {
	if m != nil && m.TotalBomb != nil {
		return *m.TotalBomb
	}
	return 0
}

func (m *DDZDeskInfo) GetCurCardType() CardType {
	if m != nil && m.CurCardType != nil {
		return *m.CurCardType
	}
	return CardType_CT_NONE
}

func (m *DDZDeskInfo) GetCurCardPivot() uint32 {
	if m != nil && m.CurCardPivot != nil {
		return *m.CurCardPivot
	}
	return 0
}

func (m *DDZDeskInfo) GetCurOutCards() []uint32 {
	if m != nil {
		return m.CurOutCards
	}
	return nil
}

// DDZResumeGameRsp 恢复对局返回
// CMD: 0x15011
type DDZResumeGameRsp struct {
	Result           *Result      `protobuf:"bytes,1,opt,name=result" json:"result,omitempty"`
	GameInfo         *DDZDeskInfo `protobuf:"bytes,2,opt,name=game_info,json=gameInfo" json:"game_info,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *DDZResumeGameRsp) Reset()                    { *m = DDZResumeGameRsp{} }
func (m *DDZResumeGameRsp) String() string            { return proto.CompactTextString(m) }
func (*DDZResumeGameRsp) ProtoMessage()               {}
func (*DDZResumeGameRsp) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{17} }

func (m *DDZResumeGameRsp) GetResult() *Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *DDZResumeGameRsp) GetGameInfo() *DDZDeskInfo {
	if m != nil {
		return m.GameInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*Result)(nil), "room.Result")
	proto.RegisterType((*NextStage)(nil), "room.NextStage")
	proto.RegisterType((*DDZStartGameNtf)(nil), "room.DDZStartGameNtf")
	proto.RegisterType((*DDZDealNtf)(nil), "room.DDZDealNtf")
	proto.RegisterType((*DDZGrabLordReq)(nil), "room.DDZGrabLordReq")
	proto.RegisterType((*DDZGrabLordNtf)(nil), "room.DDZGrabLordNtf")
	proto.RegisterType((*DDZLordNtf)(nil), "room.DDZLordNtf")
	proto.RegisterType((*DDZDoubleReq)(nil), "room.DDZDoubleReq")
	proto.RegisterType((*DDZDoubleNtf)(nil), "room.DDZDoubleNtf")
	proto.RegisterType((*DDZPlayCardReq)(nil), "room.DDZPlayCardReq")
	proto.RegisterType((*DDZPlayCardRsp)(nil), "room.DDZPlayCardRsp")
	proto.RegisterType((*DDZPlayCardNtf)(nil), "room.DDZPlayCardNtf")
	proto.RegisterType((*DDZBillPlayerInfo)(nil), "room.DDZBillPlayerInfo")
	proto.RegisterType((*DDZGameOverNtf)(nil), "room.DDZGameOverNtf")
	proto.RegisterType((*DDZResumeGameReq)(nil), "room.DDZResumeGameReq")
	proto.RegisterType((*DDZPlayerInfo)(nil), "room.DDZPlayerInfo")
	proto.RegisterType((*DDZDeskInfo)(nil), "room.DDZDeskInfo")
	proto.RegisterType((*DDZResumeGameRsp)(nil), "room.DDZResumeGameRsp")
	proto.RegisterEnum("room.DDZStage", DDZStage_name, DDZStage_value)
	proto.RegisterEnum("room.CardType", CardType_name, CardType_value)
	proto.RegisterEnum("room.GrabLordType", GrabLordType_name, GrabLordType_value)
	proto.RegisterEnum("room.DoubleType", DoubleType_name, DoubleType_value)
}

func init() { proto.RegisterFile("game_ddz.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 1297 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x56, 0xdd, 0x8e, 0xdb, 0x44,
	0x14, 0xc6, 0x76, 0x7e, 0xec, 0xe3, 0x24, 0xeb, 0x0e, 0x4b, 0x09, 0xad, 0x2a, 0x16, 0xb3, 0xa0,
	0xd5, 0x96, 0x6e, 0x45, 0x0a, 0x5c, 0x70, 0x81, 0x94, 0xc4, 0x51, 0x1a, 0x11, 0x92, 0x95, 0x93,
	0x22, 0x91, 0x0b, 0x2c, 0x27, 0x9e, 0x6e, 0x2d, 0x1c, 0x3b, 0xd8, 0xce, 0xb6, 0xe5, 0x1d, 0xe0,
	0x05, 0xfa, 0x1c, 0x70, 0xc3, 0x13, 0xf1, 0x0e, 0x20, 0xa1, 0x39, 0x33, 0x76, 0xec, 0xb0, 0x6d,
	0xca, 0xdd, 0x9c, 0xf3, 0x9d, 0x39, 0xf3, 0x9d, 0x9f, 0x39, 0x33, 0xd0, 0xba, 0x72, 0xd7, 0xd4,
	0xf1, 0xbc, 0x5f, 0x2e, 0x36, 0x71, 0x94, 0x46, 0xa4, 0x12, 0x47, 0xd1, 0xfa, 0x0e, 0x2c, 0xdd,
	0x84, 0x72, 0x8d, 0xf9, 0x0d, 0xd4, 0x6c, 0x9a, 0x6c, 0x83, 0x94, 0x7c, 0x00, 0x2a, 0x8d, 0x63,
	0x67, 0x15, 0x79, 0xb4, 0x2d, 0x9d, 0x48, 0x67, 0x4d, 0xbb, 0x4e, 0xe3, 0xb8, 0x1f, 0x79, 0x34,
	0x83, 0x3c, 0x9a, 0xac, 0xda, 0xf2, 0x89, 0x74, 0xa6, 0x21, 0x64, 0xd1, 0x64, 0x65, 0x0e, 0x40,
	0x9b, 0xd0, 0x17, 0xe9, 0x2c, 0x75, 0xaf, 0x28, 0x39, 0x85, 0x6a, 0xc2, 0x16, 0xb8, 0xbf, 0xd5,
	0x69, 0x5d, 0xb0, 0xe3, 0x2e, 0x2c, 0x6b, 0x81, 0xb0, 0xcd, 0x41, 0x42, 0xa0, 0x92, 0xfa, 0x6b,
	0x8a, 0x9e, 0x9a, 0x36, 0xae, 0xcd, 0x1f, 0xe1, 0x88, 0x9b, 0xc5, 0xe9, 0xd0, 0x5d, 0xd3, 0x49,
	0xfa, 0x94, 0xdc, 0x05, 0x6d, 0x13, 0xb8, 0x2f, 0x69, 0xec, 0xf8, 0x1e, 0x3a, 0xac, 0xd8, 0x2a,
	0x57, 0x8c, 0x3c, 0x72, 0x01, 0x10, 0xd2, 0x17, 0xa9, 0xc3, 0x8f, 0x63, 0x9e, 0xf4, 0xce, 0x11,
	0x3f, 0x2e, 0xa7, 0x63, 0x6b, 0x61, 0xb6, 0x34, 0x5f, 0x00, 0x58, 0xd6, 0xc2, 0xa2, 0x6e, 0xc0,
	0x5c, 0x1f, 0x43, 0x75, 0xe5, 0xc6, 0x5e, 0xd2, 0x96, 0x4e, 0x94, 0xb3, 0xa6, 0xcd, 0x05, 0x72,
	0x0a, 0xad, 0x95, 0x1b, 0x04, 0xce, 0xee, 0x54, 0x19, 0x4f, 0x6d, 0x30, 0xed, 0xe5, 0xcd, 0x27,
	0x2b, 0x07, 0x4f, 0x3e, 0x85, 0x96, 0x65, 0x2d, 0x86, 0xb1, 0xbb, 0x1c, 0x47, 0xb1, 0x67, 0xd3,
	0x9f, 0x59, 0xfc, 0x57, 0xb1, 0xbb, 0xc4, 0x98, 0x54, 0x1b, 0xd7, 0xe6, 0x1f, 0x52, 0xc9, 0xec,
	0x60, 0xfc, 0x99, 0x0f, 0x79, 0xe7, 0x83, 0xdc, 0x03, 0x48, 0xa3, 0xd4, 0x0d, 0x1c, 0x44, 0x14,
	0xcc, 0xae, 0x86, 0x1a, 0xe6, 0x96, 0x85, 0x87, 0xc4, 0x77, 0x4e, 0x2b, 0x3c, 0x3c, 0xa6, 0x7d,
	0x4d, 0x78, 0xd5, 0x83, 0xe1, 0xfd, 0x26, 0x61, 0x66, 0xdf, 0x8a, 0x74, 0x99, 0xa0, 0xbc, 0x4f,
	0xf0, 0x18, 0xaa, 0x9e, 0xbf, 0x71, 0xfd, 0xb6, 0xc2, 0xab, 0x82, 0xc2, 0x1e, 0xa1, 0xca, 0x41,
	0x42, 0xf7, 0xa1, 0xc1, 0x2a, 0x1d, 0x6d, 0x97, 0x01, 0x65, 0xd9, 0xbe, 0x0b, 0x9a, 0x9f, 0x38,
	0x1e, 0xca, 0x22, 0xe5, 0xaa, 0x9f, 0x70, 0xdc, 0x7c, 0x25, 0x15, 0xac, 0x0f, 0xf2, 0x2f, 0xb9,
	0x92, 0xcb, 0xae, 0xc8, 0x47, 0xd0, 0xe0, 0xc1, 0x09, 0x9c, 0xe7, 0x5f, 0x47, 0x9d, 0x30, 0xf9,
	0xbf, 0xa1, 0xcc, 0xb0, 0x27, 0x58, 0x69, 0xfa, 0x2e, 0x6f, 0x9d, 0x9b, 0x1b, 0xf7, 0x3e, 0x68,
	0x6c, 0xe1, 0xa4, 0x2f, 0x37, 0x9c, 0x57, 0x7e, 0xf5, 0xd8, 0xbe, 0xf9, 0xcb, 0x0d, 0xb5, 0xd5,
	0x95, 0x58, 0x99, 0x5f, 0x95, 0x9d, 0x26, 0x1b, 0x72, 0x0a, 0xb5, 0x18, 0x47, 0x00, 0x06, 0xac,
	0x77, 0x1a, 0x7c, 0x2f, 0x1f, 0x0b, 0xb6, 0xc0, 0xcc, 0x5f, 0xe5, 0xd2, 0xc6, 0x83, 0xc9, 0xca,
	0xa9, 0xca, 0xaf, 0xa5, 0xaa, 0xbc, 0x99, 0xea, 0xae, 0x5f, 0x96, 0xd1, 0x7a, 0x89, 0xf9, 0xca,
	0xfa, 0xa5, 0x17, 0xad, 0x6f, 0x6a, 0xe8, 0xea, 0xc1, 0x86, 0xae, 0x1d, 0x4a, 0x3a, 0xf9, 0x14,
	0x8e, 0x72, 0x86, 0xce, 0xc6, 0xbf, 0x8e, 0xd2, 0x76, 0x1d, 0x4f, 0x6e, 0x66, 0xbc, 0x2e, 0x99,
	0xd2, 0xfc, 0x47, 0x82, 0x5b, 0x96, 0xb5, 0xe8, 0xf9, 0xf9, 0x6c, 0x08, 0x9f, 0x46, 0x6f, 0x4e,
	0x89, 0x01, 0xca, 0x73, 0x3f, 0x14, 0x9d, 0xc3, 0x96, 0xec, 0x1a, 0xb3, 0x59, 0x8c, 0x99, 0xa8,
	0xda, 0xb8, 0x26, 0x77, 0x40, 0x5d, 0x6f, 0x83, 0xd4, 0xdf, 0x04, 0xbc, 0x47, 0xaa, 0x76, 0x2e,
	0xb3, 0xa4, 0x26, 0xab, 0x28, 0xe6, 0x17, 0x53, 0xb1, 0xb9, 0x40, 0x3e, 0x86, 0xe6, 0x6a, 0x1b,
	0xc7, 0x34, 0x4c, 0x1d, 0x8e, 0xd6, 0x10, 0x6d, 0x08, 0xe5, 0x0c, 0x8d, 0x08, 0x54, 0x82, 0x28,
	0xf6, 0x30, 0x18, 0xd5, 0xc6, 0x35, 0x63, 0x1b, 0x6d, 0x53, 0x87, 0xd7, 0x49, 0xc5, 0x3a, 0xa9,
	0xd1, 0x36, 0xed, 0x63, 0xa9, 0xee, 0x01, 0x3c, 0x73, 0x43, 0x4f, 0xa0, 0x1a, 0xa2, 0x1a, 0xd3,
	0x20, 0x6c, 0xfe, 0x29, 0x26, 0x96, 0xbb, 0xa6, 0xd3, 0x6b, 0x1a, 0x8b, 0x7e, 0x78, 0xee, 0x87,
	0x61, 0x29, 0x78, 0xae, 0x18, 0x79, 0xe4, 0x36, 0xd4, 0x92, 0x4d, 0xec, 0x87, 0x57, 0x22, 0x7e,
	0x21, 0x91, 0x0f, 0x41, 0x77, 0xc3, 0xd4, 0x77, 0x04, 0xa8, 0x20, 0x08, 0x4c, 0x35, 0xe3, 0x06,
	0xa7, 0xd0, 0x4a, 0x9e, 0x45, 0xcf, 0x1d, 0x24, 0x83, 0x0f, 0x07, 0xef, 0x84, 0x06, 0xd3, 0x3e,
	0x76, 0x43, 0x6f, 0xee, 0xaf, 0x29, 0x79, 0x00, 0xd5, 0xa5, 0x1f, 0x04, 0x49, 0xbb, 0x7a, 0xa2,
	0x9c, 0xe9, 0x9d, 0xf7, 0xf3, 0xa7, 0xa7, 0x5c, 0x20, 0x9b, 0x5b, 0x99, 0x9f, 0x81, 0x61, 0x59,
	0x0b, 0xd6, 0xe2, 0x6b, 0xca, 0x42, 0x60, 0x97, 0xab, 0x0d, 0xf5, 0x98, 0x26, 0x34, 0xbe, 0xce,
	0xdf, 0x3f, 0x21, 0x9a, 0xbf, 0xcb, 0xd0, 0x14, 0xbd, 0x2f, 0xea, 0xfc, 0x25, 0xe8, 0x59, 0x9d,
	0xc3, 0xa7, 0x91, 0xb8, 0x38, 0xc7, 0xe2, 0xe2, 0x44, 0xd1, 0xba, 0x70, 0x22, 0x6c, 0x4a, 0xed,
	0xb1, 0x4b, 0xb8, 0xfc, 0xc6, 0x84, 0x2b, 0x7b, 0x09, 0xcf, 0x0b, 0x58, 0x29, 0x14, 0xb0, 0x0d,
	0xf5, 0x74, 0x1b, 0x5d, 0x6d, 0xdd, 0x10, 0x3b, 0x42, 0xb5, 0x33, 0x91, 0x3c, 0x04, 0x8d, 0x4d,
	0x59, 0x07, 0xb7, 0xd4, 0xf0, 0xa2, 0x11, 0x4e, 0x2f, 0x7b, 0x63, 0xf8, 0x65, 0xbb, 0x12, 0x12,
	0x39, 0x83, 0x9a, 0x98, 0x5c, 0x75, 0xb4, 0x36, 0x44, 0x06, 0x51, 0x87, 0xb6, 0x02, 0x27, 0x67,
	0x60, 0xec, 0x78, 0x3a, 0xab, 0x68, 0x1b, 0xa6, 0x6d, 0x15, 0x13, 0xd6, 0xca, 0xd9, 0xf6, 0x99,
	0xd6, 0xfc, 0x5b, 0x06, 0x1d, 0x9f, 0xdd, 0xe4, 0x27, 0x0c, 0xff, 0x01, 0xd4, 0x79, 0x32, 0xf8,
	0x00, 0xd3, 0x3b, 0xef, 0xe6, 0x65, 0x2a, 0x24, 0x2c, 0xb3, 0x21, 0x9f, 0x64, 0xdf, 0x89, 0xd7,
	0xbc, 0xef, 0xe2, 0x3f, 0x61, 0x62, 0xfb, 0x17, 0xc6, 0x80, 0x82, 0xad, 0xa7, 0xaf, 0xb6, 0xf1,
	0x65, 0x61, 0x1a, 0xf1, 0xb7, 0xa5, 0x52, 0x7c, 0x5b, 0xca, 0x0f, 0x52, 0x75, 0xff, 0x41, 0xda,
	0x1f, 0xe9, 0xb5, 0xff, 0x8e, 0xf4, 0xf2, 0x88, 0xaa, 0xef, 0x8f, 0xa8, 0x0e, 0xa7, 0xb6, 0x1b,
	0x79, 0xea, 0x8d, 0x23, 0x8f, 0x51, 0xcd, 0x04, 0xfc, 0x86, 0x64, 0x7b, 0xf8, 0xfc, 0xd1, 0x78,
	0xbf, 0x0b, 0x23, 0x1c, 0x3f, 0x59, 0xd0, 0xbb, 0x6e, 0x02, 0x0c, 0x8c, 0x79, 0x9a, 0x8a, 0x86,
	0x32, 0x9f, 0xed, 0x37, 0xf9, 0xdb, 0x0e, 0x7b, 0x72, 0x01, 0x1a, 0xfe, 0x1c, 0xb1, 0xb9, 0x79,
	0xf6, 0x6f, 0xe5, 0xa5, 0xca, 0xca, 0x69, 0xab, 0xcc, 0x86, 0xad, 0xce, 0x5f, 0x49, 0xa0, 0x66,
	0xdf, 0x3c, 0x42, 0x70, 0x30, 0x38, 0xb3, 0x79, 0x77, 0x38, 0x70, 0x26, 0xd3, 0xc9, 0xc0, 0x78,
	0xa7, 0xac, 0xb3, 0x06, 0xdd, 0xb1, 0x21, 0x95, 0x75, 0xfd, 0xee, 0x78, 0x6c, 0xc8, 0x65, 0xdd,
	0xd0, 0xee, 0xf6, 0x0c, 0x85, 0x1c, 0x63, 0x18, 0xd9, 0xde, 0xe9, 0x93, 0xde, 0x78, 0x60, 0x54,
	0xc8, 0x7b, 0x38, 0x7e, 0x85, 0xf6, 0x72, 0xdc, 0xfd, 0x61, 0x34, 0x19, 0x1a, 0xd5, 0xb2, 0x83,
	0xe9, 0xf7, 0x03, 0xdb, 0xa8, 0x9d, 0xff, 0x25, 0x81, 0x9a, 0xa7, 0x57, 0x87, 0x7a, 0x7f, 0x9e,
	0xd1, 0x6a, 0x82, 0xd6, 0x9f, 0x3b, 0xb3, 0xd1, 0x64, 0x38, 0x1e, 0x18, 0x92, 0xc0, 0x2e, 0xbb,
	0x23, 0xdb, 0x90, 0x33, 0xec, 0xf1, 0x93, 0xc9, 0x62, 0x64, 0x28, 0xa4, 0x01, 0xaa, 0xc0, 0x66,
	0x46, 0x45, 0x80, 0x73, 0x7b, 0x74, 0x39, 0x1e, 0x18, 0x55, 0x01, 0x3e, 0xea, 0x4e, 0xac, 0xcf,
	0x8d, 0x5a, 0x41, 0xea, 0x18, 0x75, 0xd2, 0x02, 0xc8, 0x4d, 0x67, 0x86, 0x2a, 0xe4, 0x47, 0x33,
	0x66, 0x3c, 0x33, 0xb4, 0xa2, 0xdc, 0x99, 0x19, 0x20, 0xe4, 0x2f, 0x04, 0xae, 0x17, 0xe5, 0xce,
	0xcc, 0x68, 0x08, 0x92, 0xbd, 0xe9, 0x77, 0x3d, 0xa3, 0x49, 0x8e, 0x40, 0xef, 0xcf, 0x9d, 0x6f,
	0x47, 0x93, 0x21, 0x2a, 0x5a, 0xe7, 0x2e, 0x34, 0x8a, 0x17, 0x9c, 0x71, 0x19, 0x8e, 0xf3, 0x78,
	0x0d, 0x68, 0x30, 0x89, 0x25, 0x7b, 0x3c, 0xb5, 0x2d, 0x5e, 0x04, 0x8e, 0xcf, 0x99, 0x92, 0xe9,
	0xe4, 0x6c, 0x8f, 0x48, 0xff, 0x11, 0xe8, 0xc2, 0x02, 0x15, 0x95, 0xf3, 0xaf, 0x01, 0x76, 0x53,
	0x81, 0xd1, 0xb1, 0x8a, 0xf9, 0xb4, 0xe6, 0x59, 0x8d, 0x24, 0x76, 0x1c, 0x62, 0x73, 0xa1, 0x91,
	0x7b, 0xb7, 0x17, 0xc7, 0x49, 0x4a, 0xaf, 0xe9, 0xc3, 0x55, 0xe0, 0xb3, 0xf7, 0x6a, 0xb3, 0x7c,
	0xc8, 0xda, 0xea, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xda, 0xfe, 0x86, 0x57, 0xa9, 0x0c, 0x00,
	0x00,
}
