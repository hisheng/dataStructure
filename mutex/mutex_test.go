/**
@author 张海生<zhanghaisheng@qimao.com>
@dateTime   : 2021/4/15 7:47 下午
@Desc
*/
package mutex

import (
	"fmt"
	"sync"
	"testing"
)

func TestMutex(t *testing.T) {
	//var mu sync.Mutex
	var count = 0

	var wg sync.WaitGroup

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 10000; j++ {
				count++
			}
		}()
	}
	wg.Wait()
	t.Log(count)
}

func TestNew(t *testing.T) {
	//1 互斥锁的机制  临界区，临界区就是一个被共享的资源，这部分被保护起来。
	mutex := sync.Mutex{} //{0 0}
	t.Log(mutex)

	var mu sync.Mutex //{0 0}
	t.Log(mu)

	m := new(sync.Mutex) //&{0 0}
	t.Log(m)
}

type Counter struct {
	Count int
	sync.Mutex
}

func TestCount(t *testing.T) {
	var c Counter
	c.Lock()
	defer c.Unlock()
	c.Count++
	foo(&c) // 复制锁 error
}

// 为啥没法复制，但 官方 TestMutex 有这个传引用的例子
func TestMutexCopy(t *testing.T) {
	var mu sync.Mutex
	mu.Lock()
	defer mu.Unlock() //用defer 会报错，是因为foo2 函数的时候，这个锁，还么有Unlock ，此时再 Lock就出错了
	foo2(&mu)
}

func foo2(c *sync.Mutex) {
	c.Lock()
	defer c.Unlock()
	fmt.Println("in foo2")
}

// 这里Counter的参数是通过复制的方式传入的
func foo(c *Counter) {
	c.Lock()
	defer c.Unlock()
	fmt.Println("in foo")
}
