package bitmap

// OfMany creates a bitmap from a list of sub-bitmap bit positions.
// "sizes" specifies the total bits in every sub-bitmap.
//
// Since 0.1.9
func OfMany(subs [][]int32, sizes []int32) []uint64 {

	totalBits := 0
	for _, sb := range subs {
		totalBits += len(sb)
	}

	r := make([]int32, totalBits)
	base := int32(0)
	ith := 0
	for i, e := range subs {
		for _, idx := range e {
			r[ith] = base + idx
			ith++
		}
		base += sizes[i]
	}
	return Of(r, base)
}
