package bmtree

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMult(t *testing.T) {

	ta := require.New(t)

	cases := []struct {
		a, b, w uint64
		want    uint64
	}{
		{1, 1, 0, 1},
		{2, 3, 1, 3}, // 10 x 11
		{3, 3, 1, 4},
		{5, 3, 2, 3},           // 101 x 011
		{5, 5, 2, 6},           // 101 x 101
		{0xb9, 0xe2, 7, 0x145}, // 10111001 x 11100010 = 101000101
	}

	for i, c := range cases {
		got := shiftMulti(c.a, c.b, c.w)
		ta.Equal(c.want, got, "%d-th: case: %+v", i+1, c)
	}
}

var (
	A, B, C, D uint64
)

func BenchmarkMult(b *testing.B) {

	A = 0x112233b9
	B = 0x445566e2
	B = 0x40000601
	// B = 0x10001
	C = 32

	var s uint64
	for i := 0; i < b.N; i++ {
		s += shiftMulti(A, B, C)
	}
	D = s
}
