package leetcode

import . "github.com/wtchangdm/leetcode-go/structs"

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	mapping := map[*Node]*Node{}

	var clone func(n *Node) *Node
	clone = func(n *Node) *Node {
		if _, found := mapping[n]; !found {
			newNode := &Node{Val: n.Val}

			mapping[n] = newNode
			for _, neighbor := range n.Neighbors {
				newNode.Neighbors = append(newNode.Neighbors, clone(neighbor))
			}
		}

		return mapping[n]
	}

	return clone(node)
}
