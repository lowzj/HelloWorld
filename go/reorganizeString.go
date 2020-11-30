package main

import (
	"fmt"
)

// https://leetcode-cn.com/problems/reorganize-string/
// 767. 重构字符串 | Reorganize String
//------------------------------------------------------------------------------

//------------------------------------------------------------------------------
// Solution
//
// 复杂度分析:
//   * 时间:
//   * 空间:
func reorganizeString(S string) string {
	N := len(S)
	f := [26]int{}
	for _, c := range S {
		f[c-'a']++
	}
	res := ""
	for k := 0; k < N; {
		for i := 0; i < 26; i++ {
			if f[i] <= 0 {
				continue
			}
			if N%2 == 0 && f[i] > N/2 || N%2 == 1 && f[i] > (N+1)/2 {
				return ""
			}
			k++
			res += string('a' + i)
			f[i]--
		}
	}
	return res
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
