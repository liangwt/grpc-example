package main

import (
	"context"
	"log"

	pb "github.com/liangwt/note/grpc/client_streaming_rpc_example/ecommerce"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

var _ pb.OrderManagementServer = &server{}

var orders = make(map[string]pb.Order, 0)

type server struct {
	pb.UnimplementedOrderManagementServer
}

// Simple RPC
func (s *server) AddOrder(ctx context.Context, orderReq *pb.Order) (*wrapperspb.StringValue, error) {
	log.Printf("Order Added. ID : %v", orderReq.Id)

	select {
	case <-ctx.Done():
		return nil, status.Errorf(codes.Canceled, "Client cancelled, abandoning.")
	default:
	}

	orders[orderReq.Id] = *orderReq

	return &wrapperspb.StringValue{Value: "Order Added: " + orderReq.Id}, nil
}
