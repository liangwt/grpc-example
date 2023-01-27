package main

import (
	"context"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	pb "github.com/liangwt/note/grpc/ecosystem/grpc-gateway/ecommerce"
	"google.golang.org/grpc"
)

func main() {
	l, err := net.Listen("tcp", ":8009")
	if err != nil {
		panic(err)
	}

	go func() {
		s := grpc.NewServer()
		pb.RegisterOrderManagementServer(s, &server{})
		if err := s.Serve(l); err != nil {
			panic(err)
		}
	}()

	// // This is where the gRPC-Gateway proxies the requests
	// conn, err := grpc.DialContext(
	// 	context.Background(),
	// 	"0.0.0.0:8009",
	// 	grpc.WithBlock(),
	// 	grpc.WithInsecure(),
	// )

	// gwmux := runtime.NewServeMux()

	// err = pb.RegisterOrderManagementHandler(context.Background(), gwmux, conn)
	// if err != nil {
	// 	panic(err)
	// }

	mux := http.NewServeMux()
	
	{
		gwmux := runtime.NewServeMux()

		err = pb.RegisterOrderManagementHandlerServer(context.Background(), gwmux, &server{})
		if err != nil {
			panic(err)
		}
	
		mux.Handle("/", gwmux)
	}

	{
		
	}
	

	gwServer := &http.Server{
		Addr:    ":8010",
		Handler: mux,
	}

	if err := gwServer.ListenAndServe(); err != nil {
		panic(err)
	}
}
