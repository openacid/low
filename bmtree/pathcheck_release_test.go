// +build !debug

package bmtree

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPathCheck_release(t *testing.T) {

	ta := require.New(t)

	cases := []struct {
		path uint64
	}{
		// path must be consecutive "1"s
		{5},

		// path mask shorter than path bits
		{0xf<<32 + 0xe},

		// path length must <=31
		{1 << 31},
		{1 << 30},

		{0x7},
	}

	for _, c := range cases {
		ta.NotPanics(func() { pathCheck(c.path) }, "%064b", c.path)
	}
}
