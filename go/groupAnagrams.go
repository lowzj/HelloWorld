package main

import (
	"fmt"
	"strconv"
)

// https://leetcode-cn.com/problems/group-anagrams/

func groupAnagrams(strs []string) [][]string {
	n := len(strs)
	if n < 1 {
		return [][]string{}
	}

	const N = 26
	bm := [N]int{}
	m := make(map[string]int)
	toId := func(s string) string {
		for i := 0; i < N; i++ {
			bm[i] = 0
		}
		for i := 0; i < len(s); i++ {
			bm[s[i]-'a']++
		}
		res := ""
		for i := 0; i < N; i++ {
			if bm[i] > 0 {
				res += strconv.Itoa(bm[i]) + string(byte(i+'a'))
			}
		}
		return res
	}
	res := make([][]string, 0)
	for _, s := range strs {
		id := toId(s)
		idx, ok := m[id]
		if !ok {
			res = append(res, make([]string, 0))
			idx = len(res) - 1
			m[id] = idx
		}
		res[idx] = append(res[idx], s)
	}
	return res
}

// 质数乘积代替 toId
var listId = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101}

func helper(a string) int {
	ret := 1
	for i := 0; i < len(a); i++ {
		ret *= listId[int(a[i]-'a')]
	}
	return ret
}

func main() {
	cases := [][]string{
		{"eat", "tea", "tan", "ate", "nat", "bat"},
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(groupAnagrams(c))
	}
}
