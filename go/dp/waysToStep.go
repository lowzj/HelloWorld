package main

import "fmt"

// https://leetcode-cn.com/problems/three-steps-problem-lcci/

func waysToStep(n int) int {
	const mod = 1000000007
	f := []int{1, 2, 4}
	if n <= 3 {
		return f[n-1]
	}
	for i := 4; i <= n; i++ {
		f[2], f[1], f[0] = (f[2]+f[1]+f[0])%mod, f[2], f[1]
	}
	return f[2]
}

func main() {
	cases := []int{
		4, 5,
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(waysToStep(c))
	}
}
