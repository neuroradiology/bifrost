// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.11.4
// source: github.com/aperturerobotics/bifrost/transport/common/dialer/dialer.proto

package dialer

import (
	backoff "github.com/aperturerobotics/bifrost/util/backoff"
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

// DialerOpts contains options relating to dialing a statically configured peer.
type DialerOpts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Address is the address of the peer, in the format expected by the transport.
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// Backoff is the dialing backoff configuration.
	// Can be empty.
	Backoff *backoff.Backoff `protobuf:"bytes,2,opt,name=backoff,proto3" json:"backoff,omitempty"`
}

func (x *DialerOpts) Reset() {
	*x = DialerOpts{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_aperturerobotics_bifrost_transport_common_dialer_dialer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DialerOpts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DialerOpts) ProtoMessage() {}

func (x *DialerOpts) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_aperturerobotics_bifrost_transport_common_dialer_dialer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DialerOpts.ProtoReflect.Descriptor instead.
func (*DialerOpts) Descriptor() ([]byte, []int) {
	return file_github_com_aperturerobotics_bifrost_transport_common_dialer_dialer_proto_rawDescGZIP(), []int{0}
}

func (x *DialerOpts) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *DialerOpts) GetBackoff() *backoff.Backoff {
	if x != nil {
		return x.Backoff
	}
	return nil
}

var File_github_com_aperturerobotics_bifrost_transport_common_dialer_dialer_proto protoreflect.FileDescriptor

var file_github_com_aperturerobotics_bifrost_transport_common_dialer_dialer_proto_rawDesc = []byte{
	0x0a, 0x48, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x70, 0x65,
	0x72, 0x74, 0x75, 0x72, 0x65, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x69, 0x63, 0x73, 0x2f, 0x62, 0x69,
	0x66, 0x72, 0x6f, 0x73, 0x74, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x64, 0x69, 0x61, 0x6c, 0x65, 0x72, 0x2f, 0x64, 0x69,
	0x61, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x64, 0x69, 0x61, 0x6c,
	0x65, 0x72, 0x1a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61,
	0x70, 0x65, 0x72, 0x74, 0x75, 0x72, 0x65, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x69, 0x63, 0x73, 0x2f,
	0x62, 0x69, 0x66, 0x72, 0x6f, 0x73, 0x74, 0x2f, 0x75, 0x74, 0x69, 0x6c, 0x2f, 0x62, 0x61, 0x63,
	0x6b, 0x6f, 0x66, 0x66, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x6f, 0x66, 0x66, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x52, 0x0a, 0x0a, 0x44, 0x69, 0x61, 0x6c, 0x65, 0x72, 0x4f, 0x70, 0x74, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x2a, 0x0a, 0x07, 0x62, 0x61,
	0x63, 0x6b, 0x6f, 0x66, 0x66, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x62, 0x61,
	0x63, 0x6b, 0x6f, 0x66, 0x66, 0x2e, 0x42, 0x61, 0x63, 0x6b, 0x6f, 0x66, 0x66, 0x52, 0x07, 0x62,
	0x61, 0x63, 0x6b, 0x6f, 0x66, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_aperturerobotics_bifrost_transport_common_dialer_dialer_proto_rawDescOnce sync.Once
	file_github_com_aperturerobotics_bifrost_transport_common_dialer_dialer_proto_rawDescData = file_github_com_aperturerobotics_bifrost_transport_common_dialer_dialer_proto_rawDesc
)

func file_github_com_aperturerobotics_bifrost_transport_common_dialer_dialer_proto_rawDescGZIP() []byte {
	file_github_com_aperturerobotics_bifrost_transport_common_dialer_dialer_proto_rawDescOnce.Do(func() {
		file_github_com_aperturerobotics_bifrost_transport_common_dialer_dialer_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_aperturerobotics_bifrost_transport_common_dialer_dialer_proto_rawDescData)
	})
	return file_github_com_aperturerobotics_bifrost_transport_common_dialer_dialer_proto_rawDescData
}

var file_github_com_aperturerobotics_bifrost_transport_common_dialer_dialer_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_github_com_aperturerobotics_bifrost_transport_common_dialer_dialer_proto_goTypes = []interface{}{
	(*DialerOpts)(nil),      // 0: dialer.DialerOpts
	(*backoff.Backoff)(nil), // 1: backoff.Backoff
}
var file_github_com_aperturerobotics_bifrost_transport_common_dialer_dialer_proto_depIdxs = []int32{
	1, // 0: dialer.DialerOpts.backoff:type_name -> backoff.Backoff
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_github_com_aperturerobotics_bifrost_transport_common_dialer_dialer_proto_init() }
func file_github_com_aperturerobotics_bifrost_transport_common_dialer_dialer_proto_init() {
	if File_github_com_aperturerobotics_bifrost_transport_common_dialer_dialer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_aperturerobotics_bifrost_transport_common_dialer_dialer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DialerOpts); i {
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
			RawDescriptor: file_github_com_aperturerobotics_bifrost_transport_common_dialer_dialer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_aperturerobotics_bifrost_transport_common_dialer_dialer_proto_goTypes,
		DependencyIndexes: file_github_com_aperturerobotics_bifrost_transport_common_dialer_dialer_proto_depIdxs,
		MessageInfos:      file_github_com_aperturerobotics_bifrost_transport_common_dialer_dialer_proto_msgTypes,
	}.Build()
	File_github_com_aperturerobotics_bifrost_transport_common_dialer_dialer_proto = out.File
	file_github_com_aperturerobotics_bifrost_transport_common_dialer_dialer_proto_rawDesc = nil
	file_github_com_aperturerobotics_bifrost_transport_common_dialer_dialer_proto_goTypes = nil
	file_github_com_aperturerobotics_bifrost_transport_common_dialer_dialer_proto_depIdxs = nil
}
