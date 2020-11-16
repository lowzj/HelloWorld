package main

import (
	"fmt"
)

// https://leetcode-cn.com/problems/jump-game/

//------------------------------------------------------------------------------
func canJump(nums []int) bool {
	return canJump0(nums)
}

//------------------------------------------------------------------------------
// Solution 0
// dfs 求解

func canJump0(nums []int) bool {
	N := len(nums)
	if N < 2 {
		return true
	}

	mark := make([]int, N)
	var dfs func(start int) bool
	dfs = func(start int) bool {
		if start >= N-1 || mark[start] == 1 {
			return true
		}
		if mark[start] == -1 {
			return false
		}

		for i := nums[start]; i > 0; i-- {
			if dfs(start + i) {
				mark[start] = 1
				return true
			}
		}
		mark[start] = -1
		return false
	}
	return dfs(0)
}

//------------------------------------------------------------------------------
// Solution 1

func canJump1(nums []int) bool {
	N := len(nums)
	if N < 2 {
		return true
	}

	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	f := make([]int, N)
	f[N-1] = 0
	for i := N - 2; i >= 0; i-- {
		f[i] = nums[i]
		if nums[i] != 0 {
			j := 1
			for ; j <= nums[i] && i+j < N; j++ {
				f[i] = max(f[i+j]+j, f[i])
			}
		}
	}
	fmt.Println(f)
	return f[0] >= N-1
}

//------------------------------------------------------------------------------
// Solution 2
//

func canJump2(nums []int) bool {
	N := len(nums)
	if N < 2 {
		return true
	}

	target := N - 1
	for i := N - 2; i >= 0; i-- {
		if i+nums[i] >= target {
			target = i
		}
	}
	return target == 0
}

func main() {
	cases := [][]int{
		{1, 2, 3},
		{3, 2, 1, 0, 4, 10},
		{3, 0, 8, 2, 0, 0, 1},
		{2, 0, 0},
		{1, 1, 2, 2, 0, 1, 1},
		{5, 9, 3, 2, 1, 0, 2, 3, 3, 1, 0, 0},
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(canJump0(c))
		fmt.Println(canJump2(c))
	}
}
