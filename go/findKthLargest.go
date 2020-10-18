package main

import (
	"fmt"
	"math/rand"
	"time"
)

func findKthLargest(nums []int, k int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	rand.Seed(time.Now().UnixNano())
	swap := func(i, j int) { nums[i], nums[j] = nums[j], nums[i] }
	var find func(l, r int) int
	find = func(l, r int) int {
		p := rand.Intn(r-l) + l
		swap(p, l)
		p, x := l, nums[l]
		for i, j := l+1, r-1; i <= j; {
			if nums[i] >= x {
				swap(i, j)
				j--
			} else {
				nums[p] = nums[i]
				p = i
				i++
			}
		}
		nums[p] = x
		if p == n-k {
			return nums[p]
		}
		if p < n-k {
			return find(p+1, r)
		}
		return find(l, p)
	}
	return find(0, n)
}

func main() {
	cases := [][][]int{
		{
			{3, 2, 1, 5, 6, 4},
			{2},
		},
		{
			{3, 2, 3, 1, 2, 4, 5, 5, 6},
			{4},
		},
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(findKthLargest(c[0], c[1][0]))
	}
}
