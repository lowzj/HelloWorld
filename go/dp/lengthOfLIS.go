package main

import "fmt"

// https://leetcode-cn.com/problems/longest-increasing-subsequence/

//------------------------------------------------------------------------------
// Solution 1 O(n2)
// f[i] 表示在数组S[0..i]包含第 i 个元素的LIS的长度
// 则 f[i+1] = MAX(f[j0],f[j1],..,f[jk]) + 1
// 其中 j0,j1,..,jk 是数组 S[0..i] 中元素小于 S[i+1] 的下标

func lengthOfLIS(nums []int) int {
	N := len(nums)
	if N < 2 {
		return N
	}
	f := make([]int, N)
	ans := 0
	for i := 0; i < N; i++ {
		f[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && f[j]+1 > f[i] {
				f[i] = f[j] + 1
			}
		}
		if f[i] > ans {
			ans = f[i]
		}
	}

	return ans
}

func main() {
	cases := [][]int{
		{10, 9, 2, 5, 3, 7, 101, 18},
		{2, 2},
		{1, 3, 6, 7, 9, 4, 10, 5, 6},
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(lengthOfLIS(c))
	}
}
