package main

import (
	"context"
	"fmt"
	"log"
	"time"

	networknodev1 "github.com/productscience/chain-protos/gen/network_node/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Set up a connection to the server
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Create a client
	client := networknodev1.NewRecordServiceClient(conn)

	// Set a record
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	setResp, err := client.SetRecord(ctx, &networknodev1.SetRecordRequest{
		Record: &networknodev1.Record{
			Key:   "greeting",
			Value: "Hello from the chain-protos example!",
		},
	})
	if err != nil {
		log.Fatalf("could not set record: %v", err)
	}
	fmt.Printf("Set record response: %v\n", setResp.Success)

	// Set another record
	setResp, err = client.SetRecord(ctx, &networknodev1.SetRecordRequest{
		Record: &networknodev1.Record{
			Key:   "timestamp",
			Value: time.Now().String(),
		},
	})
	if err != nil {
		log.Fatalf("could not set record: %v", err)
	}
	fmt.Printf("Set record response: %v\n", setResp.Success)

	// Get a record
	getResp, err := client.GetRecord(ctx, &networknodev1.GetRecordRequest{
		Key: "greeting",
	})
	if err != nil {
		log.Fatalf("could not get record: %v", err)
	}
	fmt.Printf("Retrieved record: %s = %s\n", getResp.Record.Key, getResp.Record.Value)

	// Get another record
	getResp, err = client.GetRecord(ctx, &networknodev1.GetRecordRequest{
		Key: "timestamp",
	})
	if err != nil {
		log.Fatalf("could not get record: %v", err)
	}
	fmt.Printf("Retrieved record: %s = %s\n", getResp.Record.Key, getResp.Record.Value)
}
