package main

import (
	"container/heap"
	"fmt"
)

func furthestBuilding(heights []int, bricks int, ladders int) int {
	N := len(heights)

	h := &maxHeap{nil, ladders}
	totalBricks, total := 0, 0
	for i := 1; i < N; i++ {
		diff := heights[i] - heights[i-1]
		if diff > 0 {
			if totalBricks += h.push(diff); totalBricks > bricks {
				break
			}
		}
		total++
	}
	return total
}

var _ heap.Interface = &maxHeap{}

type maxHeap struct {
	bricks   []int
	capacity int
}

func (m maxHeap) Len() int {
	return len(m.bricks)
}

func (m maxHeap) Less(i, j int) bool {
	return m.bricks[i] > m.bricks[j]
}

func (m maxHeap) Swap(i, j int) {
	m.bricks[i], m.bricks[j] = m.bricks[j], m.bricks[i]
}

func (m *maxHeap) Push(x interface{}) {
	m.bricks = append(m.bricks, x.(int))
}

func (m *maxHeap) Pop() interface{} {
	n := len(m.bricks)
	if n == 0 {
		return nil
	}
	top := m.bricks[n-1]
	m.bricks = m.bricks[:n-1]
	return top
}

func (m *maxHeap) push(x int) int {
	heap.Push(m, x)
	if m.capacity < len(m.bricks) {
		return m.Pop().(int)
	}
	return 0
}

func main() {
	cases := []struct {
		heights []int
		bricks  int
		ladders int
	}{
		{[]int{4, 2, 7, 6, 9, 14, 12}, 5, 1},
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(furthestBuilding(c.heights, c.bricks, c.ladders))
	}
}
