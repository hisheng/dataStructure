package grpc

import (
	"context"
	"fmt"
	"zskiplist"
)

func NewZskiplistService() ZskiplistServer {
	return &ZskiplistService{}
}

type ZskiplistService struct {
	zsl *zskiplist.Zskiplist
	UnimplementedZskiplistServer
}

func (zsp *ZskiplistService) SayHello(ctx context.Context, req *HelloRequest) (*HelloReply, error) {
	fmt.Println("ssss ss")
	return &HelloReply{
		Message: req.Name,
	}, nil
}

func (zsp *ZskiplistService) Zadd(ctx context.Context, req *ZaddRequest) (*ZaddReply, error) {
	return nil, nil
}
func (zsp *ZskiplistService) List(ctx context.Context, req *ListRequest) (*ListReply, error) {
	return nil, nil
}
