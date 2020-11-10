package main

import (
	"fmt"
)

// https://leetcode-cn.com/problems/house-robber/

//------------------------------------------------------------------------------
// 设 f[i] 为从[0..i]号房屋中, 必须偷窃 i 号房屋后所能获得的最大金额.
// 因为只能隔屋偷窃, 所以 f[0], f[1] 就只能偷窃0号/1号房屋.
// 所以 f[i] 为偷窃i号房屋, 再加上 MAX(f[0],f[1],...,f[i-2]).
// 为了节约空间, 可以令:
//     cur = f[i-1], pre = f[i-2], ans = MAX(f[0],f[1],..,f[i-2])
// 则每次循环开始后, i = i+1
//     * pre 要保存f[i-1] 的值, 以供下次循环(i+1)使用
//     * cur 存储 f[i] = ans + nums[i]
//     * ans 重新计算 f[0],f[1],..,..f[i+1-2] 的最大值: ans = MAX(ans, pre)
//     * 循环结束后, cur = f[N-1], ans 为f[N-1]前的最大值, 所以 ans=MAX(anx,cur)

func rob(nums []int) int {
	N := len(nums)
	if N == 0 {
		return 0
	}
	if N == 1 {
		return nums[0]
	}

	cur, pre, ans := nums[1], nums[0], nums[0]
	for i := 2; i < N; i++ {
		pre, cur = cur, ans+nums[i]
		if ans < pre {
			ans = pre
		}
	}
	if ans > cur {
		return ans
	}
	return cur
}

func main() {
	cases := [][]int{
		{1, 2, 3, 1},
	}

	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		// solve
		fmt.Println(rob(c))
	}
}
