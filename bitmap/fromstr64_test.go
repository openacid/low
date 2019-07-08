package bitmap

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFromStr64(t *testing.T) {

	ta := require.New(t)

	abcdabcd := uint64(0x6162636461626364)
	abcdabc := uint64(0x61626364616263)
	dabcdabc := uint64(0x6461626364616263)

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

		{"abcdabcdabcd", 0, 64, 64, (abcdabcd)},
		{"abcdabcdabcd", 1, 65, 64, (abcdabcd << 1)},
		{"abcdabcdabcd", 2, 66, 64, (abcdabcd<<2 | abcdabcd>>62)},
		{"abcdabcdabcd", 3, 67, 64, (abcdabcd<<3 | abcdabcd>>61)},
		{"abcdabcdabcd", 4, 68, 64, (abcdabcd<<4 | abcdabcd>>60)},
		{"abcdabcdabcd", 5, 69, 64, (abcdabcd<<5 | abcdabcd>>59)},
		{"abcdabcdabcd", 6, 70, 64, (abcdabcd<<6 | abcdabcd>>58)},
		{"abcdabcdabcd", 7, 71, 64, (abcdabcd<<7 | abcdabcd>>57)},
		{"abcdabcdabcd", 8, 72, 64, (abcdabcd<<8 | abcdabcd>>56)},
		{"abcdabcdabcd", 9, 73, 64, (abcdabcd<<9 | abcdabcd>>55)},

		{"abcdabcdabcd", 0, 56, 56, (abcdabc)},
		{"abcdabcdabcd", 1, 57, 56, (abcdabc<<1 | dabcdabc>>63) & ((1 << 56) - 1)},
		{"abcdabcdabcd", 2, 58, 56, (abcdabc<<2 | dabcdabc>>62) & ((1 << 56) - 1)},
		{"abcdabcdabcd", 3, 59, 56, (abcdabc<<3 | dabcdabc>>61) & ((1 << 56) - 1)},
		{"abcdabcdabcd", 4, 60, 56, (abcdabc<<4 | dabcdabc>>60) & ((1 << 56) - 1)},
		{"abcdabcdabcd", 5, 61, 56, (abcdabc<<5 | dabcdabc>>59) & ((1 << 56) - 1)},
		{"abcdabcdabcd", 6, 62, 56, (abcdabc<<6 | dabcdabc>>58) & ((1 << 56) - 1)},
		{"abcdabcdabcd", 7, 63, 56, (abcdabc<<7 | dabcdabc>>57) & ((1 << 56) - 1)},
		{"abcdabcdabcd", 8, 64, 56, (abcdabc<<8 | dabcdabc>>56) & ((1 << 56) - 1)},
		{"abcdabcdabcd", 9, 65, 56, (abcdabc<<9 | dabcdabc>>55) & ((1 << 56) - 1)},

		{"abcdabcdabcd", 0, 57, 57, (abcdabc << 1)},
		{"abcdabcdabcd", 1, 58, 57, (abcdabc<<2 | dabcdabc>>62) & ((1 << 57) - 1)},
		{"abcdabcdabcd", 2, 59, 57, (abcdabc<<3 | dabcdabc>>61) & ((1 << 57) - 1)},
		{"abcdabcdabcd", 3, 60, 57, (abcdabc<<4 | dabcdabc>>60) & ((1 << 57) - 1)},
		{"abcdabcdabcd", 4, 61, 57, (abcdabc<<5 | dabcdabc>>59) & ((1 << 57) - 1)},
		{"abcdabcdabcd", 5, 62, 57, (abcdabc<<6 | dabcdabc>>58) & ((1 << 57) - 1)},
		{"abcdabcdabcd", 6, 63, 57, (abcdabc<<7 | dabcdabc>>57) & ((1 << 57) - 1)},
		{"abcdabcdabcd", 7, 64, 57, (abcdabc<<8 | dabcdabc>>56) & ((1 << 57) - 1)},
		{"abcdabcdabcd", 8, 65, 57, (abcdabc<<9 | dabcdabc>>55) & ((1 << 57) - 1)},
		{"abcdabcdabcd", 9, 66, 57, (abcdabc<<10 | dabcdabc>>54) & ((1 << 57) - 1)},
	}

	for i, c := range cases {
		gotlen, gotbm := FromStr64(c.input, c.from, c.to)
		ta.Equal(c.wantlen, gotlen, "%d-th: case: %+v", i+1, c)
		ta.Equal(c.wantbm, gotbm, "%d-th: case: %+v", i+1, c)
	}
}

var (
	FromStr64Output int
	FromStr64Input  string
)

func BenchmarkFromStr64(b *testing.B) {
	FromStr64Input = "abcdabcdabcd"
	var s uint64 = 0
	for i := 0; i < b.N; i++ {
		_, b := FromStr64(FromStr64Input, 9, 71)
		s += b
	}
	FromStr64Output = int(s)
}
