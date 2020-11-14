package main

import (
	"fmt"
)

// https://leetcode-cn.com/problems/super-egg-drop/

//-----------------------------------------------------------------------------
func superEggDrop(K int, N int) int {
	return superEggDrop01(K, N)
}

//-----------------------------------------------------------------------------
// Solution 1
// 设 f[i][j] 表示有 i 个鸡蛋, j 层楼, 确定所求楼层需要移动的最少步数.
// 对于 f[i][j], 即有 i 个鸡蛋, j 层楼, 假设m为试验楼层, 总楼层高 j:
//   1. 若鸡蛋碎了, 说明所求楼层 f 在[0,m-1]层之间, 鸡蛋数 i-1, 即 f[i-1][m-1]
//   2. 若鸡蛋没碎, 说明所求楼层 f 在[m,j]层之间, 鸡蛋数不变 i, 即 f[i][j-m]
//   3. 其中1,2只能有一种情况成立, 所以 f[i][j] 需要在下面两个值中选择:
//        a. f[i][j] = f[i-1][m-1] + 1
//        b. f[i][j] = f[i][j-m] + 1
//      而为了确定 f 楼层, 需要考虑最坏下情况要移动多少步, 即选择a,b中的最大值:
//        f[i][j] = MAX(f[i-1][m-1], f[i][j-m]) + 1
//   4. 由于0层楼肯定不会碎, 所以无需试验, 另 m-1>0 且 j-m>0, 所以 m 范围: [2,j-1]
//   5. 综上: f[i][j] = 1 + MIN(
//                    MAX(f[i-1][1],f[i][j-2]), MAX(f[i-1][2],f[i][j-3]),...,
//                    MAX(f[i-1][j-3],f[i][2]), MAX(f[i-1][j-2],f[i][1]))
//
// 考虑初始值: 只有1个鸡蛋时, 无论多少层楼, 必须从1楼开始逐层向上试验; 所以移动步数等于楼层高度.
// 复杂度分析:
//   * 空间: O(K*N)
//   * 时间: O(K*N^2)
//
// 但是 O(K*N^2) 不会被AC, 报TLE, 需要优化.
// 上面第 5 步是用 O(N) 的复杂度求最小值, 可以优化成 O(lgN), 分析如下:
// 首先可以得出的2点:
//   a. f[i][j] >= f[i][j-1] 且 f[i][j] <= f[i][j-1] + 1
//   b. f[i][j] <= f[i-1][j]
// 在 f[i-1][1..j-1] 中搜索等于 f[i][j-1] 的最小的下标 m, 则:
//   f[i][j] = MAX(f[i-1][m-1], f[i][j-m]) + 1

func superEggDrop00(K int, N int) int {
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	if N < 3 {
		return N
	}
	f := make([][]int, K+1)
	for i := 0; i <= K; i++ {
		f[i] = make([]int, N+1)
	}
	for i := 0; i <= N; i++ {
		f[1][i] = i
	}
	for i := 0; i <= K; i++ {
		f[i][1], f[i][2] = 1, 2
	}

	for i := 2; i <= K; i++ {
		for j := 2; j < N; j++ {
			for n := (j + 1) / 2; n >= 1; n-- {
				tmp := min(max(f[i-1][n], f[i][j-n]), max(f[i-1][j-n], f[i][n])) + 1
				if f[i][j+1] > 0 {
					f[i][j+1] = min(f[i][j+1], tmp)
				} else {
					f[i][j+1] = tmp
				}
			}
			print2(f[i], i)
		}
	}

	return f[K][N]
}

func superEggDrop01(K int, N int) int {
	if N < 3 {
		return N
	}

	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	biSearch := func(f []int, l, r, target int) int {
		for mid := (l + r) / 2; l <= r; {
			if f[mid] == target {
				r = mid - 1
			} else if f[mid] > target {
				r = mid
			} else {
				l = mid + 1
			}
			mid = (l + r) / 2
		}
		return l
	}

	f := make([][]int, K+1)
	for i := 0; i <= K; i++ {
		f[i] = make([]int, N+1)
		f[i][1], f[i][2] = 1, 2
	}
	for i := 0; i <= N; i++ {
		f[1][i] = i
	}
	for i := 2; i <= K; i++ {
		for j := 2; j < N; j++ {
			m := biSearch(f[i-1], 1, j, f[i][j])
			f[i][j+1] = max(f[i-1][m-1], f[i][j+1-m]) + 1
		}
	}

	return f[K][N]
}

//-----------------------------------------------------------------------------
// Solution 2
// 逆向考虑, 有 K 个鸡蛋, 移动 step 步, 能确定的最大楼层数 N 是多少.
// 令 f[i][j] 表示j个鸡蛋, 移动试验 i 步后, 能确定的最大楼层数.
// 现在考虑相同步数下多一个鸡蛋的情况 f[i][j+1]:
// 因为多了一个鸡蛋, 所以可以允许试验中多碎一个鸡蛋, 即多选择一个楼层进行试验:
//   * 若鸡蛋碎了, 则还剩 j 个鸡蛋和 i-1 步, 此时能确定的最大楼层数即 f[i-1][j]
//   * 若鸡蛋没碎, 则还剩 j+1 个鸡蛋和 i-1 步, 此时能确定的最大楼层数即 f[i-1]j+1]
//   * 包含试验楼, 所以 f[i][j+1] = f[i-1][j] + f[i-1][j+1] + 1
// 计算得到大于 N 的最小 step 值后即停止.
// 因为 f[i] 仅与 f[i-1] 数据相关, 所以可以使用一维滚动数组, 基于f[i-1], 从后向前来计算f[i]的值.
// 复杂度:
//   * 空间: O(K)
//   * 时间: 在 O(K*lgN), O(K*N) 之间
//
// 下表的数据更为直观点:
//      K |   1 |   2 |   3 |   4
// -------+-----+-----+-----+-----
// step 1 |   1 |   1 |   1 |   1
//      2 |   2 |   3 |   3 |   3
//      3 |   3 |   6 |   7 |   7
//      4 |   4 |  10 |  14 |  15
//      5 |   5 |  15 |  25 |  30
//      6 |   6 |  21 |  41 |  56
//      7 |   7 |  28 |  63 |  98
//      8 |   8 |  36 |  92 | 162

func superEggDrop1(K int, N int) int {
	if K == 1 {
		return N
	}
	f := make([]int, K+1)
	step := 0
	for ; f[K] < N; step++ {
		for i := K; i >= 1; i-- {
			f[i] += f[i-1] + 1
		}
	}
	return step
}

func print2(f []int, idx int) {
	fmt.Printf("// %6d", idx)
	for i := 0; i < len(f); i++ {
		fmt.Printf(" | %3d", f[i])
	}
	fmt.Println()
}

func main() {
	cases := [][]int{
		{4, 100},
		{7, 10000},
		{1, 2},
		{2, 6},
		{2, 9},
		{3, 14},
		{3, 31},
		{3, 32},
		{7, 10},
	}

	//n := len(cases)
	n := 1
	realCase := cases[n-1:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(superEggDrop(c[0], c[1]))
		fmt.Println(superEggDrop1(c[0], c[1]))
	}
}
