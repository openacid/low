package bmtree

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSize(t *testing.T) {

	ta := require.New(t)

	cases := []struct {
		input int32
		want  int32
	}{
		{0, 2},
		{1, 4},
		{2, 8},
		{3, 16},
		{29, 1 << 30},
	}

	for i, c := range cases {
		got := Size(c.input)
		ta.Equal(c.want, got, "%d-th: case: %+v", i+1, c)
	}
}
