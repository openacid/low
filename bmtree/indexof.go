package bmtree

// IndexOf returns the bit index in the bitmap of a key, which is a string
// from the specified bit.
//
// Since 0.1.9
func IndexOf(bitmapSize int32, s string, frombit int32) int32 {
	p := PathOf(s, frombit, Height(bitmapSize))
	return PathToIndex(bitmapSize, p)
}

// IndexesOf returns the bit index in the bitmap of a array of keys,
// from the specified bit in a string.
//
// Since 0.1.9
func IndexesOf(bitmapSize int32, keys []string, frombit int32, dedup bool) []int32 {
	paths := PathsOf(keys, frombit, Height(bitmapSize), dedup)
	indexes := make([]int32, 0)

	for _, p := range paths {
		indexes = append(indexes, PathToIndex(bitmapSize, p))
	}
	return indexes
}
