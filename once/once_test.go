/**
@author 张海生<zhanghaisheng@qimao.com>
@dateTime   : 2021/3/11 7:20 下午
@Desc
*/
package once

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
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

func TestName(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			obj := GetObj()
			t.Log(obj)
			wg.Done()
		}()
	}
	t.Log(runtime.NumGoroutine())
	wg.Wait()
}
