package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Data  int
	Left  *TreeNode
	Right *TreeNode
}

type Tree struct {
	Root *TreeNode
}

func (t *TreeNode) MaxDepth() int {
	if t == nil {
		return 0
	}
	return 1 + int(math.Max(float64(t.Left.MaxDepth()), float64(t.Right.MaxDepth())))
}

func (t *TreeNode) MinDepth() int {
	if t == nil {
		return 0
	}
	return 1 + int(math.Min(float64(t.Left.MaxDepth()), float64(t.Right.MaxDepth())))
}

func (t *TreeNode) IsBalanced() bool {
	return t.MaxDepth()-t.MinDepth() <= 1
}

func main() {
	node1 := new(TreeNode)
	node2 := new(TreeNode)
	node3 := new(TreeNode)
	node4 := new(TreeNode)
	node5 := new(TreeNode)
	node6 := new(TreeNode)
	node7 := new(TreeNode)
	node8 := new(TreeNode)
	node1.Data = 1
	node2.Data = 2
	node3.Data = 3
	node4.Data = 4
	node5.Data = 5
	node6.Data = 6
	node7.Data = 7
	node8.Data = 8

	node1.Left = node2
	node1.Right = node3
	node2.Left = node4
	node2.Right = node5
	fmt.Println(node1.IsBalanced())
}
