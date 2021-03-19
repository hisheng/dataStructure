/**
@author 张海生<zhanghaisheng@qimao.com>
@dateTime   : 2021/3/11 7:20 下午
@Desc
*/
package once

import (
	"fmt"
	"runtime"
	"strconv"
	"sync"
	"testing"
	"time"
	"unsafe"
)

type Singleton struct {
}

var singleton *Singleton
var once sync.Once

func GetObj() *Singleton {
	once.Do(func() {
		fmt.Println("create obj")
		singleton = new(Singleton)
	})
	return singleton
}

func TestOnce(t *testing.T) {
	var once1 sync.Once
	var once2 sync.Once
	fun1 := func() {
		t.Log("第一次打印")
	}
	once1.Do(fun1)

	fun2 := func() {
		t.Log("第二次打印")
	}
	once1.Do(fun2)
	//多个once 是不是互相影响----不同对象互相独立的
	t.Log("once2")
	once2.Do(fun1)
	once2.Do(fun2)

	//3 并发情况下，是否原子安全 --结论 开启了 5 个并发的 goroutine ，不管你咋么运行，始终只打印一次，至于 i 是多少，就看先执行的是哪个 g 了。
	t.Log("once3")
	var once3 sync.Once
	for i := 0; i < 5; i++ {
		go func(i int) {
			once3.Do(fun1)
			once2.Do(fun2)
		}(i)
	}

	time.Sleep(time.Second * 1)
}

func runTesk(id int) string {
	time.Sleep(10 * time.Millisecond)
	return strconv.Itoa(id)
}

func FirstResponse() string {
	num := 10
	ch := make(chan string, num)
	for i := 0; i < num; i++ {
		go func(i int) {
			ret := runTesk(i)
			ch <- ret
		}(i)
	}
	return <-ch
}

func AllResponse() string {
	num := 10
	//多个协程并发取
	ch := make(chan string, num)
	for i := 0; i < num; i++ {
		go func(i int) {
			ret := runTesk(i)
			ch <- ret
		}(i)
	}

	//一次阻塞读
	all := ""
	for i := 0; i < num; i++ {
		all += <-ch
	}
	return all
}

func TestFirstResponse(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(FirstResponse())
	}
}

func TestName(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			obj := GetObj()
			t.Logf("%x", obj)
			t.Log(unsafe.Pointer(obj)) //输出对象的地址
			t.Log(obj)
			wg.Done()
		}()
	}
	t.Log(runtime.NumGoroutine())
	wg.Wait()
}
