package main

import (
	"fmt"
)

// https://leetcode-cn.com/problems/find-all-anagrams-in-a-string/
//------------------------------------------------------------------------------

func findAnagrams(s string, p string) []int {
	N, M := len(s), len(p)

	target := [26]int{}
	pos, neg := 0, 0
	inc := func(b byte) {
		i := int(b - 'a')
		target[i]++
		if target[i] > 0 {
			pos |= 1 << uint(i)
		}
		if target[i] >= 0 {
			neg &^= 1 << uint(i)
		}
	}
	dec := func(b byte) {
		i := int(b - 'a')
		target[i]--
		if target[i] <= 0 {
			pos &^= 1 << uint(i)
		}
		if target[i] < 0 {
			neg |= 1 << uint(i)
		}
	}
	for i := 0; i < M; i++ {
		inc(p[i])
	}

	var res []int
	for i := 0; i < N; i++ {
		dec(s[i])
		if i >= M {
			inc(s[i-M])
		}
		if pos == 0 && neg == 0 {
			res = append(res, i-M+1)
		}
	}
	return res
}

func main() {
	cases := [][2]string{
		{
			"cbaebabacd",
			"abc",
		},
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(findAnagrams(c[0], c[1]))
	}
}
