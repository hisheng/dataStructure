/**
@author 张海生<zhanghaisheng@qimao.com>
@dateTime   : 2021/3/4 7:24 下午
@Desc
*/
package async

import (
	"testing"
)

//case 是无顺序的，default 是有顺序的  switch的case是有顺序的，一直下来的
func TestSelect(t *testing.T) {
	//select {
	//case ret := <-retCh1:
	//	t.Logf("result1 %s", ret)
	//case ret := <-retCh2:
	//	t.Logf("result2 %s", ret)
	//case <-time.After(time.Second * 1): //超时
	//	t.Log("time out")
	//default:
	//}
	// data,ok := <-chanel 两个返回值，来看chanenl
}

//超时
