package main

import (
	"fmt"
)

// https://leetcode-cn.com/problems/find-all-anagrams-in-a-string/

//------------------------------------------------------------------------------
// 思路其实大体上都相似, 找出字符串 s 中的连续子串，与 p 比较判断是否位字母异位词。
// 利用bit位来快速判断是否异位。
// 两个字符串是否是字母异位词的条件:
//   1. **长度相等**。因为给定了字符串p，所以窗口大小固定，只需要单指针，可以通过右指针 i 确定左
//      指针 i-len(p)。
//   2. **各个字母的出现次数相同**。因为都是小写字母，所以可以用长度26的数组target统计各个字母的
//      出现次数。统计 p 时递增target内元素, 然后以大小为 len(p) 的窗口在 s 上滑动，入窗口时
//      递减*target[s[i]-'a']--*, 并同时递增恢复出窗口元素*target[s[i-len(p)]-'a']++*。
//      如果 target 内所有元素为0，则该窗口内的 s 的子串就是与 p 是字母异位词。
//
// bit位快速判断:
//   - pos: 表示的低26位表示对应target的各元素是否是非负数
//   - neg: 表示的低26位表示对应target的各元素是否是非正数
//   - 在递增递减target元素时维护pos/neg
//   - 异位词判断条件: ***pos == 0 && neg == 0***
func findAnagrams(s string, p string) []int {
	N, M := len(s), len(p)

	target := [26]int{}
	pos, neg := 0, 0
	inc := func(b byte) {
		i := int(b - 'a')
		target[i]++
		if target[i] > 0 {
			pos |= 1 << uint(i)
		}
		if target[i] >= 0 {
			neg &^= 1 << uint(i)
		}
	}
	dec := func(b byte) {
		i := int(b - 'a')
		target[i]--
		if target[i] <= 0 {
			pos &^= 1 << uint(i)
		}
		if target[i] < 0 {
			neg |= 1 << uint(i)
		}
	}
	for i := 0; i < M; i++ {
		inc(p[i])
	}

	var res []int
	for i := 0; i < N; i++ {
		dec(s[i])
		if i >= M {
			inc(s[i-M])
		}
		if pos == 0 && neg == 0 {
			res = append(res, i-M+1)
		}
	}
	return res
}

func main() {
	cases := [][2]string{
		{
			"cbaebabacd",
			"abc",
		},
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(findAnagrams(c[0], c[1]))
	}
}
