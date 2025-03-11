#!/usr/bin/env python3

import grpc
import argparse
from network_node.v1 import network_node_pb2
from network_node.v1 import network_node_pb2_grpc

def set_record(stub, key, value):
    """Set a record with the given key and value"""
    record = network_node_pb2.Record(key=key, value=value)
    request = network_node_pb2.SetRecordRequest(record=record)
    
    try:
        response = stub.SetRecord(request)
        if response.success:
            print(f"Successfully set record: {key}={value}")
        else:
            print(f"Failed to set record: {key}={value}")
    except grpc.RpcError as e:
        print(f"RPC error: {e.code()}: {e.details()}")

def get_record(stub, key):
    """Get a record with the given key"""
    request = network_node_pb2.GetRecordRequest(key=key)
    
    try:
        response = stub.GetRecord(request)
        if response.record:
            print(f"Retrieved record: {response.record.key}={response.record.value}")
        else:
            print(f"No record found for key: {key}")
    except grpc.RpcError as e:
        print(f"RPC error: {e.code()}: {e.details()}")

def main():
    parser = argparse.ArgumentParser(description='gRPC Client for RecordService')
    parser.add_argument('--server', default='localhost:50051', help='Server address')
    parser.add_argument('--action', choices=['set', 'get'], required=True, help='Action to perform')
    parser.add_argument('--key', required=True, help='Record key')
    parser.add_argument('--value', help='Record value (required for set action)')
    
    args = parser.parse_args()
    
    # Check if value is provided for set action
    if args.action == 'set' and not args.value:
        parser.error('--value is required for set action')
    
    # Create a gRPC channel
    with grpc.insecure_channel(args.server) as channel:
        stub = network_node_pb2_grpc.RecordServiceStub(channel)
        
        if args.action == 'set':
            set_record(stub, args.key, args.value)
        else:  # get
            get_record(stub, args.key)

if __name__ == '__main__':
    main() 