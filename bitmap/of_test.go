package bitmap

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOf(t *testing.T) {

	ta := require.New(t)

	cases := []struct {
		bitPositions []int32
		wantwords    []uint64
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

		got := Of(c.bitPositions)

		ta.Equal(c.wantwords, got,
			"%d-th: case: %+v",
			i+1, c)
	}
}

func TestOf_with_n(t *testing.T) {

	ta := require.New(t)

	cases := []struct {
		bitPositions []int32
		n            int32
		wantwords    []uint64
	}{
		// 0 bit
		{
			[]int32{}, 0,
			[]uint64{},
		},
		{
			[]int32{}, -1,
			[]uint64{},
		},
		{
			[]int32{}, 1,
			[]uint64{0},
		},
		{
			[]int32{}, 63,
			[]uint64{0},
		},
		{
			[]int32{}, 64,
			[]uint64{0},
		},
		{
			[]int32{}, 65,
			[]uint64{0, 0},
		},
		// 1 bit
		{
			[]int32{0}, 0,
			[]uint64{1},
		},
		{
			[]int32{0}, -1,
			[]uint64{1},
		},
		{
			[]int32{0}, 1,
			[]uint64{1},
		},
		{
			[]int32{0}, 2,
			[]uint64{1},
		},
		{
			[]int32{0}, 63,
			[]uint64{1},
		},
		{
			[]int32{0}, 64,
			[]uint64{1},
		},
		{
			[]int32{0}, 65,
			[]uint64{1, 0},
		},

		// more than one bit
		{
			[]int32{0, 1, 2}, 0,
			[]uint64{7},
		},
		{
			[]int32{0, 1, 2}, -1,
			[]uint64{7},
		},
		{
			[]int32{0, 1, 2}, 63,
			[]uint64{7},
		},
		{
			[]int32{0, 1, 2}, 64,
			[]uint64{7},
		},
		{
			[]int32{0, 1, 2}, 65,
			[]uint64{7, 0},
		},
		{
			[]int32{0, 1, 2}, 127,
			[]uint64{7, 0},
		},
		{
			[]int32{0, 1, 2}, 128,
			[]uint64{7, 0},
		},
		{
			[]int32{0, 1, 2}, 129,
			[]uint64{7, 0, 0},
		},

		// 1 bit in 2nd word
		{
			[]int32{64}, 0,
			[]uint64{0, 1},
		},
		{
			[]int32{64}, -1,
			[]uint64{0, 1},
		},
		{
			[]int32{64}, 1,
			[]uint64{0, 1},
		},
		{
			[]int32{64}, 65,
			[]uint64{0, 1},
		},
		{
			[]int32{64}, 127,
			[]uint64{0, 1},
		},
		{
			[]int32{64}, 128,
			[]uint64{0, 1},
		},
		{
			[]int32{64}, 129,
			[]uint64{0, 1, 0},
		},
	}

	for i, c := range cases {

		got := Of(c.bitPositions, c.n)

		ta.Equal(c.wantwords, got,
			"%d-th: case: %+v",
			i+1, c)
	}
}
