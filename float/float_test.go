/**
@author 张海生<zhanghaisheng@qimao.com>
@dateTime   : 2021/3/19 2:46 下午
@Desc
*/
package float

import (
	"fmt"
	"math"
	"strconv"
	"testing"
)

//100 次的 0.01 相加 并不是 1  在程序中
func TestSum(t *testing.T) {
	var sum float32
	var a float32 = 0.01

	for i := 0; i < 100; i++ {
		sum += a
	}
	t.Log(sum)                               //0.99999934
	t.Logf("%.2f", sum)                      //1.00
	t.Logf("%b\n", math.Float32bits(0.0625)) //0 01111011 00000000000000000000000 = 1*(2负四次方) = 0.0625
	//指数位是 01111011---  123 - 127 = -4
	//指数位是 01111110---126 excess 系统是 (126-127) = -1 //对应的最大值 11111111 = (255-127) = 128
	t.Logf("%d", Btod([]byte{'0', '1', '1', '1', '1', '1', '1', '0'}))

}

// 二进制转十进制
func Btod(read_line []byte) int {

	var value_sum, value_square int

	for index, value := range read_line {

		value_int, error := strconv.Atoi(string(value))

		if error != nil {
			fmt.Println(error)
		}

		index_new := len(read_line) - 1 - index

		if index_new == 0 {
			value_square = 1

		} else {
			value_square = 2 << uint(index_new-1)
		}

		value_sum += value_int * value_square
	}
	return value_sum
}
