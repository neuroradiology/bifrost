// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0-devel
// 	protoc        v3.21.9
// source: github.com/aperturerobotics/bifrost/transport/common/kcp/kcp.proto

package kcp

import (
	reflect "reflect"
	sync "sync"

	blockcompress "github.com/aperturerobotics/bifrost/util/blockcompress"
	blockcrypt "github.com/aperturerobotics/bifrost/util/blockcrypt"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// KCPMode is the mode to set KCP to.
type KCPMode int32

const (
	// KCPMode_UNKNOWN defaults to normal mode.
	KCPMode_KCPMode_UNKNOWN KCPMode = 0
	// KCPMode_NORMAL is the normal mode.
	// NoDelay = 0
	// Interval = 40
	// Resend = 2
	// NoCongestion = 1
	KCPMode_KCPMode_NORMAL KCPMode = 1
	// KCPMode_FAST is the "fast" mode.
	// NoDelay = 0
	// Interval = 30
	// Resend = 2
	// NoCongestion = 1
	KCPMode_KCPMode_FAST KCPMode = 2
	// KCPMode_FAST2 is the "fast2" mode.
	// NoDelay = 1
	// Interval = 20
	// Resend = 2
	// NoCongestion = 1
	KCPMode_KCPMode_FAST2 KCPMode = 3
	// KCPMode_FAST3 is the "fast3" mode.
	// NoDelay = 1
	// Interval = 10
	// Resend = 2
	// NoCongestion = 1
	KCPMode_KCPMode_FAST3 KCPMode = 4
	// KCPMode_SLOW1 is the slow 1 mode.
	// NoDelay = 0
	// Interval = 100
	// Resend = 0
	// NoCongestion = 0
	KCPMode_KCPMode_SLOW1 KCPMode = 5
)

// Enum value maps for KCPMode.
var (
	KCPMode_name = map[int32]string{
		0: "KCPMode_UNKNOWN",
		1: "KCPMode_NORMAL",
		2: "KCPMode_FAST",
		3: "KCPMode_FAST2",
		4: "KCPMode_FAST3",
		5: "KCPMode_SLOW1",
	}
	KCPMode_value = map[string]int32{
		"KCPMode_UNKNOWN": 0,
		"KCPMode_NORMAL":  1,
		"KCPMode_FAST":    2,
		"KCPMode_FAST2":   3,
		"KCPMode_FAST3":   4,
		"KCPMode_SLOW1":   5,
	}
)

func (x KCPMode) Enum() *KCPMode {
	p := new(KCPMode)
	*p = x
	return p
}

func (x KCPMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (KCPMode) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_aperturerobotics_bifrost_transport_common_kcp_kcp_proto_enumTypes[0].Descriptor()
}

func (KCPMode) Type() protoreflect.EnumType {
	return &file_github_com_aperturerobotics_bifrost_transport_common_kcp_kcp_proto_enumTypes[0]
}

func (x KCPMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use KCPMode.Descriptor instead.
func (KCPMode) EnumDescriptor() ([]byte, []int) {
	return file_github_com_aperturerobotics_bifrost_transport_common_kcp_kcp_proto_rawDescGZIP(), []int{0}
}

// StreamMuxer sets the type of stream muxer to use.
type StreamMuxer int32

const (
	// StreamMuxer_UNKNOWN defaults to StreamMuxer_XTACI_SMUX
	StreamMuxer_StreamMuxer_UNKNOWN StreamMuxer = 0
	// StreamMuxer_XTACI_SMUX is the xtaci/smux muxer.
	StreamMuxer_StreamMuxer_XTACI_SMUX StreamMuxer = 1
	// StreamMuxer_YAMUX is the yamux muxer.
	StreamMuxer_StreamMuxer_YAMUX StreamMuxer = 2
)

// Enum value maps for StreamMuxer.
var (
	StreamMuxer_name = map[int32]string{
		0: "StreamMuxer_UNKNOWN",
		1: "StreamMuxer_XTACI_SMUX",
		2: "StreamMuxer_YAMUX",
	}
	StreamMuxer_value = map[string]int32{
		"StreamMuxer_UNKNOWN":    0,
		"StreamMuxer_XTACI_SMUX": 1,
		"StreamMuxer_YAMUX":      2,
	}
)

func (x StreamMuxer) Enum() *StreamMuxer {
	p := new(StreamMuxer)
	*p = x
	return p
}

func (x StreamMuxer) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StreamMuxer) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_aperturerobotics_bifrost_transport_common_kcp_kcp_proto_enumTypes[1].Descriptor()
}

func (StreamMuxer) Type() protoreflect.EnumType {
	return &file_github_com_aperturerobotics_bifrost_transport_common_kcp_kcp_proto_enumTypes[1]
}

func (x StreamMuxer) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StreamMuxer.Descriptor instead.
func (StreamMuxer) EnumDescriptor() ([]byte, []int) {
	return file_github_com_aperturerobotics_bifrost_transport_common_kcp_kcp_proto_rawDescGZIP(), []int{1}
}

// Opts are extra options for the packet conn.
type Opts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// DataShards are the number of FEC data shards to use. By adding t check
	// symbols to the data, a Reed–Solomon code can detect any combination of up
	// to t erroneous symbols, or correct up to ⌊t/2⌋ symbols. As an erasure code,
	// it can correct up to t known erasures, or it can detect and correct
	// combinations of errors and erasures. Furthermore, Reed–Solomon codes are
	// suitable as multiple-burst bit-error correcting codes, since a sequence of
	// b + 1 consecutive bit errors can affect at most two symbols of size b. The
	// choice of t is up to the designer of the code, and may be selected within
	// wide limits. Maximum is 256.
	// Recommended: 10
	// If zero, FEC is disabled.
	DataShards uint32 `protobuf:"varint,1,opt,name=data_shards,json=dataShards,proto3" json:"data_shards,omitempty"`
	// ParityShards are the number of FEC parity shards to use.
	// Recommended: 3
	ParityShards uint32 `protobuf:"varint,2,opt,name=parity_shards,json=parityShards,proto3" json:"parity_shards,omitempty"`
	// Mtu is the maximum transmission unit to use.
	// Defaults to 1350 (UDP safe packet size).
	Mtu uint32 `protobuf:"varint,3,opt,name=mtu,proto3" json:"mtu,omitempty"`
	// KcpMode is the KCP mode.
	KcpMode KCPMode `protobuf:"varint,4,opt,name=kcp_mode,json=kcpMode,proto3,enum=kcp.KCPMode" json:"kcp_mode,omitempty"`
	// BlockCrypt is the block crypto to use.
	// Defaults to AES256.
	// Uses the handshake-negotiated session key.
	BlockCrypt blockcrypt.BlockCrypt `protobuf:"varint,5,opt,name=block_crypt,json=blockCrypt,proto3,enum=blockcrypt.BlockCrypt" json:"block_crypt,omitempty"`
	// BlockCompress is the block compression to use.
	BlockCompress blockcompress.BlockCompress `protobuf:"varint,6,opt,name=block_compress,json=blockCompress,proto3,enum=blockcompress.BlockCompress" json:"block_compress,omitempty"`
	// StreamMuxer is the stream muxer to use.
	// Defaults to smux.
	StreamMuxer StreamMuxer `protobuf:"varint,7,opt,name=stream_muxer,json=streamMuxer,proto3,enum=kcp.StreamMuxer" json:"stream_muxer,omitempty"`
}

func (x *Opts) Reset() {
	*x = Opts{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_aperturerobotics_bifrost_transport_common_kcp_kcp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Opts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Opts) ProtoMessage() {}

func (x *Opts) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_aperturerobotics_bifrost_transport_common_kcp_kcp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Opts.ProtoReflect.Descriptor instead.
func (*Opts) Descriptor() ([]byte, []int) {
	return file_github_com_aperturerobotics_bifrost_transport_common_kcp_kcp_proto_rawDescGZIP(), []int{0}
}

func (x *Opts) GetDataShards() uint32 {
	if x != nil {
		return x.DataShards
	}
	return 0
}

func (x *Opts) GetParityShards() uint32 {
	if x != nil {
		return x.ParityShards
	}
	return 0
}

func (x *Opts) GetMtu() uint32 {
	if x != nil {
		return x.Mtu
	}
	return 0
}

func (x *Opts) GetKcpMode() KCPMode {
	if x != nil {
		return x.KcpMode
	}
	return KCPMode_KCPMode_UNKNOWN
}

func (x *Opts) GetBlockCrypt() blockcrypt.BlockCrypt {
	if x != nil {
		return x.BlockCrypt
	}
	return blockcrypt.BlockCrypt(0)
}

func (x *Opts) GetBlockCompress() blockcompress.BlockCompress {
	if x != nil {
		return x.BlockCompress
	}
	return blockcompress.BlockCompress(0)
}

func (x *Opts) GetStreamMuxer() StreamMuxer {
	if x != nil {
		return x.StreamMuxer
	}
	return StreamMuxer_StreamMuxer_UNKNOWN
}

var File_github_com_aperturerobotics_bifrost_transport_common_kcp_kcp_proto protoreflect.FileDescriptor

var file_github_com_aperturerobotics_bifrost_transport_common_kcp_kcp_proto_rawDesc = []byte{
	0x0a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x70, 0x65,
	0x72, 0x74, 0x75, 0x72, 0x65, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x69, 0x63, 0x73, 0x2f, 0x62, 0x69,
	0x66, 0x72, 0x6f, 0x73, 0x74, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6b, 0x63, 0x70, 0x2f, 0x6b, 0x63, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6b, 0x63, 0x70, 0x1a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x70, 0x65, 0x72, 0x74, 0x75, 0x72, 0x65, 0x72, 0x6f,
	0x62, 0x6f, 0x74, 0x69, 0x63, 0x73, 0x2f, 0x62, 0x69, 0x66, 0x72, 0x6f, 0x73, 0x74, 0x2f, 0x75,
	0x74, 0x69, 0x6c, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x72, 0x79, 0x70, 0x74, 0x2f, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x72, 0x79, 0x70, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x4a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x70, 0x65, 0x72,
	0x74, 0x75, 0x72, 0x65, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x69, 0x63, 0x73, 0x2f, 0x62, 0x69, 0x66,
	0x72, 0x6f, 0x73, 0x74, 0x2f, 0x75, 0x74, 0x69, 0x6c, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63,
	0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x6f, 0x6d,
	0x70, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xba, 0x02, 0x0a, 0x04,
	0x4f, 0x70, 0x74, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x68, 0x61,
	0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x53,
	0x68, 0x61, 0x72, 0x64, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x61, 0x72, 0x69, 0x74, 0x79, 0x5f,
	0x73, 0x68, 0x61, 0x72, 0x64, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x70, 0x61,
	0x72, 0x69, 0x74, 0x79, 0x53, 0x68, 0x61, 0x72, 0x64, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x74,
	0x75, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6d, 0x74, 0x75, 0x12, 0x27, 0x0a, 0x08,
	0x6b, 0x63, 0x70, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c,
	0x2e, 0x6b, 0x63, 0x70, 0x2e, 0x4b, 0x43, 0x50, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x07, 0x6b, 0x63,
	0x70, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x37, 0x0a, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x63, 0x72, 0x79, 0x70, 0x74, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x43, 0x72, 0x79,
	0x70, 0x74, 0x52, 0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x43, 0x72, 0x79, 0x70, 0x74, 0x12, 0x43,
	0x0a, 0x0e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x6f,
	0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x43, 0x6f, 0x6d, 0x70,
	0x72, 0x65, 0x73, 0x73, 0x52, 0x0d, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x43, 0x6f, 0x6d, 0x70, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x33, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x6d, 0x75,
	0x78, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x6b, 0x63, 0x70, 0x2e,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4d, 0x75, 0x78, 0x65, 0x72, 0x52, 0x0b, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x4d, 0x75, 0x78, 0x65, 0x72, 0x2a, 0x7d, 0x0a, 0x07, 0x4b, 0x43, 0x50, 0x4d,
	0x6f, 0x64, 0x65, 0x12, 0x13, 0x0a, 0x0f, 0x4b, 0x43, 0x50, 0x4d, 0x6f, 0x64, 0x65, 0x5f, 0x55,
	0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x4b, 0x43, 0x50, 0x4d,
	0x6f, 0x64, 0x65, 0x5f, 0x4e, 0x4f, 0x52, 0x4d, 0x41, 0x4c, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c,
	0x4b, 0x43, 0x50, 0x4d, 0x6f, 0x64, 0x65, 0x5f, 0x46, 0x41, 0x53, 0x54, 0x10, 0x02, 0x12, 0x11,
	0x0a, 0x0d, 0x4b, 0x43, 0x50, 0x4d, 0x6f, 0x64, 0x65, 0x5f, 0x46, 0x41, 0x53, 0x54, 0x32, 0x10,
	0x03, 0x12, 0x11, 0x0a, 0x0d, 0x4b, 0x43, 0x50, 0x4d, 0x6f, 0x64, 0x65, 0x5f, 0x46, 0x41, 0x53,
	0x54, 0x33, 0x10, 0x04, 0x12, 0x11, 0x0a, 0x0d, 0x4b, 0x43, 0x50, 0x4d, 0x6f, 0x64, 0x65, 0x5f,
	0x53, 0x4c, 0x4f, 0x57, 0x31, 0x10, 0x05, 0x2a, 0x59, 0x0a, 0x0b, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x4d, 0x75, 0x78, 0x65, 0x72, 0x12, 0x17, 0x0a, 0x13, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x4d, 0x75, 0x78, 0x65, 0x72, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12,
	0x1a, 0x0a, 0x16, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4d, 0x75, 0x78, 0x65, 0x72, 0x5f, 0x58,
	0x54, 0x41, 0x43, 0x49, 0x5f, 0x53, 0x4d, 0x55, 0x58, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x4d, 0x75, 0x78, 0x65, 0x72, 0x5f, 0x59, 0x41, 0x4d, 0x55, 0x58,
	0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_aperturerobotics_bifrost_transport_common_kcp_kcp_proto_rawDescOnce sync.Once
	file_github_com_aperturerobotics_bifrost_transport_common_kcp_kcp_proto_rawDescData = file_github_com_aperturerobotics_bifrost_transport_common_kcp_kcp_proto_rawDesc
)

func file_github_com_aperturerobotics_bifrost_transport_common_kcp_kcp_proto_rawDescGZIP() []byte {
	file_github_com_aperturerobotics_bifrost_transport_common_kcp_kcp_proto_rawDescOnce.Do(func() {
		file_github_com_aperturerobotics_bifrost_transport_common_kcp_kcp_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_aperturerobotics_bifrost_transport_common_kcp_kcp_proto_rawDescData)
	})
	return file_github_com_aperturerobotics_bifrost_transport_common_kcp_kcp_proto_rawDescData
}

var file_github_com_aperturerobotics_bifrost_transport_common_kcp_kcp_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_github_com_aperturerobotics_bifrost_transport_common_kcp_kcp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_github_com_aperturerobotics_bifrost_transport_common_kcp_kcp_proto_goTypes = []interface{}{
	(KCPMode)(0),                     // 0: kcp.KCPMode
	(StreamMuxer)(0),                 // 1: kcp.StreamMuxer
	(*Opts)(nil),                     // 2: kcp.Opts
	(blockcrypt.BlockCrypt)(0),       // 3: blockcrypt.BlockCrypt
	(blockcompress.BlockCompress)(0), // 4: blockcompress.BlockCompress
}
var file_github_com_aperturerobotics_bifrost_transport_common_kcp_kcp_proto_depIdxs = []int32{
	0, // 0: kcp.Opts.kcp_mode:type_name -> kcp.KCPMode
	3, // 1: kcp.Opts.block_crypt:type_name -> blockcrypt.BlockCrypt
	4, // 2: kcp.Opts.block_compress:type_name -> blockcompress.BlockCompress
	1, // 3: kcp.Opts.stream_muxer:type_name -> kcp.StreamMuxer
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_github_com_aperturerobotics_bifrost_transport_common_kcp_kcp_proto_init() }
func file_github_com_aperturerobotics_bifrost_transport_common_kcp_kcp_proto_init() {
	if File_github_com_aperturerobotics_bifrost_transport_common_kcp_kcp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_aperturerobotics_bifrost_transport_common_kcp_kcp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Opts); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_aperturerobotics_bifrost_transport_common_kcp_kcp_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_aperturerobotics_bifrost_transport_common_kcp_kcp_proto_goTypes,
		DependencyIndexes: file_github_com_aperturerobotics_bifrost_transport_common_kcp_kcp_proto_depIdxs,
		EnumInfos:         file_github_com_aperturerobotics_bifrost_transport_common_kcp_kcp_proto_enumTypes,
		MessageInfos:      file_github_com_aperturerobotics_bifrost_transport_common_kcp_kcp_proto_msgTypes,
	}.Build()
	File_github_com_aperturerobotics_bifrost_transport_common_kcp_kcp_proto = out.File
	file_github_com_aperturerobotics_bifrost_transport_common_kcp_kcp_proto_rawDesc = nil
	file_github_com_aperturerobotics_bifrost_transport_common_kcp_kcp_proto_goTypes = nil
	file_github_com_aperturerobotics_bifrost_transport_common_kcp_kcp_proto_depIdxs = nil
}
