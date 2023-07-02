package main

import (
	"context"
	"flag"
	"log"

	pb "github.com/liangwt/note/grpc/name_resolver_lb_example/ecommerce"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	xdscreds "google.golang.org/grpc/credentials/xds"

	_ "google.golang.org/grpc/xds" // To install the xds resolvers and balancers.
)

var (
	xdsCreds = flag.Bool("xds_creds", false, "whether the server should use xDS APIs to receive security configuration")
)

func main() {
	flag.Parse()
	
	creds := insecure.NewCredentials()
	// xds api也可以传输TLS证书
	if *xdsCreds {
		log.Println("Using xDS credentials...")
		var err error
		if creds, err = xdscreds.NewClientCredentials(xdscreds.ClientOptions{FallbackCreds: insecure.NewCredentials()}); err != nil {
			panic(err)
		}
	}

	conn, err := grpc.Dial("xds:///localhost:50051", grpc.WithTransportCredentials(creds))
	if err != nil {
		panic(err)
	}

	ctx, cancelFn := context.WithCancel(context.Background())
	defer cancelFn()

	c := pb.NewOrderManagementClient(conn)

	for i := 0; i < 5; i++ {
		// Add Order
		order := pb.Order{Id: "101", Items: []string{"iPhone XS", "Mac Book Pro"}, Destination: "San Jose, CA", Price: 2300.00}
		res, err := c.AddOrder(ctx, &order)
		if err != nil {
			panic(err)
		}

		log.Printf("Add Orders Res : %s", res)
	}
}