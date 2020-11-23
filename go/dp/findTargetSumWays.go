package main

import (
	"fmt"
)

// https://leetcode-cn.com/problems/target-sum/
//------------------------------------------------------------------------------

//------------------------------------------------------------------------------
// Solution
// 问题:
//   给定一个非负整数数组A: a1,a2,...,an, 和一个目标数: S.
//   现在你有两个符号+和-. 对于数组中的任意一个整数，你都可以从+或-中选择一个符号添加在前面.
//   返回可以使最终数组和为目标数 S 的所有添加符号的方法数.
// 分析:
// 因为只需要添加 + 或 -, 所以实际上就是将数组 A 分为2个子数组 A1,A2, 使其和之差的绝对值等于S:
//   S = ABS(SUM(A1)-SUM(A2))
// 又 SUM(A2) = SUM(A)-SUM(A1), 所以:
//   S = ABS(SUM(A)-2*SUM(A1))
// 并且其中一个子数组是可以为空的, 即 SUM(A1) 可以为0.
// 所以原问题就转化为: 找到 A 中的所有子数组 A1, 使等式 S==ABS(SUM(A)-2*SUM(A1)) 成立.
//   1. 计算 SUM(A). 遍历一遍数组 A 即可.
//   2. 统计所有子数组 A1 可能出现的和的组合数.
//      * 令 f[k] 表示子数组的和为 k 的组合数, 初始 f[0] = 1.
//      * 遍历原数组 A, 更新所有包含 A[i] 的和的组合数: f[k] += f[k-A[i]].
//        k 从大到小更新, 防止重复计算.
//      * k 的范围: [0,SUM(A)]
//   3. 遍历 f, 统计所有使等式成立的组合数. 为防止重复计算, 遍历范围: [0,SUM(A)/2].
//      * 当 S == ABS(SUM(A)-2*f[i]) 为真, 则 res += f[i]
// 其中 SUM(A) 的范围: [0,1000].
//
// 复杂度分析:
//   * 时间: O(n*SUM(A))
//   * 空间: O(SUM(A))
func findTargetSumWays(nums []int, S int) int {
	N := len(nums)
	sum := 0
	for _, c := range nums {
		sum += c
	}

	f := make([]int, sum+1)
	f[0] = 1
	for i := 0; i < N; i++ {
		for k := sum - nums[i]; k >= nums[i]; k-- {
			f[k] += f[k-nums[i]]
		}
	}

	res := 0
	for i := 0; i <= sum/2; i++ {
		if v := sum - 2*i; v == S || -v == S {
			res += f[i]
		}
	}
	return res
}

func main() {
	cases := []struct {
		nums []int
		s    int
	}{
		{
			[]int{1, 1, 1, 1, 1},
			3,
		},
		{
			[]int{1, 2, 1},
			0,
		},
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(findTargetSumWays(c.nums, c.s))
	}
}
