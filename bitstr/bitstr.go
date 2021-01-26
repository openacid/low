package bitstr

import (
	"bytes"
	"math/bits"
	"unsafe"

	"github.com/openacid/low/bitmap"
)

// New extracts bits from fromBit to toBit and save it into bytes.
// The fromBit must be aligned to 8, or it will be truncated to 8*n.
// The bits after toBit are discard.
//
// To describe non-8 aligned bits, a trailing byte of the effective number of
// bits in the last payload byte is appended.
//
// E.g.:
//   // "abc" = 0110 0001, 0110 0010, 0110 0011
//   New("abc", 5, 12)
//   // returns 3 byte:
//   // 0110 0001, 0110 0000, 1111, 0000
//   // ---------------       ----------
//   // payload               trailing byte:
//   //                       the higher 4 bit in the last(2nd) byte.
//
// Since 0.1.20
func New(s string, fromBit, toBit int32) []byte {

	if fromBit == toBit && fromBit&7 == 0 {
		return []byte{0xff}
	}

	fromByte := fromBit >> 3
	toByte := (toBit + 7) >> 3

	l := toByte - fromByte

	bitStr := make([]byte, l+1)
	copy(bitStr, s[fromBit>>3:toByte])

	mask := byte(bitmap.RMask[(8-toBit)&7])

	// truncate bits after `toBit`
	bitStr[l-1] &= mask

	// the last byte stores number of effective bits in the last payload byte.
	bitStr[l] = mask

	return bitStr
}

// Cmp compares two bitStr and returns -1, 0, 1 for less-than, equal, greater-than.
//
// Since 0.1.20
func Cmp(a, b []byte) int {
	la, lb := len(a), len(b)
	if la == lb {
		return bytes.Compare(a, b)
	}

	return bytes.Compare(a[:la-1], b[:lb-1])
}

func cmpBytes(a []byte, b []byte) int {

	la, lb := len(a), len(b)
	if la < 8 {
		var i int
		for i = 0; i < la; i++ {
			if a[i] < b[i] {
				return -1
			} else if a[i] > b[i] {
				return 1
			}
		}

		if i < lb {
			return -1
		}
		return 0
	}
	return bytes.Compare(a, b)
}

// CmpUpto compares a normal []byte `a` and a bitStr `b` with `a` being
// TRUNCATED upto `Len(b)`.
//
// Performance
//  >=8 bytes: TODO
//  <8 bytes: TODO
//
// Since 0.1.20
func CmpUpto(a, b []byte) int {
	la, lb := len(a), len(b)
	if lb == 1 {
		// an empty bitStr
		return 0
	}

	if la < lb-1 {
		return cmpBytes(a, b[:lb-1])
	}

	// la >= lb-1

	la = lb - 1

	rst := cmpBytes(a[:lb-2], b[:lb-2])
	if rst != 0 {
		return rst
	}

	// compare the last byte

	bytea := a[la-1] & b[lb-1]
	byteb := b[la-1]

	if bytea > byteb {
		return 1
	} else if bytea < byteb {
		return -1
	}

	return 0
}

// StrCmpUpto is similar to CmpUpto, except the first arg is a string instead of
// []byte.
//
// Since 0.1.20
func StrCmpUpto(a string, b []byte) int {
	return CmpUpto(*(*[]byte)(unsafe.Pointer(&a)), b)
}

// Len returns the number of payload bits in a bitStr.
//
// Since 0.1.20
func Len(bs []byte) int32 {
	l := len(bs)

	// -16: exclude the tailing byte, and the last payload byte
	return int32(l)<<3 - 16 + int32(bits.OnesCount8(bs[l-1]))
}
