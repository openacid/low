package bmtree

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMaxbits(t *testing.T) {

	ta := require.New(t)

	cases := []struct {
		input uint64
		want  int32
	}{
		{0x0000000000000000, 0},
		{0x0000000000000040, 7},
		{0x0000000000000070, 7},
		{0x0000007000000070, 7},
		{0x0000006000000070, 7},
		{0x0000000500000070, 7},
		{0x000000050000007f, 7},
		{0x0000000b0000007f, 7},
		{0x00000000ffffffff, 32},
		{0x00000012ffffffff, 32},
		{0x00000012fffffffc, 32},
	}

	for i, c := range cases {
		got := PathHeight(c.input)
		ta.Equal(c.want, got, "%d-th: case: %+v", i+1, c)
	}
}
