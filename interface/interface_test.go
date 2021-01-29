/**
@author 张海生<zhanghaisheng@qimao.com>
@dateTime   : 2021/1/28 7:27 下午
@Desc
*/
package _interface

import "testing"

type Programmer interface {
	WriteCode() string
}

type GoP struct {
}

func (gop GoP) WriteCode() string {
	return "h"
}

func TestInter(t *testing.T) {
	g := new(GoP)
	t.Log(g.WriteCode())

}
