package main

type Node struct {
	Value int
	Right *Node
	Left  *Node
}

func hasPathSum(node *Node, t int) bool {
	if node == nil {
		return false
	}

	if node.Left == nil && node.Right == nil {
		return node.Value == t
	}

	if node.Left != nil {
		node.Left.Value += node.Value
	}
	if node.Right != nil {
		node.Right.Value += node.Value
	}
	hasLeft := hasPathSum(node.Left, t)
	hasRight := hasPathSum(node.Right, t)

	return hasLeft || hasRight
}
