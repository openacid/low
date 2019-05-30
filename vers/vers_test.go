package vers

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsCompatible(t *testing.T) {

	ta := require.New(t)

	cases := []struct {
		input string
		spec  []string
		want  bool
	}{
		{"1.0.0", []string{"==1.0.0"}, true},
		{"1.0.0", []string{"==1.0.1"}, false},
		{"1.0.0", []string{"==1.0.0", "==1.0.1"}, true},
		{"1.2.3", []string{`>1.0.0 <2.0.0`, `>3.0.0 !4.2.1`}, true},
		{"1.9.9", []string{`>1.0.0 <2.0.0`, `>3.0.0 !4.2.1`}, true},
		{"3.1.1", []string{`>1.0.0 <2.0.0`, `>3.0.0 !4.2.1`}, true},
		{"2.1.1", []string{`>1.0.0 <2.0.0`, `>3.0.0 !4.2.1`}, false},
		{"4.2.1", []string{`>1.0.0 <2.0.0`, `>3.0.0 !4.2.1`}, false},
	}

	for i, c := range cases {
		got := IsCompatible(c.input, c.spec)
		ta.Equal(c.want, got, "%d-th: case: %+v", i+1, c)
	}
}
