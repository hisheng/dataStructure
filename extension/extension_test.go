/**
@author 张海生<zhanghaisheng@qimao.com>
@dateTime   : 2021/1/28 7:42 下午
@Desc
*/
package extension

import (
	"fmt"
	"testing"
)

type Parent struct {
	Age int
}

func (p Parent) Speak() {
	fmt.Println("...")
}

func (p Parent) SpeakTo(name string) {
	p.Speak()
	fmt.Println(name)
}

///LSP 继承

type Dog struct {
	Parent //Pet 是嵌套类型   golang，不能循环嵌套
}

func TestDog(t *testing.T) {
	dog := new(Dog)
	t.Log(dog)
	t.Log(dog.Age)

	//断言
	//v, ok := p.(int) //ok = true

}
