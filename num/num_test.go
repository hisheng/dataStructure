/**
@author 张海生<zhanghaisheng@qimao.com>
@dateTime   : 2021/1/20 11:08 上午
@Desc
*/
package num

import (
	"strconv"
	"testing"
)

func TestNum(t *testing.T) {
	disId, err := strconv.Atoi("0")
	t.Log(disId, err)

	disId, err = strconv.Atoi("")
	t.Log(disId, err)

	disId, err = strconv.Atoi("a")
	t.Log(disId, err)

	disId, err = strconv.Atoi("1a")
	t.Log(disId, err)

	disId, err = strconv.Atoi("a1")
	t.Log(disId, err)

	//结果是 "数字" 这种类型 err 是没有的 转成对应的数字
	// 其他的字符串 对应没有数字的，都会转成0 ，有err
}
