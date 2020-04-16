package bloom_filter

type BitSets []int64

func NewBitSets(n uint) BitSets {
	bs := make(BitSets, n/64+1)
	return bs
}

func (bs BitSets) Set(index uint) {
	index, bit := index/64, index%64
	bs[index] |= 1 << bit
}

func (bs BitSets) Unset(index uint) {
	index, bit := index/64, index%64
	bs[index] ^= 1 << bit
}

func (bs BitSets) IsSet(index uint) bool {
	index, bit := index/64, index%64
	num := bs[index]
	return (num | (1 << bit)) == num
}

//size byte
//func (bs BitSets) Size() uint64 {
//	return uint64(len(bs)) * 4
//}
