# gRPC Client Example

This example demonstrates how to create a gRPC client that communicates with the server using the protocol definitions from the chain-protos repository.

## Overview

This is a simple gRPC client that connects to the server implemented in the parent directory. It demonstrates how to:

1. Connect to a gRPC server
2. Use the generated client code to call RPC methods
3. Handle the responses

## Running the Example

1. First, make sure the server is running:

   ```bash
   # From the examples/go-usage directory
   go run main.go
   ```

2. In another terminal, run the client:

   ```bash
   # From the examples/go-usage/client directory
   go run main.go
   ```

3. You should see output showing the client setting and retrieving records from the server.

## How It Works

The client:

1. Creates a connection to the gRPC server running on localhost:50051
2. Uses the generated client code to create a `RecordServiceClient`
3. Calls the `SetRecord` method to store data on the server
4. Calls the `GetRecord` method to retrieve data from the server

This demonstrates how easy it is to use the generated code from the chain-protos repository to communicate between services. 