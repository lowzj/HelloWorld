package main

import "fmt"

// https://leetcode-cn.com/problems/find-all-numbers-disappeared-in-an-array/

func findDisappearedNumbers(nums []int) []int {
	N := len(nums)
	if N == 0 {
		return nil
	}

	swap := func(i, j int) { nums[i], nums[j] = nums[j], nums[i] }
	for i := 0; i < N; {
		if nums[i] == i+1 || nums[i] == nums[nums[i]-1] {
			i++
		} else {
			swap(i, nums[i]-1)
		}
	}
	var res []int
	for i := 0; i < N; i++ {
		if nums[i] != i+1 {
			res = append(res, i+1)
		}
	}
	return res
}

func main() {
	cases := [][]int{
		{4, 3, 2, 7, 8, 2, 3, 1},
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(findDisappearedNumbers(c))
	}
}
