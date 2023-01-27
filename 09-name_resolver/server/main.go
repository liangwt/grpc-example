package main

import (
	"net"

	pb "github.com/liangwt/note/grpc/name_resolver_lb_example/ecommerce"
	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()
	pb.RegisterOrderManagementServer(s, &server{})

	lit, err := net.Listen("tcp", ":8010")
	if err != nil {
		panic(err)
	}

	if err := s.Serve(lit); err != nil {
		panic(err)
	}
}
