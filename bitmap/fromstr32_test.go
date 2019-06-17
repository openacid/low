package bitmap

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFromStr32(t *testing.T) {

	ta := require.New(t)

	abcdabcd := uint64(0x6162636461626364)
	mask24 := (uint64(1) << 24) - 1
	mask25 := (uint64(1) << 25) - 1
	mask32 := (uint64(1) << 32) - 1

	cases := []struct {
		input    string
		from, to int32
		wantlen  int32
		wantbm   uint64
	}{
		// a: 0110 0001
		{"", 0, 0, 0, 0},
		{"", 0, 1, 0, 0},
		{"", 0, 32, 0, 0},
		{"a", 0, 0, 0, 0},
		{"a", 0, 1, 1, 0x00},
		{"a", 0, 2, 2, 0x01},
		{"a", 0, 5, 5, 0x0c},
		{"a", 0, 6, 6, 0x18},
		{"a", 0, 7, 7, 0x30},
		{"a", 0, 8, 8, 0x61},
		{"a", 2, 2, 0, 0x00},
		{"a", 2, 3, 1, 0x01},
		{"a", 2, 4, 2, 0x02},
		{"a", 2, 7, 5, 0x10},
		{"a", 2, 8, 6, 0x21},
		{"a", 2, 10, 6, 0x84},

		// "to" over string end
		{"abc", 10, 42, 14, 0x2263 << 18},
		{"abc", 23, 31, 1, 1 << 7},
		{"abcdefgh", 31, 63, 32, (0x65666768 >> 1)},
		{"abcdefgh", 32, 64, 32, 0x65666768},
		{"abcdefgh", 33, 65, 31, (0x65666768 << 1)},

		{"abcdabcdabcd", 0, 32, 32, (abcdabcd >> 32) & mask32},
		{"abcdabcdabcd", 1, 33, 32, (abcdabcd >> 31) & mask32},
		{"abcdabcdabcd", 2, 34, 32, (abcdabcd >> 30) & mask32},
		{"abcdabcdabcd", 3, 35, 32, (abcdabcd >> 29) & mask32},
		{"abcdabcdabcd", 4, 36, 32, (abcdabcd >> 28) & mask32},
		{"abcdabcdabcd", 5, 37, 32, (abcdabcd >> 27) & mask32},
		{"abcdabcdabcd", 6, 38, 32, (abcdabcd >> 26) & mask32},
		{"abcdabcdabcd", 7, 39, 32, (abcdabcd >> 25) & mask32},
		{"abcdabcdabcd", 8, 40, 32, (abcdabcd >> 24) & mask32},
		{"abcdabcdabcd", 9, 41, 32, (abcdabcd >> 23) & mask32},

		{"abcdabcdabcd", 0, 24, 24, (abcdabcd >> 40) & mask24},
		{"abcdabcdabcd", 1, 25, 24, (abcdabcd >> 39) & mask24},
		{"abcdabcdabcd", 2, 26, 24, (abcdabcd >> 38) & mask24},
		{"abcdabcdabcd", 3, 27, 24, (abcdabcd >> 37) & mask24},
		{"abcdabcdabcd", 4, 28, 24, (abcdabcd >> 36) & mask24},
		{"abcdabcdabcd", 5, 29, 24, (abcdabcd >> 35) & mask24},
		{"abcdabcdabcd", 6, 30, 24, (abcdabcd >> 34) & mask24},
		{"abcdabcdabcd", 7, 31, 24, (abcdabcd >> 33) & mask24},
		{"abcdabcdabcd", 8, 32, 24, (abcdabcd >> 32) & mask24},
		{"abcdabcdabcd", 9, 33, 24, (abcdabcd >> 31) & mask24},

		{"abcdabcdabcd", 0, 25, 25, (abcdabcd >> 39) & mask25},
		{"abcdabcdabcd", 1, 26, 25, (abcdabcd >> 38) & mask25},
		{"abcdabcdabcd", 2, 27, 25, (abcdabcd >> 37) & mask25},
		{"abcdabcdabcd", 3, 28, 25, (abcdabcd >> 36) & mask25},
		{"abcdabcdabcd", 4, 29, 25, (abcdabcd >> 35) & mask25},
		{"abcdabcdabcd", 5, 30, 25, (abcdabcd >> 34) & mask25},
		{"abcdabcdabcd", 6, 31, 25, (abcdabcd >> 33) & mask25},
		{"abcdabcdabcd", 7, 32, 25, (abcdabcd >> 32) & mask25},
		{"abcdabcdabcd", 8, 33, 25, (abcdabcd >> 31) & mask25},
		{"abcdabcdabcd", 9, 34, 25, (abcdabcd >> 30) & mask25},
	}

	for i, c := range cases {
		gotlen, gotbm := FromStr32(c.input, c.from, c.to)
		ta.Equal(c.wantlen, gotlen, "%d-th: case: %+v", i+1, c)
		ta.Equal(c.wantbm, gotbm, "%d-th: case: %+v", i+1, c)
	}
}

var (
	FromStr32Output int
	FromStr32Input  string
)

func BenchmarkFromStr32(b *testing.B) {
	FromStr32Input = "abcdabcdabcd"
	var s uint64 = 0
	for i := 0; i < b.N; i++ {
		_, b := FromStr32(FromStr32Input, 9, 41)
		s += b
	}
	FromStr32Output = int(s)
}

func BenchmarkStringToBytes(b *testing.B) {
	FromStr32Input = "abcdabcdabcd"
	var s uint64 = 0
	for i := 0; i < b.N; i++ {
		b := []byte(FromStr32Input)
		s += uint64(b[2])
	}
	FromStr32Output = int(s)
}
