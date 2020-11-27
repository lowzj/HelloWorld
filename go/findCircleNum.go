package main

import (
	"fmt"
)

// https://leetcode-cn.com/problems/friend-circles/
// 547. 朋友圈 | Friend Circles
//------------------------------------------------------------------------------

//------------------------------------------------------------------------------
// Solution
//
// 并查集
//
// 复杂度分析:
//   * 时间: O(N^3)
//   * 空间: O(N)
func findCircleNum(M [][]int) int {
	N := len(M)

	us, cnt := make([]int, N), N
	for i := 0; i < N; i++ {
		us[i] = i
	}
	find := func(p int) int {
		for ; p != us[p]; p = us[p] {
			us[p] = us[us[p]]
		}
		return p
	}
	union := func(p, q int) {
		rootP := find(p)
		rootQ := find(q)
		if rootP == rootQ {
			return
		}
		us[rootQ] = rootP
		cnt--
	}

	for i := 0; i < N-1; i++ {
		for j := i + 1; j < N; j++ {
			if M[i][j] == 1 {
				union(i, j)
			}
		}
	}
	return cnt
}

//------------------------------------------------------------------------------
// main

func main() {
	cases := [][][]int{
		{
			{1, 1, 0},
			{1, 1, 0},
			{0, 0, 1},
		},
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(findCircleNum(c))
	}
}
