package main

type land struct {
	r, c int
}

func numIslands(grid [][]byte) int {
	var lands []*land
	mark := make([][]bool, len(grid))
	for i := 0; i < len(grid); i++ {
		mark[i] = make([]bool, len(grid[i]))
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				lands = append(lands, &land{i, j})
			}
		}
	}
	var walk func(r, c int)
	walk = func(r, c int) {
		if r < 0 || r >= len(grid) ||
			c < 0 || c >= len(grid[r]) ||
			mark[r][c] || grid[r][c] != '1' {
			return
		}
		mark[r][c] = true
		walk(r-1, c) // up
		walk(r+1, c) // down
		walk(r, c-1) // left
		walk(r, c+1) // right
	}

	res := 0
	for pos := 0; pos < len(lands); pos++ {
		l := lands[pos]
		if mark[l.r][l.c] {
			continue
		}
		walk(l.r, l.c)
		res++
	}

	return res
}
