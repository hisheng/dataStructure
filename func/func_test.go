package tfunc

import (
	"testing"
)

type Options struct {
	Nodes    []string
	Database string
	Table    string
}

type Option func(o *Options)

// 设置 table 方法1
func (o *Options) SetTable(t string) {
	o.Table = t
}

// 设置 database 方法1
func (o *Options) SetDatabase(d string) {
	o.Database = d
}

// 设置 table 方法2
func Table(t string) Option {
	return func(o *Options) {
		o.Table = t
	}
}

// 设置 database 方法2
func Database(d string) Option {
	return func(o *Options) {
		o.Table = d
	}
}

type Store interface {
	Write(Options, ...Option)
}

func TestFunc(t *testing.T) {
	//设置table 方法1
	ops := new(Options)
	ops.SetTable("t_name")
	ops.SetDatabase("d")
	//设置table 方法2 有啥好处?
	Table("t_name")
	Database("ss")
	//Write(ops, Table("t"), Database("d")) 封装属性，如果有一些属性很复杂的话，但为什么不用第一种方法？
}
