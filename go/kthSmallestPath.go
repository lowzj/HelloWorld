package main

import (
	"fmt"
	"math/big"
)

// https://leetcode-cn.com/problems/kth-smallest-instructions/

func kthSmallestPath(destination []int, k int) string {
	if len(destination) != 2 {
		return ""
	}

	zero := big.NewInt(0)
	cal := func(n, m, k int) (int, bool) {
		b := big.NewInt(1).Binomial(int64(n), int64(m))
		b.Sub(b, big.NewInt(int64(k)))
		if b.Cmp(zero) < 0 {
			return -int(b.Int64()), true
		}
		return 0, false
	}

	N, M := destination[1], destination[0]
	res := make([]byte, N+M)

	i, n, m := 0, N, M
	for ; i < N+M && k != 0 && n > 0 && m > 0; i++ {
		// 'H' 开头的组合个数:
		// 从剩余的 n+m-1 个中选则 m 个位置填上 'V', 其余填'H'
		v, ok := cal(n+m-1, m, k)
		if !ok {
			res[i] = 'H'
			n--
		} else {
			res[i] = 'V'
			m--
			k = v
		}
	}
	for j := 0; j < m; j, i = j+1, i+1 {
		res[i] = 'V'
	}
	for j := 0; j < n; j, i = j+1, i+1 {
		res[i] = 'H'
	}
	return string(res)
}

// HHHVV 00011 3
// HHVHV 00101 5   2
// HHVVH 00110 6   1
// HVHHV 01001 9   3
// HVHVH 01010 10  1
// HVVHH 01100 12  2
// VHHHV 10001 17  5
// VHHVH 10010 18  1
// VHVHH 10100 20  2
// VVHHH 11000 24  4

func main() {
	cases := [][]int{
		{2, 3, 8},
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(kthSmallestPath(c[:2], c[2]))
	}
}
