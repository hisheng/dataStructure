package string

import (
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

}
