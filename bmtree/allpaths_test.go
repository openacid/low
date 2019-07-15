package bmtree

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAllPaths(t *testing.T) {

	ta := require.New(t)

	cases := []struct {
		bitmapSize int32
		want       []string
	}{
		{1, []string{""}},
		{2, []string{"0", "1"}},
		{3, []string{"", "0", "1"}},
		{4, []string{"00", "01", "10", "11"}},
		{5, []string{"", "00", "01", "10", "11"}},
		{6, []string{"0", "00", "01", "1", "10", "11"}},
		{7, []string{"", "0", "00", "01", "1", "10", "11"}},
		{8, []string{"000", "001", "010", "011", "100", "101", "110", "111"}},
		{9, []string{"", "000", "001", "010", "011", "100", "101", "110", "111"}},
	}

	for i, c := range cases {
		ps := AllPaths(c.bitmapSize, 0, 1<<63)
		got := make([]string, 0)
		for _, p := range ps {
			got = append(got, PathStr(p))
		}
		ta.Equal(c.want, got, "%d-th: case: %+v", i+1, c)
	}
}

func TestAllPaths_range(t *testing.T) {

	ta := require.New(t)

	cases := []struct {
		bitmapSize int32
		from, to   uint64
		want       []string
	}{
		{3, 0x0000000000000001, 0x0000000100000002, []string{"0", "1"}},
		{9, 0x0000000000000000, 0x0000000f00000010, []string{"", "000", "001", "010", "011", "100", "101", "110", "111"}},
		{9, 0x0000000000000000, 0x0000000100000007, []string{"", "000"}},
		{0x7fffffff, 0x3fffffff3fffffff, 0x4000000000000000, []string{"111111111111111111111111111111"}},
	}

	for i, c := range cases {
		ps := AllPaths(c.bitmapSize, c.from, c.to)
		got := make([]string, 0)
		for _, p := range ps {
			got = append(got, PathStr(p))
		}
		ta.Equal(c.want, got, "%d-th: case: %+v", i+1, c)
	}
}

func TestAllPaths_order(t *testing.T) {

	ta := require.New(t)

	for bitmapSize := int32(1); bitmapSize < 1024; bitmapSize++ {
		ps := AllPaths(bitmapSize, 0, 1<<63)
		prev := uint64(0)
		for i, p := range ps {
			if i > 0 {
				ta.True(p > prev)
			}
			prev = p
		}
	}
}
