package bitmap

import (
	"math/bits"
)

// First returns the index of the first "1" in a bit range "[from, to)".
// If no "1" is found, it returns "to".
//
// Since 0.1.9
func First(words []uint64, from, to int32) int32 {
	to--
	toWordI := to >> 6
	fromWordI := from >> 6

	for i := fromWordI; i <= toWordI; i++ {
		w := words[i]
		if i == toWordI {
			w &= MaskUpto[to&63]
		}
		if i == fromWordI {
			w &= RMask[from&63]
		}

		if w != 0 {
			return (i << 6) + int32(bits.TrailingZeros64(w))
		}
	}
	return to + 1
}

// Last returns the index of the last "1" in a bit range "[from, to)".
// If no "1" is found, it returns "from - 1".
//
// Since 0.1.9
func Last(words []uint64, from, to int32) int32 {
	to--
	toWordI := to >> 6
	fromWordI := from >> 6

	for i := toWordI; i >= fromWordI; i-- {
		w := words[i]
		if i == toWordI {
			w &= MaskUpto[to&63]
		}
		if i == fromWordI {
			w &= RMask[from&63]
		}

		if w != 0 {
			return (i << 6) + int32(63-bits.LeadingZeros64(w))
		}
	}
	return from - 1
}
