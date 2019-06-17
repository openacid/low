package bmtree

// Size returns the bitmap size.
//
// Since 0.1.9
func (b *Builder) Size() int32 {
	return b.bitmapSize
}
