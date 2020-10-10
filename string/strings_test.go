package string

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"testing"
)

//slice to string
func TestSliceToStrings(t *testing.T) {
	numSlice := make([]string, 0)
	for i := 0; i < 16; i++ {
		num := rand.Intn(10)
		numSlice = append(numSlice, strconv.Itoa(num))
	}
	s := strings.Join(numSlice, "")
	t.Log(numSlice)
	t.Log(s)

	helle := "helloname"
	t.Log(helle[0:3])

}

func TestSprint(t *testing.T) {
	t.Log(fmt.Sprint(1))
	var s string
	s = fmt.Sprint(0)
	t.Log(s)
	t.Log(fmt.Sprint("1"))
}


