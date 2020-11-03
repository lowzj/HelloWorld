package main

// https://leetcode-cn.com/problems/reorder-list/

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	slow := head
	for fast := head; fast != nil && fast.Next != nil; {
		fast = fast.Next.Next
		slow = slow.Next
	}

	halfHead, half := &ListNode{}, slow.Next
	slow.Next = nil
	for next := half; next != nil; {
		tmp := next.Next
		next.Next = halfHead.Next
		halfHead.Next = next
		next = tmp
	}

	for i, j := head, halfHead.Next; i != slow && j != nil; {
		tmpi, tmpj := i.Next, j.Next
		i.Next, j.Next = j, tmpi
		i, j = tmpi, tmpj
	}
}
