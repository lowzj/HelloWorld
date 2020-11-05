package main

import (
	"fmt"
	"sort"
)

//------------------------------------------------------------------------------
// problem 2
func maxWidthOfVerticalArea(points [][]int) int {
	N := len(points)
	if N == 0 {
		return 0
	}

	x := make([]int, N)
	for i := 0; i < N; i++ {
		x[i] = points[i][0]
	}
	sort.Ints(x)

	res := 0
	for i := 1; i < N; i++ {
		if x[i]-x[i-1] > res {
			res = x[i] - x[i-1]
		}
	}
	return res
}

//------------------------------------------------------------------------------
// problem 1

func frequencySort(nums []int) []int {
	N := len(nums)
	if N == 0 {
		return []int{}
	}

	res := make([]int, N)
	m := make(map[int]int)
	for i := 0; i < N; i++ {
		m[nums[i]]++
	}
	var arr [][2]int
	for k, v := range m {
		arr = append(arr, [2]int{v, k})
	}
	sort.Sort(&freqSort{arr})
	k := 0
	for i := 0; i < len(arr); i++ {
		for j := 0; j < arr[i][0]; j++ {
			res[k] = arr[i][1]
			k++
		}
	}
	return res
}

var _ sort.Interface = &freqSort{}

type freqSort struct {
	items [][2]int
}

func (f freqSort) Len() int {
	return len(f.items)
}

func (f freqSort) Less(i, j int) bool {
	if f.items[i][0] == f.items[j][0] {
		return f.items[i][1] > f.items[j][1]
	}
	return f.items[i][0] < f.items[j][0]
}

func (f freqSort) Swap(i, j int) {
	f.items[i], f.items[j] = f.items[j], f.items[i]
}

func main() {
	cases := [][]int{
		{},
		{},
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(c)
	}
}
