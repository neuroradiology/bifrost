package kcp

import (
	"context"

	"github.com/aperturerobotics/bifrost/stream"
	"github.com/xtaci/smux/v2"
)

// smuxWrapper wraps smux.
type smuxWrapper struct {
	*smux.Session
}

// OpenStream opens a stream.
func (m *smuxWrapper) OpenStream(ctx context.Context) (stream.Stream, error) {
	return m.Session.OpenStream()
}

// AcceptStream accepts a stream.
func (m *smuxWrapper) AcceptStream() (stream.Stream, error) {
	return m.Session.AcceptStream()
}

// _ is a type assertion
var _ streamMuxer = ((*smuxWrapper)(nil))
