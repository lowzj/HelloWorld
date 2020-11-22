package main

import (
	"fmt"
	"sort"
)

// https://leetcode-cn.com/problems/minimum-number-of-arrows-to-burst-balloons/
//------------------------------------------------------------------------------

// 其实就是求最小相交区间的个数
func findMinArrowShots(points [][]int) int {
	N := len(points)
	if N == 0 {
		return 0
	}
	sort.Sort(sortPoints(points))
	res, start, end := 1, points[0][0], points[0][1]
	for i := 1; i < N; i++ {
		s, e := points[i][0], points[i][1]
		if s > end {
			res++
			start, end = s, e
		} else {
			if start < s {
				start = s
			}
			if end > e {
				end = e
			}
		}
	}
	return res
}

var _ sort.Interface = &sortPoints{}

type sortPoints [][]int

func (s sortPoints) Len() int {
	return len(s)
}

func (s sortPoints) Less(i, j int) bool {
	if s[i][0] == s[j][0] {
		return s[i][1] < s[j][1]
	}
	return s[i][0] < s[j][0]
}

func (s sortPoints) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	cases := [][][]int{
		{},
		{{1, 2}, {2, 3}, {3, 4}, {4, 5}},
		{{1, 2}},
		{{2, 3}, {2, 3}},
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(findMinArrowShots(c))
	}
}
