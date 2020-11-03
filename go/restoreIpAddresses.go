package main

import (
	"fmt"
	"strconv"
)

func restoreIpAddresses(s string) []string {
	n := len(s)
	if n < 4 || n > 12 {
		return []string{}
	}

	pos := [4]int{}
	var res []string
	var search func(idx, cnt int)
	search = func(idx, cnt int) {
		if idx == n && cnt == 4 {
			item := make([]byte, n+3)
			for i, j := 0, 0; i < n && j < 4; i++ {
				if i == pos[j] {
					item[i+j] = '.'
					j++
				}
				item[i+j] = s[i]
			}
			res = append(res, string(item))
			return
		}
		if idx >= n || cnt >= 4 {
			return
		}

		if s[idx] == '0' {
			pos[cnt] = idx + 1
			search(idx+1, cnt+1)
			return
		}
		for i := 1; i <= 3 && idx+i <= n; i++ {
			if x, e := strconv.Atoi(s[idx : idx+i]); e == nil && x > 0 && x <= 255 {
				pos[cnt] = idx + i
				search(idx+i, cnt+1)
			}
		}
	}
	search(0, 0)
	return res
}

func main() {
	cases := []string{
		"25525511135",
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(restoreIpAddresses(c))
	}
}
