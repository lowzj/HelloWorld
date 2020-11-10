package main

import "fmt"

// https://leetcode-cn.com/problems/divisor-game/

func minCostClimbingStairs(cost []int) int {
	N := len(cost)
	if N == 0 {
		return 0
	}
	if N == 1 {
		return cost[0]
	}
	min := func(x, y int) int {
		if x > y {
			return y
		}
		return x
	}
	cur, pre := cost[1], cost[0]
	for i := 2; i < N; i++ {
		pre, cur = cur, min(cur, pre)+cost[i]
	}
	return min(cur, pre)
}

func main() {
	cases := [][]int{
		{1, 100, 1, 1, 1, 100, 1, 1, 100, 1},
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(minCostClimbingStairs(c))
	}
}
