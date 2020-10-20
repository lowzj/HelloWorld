package main

import "fmt"

// https://leetcode-cn.com/problems/backspace-string-compare/

func backspaceCompare(S string, T string) bool {
	back := func(s *string, end int) int {
		i := end
		for c := 0; i >= 0; i-- {
			if (*s)[i] == '#' {
				c++
			} else if c > 0 {
				c--
			} else {
				return i
			}
		}
		return i
	}
	i, j := len(S)-1, len(T)-1
	for i >= 0 || j >= 0 {
		i = back(&S, i)
		j = back(&T, j)
		if i < 0 && j < 0 {
			return true
		}
		if i < 0 || j < 0 || S[i] != T[j] {
			return false
		}
		i--
		j--
	}
	return i == j
}

func main() {
	cases := [][]string{
		{"ab#c", "ad#c"},
		{
			"nzp#o#g",
			"b#nzp#o#g",
		},
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(backspaceCompare(c[0], c[1]))
	}
}
