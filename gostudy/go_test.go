/**
@author 张海生<zhanghaisheng@qimao.com>
@dateTime   : 2021/1/7 7:05 下午
@Desc
*/
package gostudy

import "testing"

func TestGo(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch i {
		case 0, 2:
			t.Log("even")
		case 1, 3:
			t.Log("odd")
		default:
			t.Log("iii")
		}
	}

	var a [3]int //len(1) = 3 此时
	if len(a) == 0 {
		t.Log("a == nil")
	} else {
		t.Log("a != 0")
	}
	t.Log(a)
	var b []int
	if b == nil {
		t.Log("b ==nil")
	}
	t.Log(b)

	s2 := make([]int, 3, 5)
	t.Log(s2[0])

	t.Log("---")
	var s []int
	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s))
	}

	year := []string{"1", "2", "3"}
	aa := year
	aa[0] = "0"
	t.Log(aa, year)

	y := map[int]string{0: "0", 1: "1"}
	aaa := y
	aaa[0] = "00"
	t.Log(aaa, y)

}
