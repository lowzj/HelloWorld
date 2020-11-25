package main

import (
	"bytes"
	"fmt"
)

// https://leetcode-cn.com/problems/multiply-strings/
// 43. 字符串相乘 | Multiply Strings
//------------------------------------------------------------------------------

//------------------------------------------------------------------------------
// Solution
//
// bytes.Repeat([]byte, int) 可以用于快速申请初始化 byte 数组.
//
// 复杂度分析:
//   * 时间: O(N*M)
//   * 空间: O(N+M). 返回值所申请的空间
func multiply(a string, b string) string {
	if a == "0" || b == "0" {
		return "0"
	}

	N, M := len(a), len(b)

	res, start := bytes.Repeat([]byte{'0'}, N+M), N+M-1
	for i := N - 1; i >= 0; i-- {
		p, carry := M+i, 0
		for j := M - 1; j >= 0 || (carry > 0 && p >= 0); j-- {
			v := int(res[p]-'0') + carry
			if j >= 0 {
				v += int(a[i]-'0') * int(b[j]-'0')
			}
			res[p] = byte(v%10 + '0')
			carry = v / 10
			p--
		}
		if start > p+1 {
			start = p + 1
		}
	}
	return string(res[start:])
}

//------------------------------------------------------------------------------
// main

func main() {
	cases := [][2]string{
		{"123", "456"},
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(multiply(c[0], c[1]))
	}
}
