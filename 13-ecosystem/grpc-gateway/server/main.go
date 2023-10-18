package main

import (
	"context"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	pb "github.com/liangwt/note/grpc/ecosystem/grpc-gateway/ecommerce"
	"google.golang.org/grpc"
)

// curl -s -X POST \
// '127.0.0.1:8010/v1/addOrder' \
// --header 'Accept: */*' \
// --data '{"id": "102","items": ["Google","Baidu"],"description": "example","price": 0,"destination": "example"}'

func main() {
	grpcPort, gwPort := ":8009", ":8010"

	go func() {
		lis, err := net.Listen("tcp", grpcPort)
		if err != nil {
			panic(err)
		}

		s := grpc.NewServer()

		pb.RegisterOrderManagementServer(s, &OrderManagementImpl{})
		if err := s.Serve(lis); err != nil {
			panic(err)
		}
	}()

	// 建立一个到gRPC Port的连接
	conn, err := grpc.DialContext(
		context.Background(),
		"127.0.0.1"+grpcPort,
		grpc.WithBlock(),
		grpc.WithInsecure(),
	)
	if err != nil {
		panic(err)
	}

	gwmux := runtime.NewServeMux()
	err = pb.RegisterOrderManagementHandler(context.Background(), gwmux, conn)
	if err != nil {
		panic(err)
	}

	err = gwmux.HandlePath("GET", "/hello/{name}", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		w.Write([]byte("hello " + pathParams["name"]))
	})

	http.ListenAndServe(gwPort, gwmux)

	// 以下和http.ListenAndServe(gwPort, gwmux)等价

	// gwServer := &http.Server{
	// 	Addr:    gwPort,
	// 	Handler: gwmux,
	// }

	// if err := gwServer.ListenAndServe(); err != nil {
	// 	panic(err)
	// }
}
