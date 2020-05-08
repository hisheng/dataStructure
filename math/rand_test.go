package math

import (
	"math/rand"
	"testing"
)

func TestRand(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log("rand num 32int= ", rand.Int31n(10))
		//t.Log("rand num int = ", rand.Intn(20))
	}
}

func TestInt(t *testing.T) {
	t.Log(rand.Int())
	t.Log(rand.Int31())
	t.Log(rand.Intn(10))
	t.Log(rand.Int31n(10))
	t.Log(rand.Float32())
	t.Log(rand.Float64())

}
