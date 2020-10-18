package main

import "fmt"

// https://leetcode-cn.com/problems/permutations/

func permute(nums []int) [][]int {
	n := len(nums)
	if n == 0 {
		return [][]int{}
	}

	total := 1
	for i := 2; i <= n; i++ {
		total *= i
	}

	mark := make([]bool, n)
	perm := make([]int, n)
	res, cnt := make([][]int, total), 0

	add := func() {
		res[cnt] = make([]int, n)
		copy(res[cnt], perm)
		cnt++
	}
	var gen func(p int)
	gen = func(p int) {
		if p == n {
			add()
			return
		}
		for i := 0; i < n; i++ {
			if !mark[i] {
				perm[p] = nums[i]
				mark[i] = true
				gen(p + 1)
				mark[i] = false
			}
		}
	}
	gen(0)
	return res
}

func main() {
	cases := [][]int{
		{1, 2, 3},
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(permute(c))
	}
}
