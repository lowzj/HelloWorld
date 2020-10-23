package main

import (
	"fmt"
)

// https://leetcode-cn.com/problems/word-ladder/

func ladderLength(beginWord string, endWord string, wordList []string) int {
	n := len(wordList)
	if n == 0 {
		return 0
	}

	isOK := func(a, b string) bool {
		diff := 0
		for i := 0; i < len(a) && diff < 2; i++ {
			if a[i] != b[i] {
				diff++
			}
		}
		return diff == 1
	}

	endIdx, mark := -1, make([]bool, n)
	for i := 0; i < n; i++ {
		if wordList[i] == endWord {
			endIdx = i
		}
	}
	if endIdx < 0 {
		return 0
	}

	// dist[i]: 与 endWord 距离为i的单词下标
	dist, pos := make([][]int, n), 0
	dist[0], mark[endIdx] = []int{endIdx}, true
	for ; pos < n; pos++ {
		for i := 0; i < len(dist[pos]); i++ {
			if isOK(beginWord, wordList[dist[pos][i]]) {
				return pos + 2
			}
			for j := 0; j < n; j++ {
				if !mark[j] && isOK(wordList[dist[pos][i]], wordList[j]) {
					mark[j] = true
					dist[pos+1] = append(dist[pos+1], j)
				}
			}
		}
	}

	return 0
}

func main() {
	cases := [][][]string{
		{
			{"lost", "miss"},
			{"most", "mist", "miss", "lost", "fist", "fish"},
		},
		{
			{"leet", "code"},
			{"lest", "leet", "lose", "code", "lode", "robe", "lost"},
		},
	}

	realCase := cases[1:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(ladderLength(c[0][0], c[0][1], c[1]))
	}
}
