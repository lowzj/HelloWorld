package main

import (
	"fmt"
)

// https://leetcode-cn.com/problems/subsets/
// 78. 子集 | Subsets
//------------------------------------------------------------------------------

//------------------------------------------------------------------------------
// Solution
func subsets(nums []int) [][]int {
	N := len(nums)

	res := [][]int{{}}
	var dfs func(start, cnt int)
	values := make([]int, N)
	dfs = func(start, cnt int) {
		if start >= N || cnt == N {
			return
		}
		for i := start; i < N; i++ {
			values[cnt] = nums[i]
			tmp := make([]int, cnt+1)
			copy(tmp, values[:cnt+1])
			res = append(res, tmp)
			dfs(i+1, cnt+1)
		}
	}
	dfs(0, 0)
	return res
}

//------------------------------------------------------------------------------
// main

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
