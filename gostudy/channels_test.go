/**
@author 张海生<zhanghaisheng@qimao.com>
@dateTime   : 2021/4/16 3:24 下午
@Desc
*/
package gostudy

import (
	"fmt"
	"testing"
)

func TestChannel(t *testing.T) {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // 从 c 中接收

	fmt.Println(x, y, x+y)
	// fmt.Println(<-c)  fatal error: all goroutines are asleep - deadlock! 没有值了引发panic
}

func TestBufChannel(t *testing.T) {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	//用select 可以防止  fmt.Println(<-ch) 包错误信息
	select {
	case i := <-ch:
		t.Log(i)
		// 使用 i
	default:
		t.Log("五数据")
		// 从 c 中接收会阻塞时执行
	}

	// 不close 的话，下面再请求会panic
	close(ch)
	v, ok := <-ch
	fmt.Println(v, ok)
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}
