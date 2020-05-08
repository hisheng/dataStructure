package hash

import (
	"github.com/go-playground/assert/v2"
	"github.com/spaolacci/murmur3"
	"io"
	"testing"
)

//非加密型哈希函数
//规律性较强的key，MurmurHash的随机分布特征表现更良好
//上面的表述我感觉说的还是不够浅显，Murmur哈希算法最大的特点是：
// 1碰撞率低
// 2计算速度快

func TestMurMur3(t *testing.T) {
	t.Log("s")
	message := "hello world hisheng"
	//quick use
	hs := murmur3.Sum64([]byte(message))

	assert.Equal(t, hs, mumu3(message))
}

func mumu3(message string) uint64 {
	m3 := murmur3.New128()
	io.WriteString(m3, message)
	s, _ := m3.Sum128()
	return s
}
