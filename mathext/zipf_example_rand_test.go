package mathext

import (
	"fmt"
	"math/rand"
	"strings"
)

func ExampleZipf() {

	n := 20

	// y = C * x**(1-s),  x ∈ [a, b)
	a := float64(1)
	b := float64(20)
	s := float64(1.5)

	z := NewZipf(a, b, s)
	sampleCnt := float64(100)

	sample := make([]int, n)
	r := rand.New(rand.NewSource(44))
	for u := float64(0); u < 1; u += 1 / sampleCnt {
		v := r.Float64()
		x := int(z.Float64(v))
		sample[x]++
	}

	for _, v := range sample {
		fmt.Println("|" + strings.Repeat("*", v))
	}

	// Output:
	//
	// |
	// |*******************************
	// |****************
	// |****************
	// |*******
	// |**
	// |*****
	// |***
	// |***
	// |*******
	// |*
	// |**
	// |*
	// |***
	// |
	// |
	// |
	// |*
	// |*
	// |*
}
