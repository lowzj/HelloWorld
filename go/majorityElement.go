package main

import (
	"fmt"
)

// https://leetcode-cn.com/problems/majority-element
//------------------------------------------------------------------------------

func majorityElement(nums []int) int {
	return majorityElement1(nums)
}

//------------------------------------------------------------------------------
// Solution 1
// 摩尔投票法
// 因为多数元素的出现次数大于 N/2, 而其他所有元素的出现次数总和小于一半
// 每次出现多数元素时, 计数cnt+1, 否则cnt-1; 当cnt==0时, 表示可能的多数元素.
// 复杂度:
//   * 时间: O(n)
//   * 空间: O(1)
func majorityElement1(nums []int) int {
	res, cnt := nums[0], 0
	for _, n := range nums {
		if cnt == 0 {
			res = n
		}
		if n == res {
			cnt++
		} else {
			cnt--
		}
	}
	return res
}

//------------------------------------------------------------------------------
// Solution 0
// 排序后计数
// 复杂度分析:
//   * 时间: O(n*lgn)
//   * 空间: O(1)
func majorityElement0(nums []int) int {
	N := len(nums)
	for i, cnt := 1, 1; i < N; i++ {
		if nums[i] == nums[i-1] {
			cnt++
		} else {
			cnt = 1
		}
		if cnt > N/2 {
			return nums[i]
		}
	}
	return nums[N-1]
}

func main() {
	cases := [][]int{
		{2, 2, 1, 1, 1, 2, 2},
		{3, 1, 3},
		{7, 5, 7, 5, 7},
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(majorityElement(c))
	}
}
