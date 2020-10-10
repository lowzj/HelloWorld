package main

// https://leetcode-cn.com/problems/substring-with-concatenation-of-all-words/

func findSubstring(s string, words []string) []int {
	if len(words) == 0 || len(words[0]) <= 0 {
		return []int{}
	}
	m := make(map[string]int)
	for _, v := range words {
		m[v]++
	}
	var res []int
	wordLen := len(words[0])
	total := len(words)
	length := len(s)
	for i := 0; i <= length-wordLen*total; i++ {
		if find(s, i, wordLen, total, m) {
			res = append(res, i)
		}
	}
	return res
}

func find(s string, start, wordLen, total int,
	words map[string]int) bool {
	end := start + wordLen*total
	if end > len(s) {
		return false
	}

	m := make(map[string]int)
	count := 0
	for i := end; i >= start+wordLen; i -= wordLen {
		sub := s[i-wordLen : i]
		if v, ok := words[sub]; !ok {
			return false
		} else {
			m[sub]++
			if m[sub] > v {
				return false
			}
			count++
		}
	}
	return count == total
}
