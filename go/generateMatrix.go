package main

import (
	"fmt"
)

// https://leetcode-cn.com/problems/spiral-matrix-ii/
// 59. 螺旋矩阵 | Spiral Matrix II
//------------------------------------------------------------------------------

//------------------------------------------------------------------------------
// Solution
//
// 控制方向
//
// 复杂度分析:
//   * 时间: O(n*n)
func generateMatrix(n int) [][]int {
	direct := [4][2]int{
		{0, 1}, {1, 0}, {0, -1}, {-1, 0},
	}
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}

	next := 1
	for m := n; m > 0; m -= 2 {
		if m == 1 {
			res[n/2][n/2] = n * n
			break
		}
		start, end := (n-m)/2, n-(n-m)/2
		for i, j, d, cnt := start, start, 0, 4*(m-1); cnt > 0; cnt-- {
			res[i][j] = next
			next++

			if vi, vj := i+direct[d][0], j+direct[d][1]; vi >= end || vi < start || vj >= end || vj < start {
				d = (d + 1) % 4
			}
			i += direct[d][0]
			j += direct[d][1]
		}
	}
	return res
}

//------------------------------------------------------------------------------
// Solution
//
// 控制边界
func generateMatrix0(n int) [][]int {
	if n == 1 {
		return [][]int{{1}}
	}

	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}

	l, r, t, b := 0, n-1, 0, n-1
	total := n * n
	for num := 1; num <= total; {
		for i := l; i <= r; i, num = i+1, num+1 {
			res[t][i] = num
		}
		t++
		for i := t; i <= b; i, num = i+1, num+1 {
			res[i][r] = num
		}
		r--
		for i := r; i >= l; i, num = i-1, num+1 {
			res[b][i] = num
		}
		b--
		for i := b; i >= t; i, num = i-1, num+1 {
			res[i][l] = num
		}
		l++
	}
	return res
}

//------------------------------------------------------------------------------
// main

func main() {
	cases := []int{
		3,
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		m := generateMatrix0(c)
		for _, row := range m {
			fmt.Println(row)
		}
	}
}
