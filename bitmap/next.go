package bitmap

import (
	"math/bits"
)

// NextOne find the next "1" in range `[i, end)`.
// If there is no "1" found, it returns `end`.
//
// Since 0.1.11
func NextOne(bm []uint64, i, end int32) int32 {

	wordIdx := i >> 6
	bitIdx := i & 63
	var nxt int32 = -1

	word := bm[wordIdx] & RMask[bitIdx]
	if word != 0 {
		nxt = wordIdx<<6 + int32(bits.TrailingZeros64(word))
	} else {

		i = (i + 63) & ^63

		for ; i < end; i += 64 {

			word := bm[i>>6]
			if word != 0 {
				nxt = i + int32(bits.TrailingZeros64(word))
				break
			}
		}
	}
	if nxt >= end {
		return -1
	}
	return nxt
}

// PrevOne find the previous "1" in range `[i, end)`
// If there is no "1" found, it returns -1.
//
// Since 0.1.11
func PrevOne(bm []uint64, i, end int32) int32 {

	end--
	wordIdx := end >> 6
	bitIdx := end & 63
	var prv int32 = -1

	word := bm[wordIdx] & MaskUpto[bitIdx]
	if word != 0 {
		prv = wordIdx<<6 + 63 - int32(bits.LeadingZeros64(word))
	} else {
		end = (end & ^63) - 1

		for ; end >= i; end -= 64 {

			word := bm[end>>6]
			if word != 0 {
				prv = end - int32(bits.LeadingZeros64(word))
				break
			}
		}
	}

	if prv < i {
		return -1
	}

	return prv
}
