package bloom_filter

import (
	"crypto/sha256"
	"github.com/go-playground/assert/v2"
	"github.com/spaolacci/murmur3"
	"testing"
)

func TestHashData(t *testing.T) {
	sha_data := sha256.Sum256([]byte{'h', 'e', 'l', 'l', 'o'})
	t.Log(sha_data)
	t.Log(HashFunc([]byte{'h', 'e', 'l', 'l', 'o'}, 0)) //3182187045311309425 19位
}

func TestMemoryBloomFilter_PutString(t *testing.T) {
	memoryBloomFilter := NewMemoryBloomFilter(100, 4)
	l := uint(len(memoryBloomFilter.bs))

	t.Log(l)
	memoryBloomFilter.PutString("1")
	data := "a"
	t.Log("TestMemoryBloomFilter_PutString")
	t.Log([]byte(data))
	t.Log([]byte("A"))

	t.Log(memoryBloomFilter.k)

	for i := uint(0); i < memoryBloomFilter.k; i++ {
		memoryBloomFilter.bs.Set(HashFunc([]byte(data), i) % l)
	}

}

func TestNewMemoryBloomFilter(t *testing.T) {
	hs1 := murmur3.Sum64WithSeed([]byte("byte"), 0)
	hs2 := murmur3.Sum64WithSeed([]byte("a"), 0)
	hs3 := murmur3.Sum64WithSeed([]byte("hello"), 0)
	hs4 := murmur3.Sum64WithSeed([]byte("b"), 0)
	t.Log("murmur3 hs1 = ", hs1)
	t.Log("murmur3 hs2 = ", hs2)
	t.Log("murmur3 hs3 = ", hs3)
	t.Log("murmur3 hs4 = ", hs4)

	memoryBloomFilter := NewMemoryBloomFilter(100, 4)
	t.Log("TestNewMemoryBloomFilter")
	t.Log(memoryBloomFilter.bs)
	memoryBloomFilter.PutString("hello")
	memoryBloomFilter.PutString("world")
	memoryBloomFilter.PutString("hisheng")

	assert.Equal(t, memoryBloomFilter.HasString("hello"), true)
	assert.Equal(t, memoryBloomFilter.HasString("world"), true)
	assert.Equal(t, memoryBloomFilter.HasString("hisheng"), true)

	t.Log("hisheng2", memoryBloomFilter.HasString("hisheng2"))      //有误差！  总共就100个bit，所以会有误差，n变成100000就无误差了
	assert.Equal(t, memoryBloomFilter.HasString("hisheng2"), false) // //有误差！  要么n变成1万，或者k变大(比如10，就无误差)，k相当于多少个hash函数，这里通过不同的seed，相当于返回不同的值

	assert.Equal(t, memoryBloomFilter.HasString("hisheng3"), false)
	assert.Equal(t, memoryBloomFilter.HasString("hisheng4"), false)
	assert.Equal(t, memoryBloomFilter.HasString("world1"), false)

}
