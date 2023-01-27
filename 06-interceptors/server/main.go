package main

import (
	"net"

	pb "github.com/liangwt/note/grpc/bidirectional_streaming_rpc_example/ecommerce"
	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			orderUnaryServerInterceptor,
		),
		grpc.ChainStreamInterceptor(
			orderStreamServerInterceptor,
		),
	)

	pb.RegisterOrderManagementServer(s, &OrderManagementImpl{})

	lit, err := net.Listen("tcp", ":8009")
	if err != nil {
		panic(err)
	}

	if err := s.Serve(lit); err != nil {
		panic(err)
	}
}
