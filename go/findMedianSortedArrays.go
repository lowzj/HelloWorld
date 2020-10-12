package main

import "fmt"

// https://leetcode-cn.com/problems/median-of-two-sorted-arrays/

func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	l1, h1, l2, h2 := 0, m, 0, n
	for i, j := 0, 0; l1 <= h1 || l2 <= h2; {
		i, j = (l1+h1)/2, (l2+h2)/2
		if nums1[i] > nums2[j] {
			h1, l2 = i-1, j+1
		} else if nums1[i] < nums2[j] {
			l1, h2 = i+1, j-1
		} else {
			return float64(nums1[i])
		}
	}
	if l1 > h1 {

	}
	return 0.0
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	i, j, k := 0, 0, (m+n)/2
	count := 1
	if (m+n)&1 == 0 {
		count = 2
	}
	res := make([]int, count)
	for i+j <= k {
		if j >= n || i < m && nums1[i] <= nums2[j] {
			res[(i+j)%count] = nums1[i]
			i++
		} else if i >= m || j < n && nums1[i] > nums2[j] {
			res[(i+j)%count] = nums2[j]
			j++
		}
	}
	if count == 1 {
		return float64(res[0])
	}
	return (float64(res[0]) + float64(res[1])) / 2
}

func main() {
	cases := [][][]int{
		{
			{1, 4, 5},
			{1, 3, 4},
		},
		{
			{1}, {0},
		},
		{
			{1, 2}, {3, 4},
		},
	}

	realCase := cases[2:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		fmt.Println(findMedianSortedArrays(c[0], c[1]))
	}
}
