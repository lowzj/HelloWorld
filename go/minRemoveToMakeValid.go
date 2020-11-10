package main

import (
	"fmt"
	"sort"
)

// https://leetcode-cn.com/problems/minimum-remove-to-make-valid-parentheses/

func minRemoveToMakeValid(s string) string {
	n := len(s)
	if n < 1 {
		return ""
	}
	stack, top := make([]int, n), 0
	deleted, k := make([]int, n), 0
	for i := 0; i < n; i++ {
		switch s[i] {
		case '(':
			stack[top] = i
			top++
		case ')':
			if top > 0 {
				top--
			} else {
				deleted[k] = i
				k++
			}
		}
	}
	for ; top > 0; k++ {
		top--
		deleted[k] = stack[top]
	}
	deleted = deleted[:k]
	sort.Ints(deleted)
	res := make([]byte, n-k)
	for i, j, m := 0, 0, 0; i < n; i++ {
		if j < k && i == deleted[j] {
			j++
			continue
		}
		res[m] = s[i]
		m++
	}
	return string(res)
}

func main() {
	cases := []string{
		"lee(t(c)o)de)",
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(minRemoveToMakeValid(c))
	}
}
