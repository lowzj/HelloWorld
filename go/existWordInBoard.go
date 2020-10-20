package main

import "fmt"

// https://leetcode-cn.com/problems/word-search/

func exist(board [][]byte, word string) bool {
	n := len(board)
	if n == 0 {
		return false
	}
	m := len(board[0])
	if m == 0 {
		return false
	}

	mark := make([]bool, n*m)
	st := &wordStack{nums: make([]int, 0, 3*len(word))}

	var search func(row, col int) bool
	search = func(row, col int) bool {
		wordLen := len(word)
		st.push(row, col, 0)
		for i := m*n - 1; i >= 0; i-- {
			mark[i] = false
		}

		for st.len() > 0 {
			i, j, pos := st.pop()
			if pos < 0 {
				return false
			}
			if mark[i*m+j] || board[i][j] != word[pos] {
				mark[i*m+j] = false
				continue
			}
			mark[i*m+j] = true
			if pos+1 == wordLen {
				return true
			}

			// save the mark status of (i,j)
			st.push(i, j, pos)
			if i > 0 && !mark[(i-1)*m+j] {
				st.push(i-1, j, pos+1)
			}
			if i < n-1 && !mark[(i+1)*m+j] {
				st.push(i+1, j, pos+1)
			}
			if j > 0 && !mark[i*m+j-1] {
				st.push(i, j-1, pos+1)
			}
			if j < m-1 && !mark[i*m+j+1] {
				st.push(i, j+1, pos+1)
			}
		}
		return false
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] == word[0] {
				if search(i, j) {
					return true
				}
				for st.len() > 0 {
					st.pop()
				}
			}
		}
	}

	return false
}

type wordStack struct {
	nums []int
}

func (ws *wordStack) len() int {
	return len(ws.nums)
}

func (ws *wordStack) push(i, j, pos int) {
	ws.nums = append(ws.nums, i, j, pos)
}

func (ws *wordStack) pop() (i, j, pos int) {
	n := len(ws.nums)
	if n == 0 {
		return -1, -1, -1
	}
	i, j, pos = ws.nums[n-3], ws.nums[n-2], ws.nums[n-1]
	ws.nums = ws.nums[:n-3]
	return i, j, pos
}

func main() {
	cases := []struct {
		word  string
		board [][]byte
	}{
		{
			word: "ABCCED",
			board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
		},
		{
			word: "ABCESEEEFS",
			board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'E', 'S'},
				{'A', 'D', 'E', 'E'},
			},
		},
		{
			word: "a",
			board: [][]byte{
				{'a'},
			},
		},
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(exist(c.board, c.word))
	}
}
