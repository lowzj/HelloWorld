package main

import (
	"fmt"
	"sort"
)

// https://leetcode-cn.com/problems/reverse-pairs/
// 493. 翻转对 | Reverse Pairs
//------------------------------------------------------------------------------

//------------------------------------------------------------------------------
// Solution
//
// 离散化 + 树状数组
//
// 这里注意将原数据*2，以方便处理.
// 离散化2倍后的元素后，将原数据范围缩小到50000内，再从右至左遍历原数组，找到比当前元素小的最大元素
// 位置 idx ，再利用树状数组统计 idx 处的计数.
//
// 复杂度分析:
//   * 时间: O(N*lgN)
//   * 空间: O(N)
func reversePairs(nums []int) int {
	N := len(nums)

	// discretize and delete duplicated values
	ordered, hash, n := discretizeNums(nums)
	// Binary Indexed Trees: [1,n]
	NN := n + 1
	tr := make([]int, NN)
	add := func(i int) {
		for ; i < NN; i += i & -i {
			tr[i]++
		}
	}
	sum := func(i int) (sum int) {
		for ; i > 0; i -= i & -i {
			sum += tr[i]
		}
		return sum
	}

	res := 0
	for i := N - 1; i >= 0; i-- {
		l := sort.SearchInts(ordered, nums[i])
		res += sum(l)
		add(hash[nums[i]*2])
	}

	return res
}
func discretizeNums(nums []int) ([]int, map[int]int, int) {
	N := len(nums)
	hash := make(map[int]int)
	// de-dup
	for i := 0; i < N; i++ {
		hash[nums[i]*2] = i
	}
	ordered, n := make([]int, len(hash)), 0
	for k := range hash {
		ordered[n] = k
		n++
	}
	sort.Ints(ordered)
	for i := 0; i < n; i++ {
		// BIT starts with index 1
		hash[ordered[i]] = i + 1
	}
	return ordered, hash, n
}

//------------------------------------------------------------------------------
// main

func main() {
	cases := [][]int{
		{1, 3, 2, 3, 1},
		{2, 4, 3, 5, 1},
		{-6, -3},
		{6, 3},
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(reversePairs(c))
	}
}
