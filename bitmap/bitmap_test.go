package bitmap_test

import (
	"testing"

	"github.com/openacid/low/bitmap"
	"github.com/stretchr/testify/require"
)

func TestIndexRank64(t *testing.T) {

	ta := require.New(t)

	cases := []struct {
		bm      []uint64
		want64  []int32
		want128 []int32
	}{
		{
			[]uint64{},
			[]int32{},
			[]int32{0},
		},
		{
			[]uint64{0},
			[]int32{0},
			[]int32{0},
		},
		{
			[]uint64{1},
			[]int32{0},
			[]int32{0},
		},
		{
			[]uint64{0xffffffffffffffff},
			[]int32{0},
			[]int32{0},
		},
		{
			[]uint64{0xffffffffffffffff, 1},
			[]int32{0, 64},
			[]int32{0, 65},
		},
		{
			[]uint64{0xffffffffffffffff, 1, 1},
			[]int32{0, 64, 65},
			[]int32{0, 65},
		},
		{
			[]uint64{0xffffffffffffffff, 1, 1, 3},
			[]int32{0, 64, 65, 66},
			[]int32{0, 65, 68},
		},
		{
			[]uint64{0xffffffffffffffff, 1, 1, 3, 4},
			[]int32{0, 64, 65, 66, 68},
			[]int32{0, 65, 68},
		},
	}

	for i, c := range cases {

		idx64 := bitmap.IndexRank64(c.bm)
		ta.Equal(c.want64, idx64, "%d-th: case: %+v", i+1, c)

		idx128 := bitmap.IndexRank128(c.bm)
		ta.Equal(c.want128, idx128, "%d-th: case: %+v", i+1, c)
	}
}
