package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	res := &ListNode{}
	for p := head; p != nil; {
		tmp := p.Next
		p.Next = res.Next
		res.Next = p
		p = tmp
	}
	return res.Next
}

func main() {
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
