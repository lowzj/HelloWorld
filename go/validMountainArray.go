package main

import "fmt"

// https://leetcode-cn.com/problems/valid-mountain-array/

func validMountainArray(A []int) bool {
	N := len(A)
	if N < 3 {
		return false
	}
	flag := true
	for i := 1; i < N; i++ {
		if A[i] == A[i-1] {
			return false
		}
		if i == 1 && A[i] < A[i-1] {
			return false
		}
		if A[i] < A[i-1] {
			flag = false
		}
		if !flag && A[i] > A[i-1] {
			return false
		}
	}
	return true
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
