package sigbits

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEOLs(t *testing.T) {

	ta := require.New(t)

	keys := []string{
		"aa",         // 16,
		"ab",         // 16,
		"abc",        // 24,
		"abd",        // 24,
		"ac",         // 16,
		"b",          // 8,
		"bbbbbbbb",   // 64,
		"bbbbbbbba",  // 72,
		"bbbbbbbbaa", // 80,
		"bbbbbbbbab", // 80,
		"bbbbbbbbc",  // 72,
		"c",          // 8,
	}

	l := int32(len(keys))
	sb := New(keys)

	cases := []struct {
		s, e, frombit, tobit int32
		dedup                bool
		want                 []int32
	}{

		{0, l, 0, -1, false, []int32{8, 8, 16, 16, 16, 24, 24, 64, 72, 72, 80, 80}},
		{1, l, 0, -1, false, []int32{8, 8, 16, 16, 24, 24, 64, 72, 72, 80, 80}},
		{1, l - 1, 0, -1, false, []int32{8, 16, 16, 24, 24, 64, 72, 72, 80, 80}},
		{1, l - 1, 0, -1, true, []int32{8, 16, 24, 64, 72, 80}},
		{l, l, 0, -1, false, []int32{}},
		{0, 0, 0, -1, false, []int32{}},

		{l, l, 0, -1, true, []int32{}},
		{0, 0, 0, -1, true, []int32{}},
		{5, l - 3, 0, -1, true, []int32{8, 64, 72, 80}},
		{5, l - 3, 5, -1, true, []int32{3, 59, 67, 75}},
		{5, l - 3, 12, -1, true, []int32{-4, 52, 60, 68}},

		{5, l - 3, 0, 32, true, []int32{8, 32}},
		{5, l - 3, 0, 70, true, []int32{8, 64, 70}},
		{5, l - 3, 5, 32, true, []int32{3, 27}},
		{5, l - 3, 12, 75, true, []int32{-4, 52, 60, 63}},
	}

	for i, c := range cases {
		got := sb.EOLs(c.s, c.e, c.frombit, c.tobit, c.dedup)
		ta.Equal(c.want, got, "%d-th: case: %+v", i+1, c)
	}
}
