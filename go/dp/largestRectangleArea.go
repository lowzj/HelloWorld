package main

import (
	"fmt"
)

// https://leetcode-cn.com/problems/largest-rectangle-in-histogram/
// 84. 柱状图中的最大矩形 | Largest Rectangle in Histogram
// 单调栈同类型:
//   1. 接雨水: https://leetcode-cn.com/problems/trapping-rain-water/
//   2. 每日温度: https://leetcode-cn.com/problems/daily-temperatures/
//   3. 下一个更大元素I: https://leetcode-cn.com/problems/next-greater-element-i/
//   4. 去除重复字母: https://leetcode-cn.com/problems/remove-duplicate-letters/
//   5. 股票价格跨度: https://leetcode-cn.com/problems/online-stock-span/
//   6. 移掉K位数字: https://leetcode-cn.com/problems/remove-k-digits/
//   7. 最短无序连续子数组: https://leetcode-cn.com/problems/shortest-unsorted-continuous-subarray/
//   8. 最大矩形: https://leetcode-cn.com/problems/maximal-rectangle/
//------------------------------------------------------------------------------
func largestRectangleArea(heights []int) int {
	return largestRectangleArea1(heights)
}

//------------------------------------------------------------------------------
// Solution 2
//
//       +--+
//       |  +--+
//    +--|  |  |
//    |  |  |  +--+
//  __|  |  |  |  |_
//     0  1  2  3
//
// 以某根柱子的高度围成的矩形面积, 要求其周围的柱子都比它高, 否则, 这个矩形只包含自己.
// 例如:
//   * 以0号柱子高度: 可以跟1,2号围成一个最大矩形, 与3号柱子不行, 因为比3号高.
//   * 以1号柱子高度: 只能自己玩, 最大矩形面积是其高度*1.
// 所以要确定以柱子 H[i] 的高度围成的最大矩形, 就要找到其左右最近的比其小的柱子.
//
// 递增栈: 递增存储待计算的柱子的下标. 因为栈中是递增的, 即对于栈顶柱子, 左边比其小的柱子已确定,
// 只要待入栈的柱子比栈顶低, 则栈顶柱子能围成的最大矩形就可以确定下来了, 此时就可以将其弹出.
//
// 复杂度分析:
//   * 时间: O(N)
//   * 空间: O(N)
func largestRectangleArea2(heights []int) int {
	N := len(heights)
	if N == 0 {
		return 0
	}

	st, p := make([]int, N+2), 0
	push := func(v int) {
		st[p] = v
		p++
	}
	pop := func() int {
		p--
		return st[p]
	}
	top := func() int {
		return st[p-1]
	}

	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	// add virtual elements to head and tail
	get := func(i int) int {
		if i == 0 || i == N+1 {
			return 0
		}
		return heights[i-1]
	}

	res := 0
	for i := 0; i < N+2; i++ {
		for p > 0 && get(top()) > get(i) {
			// pop() must before top()
			res = max(get(pop())*(i-top()-1), res)
		}
		push(i)
	}
	return res
}

//------------------------------------------------------------------------------
// Solution 1
func largestRectangleArea1(heights []int) int {
	N := len(heights)
	if N == 0 {
		return 0
	}
	st, pos := make([][2]int, N), 0
	push := func(i, v int) {
		st[pos] = [2]int{i, v}
		pos++
	}
	pop := func() (int, int) {
		pos--
		return st[pos][0], st[pos][1]
	}
	abs := func(x int) int {
		if x > 0 {
			return x
		}
		return -x
	}
	f := make([]int, N)
	cnt := 0
	cal := func(i int) {
		pre := i
		for pos > 0 {
			cnt++
			if j, h := pop(); h >= heights[i] {
				f[i] += abs(j-pre) * heights[i]
				pre = j
			} else {
				push(j, h)
				break
			}
		}
		push(pre, heights[i])
	}
	for i := 0; i < N; i++ {
		f[i] = heights[i]
		cal(i)
	}

	res := 0
	pos = 0
	for i := N - 1; i >= 0; i-- {
		cal(i)
		if f[i] > res {
			res = f[i]
		}
	}

	return res
}

//------------------------------------------------------------------------------
// Solution 0
// 暴力解:
// 遍历每根柱子 i, 并向左右延伸, 直至其高度低于当前柱子 H[i], 确定左右边界:l,r;
// 则由柱子 i 构成的最大面积为: S[i] = H[i] * (r-l+1);
// 最后取所有 S 的最大值: MAX(S[0],S[1],...,S[n-1]).
//
// 复杂度:
//   * 时间: O(n^2)
//   * 空间: O(1)
func largestRectangleArea0(heights []int) int {
	N := len(heights)

	res := 0
	for i := 0; i < N; i++ {
		area := heights[i]
		for j := i - 1; j >= 0 && heights[j] >= heights[i]; j-- {
			area += heights[i]
		}
		for j := i + 1; j < N && heights[j] >= heights[i]; j++ {
			area += heights[i]
		}
		if area > res {
			res = area
		}
	}
	return res
}

func main() {
	cases := [][]int{
		{2, 1, 5, 6, 2, 3},
		{4, 2, 0, 3, 2, 4, 3, 4},
		{0},
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(largestRectangleArea(c))
		fmt.Println(largestRectangleArea0(c))
		fmt.Println(largestRectangleArea2(c))
	}
}
