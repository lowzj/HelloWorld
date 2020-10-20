package main

import "fmt"

// https://leetcode-cn.com/problems/permutations-ii/
// 有重复元素的全排列

func permuteUnique(nums []int) [][]int {
	n := len(nums)
	if n == 0 {
		return [][]int{}
	}

	count := make(map[int]int)
	for i := 0; i < n; i++ {
		count[nums[i]]++
	}

	mark := make(map[int]int, n)
	perm := make([]int, n)
	var res [][]int

	add := func() {
		tmp := make([]int, n)
		copy(tmp, perm)
		res = append(res, tmp)
	}
	var gen func(p int)
	gen = func(p int) {
		if p == n {
			add()
			return
		}
		for k, _ := range count {
			if mark[k] < count[k] {
				perm[p] = k
				mark[k]++
				gen(p + 1)
				mark[k]--
			}
		}
	}
	gen(0)
	return res

}

func main() {
	cases := [][]int{
		{1, 1, 3},
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(permuteUnique(c))
	}
}
