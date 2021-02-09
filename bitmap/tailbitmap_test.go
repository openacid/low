package bitmap

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewTailBitmap(t *testing.T) {

	ta := require.New(t)

	cases := []struct {
		input int64
		want  *TailBitmap
	}{
		{
			input: 0,
			want: &TailBitmap{
				Offset:    0,
				Words:     make([]uint64, 0, 1024),
				reclaimed: 0,
			},
		},
		{
			input: 999999,
			want: &TailBitmap{
				Offset:    999999,
				Words:     make([]uint64, 0, 1024),
				reclaimed: 999999,
			},
		},
	}

	for i, c := range cases {
		got := NewTailBitmap(c.input)
		ta.Equal(c.want, got, "%d-th: case: %+v", i+1, c)
	}
}

func TestTailBitmap_Compact(t *testing.T) {

	ta := require.New(t)

	allOnes1024 := make([]uint64, 1024)
	for i, _ := range allOnes1024 {
		allOnes1024[i] = 0xffffffffffffffff
	}

	cases := []struct {
		input *TailBitmap
		want  *TailBitmap
	}{
		{
			input: &TailBitmap{
				Offset:    0,
				Words:     []uint64{0xffffffffffffffff},
				reclaimed: 0,
			},
			want: &TailBitmap{
				Offset:    64,
				Words:     []uint64{},
				reclaimed: 0,
			},
		},
		{
			input: &TailBitmap{
				Offset:    64,
				Words:     []uint64{0xffffffffffffffff},
				reclaimed: 0,
			},
			want: &TailBitmap{
				Offset:    64 * 2,
				Words:     []uint64{},
				reclaimed: 0,
			},
		},
		{
			input: &TailBitmap{
				Offset:    64,
				Words:     []uint64{0xffffffffffffffff, 1},
				reclaimed: 0,
			},
			want: &TailBitmap{
				Offset:    64 * 2,
				Words:     []uint64{1},
				reclaimed: 0,
			},
		},
		{
			input: &TailBitmap{
				Offset:    64,
				Words:     allOnes1024,
				reclaimed: 0,
			},
			want: &TailBitmap{
				Offset:    64 * 1025,
				Words:     []uint64{},
				reclaimed: 64 * 1025,
			},
		},
	}

	for i, c := range cases {
		c.input.Compact()
		ta.Equal(c.want, c.input, "%d-th: case: %+v", i+1, c)
	}
}

func TestTailBitmap_Set(t *testing.T) {

	ta := require.New(t)

	allOnes1024 := make([]uint64, 1024)
	for i, _ := range allOnes1024 {
		allOnes1024[i] = 0xffffffffffffffff
	}

	cases := []struct {
		input *TailBitmap
		set   int64
		want  *TailBitmap
	}{
		{
			input: &TailBitmap{
				Offset:    0,
				Words:     []uint64{},
				reclaimed: 0,
			},
			set: 0,
			want: &TailBitmap{
				Offset:    0,
				Words:     []uint64{1},
				reclaimed: 0,
			},
		},
		{
			input: &TailBitmap{
				Offset:    64,
				Words:     []uint64{},
				reclaimed: 0,
			},
			set: 65,
			want: &TailBitmap{
				Offset:    64,
				Words:     []uint64{2},
				reclaimed: 0,
			},
		},
		{
			input: &TailBitmap{
				Offset:    64 * 2,
				Words:     []uint64{1},
				reclaimed: 0,
			},
			set: 5,
			want: &TailBitmap{
				Offset:    64 * 2,
				Words:     []uint64{1},
				reclaimed: 0,
			},
		},
		{
			input: &TailBitmap{
				Offset:    64 * 2,
				Words:     []uint64{1},
				reclaimed: 0,
			},
			set: 64*2 + 1,
			want: &TailBitmap{
				Offset:    64 * 2,
				Words:     []uint64{3},
				reclaimed: 0,
			},
		},
		{
			input: &TailBitmap{
				Offset:    64 * 2,
				Words:     []uint64{1},
				reclaimed: 0,
			},
			set: 64*3 + 2,
			want: &TailBitmap{
				Offset:    64 * 2,
				Words:     []uint64{1, 4},
				reclaimed: 0,
			},
		},
		{
			input: &TailBitmap{
				Offset:    64,
				Words:     []uint64{0xffffffffffffff7f, 1},
				reclaimed: 0,
			},
			set: 64 + 7,
			want: &TailBitmap{
				Offset:    64 * 2,
				Words:     []uint64{1},
				reclaimed: 0,
			},
		},
		{
			input: &TailBitmap{
				Offset:    64 * 1023,
				Words:     []uint64{0xffffffffffffff7f, 1},
				reclaimed: 0,
			},
			set: 64*1023 + 7,
			want: &TailBitmap{
				Offset:    64 * 1024,
				Words:     []uint64{1},
				reclaimed: 64 * 1024,
			},
		},
	}

	for i, c := range cases {
		c.input.Set(c.set)
		ta.Equal(c.want, c.input, "%d-th: case: %+v", i+1, c)
	}
}

func TestTailBitmap_Get(t *testing.T) {

	ta := require.New(t)

	allOnes1024 := make([]uint64, 1024)
	for i, _ := range allOnes1024 {
		allOnes1024[i] = 0xffffffffffffffff
	}

	cases := []struct {
		input *TailBitmap
		get   int64
		want  uint64
	}{
		{
			input: &TailBitmap{
				Offset:    64,
				Words:     []uint64{},
				reclaimed: 0,
			},
			get:  0,
			want: 1,
		},
		{
			input: &TailBitmap{
				Offset:    64,
				Words:     []uint64{},
				reclaimed: 0,
			},
			get:  1,
			want: 2,
		},
		{
			input: &TailBitmap{
				Offset:    64,
				Words:     []uint64{},
				reclaimed: 0,
			},
			get:  63,
			want: 1 << 63,
		},

		{
			input: &TailBitmap{
				Offset:    64,
				Words:     []uint64{0xffffffffffffff7f, 1},
				reclaimed: 0,
			},
			get:  64 + 7,
			want: 0,
		},
		{
			input: &TailBitmap{
				Offset:    64,
				Words:     []uint64{0xffffffffffffff7f, 1},
				reclaimed: 0,
			},
			get:  64 + 6,
			want: 1 << 6,
		},
		{
			input: &TailBitmap{
				Offset:    64,
				Words:     []uint64{0xffffffffffffff7f, 1},
				reclaimed: 0,
			},
			get:  64 + 8,
			want: 1 << 8,
		},
		{
			input: &TailBitmap{
				Offset:    64,
				Words:     []uint64{0xffffffffffffff7f, 1},
				reclaimed: 0,
			},
			get:  64*2 + 0,
			want: 1,
		},
	}

	for i, c := range cases {
		got := c.input.Get(c.get)
		ta.Equal(c.want, got, "%d-th: Get case: %+v", i+1, c)

		got1 := c.input.Get1(c.get)
		if c.want != 0 {
			ta.Equal(uint64(1), got1, "%d-th: Get1 case: %+v", i+1, c)
		} else {
			ta.Equal(uint64(0), got1, "%d-th: Get1 case: %+v", i+1, c)
		}
	}
}
