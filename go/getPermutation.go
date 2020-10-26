package main

import "fmt"

// https://leetcode-cn.com/problems/permutation-sequence/
// solution
// 基于 nextPermutation 的 solution1
// * 计算k对应的各位置的逆序对数
// * 根据逆序对数得出排列

func getPermutation(n int, k int) string {
	if n == 0 {
		return ""
	}
	if n == 1 {
		return "1"
	}

	fac := []int{1, 1, 2, 6, 24, 120, 720, 5040, 40320, 362880}
	nums := []byte("123456789")

	nums = nums[:n]
	ro := make([]int, n)
	for tmp, i := k-1, 0; i < n; i++ {
		ro[i] = tmp / fac[n-i-1]
		tmp %= fac[n-i-1]
	}
	res := make([]byte, n)
	for i := 0; i < n; i++ {
		res[i] = nums[ro[i]]
		nums = append(nums[:ro[i]], nums[ro[i]+1:]...)
	}
	return string(res)
}

func main() {
	cases := [][]int{
		{3, 3},
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(getPermutation(c[0], c[1]))
	}
}
