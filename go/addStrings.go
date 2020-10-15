package main

import "fmt"

// https://leetcode-cn.com/problems/add-strings/

func addStrings(num1 string, num2 string) string {
	m, n := len(num1), len(num2)
	maxl := m
	if n > m {
		maxl = n
	}
	res := make([]byte, maxl+1)
	for i := 0; i <= maxl; i++ {
		if i < m {
			res[maxl-i] += num1[m-1-i] - '0'
		}
		if i < n {
			res[maxl-i] += num2[n-1-i] - '0'
		}
		if res[maxl-i] >= 10 {
			res[maxl-i] -= 10
			res[maxl-i-1] = 1
		}
		res[maxl-i] += '0'
	}
	if res[0] == '0' {
		return string(res[1:])
	}
	return string(res)
}

func main() {
	cases := [][]string{
		{"1", "999"},
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(addStrings(c[0], c[1]))
	}
}
