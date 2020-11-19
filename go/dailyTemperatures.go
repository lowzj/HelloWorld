package main

import (
	"fmt"
)

// https://leetcode-cn.com/problems/daily-temperatures/
//------------------------------------------------------------------------------

func dailyTemperatures(T []int) []int {
	return dailyTemperatures0(T)
}

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
	}
}
