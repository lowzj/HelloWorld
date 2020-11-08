package main

import (
	"container/heap"
	"fmt"
)

type MedianFinder struct {
	max *intHeap
	min *intHeap
	n   int
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{
		min: &intHeap{min: true},
		max: &intHeap{min: false},
	}
}

func (mr *MedianFinder) AddNum(num int) {
	mr.n++
	if mr.max.Len() == 0 || num <= mr.max.top() {
		heap.Push(mr.max, num)
	} else {
		heap.Push(mr.min, num)
	}
	for mr.max.Len() < mr.min.Len() {
		v := heap.Pop(mr.min).(int)
		heap.Push(mr.max, v)
	}
	for mr.max.Len() > mr.min.Len()+1 {
		v := heap.Pop(mr.max).(int)
		heap.Push(mr.min, v)
	}
}

func (mr *MedianFinder) FindMedian() float64 {
	if mr.n&1 == 1 {
		return float64(mr.max.top())
	}
	a, b := mr.max.top(), mr.min.top()
	return float64(a+b) / 2
}

var _ heap.Interface = &intHeap{}

type intHeap struct {
	items []int
	min   bool
}

func (h intHeap) Len() int {
	return len(h.items)
}

func (h intHeap) Less(i, j int) bool {
	if h.min {
		return h.items[i] < h.items[j]
	}
	return h.items[i] > h.items[j]
}

func (h intHeap) Swap(i, j int) {
	h.items[i], h.items[j] = h.items[j], h.items[i]
}

func (h *intHeap) Push(x interface{}) {
	h.items = append(h.items, x.(int))
}

func (h *intHeap) Pop() interface{} {
	n := len(h.items)
	if n == 0 {
		return nil
	}
	top := h.items[n-1]
	h.items = h.items[:n-1]
	return top
}

func (h *intHeap) top() int {
	if len(h.items) == 0 {
		return 0
	}
	return h.items[0]
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */

func main() {
	o := Constructor()
	o.AddNum(1)
	o.AddNum(2)
	o.FindMedian()
	o.AddNum(3)
	o.FindMedian()
	cases := [][]int{
		{},
		{},
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(c)
	}
}
