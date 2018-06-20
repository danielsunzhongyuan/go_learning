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

func (s *Stack) Traversal() {
	node := s.Top
	for node != nil {
		fmt.Println(node.Data)
		node = node.Next
	}
	fmt.Println("### OVER ###")
}

type MyQueue struct {
	StackForPush *Stack
	StackForPop  *Stack
}

func (mq *MyQueue) Add(data int) {
	mq.StackForPush.Push(Node{Data: data})
}

func (mq *MyQueue) Remove() *Node {
	if !mq.StackForPop.IsEmpty() {
		return mq.StackForPop.Pop()
	}
	for !mq.StackForPush.IsEmpty() {
		mq.StackForPop.Push(*mq.StackForPush.Pop())
	}
	return mq.StackForPop.Pop()
}

func (mq *MyQueue) Peek() *Node {
	if !mq.StackForPop.IsEmpty() {
		return mq.StackForPop.Peek()
	}
	for !mq.StackForPush.IsEmpty() {
		mq.StackForPop.Push(*mq.StackForPush.Pop())
	}
	return mq.StackForPop.Peek()
}

func main() {
	mq := MyQueue{StackForPush: new(Stack), StackForPop: new(Stack)}
	mq.Add(1)
	fmt.Println(mq.Peek())
	mq.Add(2)
	mq.Add(3)
	fmt.Println(mq.Peek())
	fmt.Println(mq.StackForPush.Top, mq.StackForPop.Top)
	fmt.Println("remove:", mq.Remove())
	fmt.Println(mq.StackForPush.Top, mq.StackForPop.Top)
	fmt.Println(mq.Peek())
	fmt.Println(mq.StackForPush.Top, mq.StackForPop.Top)
	fmt.Println(mq.Remove())
	fmt.Println(mq.StackForPush.Top, mq.StackForPop.Top)
	fmt.Println(mq.Remove())
	fmt.Println(mq.StackForPush.Top, mq.StackForPop.Top)
	fmt.Println(mq.Remove())
	fmt.Println(mq.StackForPush.Top, mq.StackForPop.Top)
}
