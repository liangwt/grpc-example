package main

import (
	"context"
	"io"
	"log"

	pb "github.com/liangwt/note/grpc/bidirectional_streaming_rpc_example/ecommerce"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8009",
		grpc.WithInsecure(),
	)
	if err != nil {
		panic(err)
	}

	c := pb.NewOrderManagementClient(conn)
	ctx, cancelFn := context.WithCancel(context.Background())
	defer cancelFn()

	stream, err := c.ProcessOrders(ctx)
	if err != nil {
		panic(err)
	}

	go func() {
		if err := stream.Send(&wrapperspb.StringValue{Value: "101"}); err != nil {
			panic(err)
		}

		if err := stream.Send(&wrapperspb.StringValue{Value: "102"}); err != nil {
			panic(err)
		}

		if err := stream.CloseSend(); err != nil {
			panic(err)
		}
	}()

	for {
		combinedShipment, err := stream.Recv()
		if err == io.EOF {
			break
		}
		log.Println("Combined shipment : ", combinedShipment.OrderList)
	}
}
