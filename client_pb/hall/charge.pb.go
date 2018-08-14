// Code generated by protoc-gen-go. DO NOT EDIT.
// source: charge.proto

package hall // import "steve/client_pb/hall"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "steve/client_pb/common"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// ChargeItem 充值商品
type ChargeItem struct {
	ItemId               *uint64  `protobuf:"varint,1,opt,name=item_id,json=itemId" json:"item_id,omitempty"`
	ItemName             *string  `protobuf:"bytes,2,opt,name=item_name,json=itemName" json:"item_name,omitempty"`
	Price                *uint64  `protobuf:"varint,3,opt,name=price" json:"price,omitempty"`
	Tag                  *string  `protobuf:"bytes,4,opt,name=tag" json:"tag,omitempty"`
	Coin                 *uint64  `protobuf:"varint,5,opt,name=coin" json:"coin,omitempty"`
	PresentCoin          *uint64  `protobuf:"varint,6,opt,name=present_coin,json=presentCoin" json:"present_coin,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChargeItem) Reset()         { *m = ChargeItem{} }
func (m *ChargeItem) String() string { return proto.CompactTextString(m) }
func (*ChargeItem) ProtoMessage()    {}
func (*ChargeItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_charge_6adf660e554473de, []int{0}
}
func (m *ChargeItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChargeItem.Unmarshal(m, b)
}
func (m *ChargeItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChargeItem.Marshal(b, m, deterministic)
}
func (dst *ChargeItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChargeItem.Merge(dst, src)
}
func (m *ChargeItem) XXX_Size() int {
	return xxx_messageInfo_ChargeItem.Size(m)
}
func (m *ChargeItem) XXX_DiscardUnknown() {
	xxx_messageInfo_ChargeItem.DiscardUnknown(m)
}

var xxx_messageInfo_ChargeItem proto.InternalMessageInfo

func (m *ChargeItem) GetItemId() uint64 {
	if m != nil && m.ItemId != nil {
		return *m.ItemId
	}
	return 0
}

func (m *ChargeItem) GetItemName() string {
	if m != nil && m.ItemName != nil {
		return *m.ItemName
	}
	return ""
}

func (m *ChargeItem) GetPrice() uint64 {
	if m != nil && m.Price != nil {
		return *m.Price
	}
	return 0
}

func (m *ChargeItem) GetTag() string {
	if m != nil && m.Tag != nil {
		return *m.Tag
	}
	return ""
}

func (m *ChargeItem) GetCoin() uint64 {
	if m != nil && m.Coin != nil {
		return *m.Coin
	}
	return 0
}

func (m *ChargeItem) GetPresentCoin() uint64 {
	if m != nil && m.PresentCoin != nil {
		return *m.PresentCoin
	}
	return 0
}

// GetChargeInfoReq 获取充值信息请求
type GetChargeInfoReq struct {
	Platform             *common.Platform `protobuf:"varint,1,opt,name=platform,enum=common.Platform" json:"platform,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *GetChargeInfoReq) Reset()         { *m = GetChargeInfoReq{} }
func (m *GetChargeInfoReq) String() string { return proto.CompactTextString(m) }
func (*GetChargeInfoReq) ProtoMessage()    {}
func (*GetChargeInfoReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_charge_6adf660e554473de, []int{1}
}
func (m *GetChargeInfoReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetChargeInfoReq.Unmarshal(m, b)
}
func (m *GetChargeInfoReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetChargeInfoReq.Marshal(b, m, deterministic)
}
func (dst *GetChargeInfoReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetChargeInfoReq.Merge(dst, src)
}
func (m *GetChargeInfoReq) XXX_Size() int {
	return xxx_messageInfo_GetChargeInfoReq.Size(m)
}
func (m *GetChargeInfoReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetChargeInfoReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetChargeInfoReq proto.InternalMessageInfo

func (m *GetChargeInfoReq) GetPlatform() common.Platform {
	if m != nil && m.Platform != nil {
		return *m.Platform
	}
	return common.Platform_Android
}

// GetChargeInfoRsp 获取充值信息响应
type GetChargeInfoRsp struct {
	Result               *common.Result `protobuf:"bytes,1,opt,name=result" json:"result,omitempty"`
	Items                []*ChargeItem  `protobuf:"bytes,2,rep,name=items" json:"items,omitempty"`
	TodayCharge          *uint64        `protobuf:"varint,3,opt,name=today_charge,json=todayCharge" json:"today_charge,omitempty"`
	DayMaxCharge         *uint64        `protobuf:"varint,4,opt,name=day_max_charge,json=dayMaxCharge" json:"day_max_charge,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GetChargeInfoRsp) Reset()         { *m = GetChargeInfoRsp{} }
func (m *GetChargeInfoRsp) String() string { return proto.CompactTextString(m) }
func (*GetChargeInfoRsp) ProtoMessage()    {}
func (*GetChargeInfoRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_charge_6adf660e554473de, []int{2}
}
func (m *GetChargeInfoRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetChargeInfoRsp.Unmarshal(m, b)
}
func (m *GetChargeInfoRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetChargeInfoRsp.Marshal(b, m, deterministic)
}
func (dst *GetChargeInfoRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetChargeInfoRsp.Merge(dst, src)
}
func (m *GetChargeInfoRsp) XXX_Size() int {
	return xxx_messageInfo_GetChargeInfoRsp.Size(m)
}
func (m *GetChargeInfoRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetChargeInfoRsp.DiscardUnknown(m)
}

var xxx_messageInfo_GetChargeInfoRsp proto.InternalMessageInfo

func (m *GetChargeInfoRsp) GetResult() *common.Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *GetChargeInfoRsp) GetItems() []*ChargeItem {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *GetChargeInfoRsp) GetTodayCharge() uint64 {
	if m != nil && m.TodayCharge != nil {
		return *m.TodayCharge
	}
	return 0
}

func (m *GetChargeInfoRsp) GetDayMaxCharge() uint64 {
	if m != nil && m.DayMaxCharge != nil {
		return *m.DayMaxCharge
	}
	return 0
}

// ChargeReq 充值请求
type ChargeReq struct {
	ItemId               *uint64          `protobuf:"varint,1,opt,name=item_id,json=itemId" json:"item_id,omitempty"`
	Cost                 *uint64          `protobuf:"varint,2,opt,name=cost" json:"cost,omitempty"`
	Data                 []byte           `protobuf:"bytes,3,opt,name=data" json:"data,omitempty"`
	Platform             *common.Platform `protobuf:"varint,4,opt,name=platform,enum=common.Platform" json:"platform,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ChargeReq) Reset()         { *m = ChargeReq{} }
func (m *ChargeReq) String() string { return proto.CompactTextString(m) }
func (*ChargeReq) ProtoMessage()    {}
func (*ChargeReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_charge_6adf660e554473de, []int{3}
}
func (m *ChargeReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChargeReq.Unmarshal(m, b)
}
func (m *ChargeReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChargeReq.Marshal(b, m, deterministic)
}
func (dst *ChargeReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChargeReq.Merge(dst, src)
}
func (m *ChargeReq) XXX_Size() int {
	return xxx_messageInfo_ChargeReq.Size(m)
}
func (m *ChargeReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ChargeReq.DiscardUnknown(m)
}

var xxx_messageInfo_ChargeReq proto.InternalMessageInfo

func (m *ChargeReq) GetItemId() uint64 {
	if m != nil && m.ItemId != nil {
		return *m.ItemId
	}
	return 0
}

func (m *ChargeReq) GetCost() uint64 {
	if m != nil && m.Cost != nil {
		return *m.Cost
	}
	return 0
}

func (m *ChargeReq) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *ChargeReq) GetPlatform() common.Platform {
	if m != nil && m.Platform != nil {
		return *m.Platform
	}
	return common.Platform_Android
}

// ChargeRsp 充值应答
type ChargeRsp struct {
	Result               *common.Result `protobuf:"bytes,1,opt,name=result" json:"result,omitempty"`
	ObtainedCoin         *uint64        `protobuf:"varint,2,opt,name=obtained_coin,json=obtainedCoin" json:"obtained_coin,omitempty"`
	NewCoin              *uint64        `protobuf:"varint,3,opt,name=new_coin,json=newCoin" json:"new_coin,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ChargeRsp) Reset()         { *m = ChargeRsp{} }
func (m *ChargeRsp) String() string { return proto.CompactTextString(m) }
func (*ChargeRsp) ProtoMessage()    {}
func (*ChargeRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_charge_6adf660e554473de, []int{4}
}
func (m *ChargeRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChargeRsp.Unmarshal(m, b)
}
func (m *ChargeRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChargeRsp.Marshal(b, m, deterministic)
}
func (dst *ChargeRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChargeRsp.Merge(dst, src)
}
func (m *ChargeRsp) XXX_Size() int {
	return xxx_messageInfo_ChargeRsp.Size(m)
}
func (m *ChargeRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_ChargeRsp.DiscardUnknown(m)
}

var xxx_messageInfo_ChargeRsp proto.InternalMessageInfo

func (m *ChargeRsp) GetResult() *common.Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *ChargeRsp) GetObtainedCoin() uint64 {
	if m != nil && m.ObtainedCoin != nil {
		return *m.ObtainedCoin
	}
	return 0
}

func (m *ChargeRsp) GetNewCoin() uint64 {
	if m != nil && m.NewCoin != nil {
		return *m.NewCoin
	}
	return 0
}

func init() {
	proto.RegisterType((*ChargeItem)(nil), "hall.ChargeItem")
	proto.RegisterType((*GetChargeInfoReq)(nil), "hall.GetChargeInfoReq")
	proto.RegisterType((*GetChargeInfoRsp)(nil), "hall.GetChargeInfoRsp")
	proto.RegisterType((*ChargeReq)(nil), "hall.ChargeReq")
	proto.RegisterType((*ChargeRsp)(nil), "hall.ChargeRsp")
}

func init() { proto.RegisterFile("charge.proto", fileDescriptor_charge_6adf660e554473de) }

var fileDescriptor_charge_6adf660e554473de = []byte{
	// 396 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0xc1, 0x8e, 0xd3, 0x30,
	0x10, 0x55, 0xb6, 0x69, 0xb7, 0x9d, 0x84, 0x2a, 0xb2, 0x56, 0x10, 0x96, 0x4b, 0x09, 0x68, 0xd5,
	0x03, 0x4a, 0xa5, 0x7e, 0x01, 0xa2, 0x07, 0xd4, 0x03, 0x08, 0xf9, 0xc8, 0xa5, 0x72, 0x93, 0x69,
	0x1b, 0x29, 0xb6, 0x83, 0x6d, 0x68, 0xcb, 0xdf, 0x70, 0xe6, 0x27, 0x91, 0xc7, 0x69, 0x11, 0x48,
	0x08, 0xed, 0x6d, 0xe6, 0xbd, 0x97, 0xf1, 0xcc, 0x7b, 0x81, 0xb4, 0x3a, 0x08, 0xb3, 0xc7, 0xb2,
	0x33, 0xda, 0x69, 0x16, 0x1f, 0x44, 0xdb, 0xde, 0xa7, 0x68, 0x8c, 0x36, 0x36, 0x60, 0xf7, 0x69,
	0xa5, 0xa5, 0xd4, 0x2a, 0x74, 0xc5, 0x8f, 0x08, 0x60, 0x45, 0x9f, 0xac, 0x1d, 0x4a, 0xf6, 0x0c,
	0x6e, 0x1b, 0x87, 0x72, 0xd3, 0xd4, 0x79, 0x34, 0x8b, 0xe6, 0x31, 0x1f, 0xf9, 0x76, 0x5d, 0xb3,
	0x17, 0x30, 0x21, 0x42, 0x09, 0x89, 0xf9, 0xcd, 0x2c, 0x9a, 0x4f, 0xf8, 0xd8, 0x03, 0x1f, 0x85,
	0x44, 0x76, 0x07, 0xc3, 0xce, 0x34, 0x15, 0xe6, 0x03, 0xfa, 0x26, 0x34, 0x2c, 0x83, 0x81, 0x13,
	0xfb, 0x3c, 0x26, 0xb1, 0x2f, 0x19, 0x83, 0xb8, 0xd2, 0x8d, 0xca, 0x87, 0x24, 0xa3, 0x9a, 0xbd,
	0x84, 0xb4, 0x33, 0x68, 0x51, 0xb9, 0x0d, 0x71, 0x23, 0xe2, 0x92, 0x1e, 0x5b, 0xe9, 0x46, 0x15,
	0x6f, 0x21, 0x7b, 0x8f, 0xae, 0xdf, 0x52, 0xed, 0x34, 0xc7, 0x2f, 0xec, 0x0d, 0x8c, 0xbb, 0x56,
	0xb8, 0x9d, 0x36, 0x92, 0x36, 0x9d, 0x2e, 0xb3, 0xb2, 0x3f, 0xec, 0x53, 0x8f, 0xf3, 0xab, 0xa2,
	0xf8, 0x19, 0xfd, 0x3d, 0xc2, 0x76, 0xec, 0x01, 0x46, 0x06, 0xed, 0xd7, 0xd6, 0xd1, 0x80, 0x64,
	0x39, 0xbd, 0x0c, 0xe0, 0x84, 0xf2, 0x9e, 0x65, 0x0f, 0x30, 0xf4, 0x97, 0xda, 0xfc, 0x66, 0x36,
	0x98, 0x27, 0xcb, 0xac, 0xf4, 0xa6, 0x96, 0xbf, 0x4d, 0xe3, 0x81, 0xf6, 0x97, 0x38, 0x5d, 0x8b,
	0xf3, 0x26, 0x44, 0xd0, 0x9b, 0x91, 0x10, 0x16, 0xd4, 0xec, 0x35, 0x4c, 0xbd, 0x40, 0x8a, 0xd3,
	0x45, 0x14, 0x93, 0x28, 0xad, 0xc5, 0xf9, 0x83, 0x38, 0x05, 0x55, 0xf1, 0x1d, 0x26, 0xa1, 0xf2,
	0x87, 0xfe, 0x33, 0x11, 0x32, 0xd3, 0x3a, 0x0a, 0x83, 0xcc, 0xb4, 0xce, 0x63, 0xb5, 0x70, 0x82,
	0x9e, 0x4e, 0x39, 0xd5, 0x7f, 0x38, 0x15, 0xff, 0xd7, 0x29, 0x7b, 0x7d, 0xfb, 0x11, 0x0e, 0xbd,
	0x82, 0x27, 0x7a, 0xeb, 0x44, 0xa3, 0xb0, 0x0e, 0x21, 0x86, 0x9d, 0xd2, 0x0b, 0xe8, 0x53, 0x64,
	0xcf, 0x61, 0xac, 0xf0, 0x18, 0xf8, 0x60, 0xcd, 0xad, 0xc2, 0xa3, 0xa7, 0xde, 0x3d, 0xfd, 0x7c,
	0x67, 0x1d, 0x7e, 0xc3, 0x45, 0xd5, 0x36, 0xfe, 0x47, 0xe8, 0xb6, 0x0b, 0xef, 0xf1, 0xaf, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x7a, 0x07, 0xaa, 0xcb, 0xcd, 0x02, 0x00, 0x00,
}
