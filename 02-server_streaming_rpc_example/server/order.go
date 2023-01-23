package main

import (
	"fmt"
	"strings"

	pb "github.com/liangwt/note/grpc/server_streaming_rpc_example/ecommerce"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

var _ pb.OrderManagementServer = &server{}

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

type server struct {
	pb.UnimplementedOrderManagementServer
}

func (s *server) SearchOrders(query *wrapperspb.StringValue, stram pb.OrderManagement_SearchOrdersServer) error {
	for _, order := range orders {
		for _, str := range order.Items {
			if strings.Contains(str, query.Value) {
				err := stram.Send(&order)
				if err != nil {
					return fmt.Errorf("error send: %v", err)
				}
			}
		}
	}

	return nil
}
