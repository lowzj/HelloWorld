package main

import (
	"fmt"
	"strings"
)

// https://leetcode-cn.com/problems/text-justification/
//------------------------------------------------------------------------------

func fullJustify(words []string, maxWidth int) []string {
	N := len(words)
	const FMT = "%%%ds"
	var res []string
	add := func(start, end, total int) {
		line, spaces := "", maxWidth-total
		if start == end {
			line = words[end] + fmt.Sprintf(fmt.Sprintf(FMT, spaces), "")
		} else {
			line += words[start]
			for i := start + 1; i <= end; i++ {
				count := spaces / (end - i + 1)
				if count*(end-i+1) < spaces {
					count++
				}
				line += fmt.Sprintf(fmt.Sprintf(FMT, count), "") + words[i]
				spaces -= count
			}
		}
		res = append(res, line)
	}
	start, last, total := 0, -1, 0
	for i := 0; i < N; i++ {
		l := len(words[i])
		if length := total + l + (i - start); length >= maxWidth {
			if start != i && length != maxWidth {
				add(start, i-1, total)
				start, last, total = i, i-1, l
			} else {
				add(start, i, total+l)
				start, last, total = i+1, i, 0
			}
		} else {
			total += l
		}
	}
	if s := strings.Join(words[last+1:], " "); s != "" {
		if len(s) < maxWidth {
			s += fmt.Sprintf(fmt.Sprintf(FMT, maxWidth-len(s)), "")
		}
		res = append(res, s)
	}
	return res
}

func main() {
	cases := []struct {
		words    []string
		maxWidth int
	}{
		{
			[]string{"a"},
			2,
		},
		{
			words: []string{"This", "is", "an", "example",
				"of", "text", "justification."},
			maxWidth: 16,
		},
		{
			words:    []string{"What", "must", "be", "acknowledgment", "shall", "be"},
			maxWidth: 16,
		},
		{
			words: []string{"Science", "is", "what", "we", "understand", "well",
				"enough", "to", "explain", "to", "a", "computer.",
				"Art", "is", "everything", "else", "we", "do"},
			maxWidth: 20,
		},
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		res := fullJustify(c.words, c.maxWidth)
		for _, s := range res {
			fmt.Println(len(s), "|", s, "|")
		}
	}
}
