package main

import (
	"fmt"
)

// https://leetcode-cn.com/problems/flatten-binary-tree-to-linked-list/
// 114. 二叉树展开为链表 | Flatten Binary Tree to Linked List
//------------------------------------------------------------------------------

func flatten(root *TreeNode) {
	var dfs func(root *TreeNode) *TreeNode
	dfs = func(root *TreeNode) *TreeNode {
		if root == nil {
			return nil
		}
		if root.Left == nil && root.Right == nil {
			return root
		}
		l := dfs(root.Left)
		r := dfs(root.Right)
		if l != nil {
			root.Right, l.Right = root.Left, root.Right
			root.Left = nil
		}
		if r == nil {
			return l
		}
		return r
	}
	dfs(root)
}

func main() {
	cases := [][]int{
		{},
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(c)
	}
}
