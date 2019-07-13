package bmtree

// Size returns the number of bits required for the bitmap to store all of
// bit array with length from 0 to n.
//
// Since 0.1.9
func Size(n int32) int32 {
	return 1 << uint(n+1)
}
