package bitmap

// Extend allocaed additional 0-bits thus accessing (n-1)-th bit does not panic.
//
// Since 0.5.9
func Extend(bm []uint64, n int32) []uint64 {

	nword := (n + 63) >> 6

	if nword <= int32(len(bm)) {
		return bm
	}

	rst := make([]uint64, nword)
	copy(rst, bm)

	return rst
}
