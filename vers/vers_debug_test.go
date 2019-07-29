// +build debug

package vers

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCheck_invalid(t *testing.T) {

	ta := require.New(t)

	cases := []struct {
		input string
		spec  []string
	}{

		{"abc.e", []string{"1.0.0"}},
		{"1.0.0", []string{"a.b.c"}},
		{"1.0.0", []string{"1.2.3", "a.b.c"}},
	}

	for _, c := range cases {
		ta.Panics(func() { Check(c.input, c.spec...) })
	}
}
