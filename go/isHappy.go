package main

import "fmt"

// https://leetcode-cn.com/problems/happy-number/

func isHappy(n int) bool {
	N := 1001
	hs := make([]int, N)

	sum := func(x int) int {
		s := 0
		for tmp := x; tmp > 0; tmp /= 10 {
			s += (tmp % 10) * (tmp % 10)
		}
		return s
	}
	for i := 0; i < N; i++ {
		hs[i] = sum(i)
	}

	find := func(n int) bool {
		for i := 0; i < N; i++ {
			if n = sum(n); hs[n] == 1 {
				return true
			}
		}
		return false
	}
	return find(n)
}

func main() {
	cases := []int{
		19, 91,
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(isHappy(c))
	}
}
