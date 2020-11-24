package main

import (
	"fmt"
)

// https://leetcode-cn.com/problems/maximal-rectangle/
// 85. 最大矩形 | Maximal Rectangle
//------------------------------------------------------------------------------

func maximalRectangle(matrix [][]byte) int {
	return 0
}

//------------------------------------------------------------------------------
// Solution 1
// 复杂度分析:
//   * 时间: O(n^2*m)
func maximalRectangle1(matrix [][]byte) int {
	N := len(matrix)
	if N == 0 {
		return 0
	}
	M := len(matrix[0])
	if M == 0 {
		return 0
	}

	f := make([][]int, N+1)
	for i := 0; i <= N; i++ {
		f[i] = make([]int, M+1)
	}

	cal := func(i, j int) int {
		ret, min := f[i][j], f[i][j]
		for k := i - 1; k > 0 && f[k][j] != 0; k-- {
			if f[k][j] < min {
				min = f[k][j]
			}
			if v := min * (i - k + 1); v > ret {
				ret = v
			}
		}
		return ret
	}

	res := 0
	for i := 1; i <= N; i++ {
		for j := 1; j <= M; j++ {
			if matrix[i-1][j-1] == '1' {
				f[i][j] = f[i][j-1] + 1
				if v := cal(i, j); v > res {
					res = v
				}
			}
		}
	}
	return res
}

//------------------------------------------------------------------------------
// Solution 0
// 暴力解: O(n^2*m^2), 超时了-_-.
func maximalRectangle0(matrix [][]byte) int {
	N := len(matrix)
	if N == 0 {
		return 0
	}
	M := len(matrix[0])
	if M == 0 {
		return 0
	}

	f := make([][]int, N+1)
	for i := 0; i <= N; i++ {
		f[i] = make([]int, M+1)
	}

	cal := func(i, j int) int {
		ret := 0
		for li := 0; li < i; li++ {
			for lj := 0; lj < j; lj++ {
				v := f[i][j] - f[li][j] - f[i][lj] + f[li][lj]
				if v > ret && v == (i-li)*(j-lj) {
					ret = v
				}
			}
		}
		return ret
	}

	res := 0
	for i := 1; i <= N; i++ {
		for j := 1; j <= M; j++ {
			f[i][j] = f[i][j-1] + f[i-1][j] - f[i-1][j-1]
			if matrix[i-1][j-1] == '1' {
				f[i][j]++
				if v := cal(i, j); v > res {
					res = v
				}
			}
		}
	}
	return res
}

func main() {
	cases := [][][]byte{
		{
			{'1', '0', '1', '0', '0'},
			{'1', '0', '1', '1', '1'},
			{'1', '1', '1', '1', '1'},
			{'1', '0', '0', '1', '0'},
		},
	}

	realCase := cases[1:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(maximalRectangle1(c))
	}
}
