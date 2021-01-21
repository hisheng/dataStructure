/**
@author 张海生<zhanghaisheng@qimao.com>
@dateTime   : 2021/1/21 10:31 上午
@Desc
*/
package slice

import (
	"testing"
	"time"
)

func TestArray(t *testing.T) {
	a := [2]int{1, 2}
	t.Log(a, len(a), cap(a)) //[1 2] 2 2

	b := []int{1, 2}
	t.Log(b, len(b), cap(b)) //[1 2] 2 2

	//在append 的时候才会扩容,不然和数组是是一样的 ?

	t1 := time.Now()
	var a1 [2]int
	for i := 0; i < 1*100000*1000; i++ {
		a1 = [2]int{1, 2}
	}
	t.Log("a1 创建速度：" + time.Now().Sub(t1).String()) //a1 创建速度：47.415453ms
	t.Log(a1)

	t2 := time.Now()
	var b1 []int
	for i := 0; i < 1*100000*1000; i++ {
		b1 = []int{1, 2}
	}
	t.Log("b1 创建速度: " + time.Now().Sub(t2).String()) //b1 创建速度: 1.807974945s
	t.Log(b1)

	t3 := time.Now()
	for i := 0; i < 1*100000*1000; i++ {
		for _, _ = range a1 {
		}
	}
	t.Log("a1遍历:" + time.Now().Sub(t3).String()) //a1遍历:97.244402ms

	t4 := time.Now()
	for i := 0; i < 1*100000*1000; i++ {
		for _, _ = range b1 {
		}
	}
	t.Log("b1遍历:" + time.Now().Sub(t4).String()) //b1遍历:82.392376ms
}
