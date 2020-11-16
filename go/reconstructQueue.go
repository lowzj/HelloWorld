package main

import (
	"fmt"
	"sort"
)

// https://leetcode-cn.com/problems/queue-reconstruction-by-height/

//------------------------------------------------------------------------------
//
func reconstructQueue(people [][]int) [][]int {
	N := len(people)
	if N < 0 {
		return people
	}

	m := make([][]int, N)
	for i := 0; i < N; i++ {
		m[people[i][1]] = append(m[people[i][1]], people[i][0])
	}
	for i := 0; i < N; i++ {
		sort.Reverse(sort.IntSlice(m[i]))
	}

	res, pos := make([][]int, N), 0
	add := func(v, cnt int) {
		if pos == 0 {
			res[pos] = []int{v, cnt}
		}
		i := 0
		for k := cnt; i < pos && k >= 0; i++ {
			if res[i][0] >= v {
				k--
			}
			if k == -1 {
				i--
			}
		}
		for j := pos; j > i; j-- {
			res[j] = res[j-1]
		}
		res[i] = []int{v, cnt}
		pos++
	}
	for i := 0; i < N; i++ {
		for j := 0; j < len(m[i]); j++ {
			add(m[i][j], i)
		}
	}
	return res
}

func main() {
	cases := [][][]int{
		{
			{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2},
		},
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(reconstructQueue(c))
	}
}
