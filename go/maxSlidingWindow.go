package main

import (
	"container/heap"
	"fmt"
)

// https://leetcode-cn.com/problems/sliding-window-maximum/

//------------------------------------------------------------------------------

func maxSlidingWindow(nums []int, k int) []int {
	return maxSlidingWindow0(nums, k)
}

//------------------------------------------------------------------------------
// Solution 1
// f: 记录窗口内最大值, 滚动更新
// f[0..k]: MAX(nums[i..i+k]), MAX(nums[i+1..i+k]),...,MAX(nums[i+k])
// 待调教

func maxSlidingWindow1(nums []int, k int) []int {
	N := len(nums)

	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	f := make([]int, k)
	f[k-1] = nums[k-1]
	for i := k - 2; i >= 0; i-- {
		f[i] = max(nums[i], f[i+1])
	}

	res := make([]int, N-k+1)
	now, offset := 0, 0
	next := func() (int, int) {
		if k == 1 {
			return 0, 0
		}
		if now == k-1 {
			offset = -1
		} else if now == 0 {
			offset = 1
		}
		defer func() { now += offset }()
		return now, now + offset
	}
	for i := k; i < N; i++ {
		c, n := next()
		res[i-k] = f[c]
		fmt.Println("##", i, res[i-k], c, n, nums[i-k:i], f)

		if c == 0 || c == k-1 {
			f[c] = nums[i]
		} else {
			f[c] = max(f[c-offset], nums[i])
		}
		f[n] = max(f[n], f[c])
		//fmt.Println("##", f)
	}
	c, _ := next()
	//fmt.Println("##", c, f)
	res[N-k] = f[c]
	return res
}

//------------------------------------------------------------------------------
// Solution 0
// 最大堆: O(n*lgk)
// 要记录元素在堆中的相对位置, 以便更新值及调整堆

func maxSlidingWindow0(nums []int, k int) []int {
	N := len(nums)

	h := &window{}
	m := make([]*windowNode, k)
	for i := 0; i < k; i++ {
		m[i] = &windowNode{nums[i], i}
		heap.Push(h, m[i])
	}

	res := make([]int, N-k+1)
	res[0] = h.top()
	for i := k; i < N; i++ {
		node := m[i%k]
		node.v = nums[i]
		heap.Fix(h, node.i)
		res[i-k+1] = h.top()
	}
	return res
}

type windowNode struct {
	v, i int
}

func (w *windowNode) String() string {
	return fmt.Sprintf("%d:%d", w.i, w.v)
}

var _ heap.Interface = &window{}

type window struct {
	items []*windowNode
}

func (w window) Len() int {
	return len(w.items)
}

func (w window) Less(i, j int) bool {
	if w.items[i].v == w.items[j].v {
		return w.items[i].i < w.items[j].i
	}
	return w.items[i].v > w.items[j].v
}

func (w window) Swap(i, j int) {
	// i 表示在堆中的位置, 两次交换保证相对位置不变, 只是值变了
	w.items[i].i, w.items[j].i = w.items[j].i, w.items[i].i
	w.items[i], w.items[j] = w.items[j], w.items[i]
}

func (w *window) Push(x interface{}) {
	w.items = append(w.items, x.(*windowNode))
}

func (w *window) Pop() interface{} {
	n := len(w.items)
	if n == 0 {
		return nil
	}
	top := w.items[n-1]
	w.items = w.items[:n-1]
	return top
}

func (w *window) top() int {
	return w.items[0].v
}

//------------------------------------------------------------------------------

func main() {
	cases := []struct {
		nums []int
		k    int
	}{
		{
			[]int{-6, -10, -7, -1, -9, 9, -8, -4, 10, -5, 2, 9, 0, -7, 7, 4, -2, -10, 8, 7},
			7,
		},
		{
			[]int{1, 3, -1, -3, 5, 3, 6, 7},
			3,
		},
		{
			[]int{1, 3, -1, -3, -2, 3, 6, 7},
			3,
		},
		{
			[]int{1, 2, 3},
			1,
		},
	}

	realCase := cases[0:1]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(maxSlidingWindow0(c.nums, c.k))
		fmt.Println(maxSlidingWindow1(c.nums, c.k))
	}
}
