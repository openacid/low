// +build debug

package bmtree

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPathCheck(t *testing.T) {

	ta := require.New(t)

	cases := []struct {
		path      uint64
		wantpanic bool
	}{
		// path must be consecutive "1"s
		{5, true},

		// path mask shorter than path bits
		{0xf<<32 + 0xe, true},

		// path length must <=31
		{1 << 31, true},
		{1 << 30, true},

		{0x7, false},
	}

	for _, c := range cases {
		if c.wantpanic {
			ta.Panics(func() { pathCheck(c.path) }, "%064b", c.path)
		} else {
			ta.NotPanics(func() { pathCheck(c.path) }, "%064b", c.path)

		}
	}
}
