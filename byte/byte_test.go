package byte

import (
	"testing"
)

func TestByte(t *testing.T) {
	t.Log("start")
	t.Logf("%b", 0)   //0
	t.Logf("%b", 1)   //1
	t.Logf("%b", 2)   //10
	t.Logf("%b", 3)   //11
	t.Logf("%b", 244) //11

	var ziplist []byte
	//ziplist[0] = 123
	//ziplist[1] = 10
	//ziplist[1:6] = "hello"

	ziplist = []byte("hello")
	t.Log(len(ziplist))
	t.Logf("%b ", ziplist[0])
	t.Logf("%b", ziplist[1])
	t.Logf("%b", ziplist[2])
	t.Logf("%b", ziplist[3])
	t.Logf("%b", ziplist[4])
	ziplist[4] = 123
	t.Log(string(ziplist))
	t.Log(ziplist[0])
	t.Log(string(ziplist[0]))

	zips := []rune("你号世界")
	ziplist[4] = 123
	t.Log(string(zips))
	t.Log(zips[0])
	t.Log(string(zips[0]))

}
