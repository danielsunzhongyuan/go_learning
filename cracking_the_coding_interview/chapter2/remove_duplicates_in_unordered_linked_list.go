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

func DeleteDups(head *LinkedListNode) {
	table := make(map[int]bool)
	previousNode := new(LinkedListNode)
	for head != nil {
		if table[head.Data] {
			previousNode.Next = head.Next
		} else {
			table[head.Data] = true
			previousNode = head
		}
		head = head.Next
	}
}

func DeleteDupsWithoutMap(head *LinkedListNode) {
	previous := head
	current := previous.Next
	for current != nil {
		runner := head
		for runner != current {
			if runner.Data == current.Data {
				tmp := current.Next
				previous.Next = tmp
				current = tmp
				break
			}
			runner = runner.Next
		}
		if runner == current {
			previous = current
			current = current.Next
		}
	}
}

func main() {
	x1 := LinkedListNode{Data: 2, Next: nil}
	x2 := LinkedListNode{Data: 2, Next: nil}
	x3 := LinkedListNode{Data: 1, Next: nil}

	x1.Next = &x2
	x2.Next = &x3
	Traversal(&x1)
	fmt.Println("####")
	DeleteDups(&x1)
	Traversal(&x1)

	b := []int{1, 2, 3}
	for i, n := range b {
		fmt.Println(i, n)
	}
}
