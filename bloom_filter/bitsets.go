package bloom_filter

type BitSets []int64

func NewBitSets(n uint) BitSets {
	bs := make(BitSets, n/64+1)
	return bs
}

func (bs BitSets) Set(index uint) {
	index, bit := index/63, index%64
	bs[index] |= 1 << bit
}
