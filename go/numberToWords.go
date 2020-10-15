package main

import "fmt"

// https://leetcode-cn.com/problems/integer-to-english-words/

func numberToWords(num int) string {
	words := [][]string{
		{"", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine",
			"Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen",
			"Seventeen", "Eighteen", "Nineteen"},
		{"Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"},
	}
	units := []string{"", " Thousand", " Million", " Billion"}
	if num == 0 {
		return "Zero"
	}

	var res string
	for c, i := 0, num; i > 0; c, i = c+1, i/1000 {
		n, m, k := i%1000, i%100, i%10
		if n <= 0 {
			continue
		}
		var s string
		if n >= 100 {
			s = words[0][n/100] + " Hundred"
		}
		if m > 0 {
			if n >= 100 {
				s += " "
			}
			if m < 20 {
				s += words[0][m]
			} else {
				s += words[1][m/10-2]
				if k > 0 {
					s += " " + words[0][k]
				}
			}
		}
		if len(res) > 0 {
			res = " " + res
		}
		res = s + units[c] + res
	}
	return res
}

func main() {
	cases := []int{
		123,
		1234567891,
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(numberToWords(c))
	}
}
