package main

import (
	"fmt"
	"sort"
)

// https://leetcode-cn.com/problems/3sum-closest/
// 16. 最接近的3数之和 | 3Sum Closest
//------------------------------------------------------------------------------

//------------------------------------------------------------------------------
// Solution
//
// 复杂度分析:
//   * 时间: O(lgN*N^2). 双层循环 * 二分查找.
//   * 空间: O(lgN). 快排消耗.
func threeSumClosest(nums []int, target int) int {
	N := len(nums)
	if N < 3 {
		return 0
	}
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	sort.Ints(nums)
	res := nums[0] + nums[1] + nums[2]
	for i := 0; i < N-2; i++ {
		a := nums[i]
		for j := i + 1; j < N-1; j++ {
			b := nums[j]
			for cnt, diff := 0, a+b-target; cnt < 2; cnt++ {
				k := sort.SearchInts(nums[j+1:], diff) + j + 1
				if k == N {
					k--
				}
				if v := a + b + nums[k]; k > j && abs(res-target) > abs(v-target) {
					res = v
				}
				k--
				if v := a + b + nums[k]; k > j && abs(res-target) > abs(v-target) {
					res = v
				}
				diff = -diff
			}
		}
	}
	return res
}

//------------------------------------------------------------------------------
// Solution 1
// 双指针: 固定一个加数, 双指针搜索另外两个.
// 复杂度:
//   * 时间: O(N^2). 排序 O(N*lgN), 固定 O(N) * 内层双指针 O(N).
//   * 空间: O(lgN). 快排消耗.
func threeSumClosest0(nums []int, target int) int {
	N := len(nums)
	if N < 3 {
		return 0
	}
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	sort.Ints(nums)
	res := nums[0] + nums[1] + nums[2]
	for i := 0; i < N-2; i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		for j, k := i+1, N-1; j < k; {
			v := nums[i] + nums[j] + nums[k]
			if v == target {
				return target
			}
			if abs(res-target) > abs(v-target) {
				res = v
			}
			if v > target {
				k--
			} else {
				j++
			}
		}
	}

	return res
}

//------------------------------------------------------------------------------
// main

func main() {
	cases := []struct {
		nums []int
		k    int
	}{
		{
			[]int{-1, 2, 1, -4},
			1,
		},
		{
			[]int{1, 2, 4, 8, 16, 32, 64, 128},
			82,
		},
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(threeSumClosest0(c.nums, c.k))
	}
}
