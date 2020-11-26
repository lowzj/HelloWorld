package main

import (
	"fmt"
)

// https://leetcode-cn.com/problems/kth-smallest-element-in-a-bst/
// 230. 二叉搜索树中第K小的元素 | Kth Smallest Element in a BST
//------------------------------------------------------------------------------
func kthSmallest(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}

	res := 0
	var dfs func(node *TreeNode, k int) int
	dfs = func(node *TreeNode, k int) int {
		if node == nil {
			return 0
		}
		cnt := dfs(node.Left, k)
		if cnt >= k {
			return cnt
		}
		if cnt++; cnt == k {
			res = node.Val
			return cnt
		}
		cnt += dfs(node.Right, k-cnt)
		return cnt
	}
	dfs(root, k)
	return res
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
