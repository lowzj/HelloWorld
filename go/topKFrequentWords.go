package main

import (
	"container/heap"
	"fmt"
	"sort"
)

// https://leetcode-cn.com/problems/top-k-frequent-words/

func topKFrequent(words []string, k int) []string {
	n := len(words)
	if n < 2 {
		return words
	}

	h := &minHeap{}
	heap.Init(h)

	sort.Strings(words)
	pre := &heapItem{0, 0}
	for i := 0; i < n; i++ {
		if words[i] == words[pre.i] {
			pre.v++
		}
		if words[i] != words[pre.i] || i == n-1 {
			heap.Push(h, pre)
			if i == n-1 && words[i] != words[pre.i] {
				heap.Push(h, &heapItem{1, i})
			}
			for h.Len() > k {
				heap.Pop(h)
			}
			pre = &heapItem{1, i}
		}
	}

	res := make([]string, k)
	for i := k - 1; i >= 0; i-- {
		item := heap.Pop(h).(*heapItem)
		res[i] = words[item.i]
	}
	return res
}

type heapItem struct {
	v, i int
}

var _ heap.Interface = &minHeap{}

type minHeap struct {
	items []*heapItem
}

func (m *minHeap) Push(x interface{}) {
	m.items = append(m.items, x.(*heapItem))
}

func (m *minHeap) Pop() interface{} {
	n := len(m.items)
	item := m.items[n-1]
	m.items = m.items[:n-1]
	return item
}

func (m minHeap) Len() int {
	return len(m.items)
}

func (m minHeap) Less(i, j int) bool {
	ii, ij := m.items[i], m.items[j]
	// decrease by lexicographical order
	if ii.v == ij.v {
		return ii.i > ij.i
	}
	return ii.v < ij.v
}

func (m minHeap) Swap(i, j int) {
	m.items[i], m.items[j] = m.items[j], m.items[i]
}

func main() {
	cases := []struct {
		k     int
		words []string
	}{
		{
			k:     3,
			words: []string{"i", "love", "leetcode", "i", "love", "coding"},
		},
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(topKFrequent(c.words, c.k))
	}
}
