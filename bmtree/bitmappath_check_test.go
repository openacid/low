// +build debug

package bmtree

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBitmapPathMustHaveEqualHeight(t *testing.T) {

	ta := require.New(t)

	cases := []struct {
		bitmapSize int32
		path       uint64
		wantpanic  bool
	}{

		// path is too long
		{5, 4, true},

		{7, 4, true},
		{7, 5, true},
		{7, 6, true},
		{7, 7, true},
		{7, 8, true},

		// path is too short
		{7, 1, true},

		// path must be consecutive "1"s
		{0xf, 5, true},

		// path mask shorter than path bits
		{0xf, 0xf<<32 + 0xe, true},

		// bitmap height <=30
		{-1, 0, true},

		// path height <=30
		{1, 1 << 31, true},
		{1, 1 << 30, true},

		// path points to present level
		{7, 0, false},
		{7, 2, false},
		{7, 3, false},

		{5, 0, false},
		{5, 3, false},
	}

	for _, c := range cases {
		if c.wantpanic {
			ta.Panics(func() { bitmapPathMustHaveEqualHeight(c.bitmapSize, c.path) }, "%032b, %064b", c.bitmapSize, c.path)
		} else {
			ta.NotPanics(func() { bitmapPathMustHaveEqualHeight(c.bitmapSize, c.path) }, "%032b, %064b", c.bitmapSize, c.path)

		}
	}

}
