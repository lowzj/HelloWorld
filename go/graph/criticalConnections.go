package main

import "fmt"

// https://leetcode-cn.com/problems/critical-connections-in-a-network/

func criticalConnections(n int, connections [][]int) [][]int {
	N := len(connections)
	if N == 0 {
		return [][]int{}
	}
	adj := make([][]int, n)
	for _, c := range connections {
		adj[c[0]] = append(adj[c[0]], c[1])
		adj[c[1]] = append(adj[c[1]], c[0])
	}
	st, top := make([]int, n), 0
	push := func(v int) {
		st[top] = v
		top++
	}
	pop := func() int {
		top--
		return st[top]
	}
	mark, flag := make([]int, n), 0
	search := func(s, target int) bool {
		mark[s] = flag
		push(s)
		for top > 0 {
			v := pop()
			if v == target {
				return true
			}
			for _, vv := range adj[v] {
				if v != s && vv == target {
					return true
				}
				if v == s && vv == target {
					continue
				}
				if mark[vv] != flag {
					mark[vv] = flag
					push(vv)
				}
			}
		}
		return false
	}
	var res [][]int
	for _, c := range connections {
		top = 0
		flag++
		if !search(c[0], c[1]) {
			res = append(res, []int{c[0], c[1]})
		}
	}
	return res
}

func main() {
	cases := []struct {
		n    int
		conn [][]int
	}{
		{
			4,
			[][]int{{0, 1}, {1, 2}, {2, 0}, {1, 3}},
		},
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(criticalConnections(c.n, c.conn))
	}
}
