//+build !js

package main

import (
	"sync"

	"github.com/aperturerobotics/bifrost/daemon/api"
	"github.com/urfave/cli"
	"google.golang.org/grpc"
)

var clientDialAddr string
var clientCommands []cli.Command

func init() {
	clientCommands = append(clientCommands, cli.Command{
		Name:   "local-peers",
		Usage:  "returns local peer info",
		Action: runPeerInfo,
	}, cli.Command{
		Name:   "forward",
		Usage:  "Protocol ID will be forwarded to the target multiaddress",
		Action: runForwardController,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "peer-id",
				Usage:       "peer ID to match incoming streams to",
				Destination: &forwardingConf.PeerId,
			},
			&cli.StringFlag{
				Name:        "protocol-id",
				Usage:       "protocol ID to match incoming streams to",
				Destination: &forwardingConf.ProtocolId,
			},
			&cli.StringFlag{
				Name:        "target",
				Usage:       "target multiaddr to forward streams to",
				Destination: &forwardingConf.TargetMultiaddr,
			},
		},
	})
	commands = append(
		commands,
		cli.Command{
			Name:        "client",
			Usage:       "client sub-commands",
			After:       runCloseClient,
			Subcommands: clientCommands,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "dial-addr",
					Usage:       "address to dial API on",
					Destination: &clientDialAddr,
					Value:       "localhost:5110",
				},
			},
		},
	)
}

var clientMtx sync.Mutex
var client api.BifrostDaemonServiceClient
var clientConn *grpc.ClientConn

// GetClient builds / returns the client.
func GetClient() (api.BifrostDaemonServiceClient, error) {
	clientMtx.Lock()
	defer clientMtx.Unlock()

	if client != nil {
		return client, nil
	}

	var err error
	clientConn, err = grpc.Dial(clientDialAddr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	client = api.NewBifrostDaemonServiceClient(clientConn)
	return client, nil
}

func runCloseClient(ctx *cli.Context) error {
	if clientConn != nil {
		clientConn.Close()
		clientConn = nil
	}
	return nil
}
