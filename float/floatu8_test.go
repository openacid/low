package float

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestU8_upto15(t *testing.T) {

	ta := require.New(t)

	for i := uint64(0); i < 1<<4-1; i++ {
		got := U8(i)
		ta.Equal(i, uint64(got), "%d-th: case: %+v", i+1, i)

		g2 := U8decode(got)
		ta.Equal(i, g2, "%d-th: case: %+v", i+1, i)
	}

}

func TestU8(t *testing.T) {

	ta := require.New(t)

	cases := []struct {
		n        uint64
		wantu8   uint8
		wantnorm uint64
	}{
		{0, 0, 0},
		{1, 1, 1},
		{15, 15, 15},
		{16, 0x18, 16},
		{17, 0x18, 16},
		{18, 0x19, 18},
		{19, 0x19, 18},

		{15 << 15, 0xff, 15 << 15},

		// max
		{1<<19 - 1, 0xff, 15 << 15},
	}

	for i, c := range cases {
		got := U8(c.n)
		ta.Equal(c.wantu8, got, "%d-th: case: %+v", i+1, c)

		g2 := U8decode(got)
		ta.Equal(c.wantnorm, g2, "%d-th: case: %+v", i+1, c)

		gotnorm := U8normalize(c.n)
		ta.Equal(c.wantnorm, gotnorm, "%d-th: exp: %+v", i+1, c)
	}
}
