# Go Usage Example

This example demonstrates how to use the chain-protos repository in a Go project.

## Overview

This is a simple gRPC server that implements the `RecordService` defined in the `network_node/v1/network_node.proto` file from the chain-protos repository.

## Running the Example

1. First, make sure the Go code has been generated in the chain-protos repository:

   ```bash
   # From the chain-protos repository root
   make gen-go
   ```

2. Then run the example:

   ```bash
   # From the examples/go-usage directory
   go run main.go
   ```

3. The server will start on port 50051. You can connect to it with a gRPC client.

## How It Works

The example:

1. Imports the generated Go code from chain-protos repository:
   ```go
   import networknodev1 "github.com/productscience/chain-protos/gen/network_node/v1"
   ```

2. Implements the `RecordService` gRPC service that was defined in the proto file:
   ```go
   type RecordServer struct {
       networknodev1.UnimplementedRecordServiceServer
       records map[string]string
   }
   ```

3. Registers the service with a gRPC server and starts it.

## Dependency Management

The `go.mod` file includes the chain-protos repository as a dependency:

```go
require (
    github.com/productscience/chain-protos v0.1.0
    google.golang.org/grpc v1.71.0
)
```

For development, it uses a replace directive to use the local version:

```go
replace github.com/productscience/chain-protos => ../..
```

This allows you to make changes to the proto files and test them immediately without publishing. 