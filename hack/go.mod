module hack

go 1.19

replace github.com/aperturerobotics/bifrost => ../

replace (
	github.com/sirupsen/logrus => github.com/aperturerobotics/logrus v1.9.1-0.20221224130652-ff61cbb763af // aperture
	google.golang.org/protobuf => github.com/aperturerobotics/protobuf-go v1.31.1-0.20231012212426-9cf9f0f94f47 // aperture
)

require (
	github.com/golangci/golangci-lint v1.53.3
	github.com/planetscale/vtprotobuf v0.5.0
	github.com/psampaz/go-mod-outdated v0.9.0
	github.com/quic-go/quic-go v0.39.1 // indirect
	google.golang.org/protobuf v1.31.0
)

require (
	github.com/aperturerobotics/bifrost v0.17.4
	github.com/aperturerobotics/goprotowrap v0.3.0
	github.com/aperturerobotics/starpc v0.21.2
	github.com/evanw/esbuild v0.18.12
	github.com/hajimehoshi/wasmserve v0.0.0-20230504031707-cd7d23533341
	storj.io/drpc v0.0.33
)

require (
	4d63.com/gocheckcompilerdirectives v1.2.1 // indirect
	4d63.com/gochecknoglobals v0.2.1 // indirect
	filippo.io/edwards25519 v1.0.1-0.20220803165937-8c58ed0e3550 // indirect
	github.com/4meepo/tagalign v1.2.2 // indirect
	github.com/Abirdcfly/dupword v0.0.11 // indirect
	github.com/Antonboom/errname v0.1.10 // indirect
	github.com/Antonboom/nilnil v0.1.5 // indirect
	github.com/BurntSushi/toml v1.3.2 // indirect
	github.com/Djarvur/go-err113 v0.0.0-20210108212216-aea10b59be24 // indirect
	github.com/GaijinEntertainment/go-exhaustruct/v2 v2.3.0 // indirect
	github.com/Masterminds/semver v1.5.0 // indirect
	github.com/OpenPeeDeeP/depguard/v2 v2.1.0 // indirect
	github.com/alexkohler/nakedret/v2 v2.0.2 // indirect
	github.com/alexkohler/prealloc v1.0.0 // indirect
	github.com/alingse/asasalint v0.0.11 // indirect
	github.com/aperturerobotics/controllerbus v0.30.7 // indirect
	github.com/aperturerobotics/entitygraph v0.4.0 // indirect
	github.com/aperturerobotics/timestamp v0.8.2 // indirect
	github.com/aperturerobotics/ts-proto-common-types v0.2.0 // indirect
	github.com/aperturerobotics/util v1.7.6 // indirect
	github.com/ashanbrown/forbidigo v1.5.3 // indirect
	github.com/ashanbrown/makezero v1.1.1 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/bkielbasa/cyclop v1.2.1 // indirect
	github.com/blang/semver v3.5.1+incompatible // indirect
	github.com/blizzy78/varnamelen v0.8.0 // indirect
	github.com/bombsimon/wsl/v3 v3.4.0 // indirect
	github.com/breml/bidichk v0.2.4 // indirect
	github.com/breml/errchkjson v0.3.1 // indirect
	github.com/butuzov/ireturn v0.2.0 // indirect
	github.com/butuzov/mirror v1.1.0 // indirect
	github.com/cenkalti/backoff v2.2.1+incompatible // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/charithe/durationcheck v0.0.10 // indirect
	github.com/chavacava/garif v0.0.0-20230227094218-b8c73b2037b8 // indirect
	github.com/curioswitch/go-reassign v0.2.0 // indirect
	github.com/daixiang0/gci v0.10.1 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/davidlazar/go-crypto v0.0.0-20200604182044-b73af7476f6c // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.2.0 // indirect
	github.com/denis-tingaikin/go-header v0.4.3 // indirect
	github.com/esimonov/ifshort v1.0.4 // indirect
	github.com/ettle/strcase v0.1.1 // indirect
	github.com/fatih/color v1.15.0 // indirect
	github.com/fatih/structtag v1.2.0 // indirect
	github.com/firefart/nonamedreturns v1.0.4 // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/fzipp/gocyclo v0.6.0 // indirect
	github.com/go-critic/go-critic v0.8.1 // indirect
	github.com/go-task/slim-sprig v0.0.0-20230315185526-52ccab3ef572 // indirect
	github.com/go-toolsmith/astcast v1.1.0 // indirect
	github.com/go-toolsmith/astcopy v1.1.0 // indirect
	github.com/go-toolsmith/astequal v1.1.0 // indirect
	github.com/go-toolsmith/astfmt v1.1.0 // indirect
	github.com/go-toolsmith/astp v1.1.0 // indirect
	github.com/go-toolsmith/strparse v1.1.0 // indirect
	github.com/go-toolsmith/typep v1.1.0 // indirect
	github.com/go-xmlfmt/xmlfmt v1.1.2 // indirect
	github.com/gobwas/glob v0.2.3 // indirect
	github.com/gofrs/flock v0.8.1 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/golangci/check v0.0.0-20180506172741-cfe4005ccda2 // indirect
	github.com/golangci/dupl v0.0.0-20180902072040-3e9179ac440a // indirect
	github.com/golangci/go-misc v0.0.0-20220329215616-d24fe342adfe // indirect
	github.com/golangci/gofmt v0.0.0-20220901101216-f2edd75033f2 // indirect
	github.com/golangci/lint-1 v0.0.0-20191013205115-297bf364a8e0 // indirect
	github.com/golangci/maligned v0.0.0-20180506175553-b1d89398deca // indirect
	github.com/golangci/misspell v0.4.0 // indirect
	github.com/golangci/revgrep v0.0.0-20220804021717-745bb2f7c2e6 // indirect
	github.com/golangci/unconvert v0.0.0-20180507085042-28b1c447d1f4 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/google/pprof v0.0.0-20230817174616-7a8ec2ada47b // indirect
	github.com/gordonklaus/ineffassign v0.0.0-20230610083614-0e73809eb601 // indirect
	github.com/gostaticanalysis/analysisutil v0.7.1 // indirect
	github.com/gostaticanalysis/comment v1.4.2 // indirect
	github.com/gostaticanalysis/forcetypeassert v0.1.0 // indirect
	github.com/gostaticanalysis/nilerr v0.1.1 // indirect
	github.com/hashicorp/errwrap v1.0.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/go-version v1.6.0 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/hexops/gotextdiff v1.0.3 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/ipfs/go-cid v0.4.1 // indirect
	github.com/ipfs/go-log/v2 v2.5.1 // indirect
	github.com/jbenet/go-temp-err-catcher v0.1.0 // indirect
	github.com/jgautheron/goconst v1.5.1 // indirect
	github.com/jingyugao/rowserrcheck v1.1.1 // indirect
	github.com/jirfag/go-printf-func-name v0.0.0-20200119135958-7558a9eaa5af // indirect
	github.com/julz/importas v0.1.0 // indirect
	github.com/kisielk/errcheck v1.6.3 // indirect
	github.com/kisielk/gotool v1.0.0 // indirect
	github.com/kkHAIKE/contextcheck v1.1.4 // indirect
	github.com/klauspost/compress v1.17.1 // indirect
	github.com/klauspost/cpuid/v2 v2.2.5 // indirect
	github.com/kulti/thelper v0.6.3 // indirect
	github.com/kunwardeep/paralleltest v1.0.7 // indirect
	github.com/kyoh86/exportloopref v0.1.11 // indirect
	github.com/ldez/gomoddirectives v0.2.3 // indirect
	github.com/ldez/tagliatelle v0.5.0 // indirect
	github.com/leonklingele/grouper v1.1.1 // indirect
	github.com/libp2p/go-buffer-pool v0.1.1-0.20220919134021-a29bd39bcbb7 // indirect
	github.com/libp2p/go-libp2p v0.31.0 // indirect
	github.com/lufeee/execinquery v1.2.1 // indirect
	github.com/magiconair/properties v1.8.6 // indirect
	github.com/maratori/testableexamples v1.0.0 // indirect
	github.com/maratori/testpackage v1.1.1 // indirect
	github.com/matoous/godox v0.0.0-20230222163458-006bad1f9d26 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/mattn/go-runewidth v0.0.14 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.4 // indirect
	github.com/mbilski/exhaustivestruct v1.2.0 // indirect
	github.com/mgechev/revive v1.3.2 // indirect
	github.com/minio/sha256-simd v1.0.1 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/moricho/tparallel v0.3.1 // indirect
	github.com/mr-tron/base58 v1.2.0 // indirect
	github.com/multiformats/go-base32 v0.1.0 // indirect
	github.com/multiformats/go-base36 v0.2.0 // indirect
	github.com/multiformats/go-multiaddr v0.12.0 // indirect
	github.com/multiformats/go-multibase v0.2.0 // indirect
	github.com/multiformats/go-multicodec v0.9.0 // indirect
	github.com/multiformats/go-multihash v0.2.3 // indirect
	github.com/multiformats/go-multistream v0.4.1 // indirect
	github.com/multiformats/go-varint v0.0.7 // indirect
	github.com/nakabonne/nestif v0.3.1 // indirect
	github.com/nbutton23/zxcvbn-go v0.0.0-20210217022336-fa2cb2858354 // indirect
	github.com/nishanths/exhaustive v0.11.0 // indirect
	github.com/nishanths/predeclared v0.2.2 // indirect
	github.com/nunnatsa/ginkgolinter v0.12.1 // indirect
	github.com/olekukonko/tablewriter v0.0.5 // indirect
	github.com/onsi/ginkgo/v2 v2.11.0 // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	github.com/pelletier/go-toml v1.9.5 // indirect
	github.com/pelletier/go-toml/v2 v2.0.5 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/polyfloyd/go-errorlint v1.4.2 // indirect
	github.com/prometheus/client_golang v1.14.0 // indirect
	github.com/prometheus/client_model v0.4.0 // indirect
	github.com/prometheus/common v0.42.0 // indirect
	github.com/prometheus/procfs v0.9.0 // indirect
	github.com/quasilyte/go-ruleguard v0.3.19 // indirect
	github.com/quasilyte/gogrep v0.5.0 // indirect
	github.com/quasilyte/regex/syntax v0.0.0-20210819130434-b3f0c404a727 // indirect
	github.com/quasilyte/stdinfo v0.0.0-20220114132959-f7386bf02567 // indirect
	github.com/quic-go/qtls-go1-20 v0.3.4 // indirect
	github.com/rivo/uniseg v0.4.3 // indirect
	github.com/ryancurrah/gomodguard v1.3.0 // indirect
	github.com/ryanrolds/sqlclosecheck v0.4.0 // indirect
	github.com/sanposhiho/wastedassign/v2 v2.0.7 // indirect
	github.com/sashamelentyev/interfacebloat v1.1.0 // indirect
	github.com/sashamelentyev/usestdlibvars v1.23.0 // indirect
	github.com/securego/gosec/v2 v2.16.0 // indirect
	github.com/shazow/go-diff v0.0.0-20160112020656-b6b7b6733b8c // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	github.com/sivchari/containedctx v1.0.3 // indirect
	github.com/sivchari/nosnakecase v1.7.0 // indirect
	github.com/sivchari/tenv v1.7.1 // indirect
	github.com/sonatard/noctx v0.0.2 // indirect
	github.com/sourcegraph/go-diff v0.7.0 // indirect
	github.com/spaolacci/murmur3 v1.1.1-0.20190317074736-539464a789e9 // indirect
	github.com/spf13/afero v1.8.2 // indirect
	github.com/spf13/cast v1.5.0 // indirect
	github.com/spf13/cobra v1.7.0 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/spf13/viper v1.12.0 // indirect
	github.com/ssgreg/nlreturn/v2 v2.2.1 // indirect
	github.com/stbenjam/no-sprintf-host-port v0.1.1 // indirect
	github.com/stretchr/objx v0.5.0 // indirect
	github.com/stretchr/testify v1.8.4 // indirect
	github.com/subosito/gotenv v1.4.1 // indirect
	github.com/t-yuki/gocover-cobertura v0.0.0-20180217150009-aaee18c8195c // indirect
	github.com/tdakkota/asciicheck v0.2.0 // indirect
	github.com/tetafro/godot v1.4.11 // indirect
	github.com/timakin/bodyclose v0.0.0-20230421092635-574207250966 // indirect
	github.com/timonwong/loggercheck v0.9.4 // indirect
	github.com/tomarrell/wrapcheck/v2 v2.8.1 // indirect
	github.com/tommy-muehle/go-mnd/v2 v2.5.1 // indirect
	github.com/ultraware/funlen v0.0.3 // indirect
	github.com/ultraware/whitespace v0.0.5 // indirect
	github.com/uudashr/gocognit v1.0.6 // indirect
	github.com/valyala/fastjson v1.6.4 // indirect
	github.com/xen0n/gosmopolitan v1.2.1 // indirect
	github.com/yagipy/maintidx v1.0.0 // indirect
	github.com/yeya24/promlinter v0.2.0 // indirect
	github.com/ykadowak/zerologlint v0.1.2 // indirect
	github.com/zeebo/blake3 v0.2.3 // indirect
	gitlab.com/bosi/decorder v0.2.3 // indirect
	go.tmz.dev/musttag v0.7.0 // indirect
	go.uber.org/mock v0.3.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.25.0 // indirect
	golang.org/x/crypto v0.14.0 // indirect
	golang.org/x/exp v0.0.0-20231006140011-7918f672742d // indirect
	golang.org/x/exp/typeparams v0.0.0-20230224173230-c95f2b4c22f2 // indirect
	golang.org/x/mod v0.13.0 // indirect
	golang.org/x/net v0.17.0 // indirect
	golang.org/x/sync v0.4.0 // indirect
	golang.org/x/sys v0.13.0 // indirect
	golang.org/x/text v0.13.0 // indirect
	golang.org/x/tools v0.14.0 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	honnef.co/go/tools v0.4.3 // indirect
	lukechampine.com/blake3 v1.2.1 // indirect
	mvdan.cc/gofumpt v0.5.0 // indirect
	mvdan.cc/interfacer v0.0.0-20180901003855-c20040233aed // indirect
	mvdan.cc/lint v0.0.0-20170908181259-adc824a0674b // indirect
	mvdan.cc/unparam v0.0.0-20221223090309-7455f1af531d // indirect
	nhooyr.io/websocket v1.8.8-0.20221213223501-14fb98eba64e // indirect
)
