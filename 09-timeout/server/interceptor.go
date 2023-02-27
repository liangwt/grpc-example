package main

import (
	"context"
	"time"

	"google.golang.org/grpc"
)

func unaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Invoking the handler to complete the normal execution of a unary RPC.
	m, err := handler(ctx, req)

	return m, err
}
