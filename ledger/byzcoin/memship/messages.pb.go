// Code generated by protoc-gen-go. DO NOT EDIT.
// source: messages.proto

package memship

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
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

// Task is the message for a client task.
type Task struct {
	Remove               []uint32 `protobuf:"varint,1,rep,packed,name=remove,proto3" json:"remove,omitempty"`
	Addr                 []byte   `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	PublicKey            *any.Any `protobuf:"bytes,3,opt,name=publicKey,proto3" json:"publicKey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Task) Reset()         { *m = Task{} }
func (m *Task) String() string { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()    {}
func (*Task) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{0}
}

func (m *Task) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Task.Unmarshal(m, b)
}
func (m *Task) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Task.Marshal(b, m, deterministic)
}
func (m *Task) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Task.Merge(m, src)
}
func (m *Task) XXX_Size() int {
	return xxx_messageInfo_Task.Size(m)
}
func (m *Task) XXX_DiscardUnknown() {
	xxx_messageInfo_Task.DiscardUnknown(m)
}

var xxx_messageInfo_Task proto.InternalMessageInfo

func (m *Task) GetRemove() []uint32 {
	if m != nil {
		return m.Remove
	}
	return nil
}

func (m *Task) GetAddr() []byte {
	if m != nil {
		return m.Addr
	}
	return nil
}

func (m *Task) GetPublicKey() *any.Any {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func init() {
	proto.RegisterType((*Task)(nil), "memship.Task")
}

func init() {
	proto.RegisterFile("messages.proto", fileDescriptor_4dc296cbfe5ffcd5)
}

var fileDescriptor_4dc296cbfe5ffcd5 = []byte{
	// 152 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcb, 0x4d, 0x2d, 0x2e,
	0x4e, 0x4c, 0x4f, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xcf, 0x4d, 0xcd, 0x2d,
	0xce, 0xc8, 0x2c, 0x90, 0x92, 0x4c, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x07, 0x0b, 0x27, 0x95,
	0xa6, 0xe9, 0x27, 0xe6, 0x55, 0x42, 0xd4, 0x28, 0xa5, 0x71, 0xb1, 0x84, 0x24, 0x16, 0x67, 0x0b,
	0x89, 0x71, 0xb1, 0x15, 0xa5, 0xe6, 0xe6, 0x97, 0xa5, 0x4a, 0x30, 0x2a, 0x30, 0x6b, 0xf0, 0x06,
	0x41, 0x79, 0x42, 0x42, 0x5c, 0x2c, 0x89, 0x29, 0x29, 0x45, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x3c,
	0x41, 0x60, 0xb6, 0x90, 0x11, 0x17, 0x67, 0x41, 0x69, 0x52, 0x4e, 0x66, 0xb2, 0x77, 0x6a, 0xa5,
	0x04, 0xb3, 0x02, 0xa3, 0x06, 0xb7, 0x91, 0x88, 0x1e, 0xc4, 0x0a, 0x3d, 0x98, 0x15, 0x7a, 0x8e,
	0x79, 0x95, 0x41, 0x08, 0x65, 0x49, 0x6c, 0x60, 0x09, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xc9, 0xa0, 0x09, 0x26, 0xa4, 0x00, 0x00, 0x00,
}