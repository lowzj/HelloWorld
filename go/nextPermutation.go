package main

import (
	"fmt"
	"sort"
	"strconv"
)

// https://leetcode-cn.com/problems/next-permutation/

// solution 1
// 思路:
// 下面是演算了很久发现的规律, 下表是4个数字的全排列, 最右边是各排列中下标[0,1,2]对应的逆序对数.
// idx   | 0 1 2 3 | 0 1 2
// ------+---------+--------
// P( 0) | 1 2 3 4 | 0 0 0
// P( 1) | 1 2 4 3 | 0 0 1
// P( 2) | 1 3 2 4 | 0 1 0
// P( 3) | 1 3 4 2 | 0 1 1
// P( 4) | 1 4 2 3 | 0 2 0
// P( 5) | 1 4 3 2 | 0 2 1
// P( 6) | 2 1 3 4 | 1 0 0
// P( 7) | 2 1 4 3 | 1 0 1
// P( 8) | 2 3 1 4 | 1 1 0
// P( 9) | 2 3 4 1 | 1 1 1
// P(10) | 2 4 1 3 | 1 2 0
// P(11) | 2 4 3 1 | 1 2 1
// P(12) | 3 1 2 4 | 2 0 0
// P(13) | 3 1 4 2 | 2 0 1
// P(14) | 3 2 1 4 | 2 1 0
// P(15) | 3 2 4 1 | 2 1 1
// P(16) | 3 4 1 2 | 2 2 0
// P(17) | 3 4 2 1 | 2 2 1
// P(18) | 4 1 2 3 | 3 0 0
// P(19) | 4 1 3 2 | 3 0 1
// P(20) | 4 2 1 3 | 3 1 0
// P(21) | 4 2 3 1 | 3 1 1
// P(22) | 4 3 1 2 | 3 2 0
// P(23) | 4 3 2 1 | 3 2 1
// 首先了解下逆序对: 对于数组A[0,..,n-1], 若i<j有A[i]>A[j], 则(A[i],A[j])为一个逆序对.
// 那么对于 n 个数的字典序排列 P[i], 记该排列的逆序对数为RO[i].
// RO[i]包含n-1个逆序对数, 对应排列P[i]中每位数, 记 P[i][k] 的逆序对数为 RO[i][k].
// 且 RO[i][k] 的进制为 n-k, 即当 RO[i][k] == n-k 时, 需向前一位(第k-1)进1.
// 进位含义:
//		在一个排列中从左至右的第 k 个数的最大逆序对数为 n-k-1, 即右侧数字的个数;
//		当其逆序对数为 n-k 时, 则进位至(k-1)继续计算 RO[i][k-1]+1;
// RO[i][k]含义:
//		除了逆序对数, 还可理解为排列 P[i][k] 在子数列 P[i][k,..,n-1] 升序排列后的下标.
// 因为最后一位 P[i][n-1] 的逆序对数永远为0, 其进制为1, 这里省略.
// 则排列 P[i] 的下一个排列 P[i+1] 的逆序对数为:
// 	RO[i+1] = RO[i] + 1
// 然后按如下步骤取数字形成排列 P[i+1]:
//	1. 拷贝数组 A 为 B, 升序排列数组 B
//	2. k从0至n-1取剩余数列B[0,..,n-1-k]中的第 RO[i+1][k] 小的数字 B[m].
//	3. 则 A[k] = B[m], 即为新排列中的第 k 位数
//	4. 将B[m]从 B[0,..,n-k-1], 形成 B[0, n-1-(k+1)]
//	5. 重复2,3,4即可

func nextPermutation(nums []int) {
	n := len(nums)
	order := make([]int, n)
	copy(order, nums)
	sort.Ints(order)

	ro := make([]int, n)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i] >= nums[j] {
				ro[i]++
			}
		}
	}
	for i := n - 2; i >= 0; i-- {
		ro[i]++
		if ro[i]+i < n {
			break
		}
		ro[i] = 0
	}
	for i := 0; i < n; i++ {
		j := ro[i]
		nums[i] = order[j]
		order = append(order[0:j], order[j+1:]...)
	}
}

// -----------------------------------------------------------------------------
// solution 2
// idx   | 0 1 2 3 4 5 6 7 | 0 1 2 3 4 5 6 7
// ------+-----------------+----------------
// P(10) | 1 2 3 4 6 8 5 7 | 0 0 0 0 1 2 0 0
//                                        +1
// P(11) | 1 2 3 4 6 8 7 5 | 0 0 0 0 1 2 1 0
//                                        +1
// P(12) | 1 2 3 4 7 5 6 8 | 0 0 0 0 2 0 0 0
//                                        +1
// P(13) | 1 2 3 4 7 5 8 6 | 0 0 0 0 2 0 1 0
// 考察上面排列, 参考解法1中的逆序对数的进位含义, P[10]->P[11] 的转换实际上就是逆序对数增1:
//   RO[10]+1 -> RO[11]
// 而RO[11]相对于 RO[10] 变化的位为: RO[10][6,7], 就是 P[11] 相对于 P[10] 变化的位置.
// P[11]->P[12]变化位置[4,5,6,7], 逆序对数同样变: RO[11][4,5,6,7]->R[12][4,5,6,7].
// 增加最后移位逆序数方便理解, 0代表后续数组均为正序, 那么逆序对数增1可形象化理解为:
//	* 从右至左, 找到第一个正序对P[i][j,j+1], 使得P[i][j]<P[i][j+1]
//	* 在P[i][j+1,..,n-1]中找到最小的逆序对P[i][j,k], 其中j<k<=n-1, P[i][j]<P[i][k]
//	* 交换j,k数据, 使得P[i][j]>P[i][k], 同时P[i][j+1,n-1]为全逆序对
//	* 为了RO[i]只增1, 即进位效果, 将 P[i][j+1,n-1] 反转成为全正序对
//	* 如上便得到下一个排列 P[i+1]

func nextPermutation2(nums []int) {
	n := len(nums)
	i := n - 2
	for ; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			for j := n - 1; j > i; j-- {
				if nums[i] < nums[j] {
					nums[i], nums[j] = nums[j], nums[i]
					break
				}
			}
			break
		}
	}
	for j, k := i+1, len(nums)-1; j < k; j, k = j+1, k-1 {
		nums[j], nums[k] = nums[k], nums[j]
	}
}

func main() {
	cases := [][]int{
		{1, 2, 3},
		{1, 3, 2},
		{3, 2, 1},
		{1, 5, 1},
		{1, 2, 3, 9, 4, 7, 10, 8, 5},
	}
	realCase := cases[0:]
	for i, c := range realCase {
		fmt.Println("## case", i)
		fmt.Println("I:", c)
		nextPermutation2(c)
		fmt.Println("O:", c)
	}
	reversedOrder()
}

func reversedOrder() {
	c := []int{8, 7, 6, 5, 4, 3, 2, 1}
	n := len(c)
	toStr := func(a []int) string {
		s := ""
		for i := 0; i < len(a); i++ {
			s = s + " " + strconv.Itoa(a[i])
		}
		return s
	}
	f := func() string {
		count := make([]int, len(c)-1)
		for i := 0; i < n-1; i++ {
			for j := i + 1; j < n; j++ {
				if c[i] > c[j] {
					count[i]++
				}
			}
		}
		return toStr(count)
	}
	for i := 0; i < n*n; i++ {
		nextPermutation2(c)
		fmt.Printf("// P(%2d) |%s |%s\n",
			i, toStr(c), f())
	}
}
