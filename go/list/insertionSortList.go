package main

import (
	"fmt"
)

// https://leetcode-cn.com/problems/insertion-sort-list/
//------------------------------------------------------------------------------

func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	newHead := &ListNode{}
	for p := head; p != nil; {
		pos := newHead
		for ; pos.Next != nil && pos.Next.Val < p.Val; pos = pos.Next {
		}
		tmp := p.Next
		p.Next = pos.Next
		pos.Next = p
		p = tmp
	}

	return newHead.Next
}

func main() {
	cases := [][]int{
		{},
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(c)
	}
}
