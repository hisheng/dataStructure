/**
@author 张海生<zhanghaisheng@qimao.com>
@dateTime   : 2021/2/25 7:45 下午
@Desc
*/
package gotoutine

import (
	"sync"
	"testing"
	"time"
)

func TestGroutine(t *testing.T) {
	for i := 0; i < 20; i++ {
		go func(i int) {
			t.Log(i)
		}(i)
	}
	time.Sleep(time.Second * 2)
}

func TestLock(t *testing.T) {
	///并发了，共享内存的问题，count
	count := 0
	for i := 0; i < 5000; i++ {
		go func() {
			count++
		}()
	}
	t.Log(count)
	time.Sleep(time.Second * 2)
}

//代替sleep 让所有要等待的任务执行好
func TestLock2(t *testing.T) {
	///并发了，共享内存的问题，count
	var wg sync.WaitGroup
	count := 0
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			count++
			wg.Done()
		}()
	}
	wg.Wait()
	t.Log(count)
}

func TestLock3(t *testing.T) {
	///并发了，共享内存的问题，count
	var wg sync.WaitGroup
	var lock sync.Mutex
	count := 0
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			defer lock.Unlock()
			lock.Lock()
			count++
			wg.Done()
		}()
	}
	wg.Wait()
	t.Log(count)
}
