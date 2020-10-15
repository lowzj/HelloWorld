package main

import "container/heap"

// // https://leetcode-cn.com/problems/last-stone-weight

var _ heap.Interface = &minHeap{}

type minHeap struct {
	items []int
}

func (m *minHeap) Push(x interface{}) {
	m.items = append(m.items, x.(int))
}

func (m *minHeap) Pop() interface{} {
	l := len(m.items)
	if l <= 0 {
		return nil
	}
	item := m.items[l-1]
	m.items = m.items[0 : l-1]
	return item
}

func (m minHeap) Len() int {
	return len(m.items)
}

func (m minHeap) Less(i, j int) bool {
	return m.items[i] < m.items[j]
}

func (m minHeap) Swap(i, j int) {
	m.items[i], m.items[j] = m.items[j], m.items[i]
}

func lastStoneWeight(stones []int) int {
	n := len(stones)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return stones[1]
	}

	h := &minHeap{items: make([]int, len(stones))}
	heap.Init(h)
	for i := 0; i < n; i++ {
		heap.Push(h, -stones[i])
	}
	for h.Len() > 1 {
		x, y := heap.Pop(h).(int), heap.Pop(h).(int)
		if x != y {
			heap.Push(h, y-x)
		}
	}
	if h.Len() == 0 {
		return 0
	}
	return h.Pop().(int)
}
