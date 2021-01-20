package zipf

import (
	"fmt"
	"strings"
)

func ExampleAccesses() {

	n := 10

	// y = C * x**(1-s),  x ∈ [a, b)
	a := float64(1)
	s := float64(1.5)

	got := Accesses(a, s, n, 50, nil)
	fmt.Println(got)

	arr := make([]int, n)
	for _, idx := range got {
		arr[idx]++
	}
	for _, v := range arr {
		fmt.Println("|" + strings.Repeat("*", v))
	}

	// Output:
	//
	// [0 3 5 6 6 5 8 7 7 7 0 8 6 6 0 6 4 0 7 6 6 6 0 6 2 6 5 6 0 6 8 5 7 3 0 4 0 8 7 0 6 3 6 0 0 4 6 6 7 7]
	// |***********
	// |
	// |*
	// |***
	// |***
	// |****
	// |****************
	// |********
	// |****
	// |
}
