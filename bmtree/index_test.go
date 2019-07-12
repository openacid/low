package bmtree

import (
	"testing"

	"github.com/openacid/low/bitmap"
	"github.com/stretchr/testify/require"
)

func TestPathToIndex_fulltree_height_0_to_6(t *testing.T) {

	ta := require.New(t)

	cases := []struct {
		treeheight int32
	}{
		{0},
		{1},
		{2},
		{3},
		{4},
		{5},
		{6},
	}

	for _, c := range cases {
		paths := AllPaths(int32(bitmap.MaskUpto[c.treeheight]), 0, 1<<63)

		bitmapSize := bitmap.MaskUpto[c.treeheight]

		// test if generated index are consecutive ints
		prev := int32(0)
		for i, mp := range paths {
			got := PathToIndex(int32(bitmapSize), mp)
			ta.Equal(prev, got, "%d-th: treeheight: %d mp: %064b", i+1, c.treeheight, mp)

			got, has := PathToIndexLoose(int32(bitmapSize), mp)
			ta.Equal(prev, got, "%d-th: treeheight: %d mp: %064b", i+1, c.treeheight, mp)
			ta.Equal(int32(1), has)

			gotpath := IndexToPath(c.treeheight, got)
			ta.Equal(mp, gotpath, "%d-th: treeheight: %d mp: %064b", i+1, c.treeheight, mp)

			prev++
		}
	}
}

func TestPathToIndex_halftree_height_0_to_6(t *testing.T) {

	ta := require.New(t)

	cases := []struct {
		treeheight int32
	}{
		{0},
		{1},
		{2},
		{3},
		{4},
		{5},
		{6},
	}

	for _, c := range cases {
		bitmapSize := bitmap.Bit[c.treeheight]

		// test if generated index are consecutive ints

		for i := uint64(0); i < bitmapSize; i++ {
			mask := (uint64(1) << uint(c.treeheight)) - 1
			mp := i<<32 | mask
			got := PathToIndex(int32(bitmapSize), mp)
			ta.Equal(int32(i), got, "%d-th: treeheight: %d mp: %064b", i+1, c.treeheight, mp)

			got, has := PathToIndexLoose(int32(bitmapSize), mp)
			ta.Equal(int32(i), got, "%d-th: treeheight: %d mp: %064b", i+1, c.treeheight, mp)
			ta.Equal(int32(1), has)
		}
	}
}

func TestPathToIndex_partialtree(t *testing.T) {

	ta := require.New(t)

	cases := []struct {
		bitmapSize int32
	}{
		{0x01},
		{0x3},
		{0x5},
		{0x72},
		{0xfd49},
	}

	for _, c := range cases {

		paths := AllPaths(c.bitmapSize, 0, 1<<63)
		h := Height(c.bitmapSize)

		// test if generated index are consecutive ints
		prev := int32(0)
		for i, mp := range paths {
			l := PathLen(mp)
			if c.bitmapSize&int32(bitmap.Bit[l]) == 0 {
				continue
			}
			got := PathToIndex(c.bitmapSize, mp)
			ta.Equal(prev, got, "%d-th: bitmapSize: %032b mp: %064b height: %d, pathlen:%d", i+1, c.bitmapSize, mp, h, l)

			got, has := PathToIndexLoose(c.bitmapSize, mp)
			ta.Equal(prev, got, "%d-th: bitmapSize: %032b mp: %064b height: %d, pathlen:%d", i+1, c.bitmapSize, mp, h, l)
			ta.Equal(int32(1), has)

			prev++
		}
	}
}

func TestPathToIndexLoose_absentLevel(t *testing.T) {

	ta := require.New(t)

	cases := []struct {
		bitmapSize int32
		path       uint64
		wantindex  int32
		wanthas    int32
	}{
		{0x5, 0x000000000, 0, 1},
		{0x5, 0x000000002, 1, 0},
		{0x5, 0x000000003, 1, 1},
		{0x5, 0x100000003, 2, 1},
		{0x5, 0x200000002, 3, 0},
		{0x5, 0x200000003, 3, 1},
		{0x5, 0x300000003, 4, 1},
	}

	for i, c := range cases {

		got, has := PathToIndexLoose(c.bitmapSize, c.path)

		ta.Equal(c.wantindex, got, "%d-th: %+v", i+1, c)
		ta.Equal(c.wanthas, has, "%d-th: %+v", i+1, c)

	}
}
