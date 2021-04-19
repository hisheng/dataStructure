/**
@author 张海生<zhanghaisheng@qimao.com>
@dateTime   : 2021/4/15 7:47 下午
@Desc
*/
package mutex

import (
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
