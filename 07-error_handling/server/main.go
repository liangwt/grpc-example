package main

import (
	"net"

	pb "github.com/liangwt/note/grpc/error_handling/ecommerce"
	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()
	
	pb.RegisterOrderManagementServer(s, &OrderManagementImpl{})

	lis, err := net.Listen("tcp", ":8009")
	if err != nil {
		panic(err)
	}

	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
