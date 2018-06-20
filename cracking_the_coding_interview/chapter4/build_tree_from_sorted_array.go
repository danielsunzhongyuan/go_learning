package main

import "fmt"

type TreeNode struct {
	Data  int
	Left  *TreeNode
	Right *TreeNode
}

type Tree struct {
	Root *TreeNode
}

func addToTree(arr []int, start, end int) *TreeNode {
	if end < start {
		return nil
	}

	mid := (start + end) / 2
	node := &TreeNode{arr[mid], nil, nil}
	node.Left = addToTree(arr, start, mid-1)
	node.Right = addToTree(arr, mid+1, end)
	return node
}

func createMinimalBST(arr []int) *TreeNode {
	return addToTree(arr, 0, len(arr)-1)
}
func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8}
	root := createMinimalBST(a)
	fmt.Println(root.Data, root.Left.Data, root.Right.Data, root.Left.Left.Data, root.Left.Right.Data, root.Right.Left.Data, root.Right.Right.Data, root.Right.Right.Right.Data)
}
