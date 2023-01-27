package main

import (
	"context"
	"log"

	pb "github.com/liangwt/note/grpc/secure/ecommerce"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

var _ pb.OrderManagementServer = &server{}

var orders = map[string]pb.Order{
	"101": {
		Id:          "101",
		Items:       []string{},
		Description: "example",
		Price:       0,
		Destination: "example",
	},
}

type server struct {
	pb.UnimplementedOrderManagementServer
}

func (s *server) GetOrder(ctx context.Context, orderId *wrapperspb.StringValue) (*pb.Order, error) {
	log.Print("GetOrder Request -> : ", orderId)

	ord, exists := orders[orderId.Value]
	if exists {
		return &ord, status.New(codes.OK, "").Err()
	}

	return nil, status.Errorf(codes.NotFound, "Order does not exist. : ", orderId)
}
