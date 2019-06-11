package sigbits

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCountPrefixes(t *testing.T) {

	ta := require.New(t)

	cases := []struct {
		keys    []string
		to      int32
		wantmin int32
		want    []int32
	}{
		// a= 0110 0001
		// b= 0110 0011

		{[]string{"a", "b"}, 1,
			6, []int32{1}},

		{[]string{"a", "b"}, 8,
			6, []int32{1, 2, 2, 2,
				2, 2, 2, 2}},

		{[]string{"a", "b"}, 16,
			6, []int32{1, 2, 2, 2,
				2, 2, 2, 2,
				2, 2, 2, 2,
				2, 2, 2, 2}},

		{[]string{"aa", "ab"}, 1,
			14, []int32{1}},

		{[]string{"aa", "ab"}, 8,
			14, []int32{1, 2, 2, 2,
				2, 2, 2, 2}},

		{[]string{"aa", "ab"}, 16,
			14, []int32{1, 2, 2, 2,
				2, 2, 2, 2,
				2, 2, 2, 2,
				2, 2, 2, 2}},

		{[]string{"aa", "ab"}, 1,
			14, []int32{1}},
	}

	for i, c := range cases {
		firstdiffs := FirstDiffBits(c.keys)
		gotmin, got := countPrefixes(firstdiffs, c.to)
		ta.Equal(c.wantmin, gotmin, "%d-th: case: %+v", i+1, c)
		ta.Equal(c.want, got, "%d-th: case: %+v", i+1, c)
	}
}

func TestCountPrefixes_2uint64(t *testing.T) {

	ta := require.New(t)

	keys := []string{
		"aa",
		"ab",
		"abc",
		"abd",
		"ac",
		"b",
		"bbbbbbbb",
		"bbbbbbbba",
		"bbbbbbbbaa",
		"bbbbbbbbab",
		"bbbbbbbbc",
		"c",
	}

	wantfull := []int32{1,
		// 1~6 are empty
		2, 3,
		4, 4, 4, 4,
		4, 4, 5, 6,
		7, 7, 7, 7,
		7, 8, 8, 8,
		8, 8, 8, 8,
		8, 8, 8, 8,

		// 4-th byte
		8, 8, 8, 8,
		8, 8, 8, 8,
		8, 8, 8, 8,
		8, 8, 8, 8,
		8, 8, 8, 8,
		8, 8, 8, 8,
		8, 8, 8, 8,
		8, 8, 8, 8,

		// 8-th byte
		9, 9, 9, 9,
		9, 9, 10, 10,
		11, 11, 11, 11,
		11, 11, 12, 12,
		12, 12, 12, 12,
		12, 12, 12, 12,
		12, 12, 12, 12,
		12, 12, 12, 12,
	}

	for i := 1; i < len(wantfull)+1; i++ {
		firstdiffs := FirstDiffBits(keys)
		gotmin, got := countPrefixes(firstdiffs, int32(i))
		want := wantfull[:i]

		ta.Equal(int32(6), gotmin)
		ta.Equal(want, got, "%d-th: case: %+v", i, i)
	}
}
