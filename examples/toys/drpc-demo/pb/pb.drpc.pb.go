// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/aperturerobotics/bifrost/examples/toys/drpc-demo/pb/pb.drpc.proto

package pb

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// DoSomethingReq is a do something request.
type DoSomethingReq struct {
	// Msg is the message to echo.
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DoSomethingReq) Reset()         { *m = DoSomethingReq{} }
func (m *DoSomethingReq) String() string { return proto.CompactTextString(m) }
func (*DoSomethingReq) ProtoMessage()    {}
func (*DoSomethingReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_e16b7f8b4640523a, []int{0}
}

func (m *DoSomethingReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DoSomethingReq.Unmarshal(m, b)
}
func (m *DoSomethingReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DoSomethingReq.Marshal(b, m, deterministic)
}
func (m *DoSomethingReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DoSomethingReq.Merge(m, src)
}
func (m *DoSomethingReq) XXX_Size() int {
	return xxx_messageInfo_DoSomethingReq.Size(m)
}
func (m *DoSomethingReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DoSomethingReq.DiscardUnknown(m)
}

var xxx_messageInfo_DoSomethingReq proto.InternalMessageInfo

func (m *DoSomethingReq) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

// DoSomethingResp is a do something response.
type DoSomethingResp struct {
	// Msg is the echoed message.
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DoSomethingResp) Reset()         { *m = DoSomethingResp{} }
func (m *DoSomethingResp) String() string { return proto.CompactTextString(m) }
func (*DoSomethingResp) ProtoMessage()    {}
func (*DoSomethingResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_e16b7f8b4640523a, []int{1}
}

func (m *DoSomethingResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DoSomethingResp.Unmarshal(m, b)
}
func (m *DoSomethingResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DoSomethingResp.Marshal(b, m, deterministic)
}
func (m *DoSomethingResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DoSomethingResp.Merge(m, src)
}
func (m *DoSomethingResp) XXX_Size() int {
	return xxx_messageInfo_DoSomethingResp.Size(m)
}
func (m *DoSomethingResp) XXX_DiscardUnknown() {
	xxx_messageInfo_DoSomethingResp.DiscardUnknown(m)
}

var xxx_messageInfo_DoSomethingResp proto.InternalMessageInfo

func (m *DoSomethingResp) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*DoSomethingReq)(nil), "pb.DoSomethingReq")
	proto.RegisterType((*DoSomethingResp)(nil), "pb.DoSomethingResp")
}

func init() {
	proto.RegisterFile("github.com/aperturerobotics/bifrost/examples/toys/drpc-demo/pb/pb.drpc.proto", fileDescriptor_e16b7f8b4640523a)
}

var fileDescriptor_e16b7f8b4640523a = []byte{
	// 188 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x8e, 0xc1, 0x8a, 0xc2, 0x30,
	0x10, 0x86, 0xb7, 0xbb, 0xb0, 0xb0, 0x59, 0xe8, 0x2e, 0xd9, 0xcb, 0xe2, 0x49, 0xe2, 0xc5, 0x8b,
	0x09, 0xe8, 0xc5, 0x07, 0xd0, 0x8b, 0x78, 0x6a, 0x9f, 0xa0, 0x93, 0x8e, 0x6d, 0xc0, 0x38, 0x63,
	0x92, 0x8a, 0xbe, 0xbd, 0x54, 0x2f, 0x96, 0xde, 0x66, 0x7e, 0xbe, 0x9f, 0xef, 0x17, 0xfb, 0xc6,
	0xa5, 0xb6, 0x03, 0x6d, 0xc9, 0x9b, 0x8a, 0x31, 0xa4, 0x2e, 0x60, 0x20, 0xa0, 0xe4, 0x6c, 0x34,
	0xe0, 0x0e, 0x81, 0x62, 0x32, 0x78, 0xad, 0x3c, 0x1f, 0x31, 0x9a, 0x44, 0xb7, 0x68, 0xea, 0xc0,
	0x76, 0x51, 0xa3, 0x27, 0xc3, 0x60, 0x18, 0x74, 0xff, 0x6b, 0x0e, 0x94, 0x48, 0xbe, 0x33, 0x28,
	0x25, 0xf2, 0x0d, 0x95, 0xe4, 0x31, 0xb5, 0xee, 0xd4, 0x14, 0x78, 0x96, 0xbf, 0xe2, 0xc3, 0xc7,
	0xe6, 0x3f, 0x9b, 0x66, 0xf3, 0xaf, 0xa2, 0x3f, 0xd5, 0x4c, 0xfc, 0x0c, 0x98, 0xc8, 0x63, 0x68,
	0xb9, 0x13, 0xf9, 0xf6, 0x29, 0x2e, 0x31, 0x5c, 0x9c, 0x45, 0xb9, 0x16, 0xdf, 0x2f, 0x35, 0x29,
	0x35, 0x83, 0x1e, 0xba, 0x26, 0x7f, 0xa3, 0x2c, 0xb2, 0x7a, 0x83, 0xcf, 0xc7, 0xbe, 0xd5, 0x3d,
	0x00, 0x00, 0xff, 0xff, 0x16, 0x9c, 0x70, 0x6b, 0xef, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ExampleServiceClient is the client API for ExampleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ExampleServiceClient interface {
	DoSomething(ctx context.Context, in *DoSomethingReq, opts ...grpc.CallOption) (*DoSomethingResp, error)
}

type exampleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewExampleServiceClient(cc grpc.ClientConnInterface) ExampleServiceClient {
	return &exampleServiceClient{cc}
}

func (c *exampleServiceClient) DoSomething(ctx context.Context, in *DoSomethingReq, opts ...grpc.CallOption) (*DoSomethingResp, error) {
	out := new(DoSomethingResp)
	err := c.cc.Invoke(ctx, "/pb.ExampleService/DoSomething", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExampleServiceServer is the server API for ExampleService service.
type ExampleServiceServer interface {
	DoSomething(context.Context, *DoSomethingReq) (*DoSomethingResp, error)
}

// UnimplementedExampleServiceServer can be embedded to have forward compatible implementations.
type UnimplementedExampleServiceServer struct {
}

func (*UnimplementedExampleServiceServer) DoSomething(ctx context.Context, req *DoSomethingReq) (*DoSomethingResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoSomething not implemented")
}

func RegisterExampleServiceServer(s *grpc.Server, srv ExampleServiceServer) {
	s.RegisterService(&_ExampleService_serviceDesc, srv)
}

func _ExampleService_DoSomething_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DoSomethingReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleServiceServer).DoSomething(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ExampleService/DoSomething",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleServiceServer).DoSomething(ctx, req.(*DoSomethingReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _ExampleService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ExampleService",
	HandlerType: (*ExampleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DoSomething",
			Handler:    _ExampleService_DoSomething_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/aperturerobotics/bifrost/examples/toys/drpc-demo/pb/pb.drpc.proto",
}
