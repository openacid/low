package bitmap

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFmt(t *testing.T) {

	ta := require.New(t)

	cases := []struct {
		input interface{}
		want  string
	}{
		{int8(7), "11100000"},
		{uint8(7), "11100000"},
		{int16(0x0507), "11100000 10100000"},
		{uint16(0x0507), "11100000 10100000"},
		{int32(0x01030507), "11100000 10100000 11000000 10000000"},
		{uint32(0x01030507), "11100000 10100000 11000000 10000000"},
		{int64(0x0f01030507), "11100000 10100000 11000000 10000000 11110000 00000000 00000000 00000000"},
		{uint64(0x0f01030507), "11100000 10100000 11000000 10000000 11110000 00000000 00000000 00000000"},

		{[]int8{7}, "11100000"},
		{[]uint16{0x0507, 0x0102}, "11100000 10100000,01000000 10000000"},
	}

	for i, c := range cases {
		got := Fmt(c.input)
		ta.Equal(c.want, got,
			"%d-th: input: %#v; want: %#v; got: %#v",
			i+1, c.input, c.want, got)

	}
}
