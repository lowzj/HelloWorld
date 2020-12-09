package main

import (
	"fmt"
)

// https://leetcode-cn.com/problems/split-array-into-fibonacci-sequence/
// 842. 将数组拆分成斐波那契序列 | Split Array into Fibonacci Sequence
//------------------------------------------------------------------------------

//------------------------------------------------------------------------------
// Solution
func splitIntoFibonacci(S string) []int {
	N := len(S)
	if N < 3 {
		return nil
	}
	toInt := func(s,e int) int {
		res := 0
		for i := s; i < e; i++ {
			res = res*10+int(S[i]-'0')
		}
		return res
	}
	search := func(s, target int) int {
		res := 0
		for i := s; i < N; i++ {
			res = res*10 + int(S[i]-'0')
			if res > target || res >= (1<<31-1) {
				return -1
			}
			if res == target {
				return i+1
			}
		}
		return -1
	}

	res := []int{}
	for i := 1; i < N-1; i++ {
		res = res[:0]
		f := toInt(0, i)
		if f < 0 || S[0] == '0' && i > 1 {
			return nil
		}
		res = append(res, f)
		for j := i+1; j < N; j++ {
			res = res[:1]
			s := toInt(i, j)
			if s < 0 || S[i] == '0' && j-i > 1 {
				break
			}
			res = append(res, s)
			for ff, ss, t, k, pre := f,s,f + s, j, j; k >= 0 && t >= 0; {
				if t < 0 || k < N && S[pre] == '0' && k-pre > 1 {
					break
				}
				res = append(res, t)
				if pre, k = k, search(k, t); k == N {
					return res
				}
				ff, ss = ss, t
				t = ff + ss
			}
		}
	}
	return nil
}

//------------------------------------------------------------------------------
// main

func main() {
	cases := []string{
		"123456579",
		"11235813",
		"1101111",
		"0123",
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(splitIntoFibonacci(c))
	}
}
