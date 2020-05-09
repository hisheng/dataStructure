package equal

import (
	"reflect"
	"testing"
)

func TestEqual(t *testing.T) {
	//1 DeepEqual 可以比较同一类型的
	t.Log(reflect.DeepEqual(1, 1))
	t.Log(reflect.DeepEqual("1", "1"))
	t.Log(reflect.DeepEqual("1", "2"))
}
