/**
@author 张海生<zhanghaisheng@qimao.com>
@dateTime   : 2021/2/1 9:48 上午
@Desc
*/
package _return

import "testing"

func TestReturn(t *testing.T) {
	t.Log(ss())
}

func ss() (ab int) {
	return
}
