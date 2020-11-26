package main

import (
	"fmt"
	"sort"
)

// https://leetcode-cn.com/problems/maximum-gap/
// 164. 最大间距 | Maximum Gap
//------------------------------------------------------------------------------

func maximumGap(nums []int) int {
	return maximumGap0(nums)
}

//------------------------------------------------------------------------------
// Solution 2
//
// 基数排序: 最好选择 2的次幂做基数, 可以利用位运算来取模. 另外基数数组最好缓存住, 否则每次都申请
// 新的空间, 既耗时又耗空间.
//
// 复杂度分析:
//   * 时间: O(N). 准确的是 O(N+2*logB(max)*N+N): 包括求最大值, 分桶, 拷贝, 求解.
//     所以最终复杂度跟 S0 先排序差不多, 相当于 S0 选的基数是 2. 如果 max 远大于 N,
//     不如直接排序.
//     * 令 B = 2^x, 令循环次数 c = 2*logB(max) = logB(max^2) -> (2^x)^c = max^2
//       即: max = 2^(x*c/2) = (2^c)^(x/2), 而快排的循环次数: c1 = lgN -> N = 2^c1.
//       所以若 c > c1 可以选择快排, 即: max = (2^c)^(x/2) > (2^c1)^x/2 = N^(x/2).
//     * 所以只有当 x > 2, 即基数 B > 4 的时候, 基数排序才会加速, 否则不如直接快排.
//     * 若 max > N^(x/2) 时, 快排依然优于基数排序.
//     * 若基数为 10, 则 x = lg10 = 2.3, 即 max > N^1.15 时, 可以用快排.
//   * 空间: O(N). 最大可能 O(B*N)
func maximumGap2(nums []int) int {
	N := len(nums)
	if N < 2 {
		return 0
	}

	max := nums[0]
	for i := 1; i < N; i++ {
		if max < nums[i] {
			max = nums[i]
		}
	}

	const B, lg = 32, 5
	base := [B][]int{}

	if max/N > B {
		sort.Ints(nums)
		goto SOLVE
	}

	for exp, l := 1, uint(0); exp <= max; exp, l = exp<<lg, l+lg {
		for i := 0; i < N; i++ {
			id := nums[i] >> l & (B - 1)
			base[id] = append(base[id], nums[i])
		}

		for i, start := 0, 0; i < B; i++ {
			if len(base[i]) != 0 {
				copy(nums[start:], base[i])
				start += len(base[i])
				base[i] = base[i][:0]
			}
		}
	}

SOLVE:
	res := 0
	for i := 1; i < N; i++ {
		if nums[i]-nums[i-1] > res {
			res = nums[i] - nums[i-1]
		}
	}
	return res
}

//------------------------------------------------------------------------------
// Solution 1
//
// 桶排序, 桶数为数组长度 N, 桶大小为 MAX(A)/N+1, 这样可以保证有空桶, 那么必然存在两个非空的
// 相邻桶, 他们的间距大于一个桶的大小; 因此就不用再对桶内元素排序, 只用计算后一个桶的最小值与前一个
// 桶的最大值差值即可.
//
// 复杂度分析:
//   * 时间: O(n)
//   * 空间: O(n)
func maximumGap1(nums []int) int {
	N := len(nums)
	if N < 2 {
		return 0
	}
	max, min := nums[0], nums[0]
	for i := 1; i < N; i++ {
		if max < nums[i] {
			max = nums[i]
		} else if min > nums[i] {
			min = nums[i]
		}
	}

	// max,min in bucket
	bkt := make([][]int, N)
	size := (max-min)/N + 1
	for _, v := range nums {
		id := (v - min) / size
		if bkt[id] == nil {
			bkt[id] = []int{v, v}
		} else {
			if v < bkt[id][1] {
				bkt[id][1] = v
			} else if v > bkt[id][0] {
				bkt[id][0] = v
			}
		}
	}

	res := 0
	max = min
	for i := 0; i < N; i++ {
		if bkt[i] != nil {
			if bkt[i][1]-max > res {
				res = bkt[i][1] - max
			}
			max = bkt[i][0]
		}
	}
	return res
}

//------------------------------------------------------------------------------
// Solution 0
//
// 排序后计算相邻元素差值
//
// 复杂度分析:
//   * 时间: O(n*lgn)
//   * 空间: O(1)
func maximumGap0(nums []int) int {
	if len(nums) < 2 {
		return 0
	}

	sort.Ints(nums)
	res := 0
	for i, v := range nums[1:] {
		if diff := v - nums[i]; diff > res {
			res = diff
		}
	}
	return res
}

//------------------------------------------------------------------------------
// main

func main() {
	cases := [][]int{
		{3, 6, 9, 1},
		{3, 6, 9, 4, 7},
		{11, 10, 12},
		{10, 10, 11, 12},
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(maximumGap0(c))
		fmt.Println(maximumGap1(c))
		fmt.Println(maximumGap2(c))
	}
}
