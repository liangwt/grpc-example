package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
)

func orderUnaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// Pre-processing logic
	s := time.Now()

	// Invoking the handler to complete the normal execution of a unary RPC.
	m, err := handler(ctx, req)

	// Post processing logic
	log.Printf("method: %s, req: %s, resp: %s, latency: %s\n",
		info.FullMethod, req, m, time.Now().Sub(s))

	return m, err
}

// SendMsg method call.
type wrappedStream struct {
	Recv []interface{}
	Send []interface{}
	grpc.ServerStream
}

func (w *wrappedStream) RecvMsg(m interface{}) error {
	err := w.ServerStream.RecvMsg(m)

	w.Recv = append(w.Recv, m)

	return err
}

func (w *wrappedStream) SendMsg(m interface{}) error {
	err := w.ServerStream.SendMsg(m)

	w.Send = append(w.Send, m)

	return err
}

func newWrappedStream(s grpc.ServerStream) *wrappedStream {
	return &wrappedStream{
		make([]interface{}, 0),
		make([]interface{}, 0),
		s,
	}
}

func orderStreamServerInterceptor(srv interface{},
	ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {

	// Pre-processing logic
	s := time.Now()

	// Invoking the StreamHandler to complete the execution of RPC invocation
	nss := newWrappedStream(ss)
	err := handler(srv, nss)

	// Post processing logic
	log.Printf("method: %s, req: %+v, resp: %+v, latency: %s\n",
		info.FullMethod, nss.Recv, nss.Send, time.Now().Sub(s))

	return err
}
