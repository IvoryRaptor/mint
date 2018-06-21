// Code generated by protoc-gen-go. DO NOT EDIT.
// source: skin_message.proto

/*
Package skin is a generated protocol buffer package.

It is generated from these files:
	skin_message.proto

It has these top-level messages:
	SkinPayload
*/
package skin

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

type SkinPayload struct {
	Matrix string `protobuf:"bytes,1,opt,name=matrix" json:"matrix,omitempty"`
	Angler string `protobuf:"bytes,2,opt,name=angler" json:"angler,omitempty"`
	Number string `protobuf:"bytes,3,opt,name=number" json:"number,omitempty"`
	Time   int64  `protobuf:"varint,4,opt,name=time" json:"time,omitempty"`
}

func (m *SkinPayload) Reset()                    { *m = SkinPayload{} }
func (m *SkinPayload) String() string            { return proto.CompactTextString(m) }
func (*SkinPayload) ProtoMessage()               {}
func (*SkinPayload) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SkinPayload) GetMatrix() string {
	if m != nil {
		return m.Matrix
	}
	return ""
}

func (m *SkinPayload) GetAngler() string {
	if m != nil {
		return m.Angler
	}
	return ""
}

func (m *SkinPayload) GetNumber() string {
	if m != nil {
		return m.Number
	}
	return ""
}

func (m *SkinPayload) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func init() {
	proto.RegisterType((*SkinPayload)(nil), "skin.SkinPayload")
}

func init() { proto.RegisterFile("skin_message.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 131 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0xce, 0xce, 0xcc,
	0x8b, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62,
	0x01, 0x89, 0x29, 0x65, 0x72, 0x71, 0x07, 0x67, 0x67, 0xe6, 0x05, 0x24, 0x56, 0xe6, 0xe4, 0x27,
	0xa6, 0x08, 0x89, 0x71, 0xb1, 0xe5, 0x26, 0x96, 0x14, 0x65, 0x56, 0x48, 0x30, 0x2a, 0x30, 0x6a,
	0x70, 0x06, 0x41, 0x79, 0x20, 0xf1, 0xc4, 0xbc, 0xf4, 0x9c, 0xd4, 0x22, 0x09, 0x26, 0x88, 0x38,
	0x84, 0x07, 0x12, 0xcf, 0x2b, 0xcd, 0x4d, 0x4a, 0x2d, 0x92, 0x60, 0x86, 0x88, 0x43, 0x78, 0x42,
	0x42, 0x5c, 0x2c, 0x25, 0x99, 0xb9, 0xa9, 0x12, 0x2c, 0x0a, 0x8c, 0x1a, 0xcc, 0x41, 0x60, 0x76,
	0x12, 0x1b, 0xd8, 0x5e, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4a, 0x7d, 0xd0, 0x85, 0x8d,
	0x00, 0x00, 0x00,
}