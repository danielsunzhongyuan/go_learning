package merge

import (
	"fmt"
	"testing"
)

func TestMergeKLists(t *testing.T) {
	lists := []*ListNode{generateList([]int{1, 4, 5}), generateList([]int{1, 3, 4}), generateList([]int{2, 6})}

	x := MergeKLists(lists)
	traversal(x)
}

func generateList(arr []int) *ListNode {
	root := &ListNode{Val: 0, Next: nil}
	for i := len(arr) - 1; i >= 0; i-- {
		tmp := &ListNode{Val: arr[i], Next: nil}
		tmp.Next = root.Next
		root.Next = tmp
	}
	return root.Next
}

func traversal(head *ListNode) {
	tmp := head
	for tmp != nil {
		fmt.Print(tmp.Val)
		tmp = tmp.Next
	}
	fmt.Println("")
}
