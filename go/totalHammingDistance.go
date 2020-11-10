package main

import "fmt"

//------------------------------------------------------------------------------

// https://leetcode-cn.com/problems/total-hamming-distance/

func totalHammingDistance(nums []int) int {
	return totalHammingDistance2(nums)
}

func totalHammingDistance2(nums []int) int {
	const bitNum = 30
	N := len(nums)
	ans, cnt := 0, 0
	for i := 0; i < bitNum; i++ {
		u := 1 << uint(i)
		for j := 0; j < N; j++ {
			if nums[j]&u != 0 {
				cnt++
			}
		}
		ans += (N - cnt) * cnt
		cnt = 0
	}
	return ans
}

func totalHammingDistance0(nums []int) int {
	const bitNum = 32
	N := len(nums)
	bitCnt := make([]int, bitNum)
	ans := 0
	for i := 0; i < N; i++ {
		for v := uint(nums[i]); v != 0; {
			//j := bits.TrailingZeros(v)
			//bitCnt[j]++
			//v ^= 1<<j
		}
	}
	for j := 0; j < bitNum; j++ {
		ans += bitCnt[j] * (N - bitCnt[j])
	}
	return ans
}

func totalHammingDistance1(nums []int) int {
	ans := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			// 使用 bits.OnesCount 直接统计, 不会超时
			// ans += bits.OnesCount(uint(nums[i]^nums[j]))
			// 下面方法会超时
			for v := nums[i] ^ nums[j]; v > 0; v &= v - 1 {
				ans++
			}
		}
	}
	return ans
}

func main() {
	cases := [][]int{
		{4, 14, 2},
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(totalHammingDistance(c))
	}
}
