package main

import (
	"fmt"
)

// https://leetcode-cn.com/problems/create-maximum-number/
// 321. 拼接最大数 | Create Maximum Number
//------------------------------------------------------------------------------

//------------------------------------------------------------------------------
// Solution
//
// 复杂度分析:
//   * 时间:
//   * 空间:
func maxNumber(a []int, b []int, k int) []int {
	N, M := len(a), len(b)
	max := func(x []int, remain int, st []int) []int {
		n, p := len(x), 0
		k := n - remain
		for i := 0; i < n; i++ {
			for ; p > 0 && k > 0 && st[p-1] < x[i]; k-- {
				p--
			}
			if p < remain {
				st[p], p = x[i], p+1
			} else {
				k--
			}
		}
		return st[:p]
	}
	sta, stb, start := make([]int, k), make([]int, k), 0
	if k > M {
		start = k - M
	}
	var res []int
	for i := start; i <= k && i <= N; i++ {
		tmpa := max(a, i, sta)
		tmpb := max(b, k-i, stb)
		tmp := make([]int, 0, k)
		j, l := 0, 0
		for j < len(tmpa) && l < len(tmpb) {
			jj, ll := j, l
			for jj < len(tmpa) && ll < len(tmpb) && tmpa[jj] == tmpb[ll] {
				jj++
				ll++
			}
			if ll == len(tmpb) || jj != len(tmpa) && tmpa[jj] > tmpb[ll] {
				tmp = append(tmp, tmpa[j])
				j++
			} else {
				tmp = append(tmp, tmpb[l])
				l++
			}
		}
		for ; j < len(tmpa); j++ {
			tmp = append(tmp, tmpa[j])
		}
		for ; l < len(tmpb); l++ {
			tmp = append(tmp, tmpb[l])
		}
		if len(res) != k {
			res = tmp
		}
		j = 0
		for ; j < len(tmp) && res[j] == tmp[j]; j++ {
		}
		if j < len(tmp) && res[j] < tmp[j] {
			res = tmp
		}
	}
	return res
}

//------------------------------------------------------------------------------
// main

func main() {
	cases := []struct {
		a []int
		b []int
		k int
	}{
		{
			[]int{3, 4, 6, 5},
			[]int{9, 1, 2, 5, 8, 3},
			5,
		},
		{
			[]int{6, 7},
			[]int{6, 0, 4},
			5,
		},
		{
			[]int{1, 2},
			[]int{},
			2,
		},
		{
			[]int{5, 6, 8},
			[]int{6, 4, 0},
			3,
		},
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(maxNumber(c.a, c.b, c.k))
	}
}
