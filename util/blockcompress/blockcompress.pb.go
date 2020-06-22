// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.11.4
// source: github.com/aperturerobotics/bifrost/util/blockcompress/blockcompress.proto

package blockcompress

import (
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

// BlockCompress sets the type of compression to use.
type BlockCompress int32

const (
	// BlockCompress_NONE indicates no compression.
	BlockCompress_BlockCompress_NONE BlockCompress = 0
	// BlockCompress_SNAPPY indicates Snappy compression.
	BlockCompress_BlockCompress_SNAPPY BlockCompress = 1
	// BlockCompress_S2 indicates S2 compression.
	//
	// S2 is an extension of snappy. S2 is aimed for high throughput, which is why
	// it features concurrent compression for bigger payloads. Decoding is
	// compatible with Snappy compressed content, but content compressed with S2
	// cannot be decompressed by Snappy. This means that S2 can seamlessly replace
	// Snappy without converting compressed content. S2 is designed to have high
	// throughput on content that cannot be compressed. This is important so you
	// don't have to worry about spending CPU cycles on already compressed data.
	//
	// Reference: https://github.com/klauspost/compress/tree/master/s2
	BlockCompress_BlockCompress_S2 BlockCompress = 2
	// BlockCompress_LZ4 indicates LZ4 compression.
	BlockCompress_BlockCompress_LZ4 BlockCompress = 3
	// BlockCompress_ZSTD indicates z-standard compression.
	//
	// Zstandard is a real-time compression algorithm, providing high compression
	// ratios. It offers a very wide range of compression / speed trade-off, while
	// being backed by a very fast decoder. A high performance compression
	// algorithm is implemented.
	BlockCompress_BlockCompress_ZSTD BlockCompress = 4
)

// Enum value maps for BlockCompress.
var (
	BlockCompress_name = map[int32]string{
		0: "BlockCompress_NONE",
		1: "BlockCompress_SNAPPY",
		2: "BlockCompress_S2",
		3: "BlockCompress_LZ4",
		4: "BlockCompress_ZSTD",
	}
	BlockCompress_value = map[string]int32{
		"BlockCompress_NONE":   0,
		"BlockCompress_SNAPPY": 1,
		"BlockCompress_S2":     2,
		"BlockCompress_LZ4":    3,
		"BlockCompress_ZSTD":   4,
	}
)

func (x BlockCompress) Enum() *BlockCompress {
	p := new(BlockCompress)
	*p = x
	return p
}

func (x BlockCompress) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BlockCompress) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_aperturerobotics_bifrost_util_blockcompress_blockcompress_proto_enumTypes[0].Descriptor()
}

func (BlockCompress) Type() protoreflect.EnumType {
	return &file_github_com_aperturerobotics_bifrost_util_blockcompress_blockcompress_proto_enumTypes[0]
}

func (x BlockCompress) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BlockCompress.Descriptor instead.
func (BlockCompress) EnumDescriptor() ([]byte, []int) {
	return file_github_com_aperturerobotics_bifrost_util_blockcompress_blockcompress_proto_rawDescGZIP(), []int{0}
}

var File_github_com_aperturerobotics_bifrost_util_blockcompress_blockcompress_proto protoreflect.FileDescriptor

var file_github_com_aperturerobotics_bifrost_util_blockcompress_blockcompress_proto_rawDesc = []byte{
	0x0a, 0x4a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x70, 0x65,
	0x72, 0x74, 0x75, 0x72, 0x65, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x69, 0x63, 0x73, 0x2f, 0x62, 0x69,
	0x66, 0x72, 0x6f, 0x73, 0x74, 0x2f, 0x75, 0x74, 0x69, 0x6c, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x6f,
	0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x2a, 0x86, 0x01, 0x0a, 0x0d,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a,
	0x12, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x4e,
	0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x43, 0x6f,
	0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x53, 0x4e, 0x41, 0x50, 0x50, 0x59, 0x10, 0x01, 0x12,
	0x14, 0x0a, 0x10, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73,
	0x5f, 0x53, 0x32, 0x10, 0x02, 0x12, 0x15, 0x0a, 0x11, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x43, 0x6f,
	0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x4c, 0x5a, 0x34, 0x10, 0x03, 0x12, 0x16, 0x0a, 0x12,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x5a, 0x53,
	0x54, 0x44, 0x10, 0x04, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_aperturerobotics_bifrost_util_blockcompress_blockcompress_proto_rawDescOnce sync.Once
	file_github_com_aperturerobotics_bifrost_util_blockcompress_blockcompress_proto_rawDescData = file_github_com_aperturerobotics_bifrost_util_blockcompress_blockcompress_proto_rawDesc
)

func file_github_com_aperturerobotics_bifrost_util_blockcompress_blockcompress_proto_rawDescGZIP() []byte {
	file_github_com_aperturerobotics_bifrost_util_blockcompress_blockcompress_proto_rawDescOnce.Do(func() {
		file_github_com_aperturerobotics_bifrost_util_blockcompress_blockcompress_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_aperturerobotics_bifrost_util_blockcompress_blockcompress_proto_rawDescData)
	})
	return file_github_com_aperturerobotics_bifrost_util_blockcompress_blockcompress_proto_rawDescData
}

var file_github_com_aperturerobotics_bifrost_util_blockcompress_blockcompress_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_github_com_aperturerobotics_bifrost_util_blockcompress_blockcompress_proto_goTypes = []interface{}{
	(BlockCompress)(0), // 0: blockcompress.BlockCompress
}
var file_github_com_aperturerobotics_bifrost_util_blockcompress_blockcompress_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_github_com_aperturerobotics_bifrost_util_blockcompress_blockcompress_proto_init() }
func file_github_com_aperturerobotics_bifrost_util_blockcompress_blockcompress_proto_init() {
	if File_github_com_aperturerobotics_bifrost_util_blockcompress_blockcompress_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_aperturerobotics_bifrost_util_blockcompress_blockcompress_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_aperturerobotics_bifrost_util_blockcompress_blockcompress_proto_goTypes,
		DependencyIndexes: file_github_com_aperturerobotics_bifrost_util_blockcompress_blockcompress_proto_depIdxs,
		EnumInfos:         file_github_com_aperturerobotics_bifrost_util_blockcompress_blockcompress_proto_enumTypes,
	}.Build()
	File_github_com_aperturerobotics_bifrost_util_blockcompress_blockcompress_proto = out.File
	file_github_com_aperturerobotics_bifrost_util_blockcompress_blockcompress_proto_rawDesc = nil
	file_github_com_aperturerobotics_bifrost_util_blockcompress_blockcompress_proto_goTypes = nil
	file_github_com_aperturerobotics_bifrost_util_blockcompress_blockcompress_proto_depIdxs = nil
}
