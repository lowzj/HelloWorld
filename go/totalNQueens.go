package main

import "fmt"

// https://leetcode-cn.com/problems/n-queens-ii/

func totalNQueens(n int) int {
	row, col := make([]bool, n), make([]bool, n)
	left, right := make([]bool, 2*n), make([]bool, 2*n)
	isOK := func(i, j int) bool {
		return !row[i] && !col[j] && !left[i+j] && !right[n+i-j]
	}
	mark := func(i, j int) {
		row[i], col[j], left[i+j], right[n+i-j] =
			true, true, true, true
	}
	remove := func(i, j int) {
		row[i], col[j], left[i+j], right[n+i-j] =
			false, false, false, false
	}

	var queen func(row int) int
	queen = func(row int) int {
		if row >= n {
			return 1
		}
		count := 0
		for col := 0; col < n; col++ {
			if !isOK(row, col) {
				continue
			}
			mark(row, col)
			count += queen(row + 1)
			remove(row, col)
		}
		return count
	}
	return queen(0)
}

func main() {
	cases := []int{
		1, 2, 3, 4, 5, 6, 7, 8,
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(c, ":", totalNQueens(c))
	}
}
