// Package bitmap provides basic bitmap operations.
// A bitmap uses []uint64 as storage.
//
// Since 0.1.8
package bitmap

import "math/bits"

// IndexRank64 creates a rank index for a bitmap.
// rank(i) is defined as number of "1" upto position i, excluding i.
//
// It returns an index of []int32.
// Every element in it is rank(i*64)
//
// Since 0.1.8
func IndexRank64(words []uint64) []int32 {

	idx := make([]int32, len(words))
	n := int32(0)
	for i := 0; i < len(words); i++ {
		idx[i] = n
		n += int32(bits.OnesCount64(words[i]))
	}

	// clone to reduce cap to len
	idx = append(idx[:0:0], idx...)

	return idx
}

// IndexRank128 creates a rank index for a bitmap.
// rank(i) is defined as number of "1" upto position i, excluding i.
//
// It returns an index of []int32.
// Every element in it is rank(i*128).
//
// It also adds a last index item if len(words) % 2 == 0, in order to make the
// distance from any bit to the closest index be less than 64.
//
//
// Since 0.1.8
func IndexRank128(words []uint64) []int32 {

	idx := make([]int32, 0)
	n := int32(0)
	for i := 0; i < len(words); i += 2 {
		idx = append(idx, n)
		n += int32(bits.OnesCount64(words[i]))
		if i < len(words)-1 {
			n += int32(bits.OnesCount64(words[i+1]))
		}
	}

	// Need a last index to let distance from every bit to its closest index
	// <=64
	if len(words)&1 == 0 {
		idx = append(idx, n)
	}

	// clone to reduce cap to len
	idx = append(idx[:0:0], idx...)

	return idx
}
