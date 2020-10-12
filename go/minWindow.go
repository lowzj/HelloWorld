package main

import "fmt"

func minWindow(s string, t string) string {
	const N = 60 // 'z'-'A'
	sl := len(s)
	m := [N]int{}
	for i := 0; i < len(t); i++ {
		m[t[i]-'A']++
	}
	count := 0
	for i := 0; i < N; i++ {
		if m[i] > 0 {
			count++
		}
	}
	record := [N]int{}
	start, end := -1, -1
	for i := 0; i < sl; i++ {
		idx := s[i] - 'A'
		if m[idx] <= 0 {
			continue
		}

		record[idx]++
		if start < 0 {
			start = i
		}
		if record[idx] == m[idx] {
			count--
		}
		if count == 0 {
			end = i
			for ; start <= end; start++ {
				idx := s[start] - 'A'
				if m[idx] == 0 {
					continue
				}
				if record[idx] > m[idx] {
					record[idx]--
				} else {
					break
				}
			}
			break
		}
	}
	if end < 0 {
		return ""
	}

	for i, j := start, end+1; j < sl; j++ {
		jdx := s[j] - 'A'
		if m[jdx] == 0 {
			continue
		}
		record[jdx]++
		for ; i <= j; i++ {
			idx := s[i] - 'A'
			if m[idx] == 0 {
				continue
			}
			if record[idx] > m[idx] {
				record[idx]--
			} else {
				break
			}
		}
		if j-i < end-start {
			start, end = i, j
		}
	}

	return s[start : end+1]
}

func main() {
	cases := [][]string{
		{
			"ADOBECODEBBANC",
			"ABC",
		},
		{
			"aab",
			"ab",
		},
		{
			"bbaac",
			"aba",
		},
		{
			"bdab",
			"ab",
		},
	}
	for i, v := range cases {
		fmt.Printf("## case %d\nS: %s\nT: %s\n", i, v[0], v[1])
		fmt.Println("R:", minWindow(v[0], v[1]))
	}
}
