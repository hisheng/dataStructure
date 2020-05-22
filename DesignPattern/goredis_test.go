package DesignPattern

import (
	"testing"
)

type Client struct {
	cmdable
}

type cmdable func(cmd Cmder) error

func (c cmdable) Incr(key string) *IntCmd {
	cmd := NewIntCmd("incr", key)
	_ = c(cmd)
	return cmd
}

//cmd error 格式化error,让所有接口有 err()与SetErr()方法
type Cmder interface {
	SetErr(error)
	Err() error
}

type baseCmd struct {
	err  error
	args []interface{}
}

func (cmd *baseCmd) SetErr(e error) {
	cmd.err = e
}

func (cmd *baseCmd) Err() error {
	return cmd.err
}

type IntCmd struct {
	baseCmd
	val int64
}

func NewIntCmd(args ...interface{}) *IntCmd {
	return &IntCmd{
		baseCmd: baseCmd{args: args},
	}
}

func (cmd *IntCmd) Val() int64 {
	return cmd.val
}

func TestRedis(t *testing.T) {
	rdb := Client{}
	err := rdb.Incr("key").Err()
	if err != nil {
		panic(err)
	}

	rdb.Incr("key").Val()
}
