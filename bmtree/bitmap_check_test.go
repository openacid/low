// +build debug

package bmtree

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBitmapSizeCheck(t *testing.T) {

	ta := require.New(t)

	cases := []struct {
		bitmapSize int32
		wantpanic  bool
	}{

		// bitmap height <=30
		{-1, true},

		// bitmap must not be 0
		{0, true},

		// valid bitmapSize
		{0x01, false},
		{0x02, false},
		{0x03, false},
		{0x04, false},
		{0x05, false},
		{0x06, false},
		{0x07, false},
		{0x7fffffff, false},
	}

	for _, c := range cases {
		if c.wantpanic {
			ta.Panics(func() { bitmapSizeCheck(c.bitmapSize) }, "%032b", c.bitmapSize)
		} else {
			ta.NotPanics(func() { bitmapSizeCheck(c.bitmapSize) }, "%032b", c.bitmapSize)

		}
	}

}

func TestBitmapMustHaveLevel(t *testing.T) {

	ta := require.New(t)

	cases := []struct {
		bitmapSize int32
		level      int32
		wantpanic  bool
	}{

		{0x0d, 0, false},
		{0x0d, 1, true},
		{0x0d, 2, false},
		{0x0d, 3, false},
		{0x0d, 4, true},
	}

	for _, c := range cases {
		if c.wantpanic {
			ta.Panics(func() { bitmapMustHaveLevel(c.bitmapSize, c.level) }, "%032b, %d", c.bitmapSize, c.level)
		} else {
			ta.NotPanics(func() { bitmapMustHaveLevel(c.bitmapSize, c.level) }, "%032b %d", c.bitmapSize, c.level)

		}
	}

}
