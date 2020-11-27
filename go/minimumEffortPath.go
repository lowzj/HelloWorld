package main

import (
	"fmt"
	"math"
)

// https://leetcode-cn.com/problems/path-with-minimum-effort/
// 1631. 最小体力消耗路径 | Path With Minimum Effort
//------------------------------------------------------------------------------

//------------------------------------------------------------------------------
// Solution
//
// BFS: 广度搜索所有路径
//   * cost[i][j]: 表示到达节点 (i,j) 消耗的体力.
//   * 基于 (i,j), 计算4个方向的相邻节点 (i2,j2) 的消耗体力
//     * 由 (i,j) 到达 (i2,j2) 消耗: c=MAX(cost[i][j], ABS(H[i][j]-H[i2][j2]))
//     * 若 c > cost[i2][j2], 表明有条路径比其之前的更优, 需要将 (i2,j2) 压入队列, 重新计算
//       其相邻节点消耗的体力.
//     * 最后整个 cost 矩阵都是最优的, 返回cost[N-1][M-1].
//
// 复杂度分析:
//   * 时间: > O(N*M). 不好确定, 因为可能重新入队.
//   * 空间: O(N*M)
func minimumEffortPath(H [][]int) int {
	N, M := len(H), len(H[0])
	cost := make([][]int, N)
	for i := 0; i < N; i++ {
		cost[i] = make([]int, M)
		for j := 0; j < M; j++ {
			cost[i][j] = math.MaxInt32
		}
	}
	NN := N * M
	q, p, size := make([][2]int, NN), 0, 0
	push := func(i, j int) {
		q[p][0], q[p][1] = i, j
		p = (p + 1) % NN
		size++
	}
	pop := func() (int, int) {
		id := (p - size + NN) % NN
		size--
		return q[id][0], q[id][1]
	}
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	set := func(i, j, i2, j2 int) bool {
		if i2 < 0 || i2 >= N || j2 < 0 || j2 >= M {
			return false
		}
		cur := max(cost[i][j], abs(H[i][j]-H[i2][j2]))
		if cur >= cost[i2][j2] {
			return false
		}
		cost[i2][j2] = cur
		return true
	}

	ds := [][2]int{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1},
	}
	push(0, 0)
	cost[0][0] = 0
	for size > 0 {
		i, j := pop()
		for _, d := range ds {
			if set(i, j, i+d[0], j+d[1]) {
				push(i+d[0], j+d[1])
			}
		}
	}
	return cost[N-1][M-1]
}

//------------------------------------------------------------------------------
// main

func main() {
	cases := [][][]int{
		// 0
		{
			{1, 2, 1, 1, 1},
			{1, 2, 1, 2, 1},
			{1, 2, 1, 2, 1},
			{1, 2, 1, 2, 1},
			{1, 1, 1, 2, 1},
		},
		// 9
		{
			{1, 10, 6, 7, 9, 10, 4, 9},
		},
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(minimumEffortPath(c))
	}
}
