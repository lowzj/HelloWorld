package main

import "fmt"

// https://leetcode-cn.com/problems/verifying-an-alien-dictionary/

func isAlienSorted(words []string, order string) bool {
	n := len(words)
	if n < 2 {
		return true
	}
	m := make(map[byte]int)
	for i := 0; i < len(order); i++ {
		m[order[i]] = i
	}

	less := func(s1, s2 string) bool {
		i, n1, n2 := 0, len(s1), len(s2)
		for ; i < n1 && i < n2; i++ {
			if s1[i] != s2[i] {
				return m[s1[i]] < m[s2[i]]
			}
		}
		if i >= n1 {
			return true
		}
		return false
	}
	for i := 1; i < n; i++ {
		if !less(words[i-1], words[i]) {
			return false
		}
	}
	return true
}

func main() {
	cases := []struct {
		order string
		words []string
	}{
		{
			order: "hlabcdefgijkmnopqrstuvwxyz",
			words: []string{"hello", "leetcode"},
		},
		{
			order: "worldabcefghijkmnpqstuvxyz",
			words: []string{"word", "world", "row"},
		},
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(isAlienSorted(c.words, c.order))
	}
}
