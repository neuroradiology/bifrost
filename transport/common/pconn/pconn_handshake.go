package pconn

import (
	"context"
	"net"

	"github.com/aperturerobotics/bifrost/handshake/identity"
	"github.com/aperturerobotics/bifrost/handshake/identity/s2s"
	"github.com/aperturerobotics/bifrost/peer"
	"github.com/golang/protobuf/proto"
)

// defaultMaxInflightStreamEstablish is the maximum number of raw stream
// establish messages to send before applying backpressure.
var defaultMaxInflightStreamEstablish = 5

// inflightHandshake is an on-going handshake.
type inflightHandshake struct {
	ctxCancel context.CancelFunc
	hs        identity.Handshaker
	addr      net.Addr
}

// handleCompleteHandshake handles a completed handshake.
func (u *Transport) handleCompleteHandshake(
	addr net.Addr,
	result *identity.Result,
	initiator bool,
) {
	ctx := u.ctx
	as := addr.String()
	pid, _ := peer.IDFromPublicKey(result.Peer)

	le := u.le.
		WithField("remote-id", pid.Pretty()).
		WithField("remote-addr", as)

	exd := &HandshakeExtraData{}
	if err := proto.Unmarshal(result.ExtraData, exd); err != nil {
		le.WithError(err).Warn("unable to decode extra data from handshake")
		exd.Reset()
	} else {
		le = le.WithField("remote-transport-id", exd.GetLocalTransportUuid())
	}

	le.Info("handshake complete")

	u.linksMtx.Lock()
	defer u.linksMtx.Unlock()

	// TODO; re-configure link for new secret rather than closing it.
	/*
		if l, ok := u.links[as]; ok {
			le.
				Debug("userping old session with peer")
			l.Close()
		}
	*/

	var lnk *Link
	lnk = NewLink(
		ctx,
		le,
		u.opts.GetDataShards(),
		u.opts.GetParityShards(),
		u.pc.LocalAddr(),
		addr,
		u.GetUUID(),
		exd.GetLocalTransportUuid(),
		result,
		result.Secret,
		u.pc.WriteTo,
		initiator,
		func() {
			go u.handleLinkLost(as, lnk)
		},
	)
	u.links[as] = lnk
	go u.handler.HandleLinkEstablished(lnk)
}

// pushHandshaker builds a new handshaker for the address.
// it is expected that handshakesMtx is locked before calling pushHandshaker
func (u *Transport) pushHandshaker(
	ctx context.Context,
	addr net.Addr,
	inititiator bool,
) (*inflightHandshake, error) {
	as := addr.String()
	nctx, nctxCancel := context.WithTimeout(ctx, handshakeTimeout)
	hs := &inflightHandshake{ctxCancel: nctxCancel, addr: addr}
	var err error
	// TODO: construct extra data
	ed := &HandshakeExtraData{LocalTransportUuid: u.GetUUID()}
	edDat, err := proto.Marshal(ed)
	if err != nil {
		return nil, err
	}
	hs.hs, err = s2s.NewHandshaker(
		u.privKey,
		nil,
		func(data []byte) error {
			data = append(data, byte(PacketType_PacketType_HANDSHAKE))
			_, err := u.pc.WriteTo(data, addr)
			return err
		},
		nil,
		inititiator,
		edDat,
	)
	if err != nil {
		nctxCancel()
		return nil, err
	}

	if old, ok := u.handshakes[as]; ok && old.ctxCancel != nil {
		old.ctxCancel()
	}

	u.handshakes[as] = hs
	go u.processHandshake(nctx, hs, inititiator)
	return hs, nil
}

// processHandshake processes an in-flight handshake.
func (u *Transport) processHandshake(ctx context.Context, hs *inflightHandshake, initiator bool) {
	as := hs.addr.String()
	ule := u.le.WithField("addr", as)

	defer func() {
		hs.hs.Close()
		u.handshakesMtx.Lock()
		ohs := u.handshakes[as]
		if ohs == hs {
			delete(u.handshakes, as)
		}
		u.handshakesMtx.Unlock()
	}()

	res, err := hs.hs.Execute(ctx)
	if err != nil {
		if err == context.Canceled {
			return
		}

		ule.WithError(err).Warn("error handshaking")
		return
	}

	u.handleCompleteHandshake(hs.addr, res, initiator)
}
