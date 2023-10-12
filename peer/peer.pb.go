// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0-devel
// 	protoc        v3.21.9
// source: github.com/aperturerobotics/bifrost/peer/peer.proto

package peer

import (
	reflect "reflect"
	sync "sync"

	hash "github.com/aperturerobotics/bifrost/hash"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Signature contains a signature by a peer.
type Signature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// PubKey is the public key of the peer.
	// May be empty if the public key is to be inferred from context.
	PubKey []byte `protobuf:"bytes,1,opt,name=pub_key,json=pubKey,proto3" json:"pub_key,omitempty"`
	// HashType is the hash type used to hash the data.
	// The signature is then of the hash bytes (usually 32).
	HashType hash.HashType `protobuf:"varint,2,opt,name=hash_type,json=hashType,proto3,enum=hash.HashType" json:"hash_type,omitempty"`
	// SigData contains the signature data.
	// The format is defined by the key type.
	SigData []byte `protobuf:"bytes,3,opt,name=sig_data,json=sigData,proto3" json:"sig_data,omitempty"`
}

func (x *Signature) Reset() {
	*x = Signature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_aperturerobotics_bifrost_peer_peer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Signature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Signature) ProtoMessage() {}

func (x *Signature) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_aperturerobotics_bifrost_peer_peer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Signature.ProtoReflect.Descriptor instead.
func (*Signature) Descriptor() ([]byte, []int) {
	return file_github_com_aperturerobotics_bifrost_peer_peer_proto_rawDescGZIP(), []int{0}
}

func (x *Signature) GetPubKey() []byte {
	if x != nil {
		return x.PubKey
	}
	return nil
}

func (x *Signature) GetHashType() hash.HashType {
	if x != nil {
		return x.HashType
	}
	return hash.HashType(0)
}

func (x *Signature) GetSigData() []byte {
	if x != nil {
		return x.SigData
	}
	return nil
}

// SignedMsg is a message from a peer with a signature.
type SignedMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// FromPeerId is the peer identifier of the sender.
	FromPeerId string `protobuf:"bytes,1,opt,name=from_peer_id,json=fromPeerId,proto3" json:"from_peer_id,omitempty"`
	// Signature is the sender signature.
	// Should not contain PubKey, which is inferred from peer id.
	Signature *Signature `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	// Data is the signed data.
	Data []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *SignedMsg) Reset() {
	*x = SignedMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_aperturerobotics_bifrost_peer_peer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignedMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignedMsg) ProtoMessage() {}

func (x *SignedMsg) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_aperturerobotics_bifrost_peer_peer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignedMsg.ProtoReflect.Descriptor instead.
func (*SignedMsg) Descriptor() ([]byte, []int) {
	return file_github_com_aperturerobotics_bifrost_peer_peer_proto_rawDescGZIP(), []int{1}
}

func (x *SignedMsg) GetFromPeerId() string {
	if x != nil {
		return x.FromPeerId
	}
	return ""
}

func (x *SignedMsg) GetSignature() *Signature {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *SignedMsg) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_github_com_aperturerobotics_bifrost_peer_peer_proto protoreflect.FileDescriptor

var file_github_com_aperturerobotics_bifrost_peer_peer_proto_rawDesc = []byte{
	0x0a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x70, 0x65,
	0x72, 0x74, 0x75, 0x72, 0x65, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x69, 0x63, 0x73, 0x2f, 0x62, 0x69,
	0x66, 0x72, 0x6f, 0x73, 0x74, 0x2f, 0x70, 0x65, 0x65, 0x72, 0x2f, 0x70, 0x65, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x70, 0x65, 0x65, 0x72, 0x1a, 0x33, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x70, 0x65, 0x72, 0x74, 0x75, 0x72, 0x65,
	0x72, 0x6f, 0x62, 0x6f, 0x74, 0x69, 0x63, 0x73, 0x2f, 0x62, 0x69, 0x66, 0x72, 0x6f, 0x73, 0x74,
	0x2f, 0x68, 0x61, 0x73, 0x68, 0x2f, 0x68, 0x61, 0x73, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x6c, 0x0a, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x17, 0x0a,
	0x07, 0x70, 0x75, 0x62, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06,
	0x70, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x12, 0x2b, 0x0a, 0x09, 0x68, 0x61, 0x73, 0x68, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x68, 0x61, 0x73, 0x68,
	0x2e, 0x48, 0x61, 0x73, 0x68, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x68, 0x61, 0x73, 0x68, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x69, 0x67, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x73, 0x69, 0x67, 0x44, 0x61, 0x74, 0x61, 0x22, 0x70,
	0x0a, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x4d, 0x73, 0x67, 0x12, 0x20, 0x0a, 0x0c, 0x66,
	0x72, 0x6f, 0x6d, 0x5f, 0x70, 0x65, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x66, 0x72, 0x6f, 0x6d, 0x50, 0x65, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2d, 0x0a,
	0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x70, 0x65, 0x65, 0x72, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_aperturerobotics_bifrost_peer_peer_proto_rawDescOnce sync.Once
	file_github_com_aperturerobotics_bifrost_peer_peer_proto_rawDescData = file_github_com_aperturerobotics_bifrost_peer_peer_proto_rawDesc
)

func file_github_com_aperturerobotics_bifrost_peer_peer_proto_rawDescGZIP() []byte {
	file_github_com_aperturerobotics_bifrost_peer_peer_proto_rawDescOnce.Do(func() {
		file_github_com_aperturerobotics_bifrost_peer_peer_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_aperturerobotics_bifrost_peer_peer_proto_rawDescData)
	})
	return file_github_com_aperturerobotics_bifrost_peer_peer_proto_rawDescData
}

var file_github_com_aperturerobotics_bifrost_peer_peer_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_github_com_aperturerobotics_bifrost_peer_peer_proto_goTypes = []interface{}{
	(*Signature)(nil),  // 0: peer.Signature
	(*SignedMsg)(nil),  // 1: peer.SignedMsg
	(hash.HashType)(0), // 2: hash.HashType
}
var file_github_com_aperturerobotics_bifrost_peer_peer_proto_depIdxs = []int32{
	2, // 0: peer.Signature.hash_type:type_name -> hash.HashType
	0, // 1: peer.SignedMsg.signature:type_name -> peer.Signature
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_github_com_aperturerobotics_bifrost_peer_peer_proto_init() }
func file_github_com_aperturerobotics_bifrost_peer_peer_proto_init() {
	if File_github_com_aperturerobotics_bifrost_peer_peer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_aperturerobotics_bifrost_peer_peer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Signature); i {
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
		file_github_com_aperturerobotics_bifrost_peer_peer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignedMsg); i {
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
			RawDescriptor: file_github_com_aperturerobotics_bifrost_peer_peer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_aperturerobotics_bifrost_peer_peer_proto_goTypes,
		DependencyIndexes: file_github_com_aperturerobotics_bifrost_peer_peer_proto_depIdxs,
		MessageInfos:      file_github_com_aperturerobotics_bifrost_peer_peer_proto_msgTypes,
	}.Build()
	File_github_com_aperturerobotics_bifrost_peer_peer_proto = out.File
	file_github_com_aperturerobotics_bifrost_peer_peer_proto_rawDesc = nil
	file_github_com_aperturerobotics_bifrost_peer_peer_proto_goTypes = nil
	file_github_com_aperturerobotics_bifrost_peer_peer_proto_depIdxs = nil
}
