package slice

import (
	"testing"
	"time"
)

type person struct {
	name string
	age  int
}

func TestSlice(t *testing.T) {

	var ps []person
	ps = []person{
		{name: "hisheng", age: 12}, {name: "hiss", age: 11}, {name: "hiss", age: 11}, {name: "hiss", age: 11}, {name: "hiss", age: 11}, {name: "hiss", age: 11},
		{name: "aaaa", age: 22}, {name: "hiss", age: 11}, {name: "hiss", age: 11}, {name: "hiss", age: 11}, {name: "hiss", age: 11}, {name: "hiss", age: 11},
		{name: "ss", age: 21}, {name: "hiss", age: 11}, {name: "hiss", age: 11}, {name: "hiss", age: 11}, {name: "hiss", age: 11}, {name: "hiss", age: 11},
		{name: "hiss", age: 11}, {name: "hiss", age: 11}, {name: "hiss", age: 11}, {name: "hiss", age: 11}, {name: "hiss", age: 11}, {name: "hiss", age: 11},
		{name: "hiss", age: 11}, {name: "hiss", age: 11}, {name: "hiss", age: 11}, {name: "hiss", age: 11}, {name: "hiss", age: 11}, {name: "hiss", age: 11},
		{name: "hiss", age: 11}, {name: "hiss", age: 11}, {name: "hiss", age: 11}, {name: "hiss", age: 11}, {name: "hiss", age: 11}, {name: "hiss", age: 11},
		{name: "hiss", age: 11}, {name: "hiss", age: 11}, {name: "hiss", age: 11}, {name: "hiss", age: 11}, {name: "hiss", age: 11}, {name: "hiss", age: 11},
		{name: "hiss", age: 11}, {name: "hiss", age: 11}, {name: "hiss", age: 11}, {name: "hiss", age: 11}, {name: "hiss", age: 11}, {name: "hiss", age: 11},
		{name: "hiss", age: 11}, {name: "hiss", age: 11}, {name: "hiss", age: 11}, {name: "hiss", age: 11}, {name: "hiss", age: 11}, {name: "hiss", age: 11},
		{name: "hiss", age: 11}, {name: "hiss", age: 11}, {name: "hiss", age: 11}, {name: "hiss", age: 11}, {name: "hiss", age: 11}, {name: "hiss", age: 11},
		{name: "hiss", age: 11}, {name: "hiss", age: 11}, {name: "hiss", age: 11}, {name: "hiss", age: 11}, {name: "hiss", age: 11}, {name: "hiss", age: 11},
		{name: "hiss", age: 11}, {name: "hiss", age: 11}, {name: "hiss", age: 11}, {name: "hiss", age: 11}, {name: "hiss", age: 11}, {name: "hiss", age: 11},
		{name: "hiss", age: 11}, {name: "hiss", age: 11}, {name: "hiss", age: 11}, {name: "hiss", age: 11}, {name: "hiss", age: 11}, {name: "hiss", age: 11},
		{name: "hiss", age: 11}, {name: "hiss", age: 11}, {name: "hiss", age: 11}, {name: "hiss", age: 11}, {name: "hiss", age: 11}, {name: "hiss", age: 11},
		{name: "hiss", age: 11}, {name: "hiss", age: 11}, {name: "hiss", age: 11}, {name: "hiss", age: 11}, {name: "hiss", age: 11}, {name: "hiss", age: 11},
		{name: "hiss", age: 11}, {name: "hiss", age: 11}, {name: "hiss", age: 11}, {name: "hiss", age: 11}, {name: "hiss", age: 11}, {name: "hiss", age: 11},
		{name: "hiss", age: 11}, {name: "hiss", age: 11}, {name: "hiss", age: 11}, {name: "hiss", age: 11}, {name: "hiss", age: 11}, {name: "hiss", age: 11},
		{name: "hiss", age: 11}, {name: "hiss", age: 11}, {name: "hiss", age: 11}, {name: "hiss", age: 11}, {name: "hiss", age: 11}, {name: "hiss", age: 11},
		{name: "hiss", age: 11}, {name: "hiss", age: 11}, {name: "hiss", age: 11}, {name: "hiss", age: 11}, {name: "hiss", age: 11}, {name: "hiss", age: 11},
		{name: "hiss", age: 11}, {name: "hiss", age: 11}, {name: "hiss", age: 11}, {name: "hiss", age: 11}, {name: "hiss", age: 11}, {name: "hiss", age: 11},
		{name: "hiss", age: 11}, {name: "hiss", age: 11}, {name: "hiss", age: 11}, {name: "hiss", age: 11}, {name: "hiss", age: 11}, {name: "hiss", age: 11},
		{name: "hiss", age: 11}, {name: "hiss", age: 11}, {name: "hiss", age: 11}, {name: "hiss", age: 11}, {name: "hiss", age: 11}, {name: "hiss", age: 11},
	}
	var newPs []person
	t.Log(ps[0])
	t.Log(ps[0:0])
	newPs = ps[0:0]
	t.Log(newPs)

	for i := 1; i < 10; i++ {
		ps = append(ps, ps...)
	}

	//方法1 节省时间
	start := time.Now().UnixNano()
	for i, p := range ps {
		if p.age == 21 {
			newPs = ps[0:i]
			newPs = append(newPs, ps[i:]...)
			break
		}
	}
	end := time.Now().UnixNano()
	t.Log(end - start)
	//t.Log(newPs)

	newPs = make([]person, 0)
	start1 := time.Now().UnixNano()
	for _, p := range ps {
		if p.age == 21 {
			continue
		}
		newPs = append(newPs, p)
	}
	end1 := time.Now().UnixNano()
	t.Log(end1 - start1)
	//t.Log(newPs)
}
