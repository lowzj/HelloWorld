package main

import "fmt"

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
	a := 1000000007
	fmt.Println(a)
	fmt.Println((1143383095 * 15) % a)
}
