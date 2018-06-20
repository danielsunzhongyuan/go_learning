package main

import "fmt"

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

func (s *Stack) Sort() Stack {
	stack2 := Stack{}
	for !s.IsEmpty() {
		tmp := s.Pop()
		for !stack2.IsEmpty() && stack2.Peek().Data < tmp.Data {
			s.Push(*stack2.Pop())
		}
		stack2.Push(*tmp)
	}
	return stack2
}

func (s *Stack) Traversal() {
	node := s.Top
	for node != nil {
		fmt.Println(node.Data)
		node = node.Next
	}
	fmt.Println("### OVER ###")
}

func main() {
	s := Stack{}
	s.Push(Node{Data: 3})
	s.Push(Node{Data: 1})
	s.Push(Node{Data: 1})
	s.Push(Node{Data: 2})
	s.Push(Node{Data: 2})
	s.Push(Node{Data: 4})
	s.Traversal()
	s2 := s.Sort()
	s2.Traversal()
}
