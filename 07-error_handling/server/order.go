package main

import (
	"context"
	"fmt"

	pb "github.com/liangwt/note/grpc/error_handling/ecommerce"
	epb "google.golang.org/genproto/googleapis/rpc/errdetails"
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
		Destination: "example",
	},
}

type OrderManagementImpl struct {
	pb.UnimplementedOrderManagementServer
}

// Simple RPC
func (s *OrderManagementImpl) GetOrder(ctx context.Context, orderId *wrapperspb.StringValue) (*pb.Order, error) {
	ord, exists := orders[orderId.Value]
	if exists {
		return &ord, status.New(codes.OK, "ok").Err()
	}

	st := status.New(codes.InvalidArgument,
		"Order does not exist. order id: "+orderId.Value)

	details, err := st.WithDetails(
		&epb.BadRequest_FieldViolation{
			Field:       "ID",
			Description: fmt.Sprintf("Order ID received is not valid"),
		},
	)
	if err == nil {
		return nil, details.Err()
	}

	return nil, st.Err()
}
