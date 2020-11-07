package main

import "fmt"

// https://leetcode-cn.com/problems/word-break/

func wordBreak(s string, wordDict []string) bool {
	N, M := len(s), len(wordDict)
	dict := make(map[string]bool)
	for i := 0; i < M; i++ {
		dict[wordDict[i]] = true
	}

	mark := make([]int, N)
	var search func(start int) bool
	search = func(start int) bool {
		if start >= N {
			return true
		}
		if mark[start] == 1 {
			return true
		}
		if mark[start] == -1 {
			return false
		}
		for i := start + 1; i <= N; i++ {
			if dict[s[start:i]] && search(i) {
				mark[start] = 1
				return true
			}
		}
		mark[start] = -1
		return false
	}
	return search(0)
}

func main() {
	cases := []struct {
		s    string
		dict []string
	}{
		{"leetcode", []string{"leet", "code"}},
		{"applepenapple", []string{"apple", "pen"}},
		{"catsandog", []string{"cats", "dog", "sand", "and", "cat"}},
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(wordBreak(c.s, c.dict))
	}
}
