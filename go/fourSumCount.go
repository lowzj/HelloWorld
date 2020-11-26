package main

import (
	"fmt"
)

// https://leetcode-cn.com/problems/4sum-ii/
// 454. 四数相加 II | 4Sum II
//------------------------------------------------------------------------------

//------------------------------------------------------------------------------
// Solution
//
// 复杂度分析:
//   * 时间: O(N^2)
//   * 空间: O(N^2)
func fourSumCount(A []int, B []int, C []int, D []int) int {
	if len(A) == 0 {
		return 0
	}
	ab := make(map[int]int, len(A)*len(B))
	for _, a := range A {
		for _, b := range B {
			ab[a+b]++
		}
	}
	res := 0
	for _, c := range C {
		for _, d := range D {
			res += ab[-c-d]
		}
	}
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
