package main

import "fmt"

// https://leetcode-cn.com/problems/count-numbers-with-unique-digits/

//------------------------------------------------------------------------------
// 设 f[i] 为 [0,10^i) 的各个位都不相同的数的个数.
// 因为对于一个 i 位数
//   * 首位不为0, 所以首位有9种选择;
//   * 其余位可以为0, 减去首位的数字, 还有9个数字可选总共有: A(9, m-1)
//   * 所以 i 位数的个数: 9 * A(9,i-1)
//   * 则 f[i] = f[i-1] + 9 * A(9,i-1)
// 由上面分析可得, 位数大于10的数, 肯定不符合条件, 满足条件的数的个数均为 f[10]

func countNumbersWithUniqueDigits(n int) int {
	const N = 11
	f := make([]int, N)
	f[0], f[1] = 1, 10
	base, fact := 9, 9
	for i := 2; i < N; i++ {
		base *= fact
		f[i] = f[i-1] + base
		fact--
	}
	if n < N {
		return f[n]
	}
	return f[N-1]
}

func main() {
	cases := []int{
		11,
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(countNumbersWithUniqueDigits(c))
	}
}
