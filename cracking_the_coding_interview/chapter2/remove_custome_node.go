package main

import "fmt"

type LinkedListNode struct {
	Data int
	Next *LinkedListNode
}

func Traversal(head *LinkedListNode) {
	node := head
	for node != nil {
		fmt.Println(node.Data)
		node = node.Next
	}
}

func RemoveCustomNode(head *LinkedListNode, node *LinkedListNode) {
	if node == nil {
		return
	}

	for head.Next != node && head.Next != nil {
		head = head.Next
	}
	if head.Next == node {
		head.Next = node.Next
		node.Next = nil
		node = new(LinkedListNode)
	}
}

func main() {
	x1 := LinkedListNode{Data: 3, Next: nil}
	x2 := LinkedListNode{Data: 2, Next: nil}
	x3 := LinkedListNode{Data: 1, Next: nil}

	x1.Next = &x2
	x2.Next = &x3
	Traversal(&x1)
	fmt.Println("####")
	RemoveCustomNode(&x1, &x3)
	Traversal(&x1)
}
