package main

import (
	"fmt"
)

// https://leetcode-cn.com/problems/maximal-rectangle/
// 85. 最大矩形 | Maximal Rectangle
//------------------------------------------------------------------------------

func maximalRectangle(matrix [][]byte) int {
	return maximalRectangle2(matrix)
}

//------------------------------------------------------------------------------
// Solution 2
//
// 将每列上连续1的点看作一个柱子, 柱子的高度就是连续1的数量, 若最底层的点不为1, 则该列柱子高度为0.
// 那么该问题就转为求: 柱状图中的最大矩形.
// 各个行与第0行构成的柱状图, 可以用一个高度数组表示: H.
//     | 0 1 2 3 4
//   --+-----------
//   0 | 1 0 1 0 0
//   1 | 1 0 1 1 1
//   2 | 1 1 1 1 1
//   3 | 1 0 0 1 0
// 例如上面例子中各个行构成的柱状图:
//   * H[0]: [1,0,1,0,0]
//   * H[1]: [2,0,2,1,1]
//   * H[2]: [3,1,3,2,2]
//   * H[3]: [4,0,0,3,0]
// 针对每行的构成的柱状图, 可以直接复用 [柱状图中的最大矩形] 的算法, 利用递增栈 O(m) 时间得出最大
// 矩形面积, 最后选最大值返回.
// 空间优化: 因为是逐行扫描逐行计算的, 上层得出的高度数组 H[i-1] 可以复用到下一行 H[i], 所以高度
// 数组只需要申请一维数组(列数), 占用 O(M) 空间.
//
// 复杂度分析:
//   * 时间: O(N*M).
//   * 空间: O(M). 高度及栈各申请 M+2 空间. 2 为首尾哨兵.
func maximalRectangle2(matrix [][]byte) int {
	N := len(matrix)
	if N == 0 {
		return 0
	}
	M := len(matrix[0])
	if M == 0 {
		return 0
	}

	H := make([]int, M+2)
	st, p := make([]int, M+2), 0
	push := func(i int) {
		st[p] = i
		p++
	}
	pop := func() int {
		p--
		return st[p]
	}
	top := func() int {
		return st[p-1]
	}
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	res := 0
	for _, m := range matrix {
		for j := 0; j < M; j++ {
			if m[j] == '1' {
				H[j+1]++
			} else {
				H[j+1] = 0
			}
		}
		for j := 0; j < M+2; j++ {
			for p > 0 && H[top()] > H[j] {
				res = max(H[pop()]*(j-top()-1), res)
			}
			push(j)
		}
		// reset stack
		p = 0
	}

	return res
}

//------------------------------------------------------------------------------
// Solution 1
//
// 优化 Solution 0. S0 没有记录连续1这个信息, 所以在固定一个点计算连续1矩形面积的时候, 又需要
// O(n*m) 的时间, 最终导致 O(n^2*m^2) 的时间复杂度.
// 所以该方法就是保存每行连续1的信息, 优化固定点后的计算效率, 降为 O(n).
//
// 图解:
//     | 0 1 2 3 4
//   --+-----------
//   0 | 1 0 1 0 0
//   1 | 1 0 1 1 1
//   2 | 1 1 1 1 1
//   3 | 1 0 0 1 0
// 逐行计算连续1的数目, 得到如下矩阵 f:
//     | 0 1 2 3 4
//   --+-----------                     列数      行数
//   0 | 1 0 1 0 0 <- f[0][4] <- s[0] = MIN(0,5)*(2-0+1) = 0
//   1 | 1 0 1 2 3 <- f[1][4] <- s[1] = MIN(3,5)*(2-1+1) = 6
//   2 | 1 2 3 4 5 <- f[2][4] <- s[2] = MIN(5,5)*(2-2+1) = 5
//   3 | 1 0 0 1 0
// 对于矩阵 f 中的每个点 f[i][j], 计算以 (i,j) 为右下角到左上方所能构成的最大面积S[i][j]:
//   * 固定列向上遍历 f, 即 j 不变
//   * 设 s[k] 表示第 k 行的 j 列左方与第 i 行的 j 列左方所能构成的连续1的矩行面积(包括j列):
//     s[k] = MIN(f[k][j],f[i][j])*(i-k+1), k 属于 [0,i].
//   * S[i][j] = MAX(s[0],s[1],...,s[i])
// 最大矩形面积就是: MAX(S).
// 实际编写代码中, 数组 s 以及矩阵 S 不用申请空间, 因为只需要记录最值, 用 O(1) 空间即可.
//
// 复杂度分析:
//   * 时间: O(n^2*m). 每行连续1的数目: O(n*m); 每个点 S[i][j]: O(n), 所有点 O(n*m*n).
//   * 空间: O(n*m). 矩阵 f 占用空间 n*m.
func maximalRectangle1(matrix [][]byte) int {
	N := len(matrix)
	if N == 0 {
		return 0
	}
	M := len(matrix[0])
	if M == 0 {
		return 0
	}

	f := make([][]int, N+1)
	for i := 0; i <= N; i++ {
		f[i] = make([]int, M+1)
	}

	cal := func(i, j int) int {
		ret, min := f[i][j], f[i][j]
		for k := i - 1; k > 0 && f[k][j] != 0; k-- {
			if f[k][j] < min {
				min = f[k][j]
			}
			if v := min * (i - k + 1); v > ret {
				ret = v
			}
		}
		return ret
	}

	res := 0
	for i := 1; i <= N; i++ {
		for j := 1; j <= M; j++ {
			if matrix[i-1][j-1] == '1' {
				f[i][j] = f[i][j-1] + 1
				if v := cal(i, j); v > res {
					res = v
				}
			}
		}
	}
	return res
}

//------------------------------------------------------------------------------
// Solution 0
// 暴力解: O(n^2*m^2), 超时了-_-.
func maximalRectangle0(matrix [][]byte) int {
	N := len(matrix)
	if N == 0 {
		return 0
	}
	M := len(matrix[0])
	if M == 0 {
		return 0
	}

	f := make([][]int, N+1)
	for i := 0; i <= N; i++ {
		f[i] = make([]int, M+1)
	}

	cal := func(i, j int) int {
		ret := 0
		for li := 0; li < i; li++ {
			for lj := 0; lj < j; lj++ {
				v := f[i][j] - f[li][j] - f[i][lj] + f[li][lj]
				if v > ret && v == (i-li)*(j-lj) {
					ret = v
				}
			}
		}
		return ret
	}

	res := 0
	for i := 1; i <= N; i++ {
		for j := 1; j <= M; j++ {
			f[i][j] = f[i][j-1] + f[i-1][j] - f[i-1][j-1]
			if matrix[i-1][j-1] == '1' {
				f[i][j]++
				if v := cal(i, j); v > res {
					res = v
				}
			}
		}
	}
	return res
}

func main() {
	cases := [][][]byte{
		{
			{'1', '0', '1', '0', '0'},
			{'1', '0', '1', '1', '1'},
			{'1', '1', '1', '1', '1'},
			{'1', '0', '0', '1', '0'},
		},
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(maximalRectangle0(c))
		fmt.Println(maximalRectangle2(c))
	}
}
