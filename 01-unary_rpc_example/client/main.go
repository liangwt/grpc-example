package main

import (
	"context"
	"log"
	"time"

	pb "github.com/liangwt/note/grpc/unary_rpc_example/ecommerce"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8009", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := pb.NewOrderManagementClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Add Order
	order := pb.Order{Id: "101", Items: []string{"iPhone XS", "Mac Book Pro"}, Destination: "San Jose, CA", Price: 2300.00}
	res, err := client.AddOrder(ctx, &order)
	if err != nil {
		panic(err)
	}

	log.Print("AddOrder Response -> ", res.Value)

	// Get Order
	retrievedOrder, err := client.GetOrder(ctx, &wrapperspb.StringValue{Value: "101"})
	if err != nil {
		panic(err)
	}

	log.Print("GetOrder Response -> : ", retrievedOrder)
}
