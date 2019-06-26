package bmtree

import "fmt"

// IndexStr returns a binary form string representation of a bit position.
// E.g.
//
// Since 0.1.9
func (b *builder) IndexStr(i int32) string {
	p := IndexToPath(b.treeHeight, i)
	l := PathLen(p)

	if l == 0 {
		return ""
	}

	bs := p >> uint(32+b.treeHeight-l)
	return fmt.Sprintf("%0[1]*[2]b", l, bs)
}
