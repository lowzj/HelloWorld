package main

import "fmt"

// https://leetcode-cn.com/problems/minimum-path-sum/

func minPathSum(grid [][]int) int {
	N := len(grid)
	if N == 0 {
		return 0
	}
	M := len(grid[0])
	if M == 0 {
		return 0
	}

	min := func(x, y int) int {
		if x > y {
			return y
		}
		return x
	}
	minSum := make([][]int, N)
	for i := 0; i < N; i++ {
		minSum[i] = make([]int, M)
	}
	minSum[0][0] = grid[0][0]
	for i := 1; i < N; i++ {
		minSum[i][0] += minSum[i-1][0] + grid[i][0]
	}
	for i := 1; i < M; i++ {
		minSum[0][i] += minSum[0][i-1] + grid[0][i]
	}
	for i := 1; i < N; i++ {
		for j := 1; j < M; j++ {
			minSum[i][j] = min(minSum[i-1][j], minSum[i][j-1]) + grid[i][j]
		}
	}
	return minSum[N-1][M-1]
}

func main() {
	cases := [][][]int{
		{
			{1, 3, 1},
			{1, 5, 1},
			{4, 2, 1},
		},
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(minPathSum(c))
	}
}
