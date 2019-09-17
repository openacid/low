package bitmap

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIndexRank(t *testing.T) {

	ta := require.New(t)

	cases := []struct {
		bm      []uint64
		want64  []int32
		want128 []int32
	}{
		{
			[]uint64{},
			[]int32{0},
			[]int32{0},
		},
		{
			[]uint64{0},
			[]int32{0, 0},
			[]int32{0},
		},
		{
			[]uint64{1},
			[]int32{0, 1},
			[]int32{0},
		},
		{
			[]uint64{0xffffffffffffffff},
			[]int32{0, 64},
			[]int32{0},
		},
		{
			[]uint64{0xffffffffffffffff, 1},
			[]int32{0, 64, 65},
			[]int32{0, 65},
		},
		{
			[]uint64{0xffffffffffffffff, 1, 1},
			[]int32{0, 64, 65, 66},
			[]int32{0, 65},
		},
		{
			[]uint64{0xffffffffffffffff, 1, 1, 3},
			[]int32{0, 64, 65, 66, 68},
			[]int32{0, 65, 68},
		},
		{
			[]uint64{0xffffffffffffffff, 1, 1, 3, 4},
			[]int32{0, 64, 65, 66, 68, 69},
			[]int32{0, 65, 68},
		},
	}

	for i, c := range cases {

		idx64 := IndexRank64(c.bm)
		ta.Equal(c.want64, idx64, "%d-th: case: %+v", i+1, c)

		idx128 := IndexRank128(c.bm)
		ta.Equal(c.want128, idx128, "%d-th: case: %+v", i+1, c)

		// test Rank64 and Rank128
		cnt := int32(0)
		cntExcludeI := int32(0)
		for j := 0; j < len(c.bm)*64; j++ {
			if c.bm[j>>6]&(1<<uint(j&63)) != 0 {
				cnt++
			}

			rExc, isSet := Rank64(c.bm, idx64, int32(j))
			ta.Equal(cntExcludeI, rExc, "bm: %+v, idx64:%+v j:%d", c.bm, idx64, j)
			ta.Equal(cnt-cntExcludeI, isSet, "bm: %+v, idx64:%+v j:%d", c.bm, idx64, j)

			rExc, isSet = Rank128(c.bm, idx128, int32(j))
			ta.Equal(cntExcludeI, rExc)
			ta.Equal(cnt-cntExcludeI, isSet)

			cntExcludeI = cnt
		}
	}
}

func TestRank_panic(t *testing.T) {

	ta := require.New(t)

	nums := []int32{1, 3, 64, 129}
	bm := Of(nums)

	idx64 := IndexRank64(bm)
	idx128 := IndexRank64(bm)

	ta.Panics(func() { Rank64(bm, idx64, -1) })
	ta.Panics(func() { Rank128(bm, idx128, -1) })

	// no panic
	_, _ = Rank64(bm, idx64, 130)
	_, _ = Rank128(bm, idx128, 191)

	ta.Panics(func() { Rank64(bm, idx64, 192) })
	ta.Panics(func() { Rank128(bm, idx128, 192) })
}
