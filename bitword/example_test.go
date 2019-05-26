package bitword

import "fmt"

func Example() {
	bw4 := BitWord[4]

	fmt.Println(bw4.FromStr("abc"))
	fmt.Println(bw4.ToStr([]byte{6, 1, 6, 2, 6, 3}))
	fmt.Println(bw4.Get("abc", 1))
	fmt.Println(bw4.FirstDiff("abc", "abd", 0, -1))

	// Output:
	// [6 1 6 2 6 3]
	// abc
	// 1
	// 5

}
