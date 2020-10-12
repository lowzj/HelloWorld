package main

import (
	"container/list"
	"fmt"
	"math"
	"sort"
)

// https://leetcode-cn.com/problems/container-with-most-water/

// -----------------------------------------------------------------------------
// solution 1
// 两条线段围成的容器面积, 是由较短线段决定的.
// 对于最短线段, 剩余线段与其围成的容器面积, 与他们的距离成正比, 不用考虑高度.
// 所以对于当前线段集合中的最短线段来说, 与其最长距离的线段，围成的容器面积最大.
// 思路如下:
// 1. 二元组表示线段 A[k](h, x), h 代表高度, x 代表 x 轴坐标, 用于计算距离
// 2. 对所有线段的二元组按高度生序排序之后, 从0开始遍历该集合(最短开始遍历)
// 3. 对于最短线段 A[k] 的最远左边线段 A[i] 和最远右边线段 A[j]
// 4. 则线段 A[k] 能够组成的最大面积:
//		S[k] = A[k].h * MAX(A[k].x-A[i].x, A[j].x-A[k].x)
// 5. 剔除线段 A[k]
// 6. 并重复3,4,5步骤, 最终可求得 MAX(S[0],...,S[n-1])
// 第3步中, 可以使用最小栈, 在O(1)时间内求得最远距离:
// * 新建两个最小栈: left, right, 分别用于求最远左边线段和最远右边线段
// * 按排序后的 A 从n-1到0
// * 将轴坐标 A[k].x 倒序压入栈 left
// * 将轴坐标 -A[k].x 倒序压入栈 right

func maxArea(height []int) int {
	hl := len(height)

	lines := make(Lines, hl)
	for i := 0; i < hl; i++ {
		lines[i] = &line{height[i], i}
	}
	sort.Sort(lines)
	left, right := newMinStack(), newMinStack()
	for i := lines.Len() - 1; i >= 0; i-- {
		left.Push(lines[i].i)
		right.Push(-lines[i].i)
	}
	area := 0
	for i := 0; i < lines.Len(); i++ {
		area = max(area,
			lines[i].h*max(lines[i].i-left.min, -right.min-lines[i].i))
		left.Pop()
		right.Pop()
	}
	return area
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func newMinStack() *minStack {
	return &minStack{
		l:   list.New(),
		min: math.MaxInt64,
	}
}

type minStack struct {
	l   *list.List
	min int
}

func (ms *minStack) Push(v int) {
	ms.l.PushFront(v - ms.min)
	if ms.min > v {
		ms.min = v
	}
}
func (ms *minStack) Pop() int {
	v := ms.l.Remove(ms.l.Front()).(int)
	res := v + ms.min
	if v < 0 {
		res = ms.min
		ms.min = res - v
	}
	return res
}
func (ms *minStack) Top() int {
	v := ms.l.Front().Value.(int)
	if v < 0 {
		return ms.min
	}
	return v + ms.min
}
func (ms *minStack) Min() int {
	return ms.min
}

type line struct {
	h, i int
}

var _ sort.Interface = &Lines{}

type Lines []*line

func (l Lines) Len() int {
	return len(l)
}

func (l Lines) Less(i, j int) bool {
	return l[i].h < l[j].h
}

func (l Lines) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

// -----------------------------------------------------------------------------
// solution 2
// 解法1中从线段高度入手: 确定线段高度, 找最长距离
// 换个角度, 确定距离后, 再比较高度.
// 用 S[i,j] 表示两条线段 H[i], H[j] 围成的最大容器面积:
//	S[i,j] = |j-i| * MIN(H[i],H[j])
// 假设 H[i] < H[j], 则对于线段 i 有:
//	* 任意与其距离小于等于 |j-i| 的线段 k, 组成的容器面积 S[i,k] <= S[i,j]
// 因为 |k-i|<=|j-i|, 又 MIN(H[i],H[k])<=MIN(H[i],H[j]), 所以 S[i,k]<=S[i,j].
// 这表明, 对于较短线段 i, 无需再检查与其距离在 |j-i| 内的所有其他线段.
// 所以可以从相距最大的两个线段开始, 计算最大面积S:
// 1. 两个游标 i[0,..,j),j[n-1,..,i), 分别代表从左到右, 从右到左滑动.
// 2. H[i,..,j] 包含了当前循环中的所有线段高度, |j-i| 表示当前集合中的最长距离.
// 3. 线段 i,j 组成的容器面积: S[i,j]=|j-i|*MIN(H[i],H[j])
// 4. 最大面积 S=MAX(S,S[i,j])
// 5. 从当前集合中剔除较短线段(无需检查其与剩余所有线段): 若 H[i]<H[j], 则 i++, 反之 j--
// 6. 重复3,4,5即求得最大面积S

func maxArea2(height []int) int {
	n := len(height)
	s := 0
	for i, j := 0, n-1; i < j; i++ {
		if height[i] < height[j] {
			s = max(s, height[i]*(j-i))
			i++
		} else {
			s = max(s, height[j]*(j-i))
			j--
		}
	}
	return s
}

func main() {
	cases := [][]int{
		{1, 8, 6, 2, 5, 4, 8, 3, 7},
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		fmt.Println(maxArea(c))
	}
}
