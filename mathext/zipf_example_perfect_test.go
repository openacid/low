package mathext

import (
	"fmt"
	"strings"
)

func ExampleZipf_perfect() {

	n := 20

	// y = C * x**(1-s),  x ∈ [a, b)
	a := float64(1)
	b := float64(20)
	s := float64(1.5)

	z := NewZipf(a, b, s)
	sampleCnt := float64(100)

	sample := make([]int, n)
	for u := float64(0); u < 1; u += 1 / sampleCnt {
		x := int(z.Float64(u))
		sample[x]++
	}

	for _, v := range sample {
		fmt.Println("|" + strings.Repeat("*", v))
	}

	// Output:
	//
	// |
	// |**************************************
	// |*****************
	// |**********
	// |*******
	// |*****
	// |****
	// |***
	// |**
	// |***
	// |*
	// |**
	// |**
	// |*
	// |*
	// |*
	// |*
	// |*
	// |*
	// |
}
