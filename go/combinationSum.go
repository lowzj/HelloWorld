package main

import (
	"fmt"
	"sort"
)

// https://leetcode-cn.com/problems/combination-sum/

func combinationSum(candidates []int, target int) [][]int {
	N := len(candidates)
	if N == 0 {
		return nil
	}

	sort.Ints(candidates)

	var res [][]int
	count := make([]int, N)
	add := func(n int) {
		var tmp []int
		for i := 0; i < n; i++ {
			for j := 0; j < count[i]; j++ {
				tmp = append(tmp, candidates[i])
			}
		}
		res = append(res, tmp)
	}
	var dfs func(start, target int)
	dfs = func(start, target int) {
		if target == 0 {
			add(start)
		}
		for i := start; i < N; i++ {
			if target < candidates[i] {
				return
			}
			cnt := target / candidates[i]
			for j := 1; j <= cnt; j++ {
				count[i] = j
				dfs(i+1, target-candidates[i]*j)
			}
			count[i] = 0
		}
	}
	dfs(0, target)
	return res
}

func main() {
	cases := []struct {
		nums   []int
		target int
	}{
		{[]int{2, 3, 5}, 8},
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(combinationSum(c.nums, c.target))
	}
}
