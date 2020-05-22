package byte

import (
	"testing"
)

func TestByte(t *testing.T) {
	t.Log("start")
	t.Logf("%b", 0) //0
	t.Logf("%b", 1) //1
	t.Logf("%b", 2) //10
	t.Logf("%b", 3) //11
}
