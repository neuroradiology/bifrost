// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/aperturerobotics/bifrost/agent/controller/config.proto

package agent_controller

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

// Config is the agent controller config.
type Config struct {
	// PrivKey is the agent private key, pem format.
	// The libp2p key is in the bytes field.
	// --- BEGIN LIBP2P PRIVATE KEY ---
	PrivKey string `protobuf:"bytes,1,opt,name=priv_key,json=privKey,proto3" json:"priv_key,omitempty"`
	// NodePeerID constrains the node peer ID.
	// If empty, attaches to whatever node is running.
	NodePeerId           string   `protobuf:"bytes,2,opt,name=node_peer_id,json=nodePeerId,proto3" json:"node_peer_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_436870aa88de69a3, []int{0}
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

func (m *Config) GetPrivKey() string {
	if m != nil {
		return m.PrivKey
	}
	return ""
}

func (m *Config) GetNodePeerId() string {
	if m != nil {
		return m.NodePeerId
	}
	return ""
}

func init() {
	proto.RegisterType((*Config)(nil), "agent.controller.Config")
}

func init() {
	proto.RegisterFile("github.com/aperturerobotics/bifrost/agent/controller/config.proto", fileDescriptor_436870aa88de69a3)
}

var fileDescriptor_436870aa88de69a3 = []byte{
	// 161 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x72, 0x4c, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0x2c, 0x48, 0x2d, 0x2a, 0x29, 0x2d, 0x4a, 0x2d,
	0xca, 0x4f, 0xca, 0x2f, 0xc9, 0x4c, 0x2e, 0xd6, 0x4f, 0xca, 0x4c, 0x2b, 0xca, 0x2f, 0x2e, 0xd1,
	0x4f, 0x4c, 0x4f, 0xcd, 0x2b, 0xd1, 0x4f, 0xce, 0xcf, 0x2b, 0x29, 0xca, 0xcf, 0xc9, 0x49, 0x2d,
	0x02, 0x31, 0xd3, 0x32, 0xd3, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x04, 0xc0, 0xd2, 0x7a,
	0x08, 0x69, 0x25, 0x57, 0x2e, 0x36, 0x67, 0xb0, 0x0a, 0x21, 0x49, 0x2e, 0x8e, 0x82, 0xa2, 0xcc,
	0xb2, 0xf8, 0xec, 0xd4, 0x4a, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x76, 0x10, 0xdf, 0x3b,
	0xb5, 0x52, 0x48, 0x81, 0x8b, 0x27, 0x2f, 0x3f, 0x25, 0x35, 0xbe, 0x20, 0x35, 0xb5, 0x28, 0x3e,
	0x33, 0x45, 0x82, 0x09, 0x2c, 0xcd, 0x05, 0x12, 0x0b, 0x48, 0x4d, 0x2d, 0xf2, 0x4c, 0x49, 0x62,
	0x03, 0x9b, 0x6f, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x25, 0x35, 0x74, 0x18, 0xa4, 0x00, 0x00,
	0x00,
}
