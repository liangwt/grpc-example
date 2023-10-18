package main

import (
	"context"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	pb "github.com/liangwt/note/grpc/ecosystem/grpc-gateway/ecommerce"
)

func main() {
	gwmux := runtime.NewServeMux()

	err := pb.RegisterOrderManagementHandlerServer(context.Background(), gwmux, &OrderManagementImpl{})
	if err != nil {
		panic(err)
	}

	err = gwmux.HandlePath("GET", "/hello/{name}", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		w.Write([]byte("hello " + pathParams["name"]))
	})

	http.ListenAndServe(":8010", gwmux)
}
