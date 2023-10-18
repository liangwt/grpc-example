package main

import (
	"context"
	"log"

	pb "github.com/liangwt/note/grpc/ecosystem/grpc-gateway/ecommerce"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

var _ pb.OrderManagementServer = &OrderManagementImpl{}

var orders = map[string]pb.Order{
	"101": {
		Id: "101",
		Items: []string{
			"Google",
			"Baidu",
		},
		Description: "example",
		Price:       0,
		Destination: &wrapperspb.StringValue{
			Value: "example",
		},
	},
}

type OrderManagementImpl struct {
	pb.UnimplementedOrderManagementServer
}

func (s *OrderManagementImpl) AddOrder2(ctx context.Context, orderReq *pb.OrderRequest) (*wrapperspb.StringValue, error) {
	log.Printf("Order Added. ID : %v", orderReq.Order.Id)
	orders[orderReq.Order.Id] = *orderReq.Order
	return &wrapperspb.StringValue{Value: "Order Added: " + orderReq.Order.Id}, nil
}

// Simple RPC
func (s *OrderManagementImpl) AddOrder1(ctx context.Context, orderReq *pb.Order) (*wrapperspb.StringValue, error) {
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
