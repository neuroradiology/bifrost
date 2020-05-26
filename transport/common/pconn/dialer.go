package pconn

import (
	"context"
	"errors"
	"net"

	"github.com/aperturerobotics/bifrost/peer"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/lucas-clemente/quic-go"
)

// dialer represents a ongoing attempt to dial an address
type dialer struct {
	// t is the transport
	t *Transport
	// rootCtx is the root context
	rootCtx context.Context
	// ctx is the dialer ctx
	ctx context.Context
	// ctxCancel cancels ctx
	ctxCancel context.CancelFunc
	// peerID is the peer id
	peerID peer.ID
	// addr is the net.Addr
	addr net.Addr
	// address is the address string
	address string
}

// newDialer constructs a new dialer.
func newDialer(
	rctx context.Context,
	t *Transport,
	peerID peer.ID,
	addr net.Addr,
	address string,
) (*dialer, error) {
	d := &dialer{
		t:       t,
		rootCtx: rctx,
		peerID:  peerID,
		addr:    addr,
		address: address,
	}
	d.ctx, d.ctxCancel = context.WithCancel(rctx)
	return d, nil
}

// execute executes the dialer
func (d *dialer) execute() (*Link, error) {
	ctx := d.ctx
	defer d.ctxCancel()
	tlsConf, keyCh := d.t.identity.ConfigForPeer(d.peerID)
	quicConfig := defaultQuicConfig()
	d.t.le.Debugf("quic dialing peer address: %s", d.address)
	sess, err := quic.DialContext(ctx, d.t.pc, d.addr, d.address, tlsConf, quicConfig)
	if err != nil {
		return nil, err
	}
	var remotePubKey crypto.PubKey
	select {
	case remotePubKey = <-keyCh:
	case <-ctx.Done():
		return nil, ctx.Err()
	}
	if remotePubKey == nil {
		return nil, errors.New("expected remote pub key to be set")
	}

	return d.t.handleSession(d.rootCtx, sess)
}
