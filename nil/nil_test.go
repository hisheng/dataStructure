/**
@author 张海生<zhanghaisheng@qimao.com>
@dateTime   : 2021/2/7 3:34 下午
@Desc
*/
package nil

import "testing"

func TestNilValue(t *testing.T) {
	t.Log("nil")
	var s []int      //nil
	s = append(s, 1) //[1]
	t.Log(s)

	//var m map[string]int //nil
	//m["age"] = 1   这样写会有error 正确写法是下面
	m := make(map[string]int, 99)
	m["age"] = 1
	t.Log(m) //map[age:1]

	var n interface{}
	if n == nil {
		t.Log("n == nil") //n == nil
	}
	t.Log(n) //<nil>
}
