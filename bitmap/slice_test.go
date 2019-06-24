package bitmap

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSlice(t *testing.T) {

	ta := require.New(t)

	ta.Panics(func() { Slice(Of([]int32{0}), 0, 65) })
	ta.Panics(func() { Slice(Of([]int32{0}), 64, 65) })

	cases := []struct {
		input    []int32
		from, to int32
		wantarr  []int32
	}{
		{
			[]int32{},
			0, 0,
			[]int32{},
		},
		{
			[]int32{0},
			0, 0,
			[]int32{},
		},
		{
			[]int32{0},
			0, 1,
			[]int32{0},
		},
		{
			[]int32{0},
			0, 2,
			[]int32{0},
		},
		{
			[]int32{0},
			0, 64,
			[]int32{0},
		},
		{
			[]int32{0, 1, 2},
			0, 3,
			[]int32{0, 1, 2},
		},
		{
			[]int32{0, 1, 2},
			0, 64,
			[]int32{0, 1, 2},
		},
		{
			[]int32{0, 1, 2},
			1, 64,
			[]int32{0, 1},
		},
		{
			[]int32{0, 1, 2},
			2, 64,
			[]int32{0},
		},
		{
			[]int32{0, 1, 2},
			3, 64,
			[]int32{},
		},
		{
			[]int32{0, 1, 2},
			4, 64,
			[]int32{},
		},
		{
			[]int32{64, 66},
			63, 67,
			[]int32{1, 3},
		},
		{
			[]int32{64, 66},
			64, 67,
			[]int32{0, 2},
		},
	}

	for i, c := range cases {

		words := Of(c.input)
		sl := Slice(words, c.from, c.to)
		gotarr := ToArray(sl)

		ta.Equal(c.wantarr, gotarr, "%d-th: case: %+v", i+1, c)
	}
}
