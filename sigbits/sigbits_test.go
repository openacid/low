package sigbits

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {

	ta := require.New(t)

	cases := []struct {
		input []string
		want  []int32
	}{
		{[]string{
			"",
			"12345678",
			"12345678a",
			"12345678aa",
			"12345678aab",
			"12345678aac",
			"a",
			"aa",
			"aaa",
			"aac",
			"ab",
			"b",
		},
			[]int32{
				0,
				64,
				72,
				80,
				87,
				1,
				8,
				16,
				22,
				14,
				6,
			},
		},
	}

	for i, c := range cases {
		got := New(c.input)
		ta.Equal(c.want, got.sigbits, "%d-th: case: %+v", i+1, c)
	}
}
