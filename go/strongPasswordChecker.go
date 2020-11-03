package main

import "fmt"

func strongPasswordChecker(s string) int {
	N := len(s)
	if N < 3 {
		return 6 - N
	}
	fmt.Println(N)

	lower, upper, digit := false, false, false
	deleteCnt := N - 20
	if deleteCnt < 0 {
		deleteCnt = 0
	}
	dupCnt, preI := 0, 0
	for i, delCnt := 0, deleteCnt; i < N; i++ {
		if s[i] != s[preI] {
			preI = i
		}
		if i-preI >= 2 {
			if delCnt > 0 {
				delCnt--
				continue
			}
			dupCnt++
			preI = i + 1
		}
		if s[i] >= 'a' && s[i] <= 'z' {
			lower = true
		}
		if s[i] >= 'A' && s[i] <= 'Z' {
			upper = true
		}
		if s[i] >= '0' && s[i] <= '9' {
			digit = true
		}
	}
	if N+dupCnt < 6 {
		dupCnt = 6 - N
	}
	result := dupCnt + deleteCnt
	addOp := func(flag bool) {
		if !flag && dupCnt <= 0 {
			result++
		}
		if !flag && dupCnt > 0 {
			dupCnt--
		}
	}
	addOp(lower)
	addOp(upper)
	addOp(digit)
	return result
}

func main() {
	cases := []string{
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
		"ABABABABABABABABABAB1",
		"aaaabbaaabbaaa123456A",
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(strongPasswordChecker(c))
	}
}
