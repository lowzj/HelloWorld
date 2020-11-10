package main

import "fmt"

// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/

//------------------------------------------------------------------------------
// 可以转为最大连续子序列和

func maxProfit(prices []int) int {
	N := len(prices)
	if N < 2 {
		return 0
	}
	diff := make([]int, N-1)
	for i := 1; i < N; i++ {
		diff[i-1] = prices[i] - prices[i-1]
	}
	ans, sum := 0, 0
	for i := 0; i < N-1; i++ {
		sum += diff[i]
		if sum > ans {
			ans = sum
		}
		if sum < 0 {
			sum = 0
		}
	}
	return ans
}

// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/

func maxProfit2(prices []int) int {
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
