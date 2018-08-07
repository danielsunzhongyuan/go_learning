package main

/*
Given a binary tree, return the bottom-up level order traversal of its nodes' values. (ie, from left to right, level by level from leaf to root).

For example:
Given binary tree [3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7
return its bottom-up level order traversal as:
[
  [15,7],
  [9,20],
  [3]
]
*/

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return make([][]int, 0)
	}
	result := make([][]int, 0)
	var tmpNumber []int
	var tmp []*TreeNode
	var level []*TreeNode
	tmpNumber = append(tmpNumber, root.Val)
	result = append(result, tmpNumber)
	level = append(level, root)
	for len(level) > 0 {
		tmpNumber = make([]int, 0)
		tmp = make([]*TreeNode, 0)
		for _, node := range level {
			if node.Left != nil {
				tmp = append(tmp, node.Left)
				tmpNumber = append(tmpNumber, node.Left.Val)
			}
			if node.Right != nil {
				tmp = append(tmp, node.Right)
				tmpNumber = append(tmpNumber, node.Right.Val)
			}
		}
		level = tmp
		if len(tmpNumber) > 0 {
			result = append(result, tmpNumber)
		}
	}
	i, j := 0, len(result)-1
	for i < j {
		result[i], result[j] = result[j], result[i]
		i++
		j--
	}
	return result
}
