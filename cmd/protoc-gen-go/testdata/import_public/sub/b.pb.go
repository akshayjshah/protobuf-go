// Code generated by protoc-gen-go. DO NOT EDIT.
// source: import_public/sub/b.proto

package sub

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type M2 struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *M2) Reset()         { *m = M2{} }
func (m *M2) String() string { return proto.CompactTextString(m) }
func (*M2) ProtoMessage()    {}
func (*M2) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc66afda3d7c2232, []int{0}
}

func (m *M2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_M2.Unmarshal(m, b)
}
func (m *M2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_M2.Marshal(b, m, deterministic)
}
func (m *M2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_M2.Merge(m, src)
}
func (m *M2) XXX_Size() int {
	return xxx_messageInfo_M2.Size(m)
}
func (m *M2) XXX_DiscardUnknown() {
	xxx_messageInfo_M2.DiscardUnknown(m)
}

var xxx_messageInfo_M2 proto.InternalMessageInfo

func init() {
	proto.RegisterType((*M2)(nil), "goproto.protoc.import_public.sub.M2")
}

func init() { proto.RegisterFile("import_public/sub/b.proto", fileDescriptor_fc66afda3d7c2232) }

var fileDescriptor_fc66afda3d7c2232 = []byte{
	// 134 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcc, 0xcc, 0x2d, 0xc8,
	0x2f, 0x2a, 0x89, 0x2f, 0x28, 0x4d, 0xca, 0xc9, 0x4c, 0xd6, 0x2f, 0x2e, 0x4d, 0xd2, 0x4f, 0xd2,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x52, 0x48, 0xcf, 0x07, 0x33, 0x20, 0xdc, 0x64, 0x3d, 0x14,
	0x95, 0x7a, 0xc5, 0xa5, 0x49, 0x4a, 0x2c, 0x5c, 0x4c, 0xbe, 0x46, 0x4e, 0x3e, 0x51, 0x5e, 0xe9,
	0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0xe9, 0xf9, 0x39, 0x89, 0x79, 0xe9,
	0xfa, 0x60, 0x3d, 0x49, 0xa5, 0x69, 0xfa, 0x65, 0x46, 0xfa, 0xc9, 0xb9, 0x29, 0x10, 0x7e, 0xb2,
	0x6e, 0x7a, 0x6a, 0x9e, 0x6e, 0x7a, 0xbe, 0x7e, 0x49, 0x6a, 0x71, 0x49, 0x4a, 0x62, 0x49, 0xa2,
	0x3e, 0x86, 0xed, 0x49, 0x6c, 0x60, 0x95, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x46, 0xb8,
	0x19, 0x87, 0x99, 0x00, 0x00, 0x00,
}