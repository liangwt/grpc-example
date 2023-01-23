package main

import (
	"fmt"
	"log"

	pb "github.com/liangwt/note/grpc/error_handling/error"
	epb "google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	// 发送错误响应
	err := InvokRPC2()

	// 错误转回 status
	// 转换有可能失败
	st, ok := status.FromError(err)
	if ok {
		fmt.Println(st.Code(), st.Message())
	}

	if st.Code() == codes.InvalidArgument {
		for _, d := range st.Details() {
			switch info := d.(type) {
			case *epb.BadRequest_FieldViolation:
				log.Printf("Request Field Invalid: %s", info)
			default:
				log.Printf("Unexpected error type: %s", info)
			}
		}
	} else {
		log.Printf("Unhandled error : %s ", st.String())
	}
}

func InvokRPC1() error {
	st := status.New(codes.InvalidArgument, "invalid args")

	if details, err := st.WithDetails(&pb.BizError{}); err == nil {
		return details.Err()
	}

	return st.Err()
}

func InvokRPC2() error {
	st := status.New(codes.InvalidArgument, "invalid args")

	if details, err := st.WithDetails(
		&epb.BadRequest_FieldViolation{
			Field:       "ID",
			Description: fmt.Sprintf("Order ID received is not valid"),
		},
	); err == nil {
		return details.Err()
	}

	return st.Err()
}
