package grpc

import (
	"context"
	"fmt"
	"strconv"
	"zskiplist"
)

func NewZskiplistService() ZskiplistServer {
	return &ZskiplistService{
		zsl: zskiplist.ZslCreate(),
	}
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
	score, err := strconv.ParseFloat(req.Score, 64)
	if err != nil {
		return nil, err
	}
	zskiplist.ZslInsert(zsp.zsl, score, req.Ele)
	var list []*Zskip
	x := zsp.zsl.Tail
	for x.Backward != nil {
		z := &Zskip{
			Ele:   x.Ele,
			Score: strconv.FormatFloat(x.Score, 'E', -1, 64),
		}
		list = append(list, z)
		x = x.Backward
	}

	reply := &ZaddReply{
		List: list,
	}
	return reply, nil
}
func (zsp *ZskiplistService) List(ctx context.Context, req *ListRequest) (*ListReply, error) {
	var list []*Zskip
	x := zsp.zsl.Tail
	for x.Backward != nil {
		z := &Zskip{
			Ele:   x.Ele,
			Score: strconv.FormatFloat(x.Score, 'E', -1, 64),
		}
		list = append(list, z)
		x = x.Backward
	}

	reply := &ListReply{
		List: list,
	}
	return reply, nil
}
