package main

import (
	"context"
	"io"
	"log"

	pb "github.com/liangwt/note/grpc/server_streaming_rpc_example/ecommerce"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func main(){
	conn, err := grpc.Dial("127.0.0.1:8009", grpc.WithInsecure())
	if err != nil{
		panic(err)
	}

	c := pb.NewOrderManagementClient(conn)
	ctx, cancelFn := context.WithCancel(context.Background())
	defer cancelFn()

	stream, err := c.SearchOrders(ctx, &wrapperspb.StringValue{Value: "Google"})
	if err != nil{
		panic(err)
	}

	for{
		order, err := stream.Recv()
		if err == io.EOF{
			break
		}

		log.Println("Search Result: ", order)
	}
}