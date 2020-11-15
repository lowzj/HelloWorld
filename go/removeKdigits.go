package main

import (
	"fmt"
)

// https://leetcode-cn.com/problems/remove-k-digits/

//------------------------------------------------------------------------------
// Solution 1
// 正序遍历, 统计逆序对, 游标 i 处将 i 之前的所有大于 num[i] 的元素删除, 直至 k 次为止, 不够
// k 次则继续向后遍历.

func removeKdigits(num string, k int) string {
	N := len(num)
	if k >= N {
		return "0"
	}

	cnt := make([]int, 10)
	inc := func(b byte) { cnt[b-'0']++ }
	dec := func(b byte) { cnt[b-'0']-- }
	sum := func(b byte) (sum int) {
		for i := int(b - '0'); i >= 0; i-- {
			sum += cnt[i]
		}
		return sum
	}

	pos, remove := make([][]int, 10), make(map[int]bool)
	get := func(b byte) int {
		i, n := int(b-'0'), len(pos[b-'0'])
		if n <= 0 {
			return -1
		}
		idx := pos[i][0]
		pos[i][0] = pos[i][n-1]
		if n > 2 {
			copy(pos[i], pos[i][1:])
		}
		pos[i] = pos[i][:n-1]
		dec(b)
		return idx
	}
	put := func(b byte, i int) {
		pos[b-'0'] = append(pos[b-'0'], i)
		inc(b)
	}
	del := func() int {
		for i := 9; i >= 0; i-- {
			if idx := get(byte('0' + i)); idx >= 0 {
				remove[idx] = true
				return idx
			}
		}
		return -1
	}

	max := byte('9')
	for i := 0; i < N && k > 0; i++ {
		put(num[i], i)
		for r := sum(max) - sum(num[i]); r > 0 && k > 0; r-- {
			// has reversed order
			k--
			del()
		}

		// delete leading zeros
		if zero := sum('0'); zero > 0 && sum(max) == zero {
			for ; zero > 0; zero-- {
				del()
			}
		}
	}
	for ; k > 0; k-- {
		del()
	}
	if len(remove) == N {
		return "0"
	}
	res := make([]byte, N-len(remove))
	for i, k := 0, 0; i < N; i++ {
		if !remove[i] {
			res[k] = num[i]
			k++
		}
	}
	return string(res)
}

func main() {
	cases := []struct {
		num string
		k   int
	}{
		{"1432219", 3},
		{"123456789", 3},
		{"10200", 2},
		{"10", 1},
		{"99123456789", 3},
		{"1234567890", 9},
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(removeKdigits(c.num, c.k))
	}
}
