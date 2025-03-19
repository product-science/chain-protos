package main

import (
	"context"
	"fmt"
	"log"
	"net"

	networknodev1 "github.com/productscience/chain-protos/gen/network_node/v1"
	"google.golang.org/grpc"
)

// RecordServer implements the RecordService
type RecordServer struct {
	networknodev1.UnimplementedRecordServiceServer
	records map[string]string
}

// SetRecord implements the SetRecord RPC method
func (s *RecordServer) SetRecord(ctx context.Context, req *networknodev1.SetRecordRequest) (*networknodev1.SetRecordResponse, error) {
	if req.Record == nil {
		return &networknodev1.SetRecordResponse{Success: false}, fmt.Errorf("record cannot be nil")
	}

	s.records[req.Record.Key] = req.Record.Value
	fmt.Printf("Set record: %s = %s\n", req.Record.Key, req.Record.Value)

	return &networknodev1.SetRecordResponse{Success: true}, nil
}

// GetRecord implements the GetRecord RPC method
func (s *RecordServer) GetRecord(ctx context.Context, req *networknodev1.GetRecordRequest) (*networknodev1.GetRecordResponse, error) {
	value, exists := s.records[req.Key]
	if !exists {
		return &networknodev1.GetRecordResponse{}, fmt.Errorf("key not found: %s", req.Key)
	}

	fmt.Printf("Get record: %s = %s\n", req.Key, value)

	return &networknodev1.GetRecordResponse{
		Record: &networknodev1.Record{
			Key:   req.Key,
			Value: value,
		},
	}, nil
}

func main() {
	// Create a new server instance
	server := &RecordServer{
		records: make(map[string]string),
	}

	// Create a gRPC server
	grpcServer := grpc.NewServer()

	// Register our service
	networknodev1.RegisterRecordServiceServer(grpcServer, server)

	// Start listening
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	fmt.Println("Starting gRPC server on :50051")
	fmt.Println("This example demonstrates using the chain-protos generated code.")
	fmt.Println("To use this service, connect with a gRPC client to localhost:50051")

	// Start the server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
