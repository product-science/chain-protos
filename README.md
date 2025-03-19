# Chain Protos

This repository contains Protocol Buffer (proto) definitions for various services and APIs. It serves as a centralized place for sharing these definitions across multiple projects.

## Repository Structure

```
chain-protos/
├── network_node/            # Service definitions for network node
│   └── v1/                  # Version 1 of network node API
│       └── network_node.proto
├── gen/                     # Directory for generated code (included in the repo)
└── python/                  # Python-specific generated code
```

## Usage

There are two main approaches to using this repository in your projects:

### Option 1: As a Proto-only Dependency

With this approach, you:
1. Import this repository as a dependency
2. Generate the code in your own project

This gives you more control over code generation parameters and avoids version mismatches.

#### Go Projects

Add this repo as a dependency:

```bash
# In your Go project
go get github.com/productscience/chain-protos
```

Generate code in your project:

```bash
# In your Go project
protoc --go_out=. --go_opt=paths=source_relative \
  --go-grpc_out=. --go-grpc_opt=paths=source_relative \
  $(go list -f '{{ .Dir }}' -m github.com/productscience/chain-protos)/network_node/v1/network_node.proto
```

You can also add this to your Makefile:

```makefile
CHAIN_PROTOS_DIR=$(shell go list -f '{{ .Dir }}' -m github.com/productscience/chain-protos)

gen-protos:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		$(CHAIN_PROTOS_DIR)/network_node/v1/network_node.proto
```

#### Non-Go Projects

Clone this repository and reference the proto files:

```bash
# Clone the repository
git clone https://github.com/productscience/chain-protos.git

# Generate code for your language
protoc --proto_path=chain-protos \
  --python_out=. --grpc_python_out=. \
  chain-protos/network_node/v1/network_node.proto
```

### Option 2: Using Pre-generated Code (Recommended for Go)

This repository includes pre-generated Go code checked into version control, making it easy to use as a direct dependency:

```bash
# In your Go project
go get github.com/productscience/chain-protos
```

Then import the generated code directly:

```go
import networknodev1 "github.com/productscience/chain-protos/gen/network_node/v1"
```

This approach is simpler and ensures all consumers use the exact same generated code, preventing subtle incompatibilities between different versions of the code generator.

## Developing Proto Files

When adding or modifying proto files:

1. Place new proto files in a directory structure that reflects their package:
   ```
   <service-name>/<version>/<service-name>.proto
   ```

2. For Go, use the following format for the `go_package` option:
   ```protobuf
   option go_package = "github.com/productscience/chain-protos/<service-name>/<version>;<packagename>";
   ```
   Where `<packagename>` is a shortened version of the package path.

3. Run code generation:
   ```bash
   make gen-go
   ```

## Understanding the go_package Option

The `go_package` option has two parts separated by a semicolon:
- The import path: `github.com/productscience/chain-protos/network_node/v1`
- The package name: `networknodev1`

Example:
```protobuf
option go_package = "github.com/productscience/chain-protos/network_node/v1;networknodev1";
```

The import path tells Go where to find the package, while the package name is used in the generated code.

## Versioning

This repository follows semantic versioning. When making changes:

- For backward-compatible additions, increment the minor version
- For backward-incompatible changes, increment the major version
- Create new proto files in new version directories for major changes

## For Cosmos SDK Projects

For detailed instructions on how to use these protos in Cosmos SDK projects, see [COSMOS_USAGE.md](COSMOS_USAGE.md).

## Examples

### Go gRPC Server and Client

A complete example showing how to use this repository in a Go project can be found in the [examples/go-usage](examples/go-usage) directory. It demonstrates:

- Setting up a Go project that depends on chain-protos
- Implementing a gRPC server using the generated code
- Creating a gRPC client that communicates with the server

To run the example:

1. Generate the Go code:
   ```bash
   make gen-go
   ```

2. Run the server:
   ```bash
   cd examples/go-usage
   go run main.go
   ```

3. In another terminal, run the client:
   ```bash
   cd examples/go-usage/client
   go run main.go
   ```

This example provides a practical demonstration of how to use the proto definitions in a real project.

## Contributing

We welcome contributions! Please see our [Contributing Guide](CONTRIBUTING.md) for details on how to submit changes and the process for reviewing them. 