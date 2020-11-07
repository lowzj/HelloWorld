package main

import (
	"container/list"
	"fmt"
)

// https://leetcode-cn.com/problems/rotting-oranges/

func orangesRotting(grid [][]int) int {
	N := len(grid)
	if N == 0 {
		return 0
	}
	M := len(grid[0])
	if M == 0 {
		return 0
	}

	q := &queue{list.New()}
	cnt := 0
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			if grid[i][j] == 2 {
				q.push(i, j, 0)
			}
			if grid[i][j] > 0 {
				cnt++
			}
		}
	}

	res := 0
	push := func(i, j, v int) {
		if i < 0 || i >= N || j < 0 || j >= M ||
			grid[i][j] == 2 || grid[i][j] == 0 {
			return
		}
		grid[i][j] = 2
		if v > res {
			res = v
		}
		q.push(i, j, v)
	}
	for q.len() > 0 {
		i, j, v := q.pop()
		cnt--
		push(i-1, j, v+1)
		push(i+1, j, v+1)
		push(i, j-1, v+1)
		push(i, j+1, v+1)
	}
	if cnt == 0 {
		return res
	}
	return -1
}

type queueNode struct {
	i, j, v int
}
type queue struct {
	items *list.List
}

func (q *queue) len() int {
	return q.items.Len()
}
func (q *queue) push(i, j, v int) {
	q.items.PushFront(&queueNode{i, j, v})
}
func (q *queue) pop() (i, j, v int) {
	top := q.items.Remove(q.items.Back()).(*queueNode)
	return top.i, top.j, top.v
}

func main() {
	cases := [][][]int{
		{
			{2, 1, 1},
			{1, 1, 0},
			{0, 1, 1},
		},
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(orangesRotting(c))
	}
}
