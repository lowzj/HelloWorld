package main

import "fmt"

func main() {
	cases := [][][]int{
		{
			{8}, {5, 7, 7, 8, 8, 10},
		},
		{
			{8}, {5, 7, 7, 8, 8, 10},
		},
	}

	realCase := cases[1:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(c)
	}
}
