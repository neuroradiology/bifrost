// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0-devel
// 	protoc        v3.21.9
// source: github.com/aperturerobotics/bifrost/daemon/api/controller/controller.proto

package bifrost_api_controller

import (
	reflect "reflect"
	sync "sync"

	api "github.com/aperturerobotics/bifrost/daemon/api"
	api1 "github.com/aperturerobotics/controllerbus/bus/api"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Config configures the API.
type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ListenAddr is the address to listen on for connections.
	ListenAddr string `protobuf:"bytes,1,opt,name=listen_addr,json=listenAddr,proto3" json:"listen_addr,omitempty"`
	// ApiConfig are api config options.
	ApiConfig *api.Config `protobuf:"bytes,2,opt,name=api_config,json=apiConfig,proto3" json:"api_config,omitempty"`
	// DisableBusApi disables the bus api.
	DisableBusApi bool `protobuf:"varint,3,opt,name=disable_bus_api,json=disableBusApi,proto3" json:"disable_bus_api,omitempty"`
	// BusApiConfig are controller-bus bus api config options.
	// BusApiConfig are options for controller bus api.
	BusApiConfig *api1.Config `protobuf:"bytes,4,opt,name=bus_api_config,json=busApiConfig,proto3" json:"bus_api_config,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_aperturerobotics_bifrost_daemon_api_controller_controller_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_aperturerobotics_bifrost_daemon_api_controller_controller_proto_msgTypes[0]
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
	return file_github_com_aperturerobotics_bifrost_daemon_api_controller_controller_proto_rawDescGZIP(), []int{0}
}

func (x *Config) GetListenAddr() string {
	if x != nil {
		return x.ListenAddr
	}
	return ""
}

func (x *Config) GetApiConfig() *api.Config {
	if x != nil {
		return x.ApiConfig
	}
	return nil
}

func (x *Config) GetDisableBusApi() bool {
	if x != nil {
		return x.DisableBusApi
	}
	return false
}

func (x *Config) GetBusApiConfig() *api1.Config {
	if x != nil {
		return x.BusApiConfig
	}
	return nil
}

var File_github_com_aperturerobotics_bifrost_daemon_api_controller_controller_proto protoreflect.FileDescriptor

var file_github_com_aperturerobotics_bifrost_daemon_api_controller_controller_proto_rawDesc = []byte{
	0x0a, 0x4a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x70, 0x65,
	0x72, 0x74, 0x75, 0x72, 0x65, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x69, 0x63, 0x73, 0x2f, 0x62, 0x69,
	0x66, 0x72, 0x6f, 0x73, 0x74, 0x2f, 0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x62, 0x69,
	0x66, 0x72, 0x6f, 0x73, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x6c, 0x65, 0x72, 0x1a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x61, 0x70, 0x65, 0x72, 0x74, 0x75, 0x72, 0x65, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x69, 0x63,
	0x73, 0x2f, 0x62, 0x69, 0x66, 0x72, 0x6f, 0x73, 0x74, 0x2f, 0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3b,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x70, 0x65, 0x72, 0x74,
	0x75, 0x72, 0x65, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x69, 0x63, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x62, 0x75, 0x73, 0x2f, 0x62, 0x75, 0x73, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbc, 0x01, 0x0a, 0x06,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e,
	0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c, 0x69, 0x73,
	0x74, 0x65, 0x6e, 0x41, 0x64, 0x64, 0x72, 0x12, 0x32, 0x0a, 0x0a, 0x61, 0x70, 0x69, 0x5f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x62, 0x69,
	0x66, 0x72, 0x6f, 0x73, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x52, 0x09, 0x61, 0x70, 0x69, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x26, 0x0a, 0x0f, 0x64,
	0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x62, 0x75, 0x73, 0x5f, 0x61, 0x70, 0x69, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x42, 0x75, 0x73,
	0x41, 0x70, 0x69, 0x12, 0x35, 0x0a, 0x0e, 0x62, 0x75, 0x73, 0x5f, 0x61, 0x70, 0x69, 0x5f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x62, 0x75,
	0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0c, 0x62, 0x75,
	0x73, 0x41, 0x70, 0x69, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_github_com_aperturerobotics_bifrost_daemon_api_controller_controller_proto_rawDescOnce sync.Once
	file_github_com_aperturerobotics_bifrost_daemon_api_controller_controller_proto_rawDescData = file_github_com_aperturerobotics_bifrost_daemon_api_controller_controller_proto_rawDesc
)

func file_github_com_aperturerobotics_bifrost_daemon_api_controller_controller_proto_rawDescGZIP() []byte {
	file_github_com_aperturerobotics_bifrost_daemon_api_controller_controller_proto_rawDescOnce.Do(func() {
		file_github_com_aperturerobotics_bifrost_daemon_api_controller_controller_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_aperturerobotics_bifrost_daemon_api_controller_controller_proto_rawDescData)
	})
	return file_github_com_aperturerobotics_bifrost_daemon_api_controller_controller_proto_rawDescData
}

var file_github_com_aperturerobotics_bifrost_daemon_api_controller_controller_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_github_com_aperturerobotics_bifrost_daemon_api_controller_controller_proto_goTypes = []interface{}{
	(*Config)(nil),      // 0: bifrost.api.controller.Config
	(*api.Config)(nil),  // 1: bifrost.api.Config
	(*api1.Config)(nil), // 2: bus.api.Config
}
var file_github_com_aperturerobotics_bifrost_daemon_api_controller_controller_proto_depIdxs = []int32{
	1, // 0: bifrost.api.controller.Config.api_config:type_name -> bifrost.api.Config
	2, // 1: bifrost.api.controller.Config.bus_api_config:type_name -> bus.api.Config
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_github_com_aperturerobotics_bifrost_daemon_api_controller_controller_proto_init() }
func file_github_com_aperturerobotics_bifrost_daemon_api_controller_controller_proto_init() {
	if File_github_com_aperturerobotics_bifrost_daemon_api_controller_controller_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_aperturerobotics_bifrost_daemon_api_controller_controller_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_github_com_aperturerobotics_bifrost_daemon_api_controller_controller_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_aperturerobotics_bifrost_daemon_api_controller_controller_proto_goTypes,
		DependencyIndexes: file_github_com_aperturerobotics_bifrost_daemon_api_controller_controller_proto_depIdxs,
		MessageInfos:      file_github_com_aperturerobotics_bifrost_daemon_api_controller_controller_proto_msgTypes,
	}.Build()
	File_github_com_aperturerobotics_bifrost_daemon_api_controller_controller_proto = out.File
	file_github_com_aperturerobotics_bifrost_daemon_api_controller_controller_proto_rawDesc = nil
	file_github_com_aperturerobotics_bifrost_daemon_api_controller_controller_proto_goTypes = nil
	file_github_com_aperturerobotics_bifrost_daemon_api_controller_controller_proto_depIdxs = nil
}
