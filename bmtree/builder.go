package bmtree

// Builder provides functionalities of creating a bitmap from binary tree
// searching paths,
// and extracting paths from a bitmap.
//
// Since 0.1.9
type Builder struct {
	bitmapSize int32
	treeHeight int32
}

// NewBuilder creates a *Builder.
//
// Since 0.1.9
func NewBuilder(bitmapSize int32) *Builder {
	b := &Builder{
		bitmapSize: bitmapSize,
		treeHeight: Height(bitmapSize),
	}
	return b
}
