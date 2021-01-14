/**
@author 张海生<zhanghaisheng@qimao.com>
@dateTime   : 2021/1/8 9:17 上午
@Desc
*/
package gostudy

import "testing"

func TestSlice(t *testing.T) {
	year := []string{"1", "2", "3"}
	aa := year[:1]
	aa = append(aa, "4")
	t.Log(aa, year)
}
