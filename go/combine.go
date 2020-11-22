package main

import (
	"fmt"
)

// https://leetcode-cn.com/problems/combinations/
//------------------------------------------------------------------------------

func combine(n int, k int) [][]int {
	st, top := make([]int, n*n), 0
	push := func(v, cnt int) {
		st[top] = v + (n+1)*cnt
		top++
	}
	pop := func() (int, int) {
		top--
		v, cnt := st[top]%(n+1), st[top]/(n+1)
		return v, cnt
	}

	var res [][]int
	comb := make([]int, k)
	push(0, -1)
	for top > 0 {
		v, cnt := pop()
		if cnt >= 0 {
			comb[cnt] = v
		}
		if cnt == k-1 {
			tmp := make([]int, k)
			copy(tmp, comb)
			res = append(res, tmp)
			continue
		}
		for i := v + 1; i <= n; i++ {
			push(i, cnt+1)
		}
	}
	return res
}
func combine0(n int, k int) [][]int {
	comb := make([]int, k)
	var res [][]int
	add := func() {
		tmp := make([]int, k)
		copy(tmp, comb)
		res = append(res, tmp)
	}
	var dfs func(start, cnt int)
	dfs = func(start, cnt int) {
		if cnt == k {
			add()
			return
		}
		for i := start; i <= n; i++ {
			comb[cnt] = i
			dfs(i+1, cnt+1)
		}
	}
	dfs(1, 0)
	return res
}

func main() {
	cases := [][]int{
		{4, 2},
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(combine0(c[0], c[1]))
	}
}
