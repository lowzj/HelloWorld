package main

import (
	"fmt"
	"sort"
)

// https://leetcode-cn.com/problems/reorganize-string/
// 767. 重构字符串 | Reorganize String
//------------------------------------------------------------------------------

//------------------------------------------------------------------------------
// Solution
//
// 复杂度分析:
//   * 时间: O(N+26lg26)
//   * 空间: O(N)
func reorganizeString(S string) string {
	N := len(S)
	f := make([]int, 26)
	for _, c := range S {
		f[c-'a'] = (f[c-'a']/100+1)*100 + int(c-'a')
	}
	sort.Ints(f)
	res, k, start := make([]byte, N), 0, 0
	for i := 25; i >= 0 && f[i] > 0; i-- {
		v, b := f[i]/100, byte(f[i]%100+'a')
		if N%2 == 0 && v > N/2 || N%2 == 1 && v > (N+1)/2 {
			return ""
		}
		for j := v; j > 0; j-- {
			res[k] = b
			k += 2
			if k >= N {
				start++
				k = start
			}
		}
	}
	return string(res)
}

//------------------------------------------------------------------------------
// main

func main() {
	cases := []string{
		"aab",
		"aaab",
		"vvvlo",
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(reorganizeString(c))
	}
}
