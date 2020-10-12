package main

import "fmt"

// https://leetcode-cn.com/problems/search-in-rotated-sorted-array/

func search(nums []int, target int) int {
	l, h, mid := 0, len(nums)-1, 0
	for l <= h {
		mid = (l + h) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < nums[l] {
			if nums[mid] > target || target > nums[h] {
				h = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if nums[mid] < target || target < nums[l] {
				l = mid + 1
			} else {
				h = mid - 1
			}
		}
	}
	return -1
}

func main() {
	cases := [][]int{
		{3, 1},
		{7, 8, 9, 10, 1, 3, 4},
	}
	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		fmt.Println("I:", c)
		fmt.Println("O:", search(c, 1))
	}
}
