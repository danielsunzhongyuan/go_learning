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

func FindNthLast(head *LinkedListNode, n int) *LinkedListNode {
	if head == nil || n < 1 {
		return nil
	}
	cur, nthLatter := head, head
	for n > 0 && nthLatter != nil {
		nthLatter = nthLatter.Next
		n -= 1
	}
	if n > 0 && nthLatter == nil {
		return nil
	}
	for nthLatter != nil {
		nthLatter = nthLatter.Next
		cur = cur.Next
	}
	return cur
}

func main() {
	x1 := LinkedListNode{Data: 3, Next: nil}
	x2 := LinkedListNode{Data: 2, Next: nil}
	x3 := LinkedListNode{Data: 1, Next: nil}

	x1.Next = &x2
	x2.Next = &x3
	Traversal(&x1)
	fmt.Println("####")
	nthNode := FindNthLast(&x1, 3)
	fmt.Println(nthNode)
}
