// Code generated by protoc-gen-go. DO NOT EDIT.
// source: transport/proto/handshake.proto

package proto

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

type HandshakeRequest struct {
	// protocol version number
	Version uint32 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	// string form of the sender address (e.g. "192.0.2.1:25", "[2001:db8::1]:80")
	From string `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	// string form of the recipient address
	To string `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
	// unix time
	Timestamp            int64    `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HandshakeRequest) Reset()         { *m = HandshakeRequest{} }
func (m *HandshakeRequest) String() string { return proto.CompactTextString(m) }
func (*HandshakeRequest) ProtoMessage()    {}
func (*HandshakeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d7101ffe19b05443, []int{0}
}

func (m *HandshakeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HandshakeRequest.Unmarshal(m, b)
}
func (m *HandshakeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HandshakeRequest.Marshal(b, m, deterministic)
}
func (m *HandshakeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HandshakeRequest.Merge(m, src)
}
func (m *HandshakeRequest) XXX_Size() int {
	return xxx_messageInfo_HandshakeRequest.Size(m)
}
func (m *HandshakeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HandshakeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HandshakeRequest proto.InternalMessageInfo

func (m *HandshakeRequest) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *HandshakeRequest) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *HandshakeRequest) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *HandshakeRequest) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type HandshakeResponse struct {
	// hash of the ping packet
	ReqHash              []byte   `protobuf:"bytes,1,opt,name=req_hash,json=reqHash,proto3" json:"req_hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HandshakeResponse) Reset()         { *m = HandshakeResponse{} }
func (m *HandshakeResponse) String() string { return proto.CompactTextString(m) }
func (*HandshakeResponse) ProtoMessage()    {}
func (*HandshakeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d7101ffe19b05443, []int{1}
}

func (m *HandshakeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HandshakeResponse.Unmarshal(m, b)
}
func (m *HandshakeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HandshakeResponse.Marshal(b, m, deterministic)
}
func (m *HandshakeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HandshakeResponse.Merge(m, src)
}
func (m *HandshakeResponse) XXX_Size() int {
	return xxx_messageInfo_HandshakeResponse.Size(m)
}
func (m *HandshakeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HandshakeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HandshakeResponse proto.InternalMessageInfo

func (m *HandshakeResponse) GetReqHash() []byte {
	if m != nil {
		return m.ReqHash
	}
	return nil
}

func init() {
	proto.RegisterType((*HandshakeRequest)(nil), "proto.HandshakeRequest")
	proto.RegisterType((*HandshakeResponse)(nil), "proto.HandshakeResponse")
}

func init() { proto.RegisterFile("transport/proto/handshake.proto", fileDescriptor_d7101ffe19b05443) }

var fileDescriptor_d7101ffe19b05443 = []byte{
	// 203 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x5c, 0x8f, 0x3f, 0x4b, 0x04, 0x31,
	0x10, 0xc5, 0xb9, 0x3f, 0x7a, 0xde, 0xa0, 0xa2, 0xa9, 0x22, 0x08, 0xca, 0x55, 0x82, 0xb8, 0x29,
	0xfc, 0x06, 0x56, 0x57, 0xa7, 0xb4, 0x91, 0xec, 0x3a, 0x6e, 0x82, 0x26, 0x93, 0xcd, 0x64, 0xfd,
	0xfc, 0x0e, 0x81, 0x45, 0xb1, 0x7a, 0xef, 0xf7, 0x0b, 0xe4, 0x25, 0x70, 0x57, 0x8b, 0x4b, 0x9c,
	0xa9, 0x54, 0x93, 0x0b, 0x55, 0x32, 0xde, 0xa5, 0x77, 0xf6, 0xee, 0x13, 0xbb, 0xc6, 0xea, 0xa4,
	0xc5, 0x21, 0xc1, 0xd5, 0x71, 0x39, 0xb1, 0x38, 0xcd, 0xc8, 0x55, 0x69, 0xd8, 0x7d, 0x63, 0xe1,
	0x40, 0x49, 0xaf, 0xee, 0x57, 0x0f, 0x17, 0x76, 0x41, 0xa5, 0x60, 0xfb, 0x51, 0x28, 0xea, 0xb5,
	0xe8, 0xbd, 0x6d, 0x5d, 0x5d, 0xc2, 0xba, 0x92, 0xde, 0x34, 0x23, 0x4d, 0xdd, 0xc2, 0xbe, 0x86,
	0x28, 0xf7, 0xb8, 0x98, 0xf5, 0x56, 0xf4, 0xc6, 0xfe, 0x8a, 0x43, 0x07, 0xd7, 0x7f, 0xf6, 0xe4,
	0x81, 0x89, 0x51, 0xdd, 0xc0, 0x59, 0xc1, 0xe9, 0xcd, 0x3b, 0xf6, 0x6d, 0xf1, 0xdc, 0xee, 0x84,
	0x8f, 0x82, 0x2f, 0x4f, 0xaf, 0x8f, 0x63, 0xa8, 0x7e, 0xee, 0xbb, 0x81, 0xa2, 0x19, 0x5c, 0x26,
	0x66, 0xfc, 0x42, 0x33, 0x4a, 0x86, 0x6c, 0xfe, 0xfd, 0xb2, 0x3f, 0x6d, 0xf1, 0xfc, 0x13, 0x00,
	0x00, 0xff, 0xff, 0x51, 0xe0, 0x08, 0xd0, 0xff, 0x00, 0x00, 0x00,
}