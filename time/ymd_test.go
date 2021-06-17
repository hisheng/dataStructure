/**
@author 张海生<zhanghaisheng@qimao.com>
@dateTime   : 2021/5/31 2:39 下午
@Desc
*/
package time

import (
	"testing"
	"time"
)

func TestYmd(t *testing.T) {
	t.Log(time.Now().Format("2006-01-02 15:04:05"))
	t.Log(time.Now().Format("20060102"))
	// 4 的倍数
}
