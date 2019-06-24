package bitmap

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFirstLast(t *testing.T) {

	ta := require.New(t)

	ta.Panics(func() { Last([]uint64{}, 0, 1) })

	cases := []struct {
		input     []uint64
		from, to  int32
		wantfirst int32
		wantlast  int32
	}{
		{nil, 0, 0, 0, -1},
		{[]uint64{}, 0, 0, 0, -1},
		{[]uint64{0}, 0, 0, 0, -1},
		{[]uint64{0}, 0, 1, 1, -1},
		{[]uint64{0}, 0, 64, 64, -1},
		{[]uint64{0}, 1, 1, 1, 0},
		{[]uint64{0}, 1, 2, 2, 0},
		{[]uint64{0, 0}, 0, 0, 0, -1},
		{[]uint64{0, 0}, 1, 65, 65, 0},
		{[]uint64{0, 0}, 1, 128, 128, 0},
		{[]uint64{1}, 0, 1, 0, 0},
		{[]uint64{1}, 0, 2, 0, 0},
		{[]uint64{1}, 1, 2, 2, 0},
		{[]uint64{2}, 1, 6, 1, 1},
		{[]uint64{2}, 2, 6, 6, 1},
		{[]uint64{4}, 1, 6, 2, 2},
		{[]uint64{4}, 2, 6, 2, 2},
		{[]uint64{1 << 63}, 0, 63, 63, -1},
		{[]uint64{1 << 63}, 0, 64, 63, 63},
		{[]uint64{0, 2}, 0, 128, 65, 65},
		{[]uint64{0, 2, 0}, 0, 128, 65, 65},
		{[]uint64{0, 2, 0, 0}, 0, 128, 65, 65},
		{[]uint64{0, 0xf0, 0xf0, 0}, 0, 128, 68, 71},
		{[]uint64{0, 0xf0, 0xf0, 0}, 69, 71, 69, 70},
	}

	for i, c := range cases {
		gotlast := Last(c.input, c.from, c.to)
		ta.Equal(c.wantlast, gotlast, "%d-th: case: %+v", i+1, c)

		gotfirst := First(c.input, c.from, c.to)
		ta.Equal(c.wantfirst, gotfirst, "%d-th: case: %+v", i+1, c)
	}
}
