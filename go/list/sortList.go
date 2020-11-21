package main

import (
	"fmt"
	"math/rand"
	"time"
)

// https://leetcode-cn.com/problems/sort-list/
//------------------------------------------------------------------------------

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	n := 0
	for p := head; p != nil; p = p.Next {
		n++
	}
	newHead := &ListNode{Next: head}

	rand.Seed(time.Now().UnixNano())
	quickSort(newHead, nil, n)
	return newHead.Next
}

func quickSort(l, r *ListNode, n int) {
	if l == nil || l.Next == nil || l.Next == r || n <= 1 {
		return
	}
	p := l
	for pos := rand.Intn(n); p.Next != nil && p.Next != r && pos > 0; {
		pos--
		p = p.Next
	}
	if p != l {
		tmp := p.Next
		p.Next = tmp.Next
		tmp.Next = l.Next
		l.Next = tmp
		p = tmp
	} else {
		p = p.Next
	}

	left, right := 0, 0
	for pre := l.Next; pre.Next != nil && pre.Next != r; {
		node := pre.Next
		if p.Val > node.Val {
			pre.Next = node.Next
			node.Next = l.Next
			l.Next = node
			left++
		} else {
			pre = pre.Next
			right++
		}
	}
	quickSort(l, p, left)
	quickSort(p, r, right)
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
