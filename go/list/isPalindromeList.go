package main

// https://leetcode-cn.com/problems/palindrome-linked-list/

// -----------------------------------------------------------------------------
// solution 1
// 因为要达成 O(1) 的空间复杂度, 可以先反转一半的链表, 再顺序比较, 最后再恢复链表

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return false
	}
	if head.Next == nil {
		return true
	}

	reverse := func(node *ListNode, count int) (start, half *ListNode) {
		if count < 0 {
			return nil, nil
		}
		if count == 1 {
			return node, node
		}
		start, half = &ListNode{}, node
		for i := 0; i < count; i++ {
			tmp := half
			half = half.Next
			tmp.Next = start.Next
			start.Next = tmp
		}
		return start.Next, half
	}

	n := 0
	for tmp := head; tmp != nil; tmp = tmp.Next {
		n++
	}

	start, half := reverse(head, n/2)

	i, j, ok := start, half, true
	if n&1 == 1 {
		j = j.Next
	}
	for ; i != half && j != nil; i, j = i.Next, j.Next {
		if i.Val != j.Val {
			ok = false
			break
		}
	}
	return ok
}

// -----------------------------------------------------------------------------
// solution 2
// 独特的思路, 来自:
// https://leetcode-cn.com/problems/palindrome-linked-list/solution/ha-xi-bian-li-yi-ci-jiu-xing-by-tcan1993/
// 如果是回文，那么无论正序还是反序, 其多项式乘积是相等的, 以此来判定.
// 极端情况下也可能非回文的正反序乘积也一样, 但是此题该方法能AC!

func isPalindrome2(head *ListNode) bool {
	seed, exp := 3, 1
	hash1, hash2 := 0, 0
	for p := head; p != nil; p = p.Next {
		hash1 = hash1*seed + p.Val
		hash2 = hash2 + p.Val*exp
		exp *= seed
	}
	return hash1 == hash2
}
