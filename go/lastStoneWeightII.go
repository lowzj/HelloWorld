package main

import (
	"fmt"
	"sort"
)

// https://leetcode-cn.com/problems/last-stone-weight-ii/

func lastStoneWeightII(stones []int) int {
	if len(stones) == 0 {
		return 0
	}
	sort.Ints(stones)
	sum, n, min := 0, len(stones), stones[0]
	for i := 0; i < n; i++ {
		sum += stones[i]
	}

	sums := map[int]bool{0: true}
	preSum := []int{0}
	for i := 0; i < n; i++ {
		var sumed []int
		for j := 0; j < len(preSum); j++ {
			tmp := stones[i] + preSum[j]
			if tmp > sum/2 {
				continue
			}
			if _, ok := sums[tmp]; ok {
				continue
			}
			sumed = append(sumed, preSum[j]+stones[i])
			sums[tmp] = true
		}
		preSum = append(preSum, sumed...)
	}
	for i := sum / 2; i >= min; i-- {
		if _, ok := sums[i]; ok {
			return sum - 2*i
		}
	}
	return 0
}

func main() {
	cases := [][]int{
		{2, 7, 4, 1, 8, 1},
		{1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 14, 23, 37, 61, 98},
		{31, 26, 33, 21, 40},
		{2, 1, 2, 2, 2},
		{21, 60, 61, 20, 31},
	}

	realCase := cases[4:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(lastStoneWeightII(c))
	}
}
