// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: polestar.proto

package proto

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type RpcCode int32

const (
	RpcCode_OK            RpcCode = 0
	RpcCode_AccessDenied  RpcCode = 1001
	RpcCode_NotFound      RpcCode = 1002
	RpcCode_NoEnoughMoney RpcCode = 1003
	RpcCode_InternalError RpcCode = 1004
	RpcCode_WrongPassword RpcCode = 1005
	RpcCode_InvalidParam  RpcCode = 1006
)

var RpcCode_name = map[int32]string{
	0:    "OK",
	1001: "AccessDenied",
	1002: "NotFound",
	1003: "NoEnoughMoney",
	1004: "InternalError",
	1005: "WrongPassword",
	1006: "InvalidParam",
}

var RpcCode_value = map[string]int32{
	"OK":            0,
	"AccessDenied":  1001,
	"NotFound":      1002,
	"NoEnoughMoney": 1003,
	"InternalError": 1004,
	"WrongPassword": 1005,
	"InvalidParam":  1006,
}

func (x RpcCode) String() string {
	return proto.EnumName(RpcCode_name, int32(x))
}

func (RpcCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b39f9053dd603bf0, []int{0}
}

func init() {
	proto.RegisterEnum("walletpb.RpcCode", RpcCode_name, RpcCode_value)
}

func init() { proto.RegisterFile("polestar.proto", fileDescriptor_b39f9053dd603bf0) }

var fileDescriptor_b39f9053dd603bf0 = []byte{
	// 256 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8f, 0xcd, 0x4a, 0x03, 0x31,
	0x10, 0xc7, 0x5d, 0x0f, 0xdd, 0x12, 0xac, 0xc4, 0x1c, 0x3d, 0xec, 0x03, 0x08, 0x76, 0x11, 0x15,
	0xbc, 0xfa, 0x51, 0xa5, 0x88, 0xb5, 0x78, 0x11, 0xbc, 0x48, 0x76, 0x33, 0xa6, 0x81, 0x6c, 0x66,
	0x99, 0x64, 0xb7, 0x78, 0xf5, 0xe9, 0x3c, 0xfa, 0x08, 0xba, 0x37, 0x3f, 0xf1, 0x11, 0xc4, 0x28,
	0xa5, 0xa7, 0xf9, 0xff, 0x7f, 0xcc, 0x0c, 0xfc, 0xd8, 0x7a, 0x8d, 0x16, 0x7c, 0x90, 0x34, 0xac,
	0x09, 0x03, 0x8a, 0xfe, 0x5c, 0x5a, 0x0b, 0xa1, 0x2e, 0x36, 0xb7, 0xb5, 0x09, 0xb3, 0xa6, 0x18,
	0x96, 0x58, 0xe5, 0x1a, 0x35, 0xe6, 0x71, 0xa1, 0x68, 0xee, 0x62, 0x8b, 0x25, 0xa6, 0xbf, 0xc3,
	0xad, 0x87, 0x84, 0xa5, 0x57, 0x75, 0x79, 0x8c, 0x0a, 0x44, 0x8f, 0xad, 0x5e, 0x9e, 0xf3, 0x15,
	0xb1, 0xc1, 0xd6, 0x0e, 0xcb, 0x12, 0xbc, 0x3f, 0x01, 0x67, 0x40, 0xf1, 0xd7, 0x54, 0x0c, 0x58,
	0x7f, 0x82, 0xe1, 0x14, 0x1b, 0xa7, 0xf8, 0x5b, 0x2a, 0x04, 0x1b, 0x4c, 0x70, 0xe4, 0xb0, 0xd1,
	0xb3, 0x0b, 0x74, 0x70, 0xcf, 0xdf, 0x23, 0x1b, 0xbb, 0x00, 0xe4, 0xa4, 0x1d, 0x11, 0x21, 0xf1,
	0x8f, 0xc8, 0xae, 0x09, 0x9d, 0x9e, 0x4a, 0xef, 0xe7, 0x48, 0x8a, 0x7f, 0xa6, 0xbf, 0xdf, 0xc7,
	0xae, 0x95, 0xd6, 0xa8, 0xa9, 0x24, 0x59, 0xf1, 0xaf, 0xf4, 0xe8, 0xec, 0xfb, 0x25, 0x4b, 0x1e,
	0xbb, 0x2c, 0x79, 0xea, 0xb2, 0xe4, 0xb9, 0xcb, 0x92, 0x9b, 0xfd, 0x25, 0x8b, 0xd2, 0xb4, 0x10,
	0x76, 0xf6, 0x0e, 0xf2, 0x85, 0xb3, 0xa9, 0x16, 0xf9, 0xd6, 0x03, 0xb5, 0x40, 0xff, 0x96, 0xbd,
	0x38, 0x76, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x95, 0xd2, 0xf6, 0xeb, 0x1f, 0x01, 0x00, 0x00,
}
