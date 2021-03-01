package randext

import (
	"math/bits"
	"math/rand"
)

// Geo generates a non-negative random number so that chance of x is p times the
// chance of x+1.
// p must be a power of 2.
//   ~ 30 ns
//
// Since 0.1.22
func Geo(p int) int {
	v := rand.Uint64()

	zeros := bits.LeadingZeros64(v)
	return zeros / bits.TrailingZeros(uint(p))
}

// geoByLoop is a naive impl of geometric random number.
// p=2: 40 ns
// p=4: 23 ns
// p=8: 23 ns
func geoByLoop(p int) int {
	q := float64(1) / float64(p)
	x := 0
	for rand.Float64() < q {
		x++
	}

	return x
}
