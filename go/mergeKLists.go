package main

import (
	"container/heap"
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

var _ heap.Interface = &minHeap{}

type minHeap struct {
	items []*ListNode
}

func (m *minHeap) Push(x interface{}) {
	m.items = append(m.items, x.(*ListNode))
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
	return m.items[i].Val < m.items[j].Val
}

func (m minHeap) Swap(i, j int) {
	m.items[i], m.items[j] = m.items[j], m.items[i]
}

func mergeKLists(lists []*ListNode) *ListNode {
	ll := len(lists)
	prev := &ListNode{}
	if ll == 0 {
		return nil
	}

	m := &minHeap{}
	heap.Init(m)

	for _, node := range lists {
		if node != nil {
			heap.Push(m, node)
		}
	}

	res := prev
	for m.Len() > 0 {
		top := heap.Pop(m).(*ListNode)
		if top.Next != nil {
			heap.Push(m, top.Next)
		}
		prev.Next = &ListNode{top.Val, nil}
		prev = prev.Next
		top = nil
	}
	return res.Next
}

func transfer(data [][]int) {
	lists := make([]*ListNode, len(data))
	count := 0
	for i := 0; i < len(data); i++ {
		head := &ListNode{}
		for j := len(data[i]) - 1; j >= 0; j-- {
			node := &ListNode{data[i][j], head.Next}
			head.Next = node
			count++
		}
		lists[i] = head.Next
	}
	res := mergeKLists(lists)
	m := make([]int, count)
	for i, node := 0, res; node != nil; node, i = node.Next, i+1 {
		m[i] = node.Val
	}
	fmt.Println(m)
}

func main() {
	cases := [][][]int{
		{
			{1, 4, 5},
			{1, 3, 4},
			{2, 6},
		},
		{
			{1}, {0},
		},
	}
	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		transfer(c)
	}
}
