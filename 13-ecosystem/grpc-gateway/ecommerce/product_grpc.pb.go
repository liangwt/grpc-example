// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: product.proto

package ecommerce

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// OrderManagementClient is the client API for OrderManagement service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderManagementClient interface {
	GetOrder(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*Order, error)
	AddOrder1(ctx context.Context, in *Order, opts ...grpc.CallOption) (*wrapperspb.StringValue, error)
	AddOrder2(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (*wrapperspb.StringValue, error)
}

type orderManagementClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderManagementClient(cc grpc.ClientConnInterface) OrderManagementClient {
	return &orderManagementClient{cc}
}

func (c *orderManagementClient) GetOrder(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*Order, error) {
	out := new(Order)
	err := c.cc.Invoke(ctx, "/ecommerce.OrderManagement/getOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderManagementClient) AddOrder1(ctx context.Context, in *Order, opts ...grpc.CallOption) (*wrapperspb.StringValue, error) {
	out := new(wrapperspb.StringValue)
	err := c.cc.Invoke(ctx, "/ecommerce.OrderManagement/addOrder1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderManagementClient) AddOrder2(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (*wrapperspb.StringValue, error) {
	out := new(wrapperspb.StringValue)
	err := c.cc.Invoke(ctx, "/ecommerce.OrderManagement/addOrder2", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderManagementServer is the server API for OrderManagement service.
// All implementations must embed UnimplementedOrderManagementServer
// for forward compatibility
type OrderManagementServer interface {
	GetOrder(context.Context, *wrapperspb.StringValue) (*Order, error)
	AddOrder1(context.Context, *Order) (*wrapperspb.StringValue, error)
	AddOrder2(context.Context, *OrderRequest) (*wrapperspb.StringValue, error)
	mustEmbedUnimplementedOrderManagementServer()
}

// UnimplementedOrderManagementServer must be embedded to have forward compatible implementations.
type UnimplementedOrderManagementServer struct {
}

func (UnimplementedOrderManagementServer) GetOrder(context.Context, *wrapperspb.StringValue) (*Order, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrder not implemented")
}
func (UnimplementedOrderManagementServer) AddOrder1(context.Context, *Order) (*wrapperspb.StringValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddOrder1 not implemented")
}
func (UnimplementedOrderManagementServer) AddOrder2(context.Context, *OrderRequest) (*wrapperspb.StringValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddOrder2 not implemented")
}
func (UnimplementedOrderManagementServer) mustEmbedUnimplementedOrderManagementServer() {}

// UnsafeOrderManagementServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderManagementServer will
// result in compilation errors.
type UnsafeOrderManagementServer interface {
	mustEmbedUnimplementedOrderManagementServer()
}

func RegisterOrderManagementServer(s grpc.ServiceRegistrar, srv OrderManagementServer) {
	s.RegisterService(&OrderManagement_ServiceDesc, srv)
}

func _OrderManagement_GetOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrapperspb.StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderManagementServer).GetOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ecommerce.OrderManagement/getOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderManagementServer).GetOrder(ctx, req.(*wrapperspb.StringValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderManagement_AddOrder1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Order)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderManagementServer).AddOrder1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ecommerce.OrderManagement/addOrder1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderManagementServer).AddOrder1(ctx, req.(*Order))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderManagement_AddOrder2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderManagementServer).AddOrder2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ecommerce.OrderManagement/addOrder2",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderManagementServer).AddOrder2(ctx, req.(*OrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OrderManagement_ServiceDesc is the grpc.ServiceDesc for OrderManagement service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderManagement_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ecommerce.OrderManagement",
	HandlerType: (*OrderManagementServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getOrder",
			Handler:    _OrderManagement_GetOrder_Handler,
		},
		{
			MethodName: "addOrder1",
			Handler:    _OrderManagement_AddOrder1_Handler,
		},
		{
			MethodName: "addOrder2",
			Handler:    _OrderManagement_AddOrder2_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "product.proto",
}
