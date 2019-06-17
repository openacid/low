package bmtree

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBuilder_Size(t *testing.T) {

	ta := require.New(t)

	cases := []struct {
		bitmapSize int32
		want       int32
	}{
		{2<<0 - 1, 2<<0 - 1},
		{2<<1 - 1, 2<<1 - 1},
		{2<<2 - 1, 2<<2 - 1},
		{2<<3 - 1, 2<<3 - 1},
		{2<<29 - 1, 2<<29 - 1},
		{0x7f0f, 0x7f0f},
	}

	for i, c := range cases {
		got := NewBuilder(c.bitmapSize).Size()
		ta.Equal(c.want, got, "%d-th: case: %+v", i+1, c)
	}
}
