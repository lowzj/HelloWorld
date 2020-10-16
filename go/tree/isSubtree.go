package main

import (
	"fmt"
)

// https://leetcode-cn.com/problems/subtree-of-another-tree/

func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil || t == nil {
		return false
	}

	validate := func(node *TreeNode) bool {
		ss, ps := &stack{}, node
		st, pt := &stack{}, t
		for ; ps != nil && pt != nil; ps, pt = ss.pop(), st.pop() {
			if ps.Val != pt.Val {
				return false
			}
			if pt.Left == nil && ps.Left != nil ||
				pt.Left != nil && ps.Left == nil {
				return false
			}
			if pt.Right == nil && ps.Right != nil ||
				pt.Right != nil && ps.Right == nil {
				return false
			}

			ss.push(ps.Left)
			st.push(pt.Left)
			ss.push(ps.Right)
			st.push(pt.Right)
		}
		return ps == pt
	}

	ss, ps := &stack{}, s
	for ; ps != nil; ps = ss.pop() {
		if ps.Val == t.Val && validate(ps) {
			return true
		}
		ss.push(ps.Left)
		ss.push(ps.Right)
	}
	return false
}

type stack struct {
	s []*TreeNode
}

func (s *stack) push(node *TreeNode) {
	if node == nil {
		return
	}
	s.s = append(s.s, node)
}

func (s *stack) pop() *TreeNode {
	n := len(s.s)
	if n > 0 {
		top := s.s[n-1]
		s.s = s.s[:n-1]
		return top
	}
	return nil
}

func main() {
	cases := [][][]int{
		{
			{3, 4, 5, 1, 2},
			{4, 1, 2},
		},
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		s := generateTree(c[0])
		t := generateTree(c[1])
		fmt.Println(isSubtree(s, t))
	}
}
