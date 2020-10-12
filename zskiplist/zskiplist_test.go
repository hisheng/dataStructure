package zskiplist

import (
	"fmt"
	"testing"
)

func TestZslRandomLevels(t *testing.T) {
	fmt.Println("TestZslRandomLevels")
	var nums map[int]int = make(map[int]int, 0)
	var num int
	for j := 0; j < 10000; j++ {
		num = ZslRandomLevel()
		_, ok := nums[num]
		if ok {
			nums[num]++
		} else {
			nums[num] = 1
		}
	}
	fmt.Println(nums)
	fmt.Println("TestZslRandomLevels")
}

func TestZslInsert(t *testing.T) {
	fmt.Println("TestZslInsert")
	zsl := ZslCreate()
	fmt.Println(zsl)
	ZslInsert(zsl, 1, "ele")
	ZslInsert(zsl, 2, "ele2")
	fmt.Println(zsl)
	fmt.Println("TestZslInsert")
}
