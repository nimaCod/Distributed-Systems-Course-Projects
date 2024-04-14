// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: proto/ordering.proto

package proto

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

// OrderServiceClient is the client API for OrderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderServiceClient interface {
	// server streaming RPC
	GetOrderServerStreaming(ctx context.Context, in *NamesList, opts ...grpc.CallOption) (OrderService_GetOrderServerStreamingClient, error)
	// bidirectional streaming RPC
	GetOrderBidirectionalStreaming(ctx context.Context, opts ...grpc.CallOption) (OrderService_GetOrderBidirectionalStreamingClient, error)
}

type orderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderServiceClient(cc grpc.ClientConnInterface) OrderServiceClient {
	return &orderServiceClient{cc}
}

func (c *orderServiceClient) GetOrderServerStreaming(ctx context.Context, in *NamesList, opts ...grpc.CallOption) (OrderService_GetOrderServerStreamingClient, error) {
	stream, err := c.cc.NewStream(ctx, &OrderService_ServiceDesc.Streams[0], "/order_service.OrderService/GetOrderServerStreaming", opts...)
	if err != nil {
		return nil, err
	}
	x := &orderServiceGetOrderServerStreamingClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type OrderService_GetOrderServerStreamingClient interface {
	Recv() (*OrderResponse, error)
	grpc.ClientStream
}

type orderServiceGetOrderServerStreamingClient struct {
	grpc.ClientStream
}

func (x *orderServiceGetOrderServerStreamingClient) Recv() (*OrderResponse, error) {
	m := new(OrderResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *orderServiceClient) GetOrderBidirectionalStreaming(ctx context.Context, opts ...grpc.CallOption) (OrderService_GetOrderBidirectionalStreamingClient, error) {
	stream, err := c.cc.NewStream(ctx, &OrderService_ServiceDesc.Streams[1], "/order_service.OrderService/GetOrderBidirectionalStreaming", opts...)
	if err != nil {
		return nil, err
	}
	x := &orderServiceGetOrderBidirectionalStreamingClient{stream}
	return x, nil
}

type OrderService_GetOrderBidirectionalStreamingClient interface {
	Send(*OrderRequest) error
	Recv() (*OrderResponse, error)
	grpc.ClientStream
}

type orderServiceGetOrderBidirectionalStreamingClient struct {
	grpc.ClientStream
}

func (x *orderServiceGetOrderBidirectionalStreamingClient) Send(m *OrderRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *orderServiceGetOrderBidirectionalStreamingClient) Recv() (*OrderResponse, error) {
	m := new(OrderResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// OrderServiceServer is the server API for OrderService service.
// All implementations must embed UnimplementedOrderServiceServer
// for forward compatibility
type OrderServiceServer interface {
	// server streaming RPC
	GetOrderServerStreaming(*NamesList, OrderService_GetOrderServerStreamingServer) error
	// bidirectional streaming RPC
	GetOrderBidirectionalStreaming(OrderService_GetOrderBidirectionalStreamingServer) error
	mustEmbedUnimplementedOrderServiceServer()
}

// UnimplementedOrderServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOrderServiceServer struct {
}

func (UnimplementedOrderServiceServer) GetOrderServerStreaming(*NamesList, OrderService_GetOrderServerStreamingServer) error {
	return status.Errorf(codes.Unimplemented, "method GetOrderServerStreaming not implemented")
}
func (UnimplementedOrderServiceServer) GetOrderBidirectionalStreaming(OrderService_GetOrderBidirectionalStreamingServer) error {
	return status.Errorf(codes.Unimplemented, "method GetOrderBidirectionalStreaming not implemented")
}
func (UnimplementedOrderServiceServer) mustEmbedUnimplementedOrderServiceServer() {}

// UnsafeOrderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderServiceServer will
// result in compilation errors.
type UnsafeOrderServiceServer interface {
	mustEmbedUnimplementedOrderServiceServer()
}

func RegisterOrderServiceServer(s grpc.ServiceRegistrar, srv OrderServiceServer) {
	s.RegisterService(&OrderService_ServiceDesc, srv)
}

func _OrderService_GetOrderServerStreaming_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(NamesList)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(OrderServiceServer).GetOrderServerStreaming(m, &orderServiceGetOrderServerStreamingServer{stream})
}

type OrderService_GetOrderServerStreamingServer interface {
	Send(*OrderResponse) error
	grpc.ServerStream
}

type orderServiceGetOrderServerStreamingServer struct {
	grpc.ServerStream
}

func (x *orderServiceGetOrderServerStreamingServer) Send(m *OrderResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _OrderService_GetOrderBidirectionalStreaming_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(OrderServiceServer).GetOrderBidirectionalStreaming(&orderServiceGetOrderBidirectionalStreamingServer{stream})
}

type OrderService_GetOrderBidirectionalStreamingServer interface {
	Send(*OrderResponse) error
	Recv() (*OrderRequest, error)
	grpc.ServerStream
}

type orderServiceGetOrderBidirectionalStreamingServer struct {
	grpc.ServerStream
}

func (x *orderServiceGetOrderBidirectionalStreamingServer) Send(m *OrderResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *orderServiceGetOrderBidirectionalStreamingServer) Recv() (*OrderRequest, error) {
	m := new(OrderRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// OrderService_ServiceDesc is the grpc.ServiceDesc for OrderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "order_service.OrderService",
	HandlerType: (*OrderServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetOrderServerStreaming",
			Handler:       _OrderService_GetOrderServerStreaming_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetOrderBidirectionalStreaming",
			Handler:       _OrderService_GetOrderBidirectionalStreaming_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/ordering.proto",
}
