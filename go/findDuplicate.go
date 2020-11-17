package main

import (
	"fmt"
)

// https://leetcode-cn.com/problems/find-the-duplicate-number/

//------------------------------------------------------------------------------
// 类似于二分查找
//   * 数组内数字的范围是 [1,n], 即 l,r 的初始值为 1,n
//   * 求 [1,n] 的中间值 m, 并统计数组内比 m 小的个数, 值为 less; 比 m 大的个数, 记为 gt.
//   * 如果 less > m-1, 说明 [1,m-1] 内的数字有重复, 则 r = m-1, 重复上面的步骤.
//   * 如果 gt > n-m, 说明 [m+1,n] 内的数字有重复, 则 l = m+1, 重复上面步骤.
//   * 如果上面两种情况都不是, 则重复的数字为 m, 直接返回 m.

func findDuplicate(nums []int) int {
	N := len(nums)
	n := N - 1
	for l, r := 1, n; l <= r; {
		m, less, gt := (l+r)/2, 0, 0
		for i := 0; i < N; i++ {
			if nums[i] < m {
				less++
			} else if nums[i] > m {
				gt++
			}
		}
		if less > m-1 {
			r = m - 1
		} else if gt > n-m {
			l = m + 1
		} else {
			return m
		}
	}
	return 1
}

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
