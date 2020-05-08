package bloom_filter

import (
	"crypto/sha256"
)
import "github.com/spaolacci/murmur3"

type BloomFilter interface {
	Put([]byte)
	PutString(string)

	Has([]byte) bool
	HasString(string) bool

	Close()
}

type MemoryBloomFilter struct {
	k  uint
	bs BitSets
}

/*n 做多存储多少个 bit */
func NewMemoryBloomFilter(n uint, k uint) *MemoryBloomFilter {
	return &MemoryBloomFilter{
		k:  k,
		bs: NewBitSets(n),
	}
}

// Put 添加一条记录
func (filter *MemoryBloomFilter) Put(data []byte) {
	l := uint(len(filter.bs))
	//通过seed，相当于是 不同的 hash函数，得出不同的hash值
	for i := uint(0); i < filter.k; i++ {
		filter.bs.Set(HashFunc(data, i) % l)
	}
}

// Put 添加一条string记录
func (filter *MemoryBloomFilter) PutString(data string) {
	filter.Put([]byte(data))
}

// Has 推测记录是否已存在
func (filter *MemoryBloomFilter) Has(data []byte) bool {
	l := uint(len(filter.bs))

	for i := uint(0); i < filter.k; i++ {
		if !filter.bs.IsSet(HashFunc(data, i) % l) {
			return false
		}
	}

	return true
}

// Has 推测记录是否已存在
func (filter *MemoryBloomFilter) HasString(data string) bool {
	return filter.Has([]byte(data))
}

// Close 关闭bloom filter
func (filter *MemoryBloomFilter) Close() {
	filter.bs = nil
}

/*作用是 hash*/
func HashFunc(data []byte, seed uint) uint {
	sha_data := sha256.Sum256(data) //1加密 sha256 ，是不是也可以不加密?

	return uint(murmur3.Sum64WithSeed(sha_data[0:], uint32(seed))) //2 均匀随机 分布 key
}
