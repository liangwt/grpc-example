package main

import (
	// "crypto/tls"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net"

	pb "github.com/liangwt/note/grpc/secure/ecommerce"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	l, err := net.Listen("tcp", ":8009")
	if err != nil {
		panic(err)
	}

	certificate, err := tls.LoadX509KeyPair("./x509/server.crt", "./x509/server.key")
	if err != nil {
		panic(err)
	}

	// 创建CertPool，后续就用池里的证书来校验客户端证书有效性
	// 所以如果有多个客户端 可以给每个客户端使用不同的 CA 证书，来实现分别校验的目的
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("./x509/rootCa.crt")
	if err != nil {
		log.Fatal(err)
	}
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatal("failed to append certs")
	}

	// 构建基于 TLS 的 TransportCredentials
	creds := credentials.NewTLS(&tls.Config{
		// 设置证书链，允许包含一个或多个
		Certificates: []tls.Certificate{certificate},
		// 要求必须校验客户端的证书 可以根据实际情况选用其他参数
		ClientAuth: tls.RequireAndVerifyClientCert, // NOTE: this is optional!
		// 设置根证书的集合，校验方式使用 ClientAuth 中设定的模式
		ClientCAs: certPool,
	})

	s := grpc.NewServer(grpc.Creds(creds))

	pb.RegisterOrderManagementServer(s, &server{})

	if err := s.Serve(l); err != nil {
		panic(err)
	}
}
