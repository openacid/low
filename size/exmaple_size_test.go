package size_test

import "fmt"
import "github.com/openacid/low/size"

func ExampleStat() {
	type my struct {
		a []int32
		b [3]int32
	}

	v := my{
		a: []int32{1, 2, 3},
		b: [3]int32{4, 5, 6},
	}

	fmt.Println(size.Of(v))
	got := size.Stat(v, 10, 100)
	fmt.Println(got)

	// Output:
	// 48
	// size_test.my: 48
	//     a: []int32: 36
	//         0: int32: 4
	//         1: int32: 4
	//         2: int32: 4
	//     b: [3]int32: 12
	//         0: int32: 4
	//         1: int32: 4
	//         2: int32: 4
}
