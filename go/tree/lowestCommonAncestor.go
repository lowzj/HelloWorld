package main

// https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if p == q {
		return p
	}
	if p == nil {
		return q
	}
	if q == nil {
		return p
	}

	res, level := (*TreeNode)(nil), 0
	var view func(node *TreeNode, l int) int
	view = func(node *TreeNode, l int) int {
		if node == nil {
			return 0
		}

		count := 0
		if node == p || node == q {
			count++
		}
		if count < 2 && node.Left != nil {
			count += view(node.Left, l+1)
		}
		if count < 2 && node.Right != nil {
			count += view(node.Right, l+1)
		}
		if count >= 2 && l > level {
			res = node
			level = l
		}
		return count
	}
	view(root, 1)
	return res
}
