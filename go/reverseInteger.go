package main

import (
	"fmt"
	"math"
)

// https://leetcode-cn.com/problems/reverse-integer/

func reverse(x int) int {
	sign := int64(1)
	if x < 0 {
		sign = -1
	}
	res := int64(0)
	for tmp := sign * int64(x); tmp > 0; tmp /= 10 {
		res = res*10 + tmp%10
	}
	res *= sign
	if res > math.MaxInt32 || res < math.MinInt32 {
		return 0
	}
	return int(res)
}

func main() {
	cases := []int{
		123, -123, math.MaxInt32, math.MinInt32,
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(reverse(c))
	}
}
