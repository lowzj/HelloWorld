package main

import "fmt"

// https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/

// -----------------------------------------------------------------------------
// solution 1
// 暴力姐

func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]bool)
	max, count := 0, 0
	for i := 0; i < len(s); i++ {
		count = 0
		for k := range m {
			delete(m, k)
		}
		for j := i; j < i+128 && !m[s[j]]; j++ {
			m[s[j]] = true
			count++
		}
		if max < count {
			max = count
		}
		if max >= 128 {
			break
		}
	}
	return max
}

// -----------------------------------------------------------------------------
// solution 2

// lengthOfLongestSubstring2
//
// 动态规划, 设:
//     pos[s[i]]: 表示字符s[i]在s[0,..,i-1]中的上一次出现的下标, 初始为-1.
//     Sub[i]: 表示s[0,..,i]中包含s[i]的最长不重复子串.
//     F[i]:   len(Sub[i]).
// 那么, 所求整个字符串s最长不重复子串的长度: MAX(F[0],..,F[n-1])
//
// 对于已知i-1时的pos/Sub/F值, 考虑 i 的情况:
//     1. s[i] 不在Sub[i-1]中, 则:
//        * Sub[i]=Sub[i-1]+s[i+1]
//        * F[i]=F[i-1]+1
//     2. s[i] 在Sub[i-1]中, 则:
//        * Sub[i]=Sub[pos[s[i]]+1:]+s[i]
//        * F[i]=i-pos[s[i]]
// 从上面分析得, 情况2所得F[i]一定是小于等于情况1的, 所以不必遍历Sub[i-1]来判断s[i]是否在其中,
// 只需取两者最小值即可:
//    F[i] = MIN(F[i-1]+1,i-pos[s[i]])
// 最长不重复子串: maxLen = MAX(maxLen,F[i])
// 因为不用求子串, 所以不必要为Sub申请额外空间;
// F只需额外申请O(1)空间的变量来循环记录上一次得到的值.
func lengthOfLongestSubstring2(s string) int {
	pos := [256]int{}
	for i := 0; i < 256; i++ {
		pos[i] = -1
	}
	max, pre := 0, 0
	for i := 0; i < len(s); i++ {
		idx := pos[s[i]]
		pre++
		if pre > i-idx {
			pre = i - idx
		}
		if max < pre {
			max = pre
		}
		pos[s[i]] = i
	}
	return max
}

func main() {
	cases := []string{
		"abcabcbb",
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(lengthOfLongestSubstring2(c))
	}
}
