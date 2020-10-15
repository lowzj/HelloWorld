package main

import "fmt"

// https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/

func searchRange(nums []int, target int) []int {
	l, h := 0, len(nums)-1
	low, up := 0, 0
	ok := false
	for m := 0; l <= h; {
		m = (l + h) / 2
		if nums[m] == target {
			ok = true
		}
		if nums[m] >= target {
			h = m - 1
		} else {
			l = m + 1
		}
	}
	low = h + 1
	if !ok {
		return []int{-1, -1}
	}
	l, h = 0, len(nums)-1
	for m := 0; l <= h; {
		m = (l + h) / 2
		if nums[m] > target {
			h = m - 1
		} else {
			l = m + 1
		}
	}
	up = l - 1
	return []int{low, up}
}

func main() {
	cases := [][][]int{
		{
			{8}, {5, 7, 7, 8, 8, 10},
		},
		{
			{8}, {5, 7, 7, 8, 8, 10},
		},
	}

	realCase := cases[1:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		fmt.Println(searchRange(c[1], c[0][0]))
	}
}
