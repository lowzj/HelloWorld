package main

import (
	"fmt"
	"math"
)

// https://leetcode-cn.com/problems/freedom-trail/

//------------------------------------------------------------------------------

func findRotateSteps(ring string, key string) int {
	N, M := len(ring), len(key)
	min := func(e ...int) int {
		ret := math.MaxInt32
		for i := 0; i < len(e); i++ {
			if ret > e[i] {
				ret = e[i]
			}
		}
		return ret
	}
	dist := func(i, j int) int {
		if i > j {
			return min(i-j, j-i+N)
		}
		return min(j-i, i-j+N)
	}

	dict := make(map[byte][]int)
	for i := 0; i < N; i++ {
		dict[ring[i]] = append(dict[ring[i]], i)
	}

	f := make([][]int, N)
	for i := 0; i < N; i++ {
		f[i] = make([]int, M+1)
	}
	var dfs func(start, target int) int
	dfs = func(start, target int) int {
		if target == M {
			return 0
		}
		var ds []int
		for _, p := range dict[key[target]] {
			if f[p][target+1] > 0 {
				ds = append(ds, dist(start, p)+f[p][target+1])
			} else {
				d := dist(start, p) + dfs(p, target+1)
				ds = append(ds, d)
			}
		}
		f[start][target] = min(ds...)
		return f[start][target]
	}
	ans := dfs(0, 0) + len(key)
	return ans
}

func main() {
	cases := [][]string{
		{"godding", "gd"},
		{
			"caotmcaataijjxi",
			"oatjiioicitatajtijciocjcaaxaaatmctxamacaamjjx",
		},
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(findRotateSteps(c[0], c[1]))
	}
}
