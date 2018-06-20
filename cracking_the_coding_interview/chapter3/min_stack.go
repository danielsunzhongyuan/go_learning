package main

import (
	"fmt"
	"math"
)

type Node struct {
	Data int
	Next *Node
}

type Stack struct {
	Top *Node
}

func (s *Stack) Push(node Node) {
	node.Next = s.Top
	s.Top = &node
}

func (s *Stack) Pop() *Node {
	if !s.IsEmpty() {
		tmp := s.Top
		s.Top = s.Top.Next
		return tmp
	}
	return nil
}

func (s *Stack) Peek() *Node {
	return s.Top
}

func (s *Stack) IsEmpty() bool {
	return s.Top == nil
}

func (s *Stack) Traversal() {
	node := s.Top
	for node != nil {
		fmt.Println(node.Data)
		node = node.Next
	}
	fmt.Println("### OVER ###")
}

type MinStack struct {
	Stack
	Stack2 *Stack
}

func (ms *MinStack) Push(node Node) {
	if node.Data <= ms.Min() {
		ms.Stack2.Push(node)
	}
	ms.Stack.Push(node)
}

func (ms *MinStack) Pop() *Node {
	top := ms.Stack.Pop()
	if top.Data == ms.Min() {
		ms.Stack2.Pop()
	}
	return top
}

func (ms *MinStack) Min() int {
	if !ms.Stack2.IsEmpty() {
		return ms.Stack2.Peek().Data
	}
	return math.MaxInt64
}

func (ms *MinStack) Traversal() {
	fmt.Println("# Data:")
	ms.Stack.Traversal()
	fmt.Println("# Min:")
	ms.Stack2.Traversal()
	fmt.Println("# end")
}

func main() {
	minStack := &MinStack{Stack: Stack{}, Stack2: new(Stack)}
	minStack.Traversal()
	fmt.Println(minStack.Min())
	minStack.Push(Node{Data: 3})
	minStack.Push(Node{Data: 4})
	minStack.Push(Node{Data: 5})
	minStack.Push(Node{Data: 2})
	minStack.Push(Node{Data: 1})
	minStack.Push(Node{Data: 0})

	fmt.Println(minStack.Min())
	minStack.Pop()
	fmt.Println(minStack.Min())
	minStack.Pop()
	fmt.Println(minStack.Min())
	minStack.Pop()
	fmt.Println(minStack.Min())
	minStack.Pop()
	fmt.Println(minStack.Min())
}
