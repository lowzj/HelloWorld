package main

import (
	"fmt"
	"math/rand"
	"time"
)

// https://leetcode-cn.com/problems/k-closest-points-to-origin/

func kClosest(points [][]int, K int) [][]int {
	n := len(points)
	if n == 0 {
		return [][]int{}
	}

	dist := func(i int) int {
		return points[i][0]*points[i][0] + points[i][1]*points[i][1]
	}
	swap := func(i, j int) { points[i], points[j] = points[j], points[i] }

	rand.Seed(time.Now().UnixNano())
	var find func(l, r int)
	find = func(l, r int) {
		if l >= r {
			return
		}

		p := rand.Intn(r-l) + l
		tmpPoint, d := points[p], dist(p)
		swap(p, l)
		p = l
		for j := r - 1; p < j; {
			if dist(p+1) >= d {
				swap(p+1, j)
				j--
			} else {
				points[p] = points[p+1]
				p++
			}
		}
		points[p] = tmpPoint

		if p == K || p == K-1 {
			return
		}
		if p < K {
			find(p+1, r)
		} else {
			find(l, p)
		}
	}

	find(0, n)
	return points[:K]
}

func main() {
	cases := []struct {
		k      int
		points [][]int
	}{
		{
			k: 2,
			points: [][]int{
				{3, 3}, {5, -1}, {-2, 4},
			},
		},
		{
			k: 5,
			points: [][]int{
				{68, 97}, {34, -84}, {60, 100}, {2, 31}, {-27, -38},
				{-73, -74}, {-55, -39}, {62, 91}, {62, 92}, {-57, -67},
			},
		},
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(kClosest(c.points, c.k))
	}
}
