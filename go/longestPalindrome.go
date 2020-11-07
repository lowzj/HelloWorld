package main

import "fmt"

//

func longestPalindrome(s string) string {
	N := len(s)
	if N <= 1 {
		return s
	}

	res, left, right := 1, 0, 1
	for i := 1; i < 2*N-1; i++ {
		l, r := i/2, (i+1)/2
		for l >= 0 && r < N && s[l] == s[r] {
			l--
			r++
		}
		if r-l-1 > res {
			res = r - l - 1
			left = l + 1
			right = r
		}
	}
	return s[left:right]
}

func main() {
	cases := []string{
		"babad", "cbbd",
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(longestPalindrome(c))
	}
}
