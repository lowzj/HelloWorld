package main

import (
	"fmt"
)

// https://leetcode-cn.com/problems/partition-equal-subset-sum/
// 416. 分割等和子集 | Partition Equal Subset Sum
//------------------------------------------------------------------------------

//------------------------------------------------------------------------------
// Solution
//
// [0-1] 背包问题: 求数组 A 的一个子数组 Sub(A), 使其和恰好为原数组和的一半:
//   > SUM(Sub(A)) = SUM(A)/2
// 所以:
//   1. 目标值 target = SUM(A)/2, 其中 SUM(A) 必须为偶数.
//   2. 令 f[k] 表示: 加上数字 A[i] 后的和是否为 k:
//      * f[k] = f[k-A[i]] || f[k]
//   3. 数组 f 的更新顺序: target -> A[i]. 因为每个元素只能使用一次, 这样可以避免重复使用.
//      如果可以重复计算, 则更新顺序可以反过来: A[i] -> target.
//   4. 最后 f[target] 即为所求.
// 复杂度分析:
//   * 时间: O(n*SUM(A)/2)
//   * 空间: O(SUM(A)/2)
func canPartition(nums []int) bool {
	N := len(nums)
	if N < 2 {
		return N == 0
	}

	sum := 0
	for _, c := range nums {
		sum += c
	}
	if sum%2 == 1 {
		return false
	}

	f := make([]bool, sum/2+1)
	f[0] = true
	for _, c := range nums {
		for k := sum / 2; k >= c; k-- {
			f[k] = f[k] || f[k-c]
		}
		if f[sum/2] {
			return true
		}
	}
	return false
}

//------------------------------------------------------------------------------
// main

func main() {
	cases := [][]int{
		// true
		{1, 5, 11, 5},
		{},
		{14, 9, 8, 4, 3, 2},
		// false
		{1},
		{1, 2, 3, 5},
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(canPartition(c))
	}
}
