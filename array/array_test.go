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

	y := [...]int{1, 2, 3}
	t.Log(y)

	//t.FailNow() t.Fatal()  失败了， 停止本个测试，其他的测试继续执行
	//t.Fail() t.Error()  marks the function as having failed but continues execution.
}

/**
 * 测试多维数组
 *
 * @author zhanghaisheng@qimao.com
 * @dateTime   : 2021/2/7 4:09 下午
 */
func TestArrayArray(t *testing.T) {
	x := 2
	y := 4

	table := make([][]int, x)
	for i := range table {
		table[i] = make([]int, y)
	}
	t.Log(table) //[[0 0 0 0] [0 0 0 0]]
}

func TestArr(t *testing.T) {
	x := []int{1, 2}
	t.Log(x) //[1 2]
}
