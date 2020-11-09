package main

import (
	"fmt"
)

// https://leetcode-cn.com/problems/k-th-smallest-in-lexicographical-order/

func findKthNumber(n int, k int) int {

	min := func(x, y int) int {
		if x > y {
			return y
		}
		return x
	}
	count := func(pre, n int) (cnt int) {
		for p := pre + 1; pre <= n; p, pre = p*10, pre*10 {
			cnt += min(p, n+1) - pre
		}
		return cnt
	}
	pos, pre := 1, 1
	for pos < k && pre < n {
		cnt := count(pre, n)
		if pos+cnt > k {
			pre *= 10
			pos++
		} else {
			pre++
			pos += cnt
		}
	}
	return pre
}

func main() {
	cases := [][]int{
		{103, 10},
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(findKthNumber(c[0], c[1]))
	}
}
