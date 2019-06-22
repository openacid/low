package bitmap

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOfMany(t *testing.T) {

	ta := require.New(t)

	cases := []struct {
		input      [][]int32
		inputsizes []int32
		wantwords  []uint64
	}{
		{
			[][]int32{},
			[]int32{},
			[]uint64{},
		},
		{
			[][]int32{{0}},
			[]int32{1},
			[]uint64{1},
		},
		{
			[][]int32{{0}},
			[]int32{2},
			[]uint64{1},
		},
		{
			[][]int32{{0}},
			[]int32{64},
			[]uint64{1},
		},
		{
			[][]int32{{0}},
			[]int32{65},
			[]uint64{1, 0},
		},
		{
			[][]int32{{0, 1, 2}},
			[]int32{0},
			[]uint64{7},
		},
		{
			[][]int32{{0, 1, 2}},
			[]int32{1},
			[]uint64{7},
		},
		{
			[][]int32{{0}, {0, 1}},
			[]int32{1, 2},
			[]uint64{7},
		},
		{
			[][]int32{{0}, {1, 2}},
			[]int32{1, 2},
			[]uint64{13},
		},
		{
			[][]int32{{0}, {1, 2}},
			[]int32{5, 2},
			[]uint64{193},
		},
		{
			[][]int32{{0}, {1, 2}},
			[]int32{64, 64},
			[]uint64{1, 6},
		},
		{
			[][]int32{{0, 1}, {2, 63}},
			[]int32{64, 0},
			[]uint64{3, (1 << 63) + 4},
		},
		{
			[][]int32{{64}},
			[]int32{0},
			[]uint64{0, 1},
		},
		{
			[][]int32{{64}, {0}},
			[]int32{65, 1},
			[]uint64{0, 3},
		},
		{
			[][]int32{{1}, {2}, {3, 64, 129}},
			[]int32{1, 1, 129},
			[]uint64{42, 4, 8},
		},
		{ // extend to contain all bitmap
			[][]int32{{1}, {2}},
			[]int32{1, 127},
			[]uint64{0xa, 0},
		},
		{ // extend to contain all bitmap
			[][]int32{{1}, {2}},
			[]int32{1, 128},
			[]uint64{0xa, 0, 0},
		},
	}

	for i, c := range cases {

		got := OfMany(c.input, c.inputsizes)

		ta.Equal(c.wantwords, got,
			"%d-th: case: %+v",
			i+1, c)
	}
}
