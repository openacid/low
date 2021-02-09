package bitmap

// TailBitmap is a special case bitmap in which the bits before a certain index
// are all set to `1`.
// Thus we only need to store the sparse tail part.
//
// The data structure is as the following described:
//
//                      reclaimed
//                      |
//                      |     Offset
//                      |     |
//                      v     v
//                ..... X ... 01010...00111  00...
//   bitIndex:    0123...     ^              ^
//                            |              |
//                            Words[0]       Words[1]
//
// Since 0.1.22
type TailBitmap struct {
	// The bit index from which there are a `0` in it.
	// Before Offset all bits are set to `1` thus we donot need to store them.
	// Offset must be a `n*64`.
	Offset int64
	// Words is the tail part bitmap in which there are `0`.
	Words []uint64

	// reclaimed is the bit index from which the memory in Words are reclaimed.
	reclaimed int64
}

// reclaimThreshold is the size threshold in bit for reclamation of `Words`.
var reclaimThreshold = int64(1024) * 64

// NewTailBitmap creates an TailBitmap with a preset Offset and an empty
// tail bitmap.
//
// Since 0.1.22
func NewTailBitmap(offset int64) *TailBitmap {
	tb := &TailBitmap{
		Offset:    offset,
		reclaimed: offset,
		Words:     make([]uint64, 0, reclaimThreshold>>6),
	}
	return tb
}

// Compact all leading all-ones words in the bitmap.
//
// Since 0.1.22
func (tb *TailBitmap) Compact() {

	allOnes := uint64(0xffffffffffffffff)

	for len(tb.Words) > 0 && tb.Words[0] == allOnes {
		tb.Offset += 64
		tb.Words = tb.Words[1:]
	}

	if tb.Offset-tb.reclaimed >= reclaimThreshold {
		l := len(tb.Words)
		newWords := make([]uint64, l, l*2)

		copy(newWords, tb.Words)
		tb.reclaimed = tb.Offset
	}
}

// Set the bit at `idx` to `1`.
//
// Since 0.1.22
func (tb *TailBitmap) Set(idx int64) {
	if idx < tb.Offset {
		return
	}

	idx = idx - tb.Offset
	wordIdx := idx >> 6

	for int(wordIdx) >= len(tb.Words) {
		tb.Words = append(tb.Words, 0)
	}

	tb.Words[wordIdx] |= Bit[idx&63]

	if wordIdx == 0 {
		tb.Compact()
	}
}

// Get retrieves a bit at its 64-based offset.
//
// Since 0.1.22
func (tb *TailBitmap) Get(idx int64) uint64 {
	if idx < tb.Offset {
		return Bit[idx&63]
	}

	idx = idx - tb.Offset
	return tb.Words[idx>>6] & Bit[idx&63]
}

// Get1 retrieves a bit and returns a 1-bit word, i.e., putting the bit in the
// lowest bit.
//
// Since 0.1.22
func (tb *TailBitmap) Get1(idx int64) uint64 {
	if idx < tb.Offset {
		return 1
	}
	idx = idx - tb.Offset
	return (tb.Words[idx>>6] >> uint(idx&63)) & 1
}
