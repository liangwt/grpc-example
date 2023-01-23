package main

import (
	"context"
	"fmt"
	"net"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	pb "github.com/liangwt/note/grpc/authentication/ecommerce"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var (
	errMissingMetadata = status.Errorf(codes.InvalidArgument, "missing metadata")
)

func ensureValidBasicCredentials(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errMissingMetadata
	}

	tokenString := strings.TrimPrefix(md["authorization"][0], "Bearer ")

	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("154a8b3aa89d3d4c49826f6dbbbe5542b5a9fbbb"), nil
	})

	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok || !token.Valid {
		return nil, status.Errorf(codes.Unauthenticated, err.Error())
	}

	fmt.Println(claims.ID)

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
