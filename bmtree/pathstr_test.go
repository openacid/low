package bmtree

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStr(t *testing.T) {

	ta := require.New(t)

	cases := []struct {
		input uint64
		want  string
	}{
		{0x0000000000000000, ""},
		{0x0000000000000040, "0"},
		{0x0000000000000070, "000"},
		{0x0000007000000070, "111"},
		{0x0000006000000070, "110"},
		{0x0000000500000070, "000"},
		{0x000000050000007f, "0000101"},
		{0x0000000b0000007f, "0001011"},
		{0x00000000ffffffff, "00000000000000000000000000000000"},
		{0x00000012ffffffff, "00000000000000000000000000010010"},
		{0x00000012fffffffc, "000000000000000000000000000100"},
	}

	for i, c := range cases {
		got := PathStr(c.input)
		ta.Equal(c.want, got, "%d-th: case: %+v", i+1, c)
	}
}
