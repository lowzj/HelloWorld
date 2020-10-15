package main

import "fmt"

// https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}

	p := head
	for i := 0; i < n; i++ {
		p = p.Next
	}
	if p == nil {
		head = head.Next
		return head
	}
	pre := head
	for ; p.Next != nil; p = p.Next {
		pre = pre.Next
	}
	pre.Next = pre.Next.Next
	return head
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
