/**
@author 张海生<zhanghaisheng@qimao.com>
@dateTime   : 2021/1/21 7:58 下午
@Desc
*/
package tfunc

import (
	"fmt"
	"testing"
	"time"
)

func TestR(t *testing.T) {

	s := testDefer()
	fmt.Println(s)
}

func testDefer() string {
	defer dd()
	fmt.Println("aaa")
	return "vvvvv"
}
func dd() {
	fmt.Println(112)
}

func timeSpent(in func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := in(n)
		fmt.Println("s", time.Since(start).Seconds())
		return ret
	}
}
