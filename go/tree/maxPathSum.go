package main

import (
	"fmt"
)

// https://leetcode-cn.com/problems/binary-tree-maximum-path-sum/

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	maxSum := root.Val
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	// calculate the max sum containing node.Val
	var mf func(node *TreeNode) int
	mf = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		res := node.Val
		left := mf(node.Left)
		right := mf(node.Right)
		res = max(res, res+max(left, right))
		maxSum = max(maxSum, max(res, node.Val+left+right))
		return res
	}
	mf(root)
	return maxSum
}

func main() {
	cases := [][]int{
		{},
		{},
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(c)
	}
}
