package main

import (
	"fmt"
)

// https://leetcode-cn.com/problems/sort-colors/

//------------------------------------------------------------------------------

func sortColors(nums []int) {
	sortColors1(nums)
}

//------------------------------------------------------------------------------
// Solution 1
// 三游标一趟扫描
//   * i: 表示数字 0 的最右边, 待填充 0; [0,i) 的位置为 0.
//   * j: 表示数字 1 的最右边, 待填充 1; [i,j) 的位置为 1.
//   * k: 表示数字 2 的最左边, 待填充 2; (k,n-1] 的位置为 2.
// 从0开始, 由左至右移动游标 j, 判断 nums[j]:
//   0: swap(i,j), 此时 nums[i] == 0, nums[j] == 1, 所以 i,j 需要同时加 1.
//   1: 向由移动 j, j++.
//   2: 向左移动 k, k--.
// 终止条件: j <= k, 表示一趟扫描完成.
// 算法复杂度
//   * 时间: O(n)
//   * 空间: O(1)

func sortColors1(nums []int) {
	N := len(nums)
	if N <= 1 {
		return
	}
	swap := func(i, j int) {
		nums[i], nums[j] = nums[j], nums[i]
	}
	for i, j, k := 0, 0, N-1; i <= k && j <= k && k >= 0; {
		switch nums[j] {
		case 0:
			swap(i, j)
			i++
			j++
		case 1:
			j++
		case 2:
			swap(j, k)
			k--
		}
	}
}

//------------------------------------------------------------------------------
// Solution 0
// 计数统计再重写

func sortColors0(nums []int) {
	N := len(nums)
	if N <= 1 {
		return
	}

	f := [3]int{}
	for i := 0; i < N; i++ {
		f[nums[i]]++
	}

	for i, k := 0, 0; i < len(f); i++ {
		for j := 0; j < f[i]; j, k = j+1, k+1 {
			nums[k] = i
		}
	}
}

func main() {
	cases := [][]int{
		{2, 0, 2, 1, 1, 0},
		{1, 2},
		{2, 0, 1},
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		sortColors1(c)
		fmt.Println(c)
	}
}
