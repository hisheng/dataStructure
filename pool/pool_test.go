/**
@author 张海生<zhanghaisheng@qimao.com>
@dateTime   : 2021/3/18 7:34 下午
@Desc
*/
package pool

import (
	"runtime"
	"sync"
	"testing"
)

func TestPool(t *testing.T) {
	//sync pool gc 后会 释放
	pool := sync.Pool{
		New: func() interface{} {
			return 100
		},
	}

	v := pool.Get().(int)
	t.Log(v)
	pool.Put(3)
	runtime.GC()
	t.Log(pool.Get())
}

type obj struct {
}

type objPool struct {
}

//使用 buffered channel 实现对象池
// https://github.com/search?l=Go&q=%E5%AF%B9%E8%B1%A1%E6%B1%A0&type=Repositories
func NewObjPool() chan *obj {
	num := 10
	pool := make(chan *obj, num)
	for i := 0; i < num; i++ {
		pool <- &obj{}
	}
	return pool
}
