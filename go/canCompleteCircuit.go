package main

import (
	"fmt"
)

// https://leetcode-cn.com/problems/gas-station/

//------------------------------------------------------------------------------

func canCompleteCircuit(gas []int, cost []int) int {
	N := len(gas)
	for i := 0; i < N; i++ {
		sg, sc := 0, 0
		for j := 0; j < N; j++ {
			idx := (i + j) % N
			sg += gas[idx]
			sc += cost[idx]
			if sg < sc {
				break
			}
		}
		if sg >= sc {
			return i
		}
	}
	return -1
}

func main() {
	cases := [][][]int{
		{
			{1, 2, 3, 4, 5},
			{3, 4, 5, 1, 2},
		},
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(canCompleteCircuit(c[0], c[1]))
	}
}
