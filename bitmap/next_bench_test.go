package bitmap

import "testing"

var OutputNextOne int

func BenchmarkNextOne(b *testing.B) {
	bm := []uint64{0xff, 0xff, 0x33, 0x66}

	var s int32

	b.Run("FoundInCurrentWord", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s += NextOne(bm, 0, int32(i&0x7))
		}
		OutputNextOne = int(s)
	})

	b.Run("EnumerateAllBits", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s += NextOne(bm, 65, 128)
		}
		OutputNextOne = int(s)
	})
}

func BenchmarkPrevOne(b *testing.B) {
	bm := []uint64{0xff, 0xff, 0x33, 0x66}

	var s int32

	b.Run("FoundInCurrentWord", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s += NextOne(bm, 0, int32(i&0x7))
		}
		OutputNextOne = int(s)
	})

	b.Run("EnumerateAllBits", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s += NextOne(bm, 65, 128)
		}
		OutputNextOne = int(s)
	})
}
