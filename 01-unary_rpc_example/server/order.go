package main

import (
	"context"
	"log"

	pb "github.com/liangwt/note/grpc/unary_rpc_example/ecommerce"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

var _ pb.OrderManagementServer = &OrderManagementImpl{}

var orders = make(map[string]pb.Order)

type OrderManagementImpl struct {
	pb.UnimplementedOrderManagementServer
}

// Simple RPC
func (s *OrderManagementImpl) AddOrder(ctx context.Context, orderReq *pb.Order) (*wrapperspb.StringValue, error) {
	log.Printf("Order Added. ID : %v", orderReq.Id)
	orders[orderReq.Id] = *orderReq
	return &wrapperspb.StringValue{Value: "Order Added: " + orderReq.Id}, nil
}

// Simple RPC
func (s *OrderManagementImpl) GetOrder(ctx context.Context, orderId *wrapperspb.StringValue) (*pb.Order, error) {
	ord, exists := orders[orderId.Value]
	if exists {
		return &ord, status.New(codes.OK, "").Err()
	}

	return nil, status.Errorf(codes.NotFound, "Order does not exist. : ", orderId)

}
