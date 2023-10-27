// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.4
// source: modules/inventory/inventoryPb/inventoryPb.proto

package go_mc_kafka

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

// InventoryGrpcServiceClient is the client API for InventoryGrpcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InventoryGrpcServiceClient interface {
	IsAvailableToSell(ctx context.Context, in *IsAvailableToSellReq, opts ...grpc.CallOption) (*IsAvailableToSellRes, error)
}

type inventoryGrpcServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInventoryGrpcServiceClient(cc grpc.ClientConnInterface) InventoryGrpcServiceClient {
	return &inventoryGrpcServiceClient{cc}
}

func (c *inventoryGrpcServiceClient) IsAvailableToSell(ctx context.Context, in *IsAvailableToSellReq, opts ...grpc.CallOption) (*IsAvailableToSellRes, error) {
	out := new(IsAvailableToSellRes)
	err := c.cc.Invoke(ctx, "/InventoryGrpcService/IsAvailableToSell", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InventoryGrpcServiceServer is the server API for InventoryGrpcService service.
// All implementations must embed UnimplementedInventoryGrpcServiceServer
// for forward compatibility
type InventoryGrpcServiceServer interface {
	IsAvailableToSell(context.Context, *IsAvailableToSellReq) (*IsAvailableToSellRes, error)
	mustEmbedUnimplementedInventoryGrpcServiceServer()
}

// UnimplementedInventoryGrpcServiceServer must be embedded to have forward compatible implementations.
type UnimplementedInventoryGrpcServiceServer struct {
}

func (UnimplementedInventoryGrpcServiceServer) IsAvailableToSell(context.Context, *IsAvailableToSellReq) (*IsAvailableToSellRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsAvailableToSell not implemented")
}
func (UnimplementedInventoryGrpcServiceServer) mustEmbedUnimplementedInventoryGrpcServiceServer() {}

// UnsafeInventoryGrpcServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InventoryGrpcServiceServer will
// result in compilation errors.
type UnsafeInventoryGrpcServiceServer interface {
	mustEmbedUnimplementedInventoryGrpcServiceServer()
}

func RegisterInventoryGrpcServiceServer(s grpc.ServiceRegistrar, srv InventoryGrpcServiceServer) {
	s.RegisterService(&InventoryGrpcService_ServiceDesc, srv)
}

func _InventoryGrpcService_IsAvailableToSell_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsAvailableToSellReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryGrpcServiceServer).IsAvailableToSell(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/InventoryGrpcService/IsAvailableToSell",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryGrpcServiceServer).IsAvailableToSell(ctx, req.(*IsAvailableToSellReq))
	}
	return interceptor(ctx, in, info, handler)
}

// InventoryGrpcService_ServiceDesc is the grpc.ServiceDesc for InventoryGrpcService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InventoryGrpcService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "InventoryGrpcService",
	HandlerType: (*InventoryGrpcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsAvailableToSell",
			Handler:    _InventoryGrpcService_IsAvailableToSell_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "modules/inventory/inventoryPb/inventoryPb.proto",
}
