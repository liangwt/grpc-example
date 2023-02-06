package main

import (
	"context"
	"log"
	"time"

	pb "github.com/liangwt/note/grpc/error_handling/ecommerce"
	epb "google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8009", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := pb.NewOrderManagementClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Get Order
	order, err := client.GetOrder(ctx, &wrapperspb.StringValue{Value: ""})
	if err != nil {
		st, ok := status.FromError(err)
		if !ok {
			log.Println(err)
			return
		}

		switch st.Code() {
		case codes.InvalidArgument:
			for _, d := range st.Details() {
				switch info := d.(type) {
				case *epb.BadRequest_FieldViolation:
					log.Printf("Request Field Invalid: %s", info)
				default:
					log.Printf("Unexpected error type: %s", info)
				}
			}
		default:
			log.Printf("Unhandled error : %s ", st.String())
		}

		return
	}

	log.Print("GetOrder Response -> : ", order)
}
