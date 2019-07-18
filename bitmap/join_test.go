package bitmap

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestJoin(t *testing.T) {

	ta := require.New(t)

	cases := []struct {
		subs      []uint64
		size      int32
		wantwords []uint64
	}{
		{[]uint64{}, 0, []uint64{}},
		{[]uint64{}, 1, []uint64{}},
		{[]uint64{}, 2, []uint64{}},

		{[]uint64{1, 2}, 1, []uint64{1}},
		{[]uint64{1, 1, 0, 1, 0, 1}, 1, []uint64{43}},
		{[]uint64{1, 2}, 2, []uint64{9}},
		{[]uint64{1, 2}, 32, []uint64{1 + 2<<32}},
		{[]uint64{1, 2, 3}, 32, []uint64{1 + 2<<32, 3}},
		{[]uint64{1, 2, 3, 4}, 32, []uint64{1 + 2<<32, 3 + 4<<32}},
	}

	for i, c := range cases {

		got := Join(c.subs, c.size)

		ta.Equal(c.wantwords, got,
			"%d-th: case: %+v",
			i+1, c)
	}
}
