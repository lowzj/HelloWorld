package main

// https://leetcode-cn.com/problems/diameter-of-binary-tree/

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	res := 0
	var view func(node *TreeNode) int
	view = func(n *TreeNode) int {
		if n == nil || n.Left == nil && n.Right == nil {
			return 0
		}
		maxD := 0
		left := view(n.Left)
		right := view(n.Right)
		if n.Left != nil {
			maxD += left + 1
		}
		if n.Right != nil {
			maxD += right + 1
		}
		res = max(res, maxD)
		return max(left, right) + 1
	}

	view(root)
	return res
}
