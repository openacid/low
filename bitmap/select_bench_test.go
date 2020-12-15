package bitmap

import (
	"testing"
)

var InputI32 int32 = 35
var OutputI32 int32 = 0

func BenchmarkSelect(b *testing.B) {
	words := []uint64{0xffffffff, 0xffffffff, 1}
	idx := IndexSelect32(words)
	var s int32
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s += select32single(words, idx, InputI32)
		s += select32single(words, idx, InputI32+1)
	}
	OutputI32 = s
}

func BenchmarkSelect32(b *testing.B) {
	words := []uint64{0xffffffff, 0xffffffff, 1}
	idx := IndexSelect32(words)
	var s int32
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v1, _ := Select32(words, idx, InputI32)
		s += v1
	}
	OutputI32 = s
}

func BenchmarkSelect32R64(b *testing.B) {
	words := []uint64{0xffffffff, 0xffffffff, 1}
	sidx, ridx := IndexSelect32R64(words)
	var s int32
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v1, _ := Select32R64(words, sidx, ridx, InputI32)
		s += v1
	}
	OutputI32 = s
}

func BenchmarkSelectU64Indexed(b *testing.B) {

	word := uint64(0xffffffffffffffff)

	idx := indexSelectU64(word)

	var s int32

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		v1, _ := selectU64Indexed(word, idx, uint64(InputI32))
		s += v1
	}
	OutputI32 = s
}
