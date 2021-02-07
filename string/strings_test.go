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

func TestX(t *testing.T) {
	//t.Logf("%d", "\xE4")
	t.Log(strconv.ParseInt("E4", 16, 32))
}

func TestS(t *testing.T) {
	var s string = "hello"
	//s[1] = 's' //不能修改
	sb := []byte(s)
	sb[1] = 'g'
	t.Log(sb)
	t.Log(string(sb))

	s = "中华人民共和国"
	for _, i2 := range s {
		t.Logf("%[1]c %[1]d", i2)
	}
	//range 的话 是utf-8 "中华人民共和国"

	//strings.Split(s, ",")
	//strings.Join()

	var r []rune
	r = []rune("中试试")
	//r[1] = rune("哈")
	r[1] = '哈'
	t.Log(r)
	for _, i2 := range r {
		t.Logf("%[1]c %[1]d", i2)
	}

	var b []byte
	b = []byte("hello")
	b[1] = 'l'
	t.Log(string(b))

}

func TestS2(t *testing.T) {
	x := "test"
	// x[0] = 'T' error
	xbytes := []byte(x)
	xbytes[0] = 'T'
	t.Log(xbytes)         //[84 101 115 116]
	t.Log(string(xbytes)) //Test

	s := "text"
	fmt.Println(s[0])      //print 116
	fmt.Printf("%T", s[0]) //prints uint8
}
