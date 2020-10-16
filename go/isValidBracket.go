package main

import "fmt"

// https://leetcode-cn.com/problems/valid-parentheses/

func isValid(s string) bool {
	n := len(s)
	if n < 2 || n&1 == 1 {
		return false
	}
	stack, top := make([]byte, n), 0
	match := map[byte]byte{')': '(', '}': '{', ']': '['}

	for i := 0; i < len(s); i++ {
		v, ok := match[s[i]]
		if !ok {
			stack[top] = s[i]
			top++
		} else {
			if top <= 0 {
				return false
			}
			top--
			topV := stack[top]
			if v != topV {
				return false
			}
		}
	}
	return top == 0
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
