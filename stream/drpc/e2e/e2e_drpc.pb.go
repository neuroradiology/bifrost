// Code generated by protoc-gen-go-drpc. DO NOT EDIT.
// protoc-gen-go-drpc version: v0.0.33
// source: github.com/aperturerobotics/bifrost/stream/drpc/e2e/e2e.proto

package drpc_e2e

import (
	context "context"
	errors "errors"

	drpc1 "github.com/planetscale/vtprotobuf/codec/drpc"
	drpc "storj.io/drpc"
	drpcerr "storj.io/drpc/drpcerr"
)

type drpcEncoding_File_github_com_aperturerobotics_bifrost_stream_drpc_e2e_e2e_proto struct{}

func (drpcEncoding_File_github_com_aperturerobotics_bifrost_stream_drpc_e2e_e2e_proto) Marshal(msg drpc.Message) ([]byte, error) {
	return drpc1.Marshal(msg)
}

func (drpcEncoding_File_github_com_aperturerobotics_bifrost_stream_drpc_e2e_e2e_proto) Unmarshal(buf []byte, msg drpc.Message) error {
	return drpc1.Unmarshal(buf, msg)
}

type DRPCEndToEndClient interface {
	DRPCConn() drpc.Conn

	Mock(ctx context.Context, in *MockRequest) (*MockResponse, error)
}

type drpcEndToEndClient struct {
	cc drpc.Conn
}

func NewDRPCEndToEndClient(cc drpc.Conn) DRPCEndToEndClient {
	return &drpcEndToEndClient{cc}
}

func (c *drpcEndToEndClient) DRPCConn() drpc.Conn { return c.cc }

func (c *drpcEndToEndClient) Mock(ctx context.Context, in *MockRequest) (*MockResponse, error) {
	out := new(MockResponse)
	err := c.cc.Invoke(ctx, "/drpc.e2e.EndToEnd/Mock", drpcEncoding_File_github_com_aperturerobotics_bifrost_stream_drpc_e2e_e2e_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type DRPCEndToEndServer interface {
	Mock(context.Context, *MockRequest) (*MockResponse, error)
}

type DRPCEndToEndUnimplementedServer struct{}

func (s *DRPCEndToEndUnimplementedServer) Mock(context.Context, *MockRequest) (*MockResponse, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

type DRPCEndToEndDescription struct{}

func (DRPCEndToEndDescription) NumMethods() int { return 1 }

func (DRPCEndToEndDescription) Method(n int) (string, drpc.Encoding, drpc.Receiver, interface{}, bool) {
	switch n {
	case 0:
		return "/drpc.e2e.EndToEnd/Mock", drpcEncoding_File_github_com_aperturerobotics_bifrost_stream_drpc_e2e_e2e_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCEndToEndServer).
					Mock(
						ctx,
						in1.(*MockRequest),
					)
			}, DRPCEndToEndServer.Mock, true
	default:
		return "", nil, nil, nil, false
	}
}

func DRPCRegisterEndToEnd(mux drpc.Mux, impl DRPCEndToEndServer) error {
	return mux.Register(impl, DRPCEndToEndDescription{})
}

type DRPCEndToEnd_MockStream interface {
	drpc.Stream
	SendAndClose(*MockResponse) error
}

type drpcEndToEnd_MockStream struct {
	drpc.Stream
}

func (x *drpcEndToEnd_MockStream) SendAndClose(m *MockResponse) error {
	if err := x.MsgSend(m, drpcEncoding_File_github_com_aperturerobotics_bifrost_stream_drpc_e2e_e2e_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}
