package main

import "fmt"

type Node struct {
	Value int
	Right *Node
	Left  *Node
}

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

	proceed(tree)
}

func proceed(node *Node) {
	if node == nil {
		return
	}
	q := []*Node{node}

	for len(q) != 0 {
		k := len(q)
		for i := 0; i < k; i++ {
			cur := q[i]
			fmt.Print(cur.Value)

			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
		}
		q = q[k:]
	}
}
