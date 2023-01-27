package main

import (
	"context"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
	pb "github.com/liangwt/note/grpc/authentication/ecommerce"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

var _ credentials.PerRPCCredentials = JwtAuthentication{}

type JwtAuthentication struct {
	Key []byte
}

func (a JwtAuthentication) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ID:        "example",
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(-2 * time.Hour)),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(a.Key)
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"authorization": "Bearer " + tokenString,
	}, nil
}

func (b JwtAuthentication) RequireTransportSecurity() bool {
	return true
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	creds, err := credentials.NewClientTLSFromFile("./x509/rootCa.crt", "www.example.com")
	if err != nil {
		panic(err)
	}

	jwtAuth := JwtAuthentication{[]byte("154a8b3aa89d3d4c49826f6dbbbe5542b5a9fbbb")}

	conn, err := grpc.Dial("localhost:8009",
		grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(jwtAuth))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := pb.NewOrderManagementClient(conn)

	// Get Order
	retrievedOrder, err := client.GetOrder(ctx, &wrapperspb.StringValue{Value: "101"})
	if err != nil {
		panic(err)
	}

	log.Printf("GetOrder Response -> : %+v\n", retrievedOrder)
}
