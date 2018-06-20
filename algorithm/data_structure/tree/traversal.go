package tree

import "fmt"

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

func PreOrderTraversalRecursion(node *Node) {
	if node == nil {
		return
	}
	fmt.Println(node.Data)
	PreOrderTraversalRecursion(node.Left)
	PreOrderTraversalRecursion(node.Right)
}

func InOrderTraversalRecursion(node *Node) {
	if node == nil {
		return
	}
	InOrderTraversalRecursion(node.Left)
	fmt.Println(node.Data)
	InOrderTraversalRecursion(node.Right)
}

func PostOrderTraversalRecursion(node *Node) {
	if node == nil {
		return
	}
	PostOrderTraversalRecursion(node.Left)
	PostOrderTraversalRecursion(node.Right)
	fmt.Println(node.Data)
}
