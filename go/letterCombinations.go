package main

import "fmt"

// https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	nums := [10]string{"", "",
		"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

	n := len(digits)
	s := &letterStack{nums: make([]int, 0, n)}
	str, pos := make([]byte, n), 0
	res := make([]string, 0, n)
	s.push(int(digits[0]-'0'), 0)
	for s.len() > 0 {
		c, i := s.pop()
		if i == len(nums[c]) {
			pos--
			continue
		}
		str[pos] = nums[c][i]
		pos++
		s.push(c, i+1)
		if pos < n {
			s.push(int(digits[pos]-'0'), 0)
		} else if pos == n {
			res = append(res, string(str))
			pos--
		}
	}
	return res
}

type letterStack struct {
	nums []int
}

func (ls *letterStack) len() int {
	return len(ls.nums)
}

func (ls *letterStack) push(c, i int) {
	ls.nums = append(ls.nums, c*10+i)
}

func (ls *letterStack) pop() (c, i int) {
	n := len(ls.nums)
	if n == 0 {
		return 0, 0
	}
	v := ls.nums[n-1]
	ls.nums = ls.nums[:n-1]
	return v / 10, v % 10
}

func main() {
	cases := []string{
		"23",
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(letterCombinations(c))
	}
}
