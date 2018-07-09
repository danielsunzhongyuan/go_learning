package main

import "fmt"

/*
Given a linked list and a value x, partition it such that all nodes less than x come before nodes greater than or equal to x.

You should preserve the original relative order of the nodes in each of the two partitions.

Example:
Input: head = 1->4->3->2->5->2, x = 3
Output: 1->2->2->4->3->5
*/

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	firstLesser := &ListNode{}
	firstGreater := &ListNode{}
	lesser, greater := firstLesser, firstGreater
	runner := head
	for runner != nil {
		if runner.Val < x {
			lesser.Next = runner
			lesser = lesser.Next
		} else {
			greater.Next = runner
			greater = greater.Next
		}
		runner = runner.Next
	}
	if firstLesser.Next == nil || firstGreater.Next == nil {
		return head
	} else {
		greater.Next = nil
		lesser.Next = firstGreater.Next
		return firstLesser.Next
	}
}

func printList(head *ListNode) {
	runner := head
	for runner != nil {
		fmt.Print(runner.Val, " ")
		runner = runner.Next
	}
	fmt.Println()
}

func main() {
	root := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val:  2,
							Next: nil}}}}}}
	printList(root)
	printList(partition(root, 3))
}
