// +build debug

package bmtree

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPathToIndex_input(t *testing.T) {

	ta := require.New(t)

	cases := []struct {
		bitmapSize int32
		path       uint64
		wantpanic  bool
	}{
		// path to a root but tree does not store.
		{4, 1, true},
		{4, 2, true},

		// path to absent level
		{5, 1, true},
		{5, 2, true},

		// longer
		{5, 4, true},

		// path is longer than expected
		{7, 8, true},
		{7, 4, true},
		{7, 5, true},
		{7, 6, true},
		{7, 7, true},

		// path must be consecutive "1"s
		{0xf, 5, true},

		// path mask shorter than path bits
		{0xf, 0xf<<32 + 0xe, true},

		// bitmap length <=31
		{-1, 0, true},

		// path length <=31
		{1, 1 << 31, true},

		// path points to present level
		{7, 0, false},
		{7, 2, false},
		{7, 3, false},

		{5, 0, false},
		{5, 3, false},
	}

	for _, c := range cases {
		if c.wantpanic {
			ta.Panics(func() { PathToIndex(c.bitmapSize, c.path) }, "%032b, %064b", c.bitmapSize, c.path)
		} else {
			ta.NotPanics(func() { PathToIndex(c.bitmapSize, c.path) }, "%032b, %064b", c.bitmapSize, c.path)

		}
	}

}

func TestPathToIndexLoose_input(t *testing.T) {

	ta := require.New(t)

	cases := []struct {
		bitmapSize int32
		path       uint64
		wantpanic  bool
	}{
		// path to a root but tree does not store.
		{4, 1, true},
		{4, 2, true},

		// path to absent level
		{5, 1, true},
		{5, 2, true},

		// longer
		{5, 4, true},

		// path is longer than expected
		{7, 8, true},
		{7, 4, true},
		{7, 5, true},
		{7, 6, true},
		{7, 7, true},

		// path must be consecutive "1"s
		{0xf, 5, true},

		// path mask shorter than path bits
		{0xf, 0xf<<32 + 0xe, true},

		// bitmap length <=31
		{-1, 0, true},

		// path length <=31
		{1, 1 << 31, true},

		// path points to present level
		{7, 0, false},
		{7, 2, false},
		{7, 3, false},

		{5, 0, false},
		{5, 3, false},
	}

	for _, c := range cases {
		if c.wantpanic {
			ta.Panics(func() { PathToIndex(c.bitmapSize, c.path) }, "%032b, %064b", c.bitmapSize, c.path)
		} else {
			ta.NotPanics(func() { PathToIndex(c.bitmapSize, c.path) }, "%032b, %064b", c.bitmapSize, c.path)

		}
	}

}
