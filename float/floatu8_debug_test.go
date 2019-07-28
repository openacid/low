// +build debug

package float

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestU8_panic(t *testing.T) {

	ta := require.New(t)

	cases := []struct {
		n uint64
	}{

		// max
		{1 << 19},
		{1<<19 + 1},
	}

	for _, c := range cases {
		ta.Panics(func() { U8(c.n) })
		ta.Panics(func() { U8normalize(c.n) })
	}
}
