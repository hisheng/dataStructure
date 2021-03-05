/**
@author 张海生<zhanghaisheng@qimao.com>
@dateTime   : 2021/3/4 7:22 下午
@Desc
*/
package async

import (
	"log"
	"testing"
	"time"
)

func TestAsync(t *testing.T) {
	a := make(chan int, 10)
	go func() {
		for {
			select {
			case i, ok := <-a:
				time.Sleep(time.Millisecond * 500)
				log.Println(i, ok)
				if !ok {
					return
				}
			}
		}
	}()
	go func() {
		for i := 0; i < 10; i++ {
			a <- i
		}
		close(a) //a = nil 这种情况只能收到1个，nil的话，就gc回收了。
	}()
	time.Sleep(time.Second * 10)
}
