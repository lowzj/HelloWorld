package main

import "fmt"

// https://leetcode-cn.com/problems/product-of-array-except-self/

// 设
//    M[i] 是除第i个数其余数的乘积
//    P[i] 是第i个数后续所有数的乘积
//    S[i] 是第i个前序所有数的乘积
// 则 M[i] = P[i] * S[i]
// 所以经过正向反向2轮遍历将 P,S 求出, 再求 M.
// 要做到空间O(1), 不算输出占用的空间, 可以复用输出结果的空间. 先求P, 赋值到M的空间.
// 再反向遍历求S时, 额外一个变量suf记录S[i], 同时计算M[i]=M[i]*suf.

func productExceptSelf(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return []int{}
	}
	res := make([]int, n)
	res[0] = 1
	for i := 0; i < n-1; i++ {
		res[i+1] = res[i] * nums[i]
	}
	suf := 1
	for i := n - 1; i >= 0; i-- {
		res[i] = suf * res[i]
		suf *= nums[i]
	}
	return res
}

func main() {
	cases := [][]int{
		{1, 2, 3, 4},
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(productExceptSelf(c))
	}
}
