package main

import (
	"fmt"
)

// https://leetcode-cn.com/problems/gray-code/
// 89. 格雷编码 | Gray Code
//------------------------------------------------------------------------------

//------------------------------------------------------------------------------
// Solution
//
// 找规律
//     i res | bit  | xor res[i-1]
//     ------+------+-------------
//     0.  0 | 0000 |
//     1.  1 | 0001 |  1
//     2.  3 | 0011 |  2
//     3.  2 | 0010 |  1
//     4.  6 | 0110 |  4
//     5.  7 | 0111 |  1
//     6.  5 | 0101 |  2
//     7.  4 | 0100 |  1
//     8. 12 | 1100 |  8
//     9. 13 | 1101 |  1
//    10. 15 | 1111 |  2
//    11. 14 | 1110 |  1
//    12. 10 | 1010 |  4
//    13. 11 | 1011 |  1
//    14.  9 | 1001 |  2
//    15.  8 | 1000 |  1
// 复杂度分析:
//   * 时间:
//   * 空间:
func grayCode(n int) []int {
	if n == 0 {
		return []int{0}
	}
	N := 1 << uint(n)
	res := make([]int, N)
	res[0], res[1] = 0, 1
	for i, e := 1, 2; i < n; i++ {
		res[e] = e
		copy(res[e+1:], res[1:e])
		e <<= 1
	}
	for i := 1; i < N; i++ {
		res[i] ^= res[i-1]
	}
	return res
}

// 看前两列的规律: res[i] = i^(i>>1)
func grayCode1(n int) []int {
	var res []int
	N := 1 << uint(n)
	for i := 0; i < N; i++ {
		res = append(res, i^(i>>1))
	}
	return res
}

//------------------------------------------------------------------------------
// main

func main() {
	cases := []int{
		3, 4,
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(grayCode(c))
		fmt.Println(grayCode1(c))
	}
	a := []int{0, 1, 3, 2, 6, 7, 5, 4, 12, 13, 15, 14, 10, 11, 9, 8}
	for i, v := range a {
		fmt.Printf("%2d. %2d | %04b |", i, v, v)
		if i != 0 {
			fmt.Printf(" %2d\n", v^a[i-1])
		} else {
			fmt.Println()
		}
	}
}
