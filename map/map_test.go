package _map

import (
	"testing"
)

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
