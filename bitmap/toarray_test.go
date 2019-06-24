package bitmap

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestToArray(t *testing.T) {

	ta := require.New(t)

	cases := []struct {
		input     []int32
		wantwords []uint64
	}{
		{
			[]int32{},
			[]uint64{},
		},
		{
			[]int32{0},
			[]uint64{1},
		},
		{
			[]int32{0, 1, 2},
			[]uint64{7},
		},
		{
			[]int32{0, 1, 2, 63},
			[]uint64{(1 << 63) + 7},
		},
		{
			[]int32{64},
			[]uint64{0, 1},
		},
		{
			[]int32{1, 2, 3, 64, 129},
			[]uint64{0x0e, 1, 2},
		},
	}

	for i, c := range cases {

		words := Of(c.input)

		ta.Equal(c.wantwords, words,
			"%d-th: case: %+v",
			i+1, c)

		ta.Equal(c.input, ToArray(words),
			"%d-th: case: %+v",
			i+1, c)
	}
}
