package main

import (
	"container/list"
	"fmt"
)

// https://leetcode-cn.com/problems/alien-dictionary/

func alienOrder(words []string) string {
	n := len(words)
	if n == 0 {
		return ""
	}

	order := make(map[byte]byte)
	l := list.New()
	for i := 0; i < n; i++ {
		l.PushBack([]byte(words[i]))
	}

	for pos := 0; l.Len() > 0; pos++ {
		for p := l.Front(); p != nil; p = p.Next() {
			s := p.Value.([]byte)
			if len(s) <= pos {
				tmp := p.Next()
				l.Remove(p)
				p = tmp
			}
		}
	}
	return string(order['a'])
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
	var x string = "xy"
	var y string = x[:1]
	fmt.Println(x, y)
}
