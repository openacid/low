package bmtree

// IndexOf returns the bit position in the bitmap of path from the
// specified bit in a string.
//
// Since 0.1.9
func (b *Builder) IndexOf(s string, frombit int32) int32 {
	p := PathOf(s, frombit, b.treeHeight)
	return PathToIndex(b.bitmapSize, p)
}

// IndexesOf returns the bit index in the bitmap of a array of paths made
// from the specified bit in a string.
//
// Since 0.1.9
func (b *Builder) IndexesOf(keys []string, frombit int32, dedup bool) []int32 {
	paths := PathsOf(keys, frombit, b.treeHeight, dedup)
	ps := make([]int32, 0)

	for _, p := range paths {
		ps = append(ps, PathToIndex(b.bitmapSize, p))
	}
	return ps
}
