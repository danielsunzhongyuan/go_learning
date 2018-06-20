package main

/*
1个很大的树 T1（百万级别），和另一个树T2，判断T2是不是T1的子树？

Attempt #1：
把 T1和T2 各自进行前序和中序遍历，如果T1的前序遍历包含T2的前序遍历，且，T1的中序遍历包含T2的中序遍历，那么T2就是T1的子树。（因为前序+中序就可以确定一棵树了）
这里可以使用到 suffix tree （？），但是量太大，内存受不了。

Attempt #2：
如下
*/

// Attempt #2
func containsTree(t1, t2 *TreeNode) bool {
	if t2 == nil {
		return true
	}
	return subTree(t1, t2)
}

func subTree(r1, r2 *TreeNode) bool {
	if r1 == nil {
		return false
	}
	if r1.Data == r2.Data {
		if matchTree(r1, r2) {
			return true
		}
	}
	return subTree(r1.Left, r2) || subTree(r1.Right, r2)
}

func matchTree(r1, r2 *TreeNode) bool {
	if r2 == nil && r1 == nil {
		return true
	}
	if r1 == nil || r2 == nil {
		return false
	}
	if r1.Data != r2.Data {
		return false
	}
	return matchTree(r1.Left, r2.Left) && matchTree(r1.Right, r2.Right)
}
