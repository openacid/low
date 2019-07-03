package bitmap

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGet(t *testing.T) {

	ta := require.New(t)

	ta.Panics(func() { Get(nil, 0) })
	ta.Panics(func() { Get([]uint64{}, 0) })
	ta.Panics(func() { Get([]uint64{}, 1) })
	ta.Panics(func() { Get([]uint64{}, -1) })
	ta.Panics(func() { Get([]uint64{0}, -1) })
	ta.Panics(func() { Get([]uint64{0}, 64) })

	cases := []struct {
		input []uint64
		n     int32
		want  uint64
	}{
		{[]uint64{0}, 0, 0},
		{[]uint64{0}, 1, 0},
		{[]uint64{1}, 0, 1},
		{[]uint64{1}, 1, 0},
		{[]uint64{12}, 1, 0},
		{[]uint64{12}, 2, 4},
		{[]uint64{12}, 3, 8},
		{[]uint64{32, 12}, 5, 32},
		{[]uint64{32, 12}, 64, 0},
		{[]uint64{32, 12}, 66, 4},
		{[]uint64{32, 12}, 67, 8},
	}

	for i, c := range cases {
		got := Get(c.input, c.n)
		ta.Equal(c.want, got, "%d-th: case: %+v", i+1, c)

		got = Get1(c.input, c.n)
		want := uint64(0)
		if c.want != 0 {
			want = 1
		}
		ta.Equal(want, got, "%d-th: case: %+v", i+1, c)
	}
}

func TestGetw(t *testing.T) {

	ta := require.New(t)

	ta.Panics(func() { Getw(nil, 0, 1) })
	ta.Panics(func() { Getw([]uint64{}, 0, 1) })
	ta.Panics(func() { Getw([]uint64{}, 1, 1) })
	ta.Panics(func() { Getw([]uint64{}, -1, 1) })
	ta.Panics(func() { Getw([]uint64{0}, -1, 1) })
	ta.Panics(func() { Getw([]uint64{0}, 64, 1) })

	cases := []struct {
		bm   []uint64
		n    int32
		w    int32
		want uint64
	}{
		{[]uint64{0}, 0, 1, 0},
		{[]uint64{0}, 1, 1, 0},
		{[]uint64{1}, 0, 1, 1},
		{[]uint64{1}, 1, 1, 0},
		{[]uint64{12}, 1, 1, 0},
		{[]uint64{12}, 2, 1, 1},
		{[]uint64{12}, 3, 1, 1},
		{[]uint64{32, 12}, 5, 1, 1},
		{[]uint64{32, 12}, 64, 1, 0},
		{[]uint64{32, 12}, 66, 1, 1},
		{[]uint64{32, 12}, 67, 1, 1},

		{[]uint64{0}, 0, 2, 0},
		{[]uint64{0}, 1, 2, 0},
		{[]uint64{1}, 0, 2, 1},
		{[]uint64{1}, 1, 2, 0},
		{[]uint64{12}, 0, 2, 0},
		{[]uint64{12}, 1, 2, 3},
		{[]uint64{12}, 2, 2, 0},
		{[]uint64{32, 12}, 1, 2, 0},
		{[]uint64{32, 12}, 2, 2, 2},
		{[]uint64{32, 12}, 3, 2, 0},
		{[]uint64{32, 12}, 32, 2, 0},
		{[]uint64{32, 12}, 33, 2, 3},
		{[]uint64{32, 12}, 34, 2, 0},
	}

	for i, c := range cases {
		got := Getw(c.bm, c.n, c.w)
		ta.Equal(c.want, got, "%d-th: case: %+v", i+1, c)
	}
}

func TestSafeGet(t *testing.T) {

	ta := require.New(t)

	cases := []struct {
		input []uint64
		n     int32
		want  uint64
	}{
		{nil, 0, 0},
		{[]uint64{}, 0, 0},
		{[]uint64{}, 1, 0},
		{[]uint64{}, -1, 0},
		{[]uint64{0}, -1, 0},
		{[]uint64{0}, 64, 0},

		{[]uint64{0}, 0, 0},
		{[]uint64{0}, 1, 0},
		{[]uint64{1}, 0, 1},
		{[]uint64{1}, 1, 0},
		{[]uint64{12}, 1, 0},
		{[]uint64{12}, 2, 4},
		{[]uint64{12}, 3, 8},
		{[]uint64{32, 12}, 5, 32},
		{[]uint64{32, 12}, 64, 0},
		{[]uint64{32, 12}, 66, 4},
		{[]uint64{32, 12}, 67, 8},
	}

	for i, c := range cases {
		got := SafeGet(c.input, c.n)
		ta.Equal(c.want, got, "%d-th: case: %+v", i+1, c)

		got = SafeGet1(c.input, c.n)
		want := uint64(0)
		if c.want != 0 {
			want = 1
		}
		ta.Equal(want, got, "%d-th: case: %+v", i+1, c)
	}
}
