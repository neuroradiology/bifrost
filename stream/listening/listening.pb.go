// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/aperturerobotics/bifrost/stream/listening/listening.proto

package stream_listening

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

// Config configures the listening controller.
type Config struct {
	// LocalPeerId is the peer ID to forward incoming connections with.
	// Can be empty.
	LocalPeerId string `protobuf:"bytes,1,opt,name=local_peer_id,json=localPeerId,proto3" json:"local_peer_id,omitempty"`
	// RemotePeerId is the peer ID to forward incoming connections to.
	RemotePeerId string `protobuf:"bytes,2,opt,name=remote_peer_id,json=remotePeerId,proto3" json:"remote_peer_id,omitempty"`
	// ProtocolId is the protocol ID to assign to incoming connections.
	ProtocolId string `protobuf:"bytes,3,opt,name=protocol_id,json=protocolId,proto3" json:"protocol_id,omitempty"`
	// ListenMultiaddr is the listening multiaddress.
	ListenMultiaddr string `protobuf:"bytes,4,opt,name=listen_multiaddr,json=listenMultiaddr,proto3" json:"listen_multiaddr,omitempty"`
	// TransportId sets a transport ID constraint.
	// Can be empty.
	TransportId uint64 `protobuf:"varint,5,opt,name=transport_id,json=transportId,proto3" json:"transport_id,omitempty"`
	// Reliable indicates the stream should be reliable.
	Reliable bool `protobuf:"varint,6,opt,name=reliable,proto3" json:"reliable,omitempty"`
	// Encrypted indicates the stream should be encrypted.
	Encrypted            bool     `protobuf:"varint,7,opt,name=encrypted,proto3" json:"encrypted,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_95737aaff72d6d19, []int{0}
}

func (m *Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config.Unmarshal(m, b)
}
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config.Marshal(b, m, deterministic)
}
func (m *Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config.Merge(m, src)
}
func (m *Config) XXX_Size() int {
	return xxx_messageInfo_Config.Size(m)
}
func (m *Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Config proto.InternalMessageInfo

func (m *Config) GetLocalPeerId() string {
	if m != nil {
		return m.LocalPeerId
	}
	return ""
}

func (m *Config) GetRemotePeerId() string {
	if m != nil {
		return m.RemotePeerId
	}
	return ""
}

func (m *Config) GetProtocolId() string {
	if m != nil {
		return m.ProtocolId
	}
	return ""
}

func (m *Config) GetListenMultiaddr() string {
	if m != nil {
		return m.ListenMultiaddr
	}
	return ""
}

func (m *Config) GetTransportId() uint64 {
	if m != nil {
		return m.TransportId
	}
	return 0
}

func (m *Config) GetReliable() bool {
	if m != nil {
		return m.Reliable
	}
	return false
}

func (m *Config) GetEncrypted() bool {
	if m != nil {
		return m.Encrypted
	}
	return false
}

func init() {
	proto.RegisterType((*Config)(nil), "stream.listening.Config")
}

func init() {
	proto.RegisterFile("github.com/aperturerobotics/bifrost/stream/listening/listening.proto", fileDescriptor_95737aaff72d6d19)
}

var fileDescriptor_95737aaff72d6d19 = []byte{
	// 250 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0xce, 0xc1, 0x4e, 0x83, 0x40,
	0x10, 0x80, 0xe1, 0xa0, 0x15, 0xdb, 0xa1, 0x6a, 0xb3, 0x27, 0x62, 0x4c, 0xc4, 0xc6, 0x03, 0x5e,
	0xe0, 0xe0, 0x23, 0xe8, 0x85, 0x83, 0x89, 0xe1, 0x05, 0xc8, 0xc2, 0x4e, 0x71, 0x93, 0x85, 0x21,
	0xc3, 0x70, 0xf0, 0xd9, 0xbd, 0x98, 0x2e, 0x2d, 0xbd, 0xed, 0x7e, 0xfb, 0x67, 0x67, 0xe0, 0xb3,
	0xb5, 0xf2, 0x33, 0xd5, 0x59, 0x43, 0x5d, 0xae, 0x07, 0x64, 0x99, 0x18, 0x99, 0x6a, 0x12, 0xdb,
	0x8c, 0x79, 0x6d, 0x0f, 0x4c, 0xa3, 0xe4, 0xa3, 0x30, 0xea, 0x2e, 0x77, 0x76, 0x14, 0xec, 0x6d,
	0xdf, 0x5e, 0x4e, 0xd9, 0xc0, 0x24, 0xa4, 0x76, 0x73, 0x91, 0x2d, 0xbe, 0xff, 0x0b, 0x20, 0xfc,
	0xa0, 0xfe, 0x60, 0x5b, 0xb5, 0x87, 0x3b, 0x47, 0x8d, 0x76, 0xd5, 0x80, 0xc8, 0x95, 0x35, 0x71,
	0x90, 0x04, 0xe9, 0xa6, 0x8c, 0x3c, 0x7e, 0x23, 0x72, 0x61, 0xd4, 0x2b, 0xdc, 0x33, 0x76, 0x24,
	0xb8, 0x44, 0x57, 0x3e, 0xda, 0xce, 0x7a, 0xaa, 0x9e, 0x21, 0xf2, 0xf3, 0x1a, 0x72, 0xc7, 0xe4,
	0xda, 0x27, 0x70, 0xa6, 0xc2, 0xa8, 0x37, 0xd8, 0xcd, 0x2b, 0x54, 0xdd, 0xe4, 0xc4, 0x6a, 0x63,
	0x38, 0x5e, 0xf9, 0xea, 0x61, 0xf6, 0xaf, 0x33, 0xab, 0x17, 0xd8, 0x0a, 0xeb, 0x7e, 0x1c, 0x88,
	0xe5, 0xf8, 0xd9, 0x4d, 0x12, 0xa4, 0xab, 0x32, 0x5a, 0xac, 0x30, 0xea, 0x11, 0xd6, 0x8c, 0xce,
	0xea, 0xda, 0x61, 0x1c, 0x26, 0x41, 0xba, 0x2e, 0x97, 0xbb, 0x7a, 0x82, 0x0d, 0xf6, 0x0d, 0xff,
	0x0e, 0x82, 0x26, 0xbe, 0xf5, 0x8f, 0x17, 0xa8, 0x43, 0xbf, 0xd3, 0xfb, 0x7f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x5f, 0xc6, 0xcc, 0xee, 0x5e, 0x01, 0x00, 0x00,
}
