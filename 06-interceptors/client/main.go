package main

import (
	"context"
	"log"

	pb "github.com/liangwt/note/grpc/bidirectional_streaming_rpc_example/ecommerce"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8009",
		grpc.WithInsecure(),
		grpc.WithChainUnaryInterceptor(
			orderUnaryClientInterceptor,
		),
		grpc.WithChainStreamInterceptor(
			orderStreamClientInterceptor,
		),
	)
	if err != nil {
		panic(err)
	}

	c := pb.NewOrderManagementClient(conn)
	ctx, cancelFn := context.WithCancel(context.Background())
	defer cancelFn()

	// // Add Order
	// order := pb.Order{Id: "101", Items: []string{"iPhone XS", "Mac Book Pro"}, Destination: "San Jose, CA", Price: 2300.00}
	// res, err := c.AddOrder(ctx, &order)
	// if err != nil {
	// 	panic(err)
	// }

	// log.Print("AddOrder Response -> ", res.Value)

	/////////// unary RPC //////////

	// // Get Order
	// retrievedOrder, err := c.GetOrder(ctx, &wrapperspb.StringValue{Value: "101"})
	// if err != nil {
	// 	panic(err)
	// }

	// log.Print("GetOrder Response -> : ", retrievedOrder)

	/////////// 客户端流式 RPC //////////

	stream, err := c.UpdateOrders(ctx)
	if err != nil {
		panic(err)
	}

	if err := stream.Send(&pb.Order{
		Id:          "00",
		Items:       []string{"A", "B"},
		Description: "A with B",
		Price:       0.11,
		Destination: "ABC",
	}); err != nil {
		panic(err)
	}

	if err := stream.Send(&pb.Order{
		Id:          "01",
		Items:       []string{"C", "D"},
		Description: "C with D",
		Price:       1.11,
		Destination: "ABCDEFG",
	}); err != nil {
		panic(err)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		panic(err)
	}

	log.Printf("Update Orders Res : %s", res)

	/////////// 双向流式 RPC //////////

	// stream, err := c.ProcessOrders(ctx)
	// if err != nil {
	// 	panic(err)
	// }

	// go func() {
	// 	if err := stream.Send(&wrapperspb.StringValue{Value: "101"}); err != nil {
	// 		panic(err)
	// 	}

	// 	if err := stream.Send(&wrapperspb.StringValue{Value: "102"}); err != nil {
	// 		panic(err)
	// 	}

	// 	if err := stream.CloseSend(); err != nil {
	// 		panic(err)
	// 	}
	// }()

	// for {
	// 	combinedShipment, err := stream.Recv()
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	log.Println("Combined shipment : ", combinedShipment.OrderList)
	// }
}
