package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
)

func orderUnaryClientInterceptor(ctx context.Context, method string, req, reply interface{},
	cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	// Pre-processor phase
	s := time.Now()

	// Invoking the remote method
	err := invoker(ctx, method, req, reply, cc, opts...)

	// Post-processor phase
	log.Printf("method: %s, req: %s, resp: %s, latency: %s\n",
		method, req, reply, time.Now().Sub(s))

	return err
}

// SendMsg method call.
type wrappedStream struct {
	method string
	grpc.ClientStream
}

func (w *wrappedStream) RecvMsg(m interface{}) error {
	err := w.ClientStream.RecvMsg(m)

	log.Printf("method: %s, res: %s\n", w.method, m)

	return err
}

func (w *wrappedStream) SendMsg(m interface{}) error {
	err := w.ClientStream.SendMsg(m)

	log.Printf("method: %s, req: %s\n", w.method, m)

	return err
}

func newWrappedStream(method string, s grpc.ClientStream) *wrappedStream {
	return &wrappedStream{
		method,
		s,
	}
}

func orderStreamClientInterceptor(ctx context.Context, desc *grpc.StreamDesc,
	cc *grpc.ClientConn, method string, streamer grpc.Streamer,
	opts ...grpc.CallOption) (grpc.ClientStream, error) {

	// Pre-processing logic
	s := time.Now()

	cs, err := streamer(ctx, desc, cc, method, opts...)

	// Post processing logic
	log.Printf("method: %s, latency: %s\n", method, time.Now().Sub(s))

	return newWrappedStream(method, cs), err
}
