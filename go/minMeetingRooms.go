package main

import (
	"fmt"
	"sort"
)

// https://leetcode-cn.com/problems/meeting-rooms-ii/

//------------------------------------------------------------------------------

func minMeetingRooms(data [][]int) int {
	n := len(data)
	if n < 2 {
		return n
	}

	sort.Sort(intervals(data))
	var seg [][]int
	for i := 0; i < n; i++ {
		d := data[i]
		j := 0
		for ; j < len(seg); j++ {
			if d[0] >= seg[j][1] {
				seg[j][1] = d[1]
				break
			}
		}
		if j == len(seg) {
			seg = append(seg, []int{d[0], d[1]})
		}
	}
	return len(seg)
}

var _ sort.Interface = &intervals{}

type intervals [][]int

func (n intervals) Len() int {
	return len(n)
}

func (n intervals) Less(i, j int) bool {
	if n[i][0] == n[j][0] {
		return n[i][1] < n[j][1]
	}
	return n[i][0] < n[j][0]
}

func (n intervals) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func main() {
	cases := [][][]int{
		{{0, 30}, {5, 10}, {15, 20}},
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(minMeetingRooms(c))
	}
}
