// Package float provides with floating point integer number encoding/decoding.
//
// Since 0.1.9
package float

import (
	"math/bits"

	"github.com/openacid/must"
)

// U8 encodes a uint64 into a floating point integer in uint8, with reduced
// precision.
// It keeps only the highest non-zero 4*k to 4*k+3 bits,
// and the exponent k.
//
// E.g. 0xa0bf in base-16 form is:
//
//   0xa0bf = 0xa*x^3 + 0xb*x + 0xf ; x = 16
//
// To encode it into a 8-bit floating point int,
// put the exponent in the high 4 bits, and the significand in lower 4 bits:
//
//   U8(0xa0bf) = (0x03<<4) | 0xa = 0x3a
//
// And abviously:
//
//   U8(n) == n if n < 16
//
// Since 0.1.9
func U8(n uint64) uint8 {

	if n < 16 {
		return uint8(n)
	}

	exponent := (60 - bits.LeadingZeros64(n))
	must.Be.True(exponent < 16)

	significand := n >> uint(exponent)

	return uint8(exponent<<4) | uint8(significand)
}

// U8decode decodes a floating point uint8 back into a uint64.
//
// Since 0.1.9
func U8decode(n uint8) uint64 {
	return (uint64(n) & 0x0f) << uint(n>>4)
}

// U8normalize returns the most precise value of n that a floating point uint8 could
// have.
//
// Since 0.1.9
func U8normalize(n uint64) uint64 {
	return U8decode(U8(n))
}
