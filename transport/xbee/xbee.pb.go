// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.11.4
// source: github.com/aperturerobotics/bifrost/transport/xbee/xbee.proto

package xbee

import (
	dialer "github.com/aperturerobotics/bifrost/transport/common/dialer"
	kcp "github.com/aperturerobotics/bifrost/transport/common/kcp"
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Config is the configuration for the udp transport.
type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

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
	Dialers map[string]*dialer.DialerOpts `protobuf:"bytes,5,rep,name=dialers,proto3" json:"dialers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_aperturerobotics_bifrost_transport_xbee_xbee_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_aperturerobotics_bifrost_transport_xbee_xbee_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_github_com_aperturerobotics_bifrost_transport_xbee_xbee_proto_rawDescGZIP(), []int{0}
}

func (x *Config) GetTransportPeerId() string {
	if x != nil {
		return x.TransportPeerId
	}
	return ""
}

func (x *Config) GetDevicePath() string {
	if x != nil {
		return x.DevicePath
	}
	return ""
}

func (x *Config) GetDeviceBaud() int32 {
	if x != nil {
		return x.DeviceBaud
	}
	return 0
}

func (x *Config) GetPacketOpts() *kcp.Opts {
	if x != nil {
		return x.PacketOpts
	}
	return nil
}

func (x *Config) GetDialers() map[string]*dialer.DialerOpts {
	if x != nil {
		return x.Dialers
	}
	return nil
}

var File_github_com_aperturerobotics_bifrost_transport_xbee_xbee_proto protoreflect.FileDescriptor

var file_github_com_aperturerobotics_bifrost_transport_xbee_xbee_proto_rawDesc = []byte{
	0x0a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x70, 0x65,
	0x72, 0x74, 0x75, 0x72, 0x65, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x69, 0x63, 0x73, 0x2f, 0x62, 0x69,
	0x66, 0x72, 0x6f, 0x73, 0x74, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2f,
	0x78, 0x62, 0x65, 0x65, 0x2f, 0x78, 0x62, 0x65, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x04, 0x78, 0x62, 0x65, 0x65, 0x1a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x61, 0x70, 0x65, 0x72, 0x74, 0x75, 0x72, 0x65, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x69,
	0x63, 0x73, 0x2f, 0x62, 0x69, 0x66, 0x72, 0x6f, 0x73, 0x74, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x70, 0x6f, 0x72, 0x74, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6b, 0x63, 0x70, 0x2f,
	0x6b, 0x63, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x48, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x70, 0x65, 0x72, 0x74, 0x75, 0x72, 0x65, 0x72, 0x6f,
	0x62, 0x6f, 0x74, 0x69, 0x63, 0x73, 0x2f, 0x62, 0x69, 0x66, 0x72, 0x6f, 0x73, 0x74, 0x2f, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x64, 0x69, 0x61, 0x6c, 0x65, 0x72, 0x2f, 0x64, 0x69, 0x61, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xa7, 0x02, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x2a,
	0x0a, 0x11, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x70, 0x65, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x70, 0x6f, 0x72, 0x74, 0x50, 0x65, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1f, 0x0a, 0x0b, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x62, 0x61, 0x75, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x42, 0x61, 0x75, 0x64, 0x12, 0x2a, 0x0a, 0x0b,
	0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x6f, 0x70, 0x74, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x09, 0x2e, 0x6b, 0x63, 0x70, 0x2e, 0x4f, 0x70, 0x74, 0x73, 0x52, 0x0a, 0x70, 0x61,
	0x63, 0x6b, 0x65, 0x74, 0x4f, 0x70, 0x74, 0x73, 0x12, 0x33, 0x0a, 0x07, 0x64, 0x69, 0x61, 0x6c,
	0x65, 0x72, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x78, 0x62, 0x65, 0x65,
	0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x44, 0x69, 0x61, 0x6c, 0x65, 0x72, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x64, 0x69, 0x61, 0x6c, 0x65, 0x72, 0x73, 0x1a, 0x4e, 0x0a,
	0x0c, 0x44, 0x69, 0x61, 0x6c, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x28, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x64, 0x69, 0x61, 0x6c, 0x65, 0x72, 0x2e, 0x44, 0x69, 0x61, 0x6c, 0x65, 0x72, 0x4f, 0x70,
	0x74, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_aperturerobotics_bifrost_transport_xbee_xbee_proto_rawDescOnce sync.Once
	file_github_com_aperturerobotics_bifrost_transport_xbee_xbee_proto_rawDescData = file_github_com_aperturerobotics_bifrost_transport_xbee_xbee_proto_rawDesc
)

func file_github_com_aperturerobotics_bifrost_transport_xbee_xbee_proto_rawDescGZIP() []byte {
	file_github_com_aperturerobotics_bifrost_transport_xbee_xbee_proto_rawDescOnce.Do(func() {
		file_github_com_aperturerobotics_bifrost_transport_xbee_xbee_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_aperturerobotics_bifrost_transport_xbee_xbee_proto_rawDescData)
	})
	return file_github_com_aperturerobotics_bifrost_transport_xbee_xbee_proto_rawDescData
}

var file_github_com_aperturerobotics_bifrost_transport_xbee_xbee_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_github_com_aperturerobotics_bifrost_transport_xbee_xbee_proto_goTypes = []interface{}{
	(*Config)(nil),            // 0: xbee.Config
	nil,                       // 1: xbee.Config.DialersEntry
	(*kcp.Opts)(nil),          // 2: kcp.Opts
	(*dialer.DialerOpts)(nil), // 3: dialer.DialerOpts
}
var file_github_com_aperturerobotics_bifrost_transport_xbee_xbee_proto_depIdxs = []int32{
	2, // 0: xbee.Config.packet_opts:type_name -> kcp.Opts
	1, // 1: xbee.Config.dialers:type_name -> xbee.Config.DialersEntry
	3, // 2: xbee.Config.DialersEntry.value:type_name -> dialer.DialerOpts
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_github_com_aperturerobotics_bifrost_transport_xbee_xbee_proto_init() }
func file_github_com_aperturerobotics_bifrost_transport_xbee_xbee_proto_init() {
	if File_github_com_aperturerobotics_bifrost_transport_xbee_xbee_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_aperturerobotics_bifrost_transport_xbee_xbee_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
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
			RawDescriptor: file_github_com_aperturerobotics_bifrost_transport_xbee_xbee_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_aperturerobotics_bifrost_transport_xbee_xbee_proto_goTypes,
		DependencyIndexes: file_github_com_aperturerobotics_bifrost_transport_xbee_xbee_proto_depIdxs,
		MessageInfos:      file_github_com_aperturerobotics_bifrost_transport_xbee_xbee_proto_msgTypes,
	}.Build()
	File_github_com_aperturerobotics_bifrost_transport_xbee_xbee_proto = out.File
	file_github_com_aperturerobotics_bifrost_transport_xbee_xbee_proto_rawDesc = nil
	file_github_com_aperturerobotics_bifrost_transport_xbee_xbee_proto_goTypes = nil
	file_github_com_aperturerobotics_bifrost_transport_xbee_xbee_proto_depIdxs = nil
}
