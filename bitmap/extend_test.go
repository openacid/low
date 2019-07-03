package bitmap

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExtend(t *testing.T) {

	ta := require.New(t)

	cases := []struct {
		input []uint64
		n     int32
		want  []uint64
	}{
		{nil, -1, nil},
		{nil, 0, nil},
		{nil, 1, []uint64{0}},
		{[]uint64{}, -1, []uint64{}},
		{[]uint64{}, 0, []uint64{}},
		{[]uint64{}, 1, []uint64{0}},
		{[]uint64{}, 63, []uint64{0}},
		{[]uint64{}, 64, []uint64{0}},
		{[]uint64{}, 65, []uint64{0, 0}},
		{[]uint64{123}, 65, []uint64{123, 0}},
		{[]uint64{123, 0}, 65, []uint64{123, 0}},
		{[]uint64{123, 123}, 65, []uint64{123, 123}},
	}

	for i, c := range cases {
		got := Extend(c.input, c.n)
		ta.Equal(c.want, got, "%d-th: case: %+v", i+1, c)
	}
}
