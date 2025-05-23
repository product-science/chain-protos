// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: v1/network_node.proto

package network_nodev1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	NetworkNodeService_JoinTraining_FullMethodName          = "/network_node.v1.NetworkNodeService/JoinTraining"
	NetworkNodeService_GetJoinTrainingStatus_FullMethodName = "/network_node.v1.NetworkNodeService/GetJoinTrainingStatus"
	NetworkNodeService_SendHeartbeat_FullMethodName         = "/network_node.v1.NetworkNodeService/SendHeartbeat"
	NetworkNodeService_GetAliveNodes_FullMethodName         = "/network_node.v1.NetworkNodeService/GetAliveNodes"
	NetworkNodeService_SetBarrier_FullMethodName            = "/network_node.v1.NetworkNodeService/SetBarrier"
	NetworkNodeService_GetBarrierStatus_FullMethodName      = "/network_node.v1.NetworkNodeService/GetBarrierStatus"
	NetworkNodeService_SetStoreRecord_FullMethodName        = "/network_node.v1.NetworkNodeService/SetStoreRecord"
	NetworkNodeService_GetStoreRecord_FullMethodName        = "/network_node.v1.NetworkNodeService/GetStoreRecord"
	NetworkNodeService_ListStoreKeys_FullMethodName         = "/network_node.v1.NetworkNodeService/ListStoreKeys"
)

// NetworkNodeServiceClient is the client API for NetworkNodeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NetworkNodeServiceClient interface {
	JoinTraining(ctx context.Context, in *JoinTrainingRequest, opts ...grpc.CallOption) (*MLNodeTrainStatus, error)
	GetJoinTrainingStatus(ctx context.Context, in *JoinTrainingRequest, opts ...grpc.CallOption) (*MLNodeTrainStatus, error)
	SendHeartbeat(ctx context.Context, in *HeartbeatRequest, opts ...grpc.CallOption) (*HeartbeatResponse, error)
	GetAliveNodes(ctx context.Context, in *GetAliveNodesRequest, opts ...grpc.CallOption) (*GetAliveNodesResponse, error)
	SetBarrier(ctx context.Context, in *SetBarrierRequest, opts ...grpc.CallOption) (*SetBarrierResponse, error)
	GetBarrierStatus(ctx context.Context, in *GetBarrierStatusRequest, opts ...grpc.CallOption) (*GetBarrierStatusResponse, error)
	SetStoreRecord(ctx context.Context, in *SetStoreRecordRequest, opts ...grpc.CallOption) (*SetStoreRecordResponse, error)
	GetStoreRecord(ctx context.Context, in *GetStoreRecordRequest, opts ...grpc.CallOption) (*GetStoreRecordResponse, error)
	ListStoreKeys(ctx context.Context, in *StoreListKeysRequest, opts ...grpc.CallOption) (*StoreListKeysResponse, error)
}

type networkNodeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNetworkNodeServiceClient(cc grpc.ClientConnInterface) NetworkNodeServiceClient {
	return &networkNodeServiceClient{cc}
}

func (c *networkNodeServiceClient) JoinTraining(ctx context.Context, in *JoinTrainingRequest, opts ...grpc.CallOption) (*MLNodeTrainStatus, error) {
	out := new(MLNodeTrainStatus)
	err := c.cc.Invoke(ctx, NetworkNodeService_JoinTraining_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkNodeServiceClient) GetJoinTrainingStatus(ctx context.Context, in *JoinTrainingRequest, opts ...grpc.CallOption) (*MLNodeTrainStatus, error) {
	out := new(MLNodeTrainStatus)
	err := c.cc.Invoke(ctx, NetworkNodeService_GetJoinTrainingStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkNodeServiceClient) SendHeartbeat(ctx context.Context, in *HeartbeatRequest, opts ...grpc.CallOption) (*HeartbeatResponse, error) {
	out := new(HeartbeatResponse)
	err := c.cc.Invoke(ctx, NetworkNodeService_SendHeartbeat_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkNodeServiceClient) GetAliveNodes(ctx context.Context, in *GetAliveNodesRequest, opts ...grpc.CallOption) (*GetAliveNodesResponse, error) {
	out := new(GetAliveNodesResponse)
	err := c.cc.Invoke(ctx, NetworkNodeService_GetAliveNodes_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkNodeServiceClient) SetBarrier(ctx context.Context, in *SetBarrierRequest, opts ...grpc.CallOption) (*SetBarrierResponse, error) {
	out := new(SetBarrierResponse)
	err := c.cc.Invoke(ctx, NetworkNodeService_SetBarrier_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkNodeServiceClient) GetBarrierStatus(ctx context.Context, in *GetBarrierStatusRequest, opts ...grpc.CallOption) (*GetBarrierStatusResponse, error) {
	out := new(GetBarrierStatusResponse)
	err := c.cc.Invoke(ctx, NetworkNodeService_GetBarrierStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkNodeServiceClient) SetStoreRecord(ctx context.Context, in *SetStoreRecordRequest, opts ...grpc.CallOption) (*SetStoreRecordResponse, error) {
	out := new(SetStoreRecordResponse)
	err := c.cc.Invoke(ctx, NetworkNodeService_SetStoreRecord_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkNodeServiceClient) GetStoreRecord(ctx context.Context, in *GetStoreRecordRequest, opts ...grpc.CallOption) (*GetStoreRecordResponse, error) {
	out := new(GetStoreRecordResponse)
	err := c.cc.Invoke(ctx, NetworkNodeService_GetStoreRecord_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkNodeServiceClient) ListStoreKeys(ctx context.Context, in *StoreListKeysRequest, opts ...grpc.CallOption) (*StoreListKeysResponse, error) {
	out := new(StoreListKeysResponse)
	err := c.cc.Invoke(ctx, NetworkNodeService_ListStoreKeys_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NetworkNodeServiceServer is the server API for NetworkNodeService service.
// All implementations must embed UnimplementedNetworkNodeServiceServer
// for forward compatibility
type NetworkNodeServiceServer interface {
	JoinTraining(context.Context, *JoinTrainingRequest) (*MLNodeTrainStatus, error)
	GetJoinTrainingStatus(context.Context, *JoinTrainingRequest) (*MLNodeTrainStatus, error)
	SendHeartbeat(context.Context, *HeartbeatRequest) (*HeartbeatResponse, error)
	GetAliveNodes(context.Context, *GetAliveNodesRequest) (*GetAliveNodesResponse, error)
	SetBarrier(context.Context, *SetBarrierRequest) (*SetBarrierResponse, error)
	GetBarrierStatus(context.Context, *GetBarrierStatusRequest) (*GetBarrierStatusResponse, error)
	SetStoreRecord(context.Context, *SetStoreRecordRequest) (*SetStoreRecordResponse, error)
	GetStoreRecord(context.Context, *GetStoreRecordRequest) (*GetStoreRecordResponse, error)
	ListStoreKeys(context.Context, *StoreListKeysRequest) (*StoreListKeysResponse, error)
	mustEmbedUnimplementedNetworkNodeServiceServer()
}

// UnimplementedNetworkNodeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNetworkNodeServiceServer struct {
}

func (UnimplementedNetworkNodeServiceServer) JoinTraining(context.Context, *JoinTrainingRequest) (*MLNodeTrainStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinTraining not implemented")
}
func (UnimplementedNetworkNodeServiceServer) GetJoinTrainingStatus(context.Context, *JoinTrainingRequest) (*MLNodeTrainStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetJoinTrainingStatus not implemented")
}
func (UnimplementedNetworkNodeServiceServer) SendHeartbeat(context.Context, *HeartbeatRequest) (*HeartbeatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendHeartbeat not implemented")
}
func (UnimplementedNetworkNodeServiceServer) GetAliveNodes(context.Context, *GetAliveNodesRequest) (*GetAliveNodesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAliveNodes not implemented")
}
func (UnimplementedNetworkNodeServiceServer) SetBarrier(context.Context, *SetBarrierRequest) (*SetBarrierResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetBarrier not implemented")
}
func (UnimplementedNetworkNodeServiceServer) GetBarrierStatus(context.Context, *GetBarrierStatusRequest) (*GetBarrierStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBarrierStatus not implemented")
}
func (UnimplementedNetworkNodeServiceServer) SetStoreRecord(context.Context, *SetStoreRecordRequest) (*SetStoreRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetStoreRecord not implemented")
}
func (UnimplementedNetworkNodeServiceServer) GetStoreRecord(context.Context, *GetStoreRecordRequest) (*GetStoreRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStoreRecord not implemented")
}
func (UnimplementedNetworkNodeServiceServer) ListStoreKeys(context.Context, *StoreListKeysRequest) (*StoreListKeysResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListStoreKeys not implemented")
}
func (UnimplementedNetworkNodeServiceServer) mustEmbedUnimplementedNetworkNodeServiceServer() {}

// UnsafeNetworkNodeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NetworkNodeServiceServer will
// result in compilation errors.
type UnsafeNetworkNodeServiceServer interface {
	mustEmbedUnimplementedNetworkNodeServiceServer()
}

func RegisterNetworkNodeServiceServer(s grpc.ServiceRegistrar, srv NetworkNodeServiceServer) {
	s.RegisterService(&NetworkNodeService_ServiceDesc, srv)
}

func _NetworkNodeService_JoinTraining_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinTrainingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkNodeServiceServer).JoinTraining(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkNodeService_JoinTraining_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkNodeServiceServer).JoinTraining(ctx, req.(*JoinTrainingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkNodeService_GetJoinTrainingStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinTrainingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkNodeServiceServer).GetJoinTrainingStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkNodeService_GetJoinTrainingStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkNodeServiceServer).GetJoinTrainingStatus(ctx, req.(*JoinTrainingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkNodeService_SendHeartbeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HeartbeatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkNodeServiceServer).SendHeartbeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkNodeService_SendHeartbeat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkNodeServiceServer).SendHeartbeat(ctx, req.(*HeartbeatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkNodeService_GetAliveNodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAliveNodesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkNodeServiceServer).GetAliveNodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkNodeService_GetAliveNodes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkNodeServiceServer).GetAliveNodes(ctx, req.(*GetAliveNodesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkNodeService_SetBarrier_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetBarrierRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkNodeServiceServer).SetBarrier(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkNodeService_SetBarrier_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkNodeServiceServer).SetBarrier(ctx, req.(*SetBarrierRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkNodeService_GetBarrierStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBarrierStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkNodeServiceServer).GetBarrierStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkNodeService_GetBarrierStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkNodeServiceServer).GetBarrierStatus(ctx, req.(*GetBarrierStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkNodeService_SetStoreRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetStoreRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkNodeServiceServer).SetStoreRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkNodeService_SetStoreRecord_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkNodeServiceServer).SetStoreRecord(ctx, req.(*SetStoreRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkNodeService_GetStoreRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStoreRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkNodeServiceServer).GetStoreRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkNodeService_GetStoreRecord_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkNodeServiceServer).GetStoreRecord(ctx, req.(*GetStoreRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkNodeService_ListStoreKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreListKeysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkNodeServiceServer).ListStoreKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkNodeService_ListStoreKeys_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkNodeServiceServer).ListStoreKeys(ctx, req.(*StoreListKeysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NetworkNodeService_ServiceDesc is the grpc.ServiceDesc for NetworkNodeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NetworkNodeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "network_node.v1.NetworkNodeService",
	HandlerType: (*NetworkNodeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "JoinTraining",
			Handler:    _NetworkNodeService_JoinTraining_Handler,
		},
		{
			MethodName: "GetJoinTrainingStatus",
			Handler:    _NetworkNodeService_GetJoinTrainingStatus_Handler,
		},
		{
			MethodName: "SendHeartbeat",
			Handler:    _NetworkNodeService_SendHeartbeat_Handler,
		},
		{
			MethodName: "GetAliveNodes",
			Handler:    _NetworkNodeService_GetAliveNodes_Handler,
		},
		{
			MethodName: "SetBarrier",
			Handler:    _NetworkNodeService_SetBarrier_Handler,
		},
		{
			MethodName: "GetBarrierStatus",
			Handler:    _NetworkNodeService_GetBarrierStatus_Handler,
		},
		{
			MethodName: "SetStoreRecord",
			Handler:    _NetworkNodeService_SetStoreRecord_Handler,
		},
		{
			MethodName: "GetStoreRecord",
			Handler:    _NetworkNodeService_GetStoreRecord_Handler,
		},
		{
			MethodName: "ListStoreKeys",
			Handler:    _NetworkNodeService_ListStoreKeys_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/network_node.proto",
}
