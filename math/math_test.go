package math

import (
	"github.com/hisheng/dataStructure/binary_tree/data"
	"math"
	"math/rand"
	"testing"
)

func TestFloatToInt(t *testing.T) {
	x := 1.1
	t.Log("向上取整", math.Ceil(x))  //2
	t.Log("向下取整", math.Floor(x)) //1
	t.Log("四舍五入", math.Round(x)) //1

	x = 1.5
	t.Log("向上取整", math.Ceil(x))  //2
	t.Log("向下取整", math.Floor(x)) //1
	t.Log("四舍五入", math.Round(x)) //2

	x = 1.8
	t.Log("向上取整", math.Ceil(x))  //2
	t.Log("向下取整", math.Floor(x)) //1
	t.Log("四舍五入", math.Round(x)) //2

	root := data.NewNode(250)
	for i := 0; i < 14; i++ {
		n := rand.Intn(500)
		t.Log("i=", i, "的随机数是", n)
		root.Insert(data.NewNode(n))
	}
	data.PostOrderTraversal(root)
}

/*
Golang 获取两个数之间的随机数，下面的例子可以随机生成4 位数字的验证码

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

func main() {
	myrand := random(1000, 9999)
	fmt.Println(myrand)*/
