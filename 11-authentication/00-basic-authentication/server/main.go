package main

import (
	"context"
	"encoding/base64"
	"net"
	"strings"

	pb "github.com/liangwt/note/grpc/authentication/ecommerce"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var (
	errMissingMetadata = status.Errorf(codes.InvalidArgument, "missing metadata")
	errInvalidToken    = status.Errorf(codes.Unauthenticated, "invalid credentials")
)

func ensureValidBasicCredentials(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errMissingMetadata
	}

	authorization := md["authorization"]

	if len(authorization) < 1 {
		return nil, errInvalidToken
	}

	token := strings.TrimPrefix(authorization[0], "Basic ")
	if token != base64.StdEncoding.EncodeToString([]byte("admin:admin")) {
		return nil, errInvalidToken
	}

	return handler(ctx, req)
}

func main() {
	l, err := net.Listen("tcp", ":8009")
	if err != nil {
		panic(err)
	}

	creds, err := credentials.NewServerTLSFromFile("./x509/server.crt", "./x509/server.key")
	s := grpc.NewServer(
		grpc.UnaryInterceptor(ensureValidBasicCredentials),
		grpc.Creds(creds),
	)

	pb.RegisterOrderManagementServer(s, &server{})

	if err := s.Serve(l); err != nil {
		panic(err)
	}
}
