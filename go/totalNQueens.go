package main

import (
	"fmt"
)

// https://leetcode-cn.com/problems/n-queens/

func solveNQueens(n int) [][]string {
	const NotQ = -1

	row, col := make([]bool, n), make([]bool, n)
	left, right := make([]bool, 2*n), make([]bool, 2*n)
	matrix := make([]int, n)
	var res [][]string

	isOK := func(i, j int) bool {
		return !row[i] && !col[j] && !left[i+j] && !right[n+i-j]
	}
	mark := func(i, j int) {
		row[i], col[j], left[i+j], right[n+i-j] =
			true, true, true, true
		matrix[i] = j
	}
	remove := func(i, j int) {
		row[i], col[j], left[i+j], right[n+i-j] =
			false, false, false, false
		matrix[i] = NotQ
	}
	addResult := func() {
		result := make([]string, n)
		for i := 0; i < n; i++ {
			str := make([]byte, n)
			for j := 0; j < n; j++ {
				str[j] = '.'
			}
			str[matrix[i]] = 'Q'
			result[i] = string(str)
		}
		res = append(res, result)
	}

	for i := 0; i < n; i++ {
		matrix[i] = NotQ
	}

	var queen func(row int)
	queen = func(row int) {
		if row >= n {
			addResult()
			return
		}
		for col := 0; col < n; col++ {
			if !isOK(row, col) {
				continue
			}
			mark(row, col)
			queen(row + 1)
			remove(row, col)
		}
	}
	queen(0)
	return res
}

func main() {
	cases := []int{
		1, 2, 3, 4, 5, 6, 7, 8,
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(c, ":", solveNQueens(c))
	}
}
