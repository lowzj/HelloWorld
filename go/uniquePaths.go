package main

import (
	"fmt"
)

// https://leetcode-cn.com/problems/unique-paths/

//------------------------------------------------------------------------------
// 就是求组合数

func uniquePaths(m int, n int) int {
	comp := func(n, m int) int {
		res := 1
		if n-m > m {
			m = n - m
		}
		for i, k := n, m; i > n-m; i-- {
			res *= i
			for ; k > 0 && res%k == 0; k-- {
				res /= k
			}
		}
		return res
	}
	return comp(m+n-2, m-1)
}

func main() {
	cases := [][]int{
		{100, 30},
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(uniquePaths(c[0], c[1]))
	}
}
