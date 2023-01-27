package main

import (
	"context"
	"log"

	pb "github.com/liangwt/note/grpc/ecosystem/grpc-gateway/ecommerce"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func (s *server) GetOrder(ctx context.Context, req *pb.GetOrderReq) (*pb.Order, error) {
	log.Print("GetOrder Request -> : ", req)

	ord, exists := orders[req.Id.Value]
	if exists {
		return &ord, status.New(codes.OK, "").Err()
	}

	return nil, status.Errorf(codes.NotFound, "Order does not exist. : ", req)
}
