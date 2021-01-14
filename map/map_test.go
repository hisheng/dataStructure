package _map

import (
	"fmt"
	"testing"
)

type ps struct{}

func (p ps) init() {
	fmt.Println("---------ps init----")
}

func init() {
	fmt.Println("---------init----")
}

func TestMap(t *testing.T) {
	var person map[string]string
	person = make(map[string]string, 0) //必须要这个
	person["name"] = "hisneg"
	person["age"] = "22"

	//map range
	for k, v := range person {
		t.Logf("k=%s v=%s", k, v)
	}
	//map 正确的方法是
	for key := range person {
		t.Log(key)
	}

	//2 第二种map赋值
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		t.Logf("k=%s v=%s", k, v)
	}
}

func TestMapKey(t *testing.T) {
	var person = make(map[string]string, 0)
	person["name"] = "hisneg"
	person["age"] = "22"

	value, exists := person["name"]
	fmt.Println(value, exists)

	value, exists = person["name_none"]
	a := person["name_none"]
	t.Logf("%p", &a)
	t.Log(person["name_none"])
	t.Log("----")
	t.Logf("%v %v", value, exists)
	t.Log("----end")

}

func TestMapInit(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 2, 3: 3}
	t.Log(m1[2])
	t.Logf("lem m1 = %d", len(m1))

	m3 := make(map[int]int, 10) // 10 是容量，len()是0
	t.Log(m3)                   //map[] 用len() 函数判断为空 不能用nil
	if m3 != nil {
		t.Log("is not nil")
	}

	m4 := map[int]int{} //map[]
	if m4 != nil {
		t.Log("m4 is not nil")
	}
	// m4[5]  不存在会返回0
	//需要
	if v, ok := m4[5]; ok {
		t.Log(v)
	} else {
		t.Log("no ")
	}

	for k, v := range m1 {
		t.Log(k, v)
	}
	//打印的是乱序

}

func TestMapWithValue(t *testing.T) {
	m := map[int]func(op int) int{}
	m[1] = func(op int) int {
		return op
	}
	m[2] = func(op int) int {
		return op * op
	}
	m[3] = func(op int) int {
		return op * op * op
	}
	t.Log(m[1](2), m[2](2), m[3](2)) //2 4 8

	n := m
	delete(m, 1)
	t.Log(len(m)) //2
	t.Log(len(n)) //2

}

func TestMapForSet(t *testing.T) {
	set := map[int]bool{}
	set[1] = true
	if set[1] {
		t.Log("ok")
	} else {
		t.Log("no")
	}
	t.Log(set[3]) //false

	m := map[int]string{}
	t.Log(m[3]) //空字符串

	delete(set, 1)
	t.Log(len(set)) //变成0了

}
