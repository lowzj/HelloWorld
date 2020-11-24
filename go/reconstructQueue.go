package main

import (
	"fmt"
	"sort"
)

// https://leetcode-cn.com/problems/queue-reconstruction-by-height/
// 406. 根据身高重建队列 | Queue Reconstruction by Height
//------------------------------------------------------------------------------

//------------------------------------------------------------------------------
// Solution
//   * 排序: 按 k 升序, 同 k 再按身高 h 降序.
//   * 插入: 依次将排序后的 (h,k) 插入到结果集 res 中.
//     * 从左至右在 res 中查找大于等于 h 的第 k+1 个元素的位置 i, 如果没有则 i 为 len(res).
//     * 将位置 i (包括i) 逐次向后移动一位.
//     * 将 (h,k) 插入到位置 i 处: res[i] = []int{h,k}
// 注意: 要插入到 k+1 的位置上, 因为在 k 和 k+1 之间可能包含比待插入元素小的.
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
		_ = sort.Reverse(sort.IntSlice(m[i]))
	}

	res, pos := make([][]int, N), 0
	add := func(v, cnt int) {
		i := 0
		for k := cnt; i < pos; i++ {
			if res[i][0] >= v {
				k--
				if k < 0 {
					break
				}
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

//------------------------------------------------------------------------------
// 解法同上
func reconstructQueue0(people [][]int) [][]int {
	N := len(people)
	if N < 2 {
		return people
	}
	sort.Sort(peopleSorter(people))

	res, last := make([][]int, N), 0
	for _, p := range people {
		pos := 0
		for cnt := p[1]; pos < last; pos++ {
			if res[pos][0] >= p[0] {
				cnt--
				if cnt < 0 {
					break
				}
			}
		}
		for i := last; i > pos; i-- {
			res[i] = res[i-1]
		}
		res[pos] = p
		last++
	}
	return res
}

var _ sort.Interface = &peopleSorter{}

type peopleSorter [][]int

func (p peopleSorter) Len() int {
	return len(p)
}

func (p peopleSorter) Less(i, j int) bool {
	if p[i][1] == p[j][1] {
		return p[i][0] > p[j][0]
	}
	return p[i][1] < p[j][1]
}

func (p peopleSorter) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	cases := [][][]int{
		// [[5 0] [7 0] [5 2] [6 1] [4 4] [7 1]]
		{
			{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2},
		},
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(reconstructQueue0(c))
	}
}
