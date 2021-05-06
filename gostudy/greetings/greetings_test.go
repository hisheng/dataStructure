/**
@author 张海生<zhanghaisheng@qimao.com>
@dateTime   : 2021/4/16 10:52 上午
@Desc
*/
package greetings

import (
	"regexp"
	"testing"
)

func TestHelloEmpty(t *testing.T) {
	//msg, err := Hello("")
	//if msg != "" || err == nil {
	//	t.Fatal(`Hello("") = %q, %v, want "", error`, msg, err)
	//}
}

func TestHelloName(t *testing.T) {
	name := "Gladys"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Gladys")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
	}

}

const (
	mutexLocked = 1 << iota
	mutexWoken
	mutexWaiterShift
	ss
)

func TestName(t *testing.T) {
	t.Log(mutexLocked)
	t.Log(mutexWoken)
	t.Log(mutexWaiterShift)
	t.Log(ss)
}
