package bitmap

import "testing"

var DsOutput int

func BenchmarkRank64_5_bits(b *testing.B) {

	bm := Of([]int32{1, 2, 3, 64, 129})
	idx := IndexRank64(bm)

	b.ResetTimer()

	var gotrank int32
	for i := 0; i < b.N; i++ {
		upto, _ := Rank64(bm, idx, int32(i&127))
		gotrank += upto
	}

	DsOutput = int(gotrank)
}

func BenchmarkRank128_5_bits(b *testing.B) {

	bm := Of([]int32{1, 2, 3, 64, 129})
	idx := IndexRank128(bm)

	b.ResetTimer()

	var gotrank int32
	for i := 0; i < b.N; i++ {
		upto, _ := Rank128(bm, idx, int32(i&127))
		gotrank += upto
	}

	DsOutput = int(gotrank)
}

func BenchmarkRank64_64k_bits(b *testing.B) {

	n := 64 * 1024
	mask := n - 1
	indexes := make([]int32, n)
	for i := 0; i < n; i++ {
		indexes[i] = int32(i * 2)
	}
	bm := Of(indexes)
	idx := IndexRank64(bm)

	b.ResetTimer()

	var gotrank int32
	for i := 0; i < b.N; i++ {
		upto, _ := Rank64(bm, idx, int32(i&mask))
		gotrank += upto
	}

	DsOutput = int(gotrank)
}

func BenchmarkRank128_64k_bits(b *testing.B) {

	n := 64 * 1024
	mask := n - 1
	indexes := make([]int32, n)
	for i := 0; i < n; i++ {
		indexes[i] = int32(i * 2)
	}
	bm := Of(indexes)
	idx := IndexRank128(bm)

	b.ResetTimer()

	var gotrank int32
	for i := 0; i < b.N; i++ {
		upto, _ := Rank128(bm, idx, int32(i&mask))
		gotrank += upto
	}

	DsOutput = int(gotrank)
}
