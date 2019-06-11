package sigbits

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGet64Bits(t *testing.T) {

	ta := require.New(t)

	cases := []struct {
		input string
		want  uint64
	}{
		{"", 0},
		{"a", 0x61 << 56},
		{"abc", 0x616263 << 40},
		{"abcd1234", 0x6162636431323334},
		{"abcd1234x", 0x6162636431323334},
	}

	for i, c := range cases {
		got := get64Bits(c.input)
		ta.Equal(c.want, got, "%d-th: case: %+v", i+1, c)
	}
}

var (
	Get64BitsInput  string = "abcd1234x"
	Get64BitsOutput int
)

func BenchmarkGet64Bits_9(b *testing.B) {
	Get64BitsInput = "abcd1234x"
	var s uint64 = 0
	for i := 0; i < b.N; i++ {
		s += get64Bits(Get64BitsInput)
	}
	Get64BitsOutput = int(s)
}

func BenchmarkGet64Bits_7(b *testing.B) {
	Get64BitsInput = "abcd123"
	var s uint64 = 0
	for i := 0; i < b.N; i++ {
		s += get64Bits(Get64BitsInput)
	}
	Get64BitsOutput = int(s)
}

func TestFirstDiffBit(t *testing.T) {

	ta := require.New(t)

	cases := []struct {
		a, b string
		want int32
	}{
		{"", "", 0},
		{"a", "b", 6},
		{"aa", "ab", 14},
		{"aa", "aa", 16},
		{"aa", "aab", 16},
		{"aac", "aa", 16},
		{"aac", "ab", 14},
		{"aaa", "aaa", 24},
		{"12345678", "12345678", 64},
		{"12345678a", "12345678", 64},
		{"12345678", "12345678a", 64},
		{"12345678a", "12345678a", 72},
		{"12345678a", "12345678b", 70},
		{"12345678aab", "12345678aa", 80},
		{"12345678aa", "12345678aab", 80},
		{"12345678aac", "12345678aab", 87},
	}

	for i, c := range cases {
		got := sFirstDiffBit(c.a, c.b)
		ta.Equal(c.want, got, "%d-th: case: %+v", i+1, c)
	}
}

func TestFirstDiffBits(t *testing.T) {

	ta := require.New(t)

	ta.Panics(func() { FirstDiffBits([]string{}) })

	cases := []struct {
		keys []string
		want []int32
	}{
		{[]string{""}, []int32{}},
		{[]string{"a"}, []int32{}},
		{[]string{"ab"}, []int32{}},

		{[]string{"a", "b"}, []int32{6}},
		{[]string{"aa", "ab"}, []int32{14}},
		{[]string{"a", "aa"}, []int32{8}},
		{[]string{"aa", "b"}, []int32{6}},

		{[]string{"aa", "aab", "aac"}, []int32{16, 23}},

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
		}, []int32{
			14,
			16,
			21,
			15,
			6,
			8,
			64,
			72,
			78,
			70,
			7,
		}},
	}

	for i, c := range cases {
		got := FirstDiffBits(c.keys)
		ta.Equal(c.want, got, "%d-th: case: %+v", i+1, c)
	}
}
