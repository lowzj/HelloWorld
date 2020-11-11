package main

// https://leetcode-cn.com/problems/create-sorted-array-through-instructions/

// 树状数组

func createSortedArray(ins []int) int {
	const mod = 1000000007
	N := len(ins)
	NN := 200000
	tree := make([]int, NN+1)
	add := func(i int) {
		for i <= NN {
			tree[i]++
			i += i & -i
		}
	}
	sum := func(i int) int {
		ret := 0
		for i > 0 {
			ret += tree[i]
			i -= i & -i
		}
		return ret
	}
	min := func(x, y int) int {
		if x > y {
			return y
		}
		return x
	}
	ans := 0
	for i := 0; i < N; i++ {
		ans += min(sum(ins[i]-1), sum(NN)-sum(ins[i])) % mod
		add(ins[i])
	}
	return ans % mod
}
