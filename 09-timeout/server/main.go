package main

import (
	"net"
	"time"

	pb "github.com/liangwt/note/grpc/client_streaming_rpc_example/ecommerce"
	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer(
		grpc.ConnectionTimeout(3*time.Second),
		grpc.UnaryInterceptor(unaryServerInterceptor),
	)
	pb.RegisterOrderManagementServer(s, &server{})

	lit, err := net.Listen("tcp", ":8009")
	if err != nil {
		panic(err)
	}

	if err := s.Serve(lit); err != nil {
		panic(err)
	}
}
