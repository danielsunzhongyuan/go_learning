package main

/*
查找一个树上两个节点的第一个共同的祖先节点。

Attempt #1:
如果每个node里有一个指向parent的链接，那么向上查节点 p 和节点 q 的parent，直到它俩相交

Attempt #2:
从 root 节点找起，找到第一个节点，使得p和q分别在这个节点的两边（在两边的说法其实是不准确的，准确的说法是不在同一边，这样就比较容易解释下面的特殊例子了）。

特殊情况是 p 是 q 的祖先节点，那么 p 和 q 就不是在 p 的同一边了。（q 是 p 的祖先节点也类似）

Attempt #3:
我们直到，对于任意节点 r：
1. 如果p在它的一边，而q在另一边，那么r就是 first common ancestor
2. 否则，就在r的左边或右边

未完待续！这个地方有点乱。。。
*/

// Attempt #2
func commonAncestor(root, p, q *TreeNode) *TreeNode {
	if covers(root.Left, p) && covers(root.Left, q) {
		return commonAncestor(root.Left, p, q)
	}
	if covers(root.Right, p) && covers(root.Right, q) {
		return commonAncestor(root.Right, p, q)
	}
	return root
}

func covers(root, p *TreeNode) bool {
	if root == nil {
		return false
	}
	if root == p {
		return true
	}
	return covers(root.Left, p) || covers(root.Right, p)
}
