package bmtree

import (
	"fmt"
	"testing"
)

var Input int32 = 23
var Output int

var IndexOfInput uint64

func BenchmarkIndexOf(b *testing.B) {

	var s int32 = 0
	IndexOfInput = 0x0000aa000000ff00
	bitmapSize := int32(0xffff)

	for i := 0; i < b.N; i++ {
		s += PathToIndex(bitmapSize, IndexOfInput)
	}
	Output = int(s)
}

func BenchmarkPathOf(b *testing.B) {

	var s uint64 = 0

	for treeHeight := 1; treeHeight < 16; treeHeight++ {
		b.Run(fmt.Sprintf("treeHeight=%d", treeHeight), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				p := IndexToPath(4, Input)
				s += p
			}
		})

	}
	Output = int(s)
}
