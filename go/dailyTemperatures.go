package main

import (
	"fmt"
)

// https://leetcode-cn.com/problems/daily-temperatures/
// 739. 每日温度 | Daily Temperatures
//------------------------------------------------------------------------------

func dailyTemperatures(T []int) []int {
	return dailyTemperatures0(T)
}

//------------------------------------------------------------------------------
// Solution 1
//
// 需要找到比当天温度更高的未来最近的一天, 即:
//   * 对于 j, 找到 i 使 T[i] > T[j] && i > j, 且 i-j 最小.
// 递减栈: 栈中存储的是待确定元素下标, 对应的温度是递减的.
//
// 遍历数组:
//   * 入栈: 当前温度 T[i] <= T[top] 时入栈 push(i), 栈中元素为待确定.
//   * 出栈: 当 T[i] > T[top], 则表明第一个比栈顶温度大的天出现了, 栈顶可以确定了, 弹出.
// 对于最后未出栈的天, 表明找不到未来的某天比其温度更高, 根据题意等待天数为0.
// 因为结果数组初始化为 0, 所以不用处理, 也不用再加哨兵来降低代码复杂度.
//
// 复杂度:
//   * 时间: O(N)
//   * 空间: O(N)
func dailyTemperatures1(T []int) []int {
	N := len(T)

	stk, p := make([]int, N), 0

	res := make([]int, N)
	for i := 0; i < N; i++ {
		for p > 0 && T[stk[p-1]] < T[i] {
			p--
			res[stk[p]] = i - stk[p]
		}
		stk[p], p = i, p+1
	}
	return res
}

//------------------------------------------------------------------------------
// Solution 0
// 复杂度分析:
//   * 时间: O(N*max(T))
//   * 空间: O(max(T))
func dailyTemperatures0(T []int) []int {
	N := len(T)
	const M = 101
	f := make([]int, M)
	res := make([]int, N)
	for i := N - 1; i >= 0; i-- {
		for t := T[i] + 1; t < M; t++ {
			if f[t] > 0 && (res[i] == 0 || f[t]-i < res[i]) {
				res[i] = f[t] - i
			}
		}
		f[T[i]] = i
	}
	return res
}

func main() {
	cases := [][]int{
		{73, 74, 75, 71, 69, 72, 76, 73},
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(dailyTemperatures(c))
		fmt.Println(dailyTemperatures1(c))
	}
}
