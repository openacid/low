package bmtree

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBuilder_IndexStr(t *testing.T) {

	ta := require.New(t)

	cases := []struct {
		treeHeight int32
		pos        int32
		want       string
	}{
		{4, 0, ""},
		{4, 4, "0000"},
		{3, 5, "01"},
		{3, 6, "010"},
		{3, 7, "011"},
	}

	for i, c := range cases {
		bb := NewBuilder(c.treeHeight)
		got := bb.IndexStr(c.pos)
		ta.Equal(c.want, got, "%d-th: case: %+v", i+1, c)
	}
}
