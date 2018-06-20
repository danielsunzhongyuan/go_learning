package tree

import (
	"testing"
)

func TestPreOrderTraversalRecursion(t *testing.T) {
	a := []int{4, 2, 3, 1, 7, 5, 6, 8}
	root := buildTree(a)

	PreOrderTraversalRecursion(root)
	InOrderTraversalRecursion(root)
	PostOrderTraversalRecursion(root)
}

func buildTree(arr []int) *Node {
	var root *Node
	for _, num := range arr {
		root = insertNode(root, num)
	}
	return root
}

func insertNode(node *Node, num int) *Node {
	if node == nil {
		node = &Node{Data: num}
		return node
	}

	if node.Data < num {
		node.Right = insertNode(node.Right, num)
	} else {
		node.Left = insertNode(node.Left, num)
	}
	return node
}
