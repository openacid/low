package bitmap

import (
	"math/bits"

	"golang.org/x/sys/cpu"
)

var selectU64 func(uint64, uint64) uint64

func init() {
	// TODO arch and os
	if cpu.X86.HasBMI2 {
		selectU64 = selectU64WithPDEP
	} else {
		selectU64 = selectU64WithoutPDEP
	}

}

// selectU64WithPDEP finds the ith 1 in a uint64.
//
// The implementation requires a BMI-2 instruction "PDEP"(parallel deposit),
// which moves the lowest bits to positions specified by a mask.
//
// PDEP moves bits in a to specified position by mask.
//
//     src  = 0000 0101
//             .__////
//             |  /||
//             v v vv
//     mask = 0101 1100
//             v v vv
//     rst  = 0001 0100
//
// With it, we can place a "1" at the ith bit.
// Then moves it with "word" as a mask.
// Then retrieve the position of the only "1".
//
// This function is not complete that it only works on a amd64 platform.
//
// go:noescape
func selectU64WithPDEP(word uint64, ith uint64) uint64

// selectU64WithoutPDEP finds the ith 1 in a uint64.
func selectU64WithoutPDEP(w uint64, findIth uint64) uint64 {

	ith := int(findIth)
	ones := bits.OnesCount64(w)
	if ones <= ith {
		return 0xffffffffffffffff
	}

	base := int32(0)
	ww := w

	ones = bits.OnesCount32(uint32(ww))

	if ones <= ith {
		ith -= ones
		base |= 32
		ww >>= 32
	}

	ones = bits.OnesCount16(uint16(ww))

	if ones <= ith {
		ith -= ones
		base |= 16
		ww >>= 16
	}

	ones = bits.OnesCount8(uint8(ww))

	if ones <= ith {
		return uint64(int32(tbl2[(ww>>5)&(0x7f8)|uint64(ith-ones)]) + base + 8)
	} else {
		return uint64(int32(tbl2[(ww&0xff)<<3|uint64(ith)]) + base)
	}
}
