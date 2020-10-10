package main

import (
	"fmt"
)

// https://leetcode-cn.com/problems/sudoku-solver/
// 10:27 ~
// 3 个[]int 表示该数字所在的: 横\竖\九宫 可填数字
// 每个 int 的低10个bit表示 1~9 是否可用: 1可用,0不可用

// markValue: 二进制 011,1111,1110
const markValue int = 0x03fe

type value struct {
	i, j, v, t int
}

func solveSudoku(board [][]byte) {
	n := len(board)
	r := make([]int, n)
	c := make([]int, n)
	g := make([]int, n)
	s := initSet(n, r, c, g, board)

	hold := func(v *value, t int) {
		// hold number
		clear(&r[v.i], v.t)
		clear(&c[v.j], v.t)
		clear(&g[gi(v.i, v.j)], v.t)
		board[v.i][v.j] = byte('0' + v.t)
		print2(board)
	}
	release := func(v *value, t int) {
		set(&r[v.i], v.t)
		set(&c[v.j], v.t)
		set(&g[gi(v.i, v.j)], v.t)
		board[v.i][v.j] = '.'
	}

	sLen := len(s)
	for pos := 0; 0 <= pos && pos < sLen; {
		v := s[pos]
		if v.v == 0 { // firstly scan, set available numbers
			v.t, v.v = 0, r[v.i]&c[v.j]&g[gi(v.i, v.j)]
		}
		if v.v == 0 { // no available numbers, back to pre element
			v.t = 0
			pos--
			continue
		}
		if v.t > 0 && v.t <= n {
			release(v, v.t)
		}
		i := v.t + 1
		for ; i <= n; i++ {
			if !test(v.v, i) {
				continue
			}

			v.t = i
			hold(v, v.t)
			if pos+1 < sLen {
				pos++
				break
			}
			if isOver(r, c, g) {
				return
			}
		}
		if i > n {
			v.t = 0
			v.v = 0
			pos--
		}
	}
}

func isOver(r ...[]int) bool {
	for i := 0; i < len(r); i++ {
		for j := 0; j < len(r); j++ {
			if r[i][j] != 0 {
				return false
			}
		}
	}
	return true
}

func initSet(n int, row, column, grid []int, board [][]byte) []*value {
	for i := 0; i < n; i++ {
		row[i] = markValue
		column[i] = markValue
		grid[i] = markValue
	}

	var stack []*value
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] != '.' {
				v := int(board[i][j] - '0')
				clear(&row[i], v)
				clear(&column[j], v)
				clear(&grid[gi(i, j)], v)
			} else {
				stack = append(stack, &value{i, j, 0, 0})
			}
		}
	}
	print(row, column, grid, n)
	return stack
}

func print(r, c, g []int, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%d: %010b(%04d) %010b(%04d) %010b(%04d)\n", i,
			r[i], r[i],
			c[i], c[i],
			g[i], g[i])
	}
}
func print2(b [][]byte) {
	for _, v := range b {
		fmt.Println(string(v))
	}
}

func gi(i, j int) int {
	return i/3*3 + j/3
}

func test(v int, i int) bool {
	return (v & (1 << uint(i))) > 0
}

func set(v *int, i int) {
	*v |= 1 << uint(i)
}

func clear(v *int, i int) {
	*v &^= 1 << uint(i)
}

func main() {
	c := [][]byte{
		[]byte("53..7...."),
		[]byte("6..195..."),
		[]byte(".98....6."),
		[]byte("8...6...3"),
		[]byte("4..8.3..1"),
		[]byte("7...2...6"),
		[]byte(".6....28."),
		[]byte("...419..5"),
		[]byte("....8..79"),
	}
	solveSudoku(c)
	for _, v := range c {
		fmt.Println(string(v))
	}
}
