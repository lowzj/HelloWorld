package main

import (
	"fmt"
	"sort"
)

// https://leetcode-cn.com/problems/3sum/

func threeSum0(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)
	for i, prei := 0, -1; i < len(nums)-2 && nums[i] <= 0; i++ {
		if prei >= 0 && nums[prei] == nums[i] {
			continue
		}
		for j, prej := i+1, -1; j < len(nums)-1 && nums[j] <= -nums[i]; j++ {
			if prej > i && nums[prej] == nums[j] {
				continue
			}
			if nums[i]+2*nums[j] > 0 {
				break
			}
			k := biSearch(-nums[i]-nums[j], nums[j+1:])
			if k >= 0 {
				res = append(res, []int{nums[i], nums[j], nums[k+j+1]})
			}
			prej = j
		}
		prei = i
	}
	return res
}

func threeSum(nums []int) [][]int {
	var res [][]int
	if len(nums) < 3 {
		return res
	}
	sort.Ints(nums)

	for i := 0; i < len(nums)-2 && nums[i] <= 0; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		for l, r := i+1, len(nums)-1; l < r; {
			for ; l < r && l > i+1 && nums[l] == nums[l-1]; l++ {
			}
			if l >= r {
				break
			}
			if nums[i]+nums[l]+nums[r] == 0 {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				l++
				r--
			} else if nums[i]+nums[l]+nums[r] < 0 {
				l++
			} else {
				r--
			}
		}
	}
	return res
}

func biSearch(v int, nums []int) int {
	l, h := 0, len(nums)-1
	mid := 0
	for l <= h {
		mid = (l + h) / 2
		if nums[mid] == v {
			return mid
		}
		if nums[mid] > v {
			h = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}

func main() {
	cases := [][]int{
		{-1, 0, 1, 2, -1, -4},
		{-1, 0, 1, 2, -1, -4, -2, -3, 3, 0, 4},
		{0, 0, 0},
		{-2, 0, 0, 2, 2},
	}
	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		fmt.Println("I:", c)
		fmt.Println("O:", threeSum(c))
	}
}
