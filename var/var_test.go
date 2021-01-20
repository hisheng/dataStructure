/**
@author 张海生<zhanghaisheng@qimao.com>
@dateTime   : 2021/1/20 10:35 上午
@Desc
*/
package _var

import "testing"

type persion struct {
	name string
	age  int
}

func TestVar(t *testing.T) {
	t.Log("test var ")
	var p persion
	if &p == nil {
		t.Log("p == nil")
	} else {
		t.Log("p != nil")
	}

	t.Log(p) // 注意 p 在var的时候已经分配了 地址空间 所以不是nil，只有指针类型的，比如slice，map才会nil

	var m map[int]string
	if m == nil {
		t.Log("m == nil")
	} else {
		t.Log("m != nil")
	}
	t.Log(m)

	var m2 map[int]string = map[int]string{}
	if m2 == nil {
		t.Log("m2 == nil")
	} else {
		t.Log("m2 != nil")
	}
	t.Log(m2)

}
