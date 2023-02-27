package main

import (
	"context"
	"log"
	"time"

	pb "github.com/liangwt/note/grpc/client_streaming_rpc_example/ecommerce"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 6*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, "127.0.0.1:8009",
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithUnaryInterceptor(unaryClientInterceptor),
	)
	if err != nil {
		if err == context.DeadlineExceeded {
			panic(err)
		}
		panic(err)
	}

	c := pb.NewOrderManagementClient(conn)

	ctx, cancel = context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Add Order
	order := pb.Order{
		Id:          "101",
		Items:       []string{"iPhone XS", "Mac Book Pro"},
		Destination: "San Jose, CA",
		Price:       2300.00,
	}
	res, err := c.AddOrder(ctx, &order)
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.DeadlineExceeded {
			panic(err)
		}
		panic(err)
	}

	log.Println(res)
}
