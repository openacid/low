package typehelper

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestToSlice(t *testing.T) {

	ta := require.New(t)

	ta.Panics(func() { ToSlice(1) }, "int")
	ta.Panics(func() { ToSlice(nil) }, "nil")

	cases := []struct {
		input interface{}
		want  []interface{}
	}{{
		[]int{1, 2, 3},
		[]interface{}{1, 2, 3},
	}}

	for i, c := range cases {
		rst := ToSlice(c.input)
		ta.Equal(c.want, rst, "%d-th: input: %v",
			i+1, c.input)
	}
}
