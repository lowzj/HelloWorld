package main

import "fmt"

// https://leetcode-cn.com/problems/add-two-numbers/

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{0, nil}
	for n, n1, n2 := res, l1, l2; n != nil; n = n.Next {
		if n1 != nil {
			n.Val += n1.Val
			n1 = n1.Next
		}
		if n2 != nil {
			n.Val += n2.Val
			n2 = n2.Next
		}
		if n.Val < 10 && n1 == nil && n2 == nil {
			break
		}
		n.Next = &ListNode{}
		if n.Val >= 10 {
			n.Val -= 10
			n.Next.Val = 1
		}
	}
	return res
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
