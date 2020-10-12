package main

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
