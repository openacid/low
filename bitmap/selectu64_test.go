package bitmap

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSelect64Asm(t *testing.T) {

	ta := require.New(t)

	cases := []struct {
		input uint64
	}{
		{0},
		{1},
		{2},
		{3},
		{4},
		{0xf},
		{0xf0},
		{0xffffffff},
		{0xffffffff00000000},
		{0xfffffffffffffff0},
		{0xffffffffffffffff},
		{0x6668}, // 000101100110011
	}

	for i, c := range cases {

		nth := int32(-1)
		for j := 0; j < 64; j++ {

			if c.input&(1<<uint(j)) != 0 {
				nth++
				got := selectU64WithoutPDEP(c.input, uint64(nth))
				ta.Equal(uint64(j), got, "%d-th: case: %+v, select: %d", i+1, c, nth)

				got = selectU64WithPDEP(c.input, uint64(nth))
				ta.Equal(uint64(j), got, "%d-th: case: %+v, select: %d", i+1, c, nth)

				got = selectU64(c.input, uint64(nth))
				ta.Equal(uint64(j), got, "%d-th: case: %+v, select: %d", i+1, c, nth)
			}
		}
	}
}
