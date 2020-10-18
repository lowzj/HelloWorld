package main

import "fmt"

// https://leetcode-cn.com/problems/generate-parentheses/

func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}

	str := make([]byte, 2*n)
	var res []string
	var gen func(left, right int)
	gen = func(l, r int) {
		if l == n && r == n {
			res = append(res, string(str))
			return
		}

		if l < n {
			str[l+r] = '('
			gen(l+1, r)
		}
		if l > r {
			str[l+r] = ')'
			gen(l, r+1)
		}
	}

	gen(0, 0)
	return res
}

func main() {
	cases := []int{
		3,
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(generateParenthesis(c))
	}
}
