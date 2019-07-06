package tree_test

import (
	"fmt"
	"testing"

	"github.com/openacid/low/tree"
	"github.com/stretchr/testify/require"
)

type T struct {
	branches map[int][]int
}

func nodeID(node interface{}) int {
	if node == nil {
		return 0
	}
	return node.(int)
}

func (t *T) Child(node, label interface{}) interface{} {
	if node == nil && label == nil {
		// root
		return 0
	}
	if children, ok := t.branches[nodeID(node)]; ok {
		return children[label.(int)]
	}
	return nil
}

func (t *T) Labels(node interface{}) []interface{} {
	rst := []interface{}{}
	if children, ok := t.branches[nodeID(node)]; ok {
		for i := 0; i < len(children); i++ {
			rst = append(rst, i)
		}
	}
	return rst
}

func (t *T) NodeID(node interface{}) string {
	return fmt.Sprintf("%02d", nodeID(node))
}

func (t *T) NodeInfo(node interface{}) string {
	return "(foo)"
}

func (t *T) LabelInfo(label interface{}) string {
	return fmt.Sprintf("%d", label)
}

func (t *T) LeafVal(node interface{}) (interface{}, bool) {
	if _, ok := t.branches[nodeID(node)]; ok {
		return nil, false
	}
	return "leaf", true
}

func TestToString(t *testing.T) {
	// 0 --> 1 --> 3 -->5
	//               `->6
	//   `-> 2 --> 4
	tr := &T{
		branches: map[int][]int{
			0: {1, 2},
			1: {3},
			2: {4},
			3: {5, 6},
		},
	}

	rst := tree.String(tr)
	want := `
#00(foo)*2
   -0->#01(foo)
          -0->#03(foo)*2
                 -0->#05(foo)=leaf
                 -1->#06(foo)=leaf
   -1->#02(foo)
          -0->#04(foo)=leaf`[1:]
	if want != rst {
		t.Fatalf("expect: \n%v\n; but: \n%v\n", want, rst)
	}
}

func TestDepthFirst(t *testing.T) {

	ta := require.New(t)

	tr := &T{
		branches: map[int][]int{
			0: {1, 2},
			1: {3},
			2: {4},
			3: {5, 6},
		},
	}

	got := []string{}

	tree.DepthFirst(tr, func(tr tree.Tree, parent, label, node interface{}) {
		var b int
		if label == nil {
			b = 0
		} else {
			b = label.(int)
		}
		s := fmt.Sprintf("p:%s-b:%d->n:%s", tr.NodeID(parent), b, tr.NodeID(node))
		got = append(got, s)
	})

	want := []string{
		"p:03-b:0->n:05",
		"p:03-b:1->n:06",
		"p:01-b:0->n:03",
		"p:00-b:0->n:01",
		"p:02-b:0->n:04",
		"p:00-b:1->n:02",
		"p:00-b:0->n:00",
	}

	ta.Equal(want, got)
}
