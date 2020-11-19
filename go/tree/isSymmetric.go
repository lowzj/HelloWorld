package main

import (
	"fmt"
)

// https://leetcode-cn.com/problems/symmetric-tree/

//------------------------------------------------------------------------------

func isSymmetric(root *TreeNode) bool {
	var validate func(a, b *TreeNode) bool
	validate = func(a, b *TreeNode) bool {
		if a == nil && b == nil {
			return true
		}
		if a == nil || b == nil || a.Val != b.Val {
			return false
		}
		return validate(a.Left, b.Right) && validate(a.Right, b.Left)
	}
	return root == nil || validate(root.Left, root.Right)
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
