package main

import "fmt"

// https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/

func buildTree(preorder []int, inorder []int) *TreeNode {
	n, m := len(preorder), len(preorder)
	if n == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	i := 0
	for ; i < m && inorder[i] != root.Val; i++ {
	}
	if i > 0 {
		root.Left = buildTree(preorder[1:1+i], inorder[:i])
	}
	if i < m-1 {
		root.Right = buildTree(preorder[1+i:], inorder[i+1:])
	}
	return root
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
