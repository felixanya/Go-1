// Code generated by protoc-gen-go.
// source: grpclb_backend_v1/backend.proto
// DO NOT EDIT!

/*
Package grpclb_backend_v1 is a generated protocol buffer package.

It is generated from these files:
	grpclb_backend_v1/backend.proto

It has these top-level messages:
	LoadRequest
	LoadResponse
*/
package grpclb_backend_v1

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

type LoadRequest struct {
}

func (m *LoadRequest) Reset()                    { *m = LoadRequest{} }
func (m *LoadRequest) String() string            { return proto.CompactTextString(m) }
func (*LoadRequest) ProtoMessage()               {}
func (*LoadRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type LoadResponse struct {
	Score int64 `protobuf:"varint,1,opt,name=score" json:"score,omitempty"`
}

func (m *LoadResponse) Reset()                    { *m = LoadResponse{} }
func (m *LoadResponse) String() string            { return proto.CompactTextString(m) }
func (*LoadResponse) ProtoMessage()               {}
func (*LoadResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *LoadResponse) GetScore() int64 {
	if m != nil {
		return m.Score
	}
	return 0
}

func init() {
	proto.RegisterType((*LoadRequest)(nil), "grpclb.backend.v1.LoadRequest")
	proto.RegisterType((*LoadResponse)(nil), "grpclb.backend.v1.LoadResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for LoadReport service

type LoadReportClient interface {
	Load(ctx context.Context, in *LoadRequest, opts ...grpc.CallOption) (*LoadResponse, error)
}

type loadReportClient struct {
	cc *grpc.ClientConn
}

func NewLoadReportClient(cc *grpc.ClientConn) LoadReportClient {
	return &loadReportClient{cc}
}

func (c *loadReportClient) Load(ctx context.Context, in *LoadRequest, opts ...grpc.CallOption) (*LoadResponse, error) {
	out := new(LoadResponse)
	err := grpc.Invoke(ctx, "/grpclb.backend.v1.LoadReport/Load", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for LoadReport service

type LoadReportServer interface {
	Load(context.Context, *LoadRequest) (*LoadResponse, error)
}

func RegisterLoadReportServer(s *grpc.Server, srv LoadReportServer) {
	s.RegisterService(&_LoadReport_serviceDesc, srv)
}

func _LoadReport_Load_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoadReportServer).Load(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpclb.backend.v1.LoadReport/Load",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoadReportServer).Load(ctx, req.(*LoadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LoadReport_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpclb.backend.v1.LoadReport",
	HandlerType: (*LoadReportServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Load",
			Handler:    _LoadReport_Load_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpclb_backend_v1/backend.proto",
}

func init() { proto.RegisterFile("grpclb_backend_v1/backend.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 141 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x92, 0x4f, 0x2f, 0x2a, 0x48,
	0xce, 0x49, 0x8a, 0x4f, 0x4a, 0x4c, 0xce, 0x4e, 0xcd, 0x4b, 0x89, 0x2f, 0x33, 0xd4, 0x87, 0x32,
	0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x04, 0x21, 0x0a, 0xf4, 0x60, 0xa2, 0x65, 0x86, 0x4a,
	0xbc, 0x5c, 0xdc, 0x3e, 0xf9, 0x89, 0x29, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x4a, 0x2a,
	0x5c, 0x3c, 0x10, 0x6e, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0xaa, 0x90, 0x08, 0x17, 0x6b, 0x71, 0x72,
	0x7e, 0x51, 0xaa, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x73, 0x10, 0x84, 0x63, 0x14, 0xca, 0xc5, 0x05,
	0x51, 0x55, 0x90, 0x5f, 0x54, 0x22, 0xe4, 0xce, 0xc5, 0x02, 0xe2, 0x09, 0xc9, 0xe9, 0x61, 0x18,
	0xaf, 0x87, 0x64, 0xb6, 0x94, 0x3c, 0x4e, 0x79, 0x88, 0x65, 0x49, 0x6c, 0x60, 0x57, 0x1a, 0x03,
	0x02, 0x00, 0x00, 0xff, 0xff, 0x6f, 0xb2, 0x90, 0x97, 0xc8, 0x00, 0x00, 0x00,
}
