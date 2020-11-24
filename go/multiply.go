package main

import (
	"fmt"
)

// https://leetcode-cn.com/problems/multiply-strings/
// 43. 字符串相乘 | Multiply Strings
//------------------------------------------------------------------------------

//------------------------------------------------------------------------------
// Solution
//
// 复杂度分析:
//   * 时间: O(N*M)
//   * 空间: O(N+M). 返回值所申请的空间
func multiply(a string, b string) string {
	if a == "0" || b == "0" {
		return "0"
	}

	N, M := len(a), len(b)

	res, maxLen := make([]byte, N+M), 0
	for i := 0; i < N+M; i++ {
		res[i] = '0'
	}
	for i := N - 1; i >= 0; i-- {
		p, carry := N-1-i, 0
		for j := M - 1; j >= 0; j-- {
			v := int(a[i]-'0')*int(b[j]-'0') + int(res[p]-'0') + carry
			res[p] = byte(v%10 + '0')
			carry = v / 10
			p++
		}
		for carry > 0 && p < N+M {
			v := int(res[p]-'0') + carry
			res[p] = byte(v%10 + '0')
			carry = v / 10
			p++
		}
		if maxLen < p {
			maxLen = p
		}
	}
	for i := 0; i < maxLen/2; i++ {
		res[i], res[maxLen-1-i] = res[maxLen-1-i], res[i]
	}
	return string(res[:maxLen])
}

//------------------------------------------------------------------------------
// main

func main() {
	cases := [][]int{
		{},
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(c)
	}
}
