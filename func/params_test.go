/**
Go语言参数传递 是否改变原来的值
@author 张海生<zhanghaisheng@qimao.com>
@dateTime   : 2021/1/6 2:46 下午
@Desc
*/
package tfunc

import (
	"fmt"
	"testing"
)

/**
@Desc int的值不变
@author zhanghaisheng@qimao.com
@dateTime   : 2021/1/6 3:14 下午
*/
func tInt() {
	i := 10
	modifyInt(i)
	fmt.Println("int值被修改了，新值为:", i)
}

func modifyInt(i int) {
	fmt.Printf("函数里接收到的指针的内存地址是：%p\n", &i)
	i = 1
}

/**
@Desc map的值会改变
@author zhanghaisheng@qimao.com
@dateTime   : 2021/1/6 3:12 下午
*/
func tMap() {
	i := make(map[int]int, 0)
	i[0] = 0
	fmt.Println("map原来的值为:", i)
	newI := i
	modifyMap(newI)
	fmt.Println("map值被修改了，新值为:", i)
}

func modifyMap(i map[int]int) {
	i[0] = 10
}

type person struct {
	Age int
}

/**
 * struct值不变
 *
 * @author zhanghaisheng@qimao.com
 * @dateTime   : 2021/1/6 3:23 下午
 */
func tStruct() {
	i := person{
		Age: 10,
	}
	fmt.Println("struct原来的值为:", i)

	modifyStruct(i)
	fmt.Println("struct值还是原来的值", i)
}

func modifyStruct(i person) {
	i.Age = 1
}

/**
 * slice值会被修改
 *
 * @author zhanghaisheng@qimao.com
 * @dateTime   : 2021/1/6 3:29 下午
 */
func tSlice() {
	i := make([]int, 5)
	i[0] = 11
	fmt.Println("slice原来的值为:", i)

	modifySlice(i)
	fmt.Println("slice值被修改了，新值为:", i)
}

func modifySlice(i []int) {
	i[0] = 100
}

/**
 * array的值不会被修改
 *
 * @author zhanghaisheng@qimao.com
 * @dateTime   : 2021/1/6 3:31 下午
 */
func tArray() {
	var i [5]int
	i[0] = 11
	fmt.Println("array原来的值为:", i)

	modifyArray(i)
	fmt.Println("array值还是原来的值", i)
}

func modifyArray(i [5]int) {
	i[0] = 100
}

func TestParams(t *testing.T) {
	t.Log("----")
	tInt()
	tArray()
	tStruct()

	//下面的会被修改，用make的会被修改,因为make出来的是 指针类型  https://www.flysnow.org/2018/02/24/golang-function-parameters-passed-by-value.html
	tMap()
	tSlice()

}
