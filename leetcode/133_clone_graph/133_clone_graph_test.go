package leetcode

import (
	"testing"

	. "github.com/wtchangdm/leetcode-go/structs"
)

func aggregate(node *Node) []*Node {
	result := []*Node{}

	cached := map[*Node]struct{}{}

	var visit func(n *Node)
	visit = func(n *Node) {
		if _, found := cached[n]; found {
			return
		}

		cached[n] = struct{}{}
		result = append(result, n)
	}

	visit(node)
	return result
}

func TestCloneGraph(t *testing.T) {
	node1 := &Node{Val: 1}
	node2 := &Node{Val: 2}
	node3 := &Node{Val: 3}
	node4 := &Node{Val: 4}

	node1.Neighbors = []*Node{node2, node4}
	node2.Neighbors = []*Node{node1, node3}
	node3.Neighbors = []*Node{node2, node4}
	node4.Neighbors = []*Node{node1, node3}

	result := aggregate(cloneGraph(node1))
	original := aggregate(node1)

	for i := range result {
		r := result[i]
		o := original[i]
		if r == o {
			t.Error("Graph node is not cloned")
		}

		if r.Val != o.Val {
			t.Errorf("%d != %d", r.Val, o.Val)
		}
	}
}
