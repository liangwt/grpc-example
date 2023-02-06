package main

import (
	"errors"
	"fmt"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Invoke2() {
	ok := status.New(codes.OK, "ok")
	fmt.Println(ok)

	invalidArgument := status.New(codes.InvalidArgument, "invalid args")
	fmt.Println(invalidArgument)
}

var (
	ParamsErr = errors.New("params err")
	BizErr    = errors.New("biz err")
)

func Invoke(i bool) error {
	if i {
		return ParamsErr
	} else {
		return BizErr
	}
}

func main() {
	err := Invoke(true)

	if err != nil {
		switch {
		case errors.Is(err, ParamsErr):
			log.Println("params error")
		case errors.Is(err, BizErr):
			log.Println("biz error")
		}
	}
}
