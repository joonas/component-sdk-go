module github.com/wasmCloud/component-sdk-go/_examples/http-client

go 1.23.0

require (
	github.com/bytecodealliance/wasm-tools-go v0.2.0
	go.wasmcloud.dev/component v0.0.0-20240910182305-2785f866ff0f
)

require (
	github.com/coreos/go-semver v0.3.1 // indirect
	github.com/urfave/cli/v3 v3.0.0-alpha9 // indirect
	github.com/xrash/smetrics v0.0.0-20201216005158-039620a65673 // indirect
	golang.org/x/mod v0.21.0 // indirect
)

// TODO(lxf): remove this line once the module is published
replace go.wasmcloud.dev/component => ../..
