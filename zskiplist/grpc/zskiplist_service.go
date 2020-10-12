package grpc

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewZskiplistService() GreeterServer {
	return &ZskiplistService{}
}

type ZskiplistService struct {
	UnimplementedGreeterServer
}

func (*ZskiplistService) SayHello(ctx context.Context, req *HelloRequest) (*HelloReply, error) {
	fmt.Println("ssss ss")
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
