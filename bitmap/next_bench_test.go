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

	bm = []uint64{0, 0, 0, 0x66}

	b.Run("EnumerateAllBits", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s += NextOne(bm, 1, 64*4)
		}
		OutputNextOne = int(s)
	})
}

func BenchmarkPrevOne(b *testing.B) {
	bm := []uint64{0xff, 0xff, 0x33, 0x66}

	var s int32

	b.Run("FoundInCurrentWord", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s += PrevOne(bm, 0, int32(i&0x7)+1)
		}
		OutputNextOne = int(s)
	})

	bm = []uint64{3, 0, 0, 0}

	b.Run("EnumerateAllBits", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s += PrevOne(bm, 1, 64*4)
		}
		OutputNextOne = int(s)
	})
}
