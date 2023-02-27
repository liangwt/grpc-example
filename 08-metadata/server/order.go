package main

import (
	"context"
	"io"
	"log"

	pb "github.com/liangwt/note/grpc/client_streaming_rpc_example/ecommerce"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
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

	md, ok := metadata.FromIncomingContext(ctx)
	log.Printf("has: %t. md: %v", ok, md)

	orders[orderReq.Id] = *orderReq

	grpc.SetHeader(ctx, metadata.Pairs("header-key1", "val1"))

	// create and send header
	header := metadata.Pairs("header-key", "val")
	grpc.SendHeader(ctx, header)

	// create and set trailer
	trailer := metadata.Pairs("trailer-key", "val")
	grpc.SetTrailer(ctx, trailer)

	return &wrapperspb.StringValue{Value: "Order Added: " + orderReq.Id}, nil
}

// 在这段程序中，我们对每一个 Recv 都进行了处理
// 当发现 io.EOF (流关闭) 后，需要将最终的响应结果发送给客户端，同时关闭正在另外一侧等待的 Recv
func (s *server) UpdateOrders(stream pb.OrderManagement_UpdateOrdersServer) error {
	md, ok := metadata.FromIncomingContext(stream.Context())
	log.Printf("has: %t. md: %v", ok, md)

	// create and send header
	header := metadata.Pairs("header-key", "val")
	stream.SetHeader(header)
	

	// create and set trailer
	trailer := metadata.Pairs("trailer-key", "val")
	stream.SetTrailer(trailer)

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
