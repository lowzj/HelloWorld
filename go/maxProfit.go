package main

import "fmt"

// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/

func maxProfit(prices []int) int {
	N := len(prices)
	if N < 2 {
		return 0
	}

	sum := 0
	for i := 1; i < N; i++ {
		if v := prices[i] - prices[i-1]; v > 0 {
			sum += v
		}
	}
	return sum
}

func main() {
	cases := [][]int{
		{},
		{},
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(c)
	}
}
