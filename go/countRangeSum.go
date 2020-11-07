package main

import (
	"fmt"
)

// https://leetcode-cn.com/problems/count-of-range-sum/

func countRangeSum(nums []int, lower int, upper int) int {
	N := len(nums)
	if N == 0 {
		return 0
	}

	sum := make([]int, N+1)
	for i := 1; i <= N; i++ {
		sum[i] = sum[i-1] + nums[i-1]
	}
	res := 0
	for i := 0; i < N; i++ {
		for j := 1; j <= N; j++ {
			v := sum[j] - sum[i]
			if v >= lower && v <= upper {
				res++
			}
		}
	}
	return res
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
