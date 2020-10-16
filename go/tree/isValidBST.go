package main

import "fmt"

// https://leetcode-cn.com/problems/validate-binary-search-tree/

func isValidBST(root *TreeNode) bool {
	var f func(r *TreeNode) (int, int, bool)
	f = func(r *TreeNode) (int, int, bool) {
		if r == nil {
			return 0, 0, true
		}

		min, max := r.Val, r.Val
		if r.Left != nil {
			childMin, childMax, flag := f(r.Left)
			if !flag || r.Val <= childMax {
				return 0, 0, false
			}
			min = childMin
		}
		if r.Right != nil {
			childMin, childMax, flag := f(r.Right)
			if !flag || r.Val >= childMin {
				return 0, 0, false
			}
			max = childMax
		}
		return min, max, true
	}
	_, _, flag := f(root)
	return flag
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
