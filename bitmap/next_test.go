package bitmap

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNextOne(t *testing.T) {

	ta := require.New(t)

	cases := []struct {
		bm   []uint64
		from int32
		n    int32
		want int32
	}{
		{[]uint64{0}, 0, 1, -1},
		{[]uint64{0}, 0, 63, -1},
		{[]uint64{0}, 0, 64, -1},

		{[]uint64{1}, 0, 0, -1},
		{[]uint64{1}, 0, 1, 0},
		{[]uint64{1}, 0, 2, 0},
		{[]uint64{1}, 1, 2, -1},

		{[]uint64{0x02}, 0, 0, -1},
		{[]uint64{0x02}, 0, 1, -1},
		{[]uint64{0x02}, 0, 2, 1},
		{[]uint64{0x02}, 1, 2, 1},
		{[]uint64{0x02}, 2, 2, -1},
		{[]uint64{0x02}, 2, 63, -1},
		{[]uint64{0x02}, 2, 64, -1},

		{[]uint64{0xff, 0x82}, 0, 72, 0},
		{[]uint64{0xff, 0x82}, 1, 72, 1},
		{[]uint64{0xff, 0x82}, 63, 64, -1},
		{[]uint64{0xff, 0x82}, 64, 65, -1},
		{[]uint64{0xff, 0x82}, 64, 66, 65},
		{[]uint64{0xff, 0x82}, 65, 66, 65},
		{[]uint64{0xff, 0x82}, 65, 72, 65},
		{[]uint64{0xff, 0x82}, 66, 71, -1},
		{[]uint64{0xff, 0x82}, 71, 72, 71},
		{[]uint64{0xff, 0x82}, 72, 128, -1},
	}

	for i, c := range cases {
		got := NextOne(c.bm, c.from, c.n)
		ta.Equal(c.want, got, "%d-th: case: %+v", i+1, c)
	}
}

func TestPrevOne(t *testing.T) {

	ta := require.New(t)

	cases := []struct {
		bm   []uint64
		from int32
		n    int32
		want int32
	}{
		{[]uint64{0}, 0, 1, -1},
		{[]uint64{0}, 0, 63, -1},
		{[]uint64{0}, 0, 64, -1},

		{[]uint64{1}, 0, 1, 0},
		{[]uint64{1}, 0, 2, 0},
		{[]uint64{1}, 1, 2, -1},

		{[]uint64{0x02}, 0, 1, -1},
		{[]uint64{0x02}, 0, 2, 1},
		{[]uint64{0x02}, 1, 1, -1},
		{[]uint64{0x02}, 1, 2, 1},
		{[]uint64{0x02}, 1, 3, 1},
		{[]uint64{0x02}, 2, 2, -1},
		{[]uint64{0x02}, 2, 3, -1},
		{[]uint64{0x02}, 2, 63, -1},
		{[]uint64{0x02}, 2, 64, -1},

		{[]uint64{0xff, 0x82}, 0, 63, 7},
		{[]uint64{0xff, 0x82}, 0, 64, 7},
		{[]uint64{0xff, 0x82}, 0, 65, 7},
		{[]uint64{0xff, 0x82}, 0, 66, 65},
		{[]uint64{0xff, 0x82}, 0, 72, 71},
		{[]uint64{0xff, 0x82}, 1, 72, 71},
		{[]uint64{0xff, 0x82}, 1, 71, 65},
		{[]uint64{0xff, 0x82}, 1, 64, 7},
		{[]uint64{0xff, 0x82}, 63, 64, -1},
		{[]uint64{0xff, 0x82}, 64, 65, -1},
		{[]uint64{0xff, 0x82}, 64, 66, 65},
		{[]uint64{0xff, 0x82}, 65, 66, 65},
		{[]uint64{0xff, 0x82}, 65, 72, 71},
		{[]uint64{0xff, 0x82}, 66, 71, -1},
		{[]uint64{0xff, 0x82}, 71, 72, 71},
		{[]uint64{0xff, 0x82}, 72, 128, -1},

		{[]uint64{0xffffffffffffffff, 0x82}, 63, 64, 63},
		{[]uint64{0xffffffffffffffff, 0x82}, 63, 65, 63},
		{[]uint64{0xffffffffffffffff, 0x82}, 63, 66, 65},
	}

	for i, c := range cases {
		got := PrevOne(c.bm, c.from, c.n)
		ta.Equal(c.want, got, "%d-th: case: %+v", i+1, c)
	}
}

func TestPrevOne_panic(t *testing.T) {

	// zero-width range will panic

	ta := require.New(t)

	cases := []struct {
		bm   []uint64
		from int32
		n    int32
		want int32
	}{
		{[]uint64{0}, 0, 0, -1},
		{[]uint64{1}, 0, 0, -1},
	}

	for _, c := range cases {
		ta.Panics(func(){
			PrevOne(c.bm,c.from,c.n)
		}, "case: %v", c)
	}
}
