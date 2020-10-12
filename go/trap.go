package main

import "fmt"

func trap(height []int) int {
	// S(water) = S(total=maxHeight*length) -
	//            S(reverse) - S(sum(height))

	hl := len(height)
	if hl == 0 {
		return 0
	}
	left, right, sum := 0, 0, 0
	for i := 0; i < hl; i++ {
		if height[i] > height[left] {
			left = i
		}
		if height[i] >= height[right] {
			right = i
		}
		sum += height[i]
	}
	res := height[left]*hl - sum

	areaS := 0
	preMax := 0
	for i := 0; i < left; i++ {
		if i == 0 {
			preMax = height[i]
		}
		if height[i] > preMax {
			preMax = height[i]
		}
		areaS += height[left] - preMax
	}

	preMax = height[hl-1]
	for i := hl - 1; i > right; i-- {
		if i == 0 {
			preMax = height[i]
		}
		if height[i] > preMax {
			preMax = height[i]
		}
		areaS += height[left] - preMax
	}
	return res - areaS
}

func main() {
	h := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	res := trap(h)
	fmt.Println(res)
}
