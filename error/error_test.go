/**
@author 张海生<zhanghaisheng@qimao.com>
@dateTime   : 2021/2/4 7:11 下午
@Desc
*/
package error

import (
	"errors"
	"testing"
)

func TestErr(t *testing.T) {
	t.Log(errors.New("test 1"))
	//t.Error(errors.New("test 2"))

	if "a" == "a" {
		t.Log("a == a ")
	}
	//及早失败，避免嵌套！
	panic(errors.New("panic"))

	t.Log("end")
}

func TestErr2(t *testing.T) {
	defer func() {
		panic(errors.New("panic"))
	}()
	//及早失败，避免嵌套！

	t.Log("end")
}

func TestErr3(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			// 这个有点像  try catch
			t.Log("panic ")
		}
	}()
	//及早失败，避免嵌套！

	panic("error ")
}
