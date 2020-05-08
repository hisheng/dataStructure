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

}
