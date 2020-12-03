package main

import (
	"fmt"
)

// https://leetcode-cn.com/problems/count-primes/
// 204. 计数质数 | Count Primes
//------------------------------------------------------------------------------

//------------------------------------------------------------------------------
// Solution
//
// 复杂度分析:
//   * 时间: 小于 O(n*lgn). 循环次数 < 调和级数(1+1/2+1/3+...+1/n)-1 < lgn
//   * 空间: O(n)
func countPrimes(n int) int {
	if n < 2 {
		return 0
	}
	f := make([]bool, n)
	res := 0
	for i := 2; i < n; i++ {
		if !f[i] {
			res++
			for x := 2 * i; x < n; x += i {
				f[x] = true
			}
		}
	}
	return res
}

//------------------------------------------------------------------------------
// main

func main() {
	cases := []int{
		10,
		2,
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(countPrimes(c))
	}
}
