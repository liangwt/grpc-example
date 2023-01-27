package main

import (
	"context"
	"io"
	"log"

	pb "github.com/liangwt/note/grpc/name_resolver_lb_example/ecommerce"
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

	orders[orderReq.Id] = *orderReq

	return &wrapperspb.StringValue{Value: "Order Added: " + orderReq.Id}, nil
}

// 在这段程序中，我们对每一个 Recv 都进行了处理
// 当发现 io.EOF (流关闭) 后，需要将最终的响应结果发送给客户端，同时关闭正在另外一侧等待的 Recv
func (s *server) UpdateOrders(stream pb.OrderManagement_UpdateOrdersServer) error {
	ordersStr := "Updated Order IDs : "
	for {
		order, err := stream.Recv()
		if err == io.EOF {
			// Finished reading the order stream.
			return stream.SendAndClose(
				&wrapperspb.StringValue{Value: "Orders processed " + ordersStr})
		}
		
		// Update order
		orders[order.Id] = *order

		log.Println("Order ID ", order.Id, ": Updated")
		ordersStr += order.Id + ", "
	}
}
