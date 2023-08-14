package main

import "fmt"

func main() {
	tree := &Node{
		Value: 1,
		Left: &Node{
			Value: 4,
			Left:  &Node{Value: 7, Left: nil, Right: nil},
			Right: &Node{Value: 9, Left: nil, Right: nil}},
		Right: &Node{
			Value: 5,
			Left:  &Node{Value: 2, Left: nil, Right: nil},
			Right: nil}}

	start(tree)
}

func start(node *Node) {
	if node == nil {
		return
	}
	fmt.Print(node.Value)
	goWidth(node)
}

func goWidth(node *Node) {
	if node == nil {
		return
	}

	if node.Left != nil {
		fmt.Print(node.Left.Value)
	}
	if node.Right != nil {
		fmt.Print(node.Right.Value)
	}

	goWidth(node.Left)
	goWidth(node.Right)

	return
}
