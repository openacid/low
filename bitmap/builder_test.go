package bitmap

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBuilder_Extend(t *testing.T) {

	ta := require.New(t)

	cases := []struct {
		input      [][]int32
		inputSizes []int32
		wantWords  []uint64
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

		{
			b := NewBuilder(0)
			for i, sz := range c.inputSizes {
				pos := c.input[i]
				b.Extend(pos, sz)
			}

			ta.Equal(c.wantWords, b.Words,
				"%d-th: case: %+v",
				i+1, c)
		}

		{

			b := NewBuilder(100)
			for i, sz := range c.inputSizes {
				pos := c.input[i]
				b.Extend(pos, sz)
			}

			ta.Equal(c.wantWords, b.Words,
				"%d-th: case: %+v",
				i+1, c)
		}
	}
}

func TestBuilder_Set(t *testing.T) {

	ta := require.New(t)

	cases := []struct {
		input      []int32
		value      []int32
		wantWords  []uint64
		wantOffset int32
	}{
		{
			[]int32{0},
			[]int32{1},
			[]uint64{1},
			1,
		},
		{
			[]int32{0, 5},
			[]int32{1, 0},
			[]uint64{1},
			6,
		},
		{
			[]int32{0, 5, 63},
			[]int32{1, 0, 1},
			[]uint64{1 + 1<<63},
			64,
		},
		{
			[]int32{0, 5, 63, 5},
			[]int32{1, 0, 1, 1},
			[]uint64{1 + 1<<5 + 1<<63},
			64,
		},
		{
			[]int32{0, 5, 63, 5, 64},
			[]int32{1, 0, 1, 1, 0},
			[]uint64{1 + 1<<5 + 1<<63, 0},
			65,
		},
		{
			[]int32{0, 5, 63, 5, 64},
			[]int32{1, 0, 1, 1, 1},
			[]uint64{1 + 1<<5 + 1<<63, 1},
			65,
		},
	}

	for i, c := range cases {

		b := NewBuilder(0)
		for i, pos := range c.input {
			b.Set(pos, c.value[i])
		}

		ta.Equal(c.wantWords, b.Words,
			"%d-th: words, case: %+v",
			i+1, c)
		ta.Equal(c.wantOffset, b.Offset,
			"%d-th: offset, case: %+v",
			i+1, c)
	}
}
