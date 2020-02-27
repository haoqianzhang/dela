// Code generated by protoc-gen-go. DO NOT EDIT.
// source: messages.proto

package bls

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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type PublicKeyProto struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PublicKeyProto) Reset()         { *m = PublicKeyProto{} }
func (m *PublicKeyProto) String() string { return proto.CompactTextString(m) }
func (*PublicKeyProto) ProtoMessage()    {}
func (*PublicKeyProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{0}
}

func (m *PublicKeyProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublicKeyProto.Unmarshal(m, b)
}
func (m *PublicKeyProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublicKeyProto.Marshal(b, m, deterministic)
}
func (m *PublicKeyProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublicKeyProto.Merge(m, src)
}
func (m *PublicKeyProto) XXX_Size() int {
	return xxx_messageInfo_PublicKeyProto.Size(m)
}
func (m *PublicKeyProto) XXX_DiscardUnknown() {
	xxx_messageInfo_PublicKeyProto.DiscardUnknown(m)
}

var xxx_messageInfo_PublicKeyProto proto.InternalMessageInfo

func (m *PublicKeyProto) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type SignatureProto struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignatureProto) Reset()         { *m = SignatureProto{} }
func (m *SignatureProto) String() string { return proto.CompactTextString(m) }
func (*SignatureProto) ProtoMessage()    {}
func (*SignatureProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{1}
}

func (m *SignatureProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignatureProto.Unmarshal(m, b)
}
func (m *SignatureProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignatureProto.Marshal(b, m, deterministic)
}
func (m *SignatureProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignatureProto.Merge(m, src)
}
func (m *SignatureProto) XXX_Size() int {
	return xxx_messageInfo_SignatureProto.Size(m)
}
func (m *SignatureProto) XXX_DiscardUnknown() {
	xxx_messageInfo_SignatureProto.DiscardUnknown(m)
}

var xxx_messageInfo_SignatureProto proto.InternalMessageInfo

func (m *SignatureProto) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*PublicKeyProto)(nil), "bls.publicKeyProto")
	proto.RegisterType((*SignatureProto)(nil), "bls.signatureProto")
}

func init() { proto.RegisterFile("messages.proto", fileDescriptor_4dc296cbfe5ffcd5) }

var fileDescriptor_4dc296cbfe5ffcd5 = []byte{
	// 95 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcb, 0x4d, 0x2d, 0x2e,
	0x4e, 0x4c, 0x4f, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0xca, 0x29, 0x56,
	0x52, 0xe1, 0xe2, 0x2b, 0x28, 0x4d, 0xca, 0xc9, 0x4c, 0xf6, 0x4e, 0xad, 0x0c, 0x00, 0x0b, 0x0b,
	0x71, 0xb1, 0xa4, 0x24, 0x96, 0x24, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0xf0, 0x04, 0x81, 0xd9, 0x20,
	0x55, 0xc5, 0x99, 0xe9, 0x79, 0x89, 0x25, 0xa5, 0x45, 0xa9, 0x38, 0x55, 0x25, 0xb1, 0x81, 0xcd,
	0x35, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x8e, 0xdd, 0x3a, 0xea, 0x69, 0x00, 0x00, 0x00,
}