package main

import (
	"context"
	"log"

	bl "github.com/liangwt/note/grpc/name_resolver_lb_example/client/balancer"
	rs "github.com/liangwt/note/grpc/name_resolver_lb_example/client/resolver"
	pb "github.com/liangwt/note/grpc/name_resolver_lb_example/ecommerce"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/resolver"
)

func main() {
	resolver.Register(rs.NewResolverBuilder(map[string][]string{
		"cluster@callee": {
			"127.0.0.1:8009",
			"127.0.0.1:8010",
		},
	}))

	balancer.Register(bl.NewBalancerBuilder())

	conn, err := grpc.Dial("example:cluster@callee",
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(
			`{"loadBalancingPolicy":"pick_first"}`,
		),
	)
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
