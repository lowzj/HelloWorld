package main

import "fmt"

// https://leetcode-cn.com/problems/spiral-matrix/

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	m, n := len(matrix), len(matrix[0])
	D := [][]int{
		{0, 1}, {1, 0}, {0, -1}, {-1, 0},
	}
	mark := make([][]bool, m)
	for i := 0; i < m; i++ {
		mark[i] = make([]bool, n)
	}
	k := 0
	res := make([]int, m*n)
	walk := func(row, col int, d []int) (int, int) {
		i, j := row+d[0], col+d[1]
		for 0 <= i && i < m && 0 <= j && j < n && !mark[i][j] {
			res[k], mark[i][j] = matrix[i][j], true
			i, j, k = i+d[0], j+d[1], k+1
		}
		return i - d[0], j - d[1]
	}

	row, col := 0, -1
	for k < m*n {
		for i := 0; i < len(D); i++ {
			row, col = walk(row, col, D[i])
		}
	}
	return res
}

func main() {
	cases := [][][]int{
		{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		},
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(spiralOrder(c))
	}
}
