package main

import (
	"fmt"
	"math"
)

// https://leetcode-cn.com/problems/string-to-integer-atoi/

func myAtoi(s string) int {
	N, i := len(s), 0
	if N == 0 {
		return 0
	}
	for ; i < N && s[i] == ' '; i++ {
	}
	if i >= N || s[i] != '-' && s[i] != '+' && (s[i] < '0' || s[i] > '9') {
		return 0
	}
	res, sign := int64(0), int64(1)
	if s[i] == '-' || s[i] == '+' {
		if s[i] == '-' {
			sign = -1
		}
		i++
	}
	if i >= N || s[i] < '0' || s[i] > '9' {
		return 0
	}
	for ; i < N && s[i] <= '9' && s[i] >= '0'; i++ {
		res = res*10 + int64(s[i]-'0')
		if sign*res > math.MaxInt32 {
			return math.MaxInt32
		}
		if sign*res < math.MinInt32 {
			return math.MinInt32
		}
	}

	return int(res * sign)
}

func main() {
	cases := []string{
		"42",
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(myAtoi(c))
	}
}
