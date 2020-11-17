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

// 转成中文

func read(n int64) string {
	const zero = "零"
	units := []string{"", "万", "亿", "兆"}
	if n == 0 {
		return zero
	}

	result, useZero := "", false
	for i, c := n, 0; i > 0; i /= 10000 {
		tmp := readFour(int(i % 10000))
		if tmp != zero {
			if useZero {
				result = zero + result
			}
			result = tmp + units[c] + result
		}
		useZero = i%10000/1000 <= 0 && len(result) > 0
		c++
	}
	return result
}

func readFour(n int) string {
	digits := []string{"零", "一", "二", "三", "四", "五", "六", "七", "八", "九"}
	units := []string{"", "十", "百", "千"}
	if n == 0 {
		return digits[0]
	}
	res, useZero := "", false
	for i, c := n, 0; i > 0; i /= 10 {
		idx := i % 10
		if useZero && idx > 0 {
			res = digits[0] + res
		}
		if idx > 0 {
			res = digits[idx] + units[c] + res
		}
		useZero = !useZero && idx == 0 && len(res) > 0
		c++
	}
	return res
}

func main() {
	cases := []int64{
		123,
		1234567891,
		1240567000,
		1240560000,
		1000000001,
		1000000000000001,
	}

	realCase := cases[2:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		//fmt.Println(numberToWords(int(c)))
		fmt.Println(read(c))
	}
}
