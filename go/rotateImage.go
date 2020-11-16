package main

import (
	"fmt"
)

// https://leetcode-cn.com/problems/rotate-image/

//------------------------------------------------------------------------------

func rotate(matrix [][]int) {
	n := len(matrix)
	if n < 1 {
		return
	}
	type nextFunc func(r, c, l, u, s int) (int, int)
	// down, right, up, left
	next := make([]nextFunc, 4)
	next[0] = func(r, c, lower, upper, step int) (int, int) {
		step, c = step-(c-lower), lower
		r += step
		return r, c
	}
	next[1] = func(r, c, lower, upper, step int) (int, int) {
		step, r = step-(upper-r), upper
		c += step
		return r, c
	}
	next[2] = func(r, c, lower, upper, step int) (int, int) {
		step, c = step-(upper-c), upper
		r -= step
		return r, c
	}
	next[3] = func(r, c, lower, upper, step int) (int, int) {
		step, r = step-(r-lower), lower
		c -= step
		return r, c
	}

	step, start := n-1, 0
	for ; step > 0; step -= 2 {
		for i := start; i < n-start-1; i++ {
			tmp := matrix[start][i]
			r, c := start, i
			for j := 0; j < 4; j++ {
				nextR, nextC := next[j](r, c, start, n-start-1, step)
				matrix[r][c] = matrix[nextR][nextC]
				if j == 3 {
					matrix[r][c] = tmp
				}
				r, c = nextR, nextC
			}
		}
		start++
	}
}

func main() {
	cases := [][][]int{
		{
			{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16},
		},
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		rotate(c)
		for _, x := range c {
			fmt.Println(x)
		}
		fmt.Println()
	}
}
