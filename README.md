# thruput-protocol-go

Go bindings for the [Thruput protocol](https://github.com/thruputchain/thruput-protocol):
its XDR data types, constants, and RPC program clients and servers, generated
with [goxdr](https://github.com/xdrpp/goxdr). The Rust bindings live in
[thruput-protocol-rs](https://github.com/thruputchain/thruput-protocol-rs).

## Use

Requires Go 1.24 or later.

```sh
go get github.com/thruputchain/thruput-protocol-go
```

The package is named `proto`:

```go
import proto "github.com/thruputchain/thruput-protocol-go"
```

The generated types marshal through goxdr's runtime,
`github.com/xdrpp/goxdr/xdr`; see the
[goxdr manual](https://xdrpp.github.io/goxdr/pkg/github.com/xdrpp/goxdr/cmd/goxdr/goxdr.1.html).
`proto.ProtocolGitCommit` names the thruput-protocol commit the bindings were
generated from, and `proto.HogProtocol()` returns the token peers exchange to
check that they share a protocol revision.

## Layout

- `generated.go`: goxdr output. Do not edit it by hand.
- `pkg.go`: the `go:generate` line that produces `generated.go`.
- `hogp.go`: the protocol commit constants.
- `protocol/`: the [thruput-protocol](https://github.com/thruputchain/thruput-protocol)
  submodule, the `.x` sources. `build.sh` lists the files it leaves out of the
  Go bindings.
- `build.sh`: writes `pkg.go` and `hogp.go`, then regenerates `generated.go`.

## Regenerate

`generated.go` is committed, so using the package needs neither the submodule
nor goxdr. To regenerate it, check out the submodule and install the goxdr
version `go.mod` pins; any other version reorders `generated.go` wholesale.

```sh
git clone --recurse-submodules https://github.com/thruputchain/thruput-protocol-go.git
cd thruput-protocol-go
go install "github.com/xdrpp/goxdr/cmd/goxdr@$(go list -m -f '{{.Version}}' github.com/xdrpp/goxdr)"
export PATH="$(go env GOPATH)/bin:$PATH"
./build.sh          # regenerate from protocol/
./build.sh --check  # regenerate, and fail if the result differs from what is committed
```

Protocol changes belong in thruput-protocol. The `Auto-update protocol`
workflow moves the submodule to thruput-protocol `main`, regenerates, and opens
the bump pull request; CI runs `./build.sh --check`, so a hand edit to the
generated files fails it.

## Test

```sh
GOWORK=off go test ./...
go vet -composites=false ./...
```

`-composites=false` is needed because goxdr's output uses unkeyed composite
literals throughout.

## License

Apache License 2.0; see [LICENSE](LICENSE).
