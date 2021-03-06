// Code generated by protoc-gen-go. DO NOT EDIT.
// source: idserver.proto

package idsvr

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

// 返回结果
type ResultStat int32

const (
	ResultStat_SUCCEED ResultStat = 0
	ResultStat_FAILED  ResultStat = 1
)

var ResultStat_name = map[int32]string{
	0: "SUCCEED",
	1: "FAILED",
}
var ResultStat_value = map[string]int32{
	"SUCCEED": 0,
	"FAILED":  1,
}

func (x ResultStat) String() string {
	return proto.EnumName(ResultStat_name, int32(x))
}
func (ResultStat) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_idserver_d0c6421e1112f72d, []int{0}
}

// 生成一个新的playerId请求
type NewPlayerIdReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewPlayerIdReq) Reset()         { *m = NewPlayerIdReq{} }
func (m *NewPlayerIdReq) String() string { return proto.CompactTextString(m) }
func (*NewPlayerIdReq) ProtoMessage()    {}
func (*NewPlayerIdReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_idserver_d0c6421e1112f72d, []int{0}
}
func (m *NewPlayerIdReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewPlayerIdReq.Unmarshal(m, b)
}
func (m *NewPlayerIdReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewPlayerIdReq.Marshal(b, m, deterministic)
}
func (dst *NewPlayerIdReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewPlayerIdReq.Merge(dst, src)
}
func (m *NewPlayerIdReq) XXX_Size() int {
	return xxx_messageInfo_NewPlayerIdReq.Size(m)
}
func (m *NewPlayerIdReq) XXX_DiscardUnknown() {
	xxx_messageInfo_NewPlayerIdReq.DiscardUnknown(m)
}

var xxx_messageInfo_NewPlayerIdReq proto.InternalMessageInfo

// 生成一个新的playerId回复
type NewPlayerIdRsp struct {
	ErrCode              ResultStat `protobuf:"varint,1,opt,name=err_code,json=errCode,proto3,enum=idsvr.ResultStat" json:"err_code,omitempty"`
	ErrDesc              string     `protobuf:"bytes,2,opt,name=err_desc,json=errDesc,proto3" json:"err_desc,omitempty"`
	NewId                uint64     `protobuf:"varint,3,opt,name=new_id,json=newId,proto3" json:"new_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *NewPlayerIdRsp) Reset()         { *m = NewPlayerIdRsp{} }
func (m *NewPlayerIdRsp) String() string { return proto.CompactTextString(m) }
func (*NewPlayerIdRsp) ProtoMessage()    {}
func (*NewPlayerIdRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_idserver_d0c6421e1112f72d, []int{1}
}
func (m *NewPlayerIdRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewPlayerIdRsp.Unmarshal(m, b)
}
func (m *NewPlayerIdRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewPlayerIdRsp.Marshal(b, m, deterministic)
}
func (dst *NewPlayerIdRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewPlayerIdRsp.Merge(dst, src)
}
func (m *NewPlayerIdRsp) XXX_Size() int {
	return xxx_messageInfo_NewPlayerIdRsp.Size(m)
}
func (m *NewPlayerIdRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_NewPlayerIdRsp.DiscardUnknown(m)
}

var xxx_messageInfo_NewPlayerIdRsp proto.InternalMessageInfo

func (m *NewPlayerIdRsp) GetErrCode() ResultStat {
	if m != nil {
		return m.ErrCode
	}
	return ResultStat_SUCCEED
}

func (m *NewPlayerIdRsp) GetErrDesc() string {
	if m != nil {
		return m.ErrDesc
	}
	return ""
}

func (m *NewPlayerIdRsp) GetNewId() uint64 {
	if m != nil {
		return m.NewId
	}
	return 0
}

// 生成一个新的showId请求
type NewShowIdReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewShowIdReq) Reset()         { *m = NewShowIdReq{} }
func (m *NewShowIdReq) String() string { return proto.CompactTextString(m) }
func (*NewShowIdReq) ProtoMessage()    {}
func (*NewShowIdReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_idserver_d0c6421e1112f72d, []int{2}
}
func (m *NewShowIdReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewShowIdReq.Unmarshal(m, b)
}
func (m *NewShowIdReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewShowIdReq.Marshal(b, m, deterministic)
}
func (dst *NewShowIdReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewShowIdReq.Merge(dst, src)
}
func (m *NewShowIdReq) XXX_Size() int {
	return xxx_messageInfo_NewShowIdReq.Size(m)
}
func (m *NewShowIdReq) XXX_DiscardUnknown() {
	xxx_messageInfo_NewShowIdReq.DiscardUnknown(m)
}

var xxx_messageInfo_NewShowIdReq proto.InternalMessageInfo

// 生成一个新的showId回复
type NewShowIdRsp struct {
	ErrCode              ResultStat `protobuf:"varint,1,opt,name=err_code,json=errCode,proto3,enum=idsvr.ResultStat" json:"err_code,omitempty"`
	ErrDesc              string     `protobuf:"bytes,2,opt,name=err_desc,json=errDesc,proto3" json:"err_desc,omitempty"`
	NewId                uint64     `protobuf:"varint,3,opt,name=new_id,json=newId,proto3" json:"new_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *NewShowIdRsp) Reset()         { *m = NewShowIdRsp{} }
func (m *NewShowIdRsp) String() string { return proto.CompactTextString(m) }
func (*NewShowIdRsp) ProtoMessage()    {}
func (*NewShowIdRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_idserver_d0c6421e1112f72d, []int{3}
}
func (m *NewShowIdRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewShowIdRsp.Unmarshal(m, b)
}
func (m *NewShowIdRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewShowIdRsp.Marshal(b, m, deterministic)
}
func (dst *NewShowIdRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewShowIdRsp.Merge(dst, src)
}
func (m *NewShowIdRsp) XXX_Size() int {
	return xxx_messageInfo_NewShowIdRsp.Size(m)
}
func (m *NewShowIdRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_NewShowIdRsp.DiscardUnknown(m)
}

var xxx_messageInfo_NewShowIdRsp proto.InternalMessageInfo

func (m *NewShowIdRsp) GetErrCode() ResultStat {
	if m != nil {
		return m.ErrCode
	}
	return ResultStat_SUCCEED
}

func (m *NewShowIdRsp) GetErrDesc() string {
	if m != nil {
		return m.ErrDesc
	}
	return ""
}

func (m *NewShowIdRsp) GetNewId() uint64 {
	if m != nil {
		return m.NewId
	}
	return 0
}

// 生成一个新的playerId和showId请求
type NewPlayerShowIdReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewPlayerShowIdReq) Reset()         { *m = NewPlayerShowIdReq{} }
func (m *NewPlayerShowIdReq) String() string { return proto.CompactTextString(m) }
func (*NewPlayerShowIdReq) ProtoMessage()    {}
func (*NewPlayerShowIdReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_idserver_d0c6421e1112f72d, []int{4}
}
func (m *NewPlayerShowIdReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewPlayerShowIdReq.Unmarshal(m, b)
}
func (m *NewPlayerShowIdReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewPlayerShowIdReq.Marshal(b, m, deterministic)
}
func (dst *NewPlayerShowIdReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewPlayerShowIdReq.Merge(dst, src)
}
func (m *NewPlayerShowIdReq) XXX_Size() int {
	return xxx_messageInfo_NewPlayerShowIdReq.Size(m)
}
func (m *NewPlayerShowIdReq) XXX_DiscardUnknown() {
	xxx_messageInfo_NewPlayerShowIdReq.DiscardUnknown(m)
}

var xxx_messageInfo_NewPlayerShowIdReq proto.InternalMessageInfo

// 生成一个新的playerId和showId回复
type NewPlayerShowIdRsp struct {
	ErrCode              ResultStat `protobuf:"varint,1,opt,name=err_code,json=errCode,proto3,enum=idsvr.ResultStat" json:"err_code,omitempty"`
	ErrDesc              string     `protobuf:"bytes,2,opt,name=err_desc,json=errDesc,proto3" json:"err_desc,omitempty"`
	PlayerId             uint64     `protobuf:"varint,3,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
	ShowId               uint64     `protobuf:"varint,4,opt,name=show_id,json=showId,proto3" json:"show_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *NewPlayerShowIdRsp) Reset()         { *m = NewPlayerShowIdRsp{} }
func (m *NewPlayerShowIdRsp) String() string { return proto.CompactTextString(m) }
func (*NewPlayerShowIdRsp) ProtoMessage()    {}
func (*NewPlayerShowIdRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_idserver_d0c6421e1112f72d, []int{5}
}
func (m *NewPlayerShowIdRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewPlayerShowIdRsp.Unmarshal(m, b)
}
func (m *NewPlayerShowIdRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewPlayerShowIdRsp.Marshal(b, m, deterministic)
}
func (dst *NewPlayerShowIdRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewPlayerShowIdRsp.Merge(dst, src)
}
func (m *NewPlayerShowIdRsp) XXX_Size() int {
	return xxx_messageInfo_NewPlayerShowIdRsp.Size(m)
}
func (m *NewPlayerShowIdRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_NewPlayerShowIdRsp.DiscardUnknown(m)
}

var xxx_messageInfo_NewPlayerShowIdRsp proto.InternalMessageInfo

func (m *NewPlayerShowIdRsp) GetErrCode() ResultStat {
	if m != nil {
		return m.ErrCode
	}
	return ResultStat_SUCCEED
}

func (m *NewPlayerShowIdRsp) GetErrDesc() string {
	if m != nil {
		return m.ErrDesc
	}
	return ""
}

func (m *NewPlayerShowIdRsp) GetPlayerId() uint64 {
	if m != nil {
		return m.PlayerId
	}
	return 0
}

func (m *NewPlayerShowIdRsp) GetShowId() uint64 {
	if m != nil {
		return m.ShowId
	}
	return 0
}

func init() {
	proto.RegisterType((*NewPlayerIdReq)(nil), "idsvr.NewPlayerIdReq")
	proto.RegisterType((*NewPlayerIdRsp)(nil), "idsvr.NewPlayerIdRsp")
	proto.RegisterType((*NewShowIdReq)(nil), "idsvr.NewShowIdReq")
	proto.RegisterType((*NewShowIdRsp)(nil), "idsvr.NewShowIdRsp")
	proto.RegisterType((*NewPlayerShowIdReq)(nil), "idsvr.NewPlayerShowIdReq")
	proto.RegisterType((*NewPlayerShowIdRsp)(nil), "idsvr.NewPlayerShowIdRsp")
	proto.RegisterEnum("idsvr.ResultStat", ResultStat_name, ResultStat_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// IdserviceClient is the client API for Idservice service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type IdserviceClient interface {
	// 生成一个新的playerId和showId
	NewPlayerShowId(ctx context.Context, in *NewPlayerShowIdReq, opts ...grpc.CallOption) (*NewPlayerShowIdRsp, error)
	// 生成一个新的playerId
	NewPlayerId(ctx context.Context, in *NewPlayerIdReq, opts ...grpc.CallOption) (*NewPlayerIdRsp, error)
	// 生成一个新的showId
	NewShowId(ctx context.Context, in *NewShowIdReq, opts ...grpc.CallOption) (*NewShowIdRsp, error)
}

type idserviceClient struct {
	cc *grpc.ClientConn
}

func NewIdserviceClient(cc *grpc.ClientConn) IdserviceClient {
	return &idserviceClient{cc}
}

func (c *idserviceClient) NewPlayerShowId(ctx context.Context, in *NewPlayerShowIdReq, opts ...grpc.CallOption) (*NewPlayerShowIdRsp, error) {
	out := new(NewPlayerShowIdRsp)
	err := c.cc.Invoke(ctx, "/idsvr.idservice/NewPlayerShowId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *idserviceClient) NewPlayerId(ctx context.Context, in *NewPlayerIdReq, opts ...grpc.CallOption) (*NewPlayerIdRsp, error) {
	out := new(NewPlayerIdRsp)
	err := c.cc.Invoke(ctx, "/idsvr.idservice/NewPlayerId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *idserviceClient) NewShowId(ctx context.Context, in *NewShowIdReq, opts ...grpc.CallOption) (*NewShowIdRsp, error) {
	out := new(NewShowIdRsp)
	err := c.cc.Invoke(ctx, "/idsvr.idservice/NewShowId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IdserviceServer is the server API for Idservice service.
type IdserviceServer interface {
	// 生成一个新的playerId和showId
	NewPlayerShowId(context.Context, *NewPlayerShowIdReq) (*NewPlayerShowIdRsp, error)
	// 生成一个新的playerId
	NewPlayerId(context.Context, *NewPlayerIdReq) (*NewPlayerIdRsp, error)
	// 生成一个新的showId
	NewShowId(context.Context, *NewShowIdReq) (*NewShowIdRsp, error)
}

func RegisterIdserviceServer(s *grpc.Server, srv IdserviceServer) {
	s.RegisterService(&_Idservice_serviceDesc, srv)
}

func _Idservice_NewPlayerShowId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewPlayerShowIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdserviceServer).NewPlayerShowId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/idsvr.idservice/NewPlayerShowId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdserviceServer).NewPlayerShowId(ctx, req.(*NewPlayerShowIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Idservice_NewPlayerId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewPlayerIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdserviceServer).NewPlayerId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/idsvr.idservice/NewPlayerId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdserviceServer).NewPlayerId(ctx, req.(*NewPlayerIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Idservice_NewShowId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewShowIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdserviceServer).NewShowId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/idsvr.idservice/NewShowId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdserviceServer).NewShowId(ctx, req.(*NewShowIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Idservice_serviceDesc = grpc.ServiceDesc{
	ServiceName: "idsvr.idservice",
	HandlerType: (*IdserviceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewPlayerShowId",
			Handler:    _Idservice_NewPlayerShowId_Handler,
		},
		{
			MethodName: "NewPlayerId",
			Handler:    _Idservice_NewPlayerId_Handler,
		},
		{
			MethodName: "NewShowId",
			Handler:    _Idservice_NewShowId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "idserver.proto",
}

func init() { proto.RegisterFile("idserver.proto", fileDescriptor_idserver_d0c6421e1112f72d) }

var fileDescriptor_idserver_d0c6421e1112f72d = []byte{
	// 312 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x92, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0xbb, 0xda, 0xa6, 0xcd, 0x54, 0x62, 0x1d, 0x2d, 0xb6, 0x7a, 0x29, 0x01, 0xa1, 0x88,
	0xf4, 0x50, 0x0f, 0x9e, 0x3c, 0x48, 0x13, 0x21, 0x20, 0x45, 0x12, 0x3c, 0x97, 0x9a, 0x1d, 0x68,
	0xa0, 0x64, 0xd7, 0xdd, 0xd8, 0xe0, 0x73, 0xf8, 0x5e, 0x3e, 0x93, 0xb8, 0x29, 0x4d, 0x6d, 0xf4,
	0x26, 0x1e, 0xf7, 0x9b, 0x7f, 0xbf, 0x6f, 0x67, 0xc0, 0x49, 0xb8, 0x26, 0xb5, 0x22, 0x35, 0x92,
	0x4a, 0x64, 0x02, 0x1b, 0x09, 0xd7, 0x2b, 0xe5, 0x76, 0xc0, 0x99, 0x52, 0xfe, 0xb8, 0x9c, 0xbf,
	0x91, 0x0a, 0x78, 0x48, 0x2f, 0xae, 0xfc, 0xae, 0x68, 0x89, 0x57, 0xd0, 0x22, 0xa5, 0x66, 0xb1,
	0xe0, 0xd4, 0x63, 0x03, 0x36, 0x74, 0xc6, 0x47, 0x23, 0x53, 0x3d, 0x0a, 0x49, 0xbf, 0x2e, 0xb3,
	0x28, 0x9b, 0x67, 0x61, 0x93, 0x94, 0x9a, 0x08, 0x4e, 0xd8, 0x2f, 0xb2, 0x39, 0xe9, 0xb8, 0xb7,
	0x37, 0x60, 0x43, 0xdb, 0x84, 0x3c, 0xd2, 0x31, 0x76, 0xc1, 0x4a, 0x29, 0x9f, 0x25, 0xbc, 0xb7,
	0x3f, 0x60, 0xc3, 0x7a, 0xd8, 0x48, 0x29, 0x0f, 0xb8, 0xeb, 0xc0, 0xc1, 0x94, 0xf2, 0x68, 0x21,
	0xf2, 0x82, 0x20, 0xdd, 0x7e, 0xff, 0xc3, 0xfc, 0x13, 0xc0, 0x8d, 0xe3, 0x92, 0xe2, 0x9d, 0x55,
	0xe5, 0xbf, 0x84, 0x39, 0x07, 0x5b, 0x9a, 0xde, 0x25, 0x4f, 0x4b, 0xae, 0x7f, 0x1d, 0x4f, 0xa1,
	0xa9, 0x17, 0xc2, 0xa0, 0xd6, 0x4d, 0xc8, 0xd2, 0x86, 0xe0, 0xf2, 0x02, 0xa0, 0x9c, 0x83, 0x6d,
	0x68, 0x46, 0x4f, 0x93, 0x89, 0xef, 0x7b, 0x9d, 0x1a, 0x02, 0x58, 0xf7, 0x77, 0xc1, 0x83, 0xef,
	0x75, 0xd8, 0xf8, 0x83, 0x81, 0x5d, 0x2c, 0x3c, 0x89, 0x09, 0x03, 0x38, 0xdc, 0x71, 0x82, 0xfd,
	0x35, 0x74, 0xd5, 0xf8, 0xd9, 0x6f, 0x21, 0x2d, 0xdd, 0x1a, 0xde, 0x42, 0x7b, 0xeb, 0x3a, 0xb0,
	0xbb, 0x9b, 0x5b, 0xb4, 0xf8, 0x49, 0x36, 0xe5, 0x37, 0x60, 0x6f, 0x56, 0x8b, 0xc7, 0x65, 0x56,
	0x39, 0xbd, 0x2a, 0x7e, 0x15, 0x3e, 0x5b, 0xe6, 0x6a, 0xaf, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff,
	0x04, 0x1b, 0xee, 0xab, 0xc7, 0x02, 0x00, 0x00,
}
