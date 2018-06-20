package main

import "fmt"

/*
在二分查找树里，某个节点的下一个节点就是中序遍历后这个节点的下一个节点，也是比这个节点数大的最小的那个数。
这里需要存储每个节点的 parent 节点信息
*/

type TreeNode struct {
	Data   int
	Left   *TreeNode
	Right  *TreeNode
	Parent *TreeNode
}

func inorderSucc(node *TreeNode) *TreeNode {
	if node != nil {
		p := new(TreeNode)
		if node.Parent == nil || node.Right != nil {
			p = leftMostChild(node.Right)
		} else {
			p = node.Parent
			for p != nil {
				if p.Left == node {
					break
				}
				node = p
				p = node.Parent
			}
		}
		return p
	}
	return nil
}

func leftMostChild(node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}
	for node.Left != nil {
		node = node.Left
	}
	return node
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
	node1.Data = 5
	node2.Data = 2
	node3.Data = 8
	node4.Data = 1
	node5.Data = 4
	node6.Data = 6
	node7.Data = 3
	node8.Data = 7

	/*
	        5
	      /   \
	     2     8
	    / \   /
	   1   4 6
	      /   \
	     3     7
	*/
	node1.Left = node2
	node1.Right = node3
	node1.Parent = nil
	node2.Parent = node1
	node3.Parent = node1

	node2.Left = node4
	node2.Right = node5
	node4.Parent = node2
	node5.Parent = node2

	node5.Left = node7
	node7.Parent = node5

	node3.Left = node6
	node6.Parent = node3

	node6.Right = node8
	node8.Parent = node6

	fmt.Println(inorderSucc(node8))
}
