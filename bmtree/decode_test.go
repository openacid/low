package bmtree

import (
	"testing"

	"github.com/openacid/low/bitmap"
	"github.com/stretchr/testify/require"
)

func TestDecode(t *testing.T) {

	ta := require.New(t)

	cases := []struct {
		bitmapSize  int32
		bitmapIndex []int32
		want        []string
	}{
		{0x1, []int32{}, []string{}},
		{0x3, []int32{}, []string{}},
		{0x3, []int32{0}, []string{""}},
		{0x3, []int32{1}, []string{"0"}},
		{0x3, []int32{1, 2}, []string{"0", "1"}},
		{0x7, []int32{1, 2, 3, 6}, []string{"0", "00", "01", "11"}},
		{0xd, []int32{0, 1, 3, 4, 9, 12}, []string{"", "00", "001", "01", "101", "111"}},

		{0xf08, []int32{0, 63, 64, 65}, []string{"000", "0000010000", "00000100000", "00000100001"}},
	}

	for i, c := range cases {
		bm := bitmap.Of(c.bitmapIndex)
		ps := Decode(c.bitmapSize, bm)
		got := []string{}
		for _, p := range ps {
			got = append(got, PathStr(p))
		}

		ta.Equal(c.want, got, "%d-th: case: %+v", i+1, c)
	}
}

func TestDecode_fulltree_height_7(t *testing.T) {

	ta := require.New(t)

	bmidx := []int32{
		0x01,
		0x0f,
		0x1a,
		0x30,
		0x5c,
		0x65,
		0x89,
		0xc6,
	}

	paths := []uint64{
		0x0000000000000040,
		0x000000050000007f,
		0x0000000b0000007f,
		0x000000160000007f,
		0x0000002c0000007f,
		0x000000300000007f,
		0x000000420000007f,
		0x000000610000007f,
	}

	got := Decode(1<<(7+1)-1, bitmap.Of(bmidx))
	ta.Equal(paths, got)
}
