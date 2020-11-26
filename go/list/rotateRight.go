package main

import (
	"fmt"
)

// url:
// title:
//------------------------------------------------------------------------------

//------------------------------------------------------------------------------
// Solution
//
// 复杂度分析:
//   * 时间:
//   * 空间:
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}

	l, tail := 1, head
	for tail = head; tail.Next != nil; tail = tail.Next {
		l++
	}
	if k %= l; k == 0 {
		return head
	}

	pre := head
	for i := l - k - 1; i > 0; i-- {
		pre = pre.Next
	}
	tail.Next = head
	head = pre.Next
	pre.Next = nil
	return head
}

//------------------------------------------------------------------------------
// main

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
