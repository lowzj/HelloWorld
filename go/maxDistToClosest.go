package main

import "fmt"

func maxDistToClosest(seats []int) int {
	N := len(seats)
	if N == 0 {
		return 0
	}
	max, pre := 0, 0
	for i := 0; i < N; i++ {
		if seats[i] == 0 {
			continue
		}
		if seats[pre] == 0 {
			max = i
		}
		if i-pre > 2*max {
			max = (i - pre) / 2
		}
		pre = i
	}
	if pre != N-1 && N-pre-1 > max {
		max = N - pre - 1
	}
	return max
}

func main() {
	cases := [][]int{
		{},
		{},
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(c)
	}
}
