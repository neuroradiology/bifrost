// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/aperturerobotics/bifrost/transport/xbee/xbee.proto

package xbee

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import dialer "github.com/aperturerobotics/bifrost/transport/common/dialer"
import pconn "github.com/aperturerobotics/bifrost/transport/common/pconn"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Config is the configuration for the udp transport.
type Config struct {
	// NodePeerID constrains the node peer ID.
	// If empty, attaches to whatever node is running.
	NodePeerId string `protobuf:"bytes,1,opt,name=node_peer_id,json=nodePeerId" json:"node_peer_id,omitempty"`
	// DevicePath is the device path to open the serial stream.
	DevicePath string `protobuf:"bytes,2,opt,name=device_path,json=devicePath" json:"device_path,omitempty"`
	// DeviceBaud is the device baudrate.
	DeviceBaud int32 `protobuf:"varint,3,opt,name=device_baud,json=deviceBaud" json:"device_baud,omitempty"`
	// PacketOpts are options to set on the packet connection.
	// In lossy environments, set the data shards for error correction.
	PacketOpts *pconn.Opts `protobuf:"bytes,4,opt,name=packet_opts,json=packetOpts" json:"packet_opts,omitempty"`
	// Dialers maps peer IDs to dialers.
	Dialers              map[string]*dialer.DialerOpts `protobuf:"bytes,5,rep,name=dialers" json:"dialers,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_xbee_eb3967fd6108773d, []int{0}
}
func (m *Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config.Unmarshal(m, b)
}
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config.Marshal(b, m, deterministic)
}
func (dst *Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config.Merge(dst, src)
}
func (m *Config) XXX_Size() int {
	return xxx_messageInfo_Config.Size(m)
}
func (m *Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Config proto.InternalMessageInfo

func (m *Config) GetNodePeerId() string {
	if m != nil {
		return m.NodePeerId
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

func (m *Config) GetPacketOpts() *pconn.Opts {
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
	proto.RegisterFile("github.com/aperturerobotics/bifrost/transport/xbee/xbee.proto", fileDescriptor_xbee_eb3967fd6108773d)
}

var fileDescriptor_xbee_eb3967fd6108773d = []byte{
	// 303 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x50, 0xcb, 0x6a, 0xc3, 0x30,
	0x10, 0xc4, 0x79, 0x95, 0xca, 0x39, 0x14, 0x9d, 0xdc, 0x5c, 0x6a, 0x7a, 0xf2, 0xa1, 0xd8, 0x90,
	0x5c, 0x4a, 0xa1, 0x97, 0xbe, 0x68, 0x2f, 0x6d, 0xc8, 0x0f, 0x18, 0x59, 0xda, 0x24, 0x22, 0xb1,
	0x56, 0x48, 0xeb, 0xd0, 0x7c, 0x4d, 0x7f, 0xb5, 0xd8, 0x72, 0x20, 0xd7, 0xf6, 0xb2, 0x5a, 0x76,
	0x46, 0xc3, 0xcc, 0xb0, 0xc7, 0x8d, 0xa6, 0x6d, 0x53, 0xe5, 0x12, 0xeb, 0x42, 0x58, 0x70, 0xd4,
	0x38, 0x70, 0x58, 0x21, 0x69, 0xe9, 0x8b, 0x4a, 0xaf, 0x1d, 0x7a, 0x2a, 0xc8, 0x09, 0xe3, 0x2d,
	0x3a, 0x2a, 0xbe, 0x2b, 0x80, 0x6e, 0xe4, 0xd6, 0x21, 0x21, 0x1f, 0xb5, 0xfb, 0xec, 0xed, 0x6f,
	0x22, 0x12, 0xeb, 0x1a, 0x4d, 0x61, 0x25, 0x9a, 0x7e, 0x06, 0xb5, 0xd9, 0xfb, 0xbf, 0x74, 0x94,
	0x16, 0x7b, 0x70, 0xfd, 0x13, 0x94, 0x6e, 0x7f, 0x06, 0x6c, 0xf2, 0x8c, 0x66, 0xad, 0x37, 0x3c,
	0x65, 0x53, 0x83, 0x0a, 0x4a, 0x0b, 0xe0, 0x4a, 0xad, 0x92, 0x28, 0x8d, 0xb2, 0xcb, 0x15, 0x6b,
	0x6f, 0x4b, 0x00, 0xf7, 0xa1, 0xf8, 0x0d, 0x8b, 0x15, 0x1c, 0xb4, 0x84, 0xd2, 0x0a, 0xda, 0x26,
	0x83, 0x40, 0x08, 0xa7, 0xa5, 0xa0, 0xed, 0x19, 0xa1, 0x12, 0x8d, 0x4a, 0x86, 0x69, 0x94, 0x8d,
	0x4f, 0x84, 0x27, 0xd1, 0x28, 0x7e, 0xc7, 0x62, 0x2b, 0xe4, 0x0e, 0xa8, 0x44, 0x4b, 0x3e, 0x19,
	0xa5, 0x51, 0x16, 0xcf, 0xe3, 0x3c, 0x64, 0xfb, 0xb2, 0xe4, 0x57, 0x2c, 0xe0, 0xed, 0xce, 0x17,
	0xec, 0x22, 0x98, 0xf5, 0xc9, 0x38, 0x1d, 0x66, 0xf1, 0xfc, 0x3a, 0xef, 0x2a, 0x0d, 0x86, 0xf3,
	0x97, 0x80, 0xbd, 0x1a, 0x72, 0xc7, 0xd5, 0x89, 0x39, 0xfb, 0x64, 0xd3, 0x73, 0x80, 0x5f, 0xb1,
	0xe1, 0x0e, 0x8e, 0x7d, 0x9a, 0x76, 0xe5, 0x19, 0x1b, 0x1f, 0xc4, 0xbe, 0x81, 0x2e, 0x40, 0x3c,
	0xe7, 0x79, 0xdf, 0x48, 0xf8, 0xd6, 0xb9, 0x08, 0x84, 0x87, 0xc1, 0x7d, 0x54, 0x4d, 0xba, 0xa2,
	0x16, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x28, 0xec, 0xe7, 0xbf, 0x01, 0x02, 0x00, 0x00,
}
