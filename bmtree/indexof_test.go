package bmtree

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIndexOf_fulltree(t *testing.T) {

	ta := require.New(t)

	k := "abcdefgh"

	cases := []struct {
		bitmapSize, frombit int32
		idx                 string
	}{

		{bitmapSize: 2<<1 - 1, frombit: 0, idx: "00000000000000000000000000000001"},
		{bitmapSize: 2<<2 - 1, frombit: 0, idx: "00000000000000000000000000000011"},
		{bitmapSize: 2<<2 - 1, frombit: 1, idx: "00000000000000000000000000000110"},
		{bitmapSize: 2<<4 - 1, frombit: 0, idx: "00000000000000000000000000001110"},
		{bitmapSize: 2<<4 - 1, frombit: 1, idx: "00000000000000000000000000011010"},
		{bitmapSize: 2<<4 - 1, frombit: 2, idx: "00000000000000000000000000010011"},
		{bitmapSize: 2<<4 - 1, frombit: 3, idx: "00000000000000000000000000000100"},
		{bitmapSize: 2<<7 - 1, frombit: 0, idx: "00000000000000000000000001100101"},
		{bitmapSize: 2<<7 - 1, frombit: 1, idx: "00000000000000000000000011000110"},
		{bitmapSize: 2<<7 - 1, frombit: 2, idx: "00000000000000000000000010001001"},
		{bitmapSize: 2<<7 - 1, frombit: 3, idx: "00000000000000000000000000001111"},
		{bitmapSize: 2<<7 - 1, frombit: 4, idx: "00000000000000000000000000011010"},
		{bitmapSize: 2<<7 - 1, frombit: 5, idx: "00000000000000000000000000110000"},
		{bitmapSize: 2<<7 - 1, frombit: 6, idx: "00000000000000000000000001011100"},
		{bitmapSize: 2<<12 - 1, frombit: 0, idx: "00000000000000000000110000110011"},
		{bitmapSize: 2<<12 - 1, frombit: 1, idx: "00000000000000000001100001011111"},
		{bitmapSize: 2<<12 - 1, frombit: 2, idx: "00000000000000000001000010111000"},
		{bitmapSize: 2<<12 - 1, frombit: 3, idx: "00000000000000000000000101101010"},
		{bitmapSize: 2<<12 - 1, frombit: 4, idx: "00000000000000000000001011001100"},
		{bitmapSize: 2<<12 - 1, frombit: 5, idx: "00000000000000000000010110010000"},
		{bitmapSize: 2<<12 - 1, frombit: 6, idx: "00000000000000000000101100011001"},
		{bitmapSize: 2<<12 - 1, frombit: 7, idx: "00000000000000000001011000101100"},
		{bitmapSize: 2<<12 - 1, frombit: 8, idx: "00000000000000000000110001010011"},
		{bitmapSize: 2<<12 - 1, frombit: 9, idx: "00000000000000000001100010011111"},
		{bitmapSize: 2<<12 - 1, frombit: 10, idx: "00000000000000000001000100111000"},
		{bitmapSize: 2<<12 - 1, frombit: 11, idx: "00000000000000000000001001101010"},
		{bitmapSize: 2<<21 - 1, frombit: 0, idx: "00000000000110000101100010100101"},
		{bitmapSize: 2<<21 - 1, frombit: 1, idx: "00000000001100001011000100111101"},
		{bitmapSize: 2<<21 - 1, frombit: 2, idx: "00000000001000010110001001101111"},
		{bitmapSize: 2<<21 - 1, frombit: 3, idx: "00000000000000101100010011010011"},
		{bitmapSize: 2<<21 - 1, frombit: 4, idx: "00000000000001011000100110011001"},
		{bitmapSize: 2<<21 - 1, frombit: 5, idx: "00000000000010110001001100100110"},
		{bitmapSize: 2<<21 - 1, frombit: 6, idx: "00000000000101100010011001000001"},
		{bitmapSize: 2<<21 - 1, frombit: 7, idx: "00000000001011000100110001110111"},
		{bitmapSize: 2<<21 - 1, frombit: 8, idx: "00000000000110001001100011100100"},
		{bitmapSize: 2<<21 - 1, frombit: 9, idx: "00000000001100010011000110111101"},
		{bitmapSize: 2<<21 - 1, frombit: 10, idx: "00000000001000100110001101110000"},
		{bitmapSize: 2<<21 - 1, frombit: 11, idx: "00000000000001001100011011010101"},
		{bitmapSize: 2<<21 - 1, frombit: 12, idx: "00000000000010011000110110011101"},
		{bitmapSize: 2<<21 - 1, frombit: 13, idx: "00000000000100110001101100101110"},
		{bitmapSize: 2<<21 - 1, frombit: 14, idx: "00000000001001100011011001010001"},
		{bitmapSize: 2<<21 - 1, frombit: 15, idx: "00000000000011000110110010011000"},
		{bitmapSize: 2<<21 - 1, frombit: 16, idx: "00000000000110001101100100100100"},
		{bitmapSize: 2<<21 - 1, frombit: 17, idx: "00000000001100011011001000111101"},
		{bitmapSize: 2<<21 - 1, frombit: 18, idx: "00000000001000110110010001110000"},
		{bitmapSize: 2<<21 - 1, frombit: 19, idx: "00000000000001101100100011010110"},
		{bitmapSize: 2<<21 - 1, frombit: 20, idx: "00000000000011011001000110100000"},
	}

	for i, c := range cases {
		got := IndexOf(c.bitmapSize, k, c.frombit)
		gots := fmt.Sprintf("%032b", got)
		ta.Equal(c.idx, gots, "%d-th: case: %+v", i+1, c)
	}
}

func TestIndexOf_partialtree(t *testing.T) {

	ta := require.New(t)

	k := "abcdefgh"

	cases := []struct {
		bitmapSize, frombit int32
		idx                 string
	}{

		{bitmapSize: 0x0003, frombit: 0, idx: "00000000000000000000000000000001"},
		{bitmapSize: 0x0002, frombit: 0, idx: "00000000000000000000000000000000"},

		// height: 12

		{bitmapSize: 0x1fff, frombit: 0, idx: "00000000000000000000110000110011"},
		{bitmapSize: 0x1fff, frombit: 1, idx: "00000000000000000001100001011111"},
		{bitmapSize: 0x1fff, frombit: 2, idx: "00000000000000000001000010111000"},
		{bitmapSize: 0x1fff, frombit: 3, idx: "00000000000000000000000101101010"},
		{bitmapSize: 0x1fff, frombit: 4, idx: "00000000000000000000001011001100"},
		{bitmapSize: 0x1fff, frombit: 5, idx: "00000000000000000000010110010000"},
		{bitmapSize: 0x1fff, frombit: 6, idx: "00000000000000000000101100011001"},
		{bitmapSize: 0x1fff, frombit: 7, idx: "00000000000000000001011000101100"},
		{bitmapSize: 0x1fff, frombit: 8, idx: "00000000000000000000110001010011"},
		{bitmapSize: 0x1fff, frombit: 9, idx: "00000000000000000001100010011111"},
		{bitmapSize: 0x1fff, frombit: 10, idx: "00000000000000000001000100111000"},
		{bitmapSize: 0x1fff, frombit: 11, idx: "00000000000000000000001001101010"},

		// no root, all index is one smaller

		{bitmapSize: 0x1ffe, frombit: 0, idx: "00000000000000000000110000110010"},
		{bitmapSize: 0x1ffe, frombit: 1, idx: "00000000000000000001100001011110"},
		{bitmapSize: 0x1ffe, frombit: 2, idx: "00000000000000000001000010110111"},
		{bitmapSize: 0x1ffe, frombit: 3, idx: "00000000000000000000000101101001"},
		{bitmapSize: 0x1ffe, frombit: 4, idx: "00000000000000000000001011001011"},
		{bitmapSize: 0x1ffe, frombit: 5, idx: "00000000000000000000010110001111"},
		{bitmapSize: 0x1ffe, frombit: 6, idx: "00000000000000000000101100011000"},
		{bitmapSize: 0x1ffe, frombit: 7, idx: "00000000000000000001011000101011"},
		{bitmapSize: 0x1ffe, frombit: 8, idx: "00000000000000000000110001010010"},
		{bitmapSize: 0x1ffe, frombit: 9, idx: "00000000000000000001100010011110"},
		{bitmapSize: 0x1ffe, frombit: 10, idx: "00000000000000000001000100110111"},
		{bitmapSize: 0x1ffe, frombit: 11, idx: "00000000000000000000001001101001"},
	}

	for i, c := range cases {
		got := IndexOf(c.bitmapSize, k, c.frombit)
		gots := fmt.Sprintf("%032b", got)
		ta.Equal(c.idx, gots, "%d-th: case: %+v", i+1, c)
	}
}

func TestIndexesOf(t *testing.T) {

	ta := require.New(t)

	k := "abcdefgh"
	bitmapSize := int32(1<<22 - 1)

	got := IndexesOf(bitmapSize, []string{k, k}, 3, false)
	ta.Equal(2, len(got))
	ta.Equal("0000000000000000000000000000000000000000000000101100010011010011", fmt.Sprintf("%064b", got[0]))
	ta.Equal("0000000000000000000000000000000000000000000000101100010011010011", fmt.Sprintf("%064b", got[1]))

	got = IndexesOf(bitmapSize, []string{k, k}, 3, true)
	ta.Equal(1, len(got))
	ta.Equal("0000000000000000000000000000000000000000000000101100010011010011", fmt.Sprintf("%064b", got[0]))
}
