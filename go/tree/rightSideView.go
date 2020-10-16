package main

// https://leetcode-cn.com/problems/binary-tree-right-side-view/

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var res []int
	var view func(node *TreeNode, level int)
	viewedLevel := 0
	view = func(node *TreeNode, level int) {
		if node == nil {
			return
		}
		if level >= viewedLevel {
			res = append(res, node.Val)
			viewedLevel++
		}
		if node.Right != nil {
			view(node.Right, level+1)
		}
		if node.Left != nil {
			view(node.Left, level+1)
		}
	}
	view(root, 0)
	return res
}
