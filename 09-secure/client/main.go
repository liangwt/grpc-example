package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"time"

	pb "github.com/liangwt/note/grpc/secure/ecommerce"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func main() {
	// 加载客户端证书
	certificate, err := tls.LoadX509KeyPair("x509/client.crt", "x509/client.key")
	if err != nil {
		log.Fatal(err)
	}

	// 构建CertPool以校验服务端证书有效性
	b, err := ioutil.ReadFile("./x509/rootCa.crt")
	if err != nil {
		log.Fatal(err)
	}
	cp := x509.NewCertPool()
	if !cp.AppendCertsFromPEM(b) {
		log.Fatal("credentials: failed to append certificates")
	}

	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{certificate},
		ServerName:   "www.example.com",
		RootCAs:      cp,
	})

	conn, err := grpc.Dial("localhost:8009", grpc.WithTransportCredentials(creds))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := pb.NewOrderManagementClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Get Order
	retrievedOrder, err := client.GetOrder(ctx, &wrapperspb.StringValue{Value: "101"})
	if err != nil {
		panic(err)
	}

	log.Print("GetOrder Response -> : ", retrievedOrder)
}
