package main

import (
	"fmt"
	"sort"
)

func sortByBits(arr []int) []int {
	N := len(arr)
	if N == 0 {
		return []int{}
	}

	countBits := func(x int) int {
		cnt := 0
		for ; x > 0; cnt++ {
			cnt++
			x &= x - 1
		}
		return cnt
	}
	sorter := make([]*sortItem, N)
	for i := 0; i < N; i++ {
		sorter[i] = &sortItem{arr[i], countBits(arr[i])}
	}
	sort.Sort(&bitsSorter{sorter})
	res := make([]int, N)
	for i := 0; i < N; i++ {
		res[i] = sorter[i].v
	}
	return res
}

type sortItem struct {
	v, cnt int
}

func (si *sortItem) String() string {
	return fmt.Sprintf("(%d,%d)", si.v, si.cnt)
}

var _ sort.Interface = &bitsSorter{}

type bitsSorter struct {
	items []*sortItem
}

func (b bitsSorter) Len() int {
	return len(b.items)
}

func (b bitsSorter) Less(i, j int) bool {
	if b.items[i].cnt == b.items[j].cnt {
		return b.items[i].v < b.items[j].v
	}
	return b.items[i].cnt < b.items[j].cnt
}

func (b bitsSorter) Swap(i, j int) {
	b.items[i], b.items[j] = b.items[j], b.items[i]
}

func main() {
	cases := [][]int{
		{10, 100, 1000, 10000},
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(sortByBits(c))
	}
}
