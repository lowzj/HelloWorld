package main

import (
	"fmt"
)

// https://leetcode-cn.com/problems/minimum-deletions-to-make-string-balanced/
// 1653. 使字符串平衡的最少删除次数 | Minimum Deletions to Make String Balanced
//------------------------------------------------------------------------------

func minimumDeletions(s string) int {
	return minimumDeletions0(s)
}

//------------------------------------------------------------------------------
// Solution 1
//
// 分别计算前缀 a 与后缀 b 的数量, 则前缀 a 与后缀 b 的和的最大值就是删除后的最长字符串.
//   * P[i] 表示 i 位置前缀 a 的数量
//   * S[i] 表示 i 位置后缀 b 的数量
//   * N - P[i] + S[i], 就是要删除的字符数, 只需要求 MAX(P[i]+S[i]) 即可.
// 可以只用一个数组保存前缀 a 数量, 求后缀 b 的时候可以只用一个变量代替, 以节省空间.
//
// 算法复杂度:
//   * 时间: O(N)
//   * 空间: O(N)
func minimumDeletions1(s string) int {
	N := len(s)
	preA := make([]int, N+1)
	for i := 0; i < N; i++ {
		preA[i+1] = preA[i]
		if s[i] == 'a' {
			preA[i+1]++
		}
	}
	res, b := 0, 0
	for i := N - 1; i >= 0; i-- {
		if s[i] == 'b' {
			b++
		}
		if preA[i+1]+b > res {
			res = preA[i+1] + b
		}
	}
	return N - res
}

//------------------------------------------------------------------------------
// Solution 0
//
// 栈: 栈中存储待删除的字符 'b', 当遇到 'a' 时, 可以删除栈顶的一个 'b', 这样就可以保证遍历过的
// 所有字母都是按照 'aa..abb..b' 顺序排列的. 最后出栈的次数就是最少删除次数.
//
// 复杂度分析:
//   * 时间: O(N)
//   * 空间: O(1). 实际上并不需要真正的栈, 只要用一个变量表示栈是否为空即可.
func minimumDeletions0(s string) int {
	pos, cnt := 0, 0
	for _, c := range s {
		if c == 'b' {
			pos++
		} else {
			if pos > 0 {
				pos--
				cnt++
			}
		}
	}
	return cnt
}

//------------------------------------------------------------------------------
// main

func main() {
	cases := []string{
		// 2
		"aababbab",
		// 25
		"ababaaaabbbbbaaababbbbbbaaabbaababbabbbbaabbbbaabbabbabaabbbababaa",
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(minimumDeletions0(c))
		fmt.Println(minimumDeletions1(c))
	}
}
