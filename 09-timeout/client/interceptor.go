package main

import (
	"context"
	"time"

	"google.golang.org/grpc"
)

func unaryClientInterceptor(ctx context.Context, method string, req, reply interface{},
	cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Invoking the remote method
	err := invoker(ctx, method, req, reply, cc, opts...)

	return err
}
