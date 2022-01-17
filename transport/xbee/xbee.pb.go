// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/aperturerobotics/bifrost/transport/xbee/xbee.proto

package xbee

import (
	fmt "fmt"
	math "math"

	dialer "github.com/aperturerobotics/bifrost/transport/common/dialer"
	kcp "github.com/aperturerobotics/bifrost/transport/common/kcp"
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

// Config is the configuration for the udp transport.
type Config struct {
	// TransportPeerID sets the peer ID to attach the transport to.
	// If unset, attaches to any running peer with a private key.
	TransportPeerId string `protobuf:"bytes,1,opt,name=transport_peer_id,json=transportPeerId,proto3" json:"transport_peer_id,omitempty"`
	// DevicePath is the device path to open the serial stream.
	DevicePath string `protobuf:"bytes,2,opt,name=device_path,json=devicePath,proto3" json:"device_path,omitempty"`
	// DeviceBaud is the device baudrate.
	DeviceBaud int32 `protobuf:"varint,3,opt,name=device_baud,json=deviceBaud,proto3" json:"device_baud,omitempty"`
	// PacketOpts are options to set on the packet connection.
	// In lossy environments, set the data shards for error correction.
	PacketOpts *kcp.Opts `protobuf:"bytes,4,opt,name=packet_opts,json=packetOpts,proto3" json:"packet_opts,omitempty"`
	// Dialers maps peer IDs to dialers.
	Dialers              map[string]*dialer.DialerOpts `protobuf:"bytes,5,rep,name=dialers,proto3" json:"dialers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_1cb971cd2f37ccfb, []int{0}
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

func (m *Config) GetTransportPeerId() string {
	if m != nil {
		return m.TransportPeerId
	}
	return ""
}

func (m *Config) GetDevicePath() string {
	if m != nil {
		return m.DevicePath
	}
	return ""
}

func (m *Config) GetDeviceBaud() int32 {
	if m != nil {
		return m.DeviceBaud
	}
	return 0
}

func (m *Config) GetPacketOpts() *kcp.Opts {
	if m != nil {
		return m.PacketOpts
	}
	return nil
}

func (m *Config) GetDialers() map[string]*dialer.DialerOpts {
	if m != nil {
		return m.Dialers
	}
	return nil
}

func init() {
	proto.RegisterType((*Config)(nil), "xbee.Config")
	proto.RegisterMapType((map[string]*dialer.DialerOpts)(nil), "xbee.Config.DialersEntry")
}

func init() {
	proto.RegisterFile("github.com/aperturerobotics/bifrost/transport/xbee/xbee.proto", fileDescriptor_1cb971cd2f37ccfb)
}

var fileDescriptor_1cb971cd2f37ccfb = []byte{
	// 306 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x50, 0xcd, 0x4a, 0x33, 0x31,
	0x14, 0x65, 0xfa, 0xf7, 0xd1, 0xcc, 0x07, 0x6a, 0x56, 0x63, 0x37, 0x0e, 0xae, 0x86, 0x2e, 0x32,
	0xd0, 0x6e, 0x44, 0x70, 0x53, 0x15, 0x74, 0xa3, 0xa5, 0x2f, 0x30, 0xe4, 0xe7, 0xb6, 0x0d, 0x6d,
	0x27, 0x21, 0xb9, 0x53, 0xec, 0x13, 0xf9, 0x9a, 0x32, 0x49, 0x5b, 0xba, 0xd5, 0x45, 0x92, 0x43,
	0xce, 0xb9, 0x87, 0x73, 0x0f, 0x79, 0x5a, 0x69, 0x5c, 0x37, 0x82, 0x49, 0xb3, 0x2b, 0xb9, 0x05,
	0x87, 0x8d, 0x03, 0x67, 0x84, 0x41, 0x2d, 0x7d, 0x29, 0xf4, 0xd2, 0x19, 0x8f, 0x25, 0x3a, 0x5e,
	0x7b, 0x6b, 0x1c, 0x96, 0x5f, 0x02, 0x20, 0x5c, 0xcc, 0x3a, 0x83, 0x86, 0xf6, 0x5a, 0x3c, 0x9a,
	0xfd, 0xce, 0x44, 0x9a, 0xdd, 0xce, 0xd4, 0xe5, 0x46, 0xda, 0xf6, 0x44, 0xa7, 0xd1, 0xdb, 0x9f,
	0x3c, 0x94, 0xe6, 0x5b, 0x70, 0xc7, 0x27, 0x3a, 0xdd, 0x7f, 0x77, 0xc8, 0xe0, 0xd9, 0xd4, 0x4b,
	0xbd, 0xa2, 0x63, 0x72, 0x73, 0x1e, 0xa9, 0x2c, 0x80, 0xab, 0xb4, 0xca, 0x92, 0x3c, 0x29, 0x86,
	0x8b, 0xab, 0x33, 0x31, 0x07, 0x70, 0xef, 0x8a, 0xde, 0x91, 0x54, 0xc1, 0x5e, 0x4b, 0xa8, 0x2c,
	0xc7, 0x75, 0xd6, 0x09, 0x2a, 0x12, 0xbf, 0xe6, 0x1c, 0xd7, 0x17, 0x02, 0xc1, 0x1b, 0x95, 0x75,
	0xf3, 0xa4, 0xe8, 0x9f, 0x04, 0x33, 0xde, 0x28, 0x3a, 0x26, 0xa9, 0xe5, 0x72, 0x03, 0x58, 0x19,
	0x8b, 0x3e, 0xeb, 0xe5, 0x49, 0x91, 0x4e, 0x86, 0xac, 0xdd, 0xf1, 0xd3, 0xa2, 0x5f, 0x90, 0xc8,
	0xb6, 0x98, 0x4e, 0xc9, 0xbf, 0x18, 0xda, 0x67, 0xfd, 0xbc, 0x5b, 0xa4, 0x93, 0x5b, 0x16, 0x6a,
	0x8d, 0xc1, 0xd9, 0x4b, 0xe4, 0x5e, 0x6b, 0x74, 0x87, 0xc5, 0x49, 0x39, 0xfa, 0x20, 0xff, 0x2f,
	0x09, 0x7a, 0x4d, 0xba, 0x1b, 0x38, 0x1c, 0x17, 0x6a, 0x21, 0x2d, 0x48, 0x7f, 0xcf, 0xb7, 0x0d,
	0x84, 0xf8, 0xe9, 0x84, 0xb2, 0x63, 0x33, 0x71, 0x2c, 0xa4, 0x88, 0x82, 0xc7, 0xce, 0x43, 0x22,
	0x06, 0xa1, 0xb0, 0xe9, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x43, 0x5e, 0xb8, 0xac, 0x05, 0x02,
	0x00, 0x00,
}
