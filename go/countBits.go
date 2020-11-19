package main

import (
	"fmt"
)

// https://leetcode-cn.com/problems/counting-bits/
//------------------------------------------------------------------------------

//------------------------------------------------------------------------------
// Solution 1
// 令 f[i] 为数字 i 的二进制中1的个数.
// 因为 i = 2^k + j, j 范围在 [0,2^k-1], 2^k 中1的个数是1
// 所以 f[i] = 1 + f[j].
// 例如:
//   f[7] = f[2^2+3] = 1 + f[3]
//   f[8] = f[2^3+0] = 1 + f[0]
//   f[9] = f[2^3+1] = 1 + f[1]
// 即每次遇到 2 的次幂时, j 都要从0开始计数.
// 可以通过 i&(i-1) == 0 来快速判断一个数是否为 2 的次幂.
// 算法复杂度:
//   * 时间: O(n)
//   * 空间: O(1)
func countBits(num int) []int {
	res := make([]int, num+1)
	res[0] = 0
	for i, j := 1, 0; i <= num; i++ {
		if i&(i-1) == 0 {
			j = 0
		}
		res[i] = res[j] + 1
		j++
	}
	return res
}

func main() {
	cases := []int{
		4, 5, 10,
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(countBits(c))
	}
}
