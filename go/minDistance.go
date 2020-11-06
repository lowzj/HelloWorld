package main

import "fmt"

// https://leetcode-cn.com/problems/edit-distance/

// -----------------------------------------------------------------------------
// 编辑距离
// 设 D(i,j) 为字符串 S[0..i] 和 T[0..j] 的编辑距离
// 1. 考虑 D(0,0):
//   a. S[0] == T[0], 则 D(0,0) = 0, 无需改动
//   b. S[0] != T[0], 则 D(0,0) = 1, 替换S[0]为T[0], 或者替换T[0]为S[0]
// 2. 考虑 D(0,1):
//   a. S[0] == T[1], 则 D(0,1) = 1, 删除T[0]
//   b. S[0] != T[1], 则 D(0,1) = D(0,0)+1
// 3. 考虑 D(1,0), 情况同2
// 对于 2.a 此种情况, 可以理解为在字符串 S[0..-1] 与 T[0..0] 之后同时添加字符 T[1] 得到
// S[0..0] 与 T[0..1], 所以 D(0,1) = D(-1,0). (此处假设 -1 也是合法的)
//
// 基于以上3种基本结论, 在已知 D(i,j) 的条件下, 考虑 D(i,j+1) 与 D(i+1,j)
// 4. 考虑 D(i,j+1)
//   a. 若 S[i] == T[j+1], 则 D(i,j+1) = D(i-1,j), 即在 S[0..i-1] 与 T[0..j] 后
//      添加 T[j+1] 得到 S[0..i] 与 T[0..j+1]
//   b. 若 S[i] != T[j+1], 可以有如下变换得到
//      * 在 S[0..i] 与 T[i..j] 通过 D(i,j) 次变换后, 再删除 T[j+1]
//      * 在 S[0..i-1] 与 T[i..j] 通过 D(i-1,j) 次变换后, 再替换 S[i] 为 T[j+1]
//      * 在 S[0..i-1] 与 T[i..j+1] 通过 D(i-1,j+1) 次变换后, 再删除 S[i]
//      * 综上, D(i,j+1) = min(D(i,j),D(i-1,j),D(i-1,j+1)) + 1
// 5. 考虑 D(i+1,j), 同4
//
// 在实际代码编写中, 为了消除下标越界(-1)的情况:
//  * 字符串下标从 1 开始, S[0..0], T[0..0] 表示空字符串
//  * D(0,0..M) 表示 S[0..0] 与 T[0..0],T[0..1],...,T[0..M] 的编辑距离:0,1,2,...,M
//  * D(0..N,0) 表示 T[0..0] 与 S[0..0],S[0..1],...,S[0..N] 的编辑距离:0,1,2,...,N
// 最后所求的编辑距离即: D(N,M)

func minDistance(s string, t string) int {
	N, M := len(s), len(t)
	if N == 0 {
		return M
	}
	if M == 0 {
		return N
	}
	d := make([][]int, N+1)
	for i := 0; i <= N; i++ {
		d[i] = make([]int, M+1)
	}
	for i := 0; i <= N; i++ {
		d[i][0] = i
	}
	for i := 0; i <= M; i++ {
		d[0][i] = i
	}

	min := func(x, y int) int {
		if x > y {
			return y
		}
		return x
	}
	for i := 1; i <= N; i++ {
		for j := 1; j <= M; j++ {
			if s[i-1] == t[j-1] {
				d[i][j] = d[i-1][j-1]
			} else {
				d[i][j] = min(d[i-1][j-1], min(d[i-1][j], d[i][j-1])) + 1
			}
		}
	}
	return d[N][M]
}

func main() {
	cases := [][]string{
		{"horse", "rose"},
		{"rose", "horse"},
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(minDistance(c[0], c[1]))
	}
}
