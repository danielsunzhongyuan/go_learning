package merge

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeKLists(lists []*ListNode) *ListNode {
	length := len(lists)
	if length == 0 {
		return nil
	} else if length == 1 {
		return lists[0]
	} else if length == 2 {
		return mergeTwoLists(lists)
	}

	B := MergeKLists(lists[length/2:])
	A := MergeKLists(lists[:length/2])
	return mergeTwoLists([]*ListNode{A, B})
}

func mergeKLists(lists []*ListNode) *ListNode {
	ret := &ListNode{Val: 0, Next: nil}
	runner := ret

	next := getSmallestNode(lists)
	for next != nil {
		runner.Next = next
		runner = runner.Next
		next = getSmallestNode(lists)
	}

	return ret.Next
}

func mergeTwoLists(lists []*ListNode) *ListNode {
	ret := &ListNode{Val: 0, Next: nil}
	runner := ret

	next := getSmallestNode(lists)
	for next != nil {
		runner.Next = next
		runner = runner.Next
		next = getSmallestNode(lists)
	}

	return ret.Next
}

func getSmallestNode(lists []*ListNode) *ListNode {
	tmp := 10000000
	index := -1
	for ind, node := range lists {
		if node != nil && node.Val < tmp {
			tmp = node.Val
			index = ind
		}
	}
	if index == -1 {
		return nil
	}
	ret := lists[index]
	lists[index] = lists[index].Next
	return ret
}
