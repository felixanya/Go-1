// Code generated by protoc-gen-go. DO NOT EDIT.
// source: room_mgr.proto

package roommgr

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type RoomError int32

const (
	RoomError_SUCCESS RoomError = 0
	RoomError_FAILED  RoomError = 1
)

var RoomError_name = map[int32]string{
	0: "SUCCESS",
	1: "FAILED",
}
var RoomError_value = map[string]int32{
	"SUCCESS": 0,
	"FAILED":  1,
}

func (x RoomError) String() string {
	return proto.EnumName(RoomError_name, int32(x))
}
func (RoomError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_room_mgr_bc05f510f82ddb45, []int{0}
}

// 牌桌玩家
type DeskPlayer struct {
	PlayerId             uint64   `protobuf:"varint,1,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
	RobotLevel           int32    `protobuf:"varint,2,opt,name=robot_level,json=robotLevel,proto3" json:"robot_level,omitempty"`
	Seat                 uint32   `protobuf:"varint,3,opt,name=seat,proto3" json:"seat,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeskPlayer) Reset()         { *m = DeskPlayer{} }
func (m *DeskPlayer) String() string { return proto.CompactTextString(m) }
func (*DeskPlayer) ProtoMessage()    {}
func (*DeskPlayer) Descriptor() ([]byte, []int) {
	return fileDescriptor_room_mgr_bc05f510f82ddb45, []int{0}
}
func (m *DeskPlayer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeskPlayer.Unmarshal(m, b)
}
func (m *DeskPlayer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeskPlayer.Marshal(b, m, deterministic)
}
func (dst *DeskPlayer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeskPlayer.Merge(dst, src)
}
func (m *DeskPlayer) XXX_Size() int {
	return xxx_messageInfo_DeskPlayer.Size(m)
}
func (m *DeskPlayer) XXX_DiscardUnknown() {
	xxx_messageInfo_DeskPlayer.DiscardUnknown(m)
}

var xxx_messageInfo_DeskPlayer proto.InternalMessageInfo

func (m *DeskPlayer) GetPlayerId() uint64 {
	if m != nil {
		return m.PlayerId
	}
	return 0
}

func (m *DeskPlayer) GetRobotLevel() int32 {
	if m != nil {
		return m.RobotLevel
	}
	return 0
}

func (m *DeskPlayer) GetSeat() uint32 {
	if m != nil {
		return m.Seat
	}
	return 0
}

// 创建桌子的请求
type CreateDeskRequest struct {
	GameId               uint32        `protobuf:"varint,1,opt,name=game_id,json=gameId,proto3" json:"game_id,omitempty"`
	Players              []*DeskPlayer `protobuf:"bytes,2,rep,name=players,proto3" json:"players,omitempty"`
	FixBanker            bool          `protobuf:"varint,3,opt,name=fix_banker,json=fixBanker,proto3" json:"fix_banker,omitempty"`
	BankerSeat           uint32        `protobuf:"varint,4,opt,name=banker_seat,json=bankerSeat,proto3" json:"banker_seat,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *CreateDeskRequest) Reset()         { *m = CreateDeskRequest{} }
func (m *CreateDeskRequest) String() string { return proto.CompactTextString(m) }
func (*CreateDeskRequest) ProtoMessage()    {}
func (*CreateDeskRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_room_mgr_bc05f510f82ddb45, []int{1}
}
func (m *CreateDeskRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateDeskRequest.Unmarshal(m, b)
}
func (m *CreateDeskRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateDeskRequest.Marshal(b, m, deterministic)
}
func (dst *CreateDeskRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateDeskRequest.Merge(dst, src)
}
func (m *CreateDeskRequest) XXX_Size() int {
	return xxx_messageInfo_CreateDeskRequest.Size(m)
}
func (m *CreateDeskRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateDeskRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateDeskRequest proto.InternalMessageInfo

func (m *CreateDeskRequest) GetGameId() uint32 {
	if m != nil {
		return m.GameId
	}
	return 0
}

func (m *CreateDeskRequest) GetPlayers() []*DeskPlayer {
	if m != nil {
		return m.Players
	}
	return nil
}

func (m *CreateDeskRequest) GetFixBanker() bool {
	if m != nil {
		return m.FixBanker
	}
	return false
}

func (m *CreateDeskRequest) GetBankerSeat() uint32 {
	if m != nil {
		return m.BankerSeat
	}
	return 0
}

// 创建桌子的回复
type CreateDeskResponse struct {
	ErrCode              RoomError `protobuf:"varint,1,opt,name=err_code,json=errCode,proto3,enum=roommgr.RoomError" json:"err_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *CreateDeskResponse) Reset()         { *m = CreateDeskResponse{} }
func (m *CreateDeskResponse) String() string { return proto.CompactTextString(m) }
func (*CreateDeskResponse) ProtoMessage()    {}
func (*CreateDeskResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_room_mgr_bc05f510f82ddb45, []int{2}
}
func (m *CreateDeskResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateDeskResponse.Unmarshal(m, b)
}
func (m *CreateDeskResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateDeskResponse.Marshal(b, m, deterministic)
}
func (dst *CreateDeskResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateDeskResponse.Merge(dst, src)
}
func (m *CreateDeskResponse) XXX_Size() int {
	return xxx_messageInfo_CreateDeskResponse.Size(m)
}
func (m *CreateDeskResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateDeskResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateDeskResponse proto.InternalMessageInfo

func (m *CreateDeskResponse) GetErrCode() RoomError {
	if m != nil {
		return m.ErrCode
	}
	return RoomError_SUCCESS
}

func init() {
	proto.RegisterType((*DeskPlayer)(nil), "roommgr.DeskPlayer")
	proto.RegisterType((*CreateDeskRequest)(nil), "roommgr.CreateDeskRequest")
	proto.RegisterType((*CreateDeskResponse)(nil), "roommgr.CreateDeskResponse")
	proto.RegisterEnum("roommgr.RoomError", RoomError_name, RoomError_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RoomMgrClient is the client API for RoomMgr service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RoomMgrClient interface {
	CreateDesk(ctx context.Context, in *CreateDeskRequest, opts ...grpc.CallOption) (*CreateDeskResponse, error)
}

type roomMgrClient struct {
	cc *grpc.ClientConn
}

func NewRoomMgrClient(cc *grpc.ClientConn) RoomMgrClient {
	return &roomMgrClient{cc}
}

func (c *roomMgrClient) CreateDesk(ctx context.Context, in *CreateDeskRequest, opts ...grpc.CallOption) (*CreateDeskResponse, error) {
	out := new(CreateDeskResponse)
	err := c.cc.Invoke(ctx, "/roommgr.RoomMgr/CreateDesk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RoomMgrServer is the server API for RoomMgr service.
type RoomMgrServer interface {
	CreateDesk(context.Context, *CreateDeskRequest) (*CreateDeskResponse, error)
}

func RegisterRoomMgrServer(s *grpc.Server, srv RoomMgrServer) {
	s.RegisterService(&_RoomMgr_serviceDesc, srv)
}

func _RoomMgr_CreateDesk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDeskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomMgrServer).CreateDesk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/roommgr.RoomMgr/CreateDesk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomMgrServer).CreateDesk(ctx, req.(*CreateDeskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RoomMgr_serviceDesc = grpc.ServiceDesc{
	ServiceName: "roommgr.RoomMgr",
	HandlerType: (*RoomMgrServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateDesk",
			Handler:    _RoomMgr_CreateDesk_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "room_mgr.proto",
}

func init() { proto.RegisterFile("room_mgr.proto", fileDescriptor_room_mgr_bc05f510f82ddb45) }

var fileDescriptor_room_mgr_bc05f510f82ddb45 = []byte{
	// 319 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xd1, 0x4a, 0x02, 0x51,
	0x10, 0x86, 0x5d, 0x35, 0x57, 0x47, 0x14, 0x9b, 0x2e, 0x5a, 0x94, 0x68, 0x91, 0x2e, 0x96, 0x40,
	0x2f, 0xec, 0x09, 0x6a, 0xb5, 0x10, 0x0c, 0xe2, 0x2c, 0xdd, 0xb6, 0xac, 0xed, 0xb8, 0x88, 0xae,
	0x63, 0xb3, 0x5b, 0xd8, 0xb3, 0xf4, 0xb2, 0x71, 0xce, 0xa6, 0x06, 0x75, 0x37, 0xf3, 0xcf, 0x70,
	0xbe, 0x8f, 0x33, 0xd0, 0x16, 0xe6, 0x34, 0x4c, 0x13, 0x19, 0x6e, 0x85, 0x73, 0x46, 0x5b, 0xf7,
	0x69, 0x22, 0xfd, 0x17, 0x80, 0x31, 0x65, 0xab, 0xa7, 0x75, 0xf4, 0x49, 0x82, 0x3d, 0x68, 0x6c,
	0x4d, 0x15, 0x2e, 0x63, 0xc7, 0x72, 0x2d, 0xaf, 0xaa, 0xea, 0x45, 0x30, 0x8d, 0xf1, 0x12, 0x9a,
	0xc2, 0x73, 0xce, 0xc3, 0x35, 0x7d, 0xd0, 0xda, 0x29, 0xbb, 0x96, 0x77, 0xa2, 0xc0, 0x44, 0x33,
	0x9d, 0x20, 0x42, 0x35, 0xa3, 0x28, 0x77, 0x2a, 0xae, 0xe5, 0xb5, 0x94, 0xa9, 0xfb, 0x5f, 0x16,
	0x9c, 0xfa, 0x42, 0x51, 0x4e, 0x1a, 0xa3, 0xe8, 0xed, 0x9d, 0xb2, 0x1c, 0xcf, 0xc1, 0x4e, 0xa2,
	0x94, 0xf6, 0x94, 0x96, 0xaa, 0xe9, 0x76, 0x1a, 0xe3, 0x00, 0xec, 0x82, 0x97, 0x39, 0x65, 0xb7,
	0xe2, 0x35, 0x47, 0x67, 0xc3, 0x1f, 0xd3, 0xe1, 0x51, 0x53, 0xed, 0x77, 0xf0, 0x02, 0x60, 0xb1,
	0xdc, 0x85, 0xf3, 0x68, 0xb3, 0x22, 0x31, 0xdc, 0xba, 0x6a, 0x2c, 0x96, 0xbb, 0x3b, 0x13, 0x68,
	0xe3, 0x62, 0x14, 0x1a, 0xaf, 0xaa, 0x41, 0x41, 0x11, 0x05, 0xda, 0xce, 0x07, 0xfc, 0x2d, 0x97,
	0x6d, 0x79, 0x93, 0x11, 0x0e, 0xa0, 0x4e, 0x22, 0xe1, 0x2b, 0xc7, 0x64, 0xf4, 0xda, 0x23, 0x3c,
	0x58, 0x28, 0xe6, 0x74, 0x22, 0xc2, 0xa2, 0x6c, 0x12, 0xf1, 0x39, 0xa6, 0xeb, 0x2b, 0x68, 0x1c,
	0x52, 0x6c, 0x82, 0x1d, 0x3c, 0xfb, 0xfe, 0x24, 0x08, 0x3a, 0x25, 0x04, 0xa8, 0xdd, 0xdf, 0x4e,
	0x67, 0x93, 0x71, 0xc7, 0x1a, 0x29, 0xb0, 0xf5, 0xd6, 0x63, 0x22, 0xf8, 0x00, 0x70, 0xa4, 0x62,
	0xf7, 0xf0, 0xf6, 0x9f, 0x7f, 0xea, 0xf6, 0xfe, 0x9d, 0x15, 0x9a, 0xfd, 0xd2, 0xbc, 0x66, 0x8e,
	0x79, 0xf3, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xe2, 0x78, 0x17, 0x5f, 0xde, 0x01, 0x00, 0x00,
}
