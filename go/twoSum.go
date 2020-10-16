package main

import (
	"fmt"
)

// https://leetcode-cn.com/problems/two-sum/

func twoSum(nums []int, target int) []int {
	n := len(nums)
	if n < 2 {
		return []int{}
	}
	m := make(map[int]int)
	for i, v := range nums {
		m[v] = i
	}
	for i, v := range nums {
		if j, ok := m[target-v]; ok && i != j {
			return []int{i, j}
		}
	}
	return []int{}
}

func main() {
	cases := []struct {
		t    int
		nums []int
	}{
		{9, []int{2, 7, 11, 15}},
		{6, []int{3, 2, 4}},
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(twoSum(c.nums, c.t))
	}
}
