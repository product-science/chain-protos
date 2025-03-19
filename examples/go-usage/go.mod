module github.com/productscience/chain-protos/examples/go-usage

go 1.22.0

require (
	github.com/productscience/chain-protos v0.1.0
	google.golang.org/grpc v1.71.0
)

// Use local version for development
replace github.com/productscience/chain-protos => ../.. 