// Code generated by protoc-gen-go. DO NOT EDIT.
// source: game_scxz.proto

package room

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// RoomGiveUpReq 认输请求
type RoomGiveUpReq struct {
	Reserve          *uint32 `protobuf:"varint,1,opt,name=reserve" json:"reserve,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RoomGiveUpReq) Reset()                    { *m = RoomGiveUpReq{} }
func (m *RoomGiveUpReq) String() string            { return proto.CompactTextString(m) }
func (*RoomGiveUpReq) ProtoMessage()               {}
func (*RoomGiveUpReq) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }

func (m *RoomGiveUpReq) GetReserve() uint32 {
	if m != nil && m.Reserve != nil {
		return *m.Reserve
	}
	return 0
}

// RoomGiveUpRsp 认输响应
type RoomGiveUpRsp struct {
	ErrCode          *RoomError `protobuf:"varint,1,opt,name=err_code,json=errCode,enum=room.RoomError" json:"err_code,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *RoomGiveUpRsp) Reset()                    { *m = RoomGiveUpRsp{} }
func (m *RoomGiveUpRsp) String() string            { return proto.CompactTextString(m) }
func (*RoomGiveUpRsp) ProtoMessage()               {}
func (*RoomGiveUpRsp) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{1} }

func (m *RoomGiveUpRsp) GetErrCode() RoomError {
	if m != nil && m.ErrCode != nil {
		return *m.ErrCode
	}
	return RoomError_SUCCESS
}

// RoomGiveUpNtf 认输通知
type RoomGiveUpNtf struct {
	PlayerId         []uint64 `protobuf:"varint,1,rep,name=player_id,json=playerId" json:"player_id,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *RoomGiveUpNtf) Reset()                    { *m = RoomGiveUpNtf{} }
func (m *RoomGiveUpNtf) String() string            { return proto.CompactTextString(m) }
func (*RoomGiveUpNtf) ProtoMessage()               {}
func (*RoomGiveUpNtf) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{2} }

func (m *RoomGiveUpNtf) GetPlayerId() []uint64 {
	if m != nil {
		return m.PlayerId
	}
	return nil
}

// RoomBrokerPlayerContinueReq 破产玩家续费完成请求
type RoomBrokerPlayerContinueReq struct {
	Reserve          *uint32 `protobuf:"varint,1,opt,name=reserve" json:"reserve,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RoomBrokerPlayerContinueReq) Reset()                    { *m = RoomBrokerPlayerContinueReq{} }
func (m *RoomBrokerPlayerContinueReq) String() string            { return proto.CompactTextString(m) }
func (*RoomBrokerPlayerContinueReq) ProtoMessage()               {}
func (*RoomBrokerPlayerContinueReq) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{3} }

func (m *RoomBrokerPlayerContinueReq) GetReserve() uint32 {
	if m != nil && m.Reserve != nil {
		return *m.Reserve
	}
	return 0
}

// RoomBrokerPlayerContinueRsp 破产玩家续费完成响应
type RoomBrokerPlayerContinueRsp struct {
	ErrCode          *RoomError `protobuf:"varint,1,opt,name=err_code,json=errCode,enum=room.RoomError" json:"err_code,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *RoomBrokerPlayerContinueRsp) Reset()                    { *m = RoomBrokerPlayerContinueRsp{} }
func (m *RoomBrokerPlayerContinueRsp) String() string            { return proto.CompactTextString(m) }
func (*RoomBrokerPlayerContinueRsp) ProtoMessage()               {}
func (*RoomBrokerPlayerContinueRsp) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{4} }

func (m *RoomBrokerPlayerContinueRsp) GetErrCode() RoomError {
	if m != nil && m.ErrCode != nil {
		return *m.ErrCode
	}
	return RoomError_SUCCESS
}

func init() {
	proto.RegisterType((*RoomGiveUpReq)(nil), "room.RoomGiveUpReq")
	proto.RegisterType((*RoomGiveUpRsp)(nil), "room.RoomGiveUpRsp")
	proto.RegisterType((*RoomGiveUpNtf)(nil), "room.RoomGiveUpNtf")
	proto.RegisterType((*RoomBrokerPlayerContinueReq)(nil), "room.RoomBrokerPlayerContinueReq")
	proto.RegisterType((*RoomBrokerPlayerContinueRsp)(nil), "room.RoomBrokerPlayerContinueRsp")
}

func init() { proto.RegisterFile("game_scxz.proto", fileDescriptor6) }

var fileDescriptor6 = []byte{
	// 218 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x8e, 0xc1, 0x4b, 0xc3, 0x30,
	0x18, 0xc5, 0x29, 0x0e, 0x36, 0x23, 0x73, 0x10, 0x44, 0x8a, 0xbb, 0x8c, 0x9c, 0xa6, 0x48, 0x07,
	0x5e, 0x3c, 0x78, 0xdb, 0x10, 0xe9, 0x45, 0x24, 0xe0, 0xc5, 0x4b, 0xa8, 0xcd, 0x53, 0x82, 0x6d,
	0xbf, 0xf8, 0x25, 0x16, 0xf5, 0xaf, 0x97, 0x58, 0x04, 0xbd, 0x14, 0x3c, 0xbe, 0x1f, 0xbf, 0xc7,
	0x7b, 0x62, 0xf1, 0x5c, 0xb5, 0x30, 0xa1, 0x7e, 0xff, 0x2c, 0x3c, 0x53, 0x24, 0x39, 0x61, 0xa2,
	0xf6, 0xe4, 0x00, 0xcc, 0xc4, 0x03, 0x52, 0xa7, 0x62, 0xae, 0x89, 0xda, 0x1b, 0xd7, 0xe3, 0xde,
	0x6b, 0xbc, 0xca, 0x5c, 0x4c, 0x19, 0x01, 0xdc, 0x23, 0xcf, 0x56, 0xd9, 0x7a, 0xae, 0x7f, 0xa2,
	0xba, 0xfa, 0xa3, 0x06, 0x2f, 0xcf, 0xc4, 0x0c, 0xcc, 0xa6, 0x26, 0x3b, 0xb8, 0x87, 0x17, 0x8b,
	0x22, 0x2d, 0x14, 0x49, 0xbb, 0x4e, 0x23, 0x7a, 0x0a, 0xe6, 0x1d, 0x59, 0xa8, 0xf3, 0xdf, 0xe5,
	0xdb, 0xf8, 0x24, 0x97, 0x62, 0xdf, 0x37, 0xd5, 0x07, 0xd8, 0x38, 0x9b, 0x67, 0xab, 0xbd, 0xf5,
	0x44, 0xcf, 0x06, 0x50, 0x5a, 0x75, 0x29, 0x96, 0xc9, 0xde, 0x32, 0xbd, 0x80, 0xef, 0xbe, 0xe9,
	0x8e, 0xba, 0xe8, 0xba, 0x37, 0x8c, 0x7f, 0x2c, 0x47, 0x8a, 0xff, 0x7b, 0xbc, 0x3d, 0x7e, 0x38,
	0x0a, 0x11, 0x3d, 0x36, 0x75, 0xe3, 0xd0, 0x45, 0xe3, 0x1f, 0x37, 0x49, 0xfd, 0x0a, 0x00, 0x00,
	0xff, 0xff, 0xc7, 0xc3, 0xbe, 0xad, 0x56, 0x01, 0x00, 0x00,
}
