package bitmap

// Builder provides continous bitmap building.
//
// Since 0.1.19
type Builder struct {
	Words  []uint64
	Offset int32
}

// NewBuilder creates a new Builder with preallocated n bits.
//
// Since 0.1.19
func NewBuilder(n int32) *Builder {
	b := &Builder{
		Words:  make([]uint64, 0, n>>6),
		Offset: 0,
	}

	return b
}

// Extend add bits into builder, the bit position are relative to current
// Builder.Offset and the size in bit of the bitmap is size.
//
// Since 0.1.19
func (b *Builder) Extend(bitPositions []int32, size int32) {
	end := b.Offset + size
	if len(bitPositions) > 0 {
		bitEnd := bitPositions[len(bitPositions)-1]
		if bitEnd >= size {
			end = b.Offset + bitEnd + 1
		}
	}
	for int(end) > len(b.Words)<<6 {
		b.Words = append(b.Words, 0)
	}

	for _, i := range bitPositions {
		idx := b.Offset + i
		b.Words[idx>>6] |= 1 << uint(idx&63)
	}

	b.Offset += size
}
