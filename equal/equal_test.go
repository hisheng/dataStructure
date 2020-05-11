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

	//2 .特别是 struct的比较相等
	//var conf MysqlConfig
	//if err := config.Get("mysql", key).Scan(&conf); err != nil {
	//	return nil, err
	//}
	//if reflect.DeepEqual(conf, MysqlConfig{}) {
	//
	//}
}
