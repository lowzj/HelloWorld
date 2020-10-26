package main

import "fmt"

// https://leetcode-cn.com/problems/powx-n/

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}
	if n < 0 {
		return myPow(1.0/x, -n)
	}
	if n == 1 {
		return x
	}

	i, res := 1, x
	for ; i*2 <= n; i *= 2 {
		res *= res
	}
	return res * myPow(x, n-i)
}

func main() {
	cases := [][]int{
		{2, 3},
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(myPow(float64(c[0]), c[1]))
	}
}
