package main

import (
	"context"
	"log"

	pb "github.com/liangwt/note/grpc/client_streaming_rpc_example/ecommerce"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8009", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	c := pb.NewOrderManagementClient(conn)
	ctx, cancelFn := context.WithCancel(context.Background())
	defer cancelFn()

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
}
