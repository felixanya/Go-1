// Code generated by protoc-gen-go. DO NOT EDIT.
// source: errors.proto

package user

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

type ErrCode int32

const (
	ErrCode_EC_SUCCESS ErrCode = 0
	ErrCode_EC_FAIL    ErrCode = 1
	ErrCode_EC_Args    ErrCode = 2
	ErrCode_EC_EMPTY   ErrCode = 3
)

var ErrCode_name = map[int32]string{
	0: "EC_SUCCESS",
	1: "EC_FAIL",
	2: "EC_Args",
	3: "EC_EMPTY",
}
var ErrCode_value = map[string]int32{
	"EC_SUCCESS": 0,
	"EC_FAIL":    1,
	"EC_Args":    2,
	"EC_EMPTY":   3,
}

func (x ErrCode) String() string {
	return proto.EnumName(ErrCode_name, int32(x))
}
func (ErrCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_errors_5a7d5aafc675a956, []int{0}
}

func init() {
	proto.RegisterEnum("user.ErrCode", ErrCode_name, ErrCode_value)
}

func init() { proto.RegisterFile("errors.proto", fileDescriptor_errors_5a7d5aafc675a956) }

var fileDescriptor_errors_5a7d5aafc675a956 = []byte{
	// 111 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0x2d, 0x2a, 0xca,
	0x2f, 0x2a, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x29, 0x2d, 0x4e, 0x2d, 0xd2, 0x72,
	0xe4, 0x62, 0x77, 0x2d, 0x2a, 0x72, 0xce, 0x4f, 0x49, 0x15, 0xe2, 0xe3, 0xe2, 0x72, 0x75, 0x8e,
	0x0f, 0x0e, 0x75, 0x76, 0x76, 0x0d, 0x0e, 0x16, 0x60, 0x10, 0xe2, 0xe6, 0x62, 0x77, 0x75, 0x8e,
	0x77, 0x73, 0xf4, 0xf4, 0x11, 0x60, 0x84, 0x72, 0x1c, 0x8b, 0xd2, 0x8b, 0x05, 0x98, 0x84, 0x78,
	0xb8, 0x38, 0x5c, 0x9d, 0xe3, 0x5d, 0x7d, 0x03, 0x42, 0x22, 0x05, 0x98, 0x93, 0xd8, 0xc0, 0xe6,
	0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x0a, 0xe7, 0x8b, 0x2d, 0x5f, 0x00, 0x00, 0x00,
}
