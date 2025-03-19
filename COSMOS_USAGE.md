# Using Chain-Protos in Cosmos SDK Projects

This document explains how to properly use this protocol buffer repository in your Cosmos SDK projects.

## Adding as a Dependency

First, add this repository as a dependency in your Go project:

```bash
go get github.com/productscience/chain-protos
```

## Setup for Proto Generation

### 1. Update your buf.work.yaml

Make sure to include the chain-protos directory in your buf.work.yaml:

```yaml
version: v1
directories:
  - proto
  - vendor/github.com/productscience/chain-protos
```

### 2. Add to your Third Party Proto Dependencies

In your Cosmos SDK project, you typically have a `third_party/proto` directory. Create a symlink to the chain-protos repository:

```bash
mkdir -p third_party/proto/chain-protos
ln -s $(go list -f '{{ .Dir }}' -m github.com/productscience/chain-protos)/network_node third_party/proto/chain-protos/network_node
```

### 3. Update your Makefile

Add a rule to generate code from the chain-protos repository:

```makefile
CHAIN_PROTOS_DIR=$(shell go list -f '{{ .Dir }}' -m github.com/productscience/chain-protos)

proto-gen: proto-gen-go
  
proto-gen-go:
  @echo "Generating Go code from proto files..."
  @cd proto && \
  buf generate --template buf.gen.yaml --path ${CHAIN_PROTOS_DIR}/...
```

### 4. Import Proto Files in Your Proto Definitions

You can now import the proto definitions in your own proto files:

```protobuf
syntax = "proto3";

package myproject.mymodule.v1;

import "chain-protos/network_node/v1/network_node.proto";  // Import from chain-protos

option go_package = "github.com/myorg/myproject/x/mymodule/types";

message MyMessage {
  network_node.v1.Record record = 1;  // Use types from chain-protos
}
```

## Using Generated Code in Your Go Code

After generation, you can import and use the generated Go code:

```go
package mymodule

import (
  networknodev1 "github.com/productscience/chain-protos/gen/network_node/v1"
)

func ProcessRecord(record *networknodev1.Record) error {
  // Use the imported types from chain-protos
  return nil
}
```

## Understanding Go Package Path vs Layout

When working with Cosmos SDK projects, it's important to understand the difference between:

1. **Go import path:** The full path used in Go imports (e.g., `github.com/productscience/inference/x/inference/types`)
2. **File system layout:** Where the files are actually located (e.g., `inference-chain/x/inference/types`)

This is controlled by the module declaration in the `go.mod` file:

```
module github.com/productscience/inference
```

This determines that all imports starting with `github.com/productscience/inference` will be found in the root of your project.

## For a Typical Cosmos SDK Project

Here's a complete example for a typical Cosmos SDK project:

1. **In your go.mod**:
   ```
   module github.com/productscience/inference
   
   go 1.22.0
   
   require (
     github.com/productscience/chain-protos v0.1.0
     // other dependencies
   )
   ```

2. **In your proto/tx.proto**:
   ```protobuf
   syntax = "proto3";
   
   package inference.v1;
   
   import "chain-protos/network_node/v1/network_node.proto";
   
   option go_package = "github.com/productscience/inference/x/inference/types";
   
   message MyTx {
     network_node.v1.Record record = 1;
   }
   ```

3. **In your x/inference/types/mytx.go**:
   ```go
   package types
   
   import (
     networknodev1 "github.com/productscience/chain-protos/gen/network_node/v1"
   )
   
   func ProcessMyTx(tx *MyTx) {
     record := tx.Record
     // Use the record
   }
   ```

## Caveats and Best Practices

1. Always ensure that your import path and go_package option correctly reflect your module name.
2. When making changes to proto files, follow semantic versioning to avoid breaking changes.
3. Use buf to manage dependencies and generate code consistently across projects.
4. Consider using [buf.build](https://buf.build/) to publish your protos for easier dependency management. 