package main

import (
	"context"
	"log"

	pb "github.com/liangwt/note/grpc/client_streaming_rpc_example/ecommerce"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8009", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	ctx, cancelFn := context.WithCancel(context.Background())
	defer cancelFn()
	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("k1", "v1", "k2", "v2"))
	ctx = metadata.AppendToOutgoingContext(ctx, "k3", "v4")

	c := pb.NewOrderManagementClient(conn)

	// RPC using the context with new metadata.
	var header, trailer metadata.MD

	// Add Order
	order := pb.Order{Id: "101", Items: []string{"iPhone XS", "Mac Book Pro"}, Destination: "San Jose, CA", Price: 2300.00}
	res, err := c.AddOrder(ctx, &order, grpc.Header(&header), grpc.Trailer(&trailer))
	if err != nil {
		panic(err)
	}

	log.Printf("#AddOrder## header: %v. trailer: %v", header, trailer)

	//////////

	stream, err := c.UpdateOrders(ctx)
	if err != nil {
		panic(err)
	}
	// retrieve header
	header, _ = stream.Header()
	// retrieve trailer
	trailer = stream.Trailer()


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

	res, err = stream.CloseAndRecv()
	if err != nil {
		panic(err)
	}


	log.Printf("##UpdateOrders## header: %v. trailer: %v", header, trailer)

	log.Printf("Update Orders Res : %s", res)
}
