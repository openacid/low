package sigbits

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEndsOf(t *testing.T) {

	ta := require.New(t)

	cases := []struct {
		keys []string
		want []int
	}{
		{[]string{}, []int{}},
		{[]string{""}, []int{0}},
		{[]string{"a"}, []int{8}},
		{[]string{"ab"}, []int{16}},

		{[]string{"a", "b"}, []int{8, 8}},
		{[]string{"aa", "ab"}, []int{16, 16}},
		{[]string{"a", "aa"}, []int{8, 16}},
		{[]string{"aa", "b"}, []int{16, 8}},

		{[]string{"aa", "aab", "aac"}, []int{16, 24, 24}},

		{[]string{
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
		}, []int{
			16,
			16,
			24,
			24,
			16,
			8,
			64,
			72,
			80,
			80,
			72,
			8,
		}},
	}

	for i, c := range cases {
		got := endsOf(c.keys)
		ta.Equal(c.want, got, "%d-th: case: %+v", i+1, c)
	}
}
