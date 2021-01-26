package bitstr

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {

	ta := require.New(t)

	cases := []struct {
		srcStr         string
		fromBit, toBit int32
		want           string
	}{
		{"abc", 0, 16 + 8, "01100001 01100010 01100011 11111111"},
		{"abc", 0, 16 + 7, "01100001 01100010 01100010 11111110"},
		{"abc", 0, 16 + 6, "01100001 01100010 01100000 11111100"},
		{"abc", 0, 16 + 5, "01100001 01100010 01100000 11111000"},
		{"abc", 0, 16 + 4, "01100001 01100010 01100000 11110000"},
		{"abc", 0, 16 + 3, "01100001 01100010 01100000 11100000"},
		{"abc", 0, 16 + 2, "01100001 01100010 01000000 11000000"},
		{"abc", 0, 16 + 1, "01100001 01100010 00000000 10000000"},

		{"abc", 0, 8 + 8, "01100001 01100010 11111111"},
		{"abc", 0, 8 + 7, "01100001 01100010 11111110"},
		{"abc", 0, 8 + 6, "01100001 01100000 11111100"},
		{"abc", 0, 8 + 5, "01100001 01100000 11111000"},
		{"abc", 0, 8 + 4, "01100001 01100000 11110000"},
		{"abc", 0, 8 + 3, "01100001 01100000 11100000"},
		{"abc", 0, 8 + 2, "01100001 01000000 11000000"},
		{"abc", 0, 8 + 1, "01100001 00000000 10000000"},

		{"abc", 0, 0 + 8, "01100001 11111111"},
		{"abc", 0, 0 + 7, "01100000 11111110"},
		{"abc", 0, 0 + 6, "01100000 11111100"},
		{"abc", 0, 0 + 5, "01100000 11111000"},
		{"abc", 0, 0 + 4, "01100000 11110000"},
		{"abc", 0, 0 + 3, "01100000 11100000"},
		{"abc", 0, 0 + 2, "01000000 11000000"},
		{"abc", 0, 0 + 1, "00000000 10000000"},

		{"abc", 0, 0, "11111111"},

		// fromBit will be truncated to align-to-8.
		{"abc", 1, 0 + 6, "01100000 11111100"},
		{"abc", 2, 0 + 6, "01100000 11111100"},
		{"abc", 3, 0 + 6, "01100000 11111100"},
		{"abc", 4, 0 + 6, "01100000 11111100"},
		{"abc", 5, 0 + 6, "01100000 11111100"},
		{"abc", 6, 0 + 6, "01100000 11111100"},

		// from the 1-th byte:
		{"abc", 8 + 3, 8 + 6, "01100000 11111100"},
	}

	for i, c := range cases {
		got := New(c.srcStr, c.fromBit, c.toBit)
		gotStr := binFmt(got)

		ta.Equal(c.want, gotStr, "%d-th: case: %+v", i+1, c)

		gotLen := Len(got)
		ta.Equal(c.toBit-c.fromBit&^7, gotLen, "%d-th: case: %+v", i+1, c)
	}
}

func TestCmp(t *testing.T) {

	ta := require.New(t)

	{ // prefix compare
		s := "abc"
		l := int32(len(s) << 3)
		for fromBit := int32(0); fromBit < l; fromBit++ {
			for toBit := fromBit; toBit < l-1; toBit++ {
				a := New(s, fromBit, toBit)
				b := New(s, fromBit, toBit+1)

				got := Cmp(a, b)

				ta.Equal(-1, got, "a smaller one is always smaller than a longer one: from %d to %d", fromBit, toBit)
			}
		}
	}

	{ // other cases

		cases := []struct {
			a          string
			aFrom, aTo int32
			b          string
			bFrom, bTo int32
			want       int
		}{
			{ // c = 0110 0011
				"abc", 0, 16 + 5,
				"abd", 0, 16 + 5,
				0,
			},
			{ // c = 0110 0011
				"abc", 0, 16 + 6,
				"abd", 0, 16 + 6,
				-1,
			},
			{ // b = 0110 0010
				"ab", 0, 8 + 7,
				"aca", 0, 16 + 6,
				-1,
			},
			{ // b = 0110 0010
				"ab", 0, 8 + 7,
				"aca", 0, 16 + 4,
				-1,
			},
			{
				"abc", 0, 16 + 8,
				"bcd", 0, 16 + 8,
				-1,
			},
		}

		for i, c := range cases {

			a := New(c.a, c.aFrom, c.aTo)
			b := New(c.b, c.bFrom, c.bTo)

			got := Cmp(a, b)
			ta.Equal(c.want, got, "%d-th: case: %+v, a: %s, b: %s", i+1, c, binFmt(a), binFmt(b))
		}
	}
}

func TestCmpUpto(t *testing.T) {

	ta := require.New(t)

	cases := []struct {
		key  string
		src  string
		s, e int32
		want int
	}{
		// original str
		{"abc", "abc", 0, 0, 0},
		{"abc", "abc", 0, 1, 0},
		{"abc", "abc", 0, 2, 0},
		{"abc", "abc", 0, 3, 0},
		{"abc", "abc", 0, 4, 0},
		{"abc", "abc", 0, 5, 0},
		{"abc", "abc", 0, 6, 0},
		{"abc", "abc", 0, 7, 0},
		{"abc", "abc", 0, 8, 0},
		{"abc", "abc", 0, 9, 0},
		{"abc", "abc", 0, 10, 0},
		{"abc", "abc", 0, 11, 0},
		{"abc", "abc", 0, 12, 0},
		{"abc", "abc", 0, 13, 0},
		{"abc", "abc", 0, 14, 0},
		{"abc", "abc", 0, 15, 0},
		{"abc", "abc", 0, 16, 0},

		{"bc", "abc", 8, 8, 0},
		{"bc", "abc", 8, 9, 0},
		{"bc", "abc", 8, 10, 0},
		{"bc", "abc", 8, 11, 0},
		{"bc", "abc", 8, 12, 0},
		{"bc", "abc", 8, 13, 0},
		{"bc", "abc", 8, 14, 0},
		{"bc", "abc", 8, 15, 0},
		{"bc", "abc", 8, 16, 0},

		// empty str
		{"", "abc", 0, 0, 0},
		{"", "abc", 0, 1, -1},
		{"", "abc", 0, 2, -1},
		{"", "abc", 0, 3, -1},
		{"", "abc", 0, 4, -1},
		{"", "abc", 0, 5, -1},
		{"", "abc", 0, 6, -1},
		{"", "abc", 0, 7, -1},
		{"", "abc", 0, 8, -1},
		{"", "abc", 0, 9, -1},
		{"", "abc", 0, 10, -1},
		{"", "abc", 0, 11, -1},
		{"", "abc", 0, 12, -1},
		{"", "abc", 0, 13, -1},
		{"", "abc", 0, 14, -1},
		{"", "abc", 0, 15, -1},
		{"", "abc", 0, 16, -1},

		{"", "abc", 8, 8, 0},
		{"", "abc", 8, 9, -1},
		{"", "abc", 8, 10, -1},
		{"", "abc", 8, 11, -1},
		{"", "abc", 8, 12, -1},
		{"", "abc", 8, 13, -1},
		{"", "abc", 8, 14, -1},
		{"", "abc", 8, 15, -1},
		{"", "abc", 8, 16, -1},

		// smaller str
		{"abc", "bcd", 0, 0, 0},
		{"abc", "bcd", 0, 1, 0},
		{"abc", "bcd", 0, 2, 0},
		{"abc", "bcd", 0, 3, 0},
		{"abc", "bcd", 0, 4, 0},
		{"abc", "bcd", 0, 5, 0},
		{"abc", "bcd", 0, 6, 0},
		{"abc", "bcd", 0, 7, -1},
		{"abc", "bcd", 0, 8, -1},
		{"abc", "bcd", 0, 9, -1},
		{"abc", "bcd", 0, 10, -1},
		{"abc", "bcd", 0, 11, -1},
		{"abc", "bcd", 0, 12, -1},
		{"abc", "bcd", 0, 13, -1},
		{"abc", "bcd", 0, 14, -1},
		{"abc", "bcd", 0, 15, -1},
		{"abc", "bcd", 0, 16, -1},

		{"bc", "bcd", 8, 8, 0},
		{"bc", "bcd", 8, 9, 0},
		{"bc", "bcd", 8, 10, 0},
		{"bc", "bcd", 8, 11, 0},
		{"bc", "bcd", 8, 12, 0},
		{"bc", "bcd", 8, 13, 0},
		{"bc", "bcd", 8, 14, 0},
		{"bc", "bcd", 8, 15, 0},
		{"bc", "bcd", 8, 16, -1},

		// greater str
		{"bcd", "abc", 0, 0, 0},
		{"bcd", "abc", 0, 1, 0},
		{"bcd", "abc", 0, 2, 0},
		{"bcd", "abc", 0, 3, 0},
		{"bcd", "abc", 0, 4, 0},
		{"bcd", "abc", 0, 5, 0},
		{"bcd", "abc", 0, 6, 0},
		{"bcd", "abc", 0, 7, 1},
		{"bcd", "abc", 0, 8, 1},
		{"bcd", "abc", 0, 9, 1},
		{"bcd", "abc", 0, 10, 1},
		{"bcd", "abc", 0, 11, 1},
		{"bcd", "abc", 0, 12, 1},
		{"bcd", "abc", 0, 13, 1},
		{"bcd", "abc", 0, 14, 1},
		{"bcd", "abc", 0, 15, 1},
		{"bcd", "abc", 0, 16, 1},

		{"cde", "abc", 8, 8, 0},
		{"cde", "abc", 8, 9, 0},
		{"cde", "abc", 8, 10, 0},
		{"cde", "abc", 8, 11, 0},
		{"cde", "abc", 8, 12, 0},
		{"cde", "abc", 8, 13, 0},
		{"cde", "abc", 8, 14, 0},
		{"cde", "abc", 8, 15, 0},
		{"cde", "abc", 8, 16, 1},
	}

	for i, c := range cases {
		bs := New(c.src, c.s, c.e)
		fmt.Println(binFmt(bs))
		{
			got := CmpUpto([]byte(c.key), bs)
			ta.Equal(c.want, got, "%d-th: case: %+v", i+1, c)
		}

		{
			gots := StrCmpUpto(c.key, bs)
			ta.Equal(c.want, gots, "%d-th: case: %+v", i+1, c)
		}
	}
}

func binFmt(s []byte) string {
	rst := make([]string, 0, len(s))
	for _, v := range s {
		rst = append(rst, fmt.Sprintf("%08b", v))
	}

	return strings.Join(rst, " ")
}

var OutputCmpUpto int

func BenchmarkCmpUpto(b *testing.B) {

	cases := []struct {
		str            string
		bitStr         string
		fromBit, toBit int32
	}{
		{"abc", "abcdef", 0, 22},
		{"abcdefghi", "abcdefghijk", 0, 60},
		{"abcdefghijkabcdef", "abcdefghijkabcdefghijk", 0, 161},
		{"abcde", "abcdefghijk", 0, 60},
	}

	for _, c := range cases {
		b.Run(
			fmt.Sprintf("%d %d", len(c.str), c.toBit-c.fromBit),
			func(b *testing.B) {
				var s int
				bs := New(c.bitStr, c.fromBit, c.toBit)
				sbs := []byte(c.str)
				for i := 0; i < b.N; i++ {
					s += CmpUpto(sbs, bs)
				}
				OutputCmpUpto = s
			})
	}
}
