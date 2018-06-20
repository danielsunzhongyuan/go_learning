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

func levelTraversal(node TreeNode) [][]int {
	ret := [][]int{}
	queue := []TreeNode{}
	queue = append(queue, node)
	for len(queue) > 0 {
		size := len(queue)
		tmp := make([]int, 0, size)
		for size > 0 {
			node := queue[0]
			tmp = append(tmp, node.Data)
			if node.Left != nil {
				queue = append(queue, *node.Left)
			}
			if node.Right != nil {
				queue = append(queue, *node.Right)
			}
			queue = queue[1:]

			size -= 1
		}
		ret = append(ret, tmp)
	}
	return ret
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

	/*
	        1
	      /   \
	     2     3
	    / \   /
	   4   5 6
	      /   \
	     7     8
	*/
	node1.Left = node2
	node1.Right = node3
	node2.Left = node4
	node2.Right = node5
	node5.Right = node7
	node3.Left = node6
	node6.Right = node8

	fmt.Println(levelTraversal(*node1))
}
