// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/aperturerobotics/bifrost/stream/grpc/rpc/rpc.proto

package stream_grpc_rpc

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

// StreamState is state for the stream related calls.
type StreamState int32

const (
	// StreamState_NONE indicates nothing about the state
	StreamState_StreamState_NONE StreamState = 0
	// StreamState_ESTABLISHING indicates the stream is connecting.
	StreamState_StreamState_ESTABLISHING StreamState = 1
	// StreamState_ESTABLISHED indicates the stream is established.
	StreamState_StreamState_ESTABLISHED StreamState = 2
)

var StreamState_name = map[int32]string{
	0: "StreamState_NONE",
	1: "StreamState_ESTABLISHING",
	2: "StreamState_ESTABLISHED",
}

var StreamState_value = map[string]int32{
	"StreamState_NONE":         0,
	"StreamState_ESTABLISHING": 1,
	"StreamState_ESTABLISHED":  2,
}

func (x StreamState) String() string {
	return proto.EnumName(StreamState_name, int32(x))
}

func (StreamState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e858c2ad1e5f8ef8, []int{0}
}

// Data is a data packet.
type Data struct {
	// State indicates stream state in-band.
	// Data is packet data from the remote.
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	// State indicates the stream state.
	State                StreamState `protobuf:"varint,2,opt,name=state,proto3,enum=stream.grpc.rpc.StreamState" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Data) Reset()         { *m = Data{} }
func (m *Data) String() string { return proto.CompactTextString(m) }
func (*Data) ProtoMessage()    {}
func (*Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_e858c2ad1e5f8ef8, []int{0}
}

func (m *Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Data.Unmarshal(m, b)
}
func (m *Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Data.Marshal(b, m, deterministic)
}
func (m *Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Data.Merge(m, src)
}
func (m *Data) XXX_Size() int {
	return xxx_messageInfo_Data.Size(m)
}
func (m *Data) XXX_DiscardUnknown() {
	xxx_messageInfo_Data.DiscardUnknown(m)
}

var xxx_messageInfo_Data proto.InternalMessageInfo

func (m *Data) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Data) GetState() StreamState {
	if m != nil {
		return m.State
	}
	return StreamState_StreamState_NONE
}

func init() {
	proto.RegisterEnum("stream.grpc.rpc.StreamState", StreamState_name, StreamState_value)
	proto.RegisterType((*Data)(nil), "stream.grpc.rpc.Data")
}

func init() {
	proto.RegisterFile("github.com/aperturerobotics/bifrost/stream/grpc/rpc/rpc.proto", fileDescriptor_e858c2ad1e5f8ef8)
}

var fileDescriptor_e858c2ad1e5f8ef8 = []byte{
	// 201 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x4d, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0x2c, 0x48, 0x2d, 0x2a, 0x29, 0x2d, 0x4a, 0x2d,
	0xca, 0x4f, 0xca, 0x2f, 0xc9, 0x4c, 0x2e, 0xd6, 0x4f, 0xca, 0x4c, 0x2b, 0xca, 0x2f, 0x2e, 0xd1,
	0x2f, 0x2e, 0x29, 0x4a, 0x4d, 0xcc, 0xd5, 0x4f, 0x2f, 0x2a, 0x48, 0xd6, 0x87, 0x62, 0xbd, 0x82,
	0xa2, 0xfc, 0x92, 0x7c, 0x21, 0x7e, 0x88, 0x94, 0x1e, 0x48, 0x4a, 0xaf, 0xa8, 0x20, 0x59, 0xc9,
	0x8f, 0x8b, 0xc5, 0x25, 0xb1, 0x24, 0x51, 0x48, 0x88, 0x8b, 0x25, 0x25, 0xb1, 0x24, 0x51, 0x82,
	0x51, 0x81, 0x51, 0x83, 0x27, 0x08, 0xcc, 0x16, 0x32, 0xe2, 0x62, 0x2d, 0x2e, 0x49, 0x2c, 0x49,
	0x95, 0x60, 0x52, 0x60, 0xd4, 0xe0, 0x33, 0x92, 0xd1, 0x43, 0xd3, 0xac, 0x17, 0x0c, 0xe6, 0x07,
	0x83, 0xd4, 0x04, 0x41, 0x94, 0x6a, 0xc5, 0x71, 0x71, 0x23, 0x89, 0x0a, 0x89, 0x70, 0x09, 0x20,
	0x71, 0xe3, 0xfd, 0xfc, 0xfd, 0x5c, 0x05, 0x18, 0x84, 0x64, 0xb8, 0x24, 0x90, 0x45, 0x5d, 0x83,
	0x43, 0x1c, 0x9d, 0x7c, 0x3c, 0x83, 0x3d, 0x3c, 0xfd, 0xdc, 0x05, 0x18, 0x85, 0xa4, 0xb9, 0xc4,
	0xb1, 0xca, 0xba, 0xba, 0x08, 0x30, 0x25, 0xb1, 0x81, 0xfd, 0x61, 0x0c, 0x08, 0x00, 0x00, 0xff,
	0xff, 0x36, 0xcb, 0x30, 0x9d, 0x08, 0x01, 0x00, 0x00,
}
