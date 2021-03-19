/**
@author 张海生<zhanghaisheng@qimao.com>
@dateTime   : 2021/3/11 7:34 下午
@Desc
*/
package csp

import "testing"

//1.关闭一个 nil 值 channel 会引发 panic。！
func TestNew(t *testing.T) {
	var ch chan struct{}
	close(ch)
}

//2. 关闭一个已关闭的 channel 会引发 panic。
func TestClose(t *testing.T) {
	ch := make(chan struct{})
	close(ch)
	close(ch)
}

//3. 向一个已关闭的 channel 发送数据。会引发 panic。
func TestCloseSend(t *testing.T) {
	ch := make(chan struct{})
	close(ch)
	ch <- struct{}{}
}
