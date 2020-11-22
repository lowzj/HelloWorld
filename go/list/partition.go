package main

import (
	"fmt"
)

// https://leetcode-cn.com/problems/partition-list/
//------------------------------------------------------------------------------
// 所有小于 x 的节点都在大于或等于 x 的节点之前, 并保持初始相对位置
func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}
	newHead := &ListNode{}
	less, tail := newHead, newHead
	for p := head; p != nil; {
		pNext := p.Next
		if p.Val < x {
			less.Next, p.Next = p, less.Next
			if tail == less {
				tail = less.Next
			}
			less = less.Next
		} else {
			tail.Next, p.Next = p, tail.Next
			tail = tail.Next
		}
		p = pNext
	}

	return newHead.Next
}

//------------------------------------------------------------------------------
// | <--- 小于 x ---> | <-- 等于 x --> | <--- 大于 x ---> |
// 例如:
// head = 1->4->3->2->5->2, x = 3
// 结果:   1->2->2->3->4->5
func partitionII(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}
	newHead := &ListNode{}
	tail, gt := newHead, newHead
	for p := head; p != nil; {
		pNext := p.Next
		if p.Val > x {
			gt.Next, p.Next = p, gt.Next
			gt = gt.Next
		} else if p.Val <= x {
			tail.Next, p.Next = p, tail.Next
			if p.Val < x {
				tail = tail.Next
			}
			if newHead == gt {
				gt = newHead.Next
			}
		}
		p = pNext
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
