package main

import (
	"fmt"
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	v1, v2 := strings.Split(version1, "."), strings.Split(version2, ".")
	l1, l2 := len(v1), len(v2)

	i := 0
	for ; i < l1 && i < l2; i++ {
		a1, _ := strconv.Atoi(v1[i])
		a2, _ := strconv.Atoi(v2[i])
		if a1 < a2 {
			return -1
		}
		if a1 > a2 {
			return 1
		}
	}
	for ; i < l1; i++ {
		if v, _ := strconv.Atoi(v1[i]); v > 0 {
			return 1
		}
	}
	for ; i < l2; i++ {
		if v, _ := strconv.Atoi(v2[i]); v > 0 {
			return -1
		}
	}
	return 0
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
}
