package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"

	pb "github.com/productscience/chain-protos/network_node_direct/v1"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedRecordServiceServer
	mu      sync.Mutex
	records map[string]string
}

func (s *server) SetRecord(ctx context.Context, req *pb.SetRecordRequest) (*pb.SetRecordResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if req.Record == nil {
		return &pb.SetRecordResponse{Success: false}, fmt.Errorf("record is nil")
	}

	s.records[req.Record.Key] = req.Record.Value
	log.Printf("Set record: key=%s, value=%s", req.Record.Key, req.Record.Value)

	return &pb.SetRecordResponse{Success: true}, nil
}

func (s *server) GetRecord(ctx context.Context, req *pb.GetRecordRequest) (*pb.GetRecordResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	value, exists := s.records[req.Key]
	if !exists {
		return &pb.GetRecordResponse{Record: nil}, fmt.Errorf("key %s not found", req.Key)
	}

	log.Printf("Got record: key=%s, value=%s", req.Key, value)
	return &pb.GetRecordResponse{Record: &pb.Record{Key: req.Key, Value: value}}, nil
}

func main() {
	port := ":50051"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterRecordServiceServer(s, &server{
		records: make(map[string]string),
	})

	log.Printf("Server listening on port %s", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
