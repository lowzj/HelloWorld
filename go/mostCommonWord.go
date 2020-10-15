package main

import (
	"fmt"
	"strings"
)

// https://leetcode-cn.com/problems/most-common-word/

func mostCommonWord(paragraph string, banned []string) string {
	p := strings.ToLower(paragraph)
	bm, pm := make(map[string]bool), make(map[string]int)
	var most string

	for _, v := range banned {
		bm[v] = true
	}
	w := make([]byte, len(p))
	for i, k := 0, 0; i <= len(p); i++ {
		if i == len(p) || p[i] < 'a' || p[i] > 'z' {
			if k == 0 {
				continue
			}
			word := string(w[:k])
			k = 0
			if bm[word] {
				continue
			}
			pm[word]++
			if pm[most] < pm[word] {
				most = word
			}
		} else {
			w[k] = p[i]
			k++
		}
	}
	fmt.Println(bm)
	fmt.Println(pm)
	return most
}

func main() {
	cases := [][]string{
		{
			"Bob hit a ball, the hit BALL flew far after it was hit.",
			"hit",
		},
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(mostCommonWord(c[0], strings.Split(c[1], " ")))
	}
}
