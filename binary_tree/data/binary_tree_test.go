package data

//https://zhuanlan.zhihu.com/p/95610262
import (
	"fmt"
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestTree_Insert(t *testing.T) {
	//情况1 增加第1条
	btree := NewTree()
	btree.Insert(0)

	fmt.Println(btree)
	t.Log("情况1 增加第1条 btree.root.data =", btree.root.data)
	assert.Equal(t, btree.root.data, 0)

	//目前不是并发安全的，此时再执行 btree.Insert(1)  就会出错

	btree.Insert(1)
	t.Log("情况2 增加第2条 btree.root.right.data =", btree.root.right.data)
	assert.Equal(t, btree.root.right.data, 1)

	btree.Insert(-10)
	t.Log("情况3 增加第3条 btree.root.left.data =", btree.root.left.data)
	assert.Equal(t, btree.root.left.data, -10)

	btree.Insert(-9)
	t.Log("情况3 增加第3条 btree.root.left.right.data =", btree.root.left.right.data)
	assert.Equal(t, btree.root.left.right.data, -9)

	btree.Insert(-11)
	t.Log("情况3 增加第3条 btree.root.left.left.data =", btree.root.left.left.data)
	assert.Equal(t, btree.root.left.left.data, -11)

	//search -9
	fmt.Println("search -9", btree.Search(-10))

	//用一个队列来存储 nodes
	fmt.Println("PreOrder")
	PreOrder(btree.root)
	fmt.Println("InOrder")
	InOrder(btree.root)
	fmt.Println("PostOrder")
	PostOrder(btree.root)
}
