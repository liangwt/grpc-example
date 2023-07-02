package main

import (
	"context"
	"encoding/base64"
	"log"
	"time"

	pb "github.com/liangwt/note/grpc/authentication/ecommerce"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

var _ credentials.PerRPCCredentials = BasicAuthentication{}

type BasicAuthentication struct {
	password string
	username string
}

func (b BasicAuthentication) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	auth := b.username + ":" + b.password
	enc := base64.StdEncoding.EncodeToString([]byte(auth))

	return map[string]string{
		"authorization": "Basic " + enc,
	}, nil
}

func (b BasicAuthentication) RequireTransportSecurity() bool {
	return true
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	auth := BasicAuthentication{
		username: "admin",
		password: "admin",
	}

	creds, err := credentials.NewClientTLSFromFile("./x509/rootCa.crt", "www.example.com")
	if err != nil {
		panic(err)
	}

	conn, err := grpc.Dial("localhost:8009",
		grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(auth))
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
