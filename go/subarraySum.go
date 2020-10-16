package main

import (
	"fmt"
)

// https://leetcode-cn.com/problems/subarray-sum-equals-k/

//------------------------------------------------------------------------------
// solution 1

func subarraySum(nums []int, k int) int {
	n := len(nums)
	if n < 1 {
		return 0
	}

	// sum -> index array
	m := make(map[int][]int)
	// sum of first i numbers
	sums := make([]int, n)
	sums[0] = nums[0]
	m[sums[0]] = []int{0}
	for i := 1; i < n; i++ {
		sums[i] = sums[i-1] + nums[i]
		idxes, _ := m[sums[i]]
		idxes = append(idxes, i)
		m[sums[i]] = idxes
	}

	// sum of sub array from i to j:
	// Sum[i,..,j] = sums[j]-sums[i]+nums[i]
	// give i,k, search j: sums[j]=k+sums[i]-nums[i]
	count := 0
	for i := 0; i < n; i++ {
		idxes, _ := m[k+sums[i]-nums[i]]
		for j := 0; j < len(idxes); j++ {
			if i <= idxes[j] {
				count++
			}
		}
	}

	return count
}

//------------------------------------------------------------------------------
// solution 2

func subarraySum2(nums []int, k int) int {
	n := len(nums)
	if n < 1 {
		return 0
	}

	// sum -> count
	m := make(map[int]int)
	// sum of first i numbers
	// sum of sub array from i to j:
	// Sum[i,..,j] = sums[j]-sums[i]+nums[i]
	// give j,k, search i: sums[i]-nums[i]=sums[j]-k=sums[i-1]
	sum, result := 0, 0
	for j := 0; j < n; j++ {
		sum += nums[j]
		m[sum-nums[j]]++
		if count, ok := m[sum-k]; ok {
			result += count
		}
	}
	return result
}

func main() {
	cases := []struct {
		k    int
		nums []int
	}{
		{2, []int{1, 1, 1}},
		{3, []int{1, 2, 3}},
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(subarraySum2(c.nums, c.k))
	}
}
