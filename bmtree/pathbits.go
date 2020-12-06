package bmtree

// PathBits returns the searching bit, e.g., discard the height info.
//
// Since 0.1.12
func PathBits(path uint64) uint64 {
	return path >> 32
}

// PathMask returns the mask, e.g., discard the searching bits.
//
// Since 0.1.12
func PathMask(path uint64) uint64 {
	return path & 0xffffffff
}
