package main

import (
	"fmt"
	"math"
)

// https://leetcode-cn.com/problems/max-dot-product-of-two-subsequences/

//------------------------------------------------------------------------------
// 设 f[i][j] 为子序列 a[0..i] 与 b[0..j] 的最大点积和

func maxDotProduct(a []int, b []int) int {
	N, M := len(a), len(b)
	if N == 0 || M == 0 {
		return 0
	}

	max := func(e ...int) int {
		ret := math.MinInt32
		for i := 0; i < len(e); i++ {
			if ret < e[i] {
				ret = e[i]
			}
		}
		return ret
	}

	f := make([][]int, N)
	for i := 0; i < N; i++ {
		f[i] = make([]int, M)
	}
	f[0][0] = a[0] * b[0]
	for i := 1; i < N; i++ {
		f[i][0] = max()
	}
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			v := a[i] * b[j]
			if i == 0 && j == 0 {
				f[i][j] = v
			} else if i == 0 {
				f[i][j] = max(f[i][j-1], v)
			} else if j == 0 {
				f[i][j] = max(f[i-1][j], v)
			} else {
				f[i][j] = max(v, f[i-1][j-1]+v, f[i-1][j-1], f[i-1][j], f[i][j-1])
			}
		}
	}
	return f[N-1][M-1]
}

func main() {
	cases := [][][]int{
		{
			{2, 1, -2, 5},
			{3, 0, -6},
		},
		{
			{-1, -1},
			{1, 1},
		},
		{
			{3, -2},
			{2, -6, 7},
		},
		{
			{-3, -8, 3, -10, 1, 3, 9},
			{9, 2, 3, 7, -9, 1, -8, 5, -1, -1},
		},
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(maxDotProduct(c[0], c[1]))
	}
}
