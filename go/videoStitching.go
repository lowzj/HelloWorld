package main

import (
	"fmt"
	"sort"
)

// https://leetcode-cn.com/problems/video-stitching/

func videoStitching(clips [][]int, T int) int {
	n := len(clips)
	if n == 0 {
		return -1
	}

	sort.Sort(videos(clips))

	ends := make([]int, n+1)
	end, cnt := 0, 0
	for i := 0; i < n; i++ {
		if i == 0 && clips[i][0] != 0 ||
			clips[i][0] > end {
			break
		}
		if clips[i][1] > end {
			end = clips[i][1]
			for cnt > 0 && ends[cnt] > clips[i][0] && ends[cnt-1] >= clips[i][0] {
				cnt--
			}
			cnt++
			ends[cnt] = end
		}
		if end >= T {
			break
		}
	}
	if end >= T {
		return cnt
	}
	return -1
}

var _ sort.Interface = &videos{}

type videos [][]int

func (v videos) Len() int {
	return len(v)
}
func (v videos) Less(i, j int) bool {
	if v[i][0] == v[j][0] {
		return v[j][1] < v[i][1]
	}
	return v[i][0] < v[j][0]
}
func (v videos) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}

func main() {
	cases := []struct {
		clips [][]int
		t     int
	}{
		{
			clips: [][]int{{0, 2}, {4, 6}, {8, 10}, {1, 9}, {1, 5}, {5, 9}},
			t:     10,
		},
		{
			clips: [][]int{
				{0, 1}, {6, 8}, {0, 2}, {5, 6}, {0, 4}, {0, 3}, {6, 7}, {1, 3}, {4, 7}, {1, 4}, {2, 5}, {2, 6}, {3, 4}, {4, 5}, {5, 7}, {6, 9},
			},
			t: 9,
		},
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(videoStitching(c.clips, c.t))
	}
}
