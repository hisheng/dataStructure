package sort

import (
	"fmt"
	"sort"
	"testing"
)

type Person struct {
	Name string
	Age  int
}
type PersonSlice []Person

func (ps PersonSlice) Len() int           { return len(ps) }
func (ps PersonSlice) Less(i, j int) bool { return ps[i].Age < ps[j].Age }
func (ps PersonSlice) Swap(i, j int)      { ps[i], ps[j] = ps[j], ps[i] }

func TestSort(t *testing.T) {
	people := []Person{
		{"bo", 31},
		{"ao", 42},
		{"co", 17},
		{"do", 26},
	}
	//方法1
	sort.Sort(PersonSlice(people))
	fmt.Println(people)

	//方法2 slice 也同样可以用来排序
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age > people[j].Age
	})
	fmt.Println(people)

	//2 这个也得测试一下 稳定排序和不稳定排序
	//张海生 5-14 09:11:19
	//sort 排序是不稳定排序，Stable 是稳定的排序
	//
	//张海生 5-14 09:11:48
	//Stable 的底层 是 插入和归并所以它是稳定排序
	//
	//张海生 5-14 09:12:21
	//sort底层是 快排，插入，堆排序
	//
	//张海生 5-14 09:12:46
	//快速排序 ，堆排序 不是稳定的排序算法
}

type PersonMultipleSlice []PersonMultiple
type PersonMultiple struct {
	Name   string
	Age    int
	Height int
}

func (ps PersonMultipleSlice) Len() int      { return len(ps) }
func (ps PersonMultipleSlice) Swap(i, j int) { ps[i], ps[j] = ps[j], ps[i] }
func (ps PersonMultipleSlice) Less(i, j int) bool {
	if ps[i].Age < ps[j].Age {
		return true
	}
	if ps[i].Age > ps[j].Age {
		return false
	}

	return ps[i].Height < ps[j].Height
}

//多阶排序 (此时是需要稳定排序的)
//https://stackoverflow.com/questions/36122668/how-to-sort-struct-with-multiple-sort-parameters
func TestMultiple(t *testing.T) {
	//先按age 升序，然后 再按身高 升序
	ps := []PersonMultiple{
		{"bo", 31, 11},
		{"ao", 42, 20},
		{"ao", 42, 18},
		{"ao", 42, 12},
		{"ao", 42, 21},
		{"co", 17, 13},
		{"do", 26, 14},
	}
	sort.Sort(PersonMultipleSlice(ps))
	t.Log(ps)
}
