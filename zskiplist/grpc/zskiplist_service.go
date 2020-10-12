package grpc

import (
	"context"
	"fmt"
)

func NewZskiplistService() ZskiplistServer {
	return &ZskiplistService{}
}

type ZskiplistService struct {
	UnimplementedZskiplistServer
}

func (zsp *ZskiplistService) SayHello(ctx context.Context, req *HelloRequest) (*HelloReply, error) {
	fmt.Println("ssss ss")
	return &HelloReply{
		Message: req.Name,
	},nil
}
