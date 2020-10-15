package main

import (
	"fmt"
	"sort"
	"strings"
)

// https://leetcode-cn.com/problems/reorder-data-in-log-files/

var _ sort.Interface = &logFiles{}

type logFiles []string

func (l logFiles) Len() int {
	return len(l)
}

func (l logFiles) Less(i, j int) bool {
	const C = ' '
	ii := strings.IndexByte(l[i], C)
	ij := strings.IndexByte(l[j], C)
	less := l[i][ii:] < l[j][ij:]
	if !less && l[i][ii:] == l[j][ij:] {
		return l[i][:ii] < l[j][:ij]
	}
	return less
}

func (l logFiles) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func reorderLogFiles(logs []string) []string {
	const C = ' '
	n, start := len(logs), 0
	res := make([]string, n)
	allNum := func(s string) bool {
		for _, c := range s {
			if c != C && (c < '0' || c > '9') {
				return false
			}
		}
		return true
	}
	for i, e := n-1, n-1; i >= 0; i-- {
		idx := strings.IndexByte(logs[i], C)
		if allNum(logs[i][idx+1:]) {
			res[e] = logs[i]
			e--
		} else {
			res[start] = logs[i]
			start++
		}
	}
	sort.Sort(logFiles(res[:start]))
	return res
}

func main() {
	cases := [][]string{
		{"a1 9 2 3 1", "g1 act car", "zo4 4 7", "ab1 off key dog", "a8 act zoo"},
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(reorderLogFiles(c))
	}
}
