package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTree(nodes []int) *TreeNode {
	const notNode = math.MaxInt64
	n := len(nodes)
	if n < 1 {
		return nil
	}

	res := make([]*TreeNode, n)
	l := func(p int) int { return 2*p + 1 }
	r := func(p int) int { return 2*p + 2 }
	for i := 0; i < n; i++ {
		if nodes[i] == notNode {
			continue
		}
		res[i] = &TreeNode{Val: nodes[i]}
	}
	for i := 0; i < (n+1)/2; i++ {
		if res[i] == nil {
			continue
		}
		if l(i) < n {
			res[i].Left = res[l(i)]
		}
		if r(i) < n {
			res[i].Right = res[r(i)]
		}
	}
	return res[0]
}
