package bmtree

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHeight(t *testing.T) {

	ta := require.New(t)

	cases := []struct {
		bitmapSize int32
		want       int32
	}{
		{1, 0},
		{2, 1},
		{3, 1},
		{4, 2},
		{5, 2},
		{6, 2},
		{7, 2},
		{8, 3},
		{8 + 1, 3},
		{16, 4},
		{16 + 1, 4},
		{1 << 30, 30},
		{1<<30 + 1, 30},
		{1<<30 + 2, 30},
		{0x7fffffff, 30},
	}

	for i, c := range cases {
		got := Height(c.bitmapSize)
		ta.Equal(c.want, got, "%d-th: case: %+v", i+1, c)
	}
}
