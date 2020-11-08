module github.com/aperturerobotics/bifrost

go 1.13

// aperture: use protobuf 1.3.x based fork for compatibility
replace (
	github.com/golang/protobuf => github.com/aperturerobotics/go-protobuf-1.3.x v0.0.0-20200706003739-05fb54d407a9 // aperture-1.3.x
	github.com/lucas-clemente/quic-go => github.com/aperturerobotics/quic-go v0.7.1-0.20201108000031-e09830c4c869 // aperture-protobuf-1.3.x
	github.com/multiformats/go-multihash => github.com/paralin/go-multihash v0.0.11-0.20200526102400-a989a5c6678b // gopherjs-compat
	google.golang.org/genproto => google.golang.org/genproto v0.0.0-20190819201941-24fa4b261c55
	google.golang.org/grpc => google.golang.org/grpc v1.30.0
)

// aperture: use aperture-2.0 branch of fork
replace (
	github.com/nats-io/nats-server/v2 => github.com/aperturerobotics/bifrost-nats-server/v2 aperture-2.0 // aperture-2.0
	github.com/nats-io/nats.go => github.com/aperturerobotics/bifrost-nats-client aperture-2.0 // aperture-2.0
)

require (
	github.com/aperturerobotics/controllerbus v0.8.1-0.20200802060256-360612dc3698
	github.com/aperturerobotics/entitygraph v0.1.3
	github.com/aperturerobotics/timestamp v0.2.3
	github.com/blang/semver v3.5.1+incompatible
	github.com/cenkalti/backoff v2.2.1+incompatible
	github.com/djherbis/buffer v1.1.0
	github.com/frankban/quicktest v1.10.2 // indirect
	github.com/golang/protobuf v1.4.2
	github.com/golang/snappy v0.0.3-0.20201103224600-674baa8c7fc3
	github.com/gopherjs/gopherjs v0.0.0-20200217142428-fce0ec30dd00
	github.com/gopherjs/websocket v0.0.0-20191103002815-9a42957e2b3a
	github.com/gorilla/websocket v1.4.3-0.20200912193213-c3dd95aea977
	github.com/hashicorp/yamux v0.0.0-20200609203250-aecfd211c9ce
	github.com/klauspost/compress v1.11.0
	github.com/libp2p/go-libp2p-core v0.7.0
	github.com/libp2p/go-libp2p-tls v0.1.3
	github.com/lucas-clemente/quic-go v0.18.1
	github.com/mr-tron/base58 v1.2.0
	github.com/multiformats/go-multiaddr v0.2.2
	github.com/multiformats/go-multiaddr-net v0.1.5
	github.com/nats-io/nats-server/v2 v2.1.9
	github.com/nats-io/nats.go v1.10.0
	github.com/nats-io/nkeys v0.2.0
	github.com/paralin/kcp-go-lite v1.0.2-0.20190927004254-2be397fe467b
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/pauleyj/gobee v0.0.0-20190212035730-6270c53072a4
	github.com/pierrec/lz4 v2.5.2+incompatible
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.7.0
	github.com/tarm/serial v0.0.0-20180830185346-98f6abe2eb07
	github.com/templexxx/xor v0.0.0-20191217153810-f85b25db303b
	github.com/tjfoc/gmsm v1.0.1
	github.com/urfave/cli v1.22.4
	github.com/xtaci/smux v1.5.14
	golang.org/x/crypto v0.0.0-20201016220609-9e8e0b390897
	google.golang.org/grpc v1.30.0
	gortc.io/stun v1.22.2
)
