package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	return searchMatrix2(matrix, target)
}

// https://leetcode-cn.com/problems/search-a-2d-matrix/

func searchMatrix1(matrix [][]int, target int) bool {
	n := len(matrix)
	if n == 0 {
		return false
	}
	m := len(matrix[0])
	if m == 0 {
		return false
	}

	get := func(i int) int { return matrix[i/m][i%m] }
	l, h := 0, m*n-1
	for l <= h {
		mid := (l + h) / 2
		if get(mid) == target {
			return true
		}
		if get(mid) > target {
			h = mid - 1
		} else {
			l = mid + 1
		}
	}
	return false
}

// https://leetcode-cn.com/problems/search-a-2d-matrix-ii/

func searchMatrix2(matrix [][]int, target int) bool {
	n := len(matrix)
	if n == 0 {
		return false
	}
	m := len(matrix[0])
	if m == 0 {
		return false
	}

	var find func(rowStart, rowEnd, colStart, colEnd int) bool
	find = func(rs, re, cs, ce int) bool {
		if rs > re || cs > ce {
			return false
		}
		rowMid, colMid := (rs+re)/2, (cs+ce)/2
		v := matrix[rowMid][colMid]
		if v == target {
			return true
		}
		if v < target {
			//    cs  cm  ce
			//          +---+
			// rs  L L L|? ?|
			//     L L L|? ?| -> find
			// rm  L L v|G G|
			//    +-----+---+
			//    |? ? G G G|
			// re |? ? G G G| -> find
			//    +---------+
			return find(rs, rowMid, colMid+1, ce) ||
				find(rowMid+1, re, cs, ce)
		} else {
			//    cs  cm  ce
			//    +---------+
			// rs |L L L ? ?|
			//    |L L L ? ?| -> find
			//    +---+-----+
			// rm |L L|v G G
			//    |? ?|G G G
			// re |? ?|G G G  -> find
			//    +---+
			return find(rs, rowMid-1, cs, ce) ||
				find(rowMid, re, cs, colMid-1)
		}
	}

	return find(0, n-1, 0, m-1)
}

func main() {
	cases := []struct {
		t int
		m [][]int
	}{
		{
			t: 5,
			m: [][]int{
				{1, 4, 7, 11, 15},
				{2, 5, 8, 12, 19},
				{3, 6, 9, 16, 22},
				{10, 13, 14, 17, 24},
				{18, 21, 23, 26, 30},
			},
		},
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(searchMatrix(c.m, c.t))
	}
}
