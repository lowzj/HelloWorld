package main

import (
	"fmt"
	"sort"
)

// https://leetcode-cn.com/problems/merge-intervals/

var _ sort.Interface = &intervals{}

type intervals [][]int

func (n intervals) Len() int {
	return len(n)
}

func (n intervals) Less(i, j int) bool {
	return n[i][0] < n[j][0]
}

func (n intervals) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func merge(data [][]int) [][]int {
	if len(data) == 0 {
		return [][]int{}
	}
	sort.Sort(intervals(data))
	res, k := make([][]int, len(data)), 0
	res[0] = data[0]
	for i := 1; i < len(data); i++ {
		if data[i][0] > res[k][1] {
			k++
			res[k] = data[i]
		} else if res[k][1] < data[i][1] {
			res[k][1] = data[i][1]
		}
	}
	return res[:k+1]
}

func main() {
	cases := [][][]int{
		{
			{1, 3}, {2, 6}, {8, 10}, {15, 18},
		},
		{},
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(merge(c))
	}
}
