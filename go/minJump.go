package main

import (
	"fmt"
)

// https://leetcode-cn.com/problems/zui-xiao-tiao-yue-ci-shu/
// LCP 09. 最小跳跃次数
//------------------------------------------------------------------------------
func minJump(jump []int) int {
	return minJump0(jump)
}

//------------------------------------------------------------------------------
// Solution 0
//
// BFS: 因为 i 位置可以向左侧任意跳, 所以时间复杂度是 O(N^2), 可以添加一个标记 start, 表示
// [0,start) 都已经搜索完毕, 下次向左侧跳的时候, 只需从 start 位置开始扫描.
//
// 复杂度分析:
//   * 时间: O(N). 240ms
//   * 空间: O(N)
func minJump0(jump []int) int {
	N := len(jump)
	if N == 0 {
		return 0
	}
	q, p, size := make([][2]int, N), 0, 0
	push := func(i, cnt int) {
		q[p][0], q[p][1] = i, cnt
		p = (p + 1) % N
		if size < N {
			size++
		}
	}
	pop := func() (int, int) {
		i := (p - size + N) % N
		size--
		return q[i][0], q[i][1]
	}

	viewed, start := make([]bool, N), 0
	push(0, 1)
	for size > 0 {
		i, cnt := pop()
		viewed[i] = true
		if i+jump[i] >= N {
			return cnt
		}
		if idx := i + jump[i]; !viewed[idx] {
			viewed[idx] = true
			push(idx, cnt+1)
		}
		for j := start; i != 0 && j < i; j++ {
			if !viewed[j] {
				viewed[j] = true
				push(j, cnt+1)
			}
		}
		if i+1 > start {
			start = i + 1
		}
	}
	return -1
}

//------------------------------------------------------------------------------
// main

func main() {
	cases := [][]int{
		// 3
		{2, 5, 1, 1, 1, 1},
		// 6
		{3, 7, 6, 1, 4, 3, 7, 8, 1, 2, 8, 5, 9, 8, 3, 2, 7, 5, 1, 1},
		// 16
		{
			24, 100, 2, 25, 27, 70, 78, 90, 66, 67, 25, 78, 21, 65, 62, 71, 77, 61, 83, 80, 62, 84, 24,
			37, 56, 55, 91, 13, 17, 6, 3, 40, 10, 51, 3, 88, 18, 25, 92, 40, 2, 41, 8, 43, 43, 35, 70, 82,
			71, 42, 37, 37, 92, 15, 87, 53, 46, 36, 78, 51, 93, 81, 2, 3, 79, 73, 83, 32, 74, 97, 55,
			87, 39, 98, 82, 73, 38, 68, 43, 30, 5, 55, 25, 74, 96, 19, 86, 76, 23, 24, 69, 14, 82, 50,
			74, 2, 8, 91, 36, 22, 59, 100, 34, 61, 4, 20, 36, 13, 39, 49, 5, 65, 21, 96, 10, 78, 30,
			72, 9, 87, 81, 43, 52, 69, 83, 72, 16, 77, 52, 99, 37, 14, 30, 95, 28, 31, 94, 13, 34, 71,
			42, 87, 53, 80, 21, 59, 45, 56, 51, 26, 38, 67, 43, 14, 10, 44, 49, 96, 55, 48, 74, 97,
			20, 53, 50, 92, 44, 10, 99, 55, 12, 72, 6, 88, 51, 51, 40, 8, 90, 89, 61, 90, 12, 88, 50,
			37, 33, 10, 25, 87, 99, 27, 8, 81, 64, 14, 100, 20, 64, 26, 38, 78, 11, 42, 3, 36, 24, 13,
			11, 16, 50, 32, 3, 73, 95, 3, 34, 13, 96, 59, 10, 56, 62, 69, 36, 34, 61, 65, 38, 95, 81,
			2, 54, 96, 5, 23, 51, 70, 83, 54, 29, 8, 100, 89, 36, 83, 46, 24, 23, 87, 85, 97, 41, 84,
			4, 70, 81, 75, 35, 16, 9, 93, 38, 98, 20, 43, 50, 56, 72, 59, 17, 21, 21, 90, 32, 62, 82,
			61, 54, 100, 84, 85, 88, 96, 64, 96, 34, 25, 69, 62, 13, 10, 48, 21, 75, 75, 72, 50, 28,
			38, 58, 8, 35, 5, 72, 24, 76, 31, 86, 35, 27, 73, 51, 71, 38, 20, 2, 50, 30, 93, 29, 57,
			63, 44, 19, 65, 100, 99, 13, 72, 52, 87, 99, 73, 55, 44, 48, 71, 30, 58, 25, 35, 22, 8,
			11, 27, 38, 98, 37, 49, 38, 50, 54, 36, 31, 47, 100, 1, 62, 81, 13, 43, 36, 26, 92, 25,
			47, 49, 89, 6, 78, 92, 61, 77, 92, 65, 94, 39, 3, 47, 66, 62, 87, 77, 79, 99, 72, 41, 92,
			92, 38, 67, 17, 92, 6, 42, 97, 50, 49, 13, 54, 41, 99, 49, 95, 11, 68, 52, 65, 1, 38, 26,
			20, 52, 57, 87, 65, 98, 73, 13, 40, 34, 13, 9, 97, 44, 56, 36, 95, 82, 32, 77, 4, 99, 62,
			22, 2, 99, 71, 32, 63, 95, 50, 90, 78, 60, 45, 54, 32, 87, 34, 82, 38, 83, 56, 67, 60, 80,
			16, 91, 90, 3, 47, 19, 73, 71, 15, 15, 40, 100, 34, 36, 22, 20, 93, 60, 64, 41, 69, 81, 35,
			3, 50, 26, 9, 32, 52, 76, 77, 12, 83, 55, 58, 31, 5, 92, 50, 7, 43, 12, 35, 11, 33, 25, 92,
			6, 70, 1, 60, 46, 49, 74, 10, 19, 99, 18, 2, 3, 64, 56, 64, 72, 38, 59, 83, 24, 97, 2, 59,
			52, 35, 90, 58, 41, 24, 73, 24, 2, 44, 99, 100, 72, 15, 94, 65, 59, 68, 49, 27, 33, 47,
			52, 71, 19, 96, 22, 48, 12, 85, 37, 71, 38, 35, 3, 66, 46, 32, 10, 16, 17, 36, 18, 87, 62,
			91, 90, 29, 21, 3, 32, 71, 51, 57, 37, 36, 1, 33, 39, 26, 1, 44, 50, 41, 44, 10, 88, 72,
			90, 52, 15, 97, 86, 71, 30, 47, 16, 91, 99, 33, 99, 33, 52, 27, 23, 67, 62, 85, 43, 93,
			60, 50, 98, 16, 39, 99, 57, 54, 53, 26, 25, 56, 48, 15, 65, 92, 34, 57, 96, 99, 76, 39,
			99, 93, 76, 64, 32, 4, 93, 33, 51, 13, 16, 63, 28, 84, 61, 97, 8, 9, 71, 89, 19, 62, 84,
			34, 9, 90, 20, 9, 18, 95, 9, 78, 80, 65, 81, 85, 92, 29, 66, 56, 43, 81, 44, 95, 9, 44, 85,
			21, 29, 19, 80, 1, 52, 82, 20, 79, 14, 98, 3, 41, 13, 11, 36, 4, 40, 3, 12, 13, 73, 82, 80,
			85, 81, 14, 9, 25, 82, 73, 72, 49, 33, 55, 96, 99, 58, 92, 33, 8, 89, 94, 55, 88, 89, 29,
			53, 74, 67, 9, 88, 90, 44, 93, 7, 95, 43, 25, 32, 94, 75, 33, 13, 29, 9, 86, 50, 92, 57,
			36, 27, 33, 2, 49, 28, 35, 82, 77, 61, 62, 11, 49, 8, 12, 78, 66, 6, 100, 5, 74, 10, 46,
			43, 1, 95, 89, 3, 81, 66, 61, 100, 53, 25, 95, 86, 5, 23, 77, 35, 24, 25, 9, 96, 33, 69,
			43, 60, 76, 10, 66, 49, 12, 98, 54, 97, 27, 25, 26, 9, 50, 72, 59, 83, 4, 97, 88, 8, 15,
			38, 73, 57, 52, 32, 75, 32, 30, 33, 12, 6, 94, 93, 25, 44, 85, 19, 28, 45, 68, 35, 40, 30,
			47, 87, 61, 50, 32, 37, 54, 76, 79, 70, 21, 65, 93, 54, 89, 55, 53, 24, 97, 17, 42, 24,
			59, 41, 5, 38, 66, 69, 71, 27, 17, 18, 21, 57, 7, 44, 35, 86, 60, 58, 6, 1, 98, 44, 96, 68,
			87, 43, 41, 63, 40, 30, 42, 52, 74, 32, 80, 26, 75, 44, 99, 34, 3, 15, 57, 73, 82, 71, 70,
			11, 50, 69, 62, 93, 50, 71, 15, 24, 93, 90, 70, 57, 51, 25, 56, 45, 82, 10, 80, 62, 75, 8,
			32, 15, 2, 85, 87, 54, 89, 7, 32, 79, 56, 64, 80, 28, 61, 37, 9, 78, 28, 38, 3, 12, 8, 19,
			70, 30, 83, 47, 77, 13, 74, 47, 16, 73, 17, 58, 64, 97, 98, 80, 44, 89, 43, 45, 78, 20,
			80, 64, 100, 76, 75, 31, 86, 87, 93, 19, 16, 54,
		},
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(minJump(c))
	}
}
