package main

import "fmt"

// https://leetcode-cn.com/problems/valid-palindrome-ii/

func validPalindrome(s string) bool {
	n := len(s)
	if n <= 1 {
		return true
	}

	valid := func(start, end int) bool {
		if start == end {
			return true
		}
		for ; start < end; start, end = start+1, end-1 {
			if s[start] != s[end] {
				return false
			}
		}
		return true
	}

	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		if s[i] == s[j] {
			continue
		}
		if valid(i+1, j) || valid(i, j-1) {
			return true
		}
		return false
	}
	return true
}

func main() {
	cases := []string{
		"abc",
		"tcaac",
		"eeccccbebaeeabebccceea",
		"aguokepatgbnvfqmgmlcupuufxoohdfpgjdmysgvhmvffcnqxjjxqncffvmhvgsymdjgpfdhooxfuupuculmgmqfvnbgtapekouga",
		"aguokepatgbnvfqmgmlucupuufxoohdfpgjdmysgvhmvffcnqxjjxqncffvmhvgsymdjgpfdhooxfuupuclmgmqfvnbgtapekouga",
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(validPalindrome(c))
	}
}
