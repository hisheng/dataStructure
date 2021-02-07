/**
@author 张海生<zhanghaisheng@qimao.com>
@dateTime   : 2021/2/7 3:52 下午
@Desc
*/
package array

import "testing"

func TestArray(t *testing.T) {
	x := [3]int{1, 2, 3}

	func(arr [3]int) {
		arr[0] = 7
		t.Log(arr) //prints [7 2 3]
	}(x)

	t.Log(x) //prints [1 2 3] (not ok if you need [7 2 3])

	//如果是 slice 的话，会一样，array是不一样的
}
