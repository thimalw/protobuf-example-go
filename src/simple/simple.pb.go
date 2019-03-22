// Code generated by protoc-gen-go. DO NOT EDIT.
// source: simple/simple.proto

package simplepb

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

type SimpleMessage struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	IsSimple             bool     `protobuf:"varint,2,opt,name=is_simple,json=isSimple,proto3" json:"is_simple,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	SampleList           []int32  `protobuf:"varint,4,rep,packed,name=sample_list,json=sampleList,proto3" json:"sample_list,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SimpleMessage) Reset()         { *m = SimpleMessage{} }
func (m *SimpleMessage) String() string { return proto.CompactTextString(m) }
func (*SimpleMessage) ProtoMessage()    {}
func (*SimpleMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b3c868e94d57426, []int{0}
}

func (m *SimpleMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SimpleMessage.Unmarshal(m, b)
}
func (m *SimpleMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SimpleMessage.Marshal(b, m, deterministic)
}
func (m *SimpleMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimpleMessage.Merge(m, src)
}
func (m *SimpleMessage) XXX_Size() int {
	return xxx_messageInfo_SimpleMessage.Size(m)
}
func (m *SimpleMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_SimpleMessage.DiscardUnknown(m)
}

var xxx_messageInfo_SimpleMessage proto.InternalMessageInfo

func (m *SimpleMessage) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SimpleMessage) GetIsSimple() bool {
	if m != nil {
		return m.IsSimple
	}
	return false
}

func (m *SimpleMessage) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SimpleMessage) GetSampleList() []int32 {
	if m != nil {
		return m.SampleList
	}
	return nil
}

func init() {
	proto.RegisterType((*SimpleMessage)(nil), "example.simple.SimpleMessage")
}

func init() { proto.RegisterFile("simple/simple.proto", fileDescriptor_9b3c868e94d57426) }

var fileDescriptor_9b3c868e94d57426 = []byte{
	// 156 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0xce, 0xcc, 0x2d,
	0xc8, 0x49, 0xd5, 0x87, 0x50, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x7c, 0xa9, 0x15, 0x89,
	0x60, 0x2e, 0x44, 0x54, 0xa9, 0x90, 0x8b, 0x37, 0x18, 0xcc, 0xf2, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c,
	0x4f, 0x15, 0xe2, 0xe3, 0x62, 0xca, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0d, 0x62, 0xca,
	0x4c, 0x11, 0x92, 0xe6, 0xe2, 0xcc, 0x2c, 0x8e, 0x87, 0xa8, 0x96, 0x60, 0x52, 0x60, 0xd4, 0xe0,
	0x08, 0xe2, 0xc8, 0x2c, 0x86, 0xe8, 0x11, 0x12, 0xe2, 0x62, 0xc9, 0x4b, 0xcc, 0x4d, 0x95, 0x60,
	0x56, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x85, 0xe4, 0xb9, 0xb8, 0x8b, 0xc1, 0x56, 0xc4, 0xe7,
	0x64, 0x16, 0x97, 0x48, 0xb0, 0x28, 0x30, 0x6b, 0xb0, 0x06, 0x71, 0x41, 0x84, 0x7c, 0x32, 0x8b,
	0x4b, 0x9c, 0xb8, 0xa2, 0x38, 0x20, 0xc6, 0x15, 0x24, 0x25, 0xb1, 0x81, 0x5d, 0x65, 0x0c, 0x08,
	0x00, 0x00, 0xff, 0xff, 0xea, 0x70, 0x2d, 0xce, 0xac, 0x00, 0x00, 0x00,
}
