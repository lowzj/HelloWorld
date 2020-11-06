package main

import (
	"fmt"
)

// https://leetcode-cn.com/problems/coin-change/

func coinChange(coins []int, amount int) int {
	N := len(coins)
	if N == 0 {
		return -1
	}
	if amount == 0 {
		return 0
	}

	A := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		A[i] = -1
	}
	for i := 0; i < N; i++ {
		if coins[i] <= amount {
			A[coins[i]] = 1
		}
	}
	for i := 1; i <= amount; i++ {
		for j := 0; j < N; j++ {
			if i > coins[j] && A[i-coins[j]] >= 0 &&
				(A[i] == -1 || A[i] > 1+A[i-coins[j]]) {
				A[i] = 1 + A[i-coins[j]]
			}
		}
	}
	return A[amount]
}

func main() {
	cases := []struct {
		coins  []int
		amount int
	}{
		{[]int{1, 2, 5}, 11},
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(coinChange(c.coins, c.amount))
	}
}
