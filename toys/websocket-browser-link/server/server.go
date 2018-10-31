package main

import (
	"context"

	"github.com/aperturerobotics/bifrost/peer"
	"github.com/aperturerobotics/bifrost/toys/websocket-browser-link/common"
	wtpt "github.com/aperturerobotics/bifrost/transport/websocket"
	"github.com/aperturerobotics/controllerbus/bus"
	"github.com/aperturerobotics/controllerbus/controller/resolver"
	"github.com/aperturerobotics/controllerbus/directive"
	"github.com/libp2p/go-libp2p-crypto"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()
var le = logrus.NewEntry(log)
var localPrivKey crypto.PrivKey
var localPeerID peer.ID

func init() {
	log.SetLevel(logrus.DebugLevel)
}

func main() {
	ctx := context.Background()
	b, privKey, err := common.BuildCommonBus(ctx)
	if err != nil {
		panic(err)
	}

	_ = privKey
	_, wsRef, err := b.AddDirective(
		resolver.NewLoadControllerWithConfigSingleton(&wtpt.Config{
			ListenAddr: ":2015",
		}),
		bus.NewCallbackHandler(func(val directive.Value) {
			le.Infof("websocket transport resolved: %#v", val)
		}, nil, nil),
	)
	defer wsRef.Release()

	<-ctx.Done()
}
